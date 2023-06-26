package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
T2D_3d_seismic_report_id       *int  `json:"t2d_3d_seismic_report_id" default:""`
T2D_3d_seismic_report       T2D_3d_seismic_report  `json:"t2d_3d_seismic_report" default:"" gorm:"foreignKey:T2D_3d_seismic_report_id"`
}