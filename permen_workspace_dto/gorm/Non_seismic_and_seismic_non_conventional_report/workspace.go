package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Non_seismic_and_seismic_non_conventional_report_id       *int  `json:"non_seismic_and_seismic_non_conventional_report_id" default:""`
Non_seismic_and_seismic_non_conventional_report       Non_seismic_and_seismic_non_conventional_report  `json:"non_seismic_and_seismic_non_conventional_report" default:"" gorm:"foreignKey:Non_seismic_and_seismic_non_conventional_report_id"`
}