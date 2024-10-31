-- Creates users table when running migrations up
CREATE TABLE users (
    id UUID PRIMARY KEY,          -- Unique ID for each user
    created_at TIMESTAMP NOT NULL, -- When user was created
    updated_at TIMESTAMP NOT NULL, -- When user was last updated
    name TEXT NOT NULL UNIQUE     -- Username (must be unique)
);

-- Removes users table when running migrations down
DROP TABLE users;
