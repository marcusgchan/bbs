-- name: GetUser :one
SELECT * FROM users WHERE username = ?;

-- name: CreateTestEvt :exec
INSERT INTO test_events (environment, templateId, difficulty, startedAt) VALUES (?, ?, ?, ?);

-- name: CreatePlayerTemp :exec
INSERT INTO templates (playerId, data, name, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?);

-- name: CreateTestResult :one
INSERT INTO test_results (moneyEarned, endedAt) VALUES (?, ?) RETURNING id;

-- name: UpdateTestEvtWithTestRes :exec
UPDATE test_events SET testResultId = ? WHERE id = ?;

-- name: CreatePlayerTestResult :exec
INSERT INTO player_test_results (playerId, testResultId, waveDied, diedTo) VALUES (?, ?, ?, ?);

-- name: GetTestEvts :many
SELECT * FROM test_events;
