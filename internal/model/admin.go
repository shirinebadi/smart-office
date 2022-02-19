package model

type Admin struct {
	Username int
	Password string
}

func NewAdmin(username int, password string) *Admin {
	user := &Admin{Username: username, Password: password}

	return user
}

type AuthenticationInterface interface {
	AdminLogin(username int, password string) (Admin, error)
	AdminRegister(user *Admin) error
	UserLogin(username int, password string) (CentralUser, error)
}
