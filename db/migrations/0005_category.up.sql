CREATE TABLE category (
   id uuid not null PRIMARY KEY,
   name varchar(200) not null,
   photo VARCHAR ,
   state numeric(1) not null DEFAULT 1,
   created_at TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE sub_category (
    id uuid not null PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    photo VARCHAR,
    state numeric(1) not null DEFAULT 1,
    created_at TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    category_id uuid REFERENCES category(id) 
)