package dto

import (
	"time"
)

type Paleo_summary struct {
	Paleo_summary_id           string     `json:"paleo_summary_id"`
	Access_condition           *string    `json:"access_condition"`
	Active_ind                 *string    `json:"active_ind"`
	Analysis_end_date          *time.Time `json:"analysis_end_date"`
	Analysis_start_date        *time.Time `json:"analysis_start_date"`
	Confidential_ind           *string    `json:"confidential_ind"`
	Confidential_reason        *string    `json:"confidential_reason"`
	Confidential_release_date  *time.Time `json:"confidential_release_date"`
	Confidential_term          *string    `json:"confidential_term"`
	Confidential_type          *string    `json:"confidential_type"`
	Diversity_count            *int       `json:"diversity_count"`
	Effective_date             *time.Time `json:"effective_date"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Preservation_quality       *string    `json:"preservation_quality"`
	Reference_num              *string    `json:"reference_num"`
	Remark                     *string    `json:"remark"`
	Report_date                *time.Time `json:"report_date"`
	Report_title               *string    `json:"report_title"`
	Source                     *string    `json:"source"`
	Total_sample_interval      *float64   `json:"total_sample_interval"`
	Total_sample_interval_ouom *string    `json:"total_sample_interval_ouom"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
