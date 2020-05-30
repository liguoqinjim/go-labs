# 查询为空

## NOTICE
 - Association和preload联合查询的时候，是直接用变量的名字，还有struct直接转到name再加`Id`。这里可以查看源码
 - Association查询的时候需要指定`gorm:"ASSOCIATION_FOREIGNKEY:id"`，而且需要先查出来user，再查company
 - Preload不需要先查user，user和company可以一起出结果

## sql
```sql
create table t_user
(
    id       int unsigned primary key auto_increment comment '无关逻辑的主键',
    username varchar(32) default null comment '用户名'
);
insert into t_user(username)
values ('tom'),
       ('ben'),
       ('john');

create table t_company
(
    id           int unsigned primary key auto_increment comment '无关逻辑的主键',
    user_id      int         not null comment '用户id',
    company_name varchar(32) not null comment '公司名称'
);
insert into t_company (user_id, company_name)
values (1, '1号公司'),
       (1, '1号分公司'),
       (3, 'JOHN公司');
```