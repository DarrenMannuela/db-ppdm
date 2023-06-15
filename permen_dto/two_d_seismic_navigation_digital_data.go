package dto

type Two_d_seismic_navigation_digital_data struct {
	Ba_long_name       *string `json:"ba_long_name" default:""`
	Ba_type            *string `json:"ba_type" default:""`
	Area_id            *string `json:"area_id" default:""`
	Area_type          *string `json:"area_type" default:""`
	Acqtn_survey_name  *string `json:"acqtn_survey_name" default:""`
	Seis_dimension     *string `json:"seis_dimension" default:""`
	Process_date       *string `json:"process_date" default:""`
	Shot_by            *string `json:"shot_by" default:""`
	Line_name          *string `json:"line_name" default:""`
	Digital_format     *string `json:"digital_format" default:""`
	Media_type         *string `json:"media_type" default:""`
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
