package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Digital_project_file_id       *int  `json:"digital_project_file_id" default:""`
Digital_project_file       Digital_project_file  `json:"digital_project_file" default:"" gorm:"foreignKey:Digital_project_file_id"`
}