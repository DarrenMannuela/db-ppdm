package dto

type Pden_alloc_factor struct {
	From_pden_subtype  string   `json:"from_pden_subtype" default:"source"`
	From_pden_id       string   `json:"from_pden_id" default:"source"`
	From_pden_source   string   `json:"from_pden_source" default:"source"`
	To_pden_subtype    string   `json:"to_pden_subtype" default:"source"`
	To_pden_id         string   `json:"to_pden_id" default:"source"`
	To_pden_source     string   `json:"to_pden_source" default:"source"`
	Allocation_obs_no  int      `json:"allocation_obs_no" default:"1"`
	Active_ind         *string  `json:"active_ind" default:""`
	Activity_type      *string  `json:"activity_type" default:""`
	Allocation_factor  *float64 `json:"allocation_factor" default:""`
	Allocation_type    *string  `json:"allocation_type" default:""`
	Business_rule      *string  `json:"business_rule" default:""`
	Calculated_from    *string  `json:"calculated_from" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Period_type        *string  `json:"period_type" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Product_type       *string  `json:"product_type" default:""`
	Remark             *string  `json:"remark" default:""`
	Source             *string  `json:"source" default:""`
	Volume_method      *string  `json:"volume_method" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
