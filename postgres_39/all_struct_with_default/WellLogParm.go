package dto

type Well_log_parm struct {
	Uwi                  string   `json:"uwi" default:"source"`
	Well_log_id          string   `json:"well_log_id" default:"source"`
	Well_log_source      string   `json:"well_log_source" default:"source"`
	Parameter_seq_no     int      `json:"parameter_seq_no" default:"1"`
	Active_ind           *string  `json:"active_ind" default:""`
	Array_ind            *string  `json:"array_ind" default:""`
	Description          *string  `json:"description" default:""`
	Dictionary_id        *string  `json:"dictionary_id" default:""`
	Dict_parameter_id    *string  `json:"dict_parameter_id" default:""`
	Effective_date       *string  `json:"effective_date" default:""`
	Expiry_date          *string  `json:"expiry_date" default:""`
	Parameter_output     *string  `json:"parameter_output" default:""`
	Parameter_text_value *string  `json:"parameter_text_value" default:""`
	Parameter_value      *float64 `json:"parameter_value" default:""`
	Parameter_value_ouom *string  `json:"parameter_value_ouom" default:""`
	Parameter_value_uom  *string  `json:"parameter_value_uom" default:""`
	Ppdm_guid            *string  `json:"ppdm_guid" default:""`
	Remark               *string  `json:"remark" default:""`
	Reported_desc        *string  `json:"reported_desc" default:""`
	Reported_mnemonic    *string  `json:"reported_mnemonic" default:""`
	Source               *string  `json:"source" default:""`
	Row_changed_by       *string  `json:"row_changed_by" default:""`
	Row_changed_date     *string  `json:"row_changed_date" default:""`
	Row_created_by       *string  `json:"row_created_by" default:""`
	Row_created_date     *string  `json:"row_created_date" default:""`
	Row_effective_date   *string  `json:"row_effective_date" default:""`
	Row_expiry_date      *string  `json:"row_expiry_date" default:""`
	Row_quality          *string  `json:"row_quality" default:""`
}
