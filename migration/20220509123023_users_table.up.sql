CREATE TABLE users
(
    id bigserial not null primary key,
    name         varchar(255) not null,
    email        varchar(255) not null unique ,
    created_at   timestamp        not null default current_date
);