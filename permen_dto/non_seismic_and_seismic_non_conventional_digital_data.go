package dto

type Non_seismic_and_seismic_non_conventional_digital_data struct {
	Ba_long_name       *string `json:"ba_long_name" default:""`
	Ba_type            *string `json:"ba_type" default:""`
	Area_id            *string `json:"area_id" default:""`
	Area_type          *string `json:"area_type" default:""`
	Acqtn_survey_name  *string `json:"acqtn_survey_name" default:""`
	Processing_company *string `json:"processing_company" default:""`
	Start_date         *string `json:"start_date" default:""`
	Seis_station_type  *string `json:"seis_station_type" default:""`
	Create_date        *string `json:"create_date" default:""`
	Proc_set_type      *string `json:"proc_set_type" default:""`
	Media_type         *string `json:"media_type" default:""`
	Tape_number        *string `json:"tape_number" default:""`
	Digital_format     *string `json:"digital_format" default:""`
	Data_store_name    *string `json:"data_store_name" default:""`
	Original_file_name *string `json:"original_file_name" default:""`
	Password           *string `json:"password" default:""`
	Digital_size       *string `json:"digital_size" default:""`
	Digital_size_uom   *string `json:"digital_size_uom" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Qc_status          *string `json:"qc_status" default:""`
	Checked_by_ba_id   *string `json:"checked_by_ba_id" default:""`
}
