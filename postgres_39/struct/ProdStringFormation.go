package dto

import (
	"time"
)

type Prod_string_formation struct {
	Uwi                 string     `json:"uwi"`
	Prod_string_source  string     `json:"prod_string_source"`
	String_id           string     `json:"string_id"`
	Pr_str_form_obs_no  int        `json:"pr_str_form_obs_no"`
	Active_ind          *string    `json:"active_ind"`
	Allocation_factor   *float64   `json:"allocation_factor"`
	Allocation_type     *string    `json:"allocation_type"`
	Base_depth          *float64   `json:"base_depth"`
	Base_depth_ouom     *string    `json:"base_depth_ouom"`
	Completion_obs_no   *int       `json:"completion_obs_no"`
	Current_status      *string    `json:"current_status"`
	Current_status_date *time.Time `json:"current_status_date"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Pool_id             *string    `json:"pool_id"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Source              *string    `json:"source"`
	Status_type         *string    `json:"status_type"`
	Strat_name_set_id   *string    `json:"strat_name_set_id"`
	Strat_unit_id       *string    `json:"strat_unit_id"`
	Top_depth           *float64   `json:"top_depth"`
	Top_depth_ouom      *string    `json:"top_depth_ouom"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
