package dto

type Consent_cond struct {
	Consent_id         string  `json:"consent_id" default:"source"`
	Condition_id       string  `json:"condition_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Component_id       *string `json:"component_id" default:""`
	Consent_condition  *string `json:"consent_condition" default:""`
	Consent_type       *string `json:"consent_type" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
