package insert

import (
	"database/sql"
	"log"
	"strings"

	/*
		"encoding/csv"
		"fmt"
		"os"
		"strconv"
		"strings"
	*/

	"github.com/LadySqrrl/Case-Management/spreadsheet"

	// _ "github.com/go-sql-driver/mysql" links to database
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

	rows := spreadsheet.ReadSheet("files/students.csv")

	for i := range rows {

		lastName := rows[i][0]
		firstName := rows[i][1]
		dob := rows[i][2]
		studentID := lastName + firstName + dob
		dobSplit := strings.Split(dob, "/")
		dob = dobSplit[2] + "-" + dobSplit[0] + "-" + dobSplit[1]
		grade := rows[i][3]
		annualDate := rows[i][6]
		annualDateSplit := strings.Split(annualDate, "/")
		annualDate = annualDateSplit[2] + "-" + annualDateSplit[0] + "-" + annualDateSplit[1]
		reevalDate := rows[i][7]
		reevalDateSplit := strings.Split(reevalDate, "/")
		reevalDate = reevalDateSplit[2] + "-" + reevalDateSplit[0] + "-" + reevalDateSplit[1]
		weighting := rows[i][5]
		caseManager := rows[i][8]
		providerID := rows[i][4] + rows[i][8]

		{
			query := "INSERT INTO students VALUES ('" + studentID + "', '" + lastName + "', '" + firstName + "', '" + dob + "', " + grade + ", '" + annualDate + "', '" + reevalDate + "', " + weighting + ", '" + caseManager + "', '" + providerID + "');"

			if _, err := db.Exec(query); err != nil {
				log.Fatal(err)
			}
		}
	}
}
