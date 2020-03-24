-- 数据库
CREATE DATABASE IF NOT EXISTS temp_db DEFAULT CHARACTER SET utf8;
use temp_db;

-- 用户表
drop table if exists t_user;
create table t_user
(
    id          int unsigned primary key auto_increment comment '无关逻辑的主键',
    username    varchar(32) not null comment '用户名',
    password    varchar(64) not null comment '密码',
    phone       char(11)             default null comment '手机号',
    create_time datetime    not null default NOW() comment '创建时间',
    update_time datetime             default null comment '更新时间'
)
    ENGINE = InnoDB
    CHARSET = utf8
    COLLATE = utf8_general_ci
    ROW_FORMAT = Compact
    comment ='用户表';

-- 测试数据
insert into t_user(username, password)
values ('admin', '123456');

