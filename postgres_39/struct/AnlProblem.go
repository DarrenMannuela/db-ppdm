package dto

import (
	"time"
)

type Anl_problem struct {
	Analysis_id            string     `json:"analysis_id"`
	Anl_source             string     `json:"anl_source"`
	Problem_obs_no         int        `json:"problem_obs_no"`
	Active_ind             *string    `json:"active_ind"`
	Activity_set_id        *string    `json:"activity_set_id"`
	Anl_problem_ind        *string    `json:"anl_problem_ind"`
	Ba_problem_ind         *string    `json:"ba_problem_ind"`
	Confidence_type        *string    `json:"confidence_type"`
	Effective_date         *time.Time `json:"effective_date"`
	Equip_problem_ind      *string    `json:"equip_problem_ind"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Method_id              *string    `json:"method_id"`
	Method_problem_ind     *string    `json:"method_problem_ind"`
	Percent_of_sample      *float64   `json:"percent_of_sample"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Problem_desc           *string    `json:"problem_desc"`
	Problem_result         *string    `json:"problem_result"`
	Problem_result_desc    *string    `json:"problem_result_desc"`
	Problem_severity       *string    `json:"problem_severity"`
	Problem_severity_desc  *string    `json:"problem_severity_desc"`
	Problem_type           *string    `json:"problem_type"`
	Referenced_column_name *string    `json:"referenced_column_name"`
	Referenced_ppdm_guid   *string    `json:"referenced_ppdm_guid"`
	Referenced_system_id   *string    `json:"referenced_system_id"`
	Referenced_table_name  *string    `json:"referenced_table_name"`
	Remark                 *string    `json:"remark"`
	Resolution_method      *string    `json:"resolution_method"`
	Resolution_method_desc *string    `json:"resolution_method_desc"`
	Resolution_step_seq_no *int       `json:"resolution_step_seq_no"`
	Resolved_by_ba_id      *string    `json:"resolved_by_ba_id"`
	Sample_problem_ind     *string    `json:"sample_problem_ind"`
	Source                 *string    `json:"source"`
	Step_seq_no            *int       `json:"step_seq_no"`
	Valid_problem_obs_no   *int       `json:"valid_problem_obs_no"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
