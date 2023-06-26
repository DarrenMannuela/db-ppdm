package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
T2D_seismic_field_digital_data_id       *int  `json:"t2d_seismic_field_digital_data_id" default:""`
T2D_seismic_field_digital_data       T2D_seismic_field_digital_data  `json:"t2d_seismic_field_digital_data" default:"" gorm:"foreignKey:T2D_seismic_field_digital_data_id"`
}