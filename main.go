package main

// import "github.com/LadySqrrl/Case-Management/insert" and "github.com/LadySqrrl/Case-Management/insert"
import (
	//"github.com/LadySqrrl/Case-Management/insert"
	//"github.com/LadySqrrl/Case-Management/setup"
	"github.com/LadySqrrl/Case-Management/reports"
)

func main() {

	list := reports.ReevalList()
	reports.PrintReevals(list)

}
