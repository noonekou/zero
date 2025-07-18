-- 资源表
INSERT INTO t_resource (code, name, parent_code, description)
VALUES (1000000, 'statistic', 0, '数据统计');
INSERT INTO t_resource (code, name, parent_code, description)
VALUES (1000001, 'statistic:user', 1000000, '用户数据');
INSERT INTO t_resource (code, name, parent_code, description)
VALUES (2000000, 'operation', 0, '运营中心');
INSERT INTO t_resource (code, name, parent_code, description)
VALUES (2000001, 'operation:user', 2000000, '用户列表');
INSERT INTO t_resource (code, name, parent_code, description)
VALUES (3000000, 'employee', 0, '员工管理');
INSERT INTO t_resource (code, name, parent_code, description)
VALUES (3000001, 'employee:employee', 3000000, '员工');
INSERT INTO t_resource (code, name, parent_code, description)
VALUES (3000002, 'employee:role', 3000000, '角色');

-- 权限表

-- 用户数据
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('statistic:user:read', 'statistic:user', 'read', '用户数据-列表');
-- 用户列表
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('operation:user:read', 'operation:user', 'read', '用户列表-列表');
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('operation:user:create', 'operation:user', 'create', '用户列表-新建');
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('operation:user:update', 'operation:user', 'update', '用户列表-更新');
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('operation:user:delete', 'operation:user', 'delete', '用户列表-删除');
INSERT INTO t_permission (name, resource_name, action, description)
VALUES ('operation:user:view', 'operation:user', 'view', '用户列表-详情');
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
VALUES (100002, 'GET', '/v1/user/info', '用户列表-详情');
INSERT INTO t_apis (code, method, path, description)
VALUES (100003, 'GET', '/v1/user/list', '用户列表-列表');

-- API 权限表
INSERT INTO t_api_permission (api_code, permission_name)
VALUES (100000, '*');
INSERT INTO t_api_permission (api_code, permission_name)
VALUES (100001, '*');

INSERT INTO t_api_permission (api_code, permission_name)
VALUES (100002, 'operation:user:view');
INSERT INTO t_api_permission (api_code, permission_name)
VALUES (100003, 'operation:user:read');

-- 角色表
INSERT INTO t_role (name)
VALUES ('admin');

-- 角色权限表
-- 用户数据
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'statistic:user:read');

-- 用户列表
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'operation:user:read');
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'operation:user:create');
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'operation:user:update');
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'operation:user:delete');
INSERT INTO t_role_permission (role_name, permission_name)
VALUES ('admin', 'operation:user:view');

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


