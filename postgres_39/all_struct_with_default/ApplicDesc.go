package dto

type Applic_desc struct {
	Application_id      string   `json:"application_id" default:"source"`
	Description_id      string   `json:"description_id" default:"source"`
	Active_ind          *string  `json:"active_ind" default:""`
	Application_type    *string  `json:"application_type" default:""`
	Applic_desc_type    *string  `json:"applic_desc_type" default:""`
	Contact_ba_id       *string  `json:"contact_ba_id" default:""`
	Currency_conversion *float64 `json:"currency_conversion" default:""`
	Currency_ouom       *string  `json:"currency_ouom" default:""`
	Currency_uom        *string  `json:"currency_uom" default:""`
	Date_format_desc    *string  `json:"date_format_desc" default:""`
	Description         *string  `json:"description" default:""`
	Desc_date           *string  `json:"desc_date" default:""`
	Desc_value          *float64 `json:"desc_value" default:""`
	Desc_value_ouom     *string  `json:"desc_value_ouom" default:""`
	Desc_value_uom      *string  `json:"desc_value_uom" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Max_cost            *float64 `json:"max_cost" default:""`
	Max_date            *string  `json:"max_date" default:""`
	Max_value           *float64 `json:"max_value" default:""`
	Max_value_ouom      *string  `json:"max_value_ouom" default:""`
	Max_value_uom       *string  `json:"max_value_uom" default:""`
	Min_cost            *float64 `json:"min_cost" default:""`
	Min_date            *string  `json:"min_date" default:""`
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
