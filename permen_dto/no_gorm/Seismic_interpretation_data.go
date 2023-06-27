package dto

type Seismic_interpretation_data struct{

Id      int  `json:"id" default:""`
Ba_long_name   *string  `json:"ba_long_name" default:""`
Ba_type   *string  `json:"ba_type" default:""`
Area_id   *string  `json:"area_id" default:""`
Area_type   *string  `json:"area_type" default:""`
Field_name   *string  `json:"field_name" default:""`
Project_name   *string  `json:"project_name" default:""`
Interpreter   *string  `json:"interpreter" default:""`
Interp_date   *string  `json:"interp_date" default:""`
Interp_objective   *string  `json:"interp_objective" default:""`
Interp_type   *string  `json:"interp_type" default:""`
Pick_method   *string  `json:"pick_method" default:""`
Sw_application_id   *string  `json:"sw_application_id" default:""`
Application_version   *string  `json:"application_version" default:""`
Trace_position   *string  `json:"trace_position" default:""`
Media_type   *string  `json:"media_type" default:""`
Tape_number   *string  `json:"tape_number" default:""`
Digital_format   *string  `json:"digital_format" default:""`
Remark   *string  `json:"remark" default:""`
Ba_long_name_2   *string  `json:"ba_long_name_2" default:""`
Ba_type_2   *string  `json:"ba_type_2" default:""`
Data_store_name   *string  `json:"data_store_name" default:""`
Original_file_name   *string  `json:"original_file_name" default:""`
Password   *string  `json:"password" default:""`
Digital_size   *int  `json:"digital_size" default:""`
Digital_size_uom   *string  `json:"digital_size_uom" default:""`
Source   *string  `json:"source" default:""`
Qc_status   *string  `json:"qc_status" default:""`
Checked_by_ba_id   *string  `json:"checked_by_ba_id" default:""`
}