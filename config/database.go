package config

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

// import (
// 	"database/sql"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func ConnectDB() (*sql.DB, error) {
// 	dbDriver := "mysql"
// 	dbUser := "root"
// 	dbPass := ""
// 	dbName := "go_products"

// 	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
// 	return db, err
// }

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:@/go_products")
	if err != nil {
		panic(err)
	}
	log.Println("Database Connected")
	DB = db
}
