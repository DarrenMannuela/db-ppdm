package dto

import (
	"time"
)

type Resent_eco_volume struct {
	Resent_id                   string     `json:"resent_id"`
	Reserve_class_id            string     `json:"reserve_class_id"`
	Economics_run_id            string     `json:"economics_run_id"`
	Product_type                string     `json:"product_type"`
	Summary_obs_no              int        `json:"summary_obs_no"`
	Active_ind                  *string    `json:"active_ind"`
	Effective_date              *time.Time `json:"effective_date"`
	Expiry_date                 *time.Time `json:"expiry_date"`
	Ppdm_guid                   string     `json:"ppdm_guid"`
	Remaining_balance           *float64   `json:"remaining_balance"`
	Remaining_balance_date      *time.Time `json:"remaining_balance_date"`
	Remaining_balance_date_desc *string    `json:"remaining_balance_date_desc"`
	Remaining_balance_ouom      *string    `json:"remaining_balance_ouom"`
	Remaining_balance_uom       *string    `json:"remaining_balance_uom"`
	Remark                      *string    `json:"remark"`
	Source                      *string    `json:"source"`
	Row_changed_by              *string    `json:"row_changed_by"`
	Row_changed_date            *time.Time `json:"row_changed_date"`
	Row_created_by              *string    `json:"row_created_by"`
	Row_created_date            *time.Time `json:"row_created_date"`
	Row_effective_date          *time.Time `json:"row_effective_date"`
	Row_expiry_date             *time.Time `json:"row_expiry_date"`
	Row_quality                 *string    `json:"row_quality"`
}
