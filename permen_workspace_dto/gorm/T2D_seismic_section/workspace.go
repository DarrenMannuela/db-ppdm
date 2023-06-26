package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
T2D_seismic_section_id       *int  `json:"t2d_seismic_section_id" default:""`
T2D_seismic_section       T2D_seismic_section  `json:"t2d_seismic_section" default:"" gorm:"foreignKey:T2D_seismic_section_id"`
}