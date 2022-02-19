package db

import (
	"github.com/shirinebadi/smart-office/internal/model"
	"gorm.io/gorm"
)

type Mydb struct {
	DB *gorm.DB
}

func (d *Mydb) AdminLogin(username int, password string) (model.Admin, error) {
	var stored model.Admin
	err := d.DB.Where(&model.Admin{Username: username, Password: password}).First(&stored).Error

	return stored, err
}

func (d *Mydb) AdminRegister(admin *model.Admin) error {
	return d.DB.Create(admin).Error
}

func (d *Mydb) GetAllActivites() (error, []model.Activity) {
	var stored []model.Activity
	err := d.DB.Where(&model.Activity{}).Find(&stored).Error

	return err, stored
}

func (d *Mydb) UserLogin(username int, password string) (model.CentralUser, error) {
	var stored model.CentralUser
	err := d.DB.Where(&model.CentralUser{Id: username, Password: password}).First(&stored).Error

	return stored, err
}

func (d *Mydb) GetUserLight(id int) (int, error) {
	var stored model.CentralUser
	err := d.DB.Where(&model.CentralUser{Id: id}).First(&stored).Error

	return stored.Light, err
}

func (d *Mydb) UserSearch(card string) (model.LocalUser, error) {
	var stored model.LocalUser
	err := d.DB.Where(&model.LocalUser{Card: card}).First(&stored).Error

	return stored, err
}

func (d *Mydb) SetUserLight(centralUser *model.CentralUser) error {
	return d.DB.Save(centralUser).Error
}

func (d *Mydb) UserRegister(localUser *model.LocalUser) error {
	return d.DB.Create(localUser).Error
}

func (d *Mydb) UserCentralRegister(centralUser *model.CentralUser) error {
	return d.DB.Create(centralUser).Error
}

func (d *Mydb) SetActivity(activity *model.Activity) error {
	return d.DB.Create(activity).Error
}

func (d *Mydb) RegisterOffice(office *model.Office) error {
	return d.DB.Create(office).Error
}

func (d *Mydb) UpdateLightsTime(office *model.Office) error {
	return d.DB.Save(office).Error
}

func (d *Mydb) GetOffice(id int) (*model.Office, error) {
	var stored model.Office
	err := d.DB.Where(&model.Office{Id: id}).First(&stored).Error

	return &stored, err
}
