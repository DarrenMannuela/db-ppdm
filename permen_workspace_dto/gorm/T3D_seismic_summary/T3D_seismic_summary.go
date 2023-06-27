package dto

type T3D_seismic_summary struct{

Id      int  `json:"id" default:"" gorm:"primaryKey"`
Ba_long_name   *string  `json:"ba_long_name" default:""`
Ba_type   *string  `json:"ba_type" default:""`
Area_id   *string  `json:"area_id" default:""`
Area_type   *string  `json:"area_type" default:""`
Acqtn_survey_name   *string  `json:"acqtn_survey_name" default:""`
Area_size   *int  `json:"area_size" default:""`
Area_size_ouom   *string  `json:"area_size_ouom" default:""`
Acqtn_inline_bin_size   *int  `json:"acqtn_inline_bin_size" default:""`
Acqtn_inline_bin_size_ouom   *string  `json:"acqtn_inline_bin_size_ouom" default:""`
Acqtn_xline_bin_size   *int  `json:"acqtn_xline_bin_size" default:""`
Acqtn_xline_bin_size_ouom   *string  `json:"acqtn_xline_bin_size_ouom" default:""`
Seis_dimension   *string  `json:"seis_dimension" default:""`
Start_date   *string  `json:"start_date" default:""`
Shot_by   *string  `json:"shot_by" default:""`
Crew_long_name   *string  `json:"crew_long_name" default:""`
Acqtn_shot_line_spacing   *int  `json:"acqtn_shot_line_spacing" default:""`
Acqtn_shot_line_spacing_ouom   *string  `json:"acqtn_shot_line_spacing_ouom" default:""`
Rcvr_spacing   *int  `json:"rcvr_spacing" default:""`
Rcvr_spacing_ouom   *string  `json:"rcvr_spacing_ouom" default:""`
Rcvr_line_spacing   *int  `json:"rcvr_line_spacing" default:""`
Rcvr_line_spacing_ouom   *string  `json:"rcvr_line_spacing_ouom" default:""`
Seis_3d_type   *string  `json:"seis_3d_type" default:""`
Energy_type   *string  `json:"energy_type" default:""`
Fold_coverage   *int  `json:"fold_coverage" default:""`
Rcrd_channel_count   *int  `json:"rcrd_channel_count" default:""`
Rcrd_rec_length   *int  `json:"rcrd_rec_length" default:""`
Rcrd_rec_length_ouom   *string  `json:"rcrd_rec_length_ouom" default:""`
Rcrd_sample_rate   *int  `json:"rcrd_sample_rate" default:""`
Rcrd_sample_rate_ouom   *string  `json:"rcrd_sample_rate_ouom" default:""`
Alias_long_name   *string  `json:"alias_long_name" default:""`
Environment_   *string  `json:"environment_" default:""`
Remark_   *string  `json:"remark_" default:""`
Source_   *string  `json:"source_" default:""`
Qc_status   *string  `json:"qc_status" default:""`
Checked_by_ba_id   *string  `json:"checked_by_ba_id" default:""`
}