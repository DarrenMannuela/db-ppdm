package dto

type Digital_technical_report struct{

Ba_long_name       *string   `json:"ba_long_name" default:""`
Ba_type       *string   `json:"ba_type" default:""`
Area_id       *string   `json:"area_id" default:""`
Area_type       *string   `json:"area_type" default:""`
Title       *string   `json:"title" default:""`
Creator_name       *string   `json:"creator_name" default:""`
Create_date       *string   `json:"create_date" default:""`
Page_count       *string   `json:"page_count" default:""`
Media_type       *string   `json:"media_type" default:""`
Digital_format       *string   `json:"digital_format" default:""`
Document_type       *string   `json:"document_type" default:""`
Item_category       *string   `json:"item_category" default:""`
Item_sub_category       *string   `json:"item_sub_category" default:""`
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