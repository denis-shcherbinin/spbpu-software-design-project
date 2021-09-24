CREATE TABLE IF NOT EXISTS t_session
(
    "user_id" bigint REFERENCES t_user (id) ON DELETE CASCADE not null,
    "token"   text unique                                     not null
);
