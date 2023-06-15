package dto

import (
	"time"
)

type Well_drill_bit_interval struct {
	Uwi                       string     `json:"uwi"`
	Source                    string     `json:"source"`
	Bit_interval_obs_no       int        `json:"bit_interval_obs_no"`
	Active_ind                *string    `json:"active_ind"`
	Avg_force_on_bit          *float64   `json:"avg_force_on_bit"`
	Avg_force_on_bit_ouom     *string    `json:"avg_force_on_bit_ouom"`
	Avg_penetration_rate      *float64   `json:"avg_penetration_rate"`
	Avg_penetration_rate_ouom *string    `json:"avg_penetration_rate_ouom"`
	Bit_company               *string    `json:"bit_company"`
	Bit_drilled_rate          *float64   `json:"bit_drilled_rate"`
	Bit_drilled_rate_ouom     *string    `json:"bit_drilled_rate_ouom"`
	Bit_jet_count             *int       `json:"bit_jet_count"`
	Bit_jet_remark            *string    `json:"bit_jet_remark"`
	Bit_number                *string    `json:"bit_number"`
	Bit_operating_time        *float64   `json:"bit_operating_time"`
	Bit_operating_time_ouom   *string    `json:"bit_operating_time_ouom"`
	Bit_remark                *string    `json:"bit_remark"`
	Bit_size                  *float64   `json:"bit_size"`
	Bit_size_desc             *string    `json:"bit_size_desc"`
	Bit_size_ouom             *string    `json:"bit_size_ouom"`
	Cutting_structure_loc     *string    `json:"cutting_structure_loc"`
	Cutting_structure_mdc     *string    `json:"cutting_structure_mdc"`
	Cutting_structure_ti      *string    `json:"cutting_structure_ti"`
	Cutting_structure_to      *string    `json:"cutting_structure_to"`
	Distance_drilled          *float64   `json:"distance_drilled"`
	Distance_drilled_ouom     *string    `json:"distance_drilled_ouom"`
	Drill_bit_type            *string    `json:"drill_bit_type"`
	Effective_date            *time.Time `json:"effective_date"`
	Equipment_id              *string    `json:"equipment_id"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Flow_rate                 *float64   `json:"flow_rate"`
	Flow_rate_ouom            *string    `json:"flow_rate_ouom"`
	Gage_out_distance         *float64   `json:"gage_out_distance"`
	Gage_out_distance_ouom    *string    `json:"gage_out_distance_ouom"`
	Max_force_on_bit          *float64   `json:"max_force_on_bit"`
	Max_force_on_bit_ouom     *string    `json:"max_force_on_bit_ouom"`
	Max_penetration_rate      *float64   `json:"max_penetration_rate"`
	Max_penetration_rate_ouom *string    `json:"max_penetration_rate_ouom"`
	Min_force_on_bit          *float64   `json:"min_force_on_bit"`
	Min_force_on_bit_ouom     *string    `json:"min_force_on_bit_ouom"`
	Min_penetration_rate      *float64   `json:"min_penetration_rate"`
	Min_penetration_rate_ouom *string    `json:"min_penetration_rate_ouom"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Reason_pulled             *string    `json:"reason_pulled"`
	Remark                    *string    `json:"remark"`
	Reported_tfa              *float64   `json:"reported_tfa"`
	Reported_tfa_ouom         *string    `json:"reported_tfa_ouom"`
	Run_in_depth              *float64   `json:"run_in_depth"`
	Run_in_depth_ouom         *string    `json:"run_in_depth_ouom"`
	Run_out_depth             *float64   `json:"run_out_depth"`
	Run_out_depth_ouom        *string    `json:"run_out_depth_ouom"`
	Torque                    *float64   `json:"torque"`
	Torque_ouom               *string    `json:"torque_ouom"`
	Total_bit_revolutions     *int       `json:"total_bit_revolutions"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
