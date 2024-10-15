package models

// User struct is a representation of user entity
type User struct {
	Name  string   `json:"name" validate:"required"`        // Name of the user, a required field
	Email string   `json:"email" validate:"required,email"` // Email ID of the user, should be in valid email format
	Roles []string `json:"roles" validate:"required"`       // Roles assigned to the user, a required field
}
