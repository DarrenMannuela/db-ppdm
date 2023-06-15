package dto

type Legal_dls_loc struct {
	Legal_dls_loc_id       string   `json:"legal_dls_loc_id" default:"source"`
	Location_type          string   `json:"location_type" default:"source"`
	Source                 string   `json:"source" default:"source"`
	Active_ind             *string  `json:"active_ind" default:""`
	Coord_system_id        *string  `json:"coord_system_id" default:""`
	Dls_event_sequence     *string  `json:"dls_event_sequence" default:""`
	Dls_legal_subdivision  *float64 `json:"dls_legal_subdivision" default:""`
	Dls_loc_exception      *string  `json:"dls_loc_exception" default:""`
	Dls_meridian           *float64 `json:"dls_meridian" default:""`
	Dls_meridian_direction *string  `json:"dls_meridian_direction" default:""`
	Dls_range              *float64 `json:"dls_range" default:""`
	Dls_range_modifier     *string  `json:"dls_range_modifier" default:""`
	Dls_section            *float64 `json:"dls_section" default:""`
	Dls_township           *float64 `json:"dls_township" default:""`
	Dls_township_modifier  *string  `json:"dls_township_modifier" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Remark                 *string  `json:"remark" default:""`
	Uwi                    *string  `json:"uwi" default:""`
	Well_node_id           *string  `json:"well_node_id" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
