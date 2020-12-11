show databases;

use gorm_lab;
show tables;

drop table if exists t_user;
create table t_user
(
    id          int primary key auto_increment,
    username    varchar(32) not null comment '用户名',
    create_time datetime    not null default now() comment '创建时间'
);
