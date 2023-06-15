package dto

import (
	"time"
)

type Well_drill_bit_period struct {
	Uwi                       string     `json:"uwi"`
	Period_obs_no             int        `json:"period_obs_no"`
	Bit_source                string     `json:"bit_source"`
	Bit_interval_obs_no       int        `json:"bit_interval_obs_no"`
	Drill_bit_period_obs_no   int        `json:"drill_bit_period_obs_no"`
	Active_ind                *string    `json:"active_ind"`
	Avg_force_on_bit          *float64   `json:"avg_force_on_bit"`
	Avg_force_on_bit_ouom     *string    `json:"avg_force_on_bit_ouom"`
	Avg_penetration_rate      *float64   `json:"avg_penetration_rate"`
	Avg_penetration_rate_ouom *string    `json:"avg_penetration_rate_ouom"`
	Avg_rotary_rpm            *float64   `json:"avg_rotary_rpm"`
	Base_depth                *float64   `json:"base_depth"`
	Base_depth_ouom           *string    `json:"base_depth_ouom"`
	Bit_operating_time        *float64   `json:"bit_operating_time"`
	Bit_operating_time_ouom   *string    `json:"bit_operating_time_ouom"`
	Drill_op_type             *string    `json:"drill_op_type"`
	Effective_date            *time.Time `json:"effective_date"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Flow_rate                 *float64   `json:"flow_rate"`
	Flow_rate_ouom            *string    `json:"flow_rate_ouom"`
	Max_force_on_bit          *float64   `json:"max_force_on_bit"`
	Max_force_on_bit_ouom     *string    `json:"max_force_on_bit_ouom"`
	Max_penetration_rate      *float64   `json:"max_penetration_rate"`
	Max_penetration_rate_ouom *string    `json:"max_penetration_rate_ouom"`
	Max_rotary_rpm            *float64   `json:"max_rotary_rpm"`
	Min_force_on_bit          *float64   `json:"min_force_on_bit"`
	Min_force_on_bit_ouom     *string    `json:"min_force_on_bit_ouom"`
	Min_penetration_rate      *float64   `json:"min_penetration_rate"`
	Min_penetration_rate_ouom *string    `json:"min_penetration_rate_ouom"`
	Min_rotary_rpm            *float64   `json:"min_rotary_rpm"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Remark                    *string    `json:"remark"`
	Source                    *string    `json:"source"`
	Top_depth                 *float64   `json:"top_depth"`
	Top_depth_ouom            *string    `json:"top_depth_ouom"`
	Torque                    *float64   `json:"torque"`
	Torque_ouom               *string    `json:"torque_ouom"`
	Total_bit_revolution      *int       `json:"total_bit_revolution"`
	Total_drilled_dist        *float64   `json:"total_drilled_dist"`
	Total_drilled_dist_ouom   *string    `json:"total_drilled_dist_ouom"`
	Total_period_run          *float64   `json:"total_period_run"`
	Total_period_run_ouom     *string    `json:"total_period_run_ouom"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
