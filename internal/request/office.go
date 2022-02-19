package request

type Office struct {
	ID int `json:"id"`
}

type Lights struct {
	ID            int `json:"id"`
	LightsOnTime  int `json:"lightsontime"`
	LightsOffTime int `json:"lightsofftime"`
}

type Value struct {
	ID    int `json:"id"`
	Light int `json:"light"`
}
