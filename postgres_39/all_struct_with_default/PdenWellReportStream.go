package dto

type Pden_well_report_stream struct {
	Pden_subtype       string   `json:"pden_subtype" default:"source"`
	Pden_id            string   `json:"pden_id" default:"source"`
	Pden_source        string   `json:"pden_source" default:"source"`
	Uwi                string   `json:"uwi" default:"source"`
	Wrs_obs_no         int      `json:"wrs_obs_no" default:"1"`
	Active_ind         *string  `json:"active_ind" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Remark             *string  `json:"remark" default:""`
	Substance_id       *string  `json:"substance_id" default:""`
	Substance_percent  *float64 `json:"substance_percent" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}