-- name: GetPlayers :many
SELECT * FROM players;

-- name: CreatePlayer :exec
INSERT OR REPLACE INTO players (id, name) VALUES (?, ?);

-- name: GetBlocks :many
SELECT name, type, count FROM player_blocks;
