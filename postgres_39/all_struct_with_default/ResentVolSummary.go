package dto

type Resent_vol_summary struct {
	Resent_id                string   `json:"resent_id" default:"source"`
	Reserve_class_id         string   `json:"reserve_class_id" default:"source"`
	Product_type             string   `json:"product_type" default:"source"`
	Summary_id               string   `json:"summary_id" default:"source"`
	Summary_obs_no           int      `json:"summary_obs_no" default:"1"`
	Active_ind               *string  `json:"active_ind" default:""`
	Analyst_ba_id            *string  `json:"analyst_ba_id" default:""`
	Approved_by_ba_id        *string  `json:"approved_by_ba_id" default:""`
	Approved_date            *string  `json:"approved_date" default:""`
	Created_date             *string  `json:"created_date" default:""`
	Created_date_desc        *string  `json:"created_date_desc" default:""`
	Current_balance          *float64 `json:"current_balance" default:""`
	Current_balance_ouom     *string  `json:"current_balance_ouom" default:""`
	Current_balance_uom      *string  `json:"current_balance_uom" default:""`
	Decline_case_id          *string  `json:"decline_case_id" default:""`
	Effective_date           *string  `json:"effective_date" default:""`
	Expiry_date              *string  `json:"expiry_date" default:""`
	Gross_summary_ind        *string  `json:"gross_summary_ind" default:""`
	Interest_set_id          *string  `json:"interest_set_id" default:""`
	Interest_set_partner     *string  `json:"interest_set_partner" default:""`
	Interest_set_seq_no      *int     `json:"interest_set_seq_no" default:""`
	Material_balance_case_id *string  `json:"material_balance_case_id" default:""`
	Net_summary_ind          *string  `json:"net_summary_ind" default:""`
	Open_balance             *float64 `json:"open_balance" default:""`
	Open_balance_ouom        *string  `json:"open_balance_ouom" default:""`
	Open_balance_uom         *string  `json:"open_balance_uom" default:""`
	Partner_obs_no           *int     `json:"partner_obs_no" default:""`
	Pden_id                  *string  `json:"pden_id" default:""`
	Pden_product_type        *string  `json:"pden_product_type" default:""`
	Pden_source              *string  `json:"pden_source" default:""`
	Pden_subtype             *string  `json:"pden_subtype" default:""`
	Ppdm_guid                *string  `json:"ppdm_guid" default:""`
	Remark                   *string  `json:"remark" default:""`
	Report_ind               *string  `json:"report_ind" default:""`
	Source                   *string  `json:"source" default:""`
	Volume_method            *string  `json:"volume_method" default:""`
	Vol_anal_case_id         *string  `json:"vol_anal_case_id" default:""`
	Yield_parent_product     *string  `json:"yield_parent_product" default:""`
	Yield_rate               *float64 `json:"yield_rate" default:""`
	Yield_rate_ouom          *string  `json:"yield_rate_ouom" default:""`
	Yield_rate_uom           *string  `json:"yield_rate_uom" default:""`
	Row_changed_by           *string  `json:"row_changed_by" default:""`
	Row_changed_date         *string  `json:"row_changed_date" default:""`
	Row_created_by           *string  `json:"row_created_by" default:""`
	Row_created_date         *string  `json:"row_created_date" default:""`
	Row_effective_date       *string  `json:"row_effective_date" default:""`
	Row_expiry_date          *string  `json:"row_expiry_date" default:""`
	Row_quality              *string  `json:"row_quality" default:""`
}
