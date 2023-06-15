package dto

type Land_remark struct {
	Land_right_subtype string  `json:"land_right_subtype" default:"source"`
	Land_right_id      string  `json:"land_right_id" default:"source"`
	Remark_id          string  `json:"remark_id" default:"source"`
	Remark_seq_no      int     `json:"remark_seq_no" default:"1"`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Recommend_impl_ind *string `json:"recommend_impl_ind" default:""`
	Remark             *string `json:"remark" default:""`
	Remark_ba_id       *string `json:"remark_ba_id" default:""`
	Remark_date        *string `json:"remark_date" default:""`
	Remark_type        *string `json:"remark_type" default:""`
	Remark_user        *string `json:"remark_user" default:""`
	Source             *string `json:"source" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
