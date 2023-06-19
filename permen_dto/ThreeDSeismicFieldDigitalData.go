package dto

type Three_d_seismic_field_digital_data struct{

Ba_long_name       *string   `json:"ba_long_name" default:""`
Ba_type       *string   `json:"ba_type" default:""`
Area_id       *string   `json:"area_id" default:""`
Area_type       *string   `json:"area_type" default:""`
Acqtn_survey_name       *string   `json:"acqtn_survey_name" default:""`
Start_date       *string   `json:"start_date" default:""`
Shot_by       *string   `json:"shot_by" default:""`
Rcrd_rec_length       *string   `json:"rcrd_rec_length" default:""`
Rcrd_rec_length_ouom       *string   `json:"rcrd_rec_length_ouom" default:""`
Rcrd_sample_rate       *string   `json:"rcrd_sample_rate" default:""`
Rcrd_sample_rate_ouom       *string   `json:"rcrd_sample_rate_ouom" default:""`
First_seis_point_id       *string   `json:"first_seis_point_id" default:""`
Last_seis_point_id       *string   `json:"last_seis_point_id" default:""`
Create_date       *string   `json:"create_date" default:""`
Proc_set_type       *string   `json:"proc_set_type" default:""`
Field_file_number       *string   `json:"field_file_number" default:""`
Tape_number       *string   `json:"tape_number" default:""`
Rcrd_format_type       *string   `json:"rcrd_format_type" default:""`
Data_store_name       *string   `json:"data_store_name" default:""`
Original_file_name       *string   `json:"original_file_name" default:""`
Password       *string   `json:"password" default:""`
Digital_size       *string   `json:"digital_size" default:""`
Digital_size_uom       *string   `json:"digital_size_uom" default:""`
Remark       *string   `json:"remark" default:""`
Source       *string   `json:"source" default:""`
Qc_status       *string   `json:"qc_status" default:""`
Checked_by_ba_id       *string   `json:"checked_by_ba_id" default:""`
}