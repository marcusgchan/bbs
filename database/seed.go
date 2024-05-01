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
	db.Query("insert into users (username, password) values (?, ?)", os.Getenv("USERNAME"), hashedPassword)

	db.Query("insert into versions (value) values ('1.0.0')")

	db.Query("insert into players (id, name) values ('1', 'marcus')")

	db.Query(`
        insert into templates (id, playerId, data, name, createdAt, updatedAt) values
        ('1', '1', '', 'A', '2012-06-18 10:34:09', '2012-06-18 10:34:09')
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

	// No test result
	db.Query(`
        insert into player_test_results (playerId, testResultId, wavesSurvived, diedTo) values
        ('1', '4', 10, 'Bombs From Above')
    `)
	db.Query(`
        insert into test_events
        (id, environment, difficulty, templateId, startedAt, version) values
        ('4', 'lab', 'normal', '1', '2024-02-20 00:00:00', '2.0.0')
    `)
}
