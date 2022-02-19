package model

type LocalUser struct {
	Id   int `gorm:"primaryKey;auto_increment"`
	Card string
	Room int
}

type CentralUser struct {
	Id       int `gorm:"primaryKey;auto_increment"`
	Password string
	Light    int
	Office   string
	Room     int
}

func NewLocalUser(id int, card string, room int) *LocalUser {
	user := &LocalUser{Id: id, Card: card, Room: room}

	return user
}

func NewCentralUser(password string, light int, office string, room int) *CentralUser {
	user := &CentralUser{Password: password, Light: light, Office: office, Room: room}

	return user
}

type UserInterface interface {
	UserRegister(*LocalUser) error
	UserCentralRegister(*CentralUser) error
	UserSearch(card string) (LocalUser, error)
	GetUserLight(id int) (int, error)
	SetUserLight(*CentralUser) error
}
