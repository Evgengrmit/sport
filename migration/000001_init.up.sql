CREATE TABLE workout_day
(
    id           bigserial not null unique,
    title        varchar(255) not null,
    description  varchar      not null,
    scheduled_at varchar(20)    not null,
    created_at   date not null default current_date
);