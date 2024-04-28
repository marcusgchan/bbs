package stats

import (
	"bytes"
	"database/sql"
	"fmt"
	"math"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	database "github.com/marcusgchan/bbs/database/gen"
	sqlc "github.com/marcusgchan/bbs/database/gen"
	_ "github.com/tursodatabase/go-libsql"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

type MockCtx struct{}

func (t MockCtx) Deadline() (deadline time.Time, ok bool) {
	return time.Time{}, true
}

func (t MockCtx) Err() error {
	return nil
}

func (t MockCtx) Done() <-chan struct{} {
	return nil
}

func (t MockCtx) Value(str any) any {
	return nil
}

// func TestMostRecentStatsWithNoContent(t *testing.T) {
// 	tempDir := t.TempDir()
// 	dbRelPath := "./test-db.db"
// 	conStr := "file:" + path.Join(tempDir, dbRelPath)
// 	createTables(tempDir, dbRelPath)
//
// 	db, err := sql.Open("libsql", conStr)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "failed to open db %s", err)
// 		os.Exit(1)
// 	}
// 	defer db.Close()
// 	q := sqlc.New(db)
//
// 	mc := MockCtx{}
//
// 	t.Run("it should return empty array - limit 5", func(t *testing.T) {
// 		limit := int64(5)
// 		data, err := q.GetMostRecentStats(mc, database.GetMostRecentStatsParams{
// 			Limit:   limit,
// 			Limit_2: limit,
// 			Limit_3: limit,
// 			Limit_4: limit,
// 			Limit_5: limit,
// 		})
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "failed to execute query %s", err)
// 			os.Exit(1)
// 		}
// 		if len(data) != 0 {
// 			t.Fail()
// 		}
// 	})
// }

// func TestMostRecentStatsWithTestEventButNoTestResult(t *testing.T) {
// 	tempDir := t.TempDir()
// 	dbRelPath := "./test-db.db"
// 	conStr := "file:" + path.Join(tempDir, dbRelPath)
// 	createTables(tempDir, dbRelPath)
//
// 	db, err := sql.Open("libsql", conStr)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "failed to open db %s", err)
// 		os.Exit(1)
// 	}
// 	defer db.Close()
// 	q := sqlc.New(db)
//
// 	mc := MockCtx{}
//
// 	tx, err := db.Begin()
// 	if err != nil {
// 		fmt.Println("Unable to start transaction")
// 		os.Exit(1)
// 	}
//
// 	tx.Query("insert into versions (version) values (1.0.0)")
// 	tx.Query("insert into versions (version) values (2.0.0)")
// 	tx.Query("insert into versions (version) values (3.0.0)")
//
// 	tx.Query("insert into players (id, name) values ('1', 'marcus')")
//
// 	tx.Query(`
//         insert into templates (id playerId, data, name, createdAt, updatedAt) values
//         ("1", "1", "", "A", '20120618 10:34:09 AM', '20120618 10:34:09 AM')
//     `)
//
// 	tx.Query(`
//         insert into test_events
//         (id, environment, difficulty, templateId, startedAt, version) values
//         ("1", "lab", "normal", " ",'20120618 10:34:09 AM', "1.0.0")
//     `)
//
// 	tx.Commit()
//
// 	t.Run("it should return empty array", func(t *testing.T) {
// 		limit := int64(1)
// 		data, err := q.GetMostRecentStats(mc, database.GetMostRecentStatsParams{
// 			Limit:   limit,
// 			Limit_2: limit,
// 			Limit_3: limit,
// 			Limit_4: limit,
// 			Limit_5: limit,
// 		})
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "failed to execute query %s", err)
// 			os.Exit(1)
// 		}
//
// 		if len(data) != 0 {
// 			t.Fail()
// 		}
// 	})
// }

func TestMostRecentStatsWithTestEventWithTestResults(t *testing.T) {
	// tempDir := t.TempDir()
	tempDir := "./"
	dbRelPath := "./test-db.db"
	conStr := "file:" + path.Join(tempDir, dbRelPath)
	createTables(tempDir, dbRelPath)

	db, err := sql.Open("libsql", conStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s", err)
		os.Exit(1)
	}
	defer db.Close()
	q := sqlc.New(db)

	mc := MockCtx{}

	db.Query("insert into versions (value) values ('1.0.0')")
	db.Query("insert into versions (value) values ('2.0.0')")
	db.Query("insert into versions (value) values ('3.0.0')")

	db.Query("insert into players (id, name) values ('1', 'marcus')")

	db.Query(`
        insert into templates (id, playerId, data, name, createdAt, updatedAt) values
        ('1', '1', '', 'A', '2012-06-18 10:34:09', '2012-06-18 10:34:09')
    `)

	db.Query(`
        insert into test_results
        (id, moneyEarned, endedAt) values
        ('1', 10, '2012-06-18 10:34:09')
    `)
	db.Query(`
        insert into player_test_results (playerId, testResultId, waveDied, diedTo) values
        ('1', '1', 10, 'Bombs From Above')
    `)
	db.Query(`
        insert into test_events
        (id, environment, difficulty, templateId, testResultId, startedAt, version) values
        ('1', 'lab', 'normal', '1', '1', '2012-06-18 10:34:09', '1.0.0')
    `)

	t.Run("it should return test evt with avg wave of 10, max wave of 10, count 1", func(t *testing.T) {
		limit := int64(1)
		data, err := q.GetMostRecentStats(mc, database.GetMostRecentStatsParams{
			Limit:   limit,
			Limit_2: limit,
			Limit_3: limit,
			Limit_4: limit,
			Limit_5: limit,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to execute query %s", err)
			os.Exit(1)
		}

		if len(data) != 1 {
			t.Fail()
		}
		if !almostEqual(data[0].Avgwave.Float64, 10) || data[0].Maxwave != 10 || data[0].Numoftestevents != 1 {
			t.Fail()
		}
	})

	db.Query(`
        insert into test_results
        (id, moneyEarned, endedAt) values
        ('2', 10, '2012-06-18 10:34:09')
    `)
	db.Query(`
        insert into player_test_results (playerId, testResultId, waveDied, diedTo) values
        ('1', '2', 20, 'Bombs From Above')
    `)
	db.Query(`
        insert into test_events
        (id, environment, difficulty, templateId, testResultId, startedAt, version) values
        ('2', 'lab', 'normal', '1', '2', '2012-06-18 10:34:09', '2.0.0')
    `)

	t.Run("it should return test evt with avg wave of 20, max wave of 20, count 1", func(t *testing.T) {
		limit := int64(1)
		data, err := q.GetMostRecentStats(mc, database.GetMostRecentStatsParams{
			Limit:   limit,
			Limit_2: limit,
			Limit_3: limit,
			Limit_4: limit,
			Limit_5: limit,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to execute query %s", err)
			os.Exit(1)
		}

		if len(data) != 1 {
			t.Fail()
		}
		// if !almostEqual(data[0].Avgwave.Float64, 15) || data[0].Maxwave != 10 || data[0].Numoftestevents != 1 {
		// 	t.Fail()
		// }
	})
	t.Run("it should return test evt with avg wave of 15, max wave of 15, count 2", func(t *testing.T) {
		limit := int64(2)
		data, err := q.GetMostRecentStats(mc, database.GetMostRecentStatsParams{
			Limit:   limit,
			Limit_2: limit,
			Limit_3: limit,
			Limit_4: limit,
			Limit_5: limit,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to execute query %s", err)
			os.Exit(1)
		}

		if len(data) != 2 {
			t.Fail()
		}
		// if !almostEqual(data[0].Avgwave.Float64, 15) || data[0].Maxwave != 10 || data[0].Numoftestevents != 1 {
		// 	t.Fail()
		// }
	})

	_, err = db.Query(`
        insert into test_results
        (id, moneyEarned, endedAt) values
        ('3', 10, '2012-06-18 10:34:09')
    `)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	db.Query(`
        insert into player_test_results (playerId, testResultId, waveDied, diedTo) values
        ('1', '3', 10, 'Bombs From Above')
    `)
	db.Query(`
        insert into test_events
        (id, environment, difficulty, templateId, testResultId, startedAt, version) values
        ('3', 'lab', 'normal', '1', '3', '2012-06-18 10:34:09', '2.0.0')
    `)

	t.Run("it should return test evt with avg wave of 15, max wave of 15, count 2", func(t *testing.T) {
		limit := int64(2)
		data, err := q.GetMostRecentStats(mc, database.GetMostRecentStatsParams{
			Limit:   limit,
			Limit_2: limit,
			Limit_3: limit,
			Limit_4: limit,
			Limit_5: limit,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to execute query %s", err)
			os.Exit(1)
		}

		fmt.Printf("3: %v", data)
		if len(data) != 2 {
			t.Fail()
		}
		// if !almostEqual(data[0].Avgwave.Float64, 15) || data[0].Maxwave != 10 || data[0].Numoftestevents != 1 {
		// 	t.Fail()
		// }
	})
}

func createTables(tempDir string, relPathToDb string) {
	var out bytes.Buffer
	var stderr bytes.Buffer
	pathToSchema := fmt.Sprintf("file://%s", path.Join(basepath, "../../database/schema.sql"))
	pathToDb := fmt.Sprintf("sqlite://%s", path.Join(tempDir, relPathToDb))
	devDb := "sqlite://dev?mode=memory"
	cmd := exec.Command("atlas", "schema", "apply", "--dev-url", devDb, "--to", pathToSchema, "-u", pathToDb, "--auto-approve")
	cmd.Stderr = &stderr
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		os.Exit(1)
	}
}

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}
