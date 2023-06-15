package dto

import (
	"time"
)

type Hse_incident_equiv struct {
	Incident_set_id    string     `json:"incident_set_id"`
	Incident_type_id   string     `json:"incident_type_id"`
	Incident_set_id2   string     `json:"incident_set_id_2"`
	Incident_type_id2  string     `json:"incident_type_id_2"`
	Active_ind         *string    `json:"active_ind"`
	Description        *string    `json:"description"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Source_document_id *string    `json:"source_document_id"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
