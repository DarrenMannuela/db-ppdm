package dto

type Resent_eco_schedule struct {
	Resent_id           string   `json:"resent_id" default:"source"`
	Reserve_class_id    string   `json:"reserve_class_id" default:"source"`
	Economics_run_id    string   `json:"economics_run_id" default:"source"`
	Schedule_id         string   `json:"schedule_id" default:"source"`
	Active_ind          *string  `json:"active_ind" default:""`
	Economic_schedule   *string  `json:"economic_schedule" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Max_value           *float64 `json:"max_value" default:""`
	Max_value_ouom      *string  `json:"max_value_ouom" default:""`
	Max_value_uom       *string  `json:"max_value_uom" default:""`
	Min_value           *float64 `json:"min_value" default:""`
	Min_value_ouom      *string  `json:"min_value_ouom" default:""`
	Min_value_uom       *string  `json:"min_value_uom" default:""`
	Period_type         *string  `json:"period_type" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Product_ind         *string  `json:"product_ind" default:""`
	Product_type        *string  `json:"product_type" default:""`
	Remark              *string  `json:"remark" default:""`
	Schedule_date       *string  `json:"schedule_date" default:""`
	Schedule_date_desc  *string  `json:"schedule_date_desc" default:""`
	Schedule_desc       *string  `json:"schedule_desc" default:""`
	Schedule_value      *float64 `json:"schedule_value" default:""`
	Schedule_value_ouom *string  `json:"schedule_value_ouom" default:""`
	Schedule_value_uom  *string  `json:"schedule_value_uom" default:""`
	Source              *string  `json:"source" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}
