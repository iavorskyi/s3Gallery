create table if not exists users
(
    id         serial
        constraint users_pk
            primary key,
    email      varchar(32)  not null
        unique,
    first_name varchar(32),
    last_name  varchar(32),
    password   varchar(256) not null
);

alter table users
    owner to admin;