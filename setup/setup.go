package setup

import (
	"database/sql"
	"log"

	// links to database
	_ "github.com/go-sql-driver/mysql"
)

// SetUp sets up the database. Only needed on initial execution
func SetUp() {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3307)/cases?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Drop providers, students, and services tables if they already exists
	/*
		{

			query := `DROP TABLE services;`

			if _, err := db.Exec(query); err != nil {
				log.Fatal(err)
			}
		}*/
	{
		query := `DROP TABLE students;`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}
	{

		query := `DROP TABLE providers;`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{ // Create a new table
		query := `
		CREATE TABLE providers (
			providerID      VARCHAR(100)		NOT NULL    ,
			provider_name   VARCHAR(50)         NOT NULL    ,
			district        VARCHAR(100)                    ,
			building        VARCHAR(50)                     ,
			CONSTRAINT pk_providerID PRIMARY KEY (providerID)
		);`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{ // Create a new table
		query := `
		CREATE TABLE students (
			studentID           VARCHAR(100)		  NOT NULL      ,
			student_last_name   VARCHAR(50)                         ,
			student_first_name  VARCHAR(50)                         ,
			dob                 DATE                                ,
			grade               INT                                 ,
			annual_due          DATE                                ,
			reeval_due          DATE                                ,
			weighting           DECIMAL(3, 2)                       ,
			case_manager        VARCHAR(50)                         ,
			providerID    		VARCHAR(100)                        ,
			CONSTRAINT pk_student PRIMARY KEY (studentID)           ,
			CONSTRAINT fk_provider FOREIGN KEY (providerID) REFERENCES providers (providerID)
		);`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{ // Create a new table
		query := `
		CREATE TABLE services (
			serviceID               INT AUTO_INCREMENT  NOT NULL    ,
			service_type            VARCHAR (5)                     ,
			providerID              VARCHAR(100)                    ,
			service_mins            INT                             ,
			service_frequency       VARCHAR(10)                     ,
			studentID               VARCHAR(100)                    ,
			CONSTRAINT pk_service PRIMARY KEY (serviceID),
			CONSTRAINT fk_svc_provider FOREIGN KEY (providerID) REFERENCES providers (providerID),
			CONSTRAINT fk_svc_student FOREIGN KEY (studentID) REFERENCES students (studentID)
		);`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}
}
