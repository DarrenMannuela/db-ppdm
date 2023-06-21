package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Well_core_sample_id       *int  `json:"well_core_sample_id" default:""`
}