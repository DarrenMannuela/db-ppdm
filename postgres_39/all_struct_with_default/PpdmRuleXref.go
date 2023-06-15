package dto

type Ppdm_rule_xref struct {
	Dependency_id       string  `json:"dependency_id" default:"source"`
	Rule_id             string  `json:"rule_id" default:"source"`
	Rule_id2            string  `json:"rule_id_2" default:"source"`
	Xref_obs_no         int     `json:"xref_obs_no" default:"1"`
	Active_ind          *string `json:"active_ind" default:""`
	Description         *string `json:"description" default:""`
	Effective_date      *string `json:"effective_date" default:""`
	Expiry_date         *string `json:"expiry_date" default:""`
	Ppdm_guid           *string `json:"ppdm_guid" default:""`
	Remark              *string `json:"remark" default:""`
	Rule_xref_condition *string `json:"rule_xref_condition" default:""`
	Rule_xref_type      *string `json:"rule_xref_type" default:""`
	Source              *string `json:"source" default:""`
	Row_changed_by      *string `json:"row_changed_by" default:""`
	Row_changed_date    *string `json:"row_changed_date" default:""`
	Row_created_by      *string `json:"row_created_by" default:""`
	Row_created_date    *string `json:"row_created_date" default:""`
	Row_effective_date  *string `json:"row_effective_date" default:""`
	Row_expiry_date     *string `json:"row_expiry_date" default:""`
	Row_quality         *string `json:"row_quality" default:""`
}
