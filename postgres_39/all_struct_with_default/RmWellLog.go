package dto

type Rm_well_log struct {
	Info_item_subtype        string   `json:"info_item_subtype" default:"source"`
	Information_item_id      string   `json:"information_item_id" default:"source"`
	Active_ind               *string  `json:"active_ind" default:""`
	Display_interval         *float64 `json:"display_interval" default:""`
	Display_interval_uom     *string  `json:"display_interval_uom" default:""`
	Display_scale            *float64 `json:"display_scale" default:""`
	Display_scale_uom        *string  `json:"display_scale_uom" default:""`
	Effective_date           *string  `json:"effective_date" default:""`
	Expiry_date              *string  `json:"expiry_date" default:""`
	Interpreted_ind          *string  `json:"interpreted_ind" default:""`
	Ppdm_guid                *string  `json:"ppdm_guid" default:""`
	Remark                   *string  `json:"remark" default:""`
	Reported_base_depth      *float64 `json:"reported_base_depth" default:""`
	Reported_base_depth_ouom *string  `json:"reported_base_depth_ouom" default:""`
	Reported_top_depth       *float64 `json:"reported_top_depth" default:""`
	Reported_top_depth_ouom  *string  `json:"reported_top_depth_ouom" default:""`
	Source                   *string  `json:"source" default:""`
	Row_changed_by           *string  `json:"row_changed_by" default:""`
	Row_changed_date         *string  `json:"row_changed_date" default:""`
	Row_created_by           *string  `json:"row_created_by" default:""`
	Row_created_date         *string  `json:"row_created_date" default:""`
	Row_effective_date       *string  `json:"row_effective_date" default:""`
	Row_expiry_date          *string  `json:"row_expiry_date" default:""`
	Row_quality              *string  `json:"row_quality" default:""`
}
