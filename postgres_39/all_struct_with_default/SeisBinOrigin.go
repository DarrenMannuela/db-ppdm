package dto

type Seis_bin_origin struct {
	Seis_set_subtype       string   `json:"seis_set_subtype" default:"source"`
	Seis_set_id            string   `json:"seis_set_id" default:"source"`
	Bin_grid_id            string   `json:"bin_grid_id" default:"source"`
	Bin_grid_source        string   `json:"bin_grid_source" default:"source"`
	Bin_origin_id          string   `json:"bin_origin_id" default:"source"`
	Active_ind             *string  `json:"active_ind" default:""`
	Coord_acquisition_id   *string  `json:"coord_acquisition_id" default:""`
	Coord_system_id        *string  `json:"coord_system_id" default:""`
	Corner1_lat            *float64 `json:"corner_1_lat" default:""`
	Corner1_long           *float64 `json:"corner_1_long" default:""`
	Corner2_lat            *float64 `json:"corner_2_lat" default:""`
	Corner2_long           *float64 `json:"corner_2_long" default:""`
	Corner3_lat            *float64 `json:"corner_3_lat" default:""`
	Corner3_long           *float64 `json:"corner_3_long" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Exclusion_ind          *string  `json:"exclusion_ind" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Inclusion_ind          *string  `json:"inclusion_ind" default:""`
	Input_bin_grid_id      *string  `json:"input_bin_grid_id" default:""`
	Input_bin_source       *string  `json:"input_bin_source" default:""`
	Input_seis_set_id      *string  `json:"input_seis_set_id" default:""`
	Input_seis_set_subtype *string  `json:"input_seis_set_subtype" default:""`
	Local_coord_system_id  *string  `json:"local_coord_system_id" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Remark                 *string  `json:"remark" default:""`
	Source                 *string  `json:"source" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
