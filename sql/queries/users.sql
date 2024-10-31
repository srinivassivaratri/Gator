-- Creates a new user (returns the created user)
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1, -- UUID
    $2, -- Created timestamp
    $3, -- Updated timestamp
    $4  -- Username
)
RETURNING *;

-- Finds user by name (returns one user)
SELECT * FROM users
WHERE name = $1;

-- Removes all users (returns nothing)
DELETE FROM users;

-- Gets all users (returns multiple users)
SELECT * FROM users;
