package dto

type Lith_diagenesis struct {
	Lithology_log_id         string   `json:"lithology_log_id" default:"source"`
	Lith_log_source          string   `json:"lith_log_source" default:"source"`
	Depth_obs_no             int      `json:"depth_obs_no" default:"1"`
	Rock_type                string   `json:"rock_type" default:"source"`
	Rock_type_obs_no         int      `json:"rock_type_obs_no" default:"1"`
	Diagenesis_type          string   `json:"diagenesis_type" default:"source"`
	Active_ind               *string  `json:"active_ind" default:""`
	Diagenesis_percent       *float64 `json:"diagenesis_percent" default:""`
	Diagenesis_rel_abundance *string  `json:"diagenesis_rel_abundance" default:""`
	Effective_date           *string  `json:"effective_date" default:""`
	Expiry_date              *string  `json:"expiry_date" default:""`
	Ppdm_guid                *string  `json:"ppdm_guid" default:""`
	Remark                   *string  `json:"remark" default:""`
	Row_changed_by           *string  `json:"row_changed_by" default:""`
	Row_changed_date         *string  `json:"row_changed_date" default:""`
	Row_created_by           *string  `json:"row_created_by" default:""`
	Row_created_date         *string  `json:"row_created_date" default:""`
	Row_effective_date       *string  `json:"row_effective_date" default:""`
	Row_expiry_date          *string  `json:"row_expiry_date" default:""`
	Row_quality              *string  `json:"row_quality" default:""`
}
