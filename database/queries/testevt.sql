-- name: GetUser :one
SELECT * FROM users WHERE username = ?;

-- name: CreateTestEvt :exec
INSERT INTO test_events (environmentId, templateId, difficultyId, createdAt) VALUES (?, ?, ?, ?);
