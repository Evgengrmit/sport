CREATE TABLE exercise
(
    id         bigserial not null primary key,
    name       varchar   not null,
    created_at timestamp not null default current_date
);