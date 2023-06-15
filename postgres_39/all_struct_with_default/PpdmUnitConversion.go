package dto

type Ppdm_unit_conversion struct {
	From_uom_id          string   `json:"from_uom_id" default:"source"`
	To_uom_id            string   `json:"to_uom_id" default:"source"`
	Active_ind           *string  `json:"active_ind" default:""`
	Effective_date       *string  `json:"effective_date" default:""`
	Exact_conversion_ind *string  `json:"exact_conversion_ind" default:""`
	Expiry_date          *string  `json:"expiry_date" default:""`
	Factor_denominator   *float64 `json:"factor_denominator" default:""`
	Factor_numerator     *float64 `json:"factor_numerator" default:""`
	Post_offset          *float64 `json:"post_offset" default:""`
	Ppdm_guid            *string  `json:"ppdm_guid" default:""`
	Pre_offset           *float64 `json:"pre_offset" default:""`
	Remark               *string  `json:"remark" default:""`
	Source               *string  `json:"source" default:""`
	Source_document_id   *string  `json:"source_document_id" default:""`
	Unit_expression      *float64 `json:"unit_expression" default:""`
	Unit_quantity_type   *string  `json:"unit_quantity_type" default:""`
	Row_changed_by       *string  `json:"row_changed_by" default:""`
	Row_changed_date     *string  `json:"row_changed_date" default:""`
	Row_created_by       *string  `json:"row_created_by" default:""`
	Row_created_date     *string  `json:"row_created_date" default:""`
	Row_effective_date   *string  `json:"row_effective_date" default:""`
	Row_expiry_date      *string  `json:"row_expiry_date" default:""`
	Row_quality          *string  `json:"row_quality" default:""`
}
