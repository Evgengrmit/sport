CREATE TABLE workout_day
(
    id           bigserial    not null primary key,
    title        varchar(255) not null,
    description  varchar      not null,
    scheduled_at timestamp       not null,
    created_at   timestamp        not null default current_date
);
CREATE TABLE trainer
(
    id         bigserial    not null primary key,
    name       varchar(255) not null,
    avatar_url varchar      null
);
CREATE TABLE schedule
(
    id           bigserial    not null primary key,
    name         varchar(255) not null,
    scheduled_at timestamp         not null,
    trainer_id   int references trainer (id)
);
CREATE TABLE users
(
    id bigserial not null primary key,
    name         varchar(255) not null,
    email        varchar(255) not null unique ,
    created_at   timestamp        not null default current_date
);