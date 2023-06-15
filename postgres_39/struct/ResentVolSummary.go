package dto

import (
	"time"
)

type Resent_vol_summary struct {
	Resent_id                string     `json:"resent_id"`
	Reserve_class_id         string     `json:"reserve_class_id"`
	Product_type             string     `json:"product_type"`
	Summary_id               string     `json:"summary_id"`
	Summary_obs_no           int        `json:"summary_obs_no"`
	Active_ind               *string    `json:"active_ind"`
	Analyst_ba_id            *string    `json:"analyst_ba_id"`
	Approved_by_ba_id        *string    `json:"approved_by_ba_id"`
	Approved_date            *time.Time `json:"approved_date"`
	Created_date             *time.Time `json:"created_date"`
	Created_date_desc        *string    `json:"created_date_desc"`
	Current_balance          *float64   `json:"current_balance"`
	Current_balance_ouom     *string    `json:"current_balance_ouom"`
	Current_balance_uom      *string    `json:"current_balance_uom"`
	Decline_case_id          *string    `json:"decline_case_id"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Gross_summary_ind        *string    `json:"gross_summary_ind"`
	Interest_set_id          *string    `json:"interest_set_id"`
	Interest_set_partner     *string    `json:"interest_set_partner"`
	Interest_set_seq_no      *int       `json:"interest_set_seq_no"`
	Material_balance_case_id *string    `json:"material_balance_case_id"`
	Net_summary_ind          *string    `json:"net_summary_ind"`
	Open_balance             *float64   `json:"open_balance"`
	Open_balance_ouom        *string    `json:"open_balance_ouom"`
	Open_balance_uom         *string    `json:"open_balance_uom"`
	Partner_obs_no           *int       `json:"partner_obs_no"`
	Pden_id                  *string    `json:"pden_id"`
	Pden_product_type        *string    `json:"pden_product_type"`
	Pden_source              *string    `json:"pden_source"`
	Pden_subtype             *string    `json:"pden_subtype"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Remark                   *string    `json:"remark"`
	Report_ind               *string    `json:"report_ind"`
	Source                   *string    `json:"source"`
	Volume_method            *string    `json:"volume_method"`
	Vol_anal_case_id         *string    `json:"vol_anal_case_id"`
	Yield_parent_product     *string    `json:"yield_parent_product"`
	Yield_rate               *float64   `json:"yield_rate"`
	Yield_rate_ouom          *string    `json:"yield_rate_ouom"`
	Yield_rate_uom           *string    `json:"yield_rate_uom"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
