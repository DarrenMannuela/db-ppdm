package dto

type Cont_provision_text struct {
	Contract_id           string  `json:"contract_id" default:"source"`
	Provision_id          string  `json:"provision_id" default:"source"`
	Provision_text_seq_no int     `json:"provision_text_seq_no" default:"1"`
	Active_ind            *string `json:"active_ind" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
	Provision_text        *string `json:"provision_text" default:""`
	Remark                *string `json:"remark" default:""`
	Source                *string `json:"source" default:""`
	Row_changed_by        *string `json:"row_changed_by" default:""`
	Row_changed_date      *string `json:"row_changed_date" default:""`
	Row_created_by        *string `json:"row_created_by" default:""`
	Row_created_date      *string `json:"row_created_date" default:""`
	Row_effective_date    *string `json:"row_effective_date" default:""`
	Row_expiry_date       *string `json:"row_expiry_date" default:""`
	Row_quality           *string `json:"row_quality" default:""`
}
