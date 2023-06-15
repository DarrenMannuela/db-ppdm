package dto

import (
	"time"
)

type Restriction struct {
	Restriction_id      string     `json:"restriction_id"`
	Restriction_version int        `json:"restriction_version"`
	Active_ind          *string    `json:"active_ind"`
	Description         *string    `json:"description"`
	Effective_date      *time.Time `json:"effective_date"`
	End_date            *time.Time `json:"end_date"`
	End_date_event      *time.Time `json:"end_date_event"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Regulatory_act      *string    `json:"regulatory_act"`
	Remark              *string    `json:"remark"`
	Restriction_class   *string    `json:"restriction_class"`
	Restriction_name    *string    `json:"restriction_name"`
	Restriction_type    *string    `json:"restriction_type"`
	Rest_duration_type  *string    `json:"rest_duration_type"`
	Source              *string    `json:"source"`
	Source_document_id  *string    `json:"source_document_id"`
	Start_date          *time.Time `json:"start_date"`
	Start_date_event    *time.Time `json:"start_date_event"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
