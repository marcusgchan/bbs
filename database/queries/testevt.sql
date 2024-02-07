-- name: GetUser :one
SELECT * FROM users WHERE username = ?;

-- name: CreateTestEvt :exec
INSERT INTO test_events (environment, templateId, difficulty, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?);

-- name: CreatePlayerTemp :exec
INSERT INTO templates (playerId, data, name, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?);

-- name: CreateTestResult :exec
INSERT INTO test_results (moneyEarned, createdAt, updatedAt) VALUES (?, ?, ?);

-- name: CreatePlayerTestResult :exec
INSERT INTO player_test_results (playerId, testResultId, waveDied, diedTo) VALUES (?, ?, ?, ?);
