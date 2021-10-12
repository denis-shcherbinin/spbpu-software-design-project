BEGIN;

CREATE TABLE t_item
(
    id          serial primary key,
    title       text                     not null,
    description text                     not null default '',
    done        boolean                  not null default false,
    created_at  timestamp with time zone not null default now()
);

CREATE TABLE t_list_item
(
    list_id bigint REFERENCES t_list (id) ON DELETE CASCADE,
    item_id bigint REFERENCES t_item (id) ON DELETE CASCADE
);

COMMIT;
