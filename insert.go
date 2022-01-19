package casemanagement

import (
	"database/sql"
	"fmt"
	"log"
)

// InsertProviders Inserts data into the providers table
func InsertProviders() {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3307)/cases?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	providerID := "1234"
	//providerName := ""
	//distrct := ""
	//building := ""
	query := `INSERT INTO providers VALUES ("'` + providerID + `');`
	fmt.Println(query)
}

/*
	{ // Drop providers, students, and services tables if they already exists
		query := `
				INSERT INTO providers VALUES ("'` + providerID + `');`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}
}


// InsertStudents inserts data into the students table
func InsertStudents() {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3307)/cases?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	{ // Drop providers, students, and services tables if they already exists
		query := `
				DROP TABLE services;
				DROP TABLE students;
				DROP TABLE providers;
				`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{ // Create a new table
		query := `
		CREATE TABLE providers (
			providerID      INT AUTO_INCREMENT  NOT NULL    ,
			provider_name   VARCHAR(50)         NOT NULL    ,
			district        VARCHAR(100)                    ,
			building        VARCHAR(50)                     ,
			CONSTRAINT pk_provider PRIMARY KEY (providerID)
		);`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}
}

// InsertServices insers info into the services table
func InsertServices() {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3307)/cases?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	{ // Drop providers, students, and services tables if they already exists
		query := `
				DROP TABLE services;
				DROP TABLE students;
				DROP TABLE providers;
				`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{ // Create a new table
		query := `
		CREATE TABLE providers (
			providerID      INT AUTO_INCREMENT  NOT NULL    ,
			provider_name   VARCHAR(50)         NOT NULL    ,
			district        VARCHAR(100)                    ,
			building        VARCHAR(50)                     ,
			CONSTRAINT pk_provider PRIMARY KEY (providerID)
		);`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}
}
*/
