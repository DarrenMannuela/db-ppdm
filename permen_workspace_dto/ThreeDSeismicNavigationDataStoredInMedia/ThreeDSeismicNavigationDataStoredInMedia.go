package dto

type Three_d_seismic_navigation_data_stored_in_media struct{

Ba_long_name       *string   `json:"ba_long_name" default:""`
Ba_type       *string   `json:"ba_type" default:""`
Area_id       *string   `json:"area_id" default:""`
Area_type       *string   `json:"area_type" default:""`
Acqtn_survey_name       *string   `json:"acqtn_survey_name" default:""`
Seis_dimension       *string   `json:"seis_dimension" default:""`
Shot_by       *string   `json:"shot_by" default:""`
Process_date       *string   `json:"process_date" default:""`
Digital_format       *string   `json:"digital_format" default:""`
Media_type       *string   `json:"media_type" default:""`
Tape_number       *string   `json:"tape_number" default:""`
Data_store_name       *string   `json:"data_store_name" default:""`
Data_store_type       *string   `json:"data_store_type" default:""`
Location_id       *string   `json:"location_id" default:""`
Remark       *string   `json:"remark" default:""`
Source       *string   `json:"source" default:""`
Qc_status       *string   `json:"qc_status" default:""`
Checked_by_ba_id       *string   `json:"checked_by_ba_id" default:""`
}