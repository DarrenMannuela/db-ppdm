package dto

type Legal_fps_loc struct {
	Legal_fps_loc_id    string   `json:"legal_fps_loc_id" default:"source"`
	Location_type       string   `json:"location_type" default:"source"`
	Source              string   `json:"source" default:"source"`
	Active_ind          *string  `json:"active_ind" default:""`
	Coord_system_id     *string  `json:"coord_system_id" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Fps_event_sequence  *string  `json:"fps_event_sequence" default:""`
	Fps_loc_exception   *string  `json:"fps_loc_exception" default:""`
	Grid                *string  `json:"grid" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Reference_latitude  *float64 `json:"reference_latitude" default:""`
	Reference_longitude *float64 `json:"reference_longitude" default:""`
	Remark              *string  `json:"remark" default:""`
	Section             *int     `json:"section" default:""`
	Unit                *string  `json:"unit" default:""`
	Uwi                 *string  `json:"uwi" default:""`
	Well_node_id        *string  `json:"well_node_id" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}
