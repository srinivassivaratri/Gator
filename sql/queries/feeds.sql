-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds WHERE user_id = $1;

-- name: DeleteFeed :exec
DELETE FROM feeds WHERE id = $1 AND user_id = $2; 