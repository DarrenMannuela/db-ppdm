package dto

import (
	"time"
)

type Well_test_press_meas struct {
	Uwi                           string     `json:"uwi"`
	Source                        string     `json:"source"`
	Test_type                     string     `json:"test_type"`
	Run_num                       string     `json:"run_num"`
	Test_num                      string     `json:"test_num"`
	Measurement_obs_no            int        `json:"measurement_obs_no"`
	Active_ind                    *string    `json:"active_ind"`
	Effective_date                *time.Time `json:"effective_date"`
	Expiry_date                   *time.Time `json:"expiry_date"`
	Measurement_pressure          *float64   `json:"measurement_pressure"`
	Measurement_pressure_ouom     *string    `json:"measurement_pressure_ouom"`
	Measurement_temperature       *float64   `json:"measurement_temperature"`
	Measurement_temp_ouom         *string    `json:"measurement_temp_ouom"`
	Measurement_time_elapsed      *float64   `json:"measurement_time_elapsed"`
	Measurement_time_elapsed_ouom *string    `json:"measurement_time_elapsed_ouom"`
	Period_obs_no                 *int       `json:"period_obs_no"`
	Period_type                   *string    `json:"period_type"`
	Ppdm_guid                     string     `json:"ppdm_guid"`
	Recorder_id                   *string    `json:"recorder_id"`
	Remark                        *string    `json:"remark"`
	Row_changed_by                *string    `json:"row_changed_by"`
	Row_changed_date              *time.Time `json:"row_changed_date"`
	Row_created_by                *string    `json:"row_created_by"`
	Row_created_date              *time.Time `json:"row_created_date"`
	Row_effective_date            *time.Time `json:"row_effective_date"`
	Row_expiry_date               *time.Time `json:"row_expiry_date"`
	Row_quality                   *string    `json:"row_quality"`
}
