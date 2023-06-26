package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Well_samples_id       *int  `json:"well_samples_id" default:""`
}