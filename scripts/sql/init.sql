CREATE DATABASE miao
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.UTF-8'
    LC_CTYPE = 'en_US.UTF-8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;


\c miao


CREATE TABLE IF NOT EXISTS model (
    id SERIAL PRIMARY KEY,
    created_by VARCHAR(100) NOT NULL DEFAULT '',
    modified_by VARCHAR(100) NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    is_del SMALLINT NOT NULL DEFAULT 0
);


COMMENT ON TABLE model IS '公共模型，处理公共字段';


COMMENT ON COLUMN model.id IS '主键 ID';
COMMENT ON COLUMN model.created_by IS '创建人';
COMMENT ON COLUMN model.modified_by IS '修改人';
COMMENT ON COLUMN model.created_at IS '创建时间';
COMMENT ON COLUMN model.updated_at IS '修改时间';
COMMENT ON COLUMN model.deleted_at IS '删除时间';
COMMENT ON COLUMN model.is_del IS '是否删除 0 为未删除、1 为已删除';


CREATE TABLE IF NOT EXISTS miao_tag (
    id SERIAL PRIMARY KEY,
    created_by VARCHAR(100) NOT NULL DEFAULT '',
    modified_by VARCHAR(100) NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    is_del SMALLINT NOT NULL DEFAULT 0,
    name VARCHAR(100) NOT NULL DEFAULT '',
    state SMALLINT NOT NULL DEFAULT 1
);

COMMENT ON TABLE miao_tag IS '标签管理';


COMMENT ON COLUMN miao_tag.id IS '主键 ID';
COMMENT ON COLUMN miao_tag.created_by IS '创建人';
COMMENT ON COLUMN miao_tag.modified_by IS '修改人';
COMMENT ON COLUMN miao_tag.created_at IS '创建时间';
COMMENT ON COLUMN miao_tag.updated_at IS '修改时间';
COMMENT ON COLUMN miao_tag.deleted_at IS '删除时间';
COMMENT ON COLUMN miao_tag.is_del IS '是否删除 0 为未删除、1 为已删除';
COMMENT ON COLUMN miao_tag.name IS '标签名称';
COMMENT ON COLUMN miao_tag.state IS '状态 0 为禁用、1 为启用';


CREATE TABLE IF NOT EXISTS miao_article (
    id SERIAL PRIMARY KEY,
    created_by VARCHAR(100) NOT NULL DEFAULT '',
    modified_by VARCHAR(100) NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    is_del SMALLINT NOT NULL DEFAULT 0,
    title VARCHAR(100) NOT NULL DEFAULT '',
    description VARCHAR(255) NOT NULL DEFAULT '',
    content TEXT NOT NULL,
    cover_image_url VARCHAR(255) NOT NULL DEFAULT '',
    state SMALLINT NOT NULL DEFAULT 1
);


COMMENT ON TABLE miao_article IS '文章管理';


COMMENT ON COLUMN miao_article.id IS '主键 ID';
COMMENT ON COLUMN miao_article.created_by IS '创建人';
COMMENT ON COLUMN miao_article.modified_by IS '修改人';
COMMENT ON COLUMN miao_article.created_at IS '创建时间';
COMMENT ON COLUMN miao_article.updated_at IS '修改时间';
COMMENT ON COLUMN miao_article.deleted_at IS '删除时间';
COMMENT ON COLUMN miao_article.is_del IS '是否删除 0 为未删除、1 为已删除';
COMMENT ON COLUMN miao_article.title IS '文章标题';
COMMENT ON COLUMN miao_article.description IS '文章简述';
COMMENT ON COLUMN miao_article.content IS '文章内容';
COMMENT ON COLUMN miao_article.cover_image_url IS '封面图片地址';
COMMENT ON COLUMN miao_article.state IS '状态 0 为禁用、1 为启用';


CREATE TABLE IF NOT EXISTS miao_article_tag (
    id SERIAL PRIMARY KEY,
    created_by VARCHAR(100) NOT NULL DEFAULT '',
    modified_by VARCHAR(100) NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    is_del SMALLINT NOT NULL DEFAULT 0,
    article_id INT NOT NULL,
    tag_id INT NOT NULL
);


COMMENT ON TABLE miao_article_tag IS '文章标签关联';


COMMENT ON COLUMN miao_article_tag.id IS '主键 ID';
COMMENT ON COLUMN miao_article_tag.created_by IS '创建人';
COMMENT ON COLUMN miao_article_tag.modified_by IS '修改人';
COMMENT ON COLUMN miao_article_tag.created_at IS '创建时间';
COMMENT ON COLUMN miao_article_tag.updated_at IS '修改时间';
COMMENT ON COLUMN miao_article_tag.deleted_at IS '删除时间';
COMMENT ON COLUMN miao_article_tag.is_del IS '是否删除 0 为未删除、1 为已删除';
COMMENT ON COLUMN miao_article_tag.article_id IS '文章 ID';
COMMENT ON COLUMN miao_article_tag.tag_id IS '标签 ID';

-- 添加外键约束到 miao_article_tag 表
ALTER TABLE miao_article_tag
    ADD CONSTRAINT fk_article
        FOREIGN KEY (article_id) REFERENCES miao_article(id)
        ON DELETE CASCADE,
    ADD CONSTRAINT fk_tag
        FOREIGN KEY (tag_id) REFERENCES miao_tag(id)
        ON DELETE CASCADE;


CREATE TABLE IF NOT EXISTS miao_auth (
    id SERIAL PRIMARY KEY,
    app_key VARCHAR(20) NOT NULL DEFAULT '',
    app_secret VARCHAR(50) NOT NULL DEFAULT '',
    created_by VARCHAR(100) NOT NULL DEFAULT '',
    modified_by VARCHAR(100) NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    is_del SMALLINT NOT NULL DEFAULT 0
);

COMMENT ON TABLE miao_auth IS '认证管理';

COMMENT ON COLUMN miao_auth.id IS '主键 ID';
COMMENT ON COLUMN miao_auth.app_key IS 'Key';
COMMENT ON COLUMN miao_auth.app_secret IS 'Secret';
COMMENT ON COLUMN miao_auth.created_by IS '创建人';
COMMENT ON COLUMN miao_auth.modified_by IS '修改人';
COMMENT ON COLUMN miao_auth.created_at IS '创建时间';
COMMENT ON COLUMN miao_auth.updated_at IS '修改时间';
COMMENT ON COLUMN miao_auth.deleted_at IS '删除时间';
COMMENT ON COLUMN miao_auth.is_del IS '是否删除 0 为未删除、1 为已删除';