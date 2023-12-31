package dto

type Well_drill_bit_interval struct {
	Uwi                       string   `json:"uwi" default:"source"`
	Source                    string   `json:"source" default:"source"`
	Bit_interval_obs_no       int      `json:"bit_interval_obs_no" default:"1"`
	Active_ind                *string  `json:"active_ind" default:""`
	Avg_force_on_bit          *float64 `json:"avg_force_on_bit" default:""`
	Avg_force_on_bit_ouom     *string  `json:"avg_force_on_bit_ouom" default:""`
	Avg_penetration_rate      *float64 `json:"avg_penetration_rate" default:""`
	Avg_penetration_rate_ouom *string  `json:"avg_penetration_rate_ouom" default:""`
	Bit_company               *string  `json:"bit_company" default:""`
	Bit_drilled_rate          *float64 `json:"bit_drilled_rate" default:""`
	Bit_drilled_rate_ouom     *string  `json:"bit_drilled_rate_ouom" default:""`
	Bit_jet_count             *int     `json:"bit_jet_count" default:""`
	Bit_jet_remark            *string  `json:"bit_jet_remark" default:""`
	Bit_number                *string  `json:"bit_number" default:""`
	Bit_operating_time        *float64 `json:"bit_operating_time" default:""`
	Bit_operating_time_ouom   *string  `json:"bit_operating_time_ouom" default:""`
	Bit_remark                *string  `json:"bit_remark" default:""`
	Bit_size                  *float64 `json:"bit_size" default:""`
	Bit_size_desc             *string  `json:"bit_size_desc" default:""`
	Bit_size_ouom             *string  `json:"bit_size_ouom" default:""`
	Cutting_structure_loc     *string  `json:"cutting_structure_loc" default:""`
	Cutting_structure_mdc     *string  `json:"cutting_structure_mdc" default:""`
	Cutting_structure_ti      *string  `json:"cutting_structure_ti" default:""`
	Cutting_structure_to      *string  `json:"cutting_structure_to" default:""`
	Distance_drilled          *float64 `json:"distance_drilled" default:""`
	Distance_drilled_ouom     *string  `json:"distance_drilled_ouom" default:""`
	Drill_bit_type            *string  `json:"drill_bit_type" default:""`
	Effective_date            *string  `json:"effective_date" default:""`
	Equipment_id              *string  `json:"equipment_id" default:""`
	Expiry_date               *string  `json:"expiry_date" default:""`
	Flow_rate                 *float64 `json:"flow_rate" default:""`
	Flow_rate_ouom            *string  `json:"flow_rate_ouom" default:""`
	Gage_out_distance         *float64 `json:"gage_out_distance" default:""`
	Gage_out_distance_ouom    *string  `json:"gage_out_distance_ouom" default:""`
	Max_force_on_bit          *float64 `json:"max_force_on_bit" default:""`
	Max_force_on_bit_ouom     *string  `json:"max_force_on_bit_ouom" default:""`
	Max_penetration_rate      *float64 `json:"max_penetration_rate" default:""`
	Max_penetration_rate_ouom *string  `json:"max_penetration_rate_ouom" default:""`
	Min_force_on_bit          *float64 `json:"min_force_on_bit" default:""`
	Min_force_on_bit_ouom     *string  `json:"min_force_on_bit_ouom" default:""`
	Min_penetration_rate      *float64 `json:"min_penetration_rate" default:""`
	Min_penetration_rate_ouom *string  `json:"min_penetration_rate_ouom" default:""`
	Ppdm_guid                 *string  `json:"ppdm_guid" default:""`
	Reason_pulled             *string  `json:"reason_pulled" default:""`
	Remark                    *string  `json:"remark" default:""`
	Reported_tfa              *float64 `json:"reported_tfa" default:""`
	Reported_tfa_ouom         *string  `json:"reported_tfa_ouom" default:""`
	Run_in_depth              *float64 `json:"run_in_depth" default:""`
	Run_in_depth_ouom         *string  `json:"run_in_depth_ouom" default:""`
	Run_out_depth             *float64 `json:"run_out_depth" default:""`
	Run_out_depth_ouom        *string  `json:"run_out_depth_ouom" default:""`
	Torque                    *float64 `json:"torque" default:""`
	Torque_ouom               *string  `json:"torque_ouom" default:""`
	Total_bit_revolutions     *int     `json:"total_bit_revolutions" default:""`
	Row_changed_by            *string  `json:"row_changed_by" default:""`
	Row_changed_date          *string  `json:"row_changed_date" default:""`
	Row_created_by            *string  `json:"row_created_by" default:""`
	Row_created_date          *string  `json:"row_created_date" default:""`
	Row_effective_date        *string  `json:"row_effective_date" default:""`
	Row_expiry_date           *string  `json:"row_expiry_date" default:""`
	Row_quality               *string  `json:"row_quality" default:""`
}
