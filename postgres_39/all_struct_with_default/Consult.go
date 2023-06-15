package dto

type Consult struct {
	Consult_id         string  `json:"consult_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Complete_date      *string `json:"complete_date" default:""`
	Consult_reason     *string `json:"consult_reason" default:""`
	Consult_type       *string `json:"consult_type" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Period_type        *string `json:"period_type" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
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
