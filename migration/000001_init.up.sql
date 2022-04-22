CREATE TABLE workout_day
(
    id           bigserial    not null primary key,
    title        varchar(255) not null,
    description  varchar      not null,
    scheduled_at date         not null,
    created_at   date         not null default current_date
);
CREATE TABLE trainer
(
    id         bigserial    not null primary key,
    name       varchar(255) not null,
    avatar_url varchar      not null
);
CREATE TABLE schedule
(
    id           bigserial    not null primary key,
    name         varchar(255) not null,
    scheduled_at time         not null,
    trainer_id   int references trainer (id)
);
