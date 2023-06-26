package dto

type Print_well_report struct{

Id      int  `json:"id" default:""`
Ba_long_name   *string  `json:"ba_long_name" default:""`
Ba_type   *string  `json:"ba_type" default:""`
Area_id   *string  `json:"area_id" default:""`
Area_type   *string  `json:"area_type" default:""`
Field_name   *string  `json:"field_name" default:""`
Well_name   *string  `json:"well_name" default:""`
Uwi_   *string  `json:"uwi_" default:""`
Title_   *string  `json:"title_" default:""`
Creator_name   *string  `json:"creator_name" default:""`
Create_date   *string  `json:"create_date" default:""`
Media_type   *string  `json:"media_type" default:""`
Document_type   *string  `json:"document_type" default:""`
Item_category   *string  `json:"item_category" default:""`
Item_sub_category   *string  `json:"item_sub_category" default:""`
Page_count   *string  `json:"page_count" default:""`
Remark_   *string  `json:"remark_" default:""`
Data_store_name   *string  `json:"data_store_name" default:""`
Data_store_type   *string  `json:"data_store_type" default:""`
Source_   *string  `json:"source_" default:""`
Qc_status   *string  `json:"qc_status" default:""`
Checked_by_ba_id   *string  `json:"checked_by_ba_id" default:""`
}