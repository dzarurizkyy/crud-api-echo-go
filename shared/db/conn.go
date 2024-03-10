package db

import (
	"crud_product/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Connect to database
func NewInstanceDB(conf config.Configuration) *sql.DB {
	db, err := sql.Open(
		"postgres", fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
	  	 conf.DBHost, conf.DBPort, conf.DBUsername, conf.DBPassword, conf.DBName));
	
	if (err != nil) {
		panic(err);
	};
	
	// -------------------------------------------

	err = db.Ping();

	if (err != nil) {
		panic(err);
	}

	return db;
}; 