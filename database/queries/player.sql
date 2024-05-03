-- name: GetPlayers :many
SELECT * FROM players
ORDER BY updatedAt DESC
LIMIT ? OFFSET ?;

-- name: UpsertPlayer :exec
INSERT INTO players (id, name, updatedAt) 
VALUES (?, ?, ?) ON CONFLICT (id)
DO UPDATE SET name = excluded.name AND updatedAt = excluded.updatedAt;
