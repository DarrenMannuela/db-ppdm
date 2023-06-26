package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Well_data_id       *int  `json:"well_data_id" default:""`
Well_data       Well_data  `json:"well_data" default:"" gorm:"foreignKey:Well_data_id"`
}