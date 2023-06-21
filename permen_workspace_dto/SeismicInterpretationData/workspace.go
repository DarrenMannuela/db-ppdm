package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Seismic_interpretation_data_id       *int  `json:"seismic_interpretation_data_id" default:""`
}