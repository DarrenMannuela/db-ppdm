package dto

type Lith_rockpart_color struct {
	Lithology_log_id   string  `json:"lithology_log_id" default:"source"`
	Lith_log_source    string  `json:"lith_log_source" default:"source"`
	Depth_obs_no       int     `json:"depth_obs_no" default:"1"`
	Rock_type          string  `json:"rock_type" default:"source"`
	Rock_type_obs_no   int     `json:"rock_type_obs_no" default:"1"`
	Rockpart_name      string  `json:"rockpart_name" default:"source"`
	Basic_color        string  `json:"basic_color" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Color_adjective    *string `json:"color_adjective" default:""`
	Color_distribution *string `json:"color_distribution" default:""`
	Dominant_ind       *string `json:"dominant_ind" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Intensity          *string `json:"intensity" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Weathered_ind      *string `json:"weathered_ind" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
