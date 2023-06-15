package dto

import (
	"time"
)

type Sf_xref struct {
	Sf_subtype           string     `json:"sf_subtype"`
	Support_facility_id  string     `json:"support_facility_id"`
	Sf_subtype2          string     `json:"sf_subtype_2"`
	Support_facility_id2 string     `json:"support_facility_id_2"`
	Active_ind           *string    `json:"active_ind"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	From_distance        *float64   `json:"from_distance"`
	From_distance_ouom   *string    `json:"from_distance_ouom"`
	Portion_percent      *float64   `json:"portion_percent"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Source               *string    `json:"source"`
	To_distance          *float64   `json:"to_distance"`
	To_distance_ouom     *string    `json:"to_distance_ouom"`
	Xref_type            *string    `json:"xref_type"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
