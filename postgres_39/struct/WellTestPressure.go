package dto

import (
	"time"
)

type Well_test_pressure struct {
	Uwi                 string     `json:"uwi"`
	Source              string     `json:"source"`
	Test_type           string     `json:"test_type"`
	Run_num             string     `json:"run_num"`
	Test_num            string     `json:"test_num"`
	Period_type         string     `json:"period_type"`
	Period_obs_no       int        `json:"period_obs_no"`
	Active_ind          *string    `json:"active_ind"`
	Effective_date      *time.Time `json:"effective_date"`
	End_pressure        *float64   `json:"end_pressure"`
	End_pressure_ouom   *string    `json:"end_pressure_ouom"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Quality_code        *string    `json:"quality_code"`
	Recorder_id         *string    `json:"recorder_id"`
	Remark              *string    `json:"remark"`
	Start_pressure      *float64   `json:"start_pressure"`
	Start_pressure_ouom *string    `json:"start_pressure_ouom"`
	Summary_ind         *string    `json:"summary_ind"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
