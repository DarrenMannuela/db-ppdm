package dto

type Print_technical_report struct{

Ba_long_name       *string   `json:"ba_long_name" default:""`
Ba_type       *string   `json:"ba_type" default:""`
Area_id       *string   `json:"area_id" default:""`
Area_type       *string   `json:"area_type" default:""`
Title       *string   `json:"title" default:""`
Creator_name       *string   `json:"creator_name" default:""`
Create_date       *string   `json:"create_date" default:""`
Page_count       *string   `json:"page_count" default:""`
Media_type       *string   `json:"media_type" default:""`
Document_type       *string   `json:"document_type" default:""`
Item_category       *string   `json:"item_category" default:""`
Item_sub_category       *string   `json:"item_sub_category" default:""`
Data_store_name       *string   `json:"data_store_name" default:""`
Data_store_type       *string   `json:"data_store_type" default:""`
Location_id       *string   `json:"location_id" default:""`
Remark       *string   `json:"remark" default:""`
Source       *string   `json:"source" default:""`
Qc_status       *string   `json:"qc_status" default:""`
Checked_by_ba_id       *string   `json:"checked_by_ba_id" default:""`
}