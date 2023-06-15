package dto

type Strat_unit struct {
	Strat_name_set_id      string   `json:"strat_name_set_id" default:"source"`
	Strat_unit_id          string   `json:"strat_unit_id" default:"source"`
	Abbreviation           *string  `json:"abbreviation" default:""`
	Active_ind             *string  `json:"active_ind" default:""`
	Area_id                *string  `json:"area_id" default:""`
	Area_type              *string  `json:"area_type" default:""`
	Business_associate_id  *string  `json:"business_associate_id" default:""`
	Confidence_id          *string  `json:"confidence_id" default:""`
	Current_status_date    *string  `json:"current_status_date" default:""`
	Description            *string  `json:"description" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Fault_type             *string  `json:"fault_type" default:""`
	Form_code              *string  `json:"form_code" default:""`
	Group_code             *string  `json:"group_code" default:""`
	Long_name              *string  `json:"long_name" default:""`
	Ordinal_age_code       *float64 `json:"ordinal_age_code" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Preferred_ind          *string  `json:"preferred_ind" default:""`
	Remark                 *string  `json:"remark" default:""`
	Short_name             *string  `json:"short_name" default:""`
	Source                 *string  `json:"source" default:""`
	Strat_interpret_method *string  `json:"strat_interpret_method" default:""`
	Strat_status           *string  `json:"strat_status" default:""`
	Strat_type             *string  `json:"strat_type" default:""`
	Strat_unit_type        *string  `json:"strat_unit_type" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
