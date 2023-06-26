package dto

type T2D_3d_seismic_report struct{

Id      int  `json:"id" default:""`
Ba_long_name   *string  `json:"ba_long_name" default:""`
Ba_type   *string  `json:"ba_type" default:""`
Area_id   *string  `json:"area_id" default:""`
Area_type   *string  `json:"area_type" default:""`
Acqtn_survey_name   *string  `json:"acqtn_survey_name" default:""`
Start_date   *string  `json:"start_date" default:""`
Seis_dimension   *string  `json:"seis_dimension" default:""`
Line_name   *string  `json:"line_name" default:""`
Title_   *string  `json:"title_" default:""`
Creator_name   *string  `json:"creator_name" default:""`
Create_date   *string  `json:"create_date" default:""`
Item_category   *string  `json:"item_category" default:""`
Media_type   *string  `json:"media_type" default:""`
Ba_long_name_2   *string  `json:"ba_long_name_2" default:""`
Ba_type_2   *string  `json:"ba_type_2" default:""`
Data_store_name   *string  `json:"data_store_name" default:""`
Data_store_type   *string  `json:"data_store_type" default:""`
Location_id   *string  `json:"location_id" default:""`
Remark_   *string  `json:"remark_" default:""`
Source_   *string  `json:"source_" default:""`
Qc_status   *string  `json:"qc_status" default:""`
Checked_by_ba_id   *string  `json:"checked_by_ba_id" default:""`
}