package dto

type Resent_vol_revision struct {
	Resent_id            string   `json:"resent_id" default:"source"`
	Reserve_class_id     string   `json:"reserve_class_id" default:"source"`
	Revision_id          string   `json:"revision_id" default:"source"`
	Revision_obs_no      int      `json:"revision_obs_no" default:"1"`
	Active_ind           *string  `json:"active_ind" default:""`
	Analyst_ba_id        *string  `json:"analyst_ba_id" default:""`
	Approved_by_ba_id    *string  `json:"approved_by_ba_id" default:""`
	Approved_date        *string  `json:"approved_date" default:""`
	Effective_date       *string  `json:"effective_date" default:""`
	Expiry_date          *string  `json:"expiry_date" default:""`
	Gross_revision_ind   *string  `json:"gross_revision_ind" default:""`
	Interest_set_id      *string  `json:"interest_set_id" default:""`
	Interest_set_seq_no  *int     `json:"interest_set_seq_no" default:""`
	Net_revision_ind     *string  `json:"net_revision_ind" default:""`
	New_volume           *float64 `json:"new_volume" default:""`
	New_volume_ouom      *string  `json:"new_volume_ouom" default:""`
	New_volume_uom       *string  `json:"new_volume_uom" default:""`
	Partner_ba_id        *string  `json:"partner_ba_id" default:""`
	Partner_obs_no       *int     `json:"partner_obs_no" default:""`
	Ppdm_guid            *string  `json:"ppdm_guid" default:""`
	Product_type         *string  `json:"product_type" default:""`
	Project_id           *string  `json:"project_id" default:""`
	Remark               *string  `json:"remark" default:""`
	Report_ind           *string  `json:"report_ind" default:""`
	Revision_category_id *string  `json:"revision_category_id" default:""`
	Revision_date        *string  `json:"revision_date" default:""`
	Revision_date_desc   *string  `json:"revision_date_desc" default:""`
	Revision_method      *string  `json:"revision_method" default:""`
	Revision_volume      *float64 `json:"revision_volume" default:""`
	Revision_volume_ouom *string  `json:"revision_volume_ouom" default:""`
	Revision_volume_uom  *string  `json:"revision_volume_uom" default:""`
	Source               *string  `json:"source" default:""`
	Summary_id           *string  `json:"summary_id" default:""`
	Summary_obs_no       *int     `json:"summary_obs_no" default:""`
	Row_changed_by       *string  `json:"row_changed_by" default:""`
	Row_changed_date     *string  `json:"row_changed_date" default:""`
	Row_created_by       *string  `json:"row_created_by" default:""`
	Row_created_date     *string  `json:"row_created_date" default:""`
	Row_effective_date   *string  `json:"row_effective_date" default:""`
	Row_expiry_date      *string  `json:"row_expiry_date" default:""`
	Row_quality          *string  `json:"row_quality" default:""`
}