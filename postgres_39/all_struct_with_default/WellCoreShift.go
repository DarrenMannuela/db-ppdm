package dto

type Well_core_shift struct {
	Uwi                   string   `json:"uwi" default:"source"`
	Source                string   `json:"source" default:"source"`
	Core_id               string   `json:"core_id" default:"source"`
	Shift_obs_no          int      `json:"shift_obs_no" default:"1"`
	Active_ind            *string  `json:"active_ind" default:""`
	Base_depth            *float64 `json:"base_depth" default:""`
	Base_depth_ouom       *string  `json:"base_depth_ouom" default:""`
	Base_shift_depth      *float64 `json:"base_shift_depth" default:""`
	Base_shift_depth_ouom *string  `json:"base_shift_depth_ouom" default:""`
	Core_shift_company    *string  `json:"core_shift_company" default:""`
	Core_shift_ind        *string  `json:"core_shift_ind" default:""`
	Core_shift_method     *string  `json:"core_shift_method" default:""`
	Effective_date        *string  `json:"effective_date" default:""`
	Expiry_date           *string  `json:"expiry_date" default:""`
	Ppdm_guid             *string  `json:"ppdm_guid" default:""`
	Remark                *string  `json:"remark" default:""`
	Top_depth             *float64 `json:"top_depth" default:""`
	Top_depth_ouom        *string  `json:"top_depth_ouom" default:""`
	Top_shift_depth       *float64 `json:"top_shift_depth" default:""`
	Top_shift_depth_ouom  *string  `json:"top_shift_depth_ouom" default:""`
	Row_changed_by        *string  `json:"row_changed_by" default:""`
	Row_changed_date      *string  `json:"row_changed_date" default:""`
	Row_created_by        *string  `json:"row_created_by" default:""`
	Row_created_date      *string  `json:"row_created_date" default:""`
	Row_effective_date    *string  `json:"row_effective_date" default:""`
	Row_expiry_date       *string  `json:"row_expiry_date" default:""`
	Row_quality           *string  `json:"row_quality" default:""`
}
