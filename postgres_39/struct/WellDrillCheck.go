package dto

import (
	"time"
)

type Well_drill_check struct {
	Uwi                   string     `json:"uwi"`
	Period_obs_no         int        `json:"period_obs_no"`
	Drill_check_obs_no    int        `json:"drill_check_obs_no"`
	Active_ind            *string    `json:"active_ind"`
	Check_set_id          *string    `json:"check_set_id"`
	Check_type            *string    `json:"check_type"`
	Contractor_name       *string    `json:"contractor_name"`
	Contractor_rep_ba_id  *string    `json:"contractor_rep_ba_id"`
	Deficient_ind         *string    `json:"deficient_ind"`
	Deficient_period      *float64   `json:"deficient_period"`
	Deficient_period_ouom *string    `json:"deficient_period_ouom"`
	Deficient_period_uom  *string    `json:"deficient_period_uom"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Operator_name         *string    `json:"operator_name"`
	Operator_rep_ba_id    *string    `json:"operator_rep_ba_id"`
	Passed_ind            *string    `json:"passed_ind"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Recorded_time         *time.Time `json:"recorded_time"`
	Recorded_timezone     *string    `json:"recorded_timezone"`
	Remark                *string    `json:"remark"`
	Source                *string    `json:"source"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
