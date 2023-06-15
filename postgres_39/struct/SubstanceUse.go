package dto

import (
	"time"
)

type Substance_use struct {
	Substance_id            string     `json:"substance_id"`
	Substance_use_obs_no    int        `json:"substance_use_obs_no"`
	Active_ind              *string    `json:"active_ind"`
	Area_id                 *string    `json:"area_id"`
	Area_type               *string    `json:"area_type"`
	Business_associate_id   *string    `json:"business_associate_id"`
	Contract_id             *string    `json:"contract_id"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Land_right_id           *string    `json:"land_right_id"`
	Land_right_subtype      *string    `json:"land_right_subtype"`
	Ppdm_column_name        *string    `json:"ppdm_column_name"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Ppdm_system_id          *string    `json:"ppdm_system_id"`
	Ppdm_table_name         *string    `json:"ppdm_table_name"`
	Preferred_ind           *string    `json:"preferred_ind"`
	Project_id              *string    `json:"project_id"`
	Remark                  *string    `json:"remark"`
	Source                  *string    `json:"source"`
	Substance_alias_id      *string    `json:"substance_alias_id"`
	Substance_use_rule_desc *string    `json:"substance_use_rule_desc"`
	Substance_use_rule_type *string    `json:"substance_use_rule_type"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
