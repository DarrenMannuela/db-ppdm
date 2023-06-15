package dto

type Lith_rockpart_grain_size struct {
	Lithology_log_id   string  `json:"lithology_log_id" default:"source"`
	Lith_log_source    string  `json:"lith_log_source" default:"source"`
	Depth_obs_no       int     `json:"depth_obs_no" default:"1"`
	Scale_scheme       string  `json:"scale_scheme" default:"source"`
	Rock_type          string  `json:"rock_type" default:"source"`
	Rock_type_obs_no   int     `json:"rock_type_obs_no" default:"1"`
	Rockpart_name      string  `json:"rockpart_name" default:"source"`
	Grain_size         string  `json:"grain_size" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Distribution_ind   *string `json:"distribution_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Size_type_ind      *string `json:"size_type_ind" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
