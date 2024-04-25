// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: stats.sql

package database

import (
	"context"
	"database/sql"
)

const getMostRecentStats = `-- name: GetMostRecentStats :many
SELECT Avg.version, Avg.avgWave, Max.maxWave, Count.numOfTestEvents, StartDate.startDate, EndDate.endDate
    FROM (
    SELECT test_events.version, AVG(player_test_results.waveDied) as avgWave
    FROM (
        SELECT version, testResultId FROM test_events
        WHERE test_events.testResultId IS NOT NULL
        ORDER BY version DESC
        LIMIT ?
    ) as S
    JOIN test_events ON test_events.version = S.version
    JOIN player_test_results ON player_test_results.testResultId = test_events.testResultId
    WHERE test_events.testResultId IS NOT NULL
    GROUP BY test_events.version
) as Avg, (
    SELECT test_events.version, MAX(player_test_results.waveDied) as maxWave
    FROM (
        SELECT version, testResultId FROM test_events
        WHERE test_events.testResultId IS NOT NULL
        ORDER BY version DESC
        LIMIT ?
    ) as S
    JOIN test_events ON test_events.version = S.version
    JOIN player_test_results ON player_test_results.testResultId = test_events.testResultId
    WHERE test_events.testResultId IS NOT NULL
    GROUP BY test_events.version
) as Max, (
    SELECT test_events.version, COUNT(test_events.id) as numOfTestEvents
    FROM (
        SELECT version, testResultId FROM test_events
        WHERE test_events.testResultId IS NOT NULL
        ORDER BY version DESC
        LIMIT ?
    ) as S
    JOIN test_events ON test_events.version = S.version
    WHERE test_events.testResultId IS NOT NULL
    GROUP BY test_events.version
) as Count, (
    SELECT test_events.version, MIN(test_events.startedAt) as startDate
    FROM (
        SELECT version, testResultId FROM test_events
        WHERE test_events.testResultId IS NOT NULL
        ORDER BY version DESC
        LIMIT ?
    ) as S
    JOIN test_events ON test_events.version = S.version
    JOIN player_test_results ON player_test_results.testResultId = test_events.testResultId
    WHERE test_events.testResultId IS NOT NULL
    GROUP BY test_events.version
) as StartDate, (
    SELECT test_events.version, MAX(test_results.endedAt) as endDate
    FROM (
        SELECT version, testResultId FROM test_events
        WHERE test_events.testResultId IS NOT NULL
        ORDER BY version DESC
        LIMIT ?
    ) as S
    JOIN test_events ON test_events.version = S.version
    JOIN player_test_results ON player_test_results.testResultId = test_events.testResultId
    JOIN test_results ON test_events.testResultId = test_results.id
    WHERE test_events.testResultId IS NOT NULL
    GROUP BY test_events.version
) as EndDate
WHERE Avg.version = Max.version
AND Max.version = Count.version 
AND Count.version = StartDate.version 
AND StartDate.version = EndDate.version
`

type GetMostRecentStatsParams struct {
	Limit   int64
	Limit_2 int64
	Limit_3 int64
	Limit_4 int64
	Limit_5 int64
}

type GetMostRecentStatsRow struct {
	Version         string
	Avgwave         sql.NullFloat64
	Maxwave         interface{}
	Numoftestevents int64
	Startdate       interface{}
	Enddate         interface{}
}

func (q *Queries) GetMostRecentStats(ctx context.Context, arg GetMostRecentStatsParams) ([]GetMostRecentStatsRow, error) {
	rows, err := q.db.QueryContext(ctx, getMostRecentStats,
		arg.Limit,
		arg.Limit_2,
		arg.Limit_3,
		arg.Limit_4,
		arg.Limit_5,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMostRecentStatsRow
	for rows.Next() {
		var i GetMostRecentStatsRow
		if err := rows.Scan(
			&i.Version,
			&i.Avgwave,
			&i.Maxwave,
			&i.Numoftestevents,
			&i.Startdate,
			&i.Enddate,
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

const getStatsByVersion = `-- name: GetStatsByVersion :one
SELECT Avg.version, Avg.avgWave, Max.maxWave, Count.numOfTestEvents, StartDate.startDate, EndDate.endDate
FROM (
    SELECT test_events.version, AVG(player_test_results.waveDied) as avgWave
    FROM test_events
    JOIN player_test_results ON player_test_results.testResultId = test_events.testResultId
    WHERE test_events.testResultId IS NOT NULL AND test_events.version = ?
) as Avg, (
    SELECT test_events.version, MAX(player_test_results.waveDied) as maxWave
    FROM test_events
    JOIN player_test_results ON player_test_results.testResultId = test_events.testResultId
    WHERE test_events.testResultId IS NOT NULL AND test_events.version = ?
) as Max, (
    SELECT test_events.version, COUNT(test_events.id) as numOfTestEvents
    FROM test_events
    WHERE test_events.testResultId IS NOT NULL AND test_events.version = ?
) as Count, (
    SELECT test_events.version, MIN(test_events.startedAt) as startDate
    FROM test_events
    WHERE test_events.testResultId IS NOT NULL AND test_events.version = ?
) as StartDate, (
    SELECT test_events.version, MAX(test_events.endedAt) as endDate
    FROM test_events
    JOIN test_results ON test_events.testResultId = test_result.id
    WHERE test_events.testResultId IS NOT NULL AND test_events.version = ?
) as EndDate
WHERE Avg.version = Max.version 
AND Max.version = Count.version
AND Count.version = StartDate.version
AND StartDate.version = EndDate.version
`

type GetStatsByVersionParams struct {
	Version   string
	Version_2 string
	Version_3 string
	Version_4 string
	Version_5 string
}

type GetStatsByVersionRow struct {
	Version         string
	Avgwave         sql.NullFloat64
	Maxwave         interface{}
	Numoftestevents int64
	Startdate       interface{}
	Enddate         interface{}
}

func (q *Queries) GetStatsByVersion(ctx context.Context, arg GetStatsByVersionParams) (GetStatsByVersionRow, error) {
	row := q.db.QueryRowContext(ctx, getStatsByVersion,
		arg.Version,
		arg.Version_2,
		arg.Version_3,
		arg.Version_4,
		arg.Version_5,
	)
	var i GetStatsByVersionRow
	err := row.Scan(
		&i.Version,
		&i.Avgwave,
		&i.Maxwave,
		&i.Numoftestevents,
		&i.Startdate,
		&i.Enddate,
	)
	return i, err
}