-- name: GetPlayers :many
SELECT * FROM players
ORDER BY updatedAt DESC
LIMIT ? OFFSET ?;

-- name: CreatePlayer :exec
INSERT OR REPLACE INTO players (id, name) VALUES (?, ?);
