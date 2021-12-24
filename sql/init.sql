create table user
(
    id         int auto_increment,
    user_name  varchar(128)           not null,
    password   varchar(128)           not null,
    state      tinyint                not null comment '1',
    gmt_create datetime default now() not null,
    gmt_update datetime default now() not null,
    constraint user_pk
        primary key (id)
);

