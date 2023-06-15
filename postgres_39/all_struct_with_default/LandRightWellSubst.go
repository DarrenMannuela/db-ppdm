package dto

type Land_right_well_subst struct {
	Uwi                string   `json:"uwi" default:"source"`
	Land_right_subtype string   `json:"land_right_subtype" default:"source"`
	Land_right_id      string   `json:"land_right_id" default:"source"`
	Lr_well_seq_no     int      `json:"lr_well_seq_no" default:"1"`
	Substance_id       string   `json:"substance_id" default:"source"`
	Active_ind         *string  `json:"active_ind" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Percent_psu        *float64 `json:"percent_psu" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Remark             *string  `json:"remark" default:""`
	Source             *string  `json:"source" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
