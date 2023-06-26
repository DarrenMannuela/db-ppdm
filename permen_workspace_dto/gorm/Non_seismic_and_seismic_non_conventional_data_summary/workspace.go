package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Non_seismic_and_seismic_non_conventional_data_summary_id       *int  `json:"non_seismic_and_seismic_non_conventional_data_summary_id" default:""`
Non_seismic_and_seismic_non_conventional_data_summary       Non_seismic_and_seismic_non_conventional_data_summary  `json:"non_seismic_and_seismic_non_conventional_data_summary" default:"" gorm:"foreignKey:Non_seismic_and_seismic_non_conventional_data_summary_id"`
}