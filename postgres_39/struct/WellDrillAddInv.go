package dto

import (
	"time"
)

type Well_drill_add_inv struct {
	Uwi                      string     `json:"uwi"`
	Drill_period_obs_no      int        `json:"drill_period_obs_no"`
	Additive_id              string     `json:"additive_id"`
	Active_ind               *string    `json:"active_ind"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Quantity_count_opening   *float64   `json:"quantity_count_opening"`
	Quantity_count_ordered   *float64   `json:"quantity_count_ordered"`
	Quantity_count_remaining *float64   `json:"quantity_count_remaining"`
	Quantity_count_used      *float64   `json:"quantity_count_used"`
	Quantity_value_opening   *float64   `json:"quantity_value_opening"`
	Quantity_value_ordered   *float64   `json:"quantity_value_ordered"`
	Quantity_value_ouom      *string    `json:"quantity_value_ouom"`
	Quantity_value_remaining *float64   `json:"quantity_value_remaining"`
	Quantity_value_uom       *string    `json:"quantity_value_uom"`
	Quantity_value_used      *float64   `json:"quantity_value_used"`
	Remark                   *string    `json:"remark"`
	Report_time_type         *string    `json:"report_time_type"`
	Sf_subtype               *string    `json:"sf_subtype"`
	Source                   *string    `json:"source"`
	Support_facility_id      *string    `json:"support_facility_id"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
