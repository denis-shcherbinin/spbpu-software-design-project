BEGIN;

CREATE TABLE IF NOT EXISTS t_list
(
    id          serial primary key,
    title       text                     not null,
    description text                     not null default '',
    created_at  timestamp with time zone not null default now()
);

CREATE TABLE IF NOT EXISTS t_user_list
(
    user_id bigint REFERENCES t_user (id) ON DELETE CASCADE,
    list_id bigint REFERENCES t_list (id) ON DELETE CASCADE
);

COMMIT;
