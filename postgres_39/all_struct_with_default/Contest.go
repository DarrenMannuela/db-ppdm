package dto

type Contest struct {
	Contest_id         string  `json:"contest_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	End_date           *string `json:"end_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Lr_contest_type    *string `json:"lr_contest_type" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Reason             *string `json:"reason" default:""`
	Remark             *string `json:"remark" default:""`
	Resolution_date    *string `json:"resolution_date" default:""`
	Resolution_method  *string `json:"resolution_method" default:""`
	Resolution_remark  *string `json:"resolution_remark" default:""`
	Source             *string `json:"source" default:""`
	Start_date         *string `json:"start_date" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
