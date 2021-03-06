package database

type User struct {
	Username string `bson:"username",json:"username"`
	Email    string `bson:"email",json:"email"`
	Password string `bson:"password",json:"password"`
}

type AuthResponseStatus int8

const (
	Success AuthResponseStatus =  iota
	WrongPassword
	WrongUsername
	DBError
)

type AuthResponse struct {
	Status AuthResponseStatus `json:"status"`
	JWT    string             `json:"jwt"`
}

type Database interface {
	AddUser(User) error
	RemoveUser(User) error
	Authenticate(User) (AuthResponse, error)
	Disconnect() error
}
