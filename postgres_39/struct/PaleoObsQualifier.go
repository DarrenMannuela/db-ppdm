package dto

import (
	"time"
)

type Paleo_obs_qualifier struct {
	Qualifier_id         string     `json:"qualifier_id"`
	Active_ind           *string    `json:"active_ind"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Qualifier_code       *string    `json:"qualifier_code"`
	Qualifier_long_name  *string    `json:"qualifier_long_name"`
	Qualifier_owner      *string    `json:"qualifier_owner"`
	Qualifier_short_name *string    `json:"qualifier_short_name"`
	Qualifier_type       *string    `json:"qualifier_type"`
	Remark               *string    `json:"remark"`
	Source               *string    `json:"source"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
