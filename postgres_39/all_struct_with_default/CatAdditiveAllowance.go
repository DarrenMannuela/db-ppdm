package dto

type Cat_additive_allowance struct {
	Allowance_id          string  `json:"allowance_id" default:"source"`
	Active_ind            *string `json:"active_ind" default:""`
	Additive_group_id     *string `json:"additive_group_id" default:""`
	Allowed_ind           *string `json:"allowed_ind" default:""`
	Business_associate_id *string `json:"business_associate_id" default:""`
	Catalogue_additive_id *string `json:"catalogue_additive_id" default:""`
	Disallowed_ind        *string `json:"disallowed_ind" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
	Remark                *string `json:"remark" default:""`
	Source                *string `json:"source" default:""`
	Source_document_id    *string `json:"source_document_id" default:""`
	Substance_id          *string `json:"substance_id" default:""`
	Use_condition         *string `json:"use_condition" default:""`
	Row_changed_by        *string `json:"row_changed_by" default:""`
	Row_changed_date      *string `json:"row_changed_date" default:""`
	Row_created_by        *string `json:"row_created_by" default:""`
	Row_created_date      *string `json:"row_created_date" default:""`
	Row_effective_date    *string `json:"row_effective_date" default:""`
	Row_expiry_date       *string `json:"row_expiry_date" default:""`
	Row_quality           *string `json:"row_quality" default:""`
}
