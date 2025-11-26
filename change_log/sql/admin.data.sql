-- 资源表
INSERT INTO t_resource (code, name, parent_code, description)
VALUES (1000000, 'dashboard', 0, '仪表盘');
INSERT INTO t_resource (code, name, parent_code, description)
VALUES (2000000, 'employee', 0, '员工');
INSERT INTO t_resource (code, name, parent_code, description)
VALUES (2000001, 'employee:employee', 2000000, '员工');
INSERT INTO t_resource (code, name, parent_code, description)
VALUES (2000002, 'employee:role', 2000000, '角色');


-- 权限表

-- 仪表盘
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('dashboard:read', 'dashboard', 'read', '仪表盘-列表');

-- 员工
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('employee:employee:read', 'employee:employee', 'read', '员工-列表');
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('employee:employee:create', 'employee:employee', 'create', '员工-新建');
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('employee:employee:update', 'employee:employee', 'update', '员工-更新');
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('employee:employee:delete', 'employee:employee', 'delete', '员工-删除');
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('employee:employee:view', 'employee:employee', 'view', '员工-详情');
-- 角色
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('employee:role:read', 'employee:role', 'read', '角色-列表');
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('employee:role:create', 'employee:role', 'create', '角色-新建');
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('employee:role:update', 'employee:role', 'update', '角色-更新');
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('employee:role:delete', 'employee:role', 'delete', '角色-删除');
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('employee:role:view', 'employee:role', 'view', '角色-详情');


-- API 表
INSERT INTO t_apis (code, method, path, description)
VALUES (100000, 'POST', '/v1/auth/login', '员工登陆');
INSERT INTO t_apis (code, method, path, description)
VALUES (100001, 'POST', '/v1/auth/register', '员工注册');

INSERT INTO t_apis (code, method, path, description)
VALUES (100002, 'POST', '/v1/auth/permission/list', '权限列表');
INSERT INTO t_apis (code, method, path, description)
VALUES (100003, 'POST', '/v1/auth/role/add', '添加角色');
INSERT INTO t_apis (code, method, path, description)
VALUES (100004, 'POST', '/v1/auth/role/update', '更新角色');
INSERT INTO t_apis (code, method, path, description)
VALUES (100005, 'GET', '/v1/auth/role/list', '角色列表');
INSERT INTO t_apis (code, method, path, description)
VALUES (100006, 'GET', '/v1/auth/role/info', '获取角色信息');
INSERT INTO t_apis (code, method, path, description)
VALUES (100007, 'DELETE', '/v1/auth/role/delete', '删除角色');

INSERT INTO t_apis (code, method, path, description)
VALUES (200001, 'GET', '/v1/user/info', '获取用户信息');
INSERT INTO t_apis (code, method, path, description)
VALUES (200002, 'GET', '/v1/user/list', '获取用户列表');

-- API 权限表
INSERT INTO t_api_permission (api_code, permission_name)
VALUES (100000, '*');
INSERT INTO t_api_permission (api_code, permission_name)
VALUES (100001, '*');

INSERT INTO t_api_permission (api_code, permission_name) VALUES (100002, '*');
INSERT INTO t_api_permission (api_code, permission_name) VALUES (100003, 'employee:role:create');
INSERT INTO t_api_permission (api_code, permission_name) VALUES (100004, 'employee:role:update');
INSERT INTO t_api_permission (api_code, permission_name) VALUES (100005, 'employee:role:read');
INSERT INTO t_api_permission (api_code, permission_name) VALUES (100006, 'employee:role:view');
INSERT INTO t_api_permission (api_code, permission_name) VALUES (100007, 'employee:role:delete');

INSERT INTO t_api_permission (api_code, permission_name) VALUES (200001, 'employee:role:view');
INSERT INTO t_api_permission (api_code, permission_name) VALUES (200002, 'employee:role:read');

-- 角色表
INSERT INTO t_role (name)
VALUES ('admin');

-- 仪表盘
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'dashboard:read');

-- 员工
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'employee:employee:read');
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'employee:employee:create');
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'employee:employee:update');
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'employee:employee:delete');
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'employee:employee:view');

-- 角色
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'employee:role:read');
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'employee:role:create');
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'employee:role:update');
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'employee:role:delete');
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'employee:role:view');

-- 用户表
-- pwd: (5pJsw1G9r9[f)
INSERT INTO t_admin_user (username, nickname, avatar, email, phone, password, status)
VALUES ('admin', 'admin', 'https://avatars.githubusercontent.com/u/785674?v=4', 'admin@example.com', '123456', '21223e1706c109dca4af2c7b1f2fff69', 1);

-- 用户角色表
INSERT INTO t_admin_user_role (user_id, role_id)
SELECT (SELECT id FROM t_admin_user ORDER BY created_at LIMIT 1) as uid,
       (SELECT id FROM t_role ORDER BY created_at LIMIT 1) as rid;;

