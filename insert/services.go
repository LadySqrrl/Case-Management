package insert

import (
	"database/sql"
	"log"

	"github.com/LadySqrrl/Case-Management/spreadsheet"
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

	rows := spreadsheet.ReadSheet("files/services.csv")

	for i := range rows {
		serviceType := rows[i][9]
		//building := rows[i][4]
		//providerName := rows
		providerID := rows[i][4] + rows[i][6]
		minutes := rows[i][10]
		frequency := rows[i][11]
		studentID := rows[i][0] + rows[i][1] + rows[i][2]

		{
			query := "INSERT INTO services (service_type, providerID, service_mins, service_frequency, studentID) VALUES ('" + serviceType + "', '" + providerID + "', " + minutes + ", '" + frequency + "', '" + studentID + "');"

			if _, err := db.Exec(query); err != nil {
				log.Fatal(err)
			}
		}
	}
}
