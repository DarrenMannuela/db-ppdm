package dto

import (
	"time"
)

type Well_pressure_aof struct {
	Uwi                         string     `json:"uwi"`
	Source                      string     `json:"source"`
	Pressure_obs_no             int        `json:"pressure_obs_no"`
	Aof_obs_no                  int        `json:"aof_obs_no"`
	Active_ind                  *string    `json:"active_ind"`
	Analysis_type               *string    `json:"analysis_type"`
	Aof_calculate_method        *string    `json:"aof_calculate_method"`
	Aof_potential               *float64   `json:"aof_potential"`
	Aof_potential_ouom          *string    `json:"aof_potential_ouom"`
	Aof_slope                   *float64   `json:"aof_slope"`
	Bottom_hole_pressure_method *string    `json:"bottom_hole_pressure_method"`
	Caof_rate                   *float64   `json:"caof_rate"`
	Caof_rate_ouom              *string    `json:"caof_rate_ouom"`
	Choke_size_desc             *string    `json:"choke_size_desc"`
	Effective_date              *time.Time `json:"effective_date"`
	Expiry_date                 *time.Time `json:"expiry_date"`
	Flow_period_duration        *float64   `json:"flow_period_duration"`
	Flow_period_duration_ouom   *string    `json:"flow_period_duration_ouom"`
	Flow_pressure               *float64   `json:"flow_pressure"`
	Flow_pressure_ouom          *string    `json:"flow_pressure_ouom"`
	Flow_rate                   *float64   `json:"flow_rate"`
	Flow_rate_ouom              *string    `json:"flow_rate_ouom"`
	Four_point_caof_rate        *float64   `json:"four_point_caof_rate"`
	Four_point_caof_rate_ouom   *string    `json:"four_point_caof_rate_ouom"`
	Ppdm_guid                   string     `json:"ppdm_guid"`
	Remark                      *string    `json:"remark"`
	Reservoir_pressure          *float64   `json:"reservoir_pressure"`
	Reservoir_pressure_ouom     *string    `json:"reservoir_pressure_ouom"`
	Shutin_pressure_type        *string    `json:"shutin_pressure_type"`
	Test_date                   *time.Time `json:"test_date"`
	Row_changed_by              *string    `json:"row_changed_by"`
	Row_changed_date            *time.Time `json:"row_changed_date"`
	Row_created_by              *string    `json:"row_created_by"`
	Row_created_date            *time.Time `json:"row_created_date"`
	Row_effective_date          *time.Time `json:"row_effective_date"`
	Row_expiry_date             *time.Time `json:"row_expiry_date"`
	Row_quality                 *string    `json:"row_quality"`
}
