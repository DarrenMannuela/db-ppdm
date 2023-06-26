package dto

type T2D_seismic_navigation_digital_data struct{

Id      int  `json:"id" default:"" gorm:"primaryKey"`
Ba_long_name   *string  `json:"ba_long_name" default:""`
Ba_type   *string  `json:"ba_type" default:""`
Area_id   *string  `json:"area_id" default:""`
Area_type   *string  `json:"area_type" default:""`
Acqtn_survey_name   *string  `json:"acqtn_survey_name" default:""`
Seis_dimension   *string  `json:"seis_dimension" default:""`
Process_date   *string  `json:"process_date" default:""`
Shot_by   *string  `json:"shot_by" default:""`
Line_name   *string  `json:"line_name" default:""`
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