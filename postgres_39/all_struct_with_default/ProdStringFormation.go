package dto

type Prod_string_formation struct {
	Uwi                 string   `json:"uwi" default:"source"`
	Prod_string_source  string   `json:"prod_string_source" default:"source"`
	String_id           string   `json:"string_id" default:"source"`
	Pr_str_form_obs_no  int      `json:"pr_str_form_obs_no" default:"1"`
	Active_ind          *string  `json:"active_ind" default:""`
	Allocation_factor   *float64 `json:"allocation_factor" default:""`
	Allocation_type     *string  `json:"allocation_type" default:""`
	Base_depth          *float64 `json:"base_depth" default:""`
	Base_depth_ouom     *string  `json:"base_depth_ouom" default:""`
	Completion_obs_no   *int     `json:"completion_obs_no" default:""`
	Current_status      *string  `json:"current_status" default:""`
	Current_status_date *string  `json:"current_status_date" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Pool_id             *string  `json:"pool_id" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Remark              *string  `json:"remark" default:""`
	Source              *string  `json:"source" default:""`
	Status_type         *string  `json:"status_type" default:""`
	Strat_name_set_id   *string  `json:"strat_name_set_id" default:""`
	Strat_unit_id       *string  `json:"strat_unit_id" default:""`
	Top_depth           *float64 `json:"top_depth" default:""`
	Top_depth_ouom      *string  `json:"top_depth_ouom" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}
