CREATE TABLE IF NOT EXISTS movies
(
    id         bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title      varchar(50)                 NOT NULL,
    year       int                         NOT NULL,
    runtime    int                         NOT NULL,
    genres     text[]                      NOT NULL,
    version    int                         NOT NULL DEFAULT 1
);
