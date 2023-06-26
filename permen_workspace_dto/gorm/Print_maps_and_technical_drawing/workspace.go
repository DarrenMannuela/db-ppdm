package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Print_maps_and_technical_drawing_id       *int  `json:"print_maps_and_technical_drawing_id" default:""`
Print_maps_and_technical_drawing       Print_maps_and_technical_drawing  `json:"print_maps_and_technical_drawing" default:"" gorm:"foreignKey:Print_maps_and_technical_drawing_id"`
}