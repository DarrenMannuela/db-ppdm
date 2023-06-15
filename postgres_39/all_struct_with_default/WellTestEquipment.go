package dto

type Well_test_equipment struct {
	Uwi                   string   `json:"uwi" default:"source"`
	Source                string   `json:"source" default:"source"`
	Test_type             string   `json:"test_type" default:"source"`
	Run_num               string   `json:"run_num" default:"source"`
	Test_num              string   `json:"test_num" default:"source"`
	Equipment_type        string   `json:"equipment_type" default:"source"`
	Equip_obs_no          int      `json:"equip_obs_no" default:"1"`
	Active_ind            *string  `json:"active_ind" default:""`
	Effective_date        *string  `json:"effective_date" default:""`
	Equipment_id          *string  `json:"equipment_id" default:""`
	Equip_length          *float64 `json:"equip_length" default:""`
	Equip_length_ouom     *string  `json:"equip_length_ouom" default:""`
	Equip_weight          *float64 `json:"equip_weight" default:""`
	Equip_weight_ouom     *string  `json:"equip_weight_ouom" default:""`
	Expiry_date           *string  `json:"expiry_date" default:""`
	Inside_diameter       *float64 `json:"inside_diameter" default:""`
	Inside_diameter_ouom  *string  `json:"inside_diameter_ouom" default:""`
	Outside_diameter      *float64 `json:"outside_diameter" default:""`
	Outside_diameter_ouom *string  `json:"outside_diameter_ouom" default:""`
	Ppdm_guid             *string  `json:"ppdm_guid" default:""`
	Remark                *string  `json:"remark" default:""`
	Top_depth             *float64 `json:"top_depth" default:""`
	Top_depth_ouom        *string  `json:"top_depth_ouom" default:""`
	Row_changed_by        *string  `json:"row_changed_by" default:""`
	Row_changed_date      *string  `json:"row_changed_date" default:""`
	Row_created_by        *string  `json:"row_created_by" default:""`
	Row_created_date      *string  `json:"row_created_date" default:""`
	Row_effective_date    *string  `json:"row_effective_date" default:""`
	Row_expiry_date       *string  `json:"row_expiry_date" default:""`
	Row_quality           *string  `json:"row_quality" default:""`
}
