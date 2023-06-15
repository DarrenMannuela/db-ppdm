package dto

type Ba_preference_level struct {
	Business_associate_id string   `json:"business_associate_id" default:"source"`
	Preference_type       string   `json:"preference_type" default:"source"`
	Preference_obs_no     int      `json:"preference_obs_no" default:"1"`
	Level_id              string   `json:"level_id" default:"source"`
	Active_ind            *string  `json:"active_ind" default:""`
	Currency_uom          *string  `json:"currency_uom" default:""`
	Description           *string  `json:"description" default:""`
	Effective_date        *string  `json:"effective_date" default:""`
	Expiry_date           *string  `json:"expiry_date" default:""`
	Ppdm_guid             *string  `json:"ppdm_guid" default:""`
	Preference_level      *float64 `json:"preference_level" default:""`
	Preference_rule       *string  `json:"preference_rule" default:""`
	Preferred_ba_id       *string  `json:"preferred_ba_id" default:""`
	Remark                *string  `json:"remark" default:""`
	Source                *string  `json:"source" default:""`
	Wl_dictionary_id      *string  `json:"wl_dictionary_id" default:""`
	Wl_dict_curve_id      *string  `json:"wl_dict_curve_id" default:""`
	Wl_parameter_id       *string  `json:"wl_parameter_id" default:""`
	Row_changed_by        *string  `json:"row_changed_by" default:""`
	Row_changed_date      *string  `json:"row_changed_date" default:""`
	Row_created_by        *string  `json:"row_created_by" default:""`
	Row_created_date      *string  `json:"row_created_date" default:""`
	Row_effective_date    *string  `json:"row_effective_date" default:""`
	Row_expiry_date       *string  `json:"row_expiry_date" default:""`
	Row_quality           *string  `json:"row_quality" default:""`
}
