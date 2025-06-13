package users

// User represents a user without password field
// Add more fields as needed
// Adjust struct tags to match your DB columns
// Example: id, username, email

type User struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Fullname string `json:"fullname" db:"fullname"`
	Password string `json:"-" db:"password"`
}
