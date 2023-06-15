package dto

type Seis_license_cond struct {
	Seis_set_subtype    string   `json:"seis_set_subtype" default:"source"`
	Seis_set_id         string   `json:"seis_set_id" default:"source"`
	License_id          string   `json:"license_id" default:"source"`
	Condition_id        string   `json:"condition_id" default:"source"`
	Active_ind          *string  `json:"active_ind" default:""`
	Condition_code      *string  `json:"condition_code" default:""`
	Condition_desc      *string  `json:"condition_desc" default:""`
	Condition_type      *string  `json:"condition_type" default:""`
	Condition_value     *float64 `json:"condition_value" default:""`
	Condition_value_uom *string  `json:"condition_value_uom" default:""`
	Contact_ba_id       *string  `json:"contact_ba_id" default:""`
	Date_format_desc    *string  `json:"date_format_desc" default:""`
	Description         *string  `json:"description" default:""`
	Due_condition       *string  `json:"due_condition" default:""`
	Due_date            *string  `json:"due_date" default:""`
	Due_frequency       *string  `json:"due_frequency" default:""`
	Due_term            *float64 `json:"due_term" default:""`
	Due_term_uom        *string  `json:"due_term_uom" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Exempt_ind          *string  `json:"exempt_ind" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Fulfilled_by_ba_id  *string  `json:"fulfilled_by_ba_id" default:""`
	Fulfilled_date      *string  `json:"fulfilled_date" default:""`
	Fulfilled_ind       *string  `json:"fulfilled_ind" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Remark              *string  `json:"remark" default:""`
	Restriction_id      *string  `json:"restriction_id" default:""`
	Restriction_version *int     `json:"restriction_version" default:""`
	Source              *string  `json:"source" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}
