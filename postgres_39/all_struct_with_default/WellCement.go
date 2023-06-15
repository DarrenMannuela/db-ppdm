package dto

type Well_cement struct {
	Uwi                         string   `json:"uwi" default:"source"`
	Well_tube_source            string   `json:"well_tube_source" default:"source"`
	Tubing_type                 string   `json:"tubing_type" default:"source"`
	Tubing_obs_no               int      `json:"tubing_obs_no" default:"1"`
	Cement_obs_no               int      `json:"cement_obs_no" default:"1"`
	Active_ind                  *string  `json:"active_ind" default:""`
	Cement_amount               *float64 `json:"cement_amount" default:""`
	Cement_amount_ouom          *string  `json:"cement_amount_ouom" default:""`
	Cement_amount_uom           *string  `json:"cement_amount_uom" default:""`
	Cement_ba_id                *string  `json:"cement_ba_id" default:""`
	Cement_type                 *string  `json:"cement_type" default:""`
	Effective_date              *string  `json:"effective_date" default:""`
	Expiry_date                 *string  `json:"expiry_date" default:""`
	Perforation_base_depth      *float64 `json:"perforation_base_depth" default:""`
	Perforation_base_depth_ouom *string  `json:"perforation_base_depth_ouom" default:""`
	Perforation_count           *int     `json:"perforation_count" default:""`
	Perforation_per_uom         *string  `json:"perforation_per_uom" default:""`
	Perforation_top_depth       *float64 `json:"perforation_top_depth" default:""`
	Perforation_top_depth_ouom  *string  `json:"perforation_top_depth_ouom" default:""`
	Ppdm_guid                   *string  `json:"ppdm_guid" default:""`
	Recement_ind                *string  `json:"recement_ind" default:""`
	Remark                      *string  `json:"remark" default:""`
	Source                      *string  `json:"source" default:""`
	Stage_no                    *int     `json:"stage_no" default:""`
	Row_changed_by              *string  `json:"row_changed_by" default:""`
	Row_changed_date            *string  `json:"row_changed_date" default:""`
	Row_created_by              *string  `json:"row_created_by" default:""`
	Row_created_date            *string  `json:"row_created_date" default:""`
	Row_effective_date          *string  `json:"row_effective_date" default:""`
	Row_expiry_date             *string  `json:"row_expiry_date" default:""`
	Row_quality                 *string  `json:"row_quality" default:""`
}
