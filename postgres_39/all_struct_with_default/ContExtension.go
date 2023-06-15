package dto

type Cont_extension struct {
	Contract_id        string  `json:"contract_id" default:"source"`
	Extension_id       string  `json:"extension_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Description        *string `json:"description" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Extension_type     *string `json:"extension_type" default:""`
	Issued_by          *string `json:"issued_by" default:""`
	Issued_date        *string `json:"issued_date" default:""`
	Land_right_id      *string `json:"land_right_id" default:""`
	Land_right_subtype *string `json:"land_right_subtype" default:""`
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
