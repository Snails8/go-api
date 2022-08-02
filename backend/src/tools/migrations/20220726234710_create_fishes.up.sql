CREATE TABLE IF NOT EXISTS fishes
(
    id serial,
    name varchar(256) NOT NULL,
    classification text NOT NULL,
    type_id INT NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,

    PRIMARY KEY (id),
    UNIQUE (name)
)