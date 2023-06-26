package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Print_well_report_id       *int  `json:"print_well_report_id" default:""`
Print_well_report       Print_well_report  `json:"print_well_report" default:"" gorm:"foreignKey:Print_well_report_id"`
}