package dto

import (
	"time"
)

type Prod_lease_unit struct {
	Lease_unit_id            string     `json:"lease_unit_id"`
	Acreage                  *float64   `json:"acreage"`
	Acreage_ouom             *string    `json:"acreage_ouom"`
	Active_ind               *string    `json:"active_ind"`
	Area_id                  *string    `json:"area_id"`
	Area_type                *string    `json:"area_type"`
	Current_operator         *string    `json:"current_operator"`
	Current_status_date      *time.Time `json:"current_status_date"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Field_id                 *string    `json:"field_id"`
	Government_lease_unit_id *string    `json:"government_lease_unit_id"`
	Land_right_id            *string    `json:"land_right_id"`
	Land_right_subtype       *string    `json:"land_right_subtype"`
	Lease_unit_long_name     *string    `json:"lease_unit_long_name"`
	Lease_unit_short_name    *string    `json:"lease_unit_short_name"`
	Lease_unit_status        *string    `json:"lease_unit_status"`
	Lease_unit_type          *string    `json:"lease_unit_type"`
	Pool_id                  *string    `json:"pool_id"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Remark                   *string    `json:"remark"`
	Source                   *string    `json:"source"`
	Strat_name_set_id        *string    `json:"strat_name_set_id"`
	Strat_unit_id            *string    `json:"strat_unit_id"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
