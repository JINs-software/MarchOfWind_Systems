package models

type User struct {
	Username string
	Password string
	Token    string
}

var Users = map[string]User{
	"john": {"john", "password123", "sample-jwt-token"},
}
