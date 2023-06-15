package dto

import (
	"time"
)

type Pool_version struct {
	Pool_id                string     `json:"pool_id"`
	Pool_source            string     `json:"pool_source"`
	Active_ind             *string    `json:"active_ind"`
	Business_associate_id  *string    `json:"business_associate_id"`
	Current_status_date    *time.Time `json:"current_status_date"`
	Discovery_date         *time.Time `json:"discovery_date"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Field_id               *string    `json:"field_id"`
	Pool_code              *string    `json:"pool_code"`
	Pool_name              *string    `json:"pool_name"`
	Pool_name_abbreviation *string    `json:"pool_name_abbreviation"`
	Pool_status            *string    `json:"pool_status"`
	Pool_type              *string    `json:"pool_type"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Strat_name_set_id      *string    `json:"strat_name_set_id"`
	Strat_unit_id          *string    `json:"strat_unit_id"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
