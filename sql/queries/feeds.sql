-- name: CreateFeed :one
INSERT INTO
feeds (name, url, user_id)
VALUES ($1, $2, $3)
RETURNING *
;

-- name: DebugGetFeed :one
SELECT * FROM feeds
WHERE name = $1
;

-- name: GetFeeds :many
SELECT feeds.name, feeds.url, users.name as user
FROM feeds
LEFT JOIN users
ON feeds.user_id = users.id
;

-- name: DebugGetFeeds :many
SELECT * FROM feeds
;

-- name: DeleteFeeds :exec
DELETE FROM feeds
;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET (updated_at, last_fetched_at) = NOW()
WHERE id = $1
;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1
;