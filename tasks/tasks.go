// Package tasks organizes task lists to be used for planning
package tasks

import (
	"github.com/LadySqrrl/Case-Management/reports"
)

// TaskList is a struct that holds information about tasks and whether they have been completed
type TaskList struct {
	studentName    string
	dueBy          string
	noticeSent     bool
	reviewDone     bool
	obsDone        bool
	teacherIntDone bool
	studentIntDone bool
	parentIntDone  bool
	testingDone    bool
	rPageDone      bool
	pwnDone        bool
}

// CreateList creates a task list for a reevaluation
func CreateList(reeval reports.Reeval) []TaskList {
	var tasks []TaskList
	/*
		this.studentName = reeval.name
		this.dueBy = reeval.meetingDate */

	return tasks
}
