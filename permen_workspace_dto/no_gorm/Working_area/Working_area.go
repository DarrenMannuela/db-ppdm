package dto

type Working_area struct{

Id      int  `json:"id" default:""`
Area_id   *string  `json:"area_id" default:""`
Area_type   *string  `json:"area_type" default:""`
Ba_long_name   *string  `json:"ba_long_name" default:""`
Ba_type   *string  `json:"ba_type" default:""`
Effective_date   *string  `json:"effective_date" default:""`
Term_exiry_date   *string  `json:"term_exiry_date" default:""`
Contract_type   *string  `json:"contract_type" default:""`
R_granted_right_type   *string  `json:"r_granted_right_type" default:""`
Gross_size   *int  `json:"gross_size" default:""`
Gross_size_ouom   *string  `json:"gross_size_ouom" default:""`
Land_right_category   *string  `json:"land_right_category" default:""`
Producing_ind   *string  `json:"producing_ind" default:""`
Original_file_name   *string  `json:"original_file_name" default:""`
Password   *string  `json:"password" default:""`
Digital_size   *int  `json:"digital_size" default:""`
Digital_size_uom   *string  `json:"digital_size_uom" default:""`
Media_type   *string  `json:"media_type" default:""`
Data_store_name   *string  `json:"data_store_name" default:""`
Source   *string  `json:"source" default:""`
Qc_status   *string  `json:"qc_status" default:""`
Checked_by_ba_id   *string  `json:"checked_by_ba_id" default:""`
}