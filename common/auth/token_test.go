package auth

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

const (
	testSecret = "test-secret-key-for-testing"
	testUserId = int64(12345)
)

func TestGenToken(t *testing.T) {
	token, err := GenerateToken("2yB#@guNbKJDtgys", 3600, 4)
	fmt.Println(token)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

// TestGenerateToken tests the token generation functionality
func TestGenerateToken(t *testing.T) {
	tests := []struct {
		name         string
		accessSecret string
		accessExpire int64
		userId       int64
		wantErr      bool
	}{
		{
			name:         "Valid token generation",
			accessSecret: testSecret,
			accessExpire: 3600, // 1 hour
			userId:       testUserId,
			wantErr:      false,
		},
		{
			name:         "Token with long expiration",
			accessSecret: testSecret,
			accessExpire: 86400, // 24 hours
			userId:       testUserId,
			wantErr:      false,
		},
		{
			name:         "Token with different user ID",
			accessSecret: testSecret,
			accessExpire: 3600,
			userId:       99999,
			wantErr:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := GenerateToken(tt.accessSecret, tt.accessExpire, tt.userId)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, token)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)

				// Verify the token structure
				parsed, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
					return []byte(tt.accessSecret), nil
				})
				assert.NoError(t, err)
				assert.True(t, parsed.Valid)

				// Verify claims
				claims, ok := parsed.Claims.(*CustomClaims)
				assert.True(t, ok)
				assert.Equal(t, tt.userId, claims.UserId)
				assert.Equal(t, "gozero-api", claims.Issuer)
			}
		})
	}
}

// TestValidateToken tests the token validation functionality
func TestValidateToken(t *testing.T) {
	// Generate a valid token for testing
	validToken, err := GenerateToken(testSecret, 3600, testUserId)
	assert.NoError(t, err)

	// Generate an expired token
	expiredToken, err := GenerateToken(testSecret, -10, testUserId) // Already expired
	assert.NoError(t, err)

	tests := []struct {
		name         string
		accessSecret string
		token        string
		wantUserId   int64
		wantErr      bool
	}{
		{
			name:         "Valid token without Bearer prefix",
			accessSecret: testSecret,
			token:        validToken,
			wantUserId:   testUserId,
			wantErr:      false,
		},
		{
			name:         "Valid token with Bearer prefix",
			accessSecret: testSecret,
			token:        "Bearer " + validToken,
			wantUserId:   testUserId,
			wantErr:      false,
		},
		{
			name:         "Invalid token - wrong secret",
			accessSecret: "wrong-secret",
			token:        validToken,
			wantUserId:   0,
			wantErr:      true,
		},
		{
			name:         "Invalid token - malformed token",
			accessSecret: testSecret,
			token:        "invalid.token.string",
			wantUserId:   0,
			wantErr:      true,
		},
		{
			name:         "Invalid token - empty token",
			accessSecret: testSecret,
			token:        "",
			wantUserId:   0,
			wantErr:      true,
		},
		{
			name:         "Invalid token - expired token",
			accessSecret: testSecret,
			token:        expiredToken,
			wantUserId:   0,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userId, err := ValidateToken(tt.accessSecret, tt.token)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, int64(0), userId)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantUserId, userId)
			}
		})
	}
}

// TestTokenExpiration tests token expiration behavior
func TestTokenExpiration(t *testing.T) {
	t.Run("Token should be valid before expiration", func(t *testing.T) {
		token, err := GenerateToken(testSecret, 10, testUserId) // 10 seconds
		assert.NoError(t, err)

		userId, err := ValidateToken(testSecret, token)
		assert.NoError(t, err)
		assert.Equal(t, testUserId, userId)
	})

	t.Run("Token should be invalid after expiration", func(t *testing.T) {
		token, err := GenerateToken(testSecret, 1, testUserId) // 1 second
		assert.NoError(t, err)

		// Wait for token to expire
		time.Sleep(2 * time.Second)

		userId, err := ValidateToken(testSecret, token)
		assert.Error(t, err)
		assert.Equal(t, int64(0), userId)
	})
}

// TestTokenClaims tests the custom claims in the token
func TestTokenClaims(t *testing.T) {
	userId := int64(67890)
	accessExpire := int64(3600)

	token, err := GenerateToken(testSecret, accessExpire, userId)
	assert.NoError(t, err)

	// Parse the token to verify claims
	parsed, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(testSecret), nil
	})
	assert.NoError(t, err)

	claims, ok := parsed.Claims.(*CustomClaims)
	assert.True(t, ok)

	// Verify all claims
	assert.Equal(t, userId, claims.UserId)
	assert.Equal(t, "gozero-api", claims.Issuer)
	assert.NotNil(t, claims.IssuedAt)
	assert.NotNil(t, claims.ExpiresAt)

	// Verify expiration time is approximately correct (within 5 seconds tolerance)
	expectedExpiration := time.Now().Add(time.Duration(accessExpire) * time.Second)
	actualExpiration := claims.ExpiresAt.Time
	timeDiff := actualExpiration.Sub(expectedExpiration).Abs()
	assert.Less(t, timeDiff, 5*time.Second)
}

// TestBearerTokenPrefix tests Bearer prefix handling
func TestBearerTokenPrefix(t *testing.T) {
	token, err := GenerateToken(testSecret, 3600, testUserId)
	assert.NoError(t, err)

	tests := []struct {
		name       string
		token      string
		wantUserId int64
		wantErr    bool
	}{
		{
			name:       "Token without Bearer prefix",
			token:      token,
			wantUserId: testUserId,
			wantErr:    false,
		},
		{
			name:       "Token with Bearer prefix (single space)",
			token:      "Bearer " + token,
			wantUserId: testUserId,
			wantErr:    false,
		},
		{
			name:       "Token with bearer prefix (lowercase)",
			token:      "bearer " + token,
			wantUserId: 0,
			wantErr:    true, // Case sensitive, should fail
		},
		{
			name:       "Token with Bearer prefix (multiple spaces)",
			token:      "Bearer  " + token,
			wantUserId: 0,
			wantErr:    true, // Extra space should fail
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userId, err := ValidateToken(testSecret, tt.token)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantUserId, userId)
			}
		})
	}
}

// TestMultipleTokensForSameUser tests generating multiple tokens for the same user
func TestMultipleTokensForSameUser(t *testing.T) {
	token1, err1 := GenerateToken(testSecret, 3600, testUserId)
	assert.NoError(t, err1)

	// Generate another token for the same user
	time.Sleep(1 * time.Second) // Ensure different IssuedAt time
	token2, err2 := GenerateToken(testSecret, 3600, testUserId)
	assert.NoError(t, err2)

	// Tokens should be different
	assert.NotEqual(t, token1, token2)

	// Both tokens should be valid
	userId1, err := ValidateToken(testSecret, token1)
	assert.NoError(t, err)
	assert.Equal(t, testUserId, userId1)

	userId2, err := ValidateToken(testSecret, token2)
	assert.NoError(t, err)
	assert.Equal(t, testUserId, userId2)
}

// BenchmarkGenerateToken benchmarks token generation
func BenchmarkGenerateToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateToken(testSecret, 3600, testUserId)
	}
}

// BenchmarkValidateToken benchmarks token validation
func BenchmarkValidateToken(b *testing.B) {
	token, _ := GenerateToken(testSecret, 3600, testUserId)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ValidateToken(testSecret, token)
	}
}
