package dto

type Digital_project_file struct{

Id      int  `json:"id" default:"" gorm:"primaryKey"`
Ba_long_name   *string  `json:"ba_long_name" default:""`
Ba_type   *string  `json:"ba_type" default:""`
Area_id   *string  `json:"area_id" default:""`
Area_type   *string  `json:"area_type" default:""`
Field_name   *string  `json:"field_name" default:""`
Project_name   *string  `json:"project_name" default:""`
Sw_application_id   *string  `json:"sw_application_id" default:""`
Application_version   *string  `json:"application_version" default:""`
Item_category   *string  `json:"item_category" default:""`
Process_date   *string  `json:"process_date" default:""`
Interpreter_   *string  `json:"interpreter_" default:""`
Digital_format   *string  `json:"digital_format" default:""`
Media_type   *string  `json:"media_type" default:""`
Original_file_name   *string  `json:"original_file_name" default:""`
Password_   *string  `json:"password_" default:""`
Digital_size   *int  `json:"digital_size" default:""`
Digital_size_uom   *string  `json:"digital_size_uom" default:""`
Ba_long_name_2   *string  `json:"ba_long_name_2" default:""`
Ba_type_2   *string  `json:"ba_type_2" default:""`
Data_store_name   *string  `json:"data_store_name" default:""`
Remark_   *string  `json:"remark_" default:""`
Source_   *string  `json:"source_" default:""`
Qc_status   *string  `json:"qc_status" default:""`
Checked_by_ba_id   *string  `json:"checked_by_ba_id" default:""`
}