package dto

type Well_log_curve_splice struct {
	Uwi                 string   `json:"uwi" default:"source"`
	Curve_id            string   `json:"curve_id" default:"source"`
	Splice_obs_no       int      `json:"splice_obs_no" default:"1"`
	Active_ind          *string  `json:"active_ind" default:""`
	Created_by_ba_id    *string  `json:"created_by_ba_id" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Information_item_id *string  `json:"information_item_id" default:""`
	Info_item_subtype   *string  `json:"info_item_subtype" default:""`
	Max_value_index     *float64 `json:"max_value_index" default:""`
	Min_value_index     *float64 `json:"min_value_index" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Remark              *string  `json:"remark" default:""`
	Source              *string  `json:"source" default:""`
	Source_curve_id     *string  `json:"source_curve_id" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}
