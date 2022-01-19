package insert

import (
	"database/sql"
	"encoding/csv"
	"log"
	"os"

	// links to database
	_ "github.com/go-sql-driver/mysql"
)

func readSheet(name string) [][]string {
	f, err := os.Open(name)

	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", name, err.Error())
	}

	defer f.Close()

	r := csv.NewReader(f)

	r.Comma = ','

	rows, err := r.ReadAll()

	if err != nil {
		log.Fatalln("Cannot read CSV data:", err.Error())
	}

	return rows

}

// Providers Inserts data into the providers table
func Providers() {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3307)/cases?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	rows := readSheet("files/providers.csv")

	for i := range rows {

		district := rows[i][0]
		building := rows[i][1]
		providerName := rows[i][2]

		{
			query := `INSERT INTO providers (provider_name, district, building) VALUES ('` + providerName + `', '` + district + `', '` + building + `');`

			if _, err := db.Exec(query); err != nil {
				log.Fatal(err)
			}
		}
	}
}
