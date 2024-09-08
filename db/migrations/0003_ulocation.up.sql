CREATE TABLE userslocations (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    name VARCHAR(255), 
    latitude DOUBLE PRECISION, -- Kenglik
    longitude DOUBLE PRECISION, -- Uzunlik
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);