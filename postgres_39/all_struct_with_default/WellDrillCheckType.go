package dto

type Well_drill_check_type struct {
	Check_set_id       string  `json:"check_set_id" default:"source"`
	Check_type         string  `json:"check_type" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Contractor_ind     *string `json:"contractor_ind" default:""`
	Description        *string `json:"description" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Frequency          *string `json:"frequency" default:""`
	Operator_ind       *string `json:"operator_ind" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Source_document_id *string `json:"source_document_id" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
