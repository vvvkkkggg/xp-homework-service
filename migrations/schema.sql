--goose up

create table if not exists users
(
    user_id  SERIAL PRIMARY KEY,
    role     int,
    group_id int,
    name     varchar(50)
);

create table if not exists tasks
(
    task_id     SERIAL PRIMARY KEY,
    group_id    int,
    name        varchar(50),
    description varchar(500)
);

create table if not exists assignments
(
    assignment_id SERIAL PRIMARY KEY,
    task_id int,
    user_id int,
    status  int,
    feedback text,
    file    text
);

--goose down

drop table if exists users;

drop table if exists tasks;

drop table if exists assignments;