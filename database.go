package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	_ "github.com/go-sql-driver/mysql" // Import mysql library
	"errors"
)

var db *sql.DB

func InitDB(dataSourceName string, driver string) error {
	var err error

	db, err = sql.Open(driver, dataSourceName)
	if err != nil {
		return err
	}

	return db.Ping()
}

func ChangeElement(query string, args ...interface{}) (sql.Result, error) {

	//Prepare query
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, errors.New("Error to prepare query")
	}

	result, err := stmt.Exec(args...)
	if err != nil {
		return nil, errors.New("Error to exec query")
	}
	return result, nil
}

func GetElementsWithValue(query string, args ...interface{}) (*sql.Rows, error) {

	//Prepare query
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, errors.New("Error to prepare query")
	}

	row, err := stmt.Query(query, args)
	if err != nil {
		return nil, err
	}

	return row, nil
}

func GetElements(query string) (*sql.Rows, error) {

	//Prepare query
	row, err := db.Query(query)
	if err != nil {
		return nil, errors.New("Error to prepare query")
	}
	return row, nil

}

func Close() error {
	return db.Close()
}