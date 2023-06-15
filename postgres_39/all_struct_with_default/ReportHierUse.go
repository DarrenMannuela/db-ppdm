package dto

type Report_hier_use struct {
	Report_hierarchy_id  string   `json:"report_hierarchy_id" default:"source"`
	Component_id         string   `json:"component_id" default:"source"`
	Hierarchy_use_obs_no int      `json:"hierarchy_use_obs_no" default:"1"`
	Active_ind           *string  `json:"active_ind" default:""`
	Contribution_percent *float64 `json:"contribution_percent" default:""`
	Effective_date       *string  `json:"effective_date" default:""`
	Expiry_date          *string  `json:"expiry_date" default:""`
	Pden_id              *string  `json:"pden_id" default:""`
	Pden_source          *string  `json:"pden_source" default:""`
	Pden_subtype         *string  `json:"pden_subtype" default:""`
	Ppdm_guid            *string  `json:"ppdm_guid" default:""`
	Remark               *string  `json:"remark" default:""`
	Report_ind           *string  `json:"report_ind" default:""`
	Resent_id            *string  `json:"resent_id" default:""`
	Source               *string  `json:"source" default:""`
	Substance_id         *string  `json:"substance_id" default:""`
	Row_changed_by       *string  `json:"row_changed_by" default:""`
	Row_changed_date     *string  `json:"row_changed_date" default:""`
	Row_created_by       *string  `json:"row_created_by" default:""`
	Row_created_date     *string  `json:"row_created_date" default:""`
	Row_effective_date   *string  `json:"row_effective_date" default:""`
	Row_expiry_date      *string  `json:"row_expiry_date" default:""`
	Row_quality          *string  `json:"row_quality" default:""`
}
