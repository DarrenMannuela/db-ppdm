package dto

type Strat_alias struct {
	Strat_name_set_id       string  `json:"strat_name_set_id" default:"source"`
	Strat_unit_id           string  `json:"strat_unit_id" default:"source"`
	Alias_strat_name_set_id string  `json:"alias_strat_name_set_id" default:"source"`
	Alias_strat_unit_id     string  `json:"alias_strat_unit_id" default:"source"`
	Source                  string  `json:"source" default:"source"`
	Abbreviation            *string `json:"abbreviation" default:""`
	Active_ind              *string `json:"active_ind" default:""`
	Alias_long_name         *string `json:"alias_long_name" default:""`
	Alias_reason_type       *string `json:"alias_reason_type" default:""`
	Alias_short_name        *string `json:"alias_short_name" default:""`
	Alias_type              *string `json:"alias_type" default:""`
	Amended_date            *string `json:"amended_date" default:""`
	Created_date            *string `json:"created_date" default:""`
	Effective_date          *string `json:"effective_date" default:""`
	Expiry_date             *string `json:"expiry_date" default:""`
	Language_id             *string `json:"language_id" default:""`
	Original_ind            *string `json:"original_ind" default:""`
	Owner_ba_id             *string `json:"owner_ba_id" default:""`
	Ppdm_guid               *string `json:"ppdm_guid" default:""`
	Preferred_ind           *string `json:"preferred_ind" default:""`
	Reason_desc             *string `json:"reason_desc" default:""`
	Remark                  *string `json:"remark" default:""`
	Source_document_id      *string `json:"source_document_id" default:""`
	Struckoff_date          *string `json:"struckoff_date" default:""`
	Sw_application_id       *string `json:"sw_application_id" default:""`
	Use_rule                *string `json:"use_rule" default:""`
	Row_changed_by          *string `json:"row_changed_by" default:""`
	Row_changed_date        *string `json:"row_changed_date" default:""`
	Row_created_by          *string `json:"row_created_by" default:""`
	Row_created_date        *string `json:"row_created_date" default:""`
	Row_effective_date      *string `json:"row_effective_date" default:""`
	Row_expiry_date         *string `json:"row_expiry_date" default:""`
	Row_quality             *string `json:"row_quality" default:""`
}
