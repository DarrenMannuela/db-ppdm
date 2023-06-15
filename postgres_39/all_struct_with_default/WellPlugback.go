package dto

type Well_plugback struct {
	Uwi                  string   `json:"uwi" default:"source"`
	Source               string   `json:"source" default:"source"`
	Plugback_obs_no      int      `json:"plugback_obs_no" default:"1"`
	Active_ind           *string  `json:"active_ind" default:""`
	Base_depth           *float64 `json:"base_depth" default:""`
	Base_depth_ouom      *string  `json:"base_depth_ouom" default:""`
	Base_strat_unit_id   *string  `json:"base_strat_unit_id" default:""`
	Cement_amount        *float64 `json:"cement_amount" default:""`
	Cement_amount_ouom   *string  `json:"cement_amount_ouom" default:""`
	Cement_amount_uom    *string  `json:"cement_amount_uom" default:""`
	Completion_obs_no    *int     `json:"completion_obs_no" default:""`
	Completion_source    *string  `json:"completion_source" default:""`
	Effective_date       *string  `json:"effective_date" default:""`
	Expiry_date          *string  `json:"expiry_date" default:""`
	Plugback_ba_id       *string  `json:"plugback_ba_id" default:""`
	Plugback_date        *string  `json:"plugback_date" default:""`
	Plug_type            *string  `json:"plug_type" default:""`
	Ppdm_guid            *string  `json:"ppdm_guid" default:""`
	Remark               *string  `json:"remark" default:""`
	Strat_name_set_id    *string  `json:"strat_name_set_id" default:""`
	Surface_abandon_date *string  `json:"surface_abandon_date" default:""`
	Top_depth            *float64 `json:"top_depth" default:""`
	Top_depth_ouom       *string  `json:"top_depth_ouom" default:""`
	Top_found_depth      *float64 `json:"top_found_depth" default:""`
	Top_found_depth_ouom *string  `json:"top_found_depth_ouom" default:""`
	Top_strat_unit_id    *string  `json:"top_strat_unit_id" default:""`
	Row_changed_by       *string  `json:"row_changed_by" default:""`
	Row_changed_date     *string  `json:"row_changed_date" default:""`
	Row_created_by       *string  `json:"row_created_by" default:""`
	Row_created_date     *string  `json:"row_created_date" default:""`
	Row_effective_date   *string  `json:"row_effective_date" default:""`
	Row_expiry_date      *string  `json:"row_expiry_date" default:""`
	Row_quality          *string  `json:"row_quality" default:""`
}
