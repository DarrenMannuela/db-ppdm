package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
T3D_seismic_summary_id       *int  `json:"t3d_seismic_summary_id" default:""`
T3D_seismic_summary       T3D_seismic_summary  `json:"t3d_seismic_summary" default:"" gorm:"foreignKey:T3D_seismic_summary_id"`
}