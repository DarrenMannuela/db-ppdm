package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Digital_image_well_log_id       *int  `json:"digital_image_well_log_id" default:""`
Digital_image_well_log       Digital_image_well_log  `json:"digital_image_well_log" default:"" gorm:"foreignKey:Digital_image_well_log_id"`
}