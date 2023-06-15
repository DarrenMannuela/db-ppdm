package dto

import (
	"time"
)

type Well_test_period struct {
	Uwi                  string     `json:"uwi"`
	Source               string     `json:"source"`
	Test_type            string     `json:"test_type"`
	Run_num              string     `json:"run_num"`
	Test_num             string     `json:"test_num"`
	Period_type          string     `json:"period_type"`
	Period_obs_no        int        `json:"period_obs_no"`
	Active_ind           *string    `json:"active_ind"`
	Casing_pressure      *float64   `json:"casing_pressure"`
	Casing_pressure_ouom *string    `json:"casing_pressure_ouom"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Period_duration      *float64   `json:"period_duration"`
	Period_duration_ouom *string    `json:"period_duration_ouom"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Tubing_pressure      *float64   `json:"tubing_pressure"`
	Tubing_pressure_ouom *string    `json:"tubing_pressure_ouom"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
