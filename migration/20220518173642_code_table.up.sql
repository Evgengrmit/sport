CREATE TYPE auth AS ENUM ('phone', 'email');
CREATE TABLE authorization_code
(
    id            bigserial   not null primary key,
    user_id       int references users (id),
    code          varchar(4)  not null,
    attempt_count int                  default 0,
    used_at       timestamp null,
    auth_type     auth,
    expired_at    timestamp not null default (current_timestamp + interval '5 minute'),
    created_at    timestamp not null default (current_timestamp )
);