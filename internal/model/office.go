package model

type Office struct {
	Id            int
	LightsOnTime  int
	LightsOffTime int
}

type OfficeInterface interface {
	RegisterOffice(office *Office) error
	GetOffice(id int) (*Office, error)
	UpdateLightsTime(office *Office) error
}
