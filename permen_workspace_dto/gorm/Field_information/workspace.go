package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Field_information_id       *int  `json:"field_information_id" default:""`
Field_information       Field_information  `json:"field_information" default:"" gorm:"foreignKey:Field_information_id"`
}