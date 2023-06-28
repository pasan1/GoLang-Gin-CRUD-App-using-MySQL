package main

import (
	"database/sql"

	"GoCRUDApplicationMySQL/app/routes"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/MyGoTestDB")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := routes.SetupRouter(db)
	r.Run(":8880")
}
