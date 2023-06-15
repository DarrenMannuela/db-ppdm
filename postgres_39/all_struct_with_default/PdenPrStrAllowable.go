package dto

type Pden_pr_str_allowable struct {
	Pden_subtype         string   `json:"pden_subtype" default:"source"`
	Pden_id              string   `json:"pden_id" default:"source"`
	Pden_source          string   `json:"pden_source" default:"source"`
	Uwi                  string   `json:"uwi" default:"source"`
	String_source        string   `json:"string_source" default:"source"`
	String_id            string   `json:"string_id" default:"source"`
	Pden_prs_xref_seq_no int      `json:"pden_prs_xref_seq_no" default:"1"`
	Allowable_obs_no     int      `json:"allowable_obs_no" default:"1"`
	Active_ind           *string  `json:"active_ind" default:""`
	Allowable            *float64 `json:"allowable" default:""`
	Allowable_date       *string  `json:"allowable_date" default:""`
	Allowable_date_desc  *string  `json:"allowable_date_desc" default:""`
	Allowable_ouom       *string  `json:"allowable_ouom" default:""`
	Allowable_uom        *string  `json:"allowable_uom" default:""`
	Effective_date       *string  `json:"effective_date" default:""`
	Expiry_date          *string  `json:"expiry_date" default:""`
	Period_type          *string  `json:"period_type" default:""`
	Ppdm_guid            *string  `json:"ppdm_guid" default:""`
	Product_type         *string  `json:"product_type" default:""`
	Remark               *string  `json:"remark" default:""`
	Source               *string  `json:"source" default:""`
	Row_changed_by       *string  `json:"row_changed_by" default:""`
	Row_changed_date     *string  `json:"row_changed_date" default:""`
	Row_created_by       *string  `json:"row_created_by" default:""`
	Row_created_date     *string  `json:"row_created_date" default:""`
	Row_effective_date   *string  `json:"row_effective_date" default:""`
	Row_expiry_date      *string  `json:"row_expiry_date" default:""`
	Row_quality          *string  `json:"row_quality" default:""`
}
