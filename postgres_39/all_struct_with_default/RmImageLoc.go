package dto

type Rm_image_loc struct {
	Physical_item_id   string   `json:"physical_item_id" default:"source"`
	Log_section_id     string   `json:"log_section_id" default:"source"`
	Position_id        string   `json:"position_id" default:"source"`
	Active_ind         *string  `json:"active_ind" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Log_depth          *float64 `json:"log_depth" default:""`
	Log_depth_ouom     *string  `json:"log_depth_ouom" default:""`
	Log_depth_uom      *string  `json:"log_depth_uom" default:""`
	Position_type      *string  `json:"position_type" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Remark             *string  `json:"remark" default:""`
	Source             *string  `json:"source" default:""`
	X_position         *float64 `json:"x_position" default:""`
	Y_position         *float64 `json:"y_position" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
