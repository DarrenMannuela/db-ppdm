package dto

type Resent_eco_volume struct {
	Resent_id                   string   `json:"resent_id" default:"source"`
	Reserve_class_id            string   `json:"reserve_class_id" default:"source"`
	Economics_run_id            string   `json:"economics_run_id" default:"source"`
	Product_type                string   `json:"product_type" default:"source"`
	Summary_obs_no              int      `json:"summary_obs_no" default:"1"`
	Active_ind                  *string  `json:"active_ind" default:""`
	Effective_date              *string  `json:"effective_date" default:""`
	Expiry_date                 *string  `json:"expiry_date" default:""`
	Ppdm_guid                   *string  `json:"ppdm_guid" default:""`
	Remaining_balance           *float64 `json:"remaining_balance" default:""`
	Remaining_balance_date      *string  `json:"remaining_balance_date" default:""`
	Remaining_balance_date_desc *string  `json:"remaining_balance_date_desc" default:""`
	Remaining_balance_ouom      *string  `json:"remaining_balance_ouom" default:""`
	Remaining_balance_uom       *string  `json:"remaining_balance_uom" default:""`
	Remark                      *string  `json:"remark" default:""`
	Source                      *string  `json:"source" default:""`
	Row_changed_by              *string  `json:"row_changed_by" default:""`
	Row_changed_date            *string  `json:"row_changed_date" default:""`
	Row_created_by              *string  `json:"row_created_by" default:""`
	Row_created_date            *string  `json:"row_created_date" default:""`
	Row_effective_date          *string  `json:"row_effective_date" default:""`
	Row_expiry_date             *string  `json:"row_expiry_date" default:""`
	Row_quality                 *string  `json:"row_quality" default:""`
}
