CREATE TABLE IF NOT EXISTS users
(
    id serial,
    name varchar(256) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    PRIMARY KEY (id)
);