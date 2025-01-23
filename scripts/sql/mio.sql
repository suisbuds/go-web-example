-- ----------------------------
-- 创建新数据库: \i  /path/mio.sql
-- ----------------------------
CREATE DATABASE mio
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.UTF-8'
    LC_CTYPE = 'en_US.UTF-8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    TEMPLATE = template0;



\c mio

-- ----------------------------
-- 创建 model 表
-- ----------------------------
DROP TABLE IF EXISTS model;
CREATE TABLE model (
    id SERIAL PRIMARY KEY,
    created_by VARCHAR(100) NOT NULL DEFAULT '',
    modified_by VARCHAR(100) NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    state SMALLINT NOT NULL DEFAULT 1      -- 0禁用、1启用
);

COMMENT ON TABLE model IS '公共字段';

COMMENT ON COLUMN model.id IS '主键 ID';
COMMENT ON COLUMN model.created_by IS '创建人';
COMMENT ON COLUMN model.modified_by IS '修改人';
COMMENT ON COLUMN model.created_at IS '创建时间';
COMMENT ON COLUMN model.updated_at IS '修改时间';
COMMENT ON COLUMN model.deleted_at IS '删除时间';
COMMENT ON COLUMN model.state IS '状态';


-- ----------------------------
-- 创建 mio_user 表
-- ----------------------------
DROP TABLE IF EXISTS mio_user;
CREATE TABLE mio_user (
    id SERIAL PRIMARY KEY,
    created_by VARCHAR(100) NOT NULL DEFAULT '',
    modified_by VARCHAR(100) NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    state SMALLINT NOT NULL DEFAULT 1, 
    
    username VARCHAR(50) NOT NULL DEFAULT '',
    password VARCHAR(50) DEFAULT '',
    avatar VARCHAR(255) DEFAULT 'https://zbj-bucket1.oss-cn-shenzhen.aliyuncs.com/avatar.JPG',
    user_type SMALLINT NOT NULL DEFAULT 2 -- 1管理员、2为普通用户
);

COMMENT ON TABLE mio_user IS '用户管理';

COMMENT ON COLUMN mio_user.id IS '主键 ID';
COMMENT ON COLUMN mio_user.created_by IS '创建人';
COMMENT ON COLUMN mio_user.modified_by IS '修改人';
COMMENT ON COLUMN mio_user.created_at IS '创建时间';
COMMENT ON COLUMN mio_user.updated_at IS '修改时间';
COMMENT ON COLUMN mio_user.deleted_at IS '删除时间';
COMMENT ON COLUMN mio_user.state IS '状态';
COMMENT ON COLUMN mio_user.username IS '账号';
COMMENT ON COLUMN mio_user.password IS '密码';
COMMENT ON COLUMN mio_user.avatar IS '头像地址';
COMMENT ON COLUMN mio_user.user_type IS '用户类型';

-- ----------------------------
-- 创建 mio_tag 表
-- ----------------------------
DROP TABLE IF EXISTS mio_tag;
CREATE TABLE mio_tag (
    id SERIAL PRIMARY KEY,
    created_by VARCHAR(100) NOT NULL DEFAULT '',
    modified_by VARCHAR(100) NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    state SMALLINT NOT NULL DEFAULT 1, 
    
    name VARCHAR(100) DEFAULT ''
);

COMMENT ON TABLE mio_tag IS '标签管理';

COMMENT ON COLUMN mio_tag.id IS '主键 ID';
COMMENT ON COLUMN mio_tag.created_by IS '创建人';
COMMENT ON COLUMN mio_tag.modified_by IS '修改人';
COMMENT ON COLUMN mio_tag.created_at IS '创建时间';
COMMENT ON COLUMN mio_tag.updated_at IS '修改时间';
COMMENT ON COLUMN mio_tag.deleted_at IS '删除时间';
COMMENT ON COLUMN mio_tag.state IS '状态';
COMMENT ON COLUMN mio_tag.name IS '标签名称';

-- ----------------------------
-- 创建 mio_article 表
-- ----------------------------
DROP TABLE IF EXISTS mio_article;
CREATE TABLE mio_article (
    id SERIAL PRIMARY KEY,
    created_by VARCHAR(100) NOT NULL DEFAULT '',
    modified_by VARCHAR(100) NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    state SMALLINT NOT NULL DEFAULT 1, -- 0草稿、1发布、2删除
    
    tag_id INT DEFAULT 0,
    title VARCHAR(100) DEFAULT '',
    description VARCHAR(255) DEFAULT '',
    content TEXT,
    cover_image_url VARCHAR(255) DEFAULT ''
);

COMMENT ON TABLE mio_article IS '文章管理';

COMMENT ON COLUMN mio_article.id IS '主键 ID';
COMMENT ON COLUMN mio_article.created_by IS '创建人';
COMMENT ON COLUMN mio_article.modified_by IS '修改人';
COMMENT ON COLUMN mio_article.created_at IS '创建时间';
COMMENT ON COLUMN mio_article.updated_at IS '修改时间';
COMMENT ON COLUMN mio_article.deleted_at IS '删除时间';
COMMENT ON COLUMN mio_article.state IS '状态';
COMMENT ON COLUMN mio_article.tag_id IS '标签 ID';
COMMENT ON COLUMN mio_article.title IS '文章标题';
COMMENT ON COLUMN mio_article.description IS '描述';
COMMENT ON COLUMN mio_article.content IS '内容';
COMMENT ON COLUMN mio_article.cover_image_url IS '封面图片地址';

-- ----------------------------
-- 创建 mio_role 表
-- ----------------------------
DROP TABLE IF EXISTS mio_role;
-- CREATE TABLE mio_role (
--     id SERIAL PRIMARY KEY,
--     created_by VARCHAR(100) NOT NULL DEFAULT '',
--     modified_by VARCHAR(100) NOT NULL DEFAULT '',
--     created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
--     updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
--     deleted_at TIMESTAMPTZ,
--     state SMALLINT NOT NULL DEFAULT 1,
    
--     user_id INT NOT NULL,
--     user_name VARCHAR(50) NOT NULL DEFAULT '',
--     value VARCHAR(50) DEFAULT ''
-- );

-- COMMENT ON TABLE mio_role IS '角色管理';

-- COMMENT ON COLUMN mio_role.id IS '主键 ID';
-- COMMENT ON COLUMN mio_role.created_by IS '创建人';
-- COMMENT ON COLUMN mio_role.modified_by IS '修改人';
-- COMMENT ON COLUMN mio_role.created_at IS '创建时间';
-- COMMENT ON COLUMN mio_role.updated_at IS '修改时间';
-- COMMENT ON COLUMN mio_role.deleted_at IS '删除时间';
-- COMMENT ON COLUMN mio_role.state IS '状态';
-- COMMENT ON COLUMN mio_role.user_id IS '用户 ID';
-- COMMENT ON COLUMN mio_role.user_name IS '用户名';
-- COMMENT ON COLUMN mio_role.value IS '用户属性';

-- ----------------------------
-- 创建 mio_menu 表
-- ----------------------------
DROP TABLE IF EXISTS mio_menu;
-- CREATE TABLE mio_menu (
--     id SERIAL PRIMARY KEY,
--     created_by VARCHAR(100) NOT NULL DEFAULT '',
--     modified_by VARCHAR(100) NOT NULL DEFAULT '',
--     created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
--     updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
--     deleted_at TIMESTAMPTZ,
--     state SMALLINT NOT NULL DEFAULT 1,
    
--     menu_name VARCHAR(50) NOT NULL DEFAULT '',
--     url VARCHAR(255) DEFAULT '',
--     parent_id INT NOT NULL,
--     parent_name VARCHAR(50) NOT NULL DEFAULT '',
--     level SMALLINT NOT NULL DEFAULT 2 -- 菜单等级
-- );

-- COMMENT ON TABLE mio_menu IS '菜单管理';

-- COMMENT ON COLUMN mio_menu.id IS '主键 ID';
-- COMMENT ON COLUMN mio_menu.created_by IS '创建人';
-- COMMENT ON COLUMN mio_menu.modified_by IS '修改人';
-- COMMENT ON COLUMN mio_menu.created_at IS '创建时间';
-- COMMENT ON COLUMN mio_menu.updated_at IS '修改时间';
-- COMMENT ON COLUMN mio_menu.deleted_at IS '删除时间';
-- COMMENT ON COLUMN mio_menu.state IS '状态';
-- COMMENT ON COLUMN mio_menu.menu_name IS '菜单名称';
-- COMMENT ON COLUMN mio_menu.url IS '菜单地址';
-- COMMENT ON COLUMN mio_menu.parent_id IS '父级菜单 ID';
-- COMMENT ON COLUMN mio_menu.parent_name IS '父级菜单名称';
-- COMMENT ON COLUMN mio_menu.level IS '菜单级别';

-- ----------------------------
-- 创建 mio_role_menu 表
-- ----------------------------
DROP TABLE IF EXISTS mio_role_menu;
-- CREATE TABLE mio_role_menu (
--     id SERIAL PRIMARY KEY,
--     created_by VARCHAR(100) NOT NULL DEFAULT '',
--     modified_by VARCHAR(100) NOT NULL DEFAULT '',
--     created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
--     updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
--     deleted_at TIMESTAMPTZ,
--     state SMALLINT NOT NULL DEFAULT 1,
    
--     role_id INT NOT NULL,
--     menu_id INT NOT NULL
-- );

-- COMMENT ON TABLE mio_role_menu IS '角色菜单关联';

-- COMMENT ON COLUMN mio_role_menu.id IS '主键 ID';
-- COMMENT ON COLUMN mio_role_menu.created_by IS '创建人';
-- COMMENT ON COLUMN mio_role_menu.modified_by IS '修改人';
-- COMMENT ON COLUMN mio_role_menu.created_at IS '创建时间';
-- COMMENT ON COLUMN mio_role_menu.updated_at IS '修改时间';
-- COMMENT ON COLUMN mio_role_menu.deleted_at IS '删除时间';
-- COMMENT ON COLUMN mio_role_menu.state IS '状态';
-- COMMENT ON COLUMN mio_role_menu.role_id IS '角色 ID';
-- COMMENT ON COLUMN mio_role_menu.menu_id IS '菜单 ID';

-- ----------------------------
-- 插入初始数据
-- ----------------------------
INSERT INTO mio_user (id, created_by, modified_by, username, password, avatar, user_type, state, created_at, updated_at, deleted_at)
VALUES
    (1, 'admin', 'admin', 'admin', '111111', 'https://zbj-bucket1.oss-cn-shenzhen.aliyuncs.com/avatar.JPG', 1, 1, '2019-08-19 21:00:39+00', '2019-08-19 21:00:39+00', NULL),
    (2, 'test', 'test', 'test', '111111', 'https://zbj-bucket1.oss-cn-shenzhen.aliyuncs.com/avatar.JPG', 2, 1, '2019-08-19 21:00:48+00', '2019-08-19 21:00:48+00', NULL);

INSERT INTO mio_tag (id, created_by, modified_by, name, state, created_at, updated_at, deleted_at)
VALUES
    (1, 'test', 'test', '1', 1, '2019-08-18 18:56:01+00', '2019-08-18 18:56:01+00', NULL),
    (2, 'test', 'test', '2', 1, '2019-08-16 18:56:06+00', '2019-08-16 18:56:06+00', NULL),
    (3, 'test', 'test', '3', 1, '2019-08-18 18:56:09+00', '2019-08-18 18:56:09+00', NULL);

INSERT INTO mio_article (id, created_by, modified_by, tag_id, title, description, content, cover_image_url, state, created_at, updated_at, deleted_at)
VALUES
    (1, 'test-created', 'test-created', 1, 'test1', 'test-desc', 'test-content', '', 1, '2019-08-19 21:00:39+00', '2019-08-19 21:00:39+00', NULL),
    (2, 'test-created', 'test-created', 1, 'test2', 'test-desc', 'test-content', '', 2, '2019-08-19 21:00:48+00', '2019-08-19 21:00:48+00', NULL),
    (3, 'test-created', 'test-created', 1, 'test3', 'test-desc', 'test-content', '', 1, '2019-08-19 21:00:49+00', '2019-08-19 21:00:49+00', NULL);

-- INSERT INTO mio_role (id, created_by, modified_by, user_id, user_name, value, created_at, updated_at, deleted_at)
-- VALUES
--     (1, 'admin', 'admin', 1, 'admin', 'admin', '2019-08-19 21:00:39+00', '2019-08-19 21:00:39+00', NULL),
--     (2, 'admin', 'admin', 1, 'admin', 'test', '2019-08-19 21:00:39+00', '2019-08-19 21:00:39+00', NULL),
--     (3, 'test', 'test', 2, 'test', 'test', '2019-08-19 21:00:39+00', '2019-08-19 21:00:39+00', NULL);

-- INSERT INTO mio_menu (id, created_by, modified_by, menu_name, url, parent_id, parent_name, level, state, created_at, updated_at, deleted_at)
-- VALUES
--     (1, 'admin', 'admin', 'Dashboard', '/dashboard', 0, 'Root', 1, 1, '2023-01-01 00:00:00+00', '2023-01-01 00:00:00+00', NULL),
--     (2, 'admin', 'admin', 'Settings', '/settings', 0, 'Root', 1, 1, '2023-01-01 00:00:00+00', '2023-01-01 00:00:00+00', NULL);

-- INSERT INTO mio_role_menu (id, created_by, modified_by, role_id, menu_id, created_at, updated_at, deleted_at)
-- VALUES
--     (1, 'admin', 'admin', 1, 1, '2023-01-01 00:00:00+00', '2023-01-01 00:00:00+00', NULL),
--     (2, 'admin', 'admin', 1, 2, '2023-01-01 00:00:00+00', '2023-01-01 00:00:00+00', NULL),
--     (3, 'admin', 'admin', 2, 1, '2023-01-01 00:00:00+00', '2023-01-01 00:00:00+00', NULL);