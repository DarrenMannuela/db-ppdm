package dto

import (
	"time"
)

type Well_pressure struct {
	Uwi                         string     `json:"uwi"`
	Source                      string     `json:"source"`
	Pressure_obs_no             int        `json:"pressure_obs_no"`
	Active_ind                  *string    `json:"active_ind"`
	Assigned_field              *string    `json:"assigned_field"`
	Base_depth                  *float64   `json:"base_depth"`
	Base_depth_ouom             *string    `json:"base_depth_ouom"`
	Base_strat_unit_id          *string    `json:"base_strat_unit_id"`
	Effective_date              *time.Time `json:"effective_date"`
	Event_obs_no                *int       `json:"event_obs_no"`
	Event_source                *string    `json:"event_source"`
	Expiry_date                 *time.Time `json:"expiry_date"`
	Flow_casing_pressure        *float64   `json:"flow_casing_pressure"`
	Flow_casing_pressure_ouom   *string    `json:"flow_casing_pressure_ouom"`
	Flow_tubing_pressure        *float64   `json:"flow_tubing_pressure"`
	Flow_tubing_pressure_ouom   *string    `json:"flow_tubing_pressure_ouom"`
	Init_reservoir_pressure     *float64   `json:"init_reservoir_pressure"`
	Init_reservoir_press_ouom   *string    `json:"init_reservoir_press_ouom"`
	Pool_datum                  *string    `json:"pool_datum"`
	Pool_datum_depth            *float64   `json:"pool_datum_depth"`
	Pool_datum_depth_ouom       *string    `json:"pool_datum_depth_ouom"`
	Pool_id                     *string    `json:"pool_id"`
	Ppdm_guid                   string     `json:"ppdm_guid"`
	Prod_string_id              *string    `json:"prod_string_id"`
	Prod_string_source          *string    `json:"prod_string_source"`
	Pr_str_form_obs_no          *int       `json:"pr_str_form_obs_no"`
	Remark                      *string    `json:"remark"`
	Shutin_casing_pressure      *float64   `json:"shutin_casing_pressure"`
	Shutin_casing_pressure_ouom *string    `json:"shutin_casing_pressure_ouom"`
	Shutin_tubing_pressure      *float64   `json:"shutin_tubing_pressure"`
	Shutin_tubing_pressure_ouom *string    `json:"shutin_tubing_pressure_ouom"`
	Strat_name_set_id           *string    `json:"strat_name_set_id"`
	Top_depth                   *float64   `json:"top_depth"`
	Top_depth_ouom              *string    `json:"top_depth_ouom"`
	Top_strat_unit_id           *string    `json:"top_strat_unit_id"`
	Well_datum_depth            *float64   `json:"well_datum_depth"`
	Well_datum_ouom             *string    `json:"well_datum_ouom"`
	Well_test_num               *string    `json:"well_test_num"`
	Well_test_run_num           *string    `json:"well_test_run_num"`
	Well_test_source            *string    `json:"well_test_source"`
	Well_test_type              *string    `json:"well_test_type"`
	Row_changed_by              *string    `json:"row_changed_by"`
	Row_changed_date            *time.Time `json:"row_changed_date"`
	Row_created_by              *string    `json:"row_created_by"`
	Row_created_date            *time.Time `json:"row_created_date"`
	Row_effective_date          *time.Time `json:"row_effective_date"`
	Row_expiry_date             *time.Time `json:"row_expiry_date"`
	Row_quality                 *string    `json:"row_quality"`
}
