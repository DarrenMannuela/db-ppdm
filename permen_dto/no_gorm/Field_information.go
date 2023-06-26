package dto

type Field_information struct{

Id      int  `json:"id" default:""`
Area_id   *string  `json:"area_id" default:""`
Area_type   *string  `json:"area_type" default:""`
Field_name   *string  `json:"field_name" default:""`
Discovery_date   *string  `json:"discovery_date" default:""`
Field_type   *string  `json:"field_type" default:""`
Original_file_name   *string  `json:"original_file_name" default:""`
Password_   *string  `json:"password_" default:""`
Digital_size   *int  `json:"digital_size" default:""`
Digital_size_uom   *string  `json:"digital_size_uom" default:""`
Media_type   *string  `json:"media_type" default:""`
Data_store_name   *string  `json:"data_store_name" default:""`
Remark_   *string  `json:"remark_" default:""`
Qc_status   *string  `json:"qc_status" default:""`
Checked_by_ba_id   *string  `json:"checked_by_ba_id" default:""`
}