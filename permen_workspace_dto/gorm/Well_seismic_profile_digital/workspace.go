package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Well_seismic_profile_digital_id       *int  `json:"well_seismic_profile_digital_id" default:""`
Well_seismic_profile_digital       Well_seismic_profile_digital  `json:"well_seismic_profile_digital" default:"" gorm:"foreignKey:Well_seismic_profile_digital_id"`
}