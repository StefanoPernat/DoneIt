package main

import (
	"fmt"

	"github.com/StefanoPernat/DoneIt/config"
)

func main() {
	rows, err := config.DB.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	var id int
	var username string

	for rows.Next() {
		err = rows.Scan(&id, &username)

		fmt.Println(id, username)
	}

}
