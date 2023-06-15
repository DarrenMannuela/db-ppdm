package dto

type Paleo_sum_sample struct {
	Paleo_summary_id   string  `json:"paleo_summary_id" default:"source"`
	Lith_sample_id     string  `json:"lith_sample_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Diversity_count    *int    `json:"diversity_count" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	First_sample_ind   *string `json:"first_sample_ind" default:""`
	Last_sample_ind    *string `json:"last_sample_ind" default:""`
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
