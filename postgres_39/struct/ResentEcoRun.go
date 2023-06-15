package dto

import (
	"time"
)

type Resent_eco_run struct {
	Resent_id           string     `json:"resent_id"`
	Reserve_class_id    string     `json:"reserve_class_id"`
	Economics_run_id    string     `json:"economics_run_id"`
	Active_ind          *string    `json:"active_ind"`
	Currency_conversion *float64   `json:"currency_conversion"`
	Currency_ouom       *string    `json:"currency_ouom"`
	Currency_uom        *string    `json:"currency_uom"`
	Economic_forecast   *float64   `json:"economic_forecast"`
	Economic_scenario   *string    `json:"economic_scenario"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Gross_ind           *string    `json:"gross_ind"`
	Interest_set_id     *string    `json:"interest_set_id"`
	Interest_set_seq_no *int       `json:"interest_set_seq_no"`
	Net_ind             *string    `json:"net_ind"`
	Net_present_value   *float64   `json:"net_present_value"`
	Partner_ba_id       *string    `json:"partner_ba_id"`
	Partner_obs_no      *int       `json:"partner_obs_no"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Reserve_life_index  *float64   `json:"reserve_life_index"`
	Source              *string    `json:"source"`
	Tech_forecast       *float64   `json:"tech_forecast"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
