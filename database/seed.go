package database

import "log"

type User struct {
	ID       int
	username string
	password string
}

func Seed() {
	db := Connect()
	defer db.Close()

	res, err := db.Query("SELECT * FROM users")
	if err != nil {
		println("errorsdkafj")
		panic(err.Error())
	}

	for res.Next() {
		var user User
		err = res.Scan(&user.ID, &user.username, &user.password)
		if err != nil {
			panic(err.Error())
		}
		log.Printf("user: %v", user)
	}
}
