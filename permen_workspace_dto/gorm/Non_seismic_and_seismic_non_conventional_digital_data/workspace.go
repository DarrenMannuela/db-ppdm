package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Non_seismic_and_seismic_non_conventional_digital_data_id       *int  `json:"non_seismic_and_seismic_non_conventional_digital_data_id" default:""`
Non_seismic_and_seismic_non_conventional_digital_data       Non_seismic_and_seismic_non_conventional_digital_data  `json:"non_seismic_and_seismic_non_conventional_digital_data" default:"" gorm:"foreignKey:Non_seismic_and_seismic_non_conventional_digital_data_id"`
}