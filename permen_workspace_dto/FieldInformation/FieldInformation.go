package dto

type Field_information struct{

Area_id       *string   `json:"area_id" default:""`
Area_type       *string   `json:"area_type" default:""`
Field_name       *string   `json:"field_name" default:""`
Discovery_date       *string   `json:"discovery_date" default:""`
Field_type       *string   `json:"field_type" default:""`
Original_file_name       *string   `json:"original_file_name" default:""`
Password       *string   `json:"password" default:""`
Digital_size       *string   `json:"digital_size" default:""`
Digital_size_uom       *string   `json:"digital_size_uom" default:""`
Media_type       *string   `json:"media_type" default:""`
Data_store_name       *string   `json:"data_store_name" default:""`
Remark       *string   `json:"remark" default:""`
Qc_status       *string   `json:"qc_status" default:""`
Checked_by_ba_id       *string   `json:"checked_by_ba_id" default:""`
}