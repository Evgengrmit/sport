CREATE TABLE workout_result
(
    id          bigserial not null primary key,
    user_id     int references users (id),
    workout_id  int references workout_day (id),
    time_second int       not null,
    time_cap    int       not null,
    comment     varchar   not null,
    created_at  timestamp not null default current_date
);