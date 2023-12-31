package dto

type Anl_fluor struct {
	Analysis_id              string   `json:"analysis_id" default:"source"`
	Anl_source               string   `json:"anl_source" default:"source"`
	Fluor_obs_no             int      `json:"fluor_obs_no" default:"1"`
	Active_ind               *string  `json:"active_ind" default:""`
	Color_remark             *string  `json:"color_remark" default:""`
	Effective_date           *string  `json:"effective_date" default:""`
	Expiry_date              *string  `json:"expiry_date" default:""`
	Fluor_percent            *float64 `json:"fluor_percent" default:""`
	Fluor_remark             *string  `json:"fluor_remark" default:""`
	From_color               *string  `json:"from_color" default:""`
	From_color_frequency     *float64 `json:"from_color_frequency" default:""`
	From_color_frequency_uom *string  `json:"from_color_frequency_uom" default:""`
	From_color_intensity     *float64 `json:"from_color_intensity" default:""`
	From_color_intensity_uom *string  `json:"from_color_intensity_uom" default:""`
	From_mobility_type       *string  `json:"from_mobility_type" default:""`
	Hydrocarbon_show_ind     *string  `json:"hydrocarbon_show_ind" default:""`
	Ppdm_guid                *string  `json:"ppdm_guid" default:""`
	Preferred_ind            *string  `json:"preferred_ind" default:""`
	Problem_ind              *string  `json:"problem_ind" default:""`
	Remark                   *string  `json:"remark" default:""`
	Source                   *string  `json:"source" default:""`
	Step_seq_no              *int     `json:"step_seq_no" default:""`
	To_color                 *string  `json:"to_color" default:""`
	To_color_frequency       *float64 `json:"to_color_frequency" default:""`
	To_color_frequency_uom   *string  `json:"to_color_frequency_uom" default:""`
	To_color_intensity       *float64 `json:"to_color_intensity" default:""`
	To_color_intensity_uom   *string  `json:"to_color_intensity_uom" default:""`
	To_mobility_type         *string  `json:"to_mobility_type" default:""`
	Row_changed_by           *string  `json:"row_changed_by" default:""`
	Row_changed_date         *string  `json:"row_changed_date" default:""`
	Row_created_by           *string  `json:"row_created_by" default:""`
	Row_created_date         *string  `json:"row_created_date" default:""`
	Row_effective_date       *string  `json:"row_effective_date" default:""`
	Row_expiry_date          *string  `json:"row_expiry_date" default:""`
	Row_quality              *string  `json:"row_quality" default:""`
}
