package dto

type Well_node_strat_unit struct {
	Uwi                  string   `json:"uwi" default:"source"`
	Strat_name_set_id    string   `json:"strat_name_set_id" default:"source"`
	Strat_unit_id        string   `json:"strat_unit_id" default:"source"`
	Interp_id            string   `json:"interp_id" default:"source"`
	Active_ind           *string  `json:"active_ind" default:""`
	Base_offset_tvd      *float64 `json:"base_offset_tvd" default:""`
	Base_offset_tvd_ouom *string  `json:"base_offset_tvd_ouom" default:""`
	Effective_date       *string  `json:"effective_date" default:""`
	Expiry_date          *string  `json:"expiry_date" default:""`
	Interpreter_ba_id    *string  `json:"interpreter_ba_id" default:""`
	Location_type        *string  `json:"location_type" default:""`
	Node_id              *string  `json:"node_id" default:""`
	Node_position        *string  `json:"node_position" default:""`
	Pick_date            *string  `json:"pick_date" default:""`
	Pick_depth           *float64 `json:"pick_depth" default:""`
	Pick_depth_ouom      *string  `json:"pick_depth_ouom" default:""`
	Pick_location        *string  `json:"pick_location" default:""`
	Pick_method          *string  `json:"pick_method" default:""`
	Pick_method_desc     *string  `json:"pick_method_desc" default:""`
	Pick_quality         *string  `json:"pick_quality" default:""`
	Ppdm_guid            *string  `json:"ppdm_guid" default:""`
	Preferred_ind        *string  `json:"preferred_ind" default:""`
	Remark               *string  `json:"remark" default:""`
	Source               *string  `json:"source" default:""`
	Top_offset_tvd       *float64 `json:"top_offset_tvd" default:""`
	Top_offset_tvd_ouom  *string  `json:"top_offset_tvd_ouom" default:""`
	Row_changed_by       *string  `json:"row_changed_by" default:""`
	Row_changed_date     *string  `json:"row_changed_date" default:""`
	Row_created_by       *string  `json:"row_created_by" default:""`
	Row_created_date     *string  `json:"row_created_date" default:""`
	Row_effective_date   *string  `json:"row_effective_date" default:""`
	Row_expiry_date      *string  `json:"row_expiry_date" default:""`
	Row_quality          *string  `json:"row_quality" default:""`
}
