package stats

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
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

const dbName = "file:./.local.db"

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

func init() {
	fmt.Println("Creating tables...")
	createTables()
}

func TestMostRecentStatsWithNoContent(t *testing.T) {
	db, err := sql.Open("libsql", dbName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s", err)
		os.Exit(1)
	}
	defer db.Close()
	q := sqlc.New(db)

	mc := MockCtx{}

	t.Run("it should return empty array - limit 5", func(t *testing.T) {
		limit := int64(5)
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
		if len(data) != 0 {
			t.Fail()
		}
	})
}

func TestMostRecentStatsWithContent(t *testing.T) {
	db, err := sql.Open("libsql", dbName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s", err)
		os.Exit(1)
	}
	defer db.Close()
	q := sqlc.New(db)
	mc := MockCtx{}

	defer seed()()

	t.Run("it should", func(t *testing.T) {
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

		if len(data) != 0 {
			t.Fail()
		}
	})
}

func createTables() {
	var out bytes.Buffer
	var stderr bytes.Buffer
	pathToSchema := fmt.Sprintf("file://%s", path.Join(basepath, "../../database/schema.sql"))
	pathToDb := fmt.Sprintf("sqlite://%s", path.Join(basepath, "./.local.db"))
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

func seed() func() {
	return func() {
	}
}
