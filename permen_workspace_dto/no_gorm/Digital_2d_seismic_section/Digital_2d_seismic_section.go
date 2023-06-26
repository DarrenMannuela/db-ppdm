package dto

type Digital_2d_seismic_section struct{

Id      int  `json:"id" default:""`
Ba_long_name   *string  `json:"ba_long_name" default:""`
Ba_type   *string  `json:"ba_type" default:""`
Area_id   *string  `json:"area_id" default:""`
Area_type   *string  `json:"area_type" default:""`
Acqtn_survey_name   *string  `json:"acqtn_survey_name" default:""`
Processing_company   *string  `json:"processing_company" default:""`
Start_date   *string  `json:"start_date" default:""`
Line_name   *string  `json:"line_name" default:""`
First_seis_point_id   *string  `json:"first_seis_point_id" default:""`
Last_seis_point_id   *string  `json:"last_seis_point_id" default:""`
Create_date   *string  `json:"create_date" default:""`
Proc_set_type   *string  `json:"proc_set_type" default:""`
Digital_format   *string  `json:"digital_format" default:""`
Media_type   *string  `json:"media_type" default:""`
Vertical_scale   *string  `json:"vertical_scale" default:""`
Vertical_scale_ouom   *string  `json:"vertical_scale_ouom" default:""`
Horizontal_scale   *string  `json:"horizontal_scale" default:""`
Horizontal_scale_ouom   *string  `json:"horizontal_scale_ouom" default:""`
Polarity_   *string  `json:"polarity_" default:""`
Ba_long_name_2   *string  `json:"ba_long_name_2" default:""`
Ba_type_2   *string  `json:"ba_type_2" default:""`
Data_store_name   *string  `json:"data_store_name" default:""`
Original_file_name   *string  `json:"original_file_name" default:""`
Password_   *string  `json:"password_" default:""`
Sw_application_id   *string  `json:"sw_application_id" default:""`
Application_version   *string  `json:"application_version" default:""`
Digital_size   *int  `json:"digital_size" default:""`
Digital_size_uom   *string  `json:"digital_size_uom" default:""`
Remark_   *string  `json:"remark_" default:""`
Source_   *string  `json:"source_" default:""`
Qc_status   *string  `json:"qc_status" default:""`
Checked_by_ba_id   *string  `json:"checked_by_ba_id" default:""`
}