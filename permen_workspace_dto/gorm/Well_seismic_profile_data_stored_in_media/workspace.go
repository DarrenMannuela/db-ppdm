package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Well_seismic_profile_data_stored_in_media_id       *int  `json:"well_seismic_profile_data_stored_in_media_id" default:""`
Well_seismic_profile_data_stored_in_media       Well_seismic_profile_data_stored_in_media  `json:"well_seismic_profile_data_stored_in_media" default:"" gorm:"foreignKey:Well_seismic_profile_data_stored_in_media_id"`
}