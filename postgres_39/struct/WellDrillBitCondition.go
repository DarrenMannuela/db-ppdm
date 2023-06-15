package dto

import (
	"time"
)

type Well_drill_bit_condition struct {
	Uwi                 string     `json:"uwi"`
	Bit_source          string     `json:"bit_source"`
	Bit_interval_obs_no int        `json:"bit_interval_obs_no"`
	Condition_obs_no    int        `json:"condition_obs_no"`
	Active_ind          *string    `json:"active_ind"`
	Bearing_condition   *string    `json:"bearing_condition"`
	Drill_bit_condition *string    `json:"drill_bit_condition"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Source              *string    `json:"source"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
