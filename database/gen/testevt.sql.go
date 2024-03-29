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
INSERT INTO templates (playerId, data, name, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?)
`

type CreatePlayerTempParams struct {
	Playerid  string
	Data      string
	Name      string
	Createdat time.Time
	Updatedat time.Time
}

func (q *Queries) CreatePlayerTemp(ctx context.Context, arg CreatePlayerTempParams) error {
	_, err := q.db.ExecContext(ctx, createPlayerTemp,
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
