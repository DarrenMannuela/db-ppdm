package dto

type Pool_version struct {
	Pool_id                string  `json:"pool_id" default:"source"`
	Pool_source            string  `json:"pool_source" default:"source"`
	Active_ind             *string `json:"active_ind" default:""`
	Business_associate_id  *string `json:"business_associate_id" default:""`
	Current_status_date    *string `json:"current_status_date" default:""`
	Discovery_date         *string `json:"discovery_date" default:""`
	Effective_date         *string `json:"effective_date" default:""`
	Expiry_date            *string `json:"expiry_date" default:""`
	Field_id               *string `json:"field_id" default:""`
	Pool_code              *string `json:"pool_code" default:""`
	Pool_name              *string `json:"pool_name" default:""`
	Pool_name_abbreviation *string `json:"pool_name_abbreviation" default:""`
	Pool_status            *string `json:"pool_status" default:""`
	Pool_type              *string `json:"pool_type" default:""`
	Ppdm_guid              *string `json:"ppdm_guid" default:""`
	Remark                 *string `json:"remark" default:""`
	Strat_name_set_id      *string `json:"strat_name_set_id" default:""`
	Strat_unit_id          *string `json:"strat_unit_id" default:""`
	Row_changed_by         *string `json:"row_changed_by" default:""`
	Row_changed_date       *string `json:"row_changed_date" default:""`
	Row_created_by         *string `json:"row_created_by" default:""`
	Row_created_date       *string `json:"row_created_date" default:""`
	Row_effective_date     *string `json:"row_effective_date" default:""`
	Row_expiry_date        *string `json:"row_expiry_date" default:""`
	Row_quality            *string `json:"row_quality" default:""`
}
