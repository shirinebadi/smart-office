package request

type Activity struct {
	ID int `json:"id"`
}

func NewActivity(id int) *Activity {
	return &Activity{
		ID: id,
	}
}
