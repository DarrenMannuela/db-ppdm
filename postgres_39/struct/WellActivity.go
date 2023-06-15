package dto

import (
	"time"
)

type Well_activity struct {
	Uwi                        string     `json:"uwi"`
	Source                     string     `json:"source"`
	Activity_obs_no            int        `json:"activity_obs_no"`
	Active_ind                 *string    `json:"active_ind"`
	Activity_duration          *float64   `json:"activity_duration"`
	Activity_duration_ouom     *string    `json:"activity_duration_ouom"`
	Activity_product           *string    `json:"activity_product"`
	Activity_set_id            *string    `json:"activity_set_id"`
	Activity_type_id           *string    `json:"activity_type_id"`
	Base_depth                 *float64   `json:"base_depth"`
	Base_depth_ouom            *string    `json:"base_depth_ouom"`
	Base_strat_unit_id         *string    `json:"base_strat_unit_id"`
	Blowout_fluid              *string    `json:"blowout_fluid"`
	Control_date               *time.Time `json:"control_date"`
	Control_depth              *float64   `json:"control_depth"`
	Control_depth_ouom         *string    `json:"control_depth_ouom"`
	Downtime_type              *string    `json:"downtime_type"`
	Effective_date             *time.Time `json:"effective_date"`
	End_date                   *time.Time `json:"end_date"`
	End_time                   *time.Time `json:"end_time"`
	End_timezone               *string    `json:"end_timezone"`
	Event_date                 *time.Time `json:"event_date"`
	Event_depth                *float64   `json:"event_depth"`
	Event_depth_ouom           *string    `json:"event_depth_ouom"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Final_mud_density          *float64   `json:"final_mud_density"`
	Final_mud_density_ouom     *string    `json:"final_mud_density_ouom"`
	Final_mud_viscosity        *float64   `json:"final_mud_viscosity"`
	Final_mud_viscosity_ouom   *string    `json:"final_mud_viscosity_ouom"`
	Lithology                  *string    `json:"lithology"`
	Lost_circ_severity         *string    `json:"lost_circ_severity"`
	Lost_material_amount       *float64   `json:"lost_material_amount"`
	Lost_material_amount_ouom  *string    `json:"lost_material_amount_ouom"`
	Lost_material_type         *string    `json:"lost_material_type"`
	Lost_volume                *float64   `json:"lost_volume"`
	Lost_volume_ouom           *string    `json:"lost_volume_ouom"`
	Lost_volume_product        *string    `json:"lost_volume_product"`
	Lost_volume_uom            *string    `json:"lost_volume_uom"`
	Period_obs_no              *int       `json:"period_obs_no"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Prod_string_id             *string    `json:"prod_string_id"`
	Prod_string_source         *string    `json:"prod_string_source"`
	Pr_str_form_obs_no         *int       `json:"pr_str_form_obs_no"`
	Remark                     *string    `json:"remark"`
	Reported_code              *string    `json:"reported_code"`
	Reported_time_elapsed      *float64   `json:"reported_time_elapsed"`
	Reported_time_elapsed_ouom *string    `json:"reported_time_elapsed_ouom"`
	Reported_tvd               *float64   `json:"reported_tvd"`
	Reported_tvd_ouom          *string    `json:"reported_tvd_ouom"`
	Start_date                 *time.Time `json:"start_date"`
	Start_mud_density          *float64   `json:"start_mud_density"`
	Start_mud_density_ouom     *string    `json:"start_mud_density_ouom"`
	Start_mud_viscosity        *float64   `json:"start_mud_viscosity"`
	Start_mud_viscosity_ouom   *string    `json:"start_mud_viscosity_ouom"`
	Start_time                 *time.Time `json:"start_time"`
	Start_timezone             *string    `json:"start_timezone"`
	Strat_name_set_id          *string    `json:"strat_name_set_id"`
	Top_depth                  *float64   `json:"top_depth"`
	Top_depth_ouom             *string    `json:"top_depth_ouom"`
	Top_strat_unit_id          *string    `json:"top_strat_unit_id"`
	Total_time_elapsed         *float64   `json:"total_time_elapsed"`
	Total_time_elapsed_ouom    *string    `json:"total_time_elapsed_ouom"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
