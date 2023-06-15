package dto

import (
	"time"
)

type Ba_preference_level struct {
	Business_associate_id string     `json:"business_associate_id"`
	Preference_type       string     `json:"preference_type"`
	Preference_obs_no     int        `json:"preference_obs_no"`
	Level_id              string     `json:"level_id"`
	Active_ind            *string    `json:"active_ind"`
	Currency_uom          *string    `json:"currency_uom"`
	Description           *string    `json:"description"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Preference_level      *float64   `json:"preference_level"`
	Preference_rule       *string    `json:"preference_rule"`
	Preferred_ba_id       *string    `json:"preferred_ba_id"`
	Remark                *string    `json:"remark"`
	Source                *string    `json:"source"`
	Wl_dictionary_id      *string    `json:"wl_dictionary_id"`
	Wl_dict_curve_id      *string    `json:"wl_dict_curve_id"`
	Wl_parameter_id       *string    `json:"wl_parameter_id"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
