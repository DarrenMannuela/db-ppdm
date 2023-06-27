package dto

type T2D_seismic_field_data struct{

Id      int  `json:"id" default:""`
Ba_long_name   *string  `json:"ba_long_name" default:""`
Ba_type   *string  `json:"ba_type" default:""`
Area_id   *string  `json:"area_id" default:""`
Area_type   *string  `json:"area_type" default:""`
Acqtn_survey_name   *string  `json:"acqtn_survey_name" default:""`
Shot_by   *string  `json:"shot_by" default:""`
Start_date   *string  `json:"start_date" default:""`
Rcrd_rec_length   *int  `json:"rcrd_rec_length" default:""`
Rcrd_rec_length_ouom   *string  `json:"rcrd_rec_length_ouom" default:""`
Rcrd_sample_rate   *int  `json:"rcrd_sample_rate" default:""`
Rcrd_sample_rate_ouom   *string  `json:"rcrd_sample_rate_ouom" default:""`
Line_name   *string  `json:"line_name" default:""`
First_seis_point_id   *string  `json:"first_seis_point_id" default:""`
Last_seis_point_id   *string  `json:"last_seis_point_id" default:""`
Create_date   *string  `json:"create_date" default:""`
Proc_set_type   *string  `json:"proc_set_type" default:""`
Field_file_number   *string  `json:"field_file_number" default:""`
Media_type   *string  `json:"media_type" default:""`
Tape_number   *string  `json:"tape_number" default:""`
Rcrd_format_type   *string  `json:"rcrd_format_type" default:""`
Ba_long_name_2   *string  `json:"ba_long_name_2" default:""`
Ba_type_2   *string  `json:"ba_type_2" default:""`
Data_store_name   *string  `json:"data_store_name" default:""`
Data_store_type   *string  `json:"data_store_type" default:""`
Location_id   *string  `json:"location_id" default:""`
Remark   *string  `json:"remark" default:""`
Source   *string  `json:"source" default:""`
Qc_status   *string  `json:"qc_status" default:""`
Checked_by_ba_id   *string  `json:"checked_by_ba_id" default:""`
}