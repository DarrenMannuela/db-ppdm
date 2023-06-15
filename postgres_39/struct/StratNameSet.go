package dto

import (
	"time"
)

type Strat_name_set struct {
	Strat_name_set_id     string     `json:"strat_name_set_id"`
	Active_ind            *string    `json:"active_ind"`
	Area_id               *string    `json:"area_id"`
	Area_type             *string    `json:"area_type"`
	Business_associate_id *string    `json:"business_associate_id"`
	Certified_ind         *string    `json:"certified_ind"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Preferred_ind         *string    `json:"preferred_ind"`
	Remark                *string    `json:"remark"`
	Source                *string    `json:"source"`
	Strat_name_set_name   *string    `json:"strat_name_set_name"`
	Strat_name_set_type   *string    `json:"strat_name_set_type"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
