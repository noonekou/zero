-- PostgreSQL Function to handle updated_at timestamps
-- This function will be reused by multiple tables.
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

---

-- 管理端用户表 (t_admin_user)
CREATE TABLE IF NOT EXISTS t_admin_user (
  id BIGSERIAL PRIMARY KEY,
  username VARCHAR(64) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  nickname VARCHAR(64) DEFAULT '',
  avatar VARCHAR(255) DEFAULT '',
  email VARCHAR(128) DEFAULT '',
  phone VARCHAR(32) DEFAULT '',
  status SMALLINT NOT NULL DEFAULT 1,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

-- Trigger for t_admin_user updated_at
CREATE OR REPLACE TRIGGER update_t_admin_user_updated_at
BEFORE UPDATE ON t_admin_user
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Comments for t_admin_user columns
COMMENT ON COLUMN t_admin_user.id IS '用户ID';
COMMENT ON COLUMN t_admin_user.username IS '用户名';
COMMENT ON COLUMN t_admin_user.password IS '密码hash';
COMMENT ON COLUMN t_admin_user.nickname IS '昵称';
COMMENT ON COLUMN t_admin_user.avatar IS '头像';
COMMENT ON COLUMN t_admin_user.email IS '邮箱';
COMMENT ON COLUMN t_admin_user.phone IS '手机号';
COMMENT ON COLUMN t_admin_user.status IS '状态(1:正常 0:禁用)';
COMMENT ON COLUMN t_admin_user.created_at IS '创建时间';
COMMENT ON COLUMN t_admin_user.updated_at IS '更新时间';
COMMENT ON TABLE t_admin_user IS '管理端用户表';

---

-- 角色表 (t_role)
CREATE TABLE IF NOT EXISTS t_role (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL UNIQUE,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

-- Trigger for t_role updated_at
CREATE OR REPLACE TRIGGER update_t_role_updated_at
BEFORE UPDATE ON t_role
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Comments for t_role columns
COMMENT ON COLUMN t_role.id IS '角色ID';
COMMENT ON COLUMN t_role.name IS '角色名';
COMMENT ON COLUMN t_role.created_at IS '创建时间';
COMMENT ON COLUMN t_role.updated_at IS '更新时间';
COMMENT ON TABLE t_role IS '角色表';

---

-- 权限表 (t_permission)
CREATE TABLE IF NOT EXISTS t_permission (
  id BIGSERIAL PRIMARY KEY,
  code VARCHAR(128) NOT NULL UNIQUE,
  name VARCHAR(255) NOT NULL UNIQUE,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

-- Trigger for t_permission updated_at
CREATE OR REPLACE TRIGGER update_t_permission_updated_at
BEFORE UPDATE ON t_permission
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Comments for t_permission columns
COMMENT ON COLUMN t_permission.id IS '权限ID';
COMMENT ON COLUMN t_permission.code IS '权限编码(如：sys:user:list， sys:user:add)';
COMMENT ON COLUMN t_permission.name IS '权限名';
COMMENT ON COLUMN t_permission.created_at IS '创建时间';
COMMENT ON COLUMN t_permission.updated_at IS '更新时间';
COMMENT ON TABLE t_permission IS '权限表';

---

-- 用户角色表 (t_user_role)
CREATE TABLE IF NOT EXISTS t_user_role (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL,
  role_id BIGINT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
  UNIQUE (user_id, role_id) -- `UNIQUE KEY` syntax becomes `UNIQUE` constraint
);

-- Trigger for t_user_role updated_at
CREATE OR REPLACE TRIGGER update_t_user_role_updated_at
BEFORE UPDATE ON t_user_role
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Indexes for t_user_role (explicitly defined after table creation)
CREATE INDEX IF NOT EXISTS idx_user_role_user_id ON t_user_role (user_id);
CREATE INDEX IF NOT EXISTS idx_user_role_role_id ON t_user_role (role_id);

-- Comments for t_user_role columns
COMMENT ON COLUMN t_user_role.id IS '用户角色ID';
COMMENT ON COLUMN t_user_role.user_id IS '用户ID';
COMMENT ON COLUMN t_user_role.role_id IS '角色ID';
COMMENT ON COLUMN t_user_role.created_at IS '创建时间';
COMMENT ON COLUMN t_user_role.updated_at IS '更新时间';
COMMENT ON TABLE t_user_role IS '用户角色表';

---

-- 角色权限表 (t_role_permission)
CREATE TABLE IF NOT EXISTS t_role_permission (
  id BIGSERIAL PRIMARY KEY,
  role_id BIGINT NOT NULL,
  permission_id BIGINT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
  UNIQUE (role_id, permission_id) -- `UNIQUE KEY` syntax becomes `UNIQUE` constraint
);

-- Trigger for t_role_permission updated_at
CREATE OR REPLACE TRIGGER update_t_role_permission_updated_at
BEFORE UPDATE ON t_role_permission
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Indexes for t_role_permission (explicitly defined after table creation)
CREATE INDEX IF NOT EXISTS idx_role_permission_role_id ON t_role_permission (role_id);
CREATE INDEX IF NOT EXISTS idx_role_permission_permission_id ON t_role_permission (permission_id);

-- Comments for t_role_permission columns
COMMENT ON COLUMN t_role_permission.id IS '角色权限ID';
COMMENT ON COLUMN t_role_permission.role_id IS '角色ID';
COMMENT ON COLUMN t_role_permission.permission_id IS '权限ID';
COMMENT ON COLUMN t_role_permission.created_at IS '创建时间';
COMMENT ON COLUMN t_role_permission.updated_at IS '更新时间';
COMMENT ON TABLE t_role_permission IS '角色权限表';