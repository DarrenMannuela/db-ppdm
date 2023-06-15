package dto

import (
	"time"
)

type Well_test_cushion struct {
	Uwi                        string     `json:"uwi"`
	Source                     string     `json:"source"`
	Test_type                  string     `json:"test_type"`
	Run_num                    string     `json:"run_num"`
	Test_num                   string     `json:"test_num"`
	Cushion_obs_no             int        `json:"cushion_obs_no"`
	Active_ind                 *string    `json:"active_ind"`
	Cushion_gas_pressure       *float64   `json:"cushion_gas_pressure"`
	Cushion_gas_pressure_ouom  *string    `json:"cushion_gas_pressure_ouom"`
	Cushion_inhibitor_volume   *float64   `json:"cushion_inhibitor_volume"`
	Cushion_inhibitor_vol_ouom *string    `json:"cushion_inhibitor_vol_ouom"`
	Cushion_length             *float64   `json:"cushion_length"`
	Cushion_length_ouom        *string    `json:"cushion_length_ouom"`
	Cushion_type               *string    `json:"cushion_type"`
	Cushion_volume             *float64   `json:"cushion_volume"`
	Cushion_volume_ouom        *string    `json:"cushion_volume_ouom"`
	Effective_date             *time.Time `json:"effective_date"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Remark                     *string    `json:"remark"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
