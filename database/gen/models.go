// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"database/sql"
	"time"
)

type Component struct {
	Name string
	Type string
}

type Player struct {
	ID        string
	Name      string
	Createdat sql.NullTime
	Updatedat sql.NullTime
}

type PlayerComponent struct {
	Playerid  sql.NullString
	Component sql.NullString
	Count     int64
}

type PlayerTestResult struct {
	Playerid     string
	Testresultid int64
	Wavedied     int64
	Diedto       string
}

type Template struct {
	ID        string
	Playerid  string
	Data      string
	Name      string
	Createdat time.Time
	Updatedat time.Time
}

type TestEvent struct {
	ID           string
	Environment  string
	Difficulty   string
	Templateid   int64
	Testresultid sql.NullInt64
	Startedat    time.Time
}

type TestResult struct {
	ID          int64
	Moneyearned int64
	Endedat     time.Time
}

type User struct {
	Username string
	Password string
}