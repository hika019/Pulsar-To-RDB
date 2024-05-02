CREATE DATABASE test_database;

USE test_database;

CREATE TABLE users(
    id int auto_increment,
    name varchar(255) not null,
    age int,
    address text,
    PRIMARY KEY(id)
)
ENGINE = InnoDB, ROW_FORMAT = COMPRESSED;