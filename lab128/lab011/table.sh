#!/bin/bash

user="root"
passwd="123456"
tmp_file="temp.sql"
host="127.0.0.1"
port=3306
db_name="gorm_test"

create_user_table() {
  cat <<EOF >$tmp_file
	#$1是传进来的第一个参数
create table t_user
(
    id                 int unsigned primary key auto_increment comment '无关逻辑的主键',
    username           varchar(32)       default null comment '用户名',
    password           varchar(64)       default null comment '密码'
) ENGINE = InnoDB
  CHARSET = utf8
  COLLATE = utf8_general_ci
  ROW_FORMAT = Compact comment ='用户';
EOF
}

# 批量创建分表
create_block_tables() {
  cat <<EOF >$tmp_file
	#$1是传进来的第一个参数
	create table t_block_$1
(
    id          int unsigned primary key auto_increment comment '无关逻的主键',
    user_id     int         not null comment '用户id',
    wxid        varchar(32) not null comment 'wxid',
    block_type  int         not null comment '屏蔽类型，1为好友，2为群内',
    block_data1 varchar(32) not null comment '屏蔽数据',
    block_data2 varchar(32) comment '屏蔽数据'
) ENGINE = InnoDB
  CHARSET = utf8
  COLLATE = utf8_general_ci
  ROW_FORMAT = Compact comment ='聊天屏蔽规则';
EOF
}

create_user_table
cat $tmp_file | mysql -u root -p$passwd -h $host -P $port $db_name

table_index=0
while [ $table_index -lt 10 ]; do
  tbx=$(printf "%01d" $table_index)
  create_block_tables $tbx
  cat $tmp_file | mysql -u root -p$passwd -h $host -P $port $db_name
  table_index=$(expr $table_index + 1)
done

