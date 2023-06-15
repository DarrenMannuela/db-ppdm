package dto

import (
	"time"
)

type Well_mud_sample struct {
	Anl_id                    string     `json:"anl_id"`
	Anl_source                string     `json:"anl_source"`
	Step_seq_no               int        `json:"step_seq_no"`
	Mud_anl_obs_no            int        `json:"mud_anl_obs_no"`
	Active_ind                *string    `json:"active_ind"`
	Circulation_stop_date     *time.Time `json:"circulation_stop_date"`
	Circulation_stop_time     *time.Time `json:"circulation_stop_time"`
	Circulation_stop_timezone *string    `json:"circulation_stop_timezone"`
	Effective_date            *time.Time `json:"effective_date"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Fluid_flow                *float64   `json:"fluid_flow"`
	Fluid_flow_ouom           *string    `json:"fluid_flow_ouom"`
	Fluid_loss                *float64   `json:"fluid_loss"`
	Fluid_loss_ouom           *string    `json:"fluid_loss_ouom"`
	Mud_collect_reason        *string    `json:"mud_collect_reason"`
	Mud_density               *float64   `json:"mud_density"`
	Mud_density_ouom          *string    `json:"mud_density_ouom"`
	Mud_density_uom           *string    `json:"mud_density_uom"`
	Mud_ph                    *float64   `json:"mud_ph"`
	Mud_ph_ouom               *string    `json:"mud_ph_ouom"`
	Mud_sample_type           *string    `json:"mud_sample_type"`
	Mud_type                  *string    `json:"mud_type"`
	Mud_viscosity             *float64   `json:"mud_viscosity"`
	Mud_viscosity_ouom        *string    `json:"mud_viscosity_ouom"`
	Period_obs_no             *int       `json:"period_obs_no"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Preferred_ind             *string    `json:"preferred_ind"`
	Problem_ind               *string    `json:"problem_ind"`
	Remark                    *string    `json:"remark"`
	Sample_date               *time.Time `json:"sample_date"`
	Sample_depth              *float64   `json:"sample_depth"`
	Sample_depth_ouom         *string    `json:"sample_depth_ouom"`
	Sample_time               *time.Time `json:"sample_time"`
	Sample_timezone           *string    `json:"sample_timezone"`
	Source                    *string    `json:"source"`
	Source_location           *string    `json:"source_location"`
	Time_since_circ           *float64   `json:"time_since_circ"`
	Time_since_circ_ouom      *string    `json:"time_since_circ_ouom"`
	Uwi                       *string    `json:"uwi"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
