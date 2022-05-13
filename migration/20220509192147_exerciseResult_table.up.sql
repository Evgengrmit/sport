CREATE TABLE exercise_results(
    id bigserial not null primary key,
    exercise_id int references exercise(id),
    user_id int references users(id),
    comment varchar,
    created_at timestamp not null default current_date
);
