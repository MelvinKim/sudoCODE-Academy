package dto

// StudentCreationPayload
type StudentCreationPayload struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// GetStudentPayload
type GetStudentPayload struct {
	Email string `json:"email"`
}
