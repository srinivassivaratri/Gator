-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: DeleteFeed :exec
DELETE FROM feeds WHERE id = $1 AND user_id = $2;

-- name: GetFeedsWithUser :many
SELECT 
    feeds.id,
    feeds.created_at,
    feeds.updated_at,
    feeds.name,
    feeds.url,
    feeds.user_id,
    users.name as user_name
FROM feeds
JOIN users ON feeds.user_id = users.id
WHERE feeds.user_id = $1;