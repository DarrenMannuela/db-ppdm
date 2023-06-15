package dto

type Well_horiz_drill_kop struct {
	Uwi                    string   `json:"uwi" default:"source"`
	Source                 string   `json:"source" default:"source"`
	Kickoff_point_obs_no   int      `json:"kickoff_point_obs_no" default:"1"`
	Active_ind             *string  `json:"active_ind" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Kickoff_point_md       *float64 `json:"kickoff_point_md" default:""`
	Kickoff_point_md_ouom  *string  `json:"kickoff_point_md_ouom" default:""`
	Kickoff_point_tvd      *float64 `json:"kickoff_point_tvd" default:""`
	Kickoff_point_tvd_ouom *string  `json:"kickoff_point_tvd_ouom" default:""`
	Lateral_hole_id        *string  `json:"lateral_hole_id" default:""`
	Node_id                *string  `json:"node_id" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Remark                 *string  `json:"remark" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
