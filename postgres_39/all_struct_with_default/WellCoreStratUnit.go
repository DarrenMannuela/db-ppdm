package dto

type Well_core_strat_unit struct {
	Uwi                     string   `json:"uwi" default:"source"`
	Source                  string   `json:"source" default:"source"`
	Core_id                 string   `json:"core_id" default:"source"`
	Description_obs_no      int      `json:"description_obs_no" default:"1"`
	Strat_name_set_id       string   `json:"strat_name_set_id" default:"source"`
	Core_strat_unit_id      string   `json:"core_strat_unit_id" default:"source"`
	Active_ind              *string  `json:"active_ind" default:""`
	Core_strat_unit_md      *float64 `json:"core_strat_unit_md" default:""`
	Core_strat_unit_md_ouom *string  `json:"core_strat_unit_md_ouom" default:""`
	Effective_date          *string  `json:"effective_date" default:""`
	Expiry_date             *string  `json:"expiry_date" default:""`
	Ppdm_guid               *string  `json:"ppdm_guid" default:""`
	Remark                  *string  `json:"remark" default:""`
	Row_changed_by          *string  `json:"row_changed_by" default:""`
	Row_changed_date        *string  `json:"row_changed_date" default:""`
	Row_created_by          *string  `json:"row_created_by" default:""`
	Row_created_date        *string  `json:"row_created_date" default:""`
	Row_effective_date      *string  `json:"row_effective_date" default:""`
	Row_expiry_date         *string  `json:"row_expiry_date" default:""`
	Row_quality             *string  `json:"row_quality" default:""`
}
