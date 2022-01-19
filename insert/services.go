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

// Services inserts data into the services table
func Services() {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3307)/cases?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	//rows := readOrders("students.csv")

	serviceType := ""
	providerID := ""
	minutes := ""
	frequency := ""
	studentID := ""

	{
		query := `INSERT INTO services (service_type, student_first_name, dob, grade, annual_date, reeval_date, weighting, case_manager, roster_teacherID) VALUES ('` + serviceType + `', ` + providerID + `, ` + minutes + `, '` + frequency + `', ` + studentID + `);`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}
}
