package dto

import (
	"time"
)

type Consent_cond struct {
	Consent_id         string     `json:"consent_id"`
	Condition_id       string     `json:"condition_id"`
	Active_ind         *string    `json:"active_ind"`
	Component_id       *string    `json:"component_id"`
	Consent_condition  *string    `json:"consent_condition"`
	Consent_type       *string    `json:"consent_type"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
