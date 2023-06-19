package dto

type Outcrop_sample struct{

Ba_long_name       *string   `json:"ba_long_name" default:""`
Ba_type       *string   `json:"ba_type" default:""`
Area_id       *string   `json:"area_id" default:""`
Area_type       *string   `json:"area_type" default:""`
Project_name       *string   `json:"project_name" default:""`
Field_station_id       *string   `json:"field_station_id" default:""`
Longitude       *string   `json:"longitude" default:""`
Latitude       *string   `json:"latitude" default:""`
Easting       *string   `json:"easting" default:""`
Easting_ouom       *string   `json:"easting_ouom" default:""`
Northing       *string   `json:"northing" default:""`
Northing_ouom       *string   `json:"northing_ouom" default:""`
Utm_quadrant       *string   `json:"utm_quadrant" default:""`
Geodetic_datum_name       *string   `json:"geodetic_datum_name" default:""`
Sample_num       *string   `json:"sample_num" default:""`
Sample_count       *string   `json:"sample_count" default:""`
Study_type       *string   `json:"study_type" default:""`
Collected_date       *string   `json:"collected_date" default:""`
Pick_location       *string   `json:"pick_location" default:""`
Data_store_name       *string   `json:"data_store_name" default:""`
Data_store_type       *string   `json:"data_store_type" default:""`
Location_id       *string   `json:"location_id" default:""`
Remark       *string   `json:"remark" default:""`
Source       *string   `json:"source" default:""`
Qc_status       *string   `json:"qc_status" default:""`
Checked_by_ba_id       *string   `json:"checked_by_ba_id" default:""`
}