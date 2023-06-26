package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Digital_technical_report_id       *int  `json:"digital_technical_report_id" default:""`
Digital_technical_report       Digital_technical_report  `json:"digital_technical_report" default:"" gorm:"foreignKey:Digital_technical_report_id"`
}