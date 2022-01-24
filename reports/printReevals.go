package reports

import "fmt"

// PrintReevals prints the list of reevals due
func PrintReevals(list []Reeval) {
	//fmt.Println(list)
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i].name, " has a reeval due, the meeting is due by ", list[i].meetingDate, " and the reeval date is ", list[i].reevalDate, ". The roster teacher is ", list[i].providerID)
	}
}
