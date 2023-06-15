package dto

type Consent struct {
	Consent_id          string  `json:"consent_id" default:"source"`
	Active_ind          *string `json:"active_ind" default:""`
	Consent_desc        *string `json:"consent_desc" default:""`
	Consent_method_desc *string `json:"consent_method_desc" default:""`
	Consent_type        *string `json:"consent_type" default:""`
	Current_status      *string `json:"current_status" default:""`
	Effective_date      *string `json:"effective_date" default:""`
	Expiry_date         *string `json:"expiry_date" default:""`
	Ppdm_guid           *string `json:"ppdm_guid" default:""`
	Receive_date        *string `json:"receive_date" default:""`
	Remark              *string `json:"remark" default:""`
	Request_date        *string `json:"request_date" default:""`
	Source              *string `json:"source" default:""`
	Status_remark       *string `json:"status_remark" default:""`
	Row_changed_by      *string `json:"row_changed_by" default:""`
	Row_changed_date    *string `json:"row_changed_date" default:""`
	Row_created_by      *string `json:"row_created_by" default:""`
	Row_created_date    *string `json:"row_created_date" default:""`
	Row_effective_date  *string `json:"row_effective_date" default:""`
	Row_expiry_date     *string `json:"row_expiry_date" default:""`
	Row_quality         *string `json:"row_quality" default:""`
}
