package dto

type Well_horiz_drill_poe struct {
	Uwi                       string   `json:"uwi" default:"source"`
	Source                    string   `json:"source" default:"source"`
	Point_of_entry_obs_no     int      `json:"point_of_entry_obs_no" default:"1"`
	Active_ind                *string  `json:"active_ind" default:""`
	Effective_date            *string  `json:"effective_date" default:""`
	Expiry_date               *string  `json:"expiry_date" default:""`
	Lateral_hole_id           *string  `json:"lateral_hole_id" default:""`
	Node_id                   *string  `json:"node_id" default:""`
	Point_of_entry_md         *float64 `json:"point_of_entry_md" default:""`
	Point_of_entry_md_ouom    *string  `json:"point_of_entry_md_ouom" default:""`
	Point_of_entry_strat_unit *string  `json:"point_of_entry_strat_unit" default:""`
	Point_of_entry_tvd        *float64 `json:"point_of_entry_tvd" default:""`
	Point_of_entry_tvd_ouom   *string  `json:"point_of_entry_tvd_ouom" default:""`
	Ppdm_guid                 *string  `json:"ppdm_guid" default:""`
	Remark                    *string  `json:"remark" default:""`
	Strat_name_set_id         *string  `json:"strat_name_set_id" default:""`
	Row_changed_by            *string  `json:"row_changed_by" default:""`
	Row_changed_date          *string  `json:"row_changed_date" default:""`
	Row_created_by            *string  `json:"row_created_by" default:""`
	Row_created_date          *string  `json:"row_created_date" default:""`
	Row_effective_date        *string  `json:"row_effective_date" default:""`
	Row_expiry_date           *string  `json:"row_expiry_date" default:""`
	Row_quality               *string  `json:"row_quality" default:""`
}
