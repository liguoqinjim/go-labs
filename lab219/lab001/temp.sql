show databases;

create database db_test;
use db_test;

create table t_test01
(
    data_type int      not null comment '数据类型',
    `time`    datetime not null default NOW() comment '时间',
    num       int      not null comment '量'
);

select count(1)
from t_test01;

select *
from t_test01
limit 5;
# select to_days(time) from t_test01;

select sum(num), data_type
from t_test01
where to_days(time) = to_days('2020-03-18')
group by data_type;
select hour(time)
from t_test01;

select sum(num), data_type, date_format(time, "%Y%m%d%H")
from t_test01
where to_days(time) = to_days('2020-03-18')
group by data_type;

select sum(num), data_type, date_format(time, "%Y%m%d%H")
from t_test01 group by to_days(time)


