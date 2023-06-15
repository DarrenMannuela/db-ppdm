package dto

import (
	"time"
)

type Pden_alloc_factor struct {
	From_pden_subtype  string     `json:"from_pden_subtype"`
	From_pden_id       string     `json:"from_pden_id"`
	From_pden_source   string     `json:"from_pden_source"`
	To_pden_subtype    string     `json:"to_pden_subtype"`
	To_pden_id         string     `json:"to_pden_id"`
	To_pden_source     string     `json:"to_pden_source"`
	Allocation_obs_no  int        `json:"allocation_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Activity_type      *string    `json:"activity_type"`
	Allocation_factor  *float64   `json:"allocation_factor"`
	Allocation_type    *string    `json:"allocation_type"`
	Business_rule      *string    `json:"business_rule"`
	Calculated_from    *string    `json:"calculated_from"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Period_type        *string    `json:"period_type"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Product_type       *string    `json:"product_type"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Volume_method      *string    `json:"volume_method"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
