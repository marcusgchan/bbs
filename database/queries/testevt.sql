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

-- name: GetTestEvtResults :one
SELECT sqlc.embed(test_events), sqlc.embed(test_results), sqlc.embed(templates)
FROM test_events
JOIN test_results ON test_events.testResultId = test_results.id
JOIN templates ON test_events.templateId = templates.id
WHERE test_events.id = ?;

-- name: GetTestEvtPlayerResults :many
SELECT sqlc.embed(player_test_results), sqlc.embed(players)
FROM player_test_results
JOIN players ON player_test_results.playerId = players.id
WHERE player_test_results.testResultId = ?;
