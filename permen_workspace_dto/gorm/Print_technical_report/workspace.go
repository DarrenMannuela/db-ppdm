package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Print_technical_report_id       *int  `json:"print_technical_report_id" default:""`
Print_technical_report       Print_technical_report  `json:"print_technical_report" default:"" gorm:"foreignKey:Print_technical_report_id"`
}