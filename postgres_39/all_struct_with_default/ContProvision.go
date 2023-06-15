package dto

type Cont_provision struct {
	Contract_id         string  `json:"contract_id" default:"source"`
	Provision_id        string  `json:"provision_id" default:"source"`
	Active_ind          *string `json:"active_ind" default:""`
	Clause_heading      *string `json:"clause_heading" default:""`
	Clause_number       *string `json:"clause_number" default:""`
	Cont_provision_type *string `json:"cont_provision_type" default:""`
	Effective_date      *string `json:"effective_date" default:""`
	Expiry_date         *string `json:"expiry_date" default:""`
	Modified_ind        *string `json:"modified_ind" default:""`
	Ppdm_guid           *string `json:"ppdm_guid" default:""`
	Provision_desc      *string `json:"provision_desc" default:""`
	Remark              *string `json:"remark" default:""`
	Source              *string `json:"source" default:""`
	Source_document_id  *string `json:"source_document_id" default:""`
	Row_changed_by      *string `json:"row_changed_by" default:""`
	Row_changed_date    *string `json:"row_changed_date" default:""`
	Row_created_by      *string `json:"row_created_by" default:""`
	Row_created_date    *string `json:"row_created_date" default:""`
	Row_effective_date  *string `json:"row_effective_date" default:""`
	Row_expiry_date     *string `json:"row_expiry_date" default:""`
	Row_quality         *string `json:"row_quality" default:""`
}
