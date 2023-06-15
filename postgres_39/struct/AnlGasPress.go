package dto

import (
	"time"
)

type Anl_gas_press struct {
	Analysis_id          string     `json:"analysis_id"`
	Anl_source           string     `json:"anl_source"`
	Gas_anl_press_obs_no int        `json:"gas_anl_press_obs_no"`
	Active_ind           *string    `json:"active_ind"`
	Calculate_method_id  *string    `json:"calculate_method_id"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Gauge_press          *float64   `json:"gauge_press"`
	Gauge_press_ouom     *string    `json:"gauge_press_ouom"`
	Gauge_temp           *float64   `json:"gauge_temp"`
	Gauge_temp_ouom      *string    `json:"gauge_temp_ouom"`
	Measurement_location *string    `json:"measurement_location"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Problem_ind          *string    `json:"problem_ind"`
	Remark               *string    `json:"remark"`
	Source               *string    `json:"source"`
	Step_seq_no          *int       `json:"step_seq_no"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
