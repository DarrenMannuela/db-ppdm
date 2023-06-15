package dto

type Well_samples struct {
	Ba_long_name     *string `json:"ba_long_name" default:""`
	Ba_type          *string `json:"ba_type" default:""`
	Area_id          *string `json:"area_id" default:""`
	Area_type        *string `json:"area_type" default:""`
	Field_name       *string `json:"field_name" default:""`
	Well_name        *string `json:"well_name" default:""`
	Uwi              *string `json:"uwi" default:""`
	Sample_type      *string `json:"sample_type" default:""`
	Sample_num       *string `json:"sample_num" default:""`
	Sample_count     *string `json:"sample_count" default:""`
	Top_md           *string `json:"top_md" default:""`
	Top_md_ouom      *string `json:"top_md_ouom" default:""`
	Base_md          *string `json:"base_md" default:""`
	Base_md_ouom     *string `json:"base_md_ouom" default:""`
	Study_type       *string `json:"study_type" default:""`
	Data_store_name  *string `json:"data_store_name" default:""`
	Data_store_type  *string `json:"data_store_type" default:""`
	Location_id      *string `json:"location_id" default:""`
	Remark           *string `json:"remark" default:""`
	Source           *string `json:"source" default:""`
	Qc_status        *string `json:"qc_status" default:""`
	Checked_by_ba_id *string `json:"checked_by_ba_id" default:""`
}
