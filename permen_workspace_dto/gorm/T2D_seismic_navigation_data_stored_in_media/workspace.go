package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
T2D_seismic_navigation_data_stored_in_media_id       *int  `json:"t2d_seismic_navigation_data_stored_in_media_id" default:""`
T2D_seismic_navigation_data_stored_in_media       T2D_seismic_navigation_data_stored_in_media  `json:"t2d_seismic_navigation_data_stored_in_media" default:"" gorm:"foreignKey:T2D_seismic_navigation_data_stored_in_media_id"`
}