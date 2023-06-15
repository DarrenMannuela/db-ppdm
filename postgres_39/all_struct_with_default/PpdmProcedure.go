package dto

type Ppdm_procedure struct {
	System_id          string  `json:"system_id" default:"source"`
	Procedure_id       string  `json:"procedure_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Procedure_desc     *string `json:"procedure_desc" default:""`
	Procedure_name     *string `json:"procedure_name" default:""`
	Procedure_type     *string `json:"procedure_type" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Table_name         *string `json:"table_name" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
