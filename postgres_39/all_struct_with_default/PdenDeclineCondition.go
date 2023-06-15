package dto

type Pden_decline_condition struct {
	Pden_id            string   `json:"pden_id" default:"source"`
	Pden_subtype       string   `json:"pden_subtype" default:"source"`
	Pden_source        string   `json:"pden_source" default:"source"`
	Product_type       string   `json:"product_type" default:"source"`
	Case_id            string   `json:"case_id" default:"source"`
	Condition_id       string   `json:"condition_id" default:"source"`
	Period_type        string   `json:"period_type" default:"source"`
	Period_id          string   `json:"period_id" default:"source"`
	Active_ind         *string  `json:"active_ind" default:""`
	Condition_code     *string  `json:"condition_code" default:""`
	Condition_desc     *string  `json:"condition_desc" default:""`
	Condition_type     *string  `json:"condition_type" default:""`
	Condition_value    *float64 `json:"condition_value" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Remark             *string  `json:"remark" default:""`
	Source             *string  `json:"source" default:""`
	Volume_date        *string  `json:"volume_date" default:""`
	Volume_date_desc   *string  `json:"volume_date_desc" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
