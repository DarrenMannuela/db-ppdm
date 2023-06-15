package dto

import (
	"time"
)

type Ppdm_rule_enforcement struct {
	Rule_id            string     `json:"rule_id"`
	Enforcement_id     string     `json:"enforcement_id"`
	Abbreviation       *string    `json:"abbreviation"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Enforce_method     *string    `json:"enforce_method"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Long_name          *string    `json:"long_name"`
	Owner_ba_id        *string    `json:"owner_ba_id"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Procedure_id       *string    `json:"procedure_id"`
	Remark             *string    `json:"remark"`
	Rule_fail_result   *string    `json:"rule_fail_result"`
	Short_name         *string    `json:"short_name"`
	Source             *string    `json:"source"`
	Sw_application_id  *string    `json:"sw_application_id"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
