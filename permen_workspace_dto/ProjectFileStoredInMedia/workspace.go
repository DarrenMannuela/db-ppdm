package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Project_file_stored_in_media_id       *int  `json:"project_file_stored_in_media_id" default:""`
}