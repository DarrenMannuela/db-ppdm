package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Digital_maps_and_technical_drawing_id       *int  `json:"digital_maps_and_technical_drawing_id" default:""`
Digital_maps_and_technical_drawing       Digital_maps_and_technical_drawing  `json:"digital_maps_and_technical_drawing" default:"" gorm:"foreignKey:Digital_maps_and_technical_drawing_id"`
}