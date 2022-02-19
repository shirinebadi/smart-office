package model

import "time"

type Activity struct {
	Index    int `gorm:"primaryKey;auto_increment"`
	Id       int
	Office   string
	DateTime time.Time
	Type     string
}

type ActivityInterface interface {
	GetAllActivites() (error, []Activity)
	SetActivity(activity *Activity) error
}
