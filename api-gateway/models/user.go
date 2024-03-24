package models

type Role int

const (
	ROLEOne Role = iota
	ROLETwo
)

type User struct {
	UserId   int64  `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}
