package dto

type Workspace struct{

Id       int  `json:"id" default:""`
Afe_number       *int  `json:"afe_number" default:""`
Outcrop_sample_id       *int  `json:"outcrop_sample_id" default:""`
}