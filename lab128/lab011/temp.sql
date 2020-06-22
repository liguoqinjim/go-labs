	#9是传进来的第一个参数
	create table t_block_9
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
