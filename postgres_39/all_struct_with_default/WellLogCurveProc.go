package dto

type Well_log_curve_proc struct {
	Uwi                string  `json:"uwi" default:"source"`
	Curve_id           string  `json:"curve_id" default:"source"`
	Process_obs_no     int     `json:"process_obs_no" default:"1"`
	Active_ind         *string `json:"active_ind" default:""`
	Dictionary_id      *string `json:"dictionary_id" default:""`
	Dict_process_id    *string `json:"dict_process_id" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Reported_mnemonic  *string `json:"reported_mnemonic" default:""`
	Source             *string `json:"source" default:""`
	Splice_obs_no      *int    `json:"splice_obs_no" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
