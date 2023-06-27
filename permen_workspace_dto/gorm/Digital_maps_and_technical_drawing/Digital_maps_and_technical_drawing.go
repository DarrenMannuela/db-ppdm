package dto

type Digital_maps_and_technical_drawing struct{

Id      int  `json:"id" default:"" gorm:"primaryKey"`
Ba_long_name   *string  `json:"ba_long_name" default:""`
Ba_type   *string  `json:"ba_type" default:""`
Area_id   *string  `json:"area_id" default:""`
Area_type   *string  `json:"area_type" default:""`
Title   *string  `json:"title" default:""`
Creator_name   *string  `json:"creator_name" default:""`
Create_date   *string  `json:"create_date" default:""`
Digital_format   *string  `json:"digital_format" default:""`
Media_type   *string  `json:"media_type" default:""`
Map_scale   *string  `json:"map_scale" default:""`
Projection_type   *string  `json:"projection_type" default:""`
Geodetic_datum_name   *string  `json:"geodetic_datum_name" default:""`
Document_type   *string  `json:"document_type" default:""`
Item_category   *string  `json:"item_category" default:""`
Ba_long_name_2   *string  `json:"ba_long_name_2" default:""`
Ba_type_2   *string  `json:"ba_type_2" default:""`
Data_store_name   *string  `json:"data_store_name" default:""`
Original_file_name   *string  `json:"original_file_name" default:""`
Password   *string  `json:"password" default:""`
Digital_size   *int  `json:"digital_size" default:""`
Digital_size_uom   *string  `json:"digital_size_uom" default:""`
Remark   *string  `json:"remark" default:""`
Source   *string  `json:"source" default:""`
Qc_status   *string  `json:"qc_status" default:""`
Checked_by_ba_id   *string  `json:"checked_by_ba_id" default:""`
}