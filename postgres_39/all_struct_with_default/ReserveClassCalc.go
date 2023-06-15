package dto

type Reserve_class_calc struct {
	Reserve_class_id        string   `json:"reserve_class_id" default:"source"`
	Formula_id              string   `json:"formula_id" default:"source"`
	Calculation_seq_no      int      `json:"calculation_seq_no" default:"1"`
	Active_ind              *string  `json:"active_ind" default:""`
	Contribution_factor     *float64 `json:"contribution_factor" default:""`
	Effective_date          *string  `json:"effective_date" default:""`
	Expiry_date             *string  `json:"expiry_date" default:""`
	Origin_reserve_class_id *string  `json:"origin_reserve_class_id" default:""`
	Ppdm_guid               *string  `json:"ppdm_guid" default:""`
	Remark                  *string  `json:"remark" default:""`
	Source                  *string  `json:"source" default:""`
	Row_changed_by          *string  `json:"row_changed_by" default:""`
	Row_changed_date        *string  `json:"row_changed_date" default:""`
	Row_created_by          *string  `json:"row_created_by" default:""`
	Row_created_date        *string  `json:"row_created_date" default:""`
	Row_effective_date      *string  `json:"row_effective_date" default:""`
	Row_expiry_date         *string  `json:"row_expiry_date" default:""`
	Row_quality             *string  `json:"row_quality" default:""`
}
