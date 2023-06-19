package dto

type Non_seismic_and_seismic_non_conventional_data_summary struct{

Ba_long_name       *string   `json:"ba_long_name" default:""`
Ba_type       *string   `json:"ba_type" default:""`
Area_id       *string   `json:"area_id" default:""`
Area_type       *string   `json:"area_type" default:""`
Acqtn_survey_name       *string   `json:"acqtn_survey_name" default:""`
Seis_spectrum_type       *string   `json:"seis_spectrum_type" default:""`
Seis_dimension       *string   `json:"seis_dimension" default:""`
Start_date       *string   `json:"start_date" default:""`
Shot_by       *string   `json:"shot_by" default:""`
Crew_long_name       *string   `json:"crew_long_name" default:""`
Acqtn_direction       *string   `json:"acqtn_direction" default:""`
Environment       *string   `json:"environment" default:""`
Remark       *string   `json:"remark" default:""`
Source       *string   `json:"source" default:""`
Qc_status       *string   `json:"qc_status" default:""`
Checked_by_ba_id       *string   `json:"checked_by_ba_id" default:""`
}