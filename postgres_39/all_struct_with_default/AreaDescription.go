package dto

type Area_description struct {
	Area_id               string   `json:"area_id" default:"source"`
	Area_type             string   `json:"area_type" default:"source"`
	Description_obs_no    int      `json:"description_obs_no" default:"1"`
	Active_ind            *string  `json:"active_ind" default:""`
	Area_description_date *string  `json:"area_description_date" default:""`
	Area_desc_code        *string  `json:"area_desc_code" default:""`
	Area_desc_type        *string  `json:"area_desc_type" default:""`
	Average_value         *float64 `json:"average_value" default:""`
	Average_value_ouom    *string  `json:"average_value_ouom" default:""`
	Average_value_uom     *string  `json:"average_value_uom" default:""`
	Date_format_desc      *string  `json:"date_format_desc" default:""`
	Description           *string  `json:"description" default:""`
	Effective_date        *string  `json:"effective_date" default:""`
	Expiry_date           *string  `json:"expiry_date" default:""`
	Max_value             *float64 `json:"max_value" default:""`
	Max_value_ouom        *string  `json:"max_value_ouom" default:""`
	Max_value_uom         *string  `json:"max_value_uom" default:""`
	Min_value             *float64 `json:"min_value" default:""`
	Min_value_ouom        *string  `json:"min_value_ouom" default:""`
	Min_value_uom         *string  `json:"min_value_uom" default:""`
	Ppdm_guid             *string  `json:"ppdm_guid" default:""`
	Remark                *string  `json:"remark" default:""`
	Source                *string  `json:"source" default:""`
	Source_document_id    *string  `json:"source_document_id" default:""`
	Row_changed_by        *string  `json:"row_changed_by" default:""`
	Row_changed_date      *string  `json:"row_changed_date" default:""`
	Row_created_by        *string  `json:"row_created_by" default:""`
	Row_created_date      *string  `json:"row_created_date" default:""`
	Row_effective_date    *string  `json:"row_effective_date" default:""`
	Row_expiry_date       *string  `json:"row_expiry_date" default:""`
	Row_quality           *string  `json:"row_quality" default:""`
}
