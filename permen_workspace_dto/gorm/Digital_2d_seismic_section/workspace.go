package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Digital_2d_seismic_section_id       *int  `json:"digital_2d_seismic_section_id" default:""`
Digital_2d_seismic_section       Digital_2d_seismic_section  `json:"digital_2d_seismic_section" default:"" gorm:"foreignKey:Digital_2d_seismic_section_id"`
}