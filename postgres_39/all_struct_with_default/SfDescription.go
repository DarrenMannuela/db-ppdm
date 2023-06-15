package dto

type Sf_description struct {
	Sf_subtype          string   `json:"sf_subtype" default:"source"`
	Support_facility_id string   `json:"support_facility_id" default:"source"`
	Description_obs_no  int      `json:"description_obs_no" default:"1"`
	Active_ind          *string  `json:"active_ind" default:""`
	Cost                *float64 `json:"cost" default:""`
	Cost_ouom           *string  `json:"cost_ouom" default:""`
	Cost_uom            *string  `json:"cost_uom" default:""`
	Currency_conversion *float64 `json:"currency_conversion" default:""`
	Date_format_desc    *string  `json:"date_format_desc" default:""`
	Description         *string  `json:"description" default:""`
	Desc_type           *string  `json:"desc_type" default:""`
	Desc_value          *float64 `json:"desc_value" default:""`
	Desc_value_code     *string  `json:"desc_value_code" default:""`
	Desc_value_ouom     *string  `json:"desc_value_ouom" default:""`
	Desc_value_uom      *string  `json:"desc_value_uom" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Max_value           *float64 `json:"max_value" default:""`
	Max_value_ouom      *string  `json:"max_value_ouom" default:""`
	Max_value_uom       *string  `json:"max_value_uom" default:""`
	Min_value           *float64 `json:"min_value" default:""`
	Min_value_ouom      *string  `json:"min_value_ouom" default:""`
	Min_value_uom       *string  `json:"min_value_uom" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Remark              *string  `json:"remark" default:""`
	Source              *string  `json:"source" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}
