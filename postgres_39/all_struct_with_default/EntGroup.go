package dto

type Ent_group struct {
	Entitlement_id     string  `json:"entitlement_id" default:"source"`
	Security_group_id  string  `json:"security_group_id" default:"source"`
	Group_obs_no       int     `json:"group_obs_no" default:"1"`
	Access_condition   *string `json:"access_condition" default:""`
	Access_type        *string `json:"access_type" default:""`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Restriction_desc   *string `json:"restriction_desc" default:""`
	Security_desc      *string `json:"security_desc" default:""`
	Source             *string `json:"source" default:""`
	Use_desc           *string `json:"use_desc" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
