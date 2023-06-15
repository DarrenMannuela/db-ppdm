package dto

import (
	"time"
)

type Paleo_abund_scheme struct {
	Scheme_id          string     `json:"scheme_id"`
	Abund_qualifier_id string     `json:"abund_qualifier_id"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Max_count          *int       `json:"max_count"`
	Min_count          *int       `json:"min_count"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Scheme_owner_ba_id *string    `json:"scheme_owner_ba_id"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
