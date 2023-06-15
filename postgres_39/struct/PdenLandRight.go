package dto

import (
	"time"
)

type Pden_land_right struct {
	Pden_subtype          string     `json:"pden_subtype"`
	Pden_id               string     `json:"pden_id"`
	Pden_source           string     `json:"pden_source"`
	Active_ind            *string    `json:"active_ind"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Facility_id           *string    `json:"facility_id"`
	Facility_type         *string    `json:"facility_type"`
	Land_right_id         *string    `json:"land_right_id"`
	Land_right_subtype    *string    `json:"land_right_subtype"`
	No_of_gas_wells       *float64   `json:"no_of_gas_wells"`
	No_of_injection_wells *float64   `json:"no_of_injection_wells"`
	No_of_oil_wells       *float64   `json:"no_of_oil_wells"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
