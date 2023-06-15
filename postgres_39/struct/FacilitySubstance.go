package dto

import (
	"time"
)

type Facility_substance struct {
	Facility_id            string     `json:"facility_id"`
	Facility_type          string     `json:"facility_type"`
	Substance_id           string     `json:"substance_id"`
	Active_ind             *string    `json:"active_ind"`
	Average_capacity       *float64   `json:"average_capacity"`
	Capacity_ouom          *string    `json:"capacity_ouom"`
	Capacity_uom           *string    `json:"capacity_uom"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Max_capacity           *float64   `json:"max_capacity"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Strat_name_set_id      *string    `json:"strat_name_set_id"`
	Strat_unit_id          *string    `json:"strat_unit_id"`
	Substance_excluded_ind *string    `json:"substance_excluded_ind"`
	Substance_included_ind *string    `json:"substance_included_ind"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
