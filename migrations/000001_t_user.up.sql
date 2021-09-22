CREATE TABLE IF NOT EXISTS "t_user"
(
    "id"            serial primary key,
    "first_name"    text                     not null,
    "second_name"   text                     not null,
    "username"      text unique              not null,
    "password_hash" text                     not null,
    "registered_at" timestamp with time zone not null default now()
);
