package reports

import (
	"database/sql"
	"log"

	// links to database
	_ "github.com/go-sql-driver/mysql"
)

// Reeval holds information about reevaluations due
type Reeval struct {
	name        string
	reevalDate  string
	meetingDate string
	providerID  string
}

// ReevalList creates a slice of Reevals
func ReevalList() []Reeval {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3307)/cases?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(
		`SELECT CONCAT(s.student_first_name, ' ', s.student_last_name), p.provider_name, s.reeval_due,
		CASE WHEN s.annual_due < s.reeval_due THEN s.annual_due ELSE s.reeval_due END AS "Meeting" 
		FROM students s
		INNER JOIN providers p 
		ON s.providerID = p.providerID
		WHERE s.reeval_due < '2022-10-01'
		ORDER BY s.reeval_due;`)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var reevals []Reeval
	for rows.Next() {
		var r Reeval

		err := rows.Scan(&r.name, &r.providerID, &r.reevalDate, &r.meetingDate)
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

// ReevalsByMonth creates a slice of Reevals due in a given month
func ReevalsByMonth() []Reeval {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3307)/cases?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	monthStart := "2021-02-01"
	monthEnd := "2021-03-01"

	rows, err := db.Query(
		`SELECT CONCAT(s.student_first_name, ' ', s.student_last_name), p.provider_name, s.reeval_due,
		CASE WHEN s.annual_due < s.reeval_due THEN s.annual_due ELSE s.reeval_due END AS "Meeting" 
		FROM students s
		INNER JOIN providers p 
		ON s.providerID = p.providerID
		WHERE s.reeval_due BETWEEN '` + monthStart + `' AND '` + monthEnd + `' ORDER BY s.reeval_due;`)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var reevals []Reeval
	for rows.Next() {
		var r Reeval

		err := rows.Scan(&r.name, &r.providerID, &r.reevalDate, &r.meetingDate)
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
