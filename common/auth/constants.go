package auth

import "fmt"

const (
	// TokenKeyPrefix Redis 中存储 token 的 key 前缀
	TokenKeyPrefix = "token:user:"
)

// GetTokenKey 根据 userId 生成 token 在 Redis 中的 key
func GetTokenKey(userId int64) string {
	return fmt.Sprintf("%s%d", TokenKeyPrefix, userId)
}
