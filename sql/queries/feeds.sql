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
    feeds.*,
    users.name as user_name
FROM feeds
JOIN users ON feeds.user_id = users.id
WHERE feeds.user_id = $1;

-- name: GetFeedByURL :one
SELECT * FROM feeds WHERE url = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds 
ORDER BY COALESCE(last_fetched_at, TIMESTAMP '1970-01-01') ASC
LIMIT 1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = NOW(),
    updated_at = NOW()
WHERE id = $1;