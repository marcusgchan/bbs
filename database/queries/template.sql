-- name: CreatePlayerTemp :exec
INSERT INTO templates (id, playerId, data, name, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?);
