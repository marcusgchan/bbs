package database

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	db := Connect()
	defer db.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("PASSWORD")), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("could not generate hash %v", err)
		return
	}
	_, err = db.Query("insert into users (username, password) values (?, ?)", os.Getenv("USERNAME"), hashedPassword)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	db.Query("insert into versions (value) values ('1.0.0')")

	db.Query("insert into players (id, name) values ('1', 'marcus')")

	db.Query(`
        insert into templates (id, playerId, data, name, createdAt, updatedAt) values
        ('1', '1', '[{"y":-43,"x":0,"name":"wood","z":20},{"y":-35,"x":0,"name":"wood","z":16},{"y":-27,"x":0,"name":"wood","z":12},{"y":45,"x":2,"name":"wood","z":4},{"y":39,"x":2,"name":"wood","z":2},{"y":33,"x":2,"name":"wood","z":4},{"y":-19,"x":0,"name":"wood","z":2},{"y":-19,"x":6,"name":"wood","z":2},{"y":-19,"x":2,"name":"wood","z":0},{"y":-19,"x":4,"name":"wood","z":0},{"y":-19,"x":0,"name":"wood","z":4},{"y":-19,"x":2,"name":"wood","z":2},{"y":-19,"x":4,"name":"wood","z":2},{"y":-19,"x":2,"name":"wood","z":6},{"y":-19,"x":0,"name":"wood","z":6},{"y":-19,"x":2,"name":"wood","z":4},{"y":-19,"x":4,"name":"wood","z":4},{"y":-19,"x":0,"name":"wood","z":8},{"y":-19,"x":6,"name":"wood","z":0},{"y":-17,"x":-4,"name":"wood","z":2},{"y":-17,"x":-4,"name":"wood","z":4},{"y":-17,"x":0,"name":"wood","z":-2},{"y":-17,"x":2,"name":"wood","z":8},{"y":-17,"x":2,"name":"wood","z":-2},{"y":-17,"x":10,"name":"wood","z":2},{"y":-17,"x":0,"name":"wood","z":0},{"y":-17,"x":10,"name":"wood","z":4},{"y":-17,"x":10,"name":"wood","z":0},{"y":-17,"x":-2,"name":"wood","z":0},{"y":-17,"x":-2,"name":"wood","z":4},{"y":-17,"x":8,"name":"wood","z":4},{"y":-17,"x":4,"name":"wood","z":6},{"y":-17,"x":2,"name":"wood","z":6},{"y":-17,"x":-2,"name":"wood","z":2},{"y":-17,"x":2,"name":"wood","z":-4},{"y":-17,"x":8,"name":"wood","z":2},{"y":-17,"x":6,"name":"wood","z":2},{"y":-17,"x":6,"name":"wood","z":4},{"y":-17,"x":4,"name":"wood","z":4},{"y":-17,"x":8,"name":"wood","z":0},{"y":-15,"x":6,"name":"wood","z":2},{"y":-15,"x":-2,"name":"wood","z":2},{"y":-15,"x":-2,"name":"wood","z":4},{"y":-15,"x":-4,"name":"wood","z":2},{"y":-15,"x":4,"name":"wood","z":4},{"y":-15,"x":-4,"name":"wood","z":4},{"y":-15,"x":10,"name":"wood","z":2},{"y":-15,"x":4,"name":"wood","z":6},{"y":-15,"x":10,"name":"wood","z":0},{"y":-15,"x":0,"name":"wood","z":0},{"y":-15,"x":10,"name":"wood","z":4},{"y":-15,"x":2,"name":"wood","z":6},{"y":-15,"x":8,"name":"wood","z":0},{"y":-15,"x":8,"name":"wood","z":4},{"y":-15,"x":-2,"name":"wood","z":0},{"y":-15,"x":2,"name":"wood","z":8},{"y":-15,"x":2,"name":"wood","z":-4},{"y":-15,"x":2,"name":"wood","z":-2},{"y":-15,"x":8,"name":"wood","z":2},{"y":-15,"x":0,"name":"wood","z":-2},{"y":-15,"x":6,"name":"wood","z":4},{"y":-13,"x":2,"name":"wood","z":-4},{"y":-13,"x":-4,"name":"wood","z":2},{"y":-13,"x":-4,"name":"wood","z":4},{"y":-13,"x":2,"name":"wood","z":8},{"y":-13,"x":10,"name":"wood","z":0},{"y":-13,"x":-2,"name":"wood","z":0},{"y":-13,"x":8,"name":"wood","z":4},{"y":-13,"x":-2,"name":"wood","z":4},{"y":-13,"x":10,"name":"wood","z":2},{"y":-13,"x":-2,"name":"wood","z":2},{"y":-13,"x":0,"name":"wood","z":-2},{"y":-13,"x":0,"name":"wood","z":0},{"y":-13,"x":2,"name":"wood","z":-2},{"y":-13,"x":10,"name":"wood","z":4},{"y":-13,"x":8,"name":"wood","z":2},{"y":-13,"x":6,"name":"wood","z":4},{"y":-13,"x":4,"name":"wood","z":6},{"y":-13,"x":8,"name":"wood","z":0},{"y":-13,"x":4,"name":"wood","z":4},{"y":-13,"x":2,"name":"wood","z":6},{"y":-13,"x":6,"name":"wood","z":2},{"y":-11,"x":8,"name":"wood","z":4},{"y":-11,"x":8,"name":"wood","z":2},{"y":-11,"x":6,"name":"wood","z":2},{"y":-11,"x":10,"name":"wood","z":2},{"y":-11,"x":10,"name":"wood","z":4},{"y":-11,"x":8,"name":"wood","z":0},{"y":-11,"x":4,"name":"wood","z":4},{"y":-11,"x":10,"name":"wood","z":0},{"y":-11,"x":6,"name":"wood","z":4},{"y":-9,"x":6,"name":"wood","z":2},{"y":15,"x":-2,"name":"wood","z":4},{"y":13,"x":2,"name":"wood","z":4},{"y":17,"x":2,"name":"wood","z":4},{"y":21,"x":2,"name":"wood","z":4},{"y":27,"x":2,"name":"wood","z":2},{"y":37,"x":2,"name":"wood","z":4},{"y":25,"x":2,"name":"wood","z":4},{"y":29,"x":2,"name":"wood","z":4},{"y":31,"x":2,"name":"wood","z":2},{"y":35,"x":2,"name":"wood","z":2},{"y":19,"x":2,"name":"wood","z":2},{"y":23,"x":2,"name":"wood","z":2},{"y":15,"x":2,"name":"wood","z":2},{"y":17,"x":-2,"name":"wood","z":2},{"y":19,"x":0,"name":"wood","z":4},{"y":21,"x":0,"name":"wood","z":2},{"y":25,"x":-2,"name":"wood","z":2},{"y":23,"x":-2,"name":"wood","z":4},{"y":27,"x":0,"name":"wood","z":4},{"y":29,"x":0,"name":"wood","z":2},{"y":33,"x":-2,"name":"wood","z":2},{"y":31,"x":-2,"name":"wood","z":4},{"y":35,"x":0,"name":"wood","z":4},{"y":37,"x":0,"name":"wood","z":2},{"y":-17,"x":-6,"name":"wood","z":2},{"y":-15,"x":-6,"name":"wood","z":2},{"y":-13,"x":-6,"name":"wood","z":2},{"y":-9,"x":8,"name":"wood","z":2},{"y":-17,"x":-6,"name":"wood","z":4},{"y":-15,"x":-6,"name":"wood","z":4},{"y":-9,"x":8,"name":"wood","z":4},{"y":-13,"x":-6,"name":"wood","z":4},{"y":-13,"x":-4,"name":"wood","z":0},{"y":-15,"x":12,"name":"wood","z":4},{"y":-15,"x":-4,"name":"wood","z":0},{"y":-13,"x":12,"name":"wood","z":4},{"y":-13,"x":12,"name":"wood","z":2},{"y":-17,"x":-4,"name":"wood","z":0},{"y":-15,"x":12,"name":"wood","z":2},{"y":-13,"x":12,"name":"wood","z":0},{"y":-13,"x":-2,"name":"wood","z":-2},{"y":-15,"x":12,"name":"wood","z":0},{"y":-15,"x":-2,"name":"wood","z":-2},{"y":-11,"x":12,"name":"wood","z":4},{"y":-11,"x":12,"name":"wood","z":2},{"y":-17,"x":-2,"name":"wood","z":-2},{"y":-11,"x":12,"name":"wood","z":0},{"y":-9,"x":8,"name":"wood","z":0}]', 'A', '2012-06-18 10:34:09', '2012-06-18 10:34:09')
    `)

	db.Query(`
        insert into test_results
        (id, moneyEarned, endedAt) values
        ('1', 10, '2024-01-05 00:00:00')
    `)
	db.Query(`
        insert into player_test_results (playerId, testResultId, wavesSurvived, diedTo) values
        ('1', '1', 10, 'Bombs From Above')
    `)
	db.Query(`
        insert into test_events
        (id, environment, difficulty, templateId, testResultId, startedAt, version) values
        ('1', 'lab', 'normal', '1', '1', '2024-01-01 00:00:00', '1.0.0')
    `)

	db.Query("insert into versions (value) values ('2.0.0')")
	db.Query(`
        insert into test_results
        (id, moneyEarned, endedAt) values
        ('2', 10, '2012-06-18 10:34:09')
    `)
	db.Query(`
        insert into player_test_results (playerId, testResultId, wavesSurvived, diedTo) values
        ('1', '2', 20, 'Bombs From Above')
    `)
	db.Query(`
        insert into test_events
        (id, environment, difficulty, templateId, testResultId, startedAt, version) values
        ('2', 'lab', 'normal', '1', '2', '2024-02-01 00:00:00', '2.0.0')
    `)

	db.Query(`
        insert into test_results
        (id, moneyEarned, endedAt) values
        ('3', 11, '2024-02-05 00:00:00')
    `)
	db.Query(`
        insert into player_test_results (playerId, testResultId, wavesSurvived, diedTo) values
        ('1', '3', 10, 'Bombs From Above')
    `)
	db.Query(`
        insert into test_events
        (id, environment, difficulty, templateId, testResultId, startedAt, version) values
        ('3', 'lab', 'normal', '1', '3', '2024-02-01 00:00:01', '2.0.0')
    `)

	db.Query(`
        insert into test_results
        (id, moneyEarned, endedAt) values
        ('4', 11, '2024-02-05 00:00:00')
    `)
	db.Query(`
        insert into player_test_results (playerId, testResultId, wavesSurvived, diedTo) values
        ('1', '4', 10, 'Acid')
    `)
	db.Query(`
        insert into test_events
        (id, environment, difficulty, templateId, testResultId, startedAt, version) values
        ('4', 'lab', 'normal', '1', '4', '2024-02-01 00:00:01', '2.0.0')
    `)

	// No test result
	db.Query(`
        insert into test_events
        (id, environment, difficulty, templateId, startedAt, version) values
        ('5', 'lab', 'normal', '1', '2024-02-20 00:00:00', '2.0.0')
    `)
}
