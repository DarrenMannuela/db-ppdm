package dto

import (
	"time"
)

type Entitlement struct {
	Entitlement_id     string     `json:"entitlement_id"`
	Access_condition   *string    `json:"access_condition"`
	Access_cond_code   *string    `json:"access_cond_code"`
	Active_ind         *string    `json:"active_ind"`
	Description        *string    `json:"description"`
	Effective_date     *time.Time `json:"effective_date"`
	Entitlement_name   *string    `json:"entitlement_name"`
	Entitlement_type   *string    `json:"entitlement_type"`
	Expiry_action      *string    `json:"expiry_action"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Primary_term       *float64   `json:"primary_term"`
	Primary_term_ouom  *string    `json:"primary_term_ouom"`
	Remark             *string    `json:"remark"`
	Security_desc      *string    `json:"security_desc"`
	Source             *string    `json:"source"`
	Use_condition      *string    `json:"use_condition"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
