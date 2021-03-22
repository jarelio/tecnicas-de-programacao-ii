package utils

const (
	//Success constants
	GradesByStudent = "Successfully retrieved all the grades by student"
	AllGrades       = "Successfully retrieved all grades"
	GradeInserted   = "Successfully inserted grade with ID %s"
	GradeDeleted    = "Successfully deleted grade with ID %s"
	GradeEdited     = "Successfully edited grade with ID %s"
	GradeRetrieved  = "Successfully retrieved the grade"

	//Errors constants
	GradeNotFound        = "Grade not found"
	InvalidParameters    = "Invalid parameters"
	MissingParameters    = "Missing parameters"
	ValueShouldBeGreater = "Grade value should be greater than zero"
	DeleteFailed         = "Failed to delete the grade"
	EditFailed           = "Failed to edit the grade"
)
