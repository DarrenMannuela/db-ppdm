package dto

import (
	"time"
)

type Anl_analysis_step struct {
	Analysis_id           string     `json:"analysis_id"`
	Anl_source            string     `json:"anl_source"`
	Step_seq_no           int        `json:"step_seq_no"`
	Active_ind            *string    `json:"active_ind"`
	Analysis_phase        *string    `json:"analysis_phase"`
	Anl_date              *time.Time `json:"anl_date"`
	Complete_date         *time.Time `json:"complete_date"`
	Effective_date        *time.Time `json:"effective_date"`
	End_date              *time.Time `json:"end_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Final_volume          *float64   `json:"final_volume"`
	Final_volume_ouom     *string    `json:"final_volume_ouom"`
	Final_volume_percent  *float64   `json:"final_volume_percent"`
	Final_weight          *float64   `json:"final_weight"`
	Final_weight_ouom     *string    `json:"final_weight_ouom"`
	Method_id             *string    `json:"method_id"`
	Method_set_id         *string    `json:"method_set_id"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Problem_ind           *string    `json:"problem_ind"`
	Received_date         *time.Time `json:"received_date"`
	Recovered_percent     *float64   `json:"recovered_percent"`
	Remark                *string    `json:"remark"`
	Reported_date         *time.Time `json:"reported_date"`
	Results_received_date *time.Time `json:"results_received_date"`
	Results_received_ind  *string    `json:"results_received_ind"`
	Sample_fraction_type  *string    `json:"sample_fraction_type"`
	Sample_quality        *string    `json:"sample_quality"`
	Sample_quality_desc   *string    `json:"sample_quality_desc"`
	Source                *string    `json:"source"`
	Start_date            *time.Time `json:"start_date"`
	Step_completed_ind    *string    `json:"step_completed_ind"`
	Step_desc             *string    `json:"step_desc"`
	Step_quality_desc     *string    `json:"step_quality_desc"`
	Step_quality_type     *string    `json:"step_quality_type"`
	Step_requested_ind    *string    `json:"step_requested_ind"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
