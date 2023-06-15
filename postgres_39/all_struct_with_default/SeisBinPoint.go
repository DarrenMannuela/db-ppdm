package dto

type Seis_bin_point struct {
	Seis_set_subtype      string   `json:"seis_set_subtype" default:"source"`
	Seis_set_id           string   `json:"seis_set_id" default:"source"`
	Bin_grid_id           string   `json:"bin_grid_id" default:"source"`
	Bin_grid_source       string   `json:"bin_grid_source" default:"source"`
	Bin_point_id          string   `json:"bin_point_id" default:"source"`
	Active_ind            *string  `json:"active_ind" default:""`
	Easting               *float64 `json:"easting" default:""`
	Effective_date        *string  `json:"effective_date" default:""`
	Expiry_date           *string  `json:"expiry_date" default:""`
	Nline_no              *int     `json:"nline_no" default:""`
	Northing              *float64 `json:"northing" default:""`
	Ppdm_guid             *string  `json:"ppdm_guid" default:""`
	Remark                *string  `json:"remark" default:""`
	Velocity_analysis_ind *string  `json:"velocity_analysis_ind" default:""`
	Xline_no              *int     `json:"xline_no" default:""`
	Row_changed_by        *string  `json:"row_changed_by" default:""`
	Row_changed_date      *string  `json:"row_changed_date" default:""`
	Row_created_by        *string  `json:"row_created_by" default:""`
	Row_created_date      *string  `json:"row_created_date" default:""`
	Row_effective_date    *string  `json:"row_effective_date" default:""`
	Row_expiry_date       *string  `json:"row_expiry_date" default:""`
	Row_quality           *string  `json:"row_quality" default:""`
}
