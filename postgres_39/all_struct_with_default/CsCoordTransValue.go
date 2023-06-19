package dto

type Cs_coord_trans_value struct {
	Transform_id       string   `json:"transform_id" default:"source"`
	Transform_type     string   `json:"transform_type" default:"source"`
	Parameter_id       string   `json:"parameter_id" default:"source"`
	Value_id           string   `json:"value_id" default:"source"`
	Active_ind         *string  `json:"active_ind" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Remark             *string  `json:"remark" default:""`
	Source             *string  `json:"source" default:""`
	Source_document_id *string  `json:"source_document_id" default:""`
	Text_value         *string  `json:"text_value" default:""`
	Trans_value        *float64 `json:"trans_value" default:""`
	Trans_value_uom    *string  `json:"trans_value_uom" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}