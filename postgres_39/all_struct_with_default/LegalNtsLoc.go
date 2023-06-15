package dto

type Legal_nts_loc struct {
	Legal_nts_loc_id   string   `json:"legal_nts_loc_id" default:"source"`
	Location_type      string   `json:"location_type" default:"source"`
	Source             string   `json:"source" default:"source"`
	Active_ind         *string  `json:"active_ind" default:""`
	Block              *string  `json:"block" default:""`
	Coord_system_id    *string  `json:"coord_system_id" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Event_sequence     *string  `json:"event_sequence" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Letter_quadrangle  *string  `json:"letter_quadrangle" default:""`
	Nts_loc_exception  *string  `json:"nts_loc_exception" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Primary_quadrangle *float64 `json:"primary_quadrangle" default:""`
	Quarter_unit       *string  `json:"quarter_unit" default:""`
	Remark             *string  `json:"remark" default:""`
	Sixteenth          *float64 `json:"sixteenth" default:""`
	Unit               *int     `json:"unit" default:""`
	Uwi                *string  `json:"uwi" default:""`
	Well_node_id       *string  `json:"well_node_id" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
