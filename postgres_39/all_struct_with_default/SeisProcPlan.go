package dto

type Seis_proc_plan struct {
	Seis_proc_plan_id  string  `json:"seis_proc_plan_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Created_by         *string `json:"created_by" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Plan_name          *string `json:"plan_name" default:""`
	Plan_num           *string `json:"plan_num" default:""`
	Plan_owner         *string `json:"plan_owner" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
