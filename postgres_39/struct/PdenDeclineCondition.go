package dto

import (
	"time"
)

type Pden_decline_condition struct {
	Pden_id            string     `json:"pden_id"`
	Pden_subtype       string     `json:"pden_subtype"`
	Pden_source        string     `json:"pden_source"`
	Product_type       string     `json:"product_type"`
	Case_id            string     `json:"case_id"`
	Condition_id       string     `json:"condition_id"`
	Period_type        string     `json:"period_type"`
	Period_id          string     `json:"period_id"`
	Active_ind         *string    `json:"active_ind"`
	Condition_code     *string    `json:"condition_code"`
	Condition_desc     *string    `json:"condition_desc"`
	Condition_type     *string    `json:"condition_type"`
	Condition_value    *float64   `json:"condition_value"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Volume_date        *time.Time `json:"volume_date"`
	Volume_date_desc   *string    `json:"volume_date_desc"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
