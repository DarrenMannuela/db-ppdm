package dto

import (
	"time"
)

type Anl_gas_heat struct {
	Analysis_id             string     `json:"analysis_id"`
	Anl_source              string     `json:"anl_source"`
	Heat_content_obs_no     int        `json:"heat_content_obs_no"`
	Active_ind              *string    `json:"active_ind"`
	Calculated_ind          *string    `json:"calculated_ind"`
	Calculate_method_id     *string    `json:"calculate_method_id"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Heat_content_method     *string    `json:"heat_content_method"`
	Heat_content_press      *float64   `json:"heat_content_press"`
	Heat_content_press_ouom *string    `json:"heat_content_press_ouom"`
	Heat_content_temp       *float64   `json:"heat_content_temp"`
	Heat_content_temp_ouom  *string    `json:"heat_content_temp_ouom"`
	Heat_content_value      *float64   `json:"heat_content_value"`
	Heat_content_value_ouom *string    `json:"heat_content_value_ouom"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Problem_ind             *string    `json:"problem_ind"`
	Remark                  *string    `json:"remark"`
	Reported_ind            *string    `json:"reported_ind"`
	Source                  *string    `json:"source"`
	Step_seq_no             *int       `json:"step_seq_no"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
