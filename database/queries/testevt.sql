-- name: GetUser :one
SELECT * FROM users WHERE username = ?;

-- name: CreateTestEvt :exec
INSERT INTO test_events (environment, templateId, difficulty, createdAt) VALUES (?, ?, ?, ?);

-- name: CreatePlayerTemp :exec
INSERT INTO templates (playerId, data, name, createdAt) VALUES (?, ?, ?, ?);
