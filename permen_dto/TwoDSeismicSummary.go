package dto

type Two_d_seismic_summary struct{

Ba_long_name       *string   `json:"ba_long_name" default:""`
Ba_type       *string   `json:"ba_type" default:""`
Area_id       *string   `json:"area_id" default:""`
Area_type       *string   `json:"area_type" default:""`
Acqtn_survey_name       *string   `json:"acqtn_survey_name" default:""`
Seis_dimension       *string   `json:"seis_dimension" default:""`
Start_date       *string   `json:"start_date" default:""`
Shot_by       *string   `json:"shot_by" default:""`
Crew_long_name       *string   `json:"crew_long_name" default:""`
Acqtn_shotpt_interval       *string   `json:"acqtn_shotpt_interval" default:""`
Acqtn_shotpt_interval_ouom       *string   `json:"acqtn_shotpt_interval_ouom" default:""`
Rcvr_spacing       *string   `json:"rcvr_spacing" default:""`
Rcvr_spacing_ouom       *string   `json:"rcvr_spacing_ouom" default:""`
Energy_type       *string   `json:"energy_type" default:""`
Fold_coverage       *string   `json:"fold_coverage" default:""`
Rcrd_channel_count       *string   `json:"rcrd_channel_count" default:""`
Rcrd_rec_length       *string   `json:"rcrd_rec_length" default:""`
Rcrd_rec_length_ouom       *string   `json:"rcrd_rec_length_ouom" default:""`
Rcrd_sample_rate       *string   `json:"rcrd_sample_rate" default:""`
Rcrd_sample_rate_ouom       *string   `json:"rcrd_sample_rate_ouom" default:""`
Line_name       *string   `json:"line_name" default:""`
First_seis_point_id       *string   `json:"first_seis_point_id" default:""`
Last_seis_point_id       *string   `json:"last_seis_point_id" default:""`
Acqtn_direction       *string   `json:"acqtn_direction" default:""`
Line_length       *string   `json:"line_length" default:""`
Line_length_ouom       *string   `json:"line_length_ouom" default:""`
Alias_long_name       *string   `json:"alias_long_name" default:""`
Environment       *string   `json:"environment" default:""`
Remark       *string   `json:"remark" default:""`
Source       *string   `json:"source" default:""`
Qc_status       *string   `json:"qc_status" default:""`
Checked_by_ba_id       *string   `json:"checked_by_ba_id" default:""`
}