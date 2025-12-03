DELETE FROM t_role;
DELETE FROM t_role_permission;
DELETE FROM t_admin_user_role;
DELETE FROM t_admin_user;
DELETE FROM t_api_permission;
DELETE FROM t_apis;
DELETE FROM t_permission;
DELETE FROM t_resource;

-- 资源表
INSERT INTO t_resource (code, name, parent_code, description)
VALUES (1000000, 'dashboard', 0, '仪表盘');
INSERT INTO t_resource (code, name, parent_code, description)VALUES (2000000, 'system', 0, '权限');
INSERT INTO t_resource (code, name, parent_code, description)
VALUES (2000001, 'system:employee', 2000000, '员工');
INSERT INTO t_resource (code, name, parent_code, description)
VALUES (2000002, 'system:role', 2000000, '角色');

-- 权限表

-- 仪表盘
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('dashboard:read', 'dashboard', 'read', '仪表盘');

-- 员工
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('system:employee:read', 'system:employee', 'read', '员工-列表');
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('system:employee:create', 'system:employee', 'create', '员工-新建');
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('system:employee:update', 'system:employee', 'update', '员工-更新');
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('system:employee:delete', 'system:employee', 'delete', '员工-删除');
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('system:employee:view', 'system:employee', 'view', '员工-详情');
-- 角色
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('system:role:read', 'system:role', 'read', '角色-列表');
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('system:role:create', 'system:role', 'create', '角色-新建');
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('system:role:update', 'system:role', 'update', '角色-更新');
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('system:role:delete', 'system:role', 'delete', '角色-删除');
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('system:role:view', 'system:role', 'view', '角色-详情');


-- API 表
INSERT INTO t_apis (code, method, path, description)
VALUES (100000, 'POST', '/v1/auth/login', '员工登陆');
INSERT INTO t_apis (code, method, path, description)
VALUES (100001, 'POST', '/v1/auth/register', '员工注册');
INSERT INTO t_apis (code, method, path, description)
VALUES (100002, 'POST', '/v1/auth/logout', '退出登录');

INSERT INTO t_apis (code, method, path, description)
VALUES (300001, 'GET', '/v1/permission/list', '权限列表');
INSERT INTO t_apis (code, method, path, description)
VALUES (300002, 'POST', '/v1/permission/role/add', '添加角色');
INSERT INTO t_apis (code, method, path, description)
VALUES (300003, 'POST', '/v1/permission/role/update', '更新角色');
INSERT INTO t_apis (code, method, path, description)
VALUES (300004, 'GET', '/v1/permission/role/list', '角色列表');
INSERT INTO t_apis (code, method, path, description)
VALUES (300005, 'GET', '/v1/permission/role/info', '获取角色信息');
INSERT INTO t_apis (code, method, path, description)
VALUES (300006, 'DELETE', '/v1/permission/role/delete', '删除角色');

INSERT INTO t_apis (code, method, path, description)
VALUES (200001, 'GET', '/v1/user/info', '获取用户信息');
INSERT INTO t_apis (code, method, path, description)
VALUES (200002, 'GET', '/v1/user/list', '获取用户列表');
INSERT INTO t_apis (code, method, path, description)
VALUES (200003, 'POST', '/v1/user/add', '添加用户');
INSERT INTO t_apis (code, method, path, description)
VALUES (200004, 'POST', '/v1/user/update', '更新用户');
INSERT INTO t_apis (code, method, path, description)
VALUES (200005, 'DELETE', '/v1/user/delete', '删除用户');

-- API 权限表
INSERT INTO t_api_permission (api_code, permission_name)
VALUES (100000, '*');
INSERT INTO t_api_permission (api_code, permission_name)
VALUES (100001, '*');
INSERT INTO t_api_permission (api_code, permission_name) VALUES (100002, '*');

INSERT INTO t_api_permission (api_code, permission_name) VALUES (300001, '*');
INSERT INTO t_api_permission (api_code, permission_name) VALUES (300002, 'system:role:create');
INSERT INTO t_api_permission (api_code, permission_name) VALUES (300003, 'system:role:update');
INSERT INTO t_api_permission (api_code, permission_name) VALUES (300004, 'system:role:read');
INSERT INTO t_api_permission (api_code, permission_name) VALUES (300005, 'system:role:view');
INSERT INTO t_api_permission (api_code, permission_name) VALUES (300006, 'system:role:delete');

INSERT INTO t_api_permission (api_code, permission_name) VALUES (200001, '*');
INSERT INTO t_api_permission (api_code, permission_name) VALUES (200002, 'system:role:read');
INSERT INTO t_api_permission (api_code, permission_name) VALUES (200003, 'system:role:create');
INSERT INTO t_api_permission (api_code, permission_name) VALUES (200004, 'system:role:update');
INSERT INTO t_api_permission (api_code, permission_name) VALUES (200005, 'system:role:delete');

-- 角色表
INSERT INTO t_role (name)
VALUES ('admin');

-- 仪表盘
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'dashboard:read');

-- 员工
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'system:employee:read');
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'system:employee:create');
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'system:employee:update');
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'system:employee:delete');
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'system:employee:view');

-- 角色
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'system:role:read');
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'system:role:create');
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'system:role:update');
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'system:role:delete');
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'system:role:view');

-- 用户表
-- pwd: (5pJsw1G9r9[f)
INSERT INTO t_admin_user (username, nickname, avatar, email, phone, password, status)
VALUES ('admin', 'admin', 'https://avatars.githubusercontent.com/u/785674?v=4', 'admin@example.com', '123456', '21223e1706c109dca4af2c7b1f2fff69', 1);

-- 用户角色表
INSERT INTO t_admin_user_role (user_id, role_id)
SELECT (SELECT id FROM t_admin_user ORDER BY created_at LIMIT 1) as uid,
       (SELECT id FROM t_role ORDER BY created_at LIMIT 1) as rid;;

