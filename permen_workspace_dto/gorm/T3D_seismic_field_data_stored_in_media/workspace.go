package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
T3D_seismic_field_data_stored_in_media_id       *int  `json:"t3d_seismic_field_data_stored_in_media_id" default:""`
T3D_seismic_field_data_stored_in_media       T3D_seismic_field_data_stored_in_media  `json:"t3d_seismic_field_data_stored_in_media" default:"" gorm:"foreignKey:T3D_seismic_field_data_stored_in_media_id"`
}