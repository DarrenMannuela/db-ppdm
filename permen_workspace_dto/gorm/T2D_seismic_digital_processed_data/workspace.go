package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
T2D_seismic_digital_processed_data_id       *int  `json:"t2d_seismic_digital_processed_data_id" default:""`
T2D_seismic_digital_processed_data       T2D_seismic_digital_processed_data  `json:"t2d_seismic_digital_processed_data" default:"" gorm:"foreignKey:T2D_seismic_digital_processed_data_id"`
}