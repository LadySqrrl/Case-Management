package insert

import (
	"database/sql"
	"log"

	/*
		"encoding/csv"
		"fmt"
		"os"
		"strconv"
		"strings"
	*/

	// links to database
	_ "github.com/go-sql-driver/mysql"
)

// Students inserts data into the students table
func Students() {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3307)/cases?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	//rows := readOrders("students.csv")

	lastName := ""
	firstName := ""
	dob := ""
	grade := ""
	annualDate := ""
	reevalDate := ""
	weighting := ""
	caseManager := ""
	rosterTeacher := ""

	{
		query := "INSERT INTO students (student_last_name, student_first_name, dob, grade, annual_date, reeval_date, weighting, case_manager, roster_teacherID) VALUES ('" + lastName + "', '" + firstName + "', '" + dob + "', " + grade + ", '" + annualDate + "', '" + reevalDate + "', " + weighting + ", '" + weighting + ", '" + caseManager + "', " + rosterTeacher + ");"

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}
}
