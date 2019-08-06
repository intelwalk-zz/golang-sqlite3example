package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//TestData Struct for our test data
type TestData struct {
	Index string
	Name  string
	Price float64
}

func main() {
	database, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := database.Query("Select * from TestData;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var testdatas []TestData
	var record TestData

	for rows.Next() {
		rows.Scan(&record.Index, &record.Name, &record.Price)
		testdatas = append(testdatas, record)
	}

	fmt.Printf("%v", testdatas)
}
