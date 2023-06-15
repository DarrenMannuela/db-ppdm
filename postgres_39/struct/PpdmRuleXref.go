package dto

import (
	"time"
)

type Ppdm_rule_xref struct {
	Dependency_id       string     `json:"dependency_id"`
	Rule_id             string     `json:"rule_id"`
	Rule_id2            string     `json:"rule_id_2"`
	Xref_obs_no         int        `json:"xref_obs_no"`
	Active_ind          *string    `json:"active_ind"`
	Description         *string    `json:"description"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Rule_xref_condition *string    `json:"rule_xref_condition"`
	Rule_xref_type      *string    `json:"rule_xref_type"`
	Source              *string    `json:"source"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
