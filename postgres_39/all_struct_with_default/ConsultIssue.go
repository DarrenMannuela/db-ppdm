package dto

type Consult_issue struct {
	Consult_id         string  `json:"consult_id" default:"source"`
	Detail_id          string  `json:"detail_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Detail_desc        *string `json:"detail_desc" default:""`
	Discussion_id      *string `json:"discussion_id" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Issue_type         *string `json:"issue_type" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Remark_type        *string `json:"remark_type" default:""`
	Resolution         *string `json:"resolution" default:""`
	Source             *string `json:"source" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
