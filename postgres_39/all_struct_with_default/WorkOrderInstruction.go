package dto

type Work_order_instruction struct {
	Work_order_id          string   `json:"work_order_id" default:"source"`
	Instruction_id         string   `json:"instruction_id" default:"source"`
	Active_ind             *string  `json:"active_ind" default:""`
	Ba_address_obs_no      *int     `json:"ba_address_obs_no" default:""`
	Ba_address_source      *string  `json:"ba_address_source" default:""`
	Ba_obs_no              *int     `json:"ba_obs_no" default:""`
	Business_associate_id  *string  `json:"business_associate_id" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Instruction_desc       *string  `json:"instruction_desc" default:""`
	Instruction_type       *string  `json:"instruction_type" default:""`
	Instruction_value      *float64 `json:"instruction_value" default:""`
	Instruction_value_ouom *string  `json:"instruction_value_ouom" default:""`
	Instruction_value_uom  *string  `json:"instruction_value_uom" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Remark                 *string  `json:"remark" default:""`
	Source                 *string  `json:"source" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
