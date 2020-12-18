package utils

import "fmt"

const (
	JsonContentType      = "application/json"
	AllGrades            = "Successfully retrieved all grades"
	GradeInserted        = "Successfully inserted grade with ID 0"
	GradeDeleted         = "Successfully deleted grade with ID 0"
	GradeEdited          = "Successfully edited grade with ID 0"
	GradeRetrieved       = "Successfully retrieved the grade"
	NotFoundGrade        = "Grade not found"
	InvalidParameters    = "Invalid parameters"
	MissingParameters    = "Missing parameters"
	ValueShouldBeGreater = "Grade value should be greater than zero"
	DeleteFailed         = "Failed to delete the grade"
	EditFailed           = "Failed to edit the grade"
)

func ResultMessageAndData(message, data string) string {
	return fmt.Sprintf("{\"result\": {\"message\": \"%s\", \"data\": %s}}\n", message, data)
}

func ErrorMessage(message string) string {
	return fmt.Sprintf("{\"error\": \"%s\"}\n", message)
}
