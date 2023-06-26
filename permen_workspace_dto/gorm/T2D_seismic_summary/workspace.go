package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
T2D_seismic_summary_id       *int  `json:"t2d_seismic_summary_id" default:""`
T2D_seismic_summary       T2D_seismic_summary  `json:"t2d_seismic_summary" default:"" gorm:"foreignKey:T2D_seismic_summary_id"`
}