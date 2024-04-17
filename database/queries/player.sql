-- name: GetPlayers :many
SELECT * FROM players;

-- name: CreatePlayer :exec
INSERT OR REPLACE INTO players (id, name) VALUES (?, ?);
