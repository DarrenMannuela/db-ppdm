package dto

import (
	"time"
)

type Well_core_strat_unit struct {
	Uwi                     string     `json:"uwi"`
	Source                  string     `json:"source"`
	Core_id                 string     `json:"core_id"`
	Description_obs_no      int        `json:"description_obs_no"`
	Strat_name_set_id       string     `json:"strat_name_set_id"`
	Core_strat_unit_id      string     `json:"core_strat_unit_id"`
	Active_ind              *string    `json:"active_ind"`
	Core_strat_unit_md      *float64   `json:"core_strat_unit_md"`
	Core_strat_unit_md_ouom *string    `json:"core_strat_unit_md_ouom"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Remark                  *string    `json:"remark"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
