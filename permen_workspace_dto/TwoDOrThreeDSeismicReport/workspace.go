package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Two_d_or_three_d_seismic_report_id       *int  `json:"two_d_or_three_d_seismic_report_id" default:""`
}