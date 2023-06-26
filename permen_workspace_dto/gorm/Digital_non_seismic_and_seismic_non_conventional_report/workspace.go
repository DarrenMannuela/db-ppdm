package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Digital_non_seismic_and_seismic_non_conventional_report_id       *int  `json:"digital_non_seismic_and_seismic_non_conventional_report_id" default:""`
Digital_non_seismic_and_seismic_non_conventional_report       Digital_non_seismic_and_seismic_non_conventional_report  `json:"digital_non_seismic_and_seismic_non_conventional_report" default:"" gorm:"foreignKey:Digital_non_seismic_and_seismic_non_conventional_report_id"`
}