package dto

type Well_drill_period_inv struct {
	Uwi                     string   `json:"uwi" default:"source"`
	Period_obs_no           int      `json:"period_obs_no" default:"1"`
	Inventory_seq_no        int      `json:"inventory_seq_no" default:"1"`
	Active_ind              *string  `json:"active_ind" default:""`
	Cat_additive_id         *string  `json:"cat_additive_id" default:""`
	Cat_equip_id            *string  `json:"cat_equip_id" default:""`
	Effective_date          *string  `json:"effective_date" default:""`
	Equipment_id            *string  `json:"equipment_id" default:""`
	Expiry_date             *string  `json:"expiry_date" default:""`
	Inventory_material_type *string  `json:"inventory_material_type" default:""`
	Ppdm_guid               *string  `json:"ppdm_guid" default:""`
	Quantity_on_hand        *float64 `json:"quantity_on_hand" default:""`
	Quantity_on_hand_ouom   *string  `json:"quantity_on_hand_ouom" default:""`
	Quantity_on_hand_uom    *string  `json:"quantity_on_hand_uom" default:""`
	Quantity_ordered        *float64 `json:"quantity_ordered" default:""`
	Quantity_ordered_ouom   *string  `json:"quantity_ordered_ouom" default:""`
	Quantity_ordered_uom    *string  `json:"quantity_ordered_uom" default:""`
	Remark                  *string  `json:"remark" default:""`
	Report_time             *string  `json:"report_time" default:""`
	Report_timezone         *string  `json:"report_timezone" default:""`
	Report_time_type        *string  `json:"report_time_type" default:""`
	Source                  *string  `json:"source" default:""`
	Row_changed_by          *string  `json:"row_changed_by" default:""`
	Row_changed_date        *string  `json:"row_changed_date" default:""`
	Row_created_by          *string  `json:"row_created_by" default:""`
	Row_created_date        *string  `json:"row_created_date" default:""`
	Row_effective_date      *string  `json:"row_effective_date" default:""`
	Row_expiry_date         *string  `json:"row_expiry_date" default:""`
	Row_quality             *string  `json:"row_quality" default:""`
}
