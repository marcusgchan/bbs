// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: testevt.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createPlayerTemp = `-- name: CreatePlayerTemp :exec
INSERT INTO templates (id, playerId, data, name, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?)
`

type CreatePlayerTempParams struct {
	ID        string
	Playerid  string
	Data      string
	Name      string
	Createdat time.Time
	Updatedat time.Time
}

func (q *Queries) CreatePlayerTemp(ctx context.Context, arg CreatePlayerTempParams) error {
	_, err := q.db.ExecContext(ctx, createPlayerTemp,
		arg.ID,
		arg.Playerid,
		arg.Data,
		arg.Name,
		arg.Createdat,
		arg.Updatedat,
	)
	return err
}

const createPlayerTestResult = `-- name: CreatePlayerTestResult :exec
INSERT INTO player_test_results (playerId, testResultId, waveDied, diedTo) VALUES (?, ?, ?, ?)
`

type CreatePlayerTestResultParams struct {
	Playerid     string
	Testresultid int64
	Wavedied     int64
	Diedto       string
}

func (q *Queries) CreatePlayerTestResult(ctx context.Context, arg CreatePlayerTestResultParams) error {
	_, err := q.db.ExecContext(ctx, createPlayerTestResult,
		arg.Playerid,
		arg.Testresultid,
		arg.Wavedied,
		arg.Diedto,
	)
	return err
}

const createTestEvt = `-- name: CreateTestEvt :exec
INSERT INTO test_events (environment, templateId, difficulty, startedAt) VALUES (?, ?, ?, ?)
`

type CreateTestEvtParams struct {
	Environment string
	Templateid  int64
	Difficulty  string
	Startedat   time.Time
}

func (q *Queries) CreateTestEvt(ctx context.Context, arg CreateTestEvtParams) error {
	_, err := q.db.ExecContext(ctx, createTestEvt,
		arg.Environment,
		arg.Templateid,
		arg.Difficulty,
		arg.Startedat,
	)
	return err
}

const createTestResult = `-- name: CreateTestResult :one
INSERT INTO test_results (moneyEarned, endedAt) VALUES (?, ?) RETURNING id
`

type CreateTestResultParams struct {
	Moneyearned int64
	Endedat     time.Time
}

func (q *Queries) CreateTestResult(ctx context.Context, arg CreateTestResultParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createTestResult, arg.Moneyearned, arg.Endedat)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const getTestEvtPlayerResults = `-- name: GetTestEvtPlayerResults :many
SELECT player_test_results.playerid, player_test_results.testresultid, player_test_results.wavedied, player_test_results.diedto, players.id, players.name, players.createdat, players.updatedat
FROM player_test_results
JOIN players ON player_test_results.playerId = players.id
WHERE player_test_results.testResultId = ?
`

type GetTestEvtPlayerResultsRow struct {
	PlayerTestResult PlayerTestResult
	Player           Player
}

func (q *Queries) GetTestEvtPlayerResults(ctx context.Context, testresultid int64) ([]GetTestEvtPlayerResultsRow, error) {
	rows, err := q.db.QueryContext(ctx, getTestEvtPlayerResults, testresultid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTestEvtPlayerResultsRow
	for rows.Next() {
		var i GetTestEvtPlayerResultsRow
		if err := rows.Scan(
			&i.PlayerTestResult.Playerid,
			&i.PlayerTestResult.Testresultid,
			&i.PlayerTestResult.Wavedied,
			&i.PlayerTestResult.Diedto,
			&i.Player.ID,
			&i.Player.Name,
			&i.Player.Createdat,
			&i.Player.Updatedat,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTestEvtResults = `-- name: GetTestEvtResults :one
SELECT test_events.id, test_events.environment, test_events.difficulty, test_events.templateid, test_events.testresultid, test_events.startedat, test_results.id, test_results.moneyearned, test_results.endedat, templates.id, templates.playerid, templates.data, templates.name, templates.createdat, templates.updatedat
FROM test_events
JOIN test_results ON test_events.testResultId = test_results.id
JOIN templates ON test_events.templateId = templates.id
WHERE test_events.id = ?
`

type GetTestEvtResultsRow struct {
	TestEvent  TestEvent
	TestResult TestResult
	Template   Template
}

func (q *Queries) GetTestEvtResults(ctx context.Context, id string) (GetTestEvtResultsRow, error) {
	row := q.db.QueryRowContext(ctx, getTestEvtResults, id)
	var i GetTestEvtResultsRow
	err := row.Scan(
		&i.TestEvent.ID,
		&i.TestEvent.Environment,
		&i.TestEvent.Difficulty,
		&i.TestEvent.Templateid,
		&i.TestEvent.Testresultid,
		&i.TestEvent.Startedat,
		&i.TestResult.ID,
		&i.TestResult.Moneyearned,
		&i.TestResult.Endedat,
		&i.Template.ID,
		&i.Template.Playerid,
		&i.Template.Data,
		&i.Template.Name,
		&i.Template.Createdat,
		&i.Template.Updatedat,
	)
	return i, err
}

const getTestEvts = `-- name: GetTestEvts :many
SELECT id, environment, difficulty, templateid, testresultid, startedat FROM test_events
`

func (q *Queries) GetTestEvts(ctx context.Context) ([]TestEvent, error) {
	rows, err := q.db.QueryContext(ctx, getTestEvts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TestEvent
	for rows.Next() {
		var i TestEvent
		if err := rows.Scan(
			&i.ID,
			&i.Environment,
			&i.Difficulty,
			&i.Templateid,
			&i.Testresultid,
			&i.Startedat,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one
SELECT username, password FROM users WHERE username = ?
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(&i.Username, &i.Password)
	return i, err
}

const updateTestEvtWithTestRes = `-- name: UpdateTestEvtWithTestRes :exec
UPDATE test_events SET testResultId = ? WHERE id = ?
`

type UpdateTestEvtWithTestResParams struct {
	Testresultid sql.NullInt64
	ID           string
}

func (q *Queries) UpdateTestEvtWithTestRes(ctx context.Context, arg UpdateTestEvtWithTestResParams) error {
	_, err := q.db.ExecContext(ctx, updateTestEvtWithTestRes, arg.Testresultid, arg.ID)
	return err
}
