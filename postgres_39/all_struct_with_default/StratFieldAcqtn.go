package dto

type Strat_field_acqtn struct {
	Strat_field_acqtn_id string  `json:"strat_field_acqtn_id" default:"source"`
	Field_station_id     string  `json:"field_station_id" default:"source"`
	Strat_name_set_id    string  `json:"strat_name_set_id" default:"source"`
	Strat_unit_id        string  `json:"strat_unit_id" default:"source"`
	Interp_id            string  `json:"interp_id" default:"source"`
	Active_ind           *string `json:"active_ind" default:""`
	Core_id              *string `json:"core_id" default:""`
	Core_source          *string `json:"core_source" default:""`
	Description          *string `json:"description" default:""`
	Effective_date       *string `json:"effective_date" default:""`
	Expiry_date          *string `json:"expiry_date" default:""`
	Information_item_id  *string `json:"information_item_id" default:""`
	Info_item_subtype    *string `json:"info_item_subtype" default:""`
	Log_curve_id         *string `json:"log_curve_id" default:""`
	Ppdm_guid            *string `json:"ppdm_guid" default:""`
	Project_id           *string `json:"project_id" default:""`
	Remark               *string `json:"remark" default:""`
	Seis_set_id          *string `json:"seis_set_id" default:""`
	Seis_set_subtype     *string `json:"seis_set_subtype" default:""`
	Source               *string `json:"source" default:""`
	Strat_column_id      *string `json:"strat_column_id" default:""`
	Strat_column_source  *string `json:"strat_column_source" default:""`
	Uwi                  *string `json:"uwi" default:""`
	Row_changed_by       *string `json:"row_changed_by" default:""`
	Row_changed_date     *string `json:"row_changed_date" default:""`
	Row_created_by       *string `json:"row_created_by" default:""`
	Row_created_date     *string `json:"row_created_date" default:""`
	Row_effective_date   *string `json:"row_effective_date" default:""`
	Row_expiry_date      *string `json:"row_expiry_date" default:""`
	Row_quality          *string `json:"row_quality" default:""`
}
