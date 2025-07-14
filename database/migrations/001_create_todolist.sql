-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE IF NOT EXISTS todolist (
    id int NOT NULL AUTO_INCREMENT,
    title varchar(255) not null,
    description varchar(255),
    due_date timestamp,
    completed bool default(0),
    PRIMARY KEY (id),
    CONSTRAINT UC_Title UNIQUE (title)
);


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS todolist;
