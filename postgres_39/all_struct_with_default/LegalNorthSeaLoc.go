package dto

type Legal_north_sea_loc struct {
	Legal_north_sea_loc_id string   `json:"legal_north_sea_loc_id" default:"source"`
	Location_type          string   `json:"location_type" default:"source"`
	Source                 string   `json:"source" default:"source"`
	Active_ind             *string  `json:"active_ind" default:""`
	Block_no               *int     `json:"block_no" default:""`
	Block_suffix           *string  `json:"block_suffix" default:""`
	Coord_system_id        *string  `json:"coord_system_id" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Land_well_ind          *string  `json:"land_well_ind" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Principal_meridian     *string  `json:"principal_meridian" default:""`
	Quadrant               *float64 `json:"quadrant" default:""`
	Quadrant_prefix        *string  `json:"quadrant_prefix" default:""`
	Remark                 *string  `json:"remark" default:""`
	Uwi                    *string  `json:"uwi" default:""`
	Well_node_id           *string  `json:"well_node_id" default:""`
	Well_no_prefix         *string  `json:"well_no_prefix" default:""`
	Well_no_suffix         *string  `json:"well_no_suffix" default:""`
	Well_seq_no            *int     `json:"well_seq_no" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
