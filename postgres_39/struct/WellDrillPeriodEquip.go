package dto

import (
	"time"
)

type Well_drill_period_equip struct {
	Uwi                         string     `json:"uwi"`
	Period_obs_no               int        `json:"period_obs_no"`
	Equipment_id                string     `json:"equipment_id"`
	Active_ind                  *string    `json:"active_ind"`
	Avg_boiler_ph               *float64   `json:"avg_boiler_ph"`
	Avg_boiler_stack_temp       *float64   `json:"avg_boiler_stack_temp"`
	Avg_boiler_stack_temp_ouom  *string    `json:"avg_boiler_stack_temp_ouom"`
	Avg_pump_pressure           *float64   `json:"avg_pump_pressure"`
	Avg_pump_pressure_ouom      *string    `json:"avg_pump_pressure_ouom"`
	Booked_time_elapsed         *float64   `json:"booked_time_elapsed"`
	Booked_time_elapsed_ouom    *string    `json:"booked_time_elapsed_ouom"`
	Effective_date              *time.Time `json:"effective_date"`
	Equipment_obs_no            *int       `json:"equipment_obs_no"`
	Equip_ref_num               *string    `json:"equip_ref_num"`
	Expiry_date                 *time.Time `json:"expiry_date"`
	Intake_density              *float64   `json:"intake_density"`
	Intake_density_ouom         *string    `json:"intake_density_ouom"`
	Overflow_density            *float64   `json:"overflow_density"`
	Overflow_density_ouom       *string    `json:"overflow_density_ouom"`
	Ppdm_guid                   string     `json:"ppdm_guid"`
	Pump_liner_inside_diam      *float64   `json:"pump_liner_inside_diam"`
	Pump_liner_inside_diam_ouom *string    `json:"pump_liner_inside_diam_ouom"`
	Pump_stroke                 *int       `json:"pump_stroke"`
	Pump_stroke_ouom            *string    `json:"pump_stroke_ouom"`
	Reduced_pump_depth          *float64   `json:"reduced_pump_depth"`
	Reduced_pump_depth_ouom     *string    `json:"reduced_pump_depth_ouom"`
	Reduced_pump_press          *float64   `json:"reduced_pump_press"`
	Reduced_pump_press_ouom     *string    `json:"reduced_pump_press_ouom"`
	Reduced_pump_stroke         *float64   `json:"reduced_pump_stroke"`
	Reduced_pump_stroke_ouom    *string    `json:"reduced_pump_stroke_ouom"`
	Remark                      *string    `json:"remark"`
	Source                      *string    `json:"source"`
	Tubing_obs_no               *int       `json:"tubing_obs_no"`
	Tubing_source               *string    `json:"tubing_source"`
	Tubing_type                 *string    `json:"tubing_type"`
	Underflow_density           *float64   `json:"underflow_density"`
	Underflow_density_ouom      *string    `json:"underflow_density_ouom"`
	Row_changed_by              *string    `json:"row_changed_by"`
	Row_changed_date            *time.Time `json:"row_changed_date"`
	Row_created_by              *string    `json:"row_created_by"`
	Row_created_date            *time.Time `json:"row_created_date"`
	Row_effective_date          *time.Time `json:"row_effective_date"`
	Row_expiry_date             *time.Time `json:"row_expiry_date"`
	Row_quality                 *string    `json:"row_quality"`
}
