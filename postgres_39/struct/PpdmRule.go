package dto

import (
	"time"
)

type Ppdm_rule struct {
	Rule_id            string     `json:"rule_id"`
	Active_ind         *string    `json:"active_ind"`
	Current_status     *string    `json:"current_status"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Rule_class         *string    `json:"rule_class"`
	Rule_desc          *string    `json:"rule_desc"`
	Rule_purpose       *string    `json:"rule_purpose"`
	Rule_query         *string    `json:"rule_query"`
	Source             *string    `json:"source"`
	Use_condition_desc *string    `json:"use_condition_desc"`
	Use_condition_type *string    `json:"use_condition_type"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
