package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Project_file_stored_in_media_id       *int  `json:"project_file_stored_in_media_id" default:""`
Project_file_stored_in_media       Project_file_stored_in_media  `json:"project_file_stored_in_media" default:"" gorm:"foreignKey:Project_file_stored_in_media_id"`
}