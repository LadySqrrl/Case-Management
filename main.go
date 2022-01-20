package main

//import "github.com/LadySqrrl/Case-Management/insert" and "github.com/LadySqrrl/Case-Management/insert"
import (
	"github.com/LadySqrrl/Case-Management/insert"
	"github.com/LadySqrrl/Case-Management/setup"
)

func main() {

	setup.SetUp()
	insert.Providers()
	insert.Students()
}
