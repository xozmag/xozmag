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