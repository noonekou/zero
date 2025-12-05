# 批量插入优化说明

## 概述

将用户角色关系的插入操作从**循环单条插入**优化为**批量插入**,显著提升了性能。

## 修改的文件

### 1. Model 层
**文件**: `common/model/tadminuserrolemodel.go`

#### 新增接口方法
```go
BatchInsert(ctx context.Context, data []*TAdminUserRole) error
```

#### 实现细节
- 使用 PostgreSQL 的多行 INSERT 语法
- 一次性插入所有记录,减少数据库往返次数
- 自动处理空数组情况

**SQL 示例**:
```sql
INSERT INTO "public"."t_admin_user_role" (user_id, role_id, status) 
VALUES ($1, $2, $3), ($4, $5, $6), ($7, $8, $9)
```

### 2. Logic 层

#### 文件 1: `rpc/user/internal/logic/adminuserservice/adduserlogic.go`
**优化位置**: L52-L68

**优化前**:
```go
for _, v := range in.Ids {
    _, err = adminUserRoleModel.Insert(ctx, &model.TAdminUserRole{
        UserId: uid,
        RoleId: v,
    })
    if err != nil {
        return err
    }
}
```

**优化后**:
```go
if len(in.Ids) > 0 {
    userRoles := make([]*model.TAdminUserRole, 0, len(in.Ids))
    for _, roleId := range in.Ids {
        userRoles = append(userRoles, &model.TAdminUserRole{
            UserId: uid,
            RoleId: roleId,
            Status: 1,
        })
    }
    err = adminUserRoleModel.BatchInsert(ctx, userRoles)
    if err != nil {
        return err
    }
}
```

#### 文件 2: `rpc/user/internal/logic/adminuserservice/updateuserlogic.go`
**优化位置**: L67-L81

同样的优化模式应用于更新用户逻辑。

## 性能提升

### 数据库交互次数对比

| 场景 | 优化前 | 优化后 | 提升 |
|------|--------|--------|------|
| 插入 1 个角色 | 1 次 INSERT | 1 次 INSERT | 相同 |
| 插入 5 个角色 | 5 次 INSERT | 1 次 INSERT | **5x** |
| 插入 10 个角色 | 10 次 INSERT | 1 次 INSERT | **10x** |
| 插入 N 个角色 | N 次 INSERT | 1 次 INSERT | **Nx** |

### 性能优势

1. **减少网络往返**: 从 N 次减少到 1 次
2. **降低事务开销**: 单次提交替代多次提交
3. **提高吞吐量**: 批量操作更高效
4. **减少锁竞争**: 减少数据库锁的获取和释放次数

## 代码质量改进

### ✅ 优点

1. **性能优化**: 显著减少数据库交互次数
2. **代码简洁**: 逻辑更清晰,易于维护
3. **事务安全**: 在同一事务中完成,保证原子性
4. **可扩展性**: 支持任意数量的批量插入

### 🔍 注意事项

1. **空数组处理**: 已添加 `len(in.Ids) > 0` 检查
2. **默认状态**: 统一设置 `Status: 1` (启用状态)
3. **错误处理**: 保持原有的错误处理逻辑
4. **事务一致性**: 在事务中使用 `WithSession(session)` 确保一致性

## 测试建议

### 单元测试
```go
func TestBatchInsert(t *testing.T) {
    // 测试空数组
    err := model.BatchInsert(ctx, []*TAdminUserRole{})
    assert.NoError(t, err)
    
    // 测试单条记录
    err = model.BatchInsert(ctx, []*TAdminUserRole{{UserId: 1, RoleId: 1, Status: 1}})
    assert.NoError(t, err)
    
    // 测试多条记录
    roles := []*TAdminUserRole{
        {UserId: 1, RoleId: 1, Status: 1},
        {UserId: 1, RoleId: 2, Status: 1},
        {UserId: 1, RoleId: 3, Status: 1},
    }
    err = model.BatchInsert(ctx, roles)
    assert.NoError(t, err)
}
```

### 性能测试
```bash
# 使用 go benchmark 测试
go test -bench=BenchmarkBatchInsert -benchmem
```

## 最佳实践

### 何时使用批量插入

✅ **适用场景**:
- 需要插入多条相关记录
- 数据量较大(通常 > 3 条)
- 在事务中批量操作

❌ **不适用场景**:
- 单条记录插入
- 需要获取每条记录的插入 ID
- 记录之间有复杂的依赖关系

### 批量大小建议

- **小批量** (< 100): 直接使用批量插入
- **中批量** (100-1000): 考虑分批插入
- **大批量** (> 1000): 建议分批处理,每批 500-1000 条

```go
// 分批插入示例
const batchSize = 500
for i := 0; i < len(data); i += batchSize {
    end := i + batchSize
    if end > len(data) {
        end = len(data)
    }
    err := model.BatchInsert(ctx, data[i:end])
    if err != nil {
        return err
    }
}
```

## 总结

通过实现批量插入功能,我们成功地:
- ✅ 提升了插入性能 (N 倍提升)
- ✅ 减少了数据库负载
- ✅ 保持了代码的可读性和可维护性
- ✅ 确保了事务的原子性和一致性

这是一个典型的性能优化案例,展示了如何通过减少数据库交互次数来提升应用性能。
