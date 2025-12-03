# Token 生成与校验测试文档

## 概述

本测试套件为 `common/auth/token.go` 中的 JWT token 生成和校验功能提供全面的测试覆盖。

## 测试文件

- **文件路径**: `common/auth/token_test.go`
- **测试覆盖率**: 84.2%
- **测试用例数**: 6 个主要测试函数 + 2 个性能基准测试

## 测试用例详情

### 1. `TestGenerateToken` - Token 生成测试

测试 `GenerateToken` 函数的各种场景：

- ✅ 正常 token 生成
- ✅ 长过期时间的 token（24小时）
- ✅ 不同用户ID的 token 生成

**验证内容**:
- Token 字符串非空
- JWT 结构正确
- Claims 包含正确的用户ID
- Issuer 为 "gozero-api"

### 2. `TestValidateToken` - Token 校验测试

测试 `ValidateToken` 函数的各种场景：

- ✅ 有效 token（无 Bearer 前缀）
- ✅ 有效 token（带 Bearer 前缀）
- ❌ 错误的密钥
- ❌ 格式错误的 token
- ❌ 空 token
- ❌ 过期的 token

**验证内容**:
- 正确解析用户ID
- 正确处理 Bearer 前缀
- 适当的错误处理

### 3. `TestTokenExpiration` - Token 过期测试

测试 token 过期行为：

- ✅ Token 在过期前应该有效
- ✅ Token 在过期后应该无效

**测试方法**:
- 生成短期有效的 token（1秒）
- 等待过期后验证失败

### 4. `TestTokenClaims` - Token Claims 测试

验证 JWT Claims 的正确性：

- ✅ UserId 正确
- ✅ Issuer 正确
- ✅ IssuedAt 时间存在
- ✅ ExpiresAt 时间正确（误差<5秒）

### 5. `TestBearerTokenPrefix` - Bearer 前缀处理测试

测试各种前缀格式：

- ✅ 无前缀
- ✅ 标准 "Bearer " 前缀（单空格）
- ❌ 小写 "bearer " 前缀（大小写敏感）
- ❌ "Bearer  " 多空格前缀

### 6. `TestMultipleTokensForSameUser` - 多 Token 测试

验证同一用户可以有多个有效 token：

- ✅ 同一用户的多个 token 应该不同
- ✅ 所有 token 都应该有效

## 性能基准测试

### `BenchmarkGenerateToken`

测试 token 生成的性能：

```
BenchmarkGenerateToken-12    776431    1527 ns/op    2298 B/op    36 allocs/op
```

- **速度**: ~1.5 微秒/操作
- **内存**: 2298 字节/操作
- **分配次数**: 36 次/操作

### `BenchmarkValidateToken`

测试 token 校验的性能：

```
BenchmarkValidateToken-12    481900    2571 ns/op    2800 B/op    50 allocs/op
```

- **速度**: ~2.6 微秒/操作
- **内存**: 2800 字节/操作
- **分配次数**: 50 次/操作

## 运行测试

### 运行所有测试

```bash
go test -v ./common/auth/
```

### 运行测试并查看覆盖率

```bash
go test -cover ./common/auth/
```

### 生成详细覆盖率报告

```bash
go test -coverprofile=coverage.out ./common/auth/
go tool cover -html=coverage.out
```

### 运行性能基准测试

```bash
go test -bench=. -benchmem ./common/auth/
```

### 运行特定测试

```bash
# 只运行 token 生成测试
go test -v -run TestGenerateToken ./common/auth/

# 只运行 token 校验测试
go test -v -run TestValidateToken ./common/auth/
```

## 测试结果示例

```
=== RUN   TestGenerateToken
=== RUN   TestGenerateToken/Valid_token_generation
=== RUN   TestGenerateToken/Token_with_long_expiration
=== RUN   TestGenerateToken/Token_with_different_user_ID
--- PASS: TestGenerateToken (0.00s)

=== RUN   TestValidateToken
=== RUN   TestValidateToken/Valid_token_without_Bearer_prefix
=== RUN   TestValidateToken/Valid_token_with_Bearer_prefix
=== RUN   TestValidateToken/Invalid_token_-_wrong_secret
=== RUN   TestValidateToken/Invalid_token_-_malformed_token
=== RUN   TestValidateToken/Invalid_token_-_empty_token
=== RUN   TestValidateToken/Invalid_token_-_expired_token
--- PASS: TestValidateToken (0.00s)

=== RUN   TestTokenExpiration
--- PASS: TestTokenExpiration (2.00s)

=== RUN   TestTokenClaims
--- PASS: TestTokenClaims (0.00s)

=== RUN   TestBearerTokenPrefix
--- PASS: TestBearerTokenPrefix (0.00s)

=== RUN   TestMultipleTokensForSameUser
--- PASS: TestMultipleTokensForSameUser (1.00s)

PASS
ok      bookstore/common/auth   3.371s
```

## 依赖项

测试使用以下依赖：

- `github.com/stretchr/testify/assert` - 断言库
- `github.com/golang-jwt/jwt/v5` - JWT 库（被测试代码使用）

## 注意事项

1. **时间敏感测试**: `TestTokenExpiration` 和 `TestMultipleTokensForSameUser` 包含 `time.Sleep`，总测试时间约 3 秒
2. **大小写敏感**: Bearer 前缀是大小写敏感的，只有 "Bearer " 格式有效
3. **空格敏感**: "Bearer " 和 token 之间必须是单个空格
4. **Token 唯一性**: 即使是同一用户，由于 IssuedAt 时间不同，每次生成的 token 都是唯一的

## 覆盖率分析

```
GenerateToken   75.0%  - 主要未覆盖错误日志路径
ValidateToken   90.9%  - 高覆盖率，主要场景都已测试
总体覆盖率      84.2%  - 良好的测试覆盖
```

## 未来改进建议

1. 可以添加并发测试，验证多个 goroutine 同时生成/验证 token
2. 可以添加更多边界条件测试（如非常长的过期时间）
3. 可以添加安全性测试（如测试 token 无法被伪造）
