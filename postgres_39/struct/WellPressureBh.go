package dto

import (
	"time"
)

type Well_pressure_bh struct {
	Uwi                        string     `json:"uwi"`
	Source                     string     `json:"source"`
	Pressure_obs_no            int        `json:"pressure_obs_no"`
	Bhp_obs_no                 int        `json:"bhp_obs_no"`
	Active_ind                 *string    `json:"active_ind"`
	Bhfp                       *float64   `json:"bhfp"`
	Bhfp_ouom                  *string    `json:"bhfp_ouom"`
	Bhp_method                 *string    `json:"bhp_method"`
	Bh_test_code               *string    `json:"bh_test_code"`
	Datum_pressure             *float64   `json:"datum_pressure"`
	Datum_pressure_ouom        *string    `json:"datum_pressure_ouom"`
	Effective_date             *time.Time `json:"effective_date"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Packer_depth               *float64   `json:"packer_depth"`
	Packer_depth_ouom          *string    `json:"packer_depth_ouom"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Pressure_gradient          *float64   `json:"pressure_gradient"`
	Pressure_gradient_ouom     *string    `json:"pressure_gradient_ouom"`
	Recorder_datum             *float64   `json:"recorder_datum"`
	Recorder_datum_ouom        *string    `json:"recorder_datum_ouom"`
	Remark                     *string    `json:"remark"`
	Reported_run_tvd           *float64   `json:"reported_run_tvd"`
	Reported_run_tvd_ouom      *string    `json:"reported_run_tvd_ouom"`
	Run_depth                  *float64   `json:"run_depth"`
	Run_depth_ouom             *string    `json:"run_depth_ouom"`
	Run_depth_temperature      *float64   `json:"run_depth_temperature"`
	Run_depth_temperature_ouom *string    `json:"run_depth_temperature_ouom"`
	Shutin_period              *float64   `json:"shutin_period"`
	Shutin_period_ouom         *string    `json:"shutin_period_ouom"`
	Survey_date                *time.Time `json:"survey_date"`
	Test_duration              *float64   `json:"test_duration"`
	Test_duration_ouom         *string    `json:"test_duration_ouom"`
	Tubing_size_desc           *string    `json:"tubing_size_desc"`
	Well_head_pressure         *float64   `json:"well_head_pressure"`
	Well_head_pressure_ouom    *string    `json:"well_head_pressure_ouom"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
