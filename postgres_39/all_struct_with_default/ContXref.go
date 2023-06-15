package dto

type Cont_xref struct {
	Contract_id           string  `json:"contract_id" default:"source"`
	Contract_id_2         string  `json:"contract_id_2" default:"source"`
	Active_ind            *string `json:"active_ind" default:""`
	Contract_xref_type    *string `json:"contract_xref_type" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	Order_of_supercedence *int    `json:"order_of_supercedence" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
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
