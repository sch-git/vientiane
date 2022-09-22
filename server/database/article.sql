create table article
(
    id      bigint auto_increment primary key,
    title   varchar(32) not null,
    content longtext    not null,
    author  varchar(32) not null,
    ct      timestamp   not null,
    ut      timestamp   not null
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin comment '文章';