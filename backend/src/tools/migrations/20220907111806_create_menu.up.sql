CREATE TABLE IF NOT EXISTS menu
(
    id serial,
    title varchar(256) NOT NULL,
    date date NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,

    PRIMARY KEY (id)
)