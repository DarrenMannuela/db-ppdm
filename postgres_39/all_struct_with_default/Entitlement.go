package dto

type Entitlement struct {
	Entitlement_id     string   `json:"entitlement_id" default:"source"`
	Access_condition   *string  `json:"access_condition" default:""`
	Access_cond_code   *string  `json:"access_cond_code" default:""`
	Active_ind         *string  `json:"active_ind" default:""`
	Description        *string  `json:"description" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Entitlement_name   *string  `json:"entitlement_name" default:""`
	Entitlement_type   *string  `json:"entitlement_type" default:""`
	Expiry_action      *string  `json:"expiry_action" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Primary_term       *float64 `json:"primary_term" default:""`
	Primary_term_ouom  *string  `json:"primary_term_ouom" default:""`
	Remark             *string  `json:"remark" default:""`
	Security_desc      *string  `json:"security_desc" default:""`
	Source             *string  `json:"source" default:""`
	Use_condition      *string  `json:"use_condition" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
