CREATE SEQUENCE seq_user_location;

CREATE TABLE users_locations (
    id BIGINT NOT NULL DEFAULT nextval('seq_user_location') PRIMARY KEY,
    name VARCHAR(255), 
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_id uuid NOT NULL REFERENCES users(id)
);