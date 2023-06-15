package dto

import (
	"time"
)

type Anl_gas_analysis struct {
	Analysis_id                string     `json:"analysis_id"`
	Anl_source                 string     `json:"anl_source"`
	Gas_anl_obs_no             int        `json:"gas_anl_obs_no"`
	Active_ind                 *string    `json:"active_ind"`
	Effective_date             *time.Time `json:"effective_date"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Fluid_type                 *string    `json:"fluid_type"`
	Gas_gravity                *float64   `json:"gas_gravity"`
	Gas_gravity_ouom           *string    `json:"gas_gravity_ouom"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Preferred_ind              *string    `json:"preferred_ind"`
	Problem_ind                *string    `json:"problem_ind"`
	Pseudo_critical_press      *float64   `json:"pseudo_critical_press"`
	Pseudo_critical_press_ouom *string    `json:"pseudo_critical_press_ouom"`
	Pseudo_critical_temp       *float64   `json:"pseudo_critical_temp"`
	Pseudo_critical_temp_ouom  *string    `json:"pseudo_critical_temp_ouom"`
	Remark                     *string    `json:"remark"`
	Source                     *string    `json:"source"`
	Step_seq_no                *int       `json:"step_seq_no"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
