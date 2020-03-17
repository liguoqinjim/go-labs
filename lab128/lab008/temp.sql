create table t_user(
    id int primary key auto_increment,
    username varchar(6) not null
);

create table t_bill(
    id int primary key auto_increment,
    user_id int not null
);