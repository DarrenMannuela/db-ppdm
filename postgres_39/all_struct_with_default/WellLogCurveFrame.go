package dto

type Well_log_curve_frame struct {
	Uwi                string   `json:"uwi" default:"source"`
	Well_log_id        string   `json:"well_log_id" default:"source"`
	Well_log_source    string   `json:"well_log_source" default:"source"`
	Frame_id           string   `json:"frame_id" default:"source"`
	Active_ind         *string  `json:"active_ind" default:""`
	Base_depth         *float64 `json:"base_depth" default:""`
	Depth_ouom         *string  `json:"depth_ouom" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Frame_name         *string  `json:"frame_name" default:""`
	Frame_spacing      *float64 `json:"frame_spacing" default:""`
	Frame_spacing_ouom *string  `json:"frame_spacing_ouom" default:""`
	Frame_spacing_uom  *string  `json:"frame_spacing_uom" default:""`
	Gaps_ind           *string  `json:"gaps_ind" default:""`
	Index_ouom         *string  `json:"index_ouom" default:""`
	Index_uom          *string  `json:"index_uom" default:""`
	Irregular_desc     *string  `json:"irregular_desc" default:""`
	Irregular_ind      *string  `json:"irregular_ind" default:""`
	Log_direction      *string  `json:"log_direction" default:""`
	Max_index          *float64 `json:"max_index" default:""`
	Min_index          *float64 `json:"min_index" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Primary_index_type *string  `json:"primary_index_type" default:""`
	Remark             *string  `json:"remark" default:""`
	Source             *string  `json:"source" default:""`
	Top_depth          *float64 `json:"top_depth" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
