package reports

import (
	"database/sql"
	"log"

	// links to database
	_ "github.com/go-sql-driver/mysql"
)

// ReevalList creates a list of reeval due dates
func ReevalList() []string {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3307)/cases?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	var reevalList []string
	reevalList[0] = ""

	return reevalList
}
