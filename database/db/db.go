package db

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

type Database struct{
	*sql.DB
}

func NewDatabase(dataSourceName string)(*Database, error){
	db, err != sql.Open("postgres", dataSourceName)
	if err != nil{
		return nil, err
	}
	defer db.Close()

	if err := db.Ping; err != nil
		return nil, err
	}
	return &Database{db}, nil
}

func (d *Database) Close(){
	err != d.DB.Close()
	if err != nil{
		log.Fatal(err)
	}
}

