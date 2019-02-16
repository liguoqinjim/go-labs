create database go_sql_lab;

create table student(
    id int not null primary key auto_increment comment '无关逻辑的主键',
    sid int not null comment '学号',
    sname varchar(32) not null comment '姓名',
    sage int comment '年龄'
);
