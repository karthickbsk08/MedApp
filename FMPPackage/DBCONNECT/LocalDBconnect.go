package DBCONNECT

import (
	"database/sql"
	"fmt"
	"log"
)

func LocalDBConnect() (*sql.DB, error) {
	log.Println("LocalDBConnect(+)")

	lConnString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "root", "192.168.2.5", 3306, "karthickb")

	db, err := sql.Open("mysql", lConnString)
	if err != nil {
		log.Println("Open connection Failed :", err.Error())
		return db, err
	}

	log.Println("LocalDBConnect(-)")
	return db, nil

}
