package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Digital_well_report_id       *int  `json:"digital_well_report_id" default:""`
Digital_well_report       Digital_well_report  `json:"digital_well_report" default:"" gorm:"foreignKey:Digital_well_report_id"`
}