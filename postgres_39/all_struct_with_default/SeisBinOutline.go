package dto

type Seis_bin_outline struct {
	Seis_set_subtype      string   `json:"seis_set_subtype" default:"source"`
	Seis_set_id           string   `json:"seis_set_id" default:"source"`
	Bin_grid_id           string   `json:"bin_grid_id" default:"source"`
	Bin_grid_source       string   `json:"bin_grid_source" default:"source"`
	Outline_seq_no        int      `json:"outline_seq_no" default:"1"`
	Active_ind            *string  `json:"active_ind" default:""`
	Bin_outline_type      *string  `json:"bin_outline_type" default:""`
	Bin_point_id          *string  `json:"bin_point_id" default:""`
	Coord_acquisition_id  *string  `json:"coord_acquisition_id" default:""`
	Coord_system_id       *string  `json:"coord_system_id" default:""`
	Easting               *float64 `json:"easting" default:""`
	Effective_date        *string  `json:"effective_date" default:""`
	Expiry_date           *string  `json:"expiry_date" default:""`
	Local_coord_system_id *string  `json:"local_coord_system_id" default:""`
	Nline_no              *int     `json:"nline_no" default:""`
	Northing              *float64 `json:"northing" default:""`
	Ppdm_guid             *string  `json:"ppdm_guid" default:""`
	Remark                *string  `json:"remark" default:""`
	Xline_no              *int     `json:"xline_no" default:""`
	Row_changed_by        *string  `json:"row_changed_by" default:""`
	Row_changed_date      *string  `json:"row_changed_date" default:""`
	Row_created_by        *string  `json:"row_created_by" default:""`
	Row_created_date      *string  `json:"row_created_date" default:""`
	Row_effective_date    *string  `json:"row_effective_date" default:""`
	Row_expiry_date       *string  `json:"row_expiry_date" default:""`
	Row_quality           *string  `json:"row_quality" default:""`
}
