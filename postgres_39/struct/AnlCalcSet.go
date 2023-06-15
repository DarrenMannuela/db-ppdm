package dto

import (
	"time"
)

type Anl_calc_set struct {
	Calc_set_id             string     `json:"calc_set_id"`
	Active_ind              *string    `json:"active_ind"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Method_id               *string    `json:"method_id"`
	Method_set_id           *string    `json:"method_set_id"`
	Owner_ba_id             *string    `json:"owner_ba_id"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Preferred_name          *string    `json:"preferred_name"`
	Remark                  *string    `json:"remark"`
	Software_application_id *string    `json:"software_application_id"`
	Source                  *string    `json:"source"`
	Source_document_id      *string    `json:"source_document_id"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
