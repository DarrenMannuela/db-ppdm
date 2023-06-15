package dto

import (
	"time"
)

type Anl_water_salinity struct {
	Analysis_id           string     `json:"analysis_id"`
	Anl_source            string     `json:"anl_source"`
	Water_salinity_obs_no int        `json:"water_salinity_obs_no"`
	Active_ind            *string    `json:"active_ind"`
	Calculated_ind        *string    `json:"calculated_ind"`
	Calculate_method_id   *string    `json:"calculate_method_id"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Ignition_ind          *string    `json:"ignition_ind"`
	Measured_ind          *string    `json:"measured_ind"`
	Measured_temp         *float64   `json:"measured_temp"`
	Measured_temp_ouom    *string    `json:"measured_temp_ouom"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Preferred_ind         *string    `json:"preferred_ind"`
	Problem_ind           *string    `json:"problem_ind"`
	Remark                *string    `json:"remark"`
	Reported_ind          *string    `json:"reported_ind"`
	Salinity              *float64   `json:"salinity"`
	Salinity_ouom         *string    `json:"salinity_ouom"`
	Salinity_type         *string    `json:"salinity_type"`
	Salinity_uom          *string    `json:"salinity_uom"`
	Source                *string    `json:"source"`
	Step_seq_no           *int       `json:"step_seq_no"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
