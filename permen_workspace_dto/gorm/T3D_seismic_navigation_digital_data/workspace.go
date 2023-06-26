package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
T3D_seismic_navigation_digital_data_id       *int  `json:"t3d_seismic_navigation_digital_data_id" default:""`
T3D_seismic_navigation_digital_data       T3D_seismic_navigation_digital_data  `json:"t3d_seismic_navigation_digital_data" default:"" gorm:"foreignKey:T3D_seismic_navigation_digital_data_id"`
}