package dto

type Pr_lse_unit_str_hist struct {
	Uwi                string  `json:"uwi" default:"source"`
	Prod_string_source string  `json:"prod_string_source" default:"source"`
	String_id          string  `json:"string_id" default:"source"`
	Lease_unit_id      string  `json:"lease_unit_id" default:"source"`
	Status_id          string  `json:"status_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	End_date           *string `json:"end_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
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
