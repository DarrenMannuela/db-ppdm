package dto

type Sf_platform struct {
	Sf_subtype         string   `json:"sf_subtype" default:"source"`
	Platform_id        string   `json:"platform_id" default:"source"`
	Active_ind         *string  `json:"active_ind" default:""`
	Drill_slot_count   *int     `json:"drill_slot_count" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Installed_date     *string  `json:"installed_date" default:""`
	Platform_name      *string  `json:"platform_name" default:""`
	Platform_type      *string  `json:"platform_type" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Remark             *string  `json:"remark" default:""`
	Removal_date       *string  `json:"removal_date" default:""`
	Source             *string  `json:"source" default:""`
	Water_depth        *float64 `json:"water_depth" default:""`
	Water_depth_datum  *string  `json:"water_depth_datum" default:""`
	Water_depth_ouom   *string  `json:"water_depth_ouom" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
