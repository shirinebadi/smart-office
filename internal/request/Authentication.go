package request

type Authentication struct {
	Username int    `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Card     string `json:"card"`
	Room     int    `json:"room"`
	Password string `json:"password"`
}

func NewAdmin(username int, password string) *Authentication {
	admin := &Authentication{Username: username, Password: password}

	return admin
}
