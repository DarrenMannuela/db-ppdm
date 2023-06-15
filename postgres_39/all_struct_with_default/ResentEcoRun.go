package dto

type Resent_eco_run struct {
	Resent_id           string   `json:"resent_id" default:"source"`
	Reserve_class_id    string   `json:"reserve_class_id" default:"source"`
	Economics_run_id    string   `json:"economics_run_id" default:"source"`
	Active_ind          *string  `json:"active_ind" default:""`
	Currency_conversion *float64 `json:"currency_conversion" default:""`
	Currency_ouom       *string  `json:"currency_ouom" default:""`
	Currency_uom        *string  `json:"currency_uom" default:""`
	Economic_forecast   *float64 `json:"economic_forecast" default:""`
	Economic_scenario   *string  `json:"economic_scenario" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Gross_ind           *string  `json:"gross_ind" default:""`
	Interest_set_id     *string  `json:"interest_set_id" default:""`
	Interest_set_seq_no *int     `json:"interest_set_seq_no" default:""`
	Net_ind             *string  `json:"net_ind" default:""`
	Net_present_value   *float64 `json:"net_present_value" default:""`
	Partner_ba_id       *string  `json:"partner_ba_id" default:""`
	Partner_obs_no      *int     `json:"partner_obs_no" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Remark              *string  `json:"remark" default:""`
	Reserve_life_index  *float64 `json:"reserve_life_index" default:""`
	Source              *string  `json:"source" default:""`
	Tech_forecast       *float64 `json:"tech_forecast" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}
