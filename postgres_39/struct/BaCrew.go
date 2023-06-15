package dto

import (
	"time"
)

type Ba_crew struct {
	Crew_company_ba_id             string     `json:"crew_company_ba_id"`
	Crew_id                        string     `json:"crew_id"`
	Active_ind                     *string    `json:"active_ind"`
	Cost_per_unit                  *float64   `json:"cost_per_unit"`
	Cost_per_unit_currency_uom_uom *string    `json:"cost_per_unit_currency_uom_uom"`
	Cost_per_unit_uom              *string    `json:"cost_per_unit_uom"`
	Crew_abbreviation              *string    `json:"crew_abbreviation"`
	Crew_long_name                 *string    `json:"crew_long_name"`
	Crew_short_name                *string    `json:"crew_short_name"`
	Crew_type                      *string    `json:"crew_type"`
	Effective_date                 *time.Time `json:"effective_date"`
	Expiry_date                    *time.Time `json:"expiry_date"`
	Ppdm_guid                      string     `json:"ppdm_guid"`
	Remark                         *string    `json:"remark"`
	Source                         *string    `json:"source"`
	Row_changed_by                 *string    `json:"row_changed_by"`
	Row_changed_date               *time.Time `json:"row_changed_date"`
	Row_created_by                 *string    `json:"row_created_by"`
	Row_created_date               *time.Time `json:"row_created_date"`
	Row_effective_date             *time.Time `json:"row_effective_date"`
	Row_expiry_date                *time.Time `json:"row_expiry_date"`
	Row_quality                    *string    `json:"row_quality"`
}
