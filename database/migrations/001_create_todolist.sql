-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE IF NOT EXISTS todolist (
    id int NOT NULL AUTO_INCREMENT,
    title varchar(255) not null unique,
    description varchar(255),
    due_date timestamp,
    completed bool,
    PRIMARY KEY (id)
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS todolist;
