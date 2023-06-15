package dto

type Ppdm_unit_of_measure struct {
	Uom_id             string  `json:"uom_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Base_unit_ind      *string `json:"base_unit_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Source_document_id *string `json:"source_document_id" default:""`
	Uom_name           *string `json:"uom_name" default:""`
	Uom_quantity_type  *string `json:"uom_quantity_type" default:""`
	Uom_system_id      *string `json:"uom_system_id" default:""`
	Uom_usage_type     *string `json:"uom_usage_type" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
