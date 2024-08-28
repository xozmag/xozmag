create table xozmaks(
    id uuid not null primary key,
    created_at timestamp not null default current_timestamp,
    created_by uuid, 
    updated_at timestamp default current_timestamp,
    updated_by uuid,
    name varchar not null,
    location json,
    rate integer
);

CREATE TYPE gender AS ENUM ('erkak', 'ayol');

CREATE TABLE users (
    id UUID not null primary key,
    phone_number varchar(13) UNIQUE NOT NULL,
    firstname varchar,
    surname varchar,
    fathersname varchar, 
    birthdate DATE,
    gender gender,
    created_by uuid,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by uuid,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);