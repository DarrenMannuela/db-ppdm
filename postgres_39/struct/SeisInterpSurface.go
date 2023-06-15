package dto

import (
	"time"
)

type Seis_interp_surface struct {
	Surface_id              string     `json:"surface_id"`
	Active_ind              *string    `json:"active_ind"`
	Conformity_relationship *string    `json:"conformity_relationship"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Ordinal_seq_no          *int       `json:"ordinal_seq_no"`
	Overturned_ind          *string    `json:"overturned_ind"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Remark                  *string    `json:"remark"`
	Repeat_strat_occur_no   *int       `json:"repeat_strat_occur_no"`
	Repeat_strat_type       *string    `json:"repeat_strat_type"`
	Source                  *string    `json:"source"`
	Strat_name_set_id       *string    `json:"strat_name_set_id"`
	Strat_unit_id           *string    `json:"strat_unit_id"`
	Surface_name            *string    `json:"surface_name"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
