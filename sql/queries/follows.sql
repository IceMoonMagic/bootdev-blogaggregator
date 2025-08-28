-- name: CreateFeedFollow :one
WITH inserted AS (
    INSERT INTO feed_follows (
        user_id, feed_id
    ) VALUES (
        $1,
        (
            SELECT feeds.id FROM feeds
            WHERE feeds.url = $2
        )
    )
    RETURNING *
)
SELECT 
    inserted.*,
    users.name AS user_name,
    feeds.name as feed_name
FROM inserted
INNER JOIN users
    ON inserted.user_id = users.id
INNER JOIN feeds
    ON inserted.feed_id = feeds.id
;

-- name: GetFeedFollowsForUser :many
SELECT feed_follows.id, feeds.name, feeds.url
FROM feed_follows
INNER JOIN users
    ON feed_follows.user_id = users.id
INNER JOIN feeds
    ON feed_follows.feed_id = feeds.id
WHERE users.id = $1
;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE feed_follows.user_id = $1
AND feed_id = (
    SELECT id 
    FROM feeds
    WHERE url = $2
)
;