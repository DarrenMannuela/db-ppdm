package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Digital_2d_3d_seismic_report_id       *int  `json:"digital_2d_3d_seismic_report_id" default:""`
Digital_2d_3d_seismic_report       Digital_2d_3d_seismic_report  `json:"digital_2d_3d_seismic_report" default:"" gorm:"foreignKey:Digital_2d_3d_seismic_report_id"`
}