package dto

type Well_log_remark struct {
	Uwi                string  `json:"uwi" default:"source"`
	Well_log_id        string  `json:"well_log_id" default:"source"`
	Well_log_source    string  `json:"well_log_source" default:"source"`
	Remark_obs_no      int     `json:"remark_obs_no" default:"1"`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Made_by_ba_id      *string `json:"made_by_ba_id" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Remark_type        *string `json:"remark_type" default:""`
	Remark_use_type    *string `json:"remark_use_type" default:""`
	Source             *string `json:"source" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
