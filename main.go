package casemanagement

import "fmt"

func main() {
	providerID := "1234"
	//providerName := ""
	//distrct := ""
	//building := ""
	query := `INSERT INTO providers VALUES ("'` + providerID + `');`
	fmt.Println(query)
}
