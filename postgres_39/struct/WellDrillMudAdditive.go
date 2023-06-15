package dto

import (
	"time"
)

type Well_drill_mud_additive struct {
	Uwi                  string     `json:"uwi"`
	Drill_period_obs_no  int        `json:"drill_period_obs_no"`
	Additive_id          string     `json:"additive_id"`
	Additive_seq_no      int        `json:"additive_seq_no"`
	Active_ind           *string    `json:"active_ind"`
	Additive_method      *string    `json:"additive_method"`
	Additive_period      *float64   `json:"additive_period"`
	Additive_period_ouom *string    `json:"additive_period_ouom"`
	Additive_period_uom  *string    `json:"additive_period_uom"`
	Effective_date       *time.Time `json:"effective_date"`
	End_time             *time.Time `json:"end_time"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Quantity_count       *float64   `json:"quantity_count"`
	Quantity_value       *float64   `json:"quantity_value"`
	Quantity_value_ouom  *string    `json:"quantity_value_ouom"`
	Quantity_value_uom   *string    `json:"quantity_value_uom"`
	Remark               *string    `json:"remark"`
	Source               *string    `json:"source"`
	Start_time           *time.Time `json:"start_time"`
	Timezone             *string    `json:"timezone"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
