package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Non_seismic_and_seismic_non_conventional_data_stored_in_media_id       *int  `json:"non_seismic_and_seismic_non_conventional_data_stored_in_media_id" default:""`
Non_seismic_and_seismic_non_conventional_data_stored_in_media       Non_seismic_and_seismic_non_conventional_data_stored_in_media  `json:"non_seismic_and_seismic_non_conventional_data_stored_in_media" default:"" gorm:"foreignKey:Non_seismic_and_seismic_non_conventional_data_stored_in_media_id"`
}