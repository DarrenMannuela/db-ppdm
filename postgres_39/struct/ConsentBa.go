package dto

import (
	"time"
)

type Consent_ba struct {
	Consent_id            string     `json:"consent_id"`
	Consent_ba_obs_no     int        `json:"consent_ba_obs_no"`
	Active_ind            *string    `json:"active_ind"`
	Business_associate_id *string    `json:"business_associate_id"`
	Description           *string    `json:"description"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Full_name             *string    `json:"full_name"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Role                  *string    `json:"role"`
	Source                *string    `json:"source"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
