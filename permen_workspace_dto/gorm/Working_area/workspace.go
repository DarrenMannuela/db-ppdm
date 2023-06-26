package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Working_area_id       *int  `json:"working_area_id" default:""`
Working_area       Working_area  `json:"working_area" default:"" gorm:"foreignKey:Working_area_id"`
}