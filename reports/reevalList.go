package reports

import (
	"database/sql"
	"log"

	// links to database
	_ "github.com/go-sql-driver/mysql"
)

// Reeval holds information about reevaluations due
type Reeval struct {
	name       string
	reevalDate string
	annualDate string
}

// ReevalList creates a list of reeval due dates
func ReevalList() []Reeval {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3307)/cases?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT CONCAT(student_first_name, ' ', student_last_name), reeval_due, annual_due FROM students WHERE reeval_due < '2022-10-01';")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var reevals []Reeval
	for rows.Next() {
		var r Reeval

		err := rows.Scan(&r.name, &r.reevalDate, &r.annualDate)
		if err != nil {
			log.Fatal(err)
		}
		reevals = append(reevals, r)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return reevals
}
