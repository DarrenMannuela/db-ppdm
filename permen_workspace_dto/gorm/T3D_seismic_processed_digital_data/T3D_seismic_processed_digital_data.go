package dto

type T3D_seismic_processed_digital_data struct{

Id      int  `json:"id" default:"" gorm:"primaryKey"`
Ba_long_name   *string  `json:"ba_long_name" default:""`
Ba_type   *string  `json:"ba_type" default:""`
Area_id   *string  `json:"area_id" default:""`
Area_type   *string  `json:"area_type" default:""`
Acqtn_survey_name   *string  `json:"acqtn_survey_name" default:""`
Processing_company   *string  `json:"processing_company" default:""`
Start_date   *string  `json:"start_date" default:""`
Rcrd_rec_length   *int  `json:"rcrd_rec_length" default:""`
Rcrd_rec_length_ouom   *string  `json:"rcrd_rec_length_ouom" default:""`
Rcrd_sample_rate   *int  `json:"rcrd_sample_rate" default:""`
Rcrd_sample_rate_ouom   *string  `json:"rcrd_sample_rate_ouom" default:""`
First_nline_no   *int  `json:"first_nline_no" default:""`
Last_nline_no   *int  `json:"last_nline_no" default:""`
First_xline_no   *int  `json:"first_xline_no" default:""`
Last_xline_no   *int  `json:"last_xline_no" default:""`
Digital_format   *string  `json:"digital_format" default:""`
Media_type   *string  `json:"media_type" default:""`
Proc_set_type   *string  `json:"proc_set_type" default:""`
Polarity   *string  `json:"polarity" default:""`
Ba_long_name_2   *string  `json:"ba_long_name_2" default:""`
Ba_type_2   *string  `json:"ba_type_2" default:""`
Data_store_name   *string  `json:"data_store_name" default:""`
Original_file_name   *string  `json:"original_file_name" default:""`
Password   *string  `json:"password" default:""`
Sw_application_id   *string  `json:"sw_application_id" default:""`
Application_version   *string  `json:"application_version" default:""`
Digital_size   *int  `json:"digital_size" default:""`
Digital_size_uom   *string  `json:"digital_size_uom" default:""`
Remark   *string  `json:"remark" default:""`
Source   *string  `json:"source" default:""`
Qc_status   *string  `json:"qc_status" default:""`
Checked_by_ba_id   *string  `json:"checked_by_ba_id" default:""`
}