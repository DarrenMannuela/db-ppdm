package dto

type Well_log_axis_coord struct {
	Uwi                string   `json:"uwi" default:"source"`
	Curve_id           string   `json:"curve_id" default:"source"`
	Axis_id            string   `json:"axis_id" default:"source"`
	Coordinate_seq_no  int      `json:"coordinate_seq_no" default:"1"`
	Active_ind         *string  `json:"active_ind" default:""`
	Axis_value         *float64 `json:"axis_value" default:""`
	Axis_value_ouom    *string  `json:"axis_value_ouom" default:""`
	Axis_value_uom     *string  `json:"axis_value_uom" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Remark             *string  `json:"remark" default:""`
	Source             *string  `json:"source" default:""`
	Text_value         *string  `json:"text_value" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
