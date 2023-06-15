package dto

import (
	"time"
)

type Anl_accuracy struct {
	Analysis_id            string     `json:"analysis_id"`
	Anl_source             string     `json:"anl_source"`
	Accuracy_obs_no        int        `json:"accuracy_obs_no"`
	Accuracy_desc          *string    `json:"accuracy_desc"`
	Accuracy_type          *string    `json:"accuracy_type"`
	Active_ind             *string    `json:"active_ind"`
	Calculated_ind         *string    `json:"calculated_ind"`
	Calculate_method_id    *string    `json:"calculate_method_id"`
	Confidence_type        *string    `json:"confidence_type"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Measured_value_ind     *string    `json:"measured_value_ind"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Raw_value_ind          *string    `json:"raw_value_ind"`
	Referenced_column_name *string    `json:"referenced_column_name"`
	Referenced_ppdm_guid   *string    `json:"referenced_ppdm_guid"`
	Referenced_system_id   *string    `json:"referenced_system_id"`
	Referenced_table_name  *string    `json:"referenced_table_name"`
	Remark                 *string    `json:"remark"`
	Repeatability_type     *string    `json:"repeatability_type"`
	Reported_ind           *string    `json:"reported_ind"`
	Source                 *string    `json:"source"`
	Step_seq_no            *int       `json:"step_seq_no"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
