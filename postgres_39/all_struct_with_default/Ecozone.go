package dto

type Ecozone struct {
	Ecozone_id         string   `json:"ecozone_id" default:"source"`
	Active_ind         *string  `json:"active_ind" default:""`
	Base_depth         *float64 `json:"base_depth" default:""`
	Depth_ouom         *string  `json:"depth_ouom" default:""`
	Ecozone_type       *string  `json:"ecozone_type" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Preferred_name     *string  `json:"preferred_name" default:""`
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
