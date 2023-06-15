package dto

import (
	"time"
)

type Ba_preference struct {
	Business_associate_id string     `json:"business_associate_id"`
	Preference_type       string     `json:"preference_type"`
	Preference_obs_no     int        `json:"preference_obs_no"`
	Active_ind            *string    `json:"active_ind"`
	Description           *string    `json:"description"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Language              *string    `json:"language"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Source                *string    `json:"source"`
	Sw_application_id     *string    `json:"sw_application_id"`
	Wl_curve_class_id     *string    `json:"wl_curve_class_id"`
	Wl_parameter_class_id *string    `json:"wl_parameter_class_id"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
