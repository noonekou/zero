-- PostgreSQL Function to handle updated_at timestamps
-- This function will be reused by multiple tables.
CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

---

-- 管理端用户表 (t_admin_user)
CREATE TABLE IF NOT EXISTS t_admin_user
(
    id         BIGSERIAL PRIMARY KEY,
    username   VARCHAR(64)                 NOT NULL UNIQUE,
    password   VARCHAR(255)                NOT NULL,
    nickname   VARCHAR(64)                          DEFAULT '',
    avatar     VARCHAR(255)                         DEFAULT '',
    email      VARCHAR(128)                         DEFAULT '',
    phone      VARCHAR(32)                          DEFAULT '',
    status     SMALLINT                    NOT NULL DEFAULT 1,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

-- Trigger for t_admin_user updated_at
CREATE OR REPLACE TRIGGER update_t_admin_user_updated_at
    BEFORE UPDATE
    ON t_admin_user
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
CREATE TABLE IF NOT EXISTS t_role
(
    id         BIGSERIAL PRIMARY KEY,
    name       VARCHAR(20)                 NOT NULL UNIQUE,
    -- status     SMALLINT                    NOT NULL DEFAULT 1,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

ALTER TABLE t_role ADD COLUMN status SMALLINT NOT NULL DEFAULT 1;


-- Trigger for t_role updated_at
CREATE OR REPLACE TRIGGER update_t_role_updated_at
    BEFORE UPDATE
    ON t_role
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Comments for t_role columns
COMMENT ON COLUMN t_role.id IS '角色ID';
COMMENT ON COLUMN t_role.name IS '角色名';
COMMENT ON COLUMN t_role.status IS '状态(1:正常 0:禁用)';
COMMENT ON COLUMN t_role.created_at IS '创建时间';
COMMENT ON COLUMN t_role.updated_at IS '更新时间';
COMMENT ON TABLE t_role IS '角色表';

---

-- 权限表 (t_permission)
CREATE TABLE IF NOT EXISTS t_permission
(
    id            BIGSERIAL PRIMARY KEY,                       -- 权限ID，自增主键
    name          VARCHAR(100)                NOT NULL UNIQUE, -- 权限名称 (例如: 'user:read', 'article:create', 'order:delete')
    resource_name VARCHAR(100)                NOT NULL,        -- 资源名称 (例如: 'user', 'article', 'order')
    action        VARCHAR(50)                 NOT NULL,        -- 操作类型 (例如: 'read', 'create', 'update', 'delete', 'view', 'approve')
    description   TEXT,                                        -- 权限描述
    created_at    TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

-- Trigger for t_permission updated_at
CREATE OR REPLACE TRIGGER update_t_permission_updated_at
    BEFORE UPDATE
    ON t_permission
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Comments for t_permission columns
COMMENT ON COLUMN t_permission.id IS '权限ID';
COMMENT ON COLUMN t_permission.name IS '权限名';
COMMENT ON COLUMN t_permission.resource_name IS '资源名';
COMMENT ON COLUMN t_permission.action IS '操作类型';
COMMENT ON COLUMN t_permission.description IS '权限描述';
COMMENT ON COLUMN t_permission.created_at IS '创建时间';
COMMENT ON COLUMN t_permission.updated_at IS '更新时间';
COMMENT ON TABLE t_permission IS '权限表';

---

-- 用户角色表 (t_user_role)
CREATE TABLE IF NOT EXISTS t_admin_user_role
(
    id         BIGSERIAL PRIMARY KEY,
    user_id    BIGINT                      NOT NULL,
    role_id    BIGINT                      NOT NULL,
    status     SMALLINT                    NOT NULL DEFAULT 1,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    UNIQUE (user_id, role_id) -- `UNIQUE KEY` syntax becomes `UNIQUE` constraint
);

-- Trigger for t_admin_user_role updated_at
CREATE OR REPLACE TRIGGER update_t_admin_user_role_updated_at
    BEFORE UPDATE
    ON t_admin_user_role
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Indexes for t_user_role (explicitly defined after table creation)
CREATE INDEX IF NOT EXISTS idx_admin_user_role_user_id ON t_admin_user_role (user_id);
CREATE INDEX IF NOT EXISTS idx_admin_user_role_role_id ON t_admin_user_role (role_id);

-- Comments for t_user_role columns
COMMENT ON COLUMN t_admin_user_role.id IS '用户角色ID';
COMMENT ON COLUMN t_admin_user_role.user_id IS '用户ID';
COMMENT ON COLUMN t_admin_user_role.role_id IS '角色ID';
COMMENT ON COLUMN t_admin_user_role.status IS '状态(1:正常 0:禁用)';
COMMENT ON COLUMN t_admin_user_role.created_at IS '创建时间';
COMMENT ON COLUMN t_admin_user_role.updated_at IS '更新时间';
COMMENT ON TABLE t_admin_user_role IS '用户角色表';

---

-- 角色权限表 (t_role_permission)
CREATE TABLE IF NOT EXISTS t_role_permission
(
    id              BIGSERIAL PRIMARY KEY,
    role_name       VARCHAR(20)                 NOT NULL,
    permission_name VARCHAR(100)                NOT NULL,
    created_at      TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    UNIQUE (role_name, permission_name) -- `UNIQUE KEY` syntax becomes `UNIQUE` constraint
);

-- Trigger for t_role_permission updated_at
CREATE OR REPLACE TRIGGER update_t_role_permission_updated_at
    BEFORE UPDATE
    ON t_role_permission
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Indexes for t_role_permission (explicitly defined after table creation)
CREATE INDEX IF NOT EXISTS idx_role_permission_role_id ON t_role_permission (role_name);
CREATE INDEX IF NOT EXISTS idx_role_permission_permission_id ON t_role_permission (permission_name);

-- Comments for t_role_permission columns
COMMENT ON COLUMN t_role_permission.id IS '角色权限ID';
COMMENT ON COLUMN t_role_permission.role_name IS '角色名';
COMMENT ON COLUMN t_role_permission.permission_name IS '权限名';
COMMENT ON COLUMN t_role_permission.created_at IS '创建时间';
COMMENT ON COLUMN t_role_permission.updated_at IS '更新时间';
COMMENT ON TABLE t_role_permission IS '角色权限表';

-- 资源表 (t_resource)
CREATE TABLE IF NOT EXISTS t_resource
(
    id          BIGSERIAL PRIMARY KEY,
    code        INT                         NOT NULL UNIQUE,
    name        VARCHAR(20)                 NOT NULL UNIQUE,
    parent_code INT                                  DEFAULT 0,
    description VARCHAR(100)                NOT NULL,
    created_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_resource_parent_code ON t_resource (parent_code);

-- Trigger for t_resource updated_at
CREATE OR REPLACE TRIGGER update_t_resource_updated_at
    BEFORE UPDATE
    ON t_resource
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Comments for t_resource columns
COMMENT ON COLUMN t_resource.id IS '资源ID';
COMMENT ON COLUMN t_resource.code IS '资源编码';
COMMENT ON COLUMN t_resource.name IS '资源名';
COMMENT ON COLUMN t_resource.parent_code IS '父资源编码';
COMMENT ON COLUMN t_resource.description IS '资源描述';
COMMENT ON COLUMN t_resource.created_at IS '创建时间';
COMMENT ON COLUMN t_resource.updated_at IS '更新时间';
COMMENT ON TABLE t_resource IS '资源表';

-- 接口表 (t_apis)
CREATE TABLE IF NOT EXISTS t_apis
(
    id          BIGSERIAL PRIMARY KEY,
    code        INT                         NOT NULL UNIQUE, -- 自定义标识
    method      CHAR(6)                     NOT NULL,        -- 方法名: GET, POST, DELETE, OPTION, PUT
    path        VARCHAR(100)                NOT NULL,        -- 路径
    description TEXT,
    created_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_apis_method_path ON t_apis (method, path);

-- Trigger for t_apis updated_at
CREATE OR REPLACE TRIGGER update_t_apis_updated_at
    BEFORE UPDATE
    ON t_apis
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Comments for t_apis columns
COMMENT ON COLUMN t_apis.id IS '接口ID';
COMMENT ON COLUMN t_apis.code IS '接口编码';
COMMENT ON COLUMN t_apis.method IS '方法名';
COMMENT ON COLUMN t_apis.path IS '路径';
COMMENT ON COLUMN t_apis.description IS '接口描述';
COMMENT ON COLUMN t_apis.created_at IS '创建时间';
COMMENT ON COLUMN t_apis.updated_at IS '更新时间';
COMMENT ON TABLE t_apis IS '接口表';

-- 接口权限表 (t_api_permission)
CREATE TABLE IF NOT EXISTS t_api_permission
(
    id              BIGSERIAL PRIMARY KEY,
    api_code        INT                         NOT NULL, -- 方法名: GET, POST, DELETE, OPTION, PUT
    permission_name VARCHAR(100)                NOT NULL, -- 权限描述: user:read,
    created_at      TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_api_permission_api_id ON t_api_permission (api_code);

-- Trigger for t_api_permission updated_at
CREATE OR REPLACE TRIGGER update_t_api_permission_updated_at
    BEFORE UPDATE
    ON t_api_permission
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Comments for t_api_permission columns
COMMENT ON COLUMN t_api_permission.id IS '接口权限ID';
COMMENT ON COLUMN t_api_permission.api_code IS '接口编码';
COMMENT ON COLUMN t_api_permission.permission_name IS '权限名';
COMMENT ON COLUMN t_api_permission.created_at IS '创建时间';
COMMENT ON COLUMN t_api_permission.updated_at IS '更新时间';
COMMENT ON TABLE t_api_permission IS '接口权限表';
