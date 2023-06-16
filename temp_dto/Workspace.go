package dto

type Workspace struct {
	Id         int  `json:"id" default:"1"`
	Afe_number *int `json:"afe_number" default:"1"`
	Pwr_id     int  `json:"pwr_id" default:"source"`
}
