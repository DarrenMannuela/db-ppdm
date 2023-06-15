package dto

type Well_mud_sample struct {
	Anl_id                    string   `json:"anl_id" default:"source"`
	Anl_source                string   `json:"anl_source" default:"source"`
	Step_seq_no               int      `json:"step_seq_no" default:"1"`
	Mud_anl_obs_no            int      `json:"mud_anl_obs_no" default:"1"`
	Active_ind                *string  `json:"active_ind" default:""`
	Circulation_stop_date     *string  `json:"circulation_stop_date" default:""`
	Circulation_stop_time     *string  `json:"circulation_stop_time" default:""`
	Circulation_stop_timezone *string  `json:"circulation_stop_timezone" default:""`
	Effective_date            *string  `json:"effective_date" default:""`
	Expiry_date               *string  `json:"expiry_date" default:""`
	Fluid_flow                *float64 `json:"fluid_flow" default:""`
	Fluid_flow_ouom           *string  `json:"fluid_flow_ouom" default:""`
	Fluid_loss                *float64 `json:"fluid_loss" default:""`
	Fluid_loss_ouom           *string  `json:"fluid_loss_ouom" default:""`
	Mud_collect_reason        *string  `json:"mud_collect_reason" default:""`
	Mud_density               *float64 `json:"mud_density" default:""`
	Mud_density_ouom          *string  `json:"mud_density_ouom" default:""`
	Mud_density_uom           *string  `json:"mud_density_uom" default:""`
	Mud_ph                    *float64 `json:"mud_ph" default:""`
	Mud_ph_ouom               *string  `json:"mud_ph_ouom" default:""`
	Mud_sample_type           *string  `json:"mud_sample_type" default:""`
	Mud_type                  *string  `json:"mud_type" default:""`
	Mud_viscosity             *float64 `json:"mud_viscosity" default:""`
	Mud_viscosity_ouom        *string  `json:"mud_viscosity_ouom" default:""`
	Period_obs_no             *int     `json:"period_obs_no" default:""`
	Ppdm_guid                 *string  `json:"ppdm_guid" default:""`
	Preferred_ind             *string  `json:"preferred_ind" default:""`
	Problem_ind               *string  `json:"problem_ind" default:""`
	Remark                    *string  `json:"remark" default:""`
	Sample_date               *string  `json:"sample_date" default:""`
	Sample_depth              *float64 `json:"sample_depth" default:""`
	Sample_depth_ouom         *string  `json:"sample_depth_ouom" default:""`
	Sample_time               *string  `json:"sample_time" default:""`
	Sample_timezone           *string  `json:"sample_timezone" default:""`
	Source                    *string  `json:"source" default:""`
	Source_location           *string  `json:"source_location" default:""`
	Time_since_circ           *float64 `json:"time_since_circ" default:""`
	Time_since_circ_ouom      *string  `json:"time_since_circ_ouom" default:""`
	Uwi                       *string  `json:"uwi" default:""`
	Row_changed_by            *string  `json:"row_changed_by" default:""`
	Row_changed_date          *string  `json:"row_changed_date" default:""`
	Row_created_by            *string  `json:"row_created_by" default:""`
	Row_created_date          *string  `json:"row_created_date" default:""`
	Row_effective_date        *string  `json:"row_effective_date" default:""`
	Row_expiry_date           *string  `json:"row_expiry_date" default:""`
	Row_quality               *string  `json:"row_quality" default:""`
}
