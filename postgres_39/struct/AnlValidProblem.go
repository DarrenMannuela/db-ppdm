package dto

import (
	"time"
)

type Anl_valid_problem struct {
	Method_set_id      string     `json:"method_set_id"`
	Method_id          string     `json:"method_id"`
	Problem_obs_no     int        `json:"problem_obs_no"`
	Accuracy_type      *string    `json:"accuracy_type"`
	Active_ind         *string    `json:"active_ind"`
	Confidence_type    *string    `json:"confidence_type"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Problem_desc       *string    `json:"problem_desc"`
	Problem_result     *string    `json:"problem_result"`
	Problem_severity   *string    `json:"problem_severity"`
	Problem_type       *string    `json:"problem_type"`
	Remark             *string    `json:"remark"`
	Resolution_method  *string    `json:"resolution_method"`
	Source             *string    `json:"source"`
	Substance_id       *string    `json:"substance_id"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
