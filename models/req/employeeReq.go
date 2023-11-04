package req

// import "time"

type Employee struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	// Birthday *time.Time `json:"birthday" validate:"required"`
}
