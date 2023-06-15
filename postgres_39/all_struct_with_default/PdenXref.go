package dto

type Pden_xref struct {
	From_pden_subtype  string  `json:"from_pden_subtype" default:"source"`
	From_pden_id       string  `json:"from_pden_id" default:"source"`
	From_pden_source   string  `json:"from_pden_source" default:"source"`
	To_pden_subtype    string  `json:"to_pden_subtype" default:"source"`
	To_pden_id         string  `json:"to_pden_id" default:"source"`
	To_pden_source     string  `json:"to_pden_source" default:"source"`
	Xref_obs_no        int     `json:"xref_obs_no" default:"1"`
	Active_ind         *string `json:"active_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Pden_xref_type     *string `json:"pden_xref_type" default:""`
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
