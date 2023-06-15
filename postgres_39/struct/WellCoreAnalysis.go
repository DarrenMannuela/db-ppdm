package dto

import (
	"time"
)

type Well_core_analysis struct {
	Uwi                   string     `json:"uwi"`
	Source                string     `json:"source"`
	Core_id               string     `json:"core_id"`
	Analysis_obs_no       int        `json:"analysis_obs_no"`
	Active_ind            *string    `json:"active_ind"`
	Analysis_date         *time.Time `json:"analysis_date"`
	Analyzing_company     *string    `json:"analyzing_company"`
	Analyzing_company_loc *string    `json:"analyzing_company_loc"`
	Analyzing_file_num    *string    `json:"analyzing_file_num"`
	Core_analyst_name     *string    `json:"core_analyst_name"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Primary_sample_type   *string    `json:"primary_sample_type"`
	Remark                *string    `json:"remark"`
	Sample_diameter       *float64   `json:"sample_diameter"`
	Sample_diameter_ouom  *string    `json:"sample_diameter_ouom"`
	Sample_length         *float64   `json:"sample_length"`
	Sample_length_ouom    *string    `json:"sample_length_ouom"`
	Sample_shape          *string    `json:"sample_shape"`
	Second_sample_type    *string    `json:"second_sample_type"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
