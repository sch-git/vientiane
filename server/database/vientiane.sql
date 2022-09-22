# 创建数据库
create schema vientiane;

use vientiane;

# 账号表
create table vientiane_account
(
    id        bigint auto_increment comment '账号Id，主键',
    name      varchar(16) default ''                not null comment '用户名',
    password  varchar(64) default ''                not null comment '加密后的密码',
    email     varchar(64) default ''                not null comment '邮箱',
    updated_at timestamp   default current_timestamp not null,
    created_at timestamp   default current_timestamp not null comment '创建时间',
    constraint vientiane_account_pk
        primary key (id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin comment '账号信息';

# 内容表
create table vientiane_doc
(
    id          bigint auto_increment,
    content     longtext                              not null comment '文档内容 markdown',
    category_id bigint      default 0                 not null comment '分类id',
    author      varchar(32) default ''                not null comment '作者',
    created_at  timestamp   default current_timestamp not null comment '创建时间',
    updated_at  timestamp   default current_timestamp not null comment '更新时间',
    constraint vientiane_doc_pk
        primary key (id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin comment '文档';
