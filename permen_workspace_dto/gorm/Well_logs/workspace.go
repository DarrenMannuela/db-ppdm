package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Well_logs_id       *int  `json:"well_logs_id" default:""`
Well_logs       Well_logs  `json:"well_logs" default:"" gorm:"foreignKey:Well_logs_id"`
}