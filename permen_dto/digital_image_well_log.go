package dto

type Digital_image_well_log struct {
	Ba_long_name       *string `json:"ba_long_name" default:""`
	Ba_type            *string `json:"ba_type" default:""`
	Area_id            *string `json:"area_id" default:""`
	Area_type          *string `json:"area_type" default:""`
	Field_name         *string `json:"field_name" default:""`
	Well_name          *string `json:"well_name" default:""`
	Uwi                *string `json:"uwi" default:""`
	Logging_company    *string `json:"logging_company" default:""`
	Scale_ratio        *string `json:"scale_ratio" default:""`
	Media_type         *string `json:"media_type" default:""`
	Digital_format     *string `json:"digital_format" default:""`
	Well_log_class_id  *string `json:"well_log_class_id" default:""`
	Log_title          *string `json:"log_title" default:""`
	Report_log_run     *string `json:"report_log_run" default:""`
	Trip_date          *string `json:"trip_date" default:""`
	Top_depth          *string `json:"top_depth" default:""`
	Top_depth_ouom     *string `json:"top_depth_ouom" default:""`
	Base_depth         *string `json:"base_depth" default:""`
	Base_depth_ouom    *string `json:"base_depth_ouom" default:""`
	Original_file_name *string `json:"original_file_name" default:""`
	Password           *string `json:"password" default:""`
	Digital_size       *string `json:"digital_size" default:""`
	Digital_size_uom   *string `json:"digital_size_uom" default:""`
	Data_store_name    *string `json:"data_store_name" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Qc_status          *string `json:"qc_status" default:""`
	Checked_by_ba_id   *string `json:"checked_by_ba_id" default:""`
}
