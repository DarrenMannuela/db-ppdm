package dto

import (
	"time"
)

type Well_treatment struct {
	Uwi                        string     `json:"uwi"`
	Source                     string     `json:"source"`
	Treatment_type             string     `json:"treatment_type"`
	Treatment_obs_no           int        `json:"treatment_obs_no"`
	Active_ind                 *string    `json:"active_ind"`
	Additive_type              *string    `json:"additive_type"`
	Base_depth                 *float64   `json:"base_depth"`
	Base_depth_ouom            *string    `json:"base_depth_ouom"`
	Base_strat_unit_id         *string    `json:"base_strat_unit_id"`
	Completion_obs_no          *int       `json:"completion_obs_no"`
	Completion_source          *string    `json:"completion_source"`
	Effective_date             *time.Time `json:"effective_date"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Form_break_pressure        *float64   `json:"form_break_pressure"`
	Form_break_pressure_ouom   *string    `json:"form_break_pressure_ouom"`
	Injection_rate             *float64   `json:"injection_rate"`
	Injection_rate_ouom        *string    `json:"injection_rate_ouom"`
	Instant_si_pressure        *float64   `json:"instant_si_pressure"`
	Instant_si_pressure_ouom   *string    `json:"instant_si_pressure_ouom"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Proppant_agent_amount      *float64   `json:"proppant_agent_amount"`
	Proppant_agent_amount_ouom *string    `json:"proppant_agent_amount_ouom"`
	Proppant_agent_amount_uom  *string    `json:"proppant_agent_amount_uom"`
	Proppant_agent_mesh_size   *string    `json:"proppant_agent_mesh_size"`
	Proppant_agent_type        *string    `json:"proppant_agent_type"`
	Remark                     *string    `json:"remark"`
	Run_num                    *string    `json:"run_num"`
	Stage_no                   *int       `json:"stage_no"`
	Strat_name_set_id          *string    `json:"strat_name_set_id"`
	Test_num                   *string    `json:"test_num"`
	Test_type                  *string    `json:"test_type"`
	Top_depth                  *float64   `json:"top_depth"`
	Top_depth_ouom             *string    `json:"top_depth_ouom"`
	Top_strat_unit_id          *string    `json:"top_strat_unit_id"`
	Treatment_amount           *float64   `json:"treatment_amount"`
	Treatment_amount_ouom      *string    `json:"treatment_amount_ouom"`
	Treatment_amount_uom       *string    `json:"treatment_amount_uom"`
	Treatment_ba_id            *string    `json:"treatment_ba_id"`
	Treatment_date             *time.Time `json:"treatment_date"`
	Treatment_fluid_type       *string    `json:"treatment_fluid_type"`
	Treatment_pressure         *float64   `json:"treatment_pressure"`
	Treatment_pressure_ouom    *string    `json:"treatment_pressure_ouom"`
	Treatment_status           *string    `json:"treatment_status"`
	Treatment_status_type      *string    `json:"treatment_status_type"`
	Well_test_source           *string    `json:"well_test_source"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
