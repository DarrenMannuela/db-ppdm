package dto

import (
	"time"
)

type Anl_maceral struct {
	Analysis_id                 string     `json:"analysis_id"`
	Anl_source                  string     `json:"anl_source"`
	Maceral_anl_obs_no          int        `json:"maceral_anl_obs_no"`
	Active_ind                  *string    `json:"active_ind"`
	Effective_date              *time.Time `json:"effective_date"`
	Expiry_date                 *time.Time `json:"expiry_date"`
	Lithology_desc              *string    `json:"lithology_desc"`
	Maceral_amount_type         *string    `json:"maceral_amount_type"`
	Ppdm_guid                   string     `json:"ppdm_guid"`
	Preferred_ind               *string    `json:"preferred_ind"`
	Problem_ind                 *string    `json:"problem_ind"`
	Remark                      *string    `json:"remark"`
	Sample_tot_maceral_val      *float64   `json:"sample_tot_maceral_val"`
	Sample_tot_maceral_val_ouom *string    `json:"sample_tot_maceral_val_ouom"`
	Source                      *string    `json:"source"`
	Step_seq_no                 *int       `json:"step_seq_no"`
	Substance_id                *string    `json:"substance_id"`
	Total_maceral_val           *float64   `json:"total_maceral_val"`
	Total_maceral_val_ouom      *string    `json:"total_maceral_val_ouom"`
	Row_changed_by              *string    `json:"row_changed_by"`
	Row_changed_date            *time.Time `json:"row_changed_date"`
	Row_created_by              *string    `json:"row_created_by"`
	Row_created_date            *time.Time `json:"row_created_date"`
	Row_effective_date          *time.Time `json:"row_effective_date"`
	Row_expiry_date             *time.Time `json:"row_expiry_date"`
	Row_quality                 *string    `json:"row_quality"`
}
