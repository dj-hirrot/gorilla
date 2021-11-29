CREATE DATABASE IF NOT EXISTS gorilla;

USE gorilla;

CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    age INT NOT NULL
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

insert into
    users(name, age)
values
    ("山本五十六", 59);

insert into
    users(name, age)
values
    ("阿南惟幾", 58);

insert into
    users(name, age)
values
    ("米内光政", 68);

insert into
    users(name, age)
values
    ("東條英機", 63);

insert into
    users(name, age)
values
    ("鈴木貫太郎", 80);
