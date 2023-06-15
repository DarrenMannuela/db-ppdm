package dto

type Anl_valid_problem struct {
	Method_set_id      string  `json:"method_set_id" default:"source"`
	Method_id          string  `json:"method_id" default:"source"`
	Problem_obs_no     int     `json:"problem_obs_no" default:"1"`
	Accuracy_type      *string `json:"accuracy_type" default:""`
	Active_ind         *string `json:"active_ind" default:""`
	Confidence_type    *string `json:"confidence_type" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Problem_desc       *string `json:"problem_desc" default:""`
	Problem_result     *string `json:"problem_result" default:""`
	Problem_severity   *string `json:"problem_severity" default:""`
	Problem_type       *string `json:"problem_type" default:""`
	Remark             *string `json:"remark" default:""`
	Resolution_method  *string `json:"resolution_method" default:""`
	Source             *string `json:"source" default:""`
	Substance_id       *string `json:"substance_id" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
