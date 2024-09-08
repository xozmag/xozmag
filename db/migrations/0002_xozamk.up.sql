CREATE TYPE role AS ENUM ('admin', 'seller', 'user');

ALTER TABLE users
     ADD role role;