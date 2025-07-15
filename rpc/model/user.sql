-- PostgreSQL Function to handle updated_at timestamps
-- This function can be reused across multiple tables.
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

---

-- 用户表 (t_user)
CREATE TABLE IF NOT EXISTS t_user (
  id BIGSERIAL PRIMARY KEY, -- MySQL的BIGINT AUTO_INCREMENT 对应 PostgreSQL的 BIGSERIAL
  username VARCHAR(64) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  nickname VARCHAR(64) DEFAULT '',
  avatar VARCHAR(255) DEFAULT '',
  email VARCHAR(128) DEFAULT '',
  phone VARCHAR(32) DEFAULT '',
  status SMALLINT NOT NULL DEFAULT 1,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(), -- MySQL的TIMESTAMP 对应 PostgreSQL的 TIMESTAMP WITHOUT TIME ZONE
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW() -- MySQL的TIMESTAMP 对应 PostgreSQL的 TIMESTAMP WITHOUT TIME ZONE
);

-- Trigger for t_user updated_at
-- PostgreSQL 没有像 MySQL 那样的 ON UPDATE CURRENT_TIMESTAMP 语法。
-- 你需要创建一个触发器来实现自动更新 updated_at 字段。
CREATE OR REPLACE TRIGGER update_t_user_updated_at
BEFORE UPDATE ON t_user
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Comments for t_user columns (PostgreSQL 推荐的注释方式)
COMMENT ON COLUMN t_user.id IS '用户ID';
COMMENT ON COLUMN t_user.username IS '用户名';
COMMENT ON COLUMN t_user.password IS '密码hash';
COMMENT ON COLUMN t_user.nickname IS '昵称';
COMMENT ON COLUMN t_user.avatar IS '头像';
COMMENT ON COLUMN t_user.email IS '邮箱';
COMMENT ON COLUMN t_user.phone IS '手机号';
COMMENT ON COLUMN t_user.status IS '状态(1:正常 0:禁用)';
COMMENT ON COLUMN t_user.created_at IS '创建时间';
COMMENT ON COLUMN t_user.updated_at IS '更新时间';
COMMENT ON TABLE t_user IS '用户表';