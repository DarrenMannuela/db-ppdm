package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Basin_id       *int  `json:"basin_id" default:""`
Basin       Basin  `json:"basin" default:"" gorm:"foreignKey:Basin_id"`
}