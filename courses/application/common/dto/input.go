package dto

// StudentCreationPayload
type StudentCreationPayload struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// CourseCreationPayload
type CourseCreationPayload struct {
	Title       string `json:"title" `
	Price       uint   `json:"price"`
	Description string `json:"description"`
	Instructor  string `json:"instructor"`
	Category    string `json:"category"`
}

// StudentCourseAssigningPayload
type StudentCourseAssigningPayload struct {
	Email       string `json:"email" `
	CourseTitle string `json:"course_title"`
}

// GetStudentPayload
type GetStudentPayload struct {
	Email string `json:"email"`
}

// GetStudentPayload
type GetCoursePayload struct {
	CourseTitle string `json:"email"`
}
