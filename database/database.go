package database

import (
	"database/sql"
	"log"
	"os"
	//	"github.com/lib/pq"
)

var dbURL string

func CreatDB() {
	db, err := ConnDB()
	if err != nil {
		log.Fatal("Comnnection Error")
	}
	createTb := ` 
	CREATE TABLE IF NOT EXISTS customers ( 
		id SERIAL PRIMARY KEY, 
		name   TEXT,
		email TEXT,
		status TEXT
	 );`

	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("Can't create table fatal error", err.Error())
	}
}

func ConnDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	// db, err := sql.Open("postgres", "postgres://dbuvonnf:R2v1DtbVijedmzNCTHnIWAWqtE_FYw8_@satao.db.elephantsql.com:5432/dbuvonnf")
	if err != nil {
		return nil, err
	}
	return db, nil

}

// func InitDB(*sql.DB, error) {
// 	{
// 		dbURL = os.Getenv("DATABASE_URL")
// 		if len(dbURL) == 0 {
// 			log.Fatal(" Environment DATABASE_URL is empty ")
// 		}

// 		db, err := sql.Open("postgres", dbURL)
// 		if err != nil {
// 			log.Fatal("Can't connect db", err.Error())
// 			return nil, err
// 		}
// 		defer db.Close()
// 		fmt.Print("1)database : ")
// 		return db, nil

// 	}
// }
