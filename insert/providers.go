package insert

import (
	"database/sql"
	"log"

	"github.com/LadySqrrl/Case-Management/spreadsheet"

	// links to database
	_ "github.com/go-sql-driver/mysql"
)

// Providers Inserts data into the providers table
func Providers() {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3307)/cases?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	rows := spreadsheet.ReadSheet("files/providers.csv")

	for i := range rows {

		district := rows[i][1]
		building := rows[i][0]
		providerName := rows[i][2]
		providerID := providerName + building

		{
			query := `INSERT INTO providers VALUES ('` + providerID + `', '` + providerName + `', '` + district + `', '` + building + `');`

			if _, err := db.Exec(query); err != nil {
				log.Fatal(err)
			}
		}
	}
}
