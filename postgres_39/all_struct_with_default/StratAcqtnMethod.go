package dto

type Strat_acqtn_method struct {
	Strat_acqtn_method_id string  `json:"strat_acqtn_method_id" default:"source"`
	Acqtn_date            *string `json:"acqtn_date" default:""`
	Acqtn_method_type     *string `json:"acqtn_method_type" default:""`
	Active_ind            *string `json:"active_ind" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	Field_station_id      *string `json:"field_station_id" default:""`
	Interp_id             *string `json:"interp_id" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
	Remark                *string `json:"remark" default:""`
	Source                *string `json:"source" default:""`
	Strat_column_id       *string `json:"strat_column_id" default:""`
	Strat_column_source   *string `json:"strat_column_source" default:""`
	Strat_name_set_id     *string `json:"strat_name_set_id" default:""`
	Strat_unit_id         *string `json:"strat_unit_id" default:""`
	Uwi                   *string `json:"uwi" default:""`
	Row_changed_by        *string `json:"row_changed_by" default:""`
	Row_changed_date      *string `json:"row_changed_date" default:""`
	Row_created_by        *string `json:"row_created_by" default:""`
	Row_created_date      *string `json:"row_created_date" default:""`
	Row_effective_date    *string `json:"row_effective_date" default:""`
	Row_expiry_date       *string `json:"row_expiry_date" default:""`
	Row_quality           *string `json:"row_quality" default:""`
}
