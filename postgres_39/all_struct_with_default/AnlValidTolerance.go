package dto

type Anl_valid_tolerance struct {
	Method_set_id      string   `json:"method_set_id" default:"source"`
	Method_id          string   `json:"method_id" default:"source"`
	Equip_obs_no       int      `json:"equip_obs_no" default:"1"`
	Tolerance_obs_no   int      `json:"tolerance_obs_no" default:"1"`
	Active_ind         *string  `json:"active_ind" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Max_tolerance      *float64 `json:"max_tolerance" default:""`
	Max_tolerance_ouom *string  `json:"max_tolerance_ouom" default:""`
	Max_tolerance_rep  *float64 `json:"max_tolerance_rep" default:""`
	Max_tolerance_uom  *string  `json:"max_tolerance_uom" default:""`
	Min_tolerance      *float64 `json:"min_tolerance" default:""`
	Min_tolerance_ouom *string  `json:"min_tolerance_ouom" default:""`
	Min_tolerance_rep  *string  `json:"min_tolerance_rep" default:""`
	Min_tolerance_uom  *string  `json:"min_tolerance_uom" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Remark             *string  `json:"remark" default:""`
	Source             *string  `json:"source" default:""`
	Substance_id       *string  `json:"substance_id" default:""`
	Tolerance_desc     *string  `json:"tolerance_desc" default:""`
	Tolerance_type     *string  `json:"tolerance_type" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
