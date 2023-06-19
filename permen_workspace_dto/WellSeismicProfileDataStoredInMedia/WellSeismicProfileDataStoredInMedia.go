package dto

type Well_seismic_profile_data_stored_in_media struct{

Ba_long_name       *string   `json:"ba_long_name" default:""`
Ba_type       *string   `json:"ba_type" default:""`
Area_id       *string   `json:"area_id" default:""`
Area_type       *string   `json:"area_type" default:""`
Field_name       *string   `json:"field_name" default:""`
Well_name       *string   `json:"well_name" default:""`
Alias_long_name       *string   `json:"alias_long_name" default:""`
Uwi       *string   `json:"uwi" default:""`
Trip_date       *string   `json:"trip_date" default:""`
Survey_company_ba_id       *string   `json:"survey_company_ba_id" default:""`
Top_depth       *string   `json:"top_depth" default:""`
Top_depth_ouom       *string   `json:"top_depth_ouom" default:""`
Base_depth       *string   `json:"base_depth" default:""`
Base_depth_ouom       *string   `json:"base_depth_ouom" default:""`
Checkshot_survey_type       *string   `json:"checkshot_survey_type" default:""`
Vsp_type       *string   `json:"vsp_type" default:""`
Digital_format       *string   `json:"digital_format" default:""`
Media_type       *string   `json:"media_type" default:""`
Data_store_name       *string   `json:"data_store_name" default:""`
Data_store_type       *string   `json:"data_store_type" default:""`
Location_id       *string   `json:"location_id" default:""`
Remark       *string   `json:"remark" default:""`
Source       *string   `json:"source" default:""`
Qc_status       *string   `json:"qc_status" default:""`
Checked_by_ba_id       *string   `json:"checked_by_ba_id" default:""`
}