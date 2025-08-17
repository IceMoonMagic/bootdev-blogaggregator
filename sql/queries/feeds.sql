-- name: CreateFeed :one
INSERT INTO
feeds (name, url, user_id)
VALUES ($1, $2, (SELECT users.id FROM users WHERE users.name = @user_name))
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