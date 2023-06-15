package dto

type Anl_coal_rank struct {
	Coal_rank_id        string   `json:"coal_rank_id" default:"source"`
	Abbreviation        *string  `json:"abbreviation" default:""`
	Active_ind          *string  `json:"active_ind" default:""`
	Coal_rank_scheme_id *string  `json:"coal_rank_scheme_id" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Long_name           *string  `json:"long_name" default:""`
	Max_value           *float64 `json:"max_value" default:""`
	Max_value_ouom      *string  `json:"max_value_ouom" default:""`
	Max_value_uom       *string  `json:"max_value_uom" default:""`
	Min_value           *float64 `json:"min_value" default:""`
	Min_value_ouom      *string  `json:"min_value_ouom" default:""`
	Min_value_uom       *string  `json:"min_value_uom" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Remark              *string  `json:"remark" default:""`
	Short_name          *string  `json:"short_name" default:""`
	Source              *string  `json:"source" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}
