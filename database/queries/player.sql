-- name: GetPlayers :many
SELECT * FROM players;

-- name: CreatePlayer :exec
INSERT INTO players (id, name) VALUES (?, ?);
