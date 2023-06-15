package dto

import (
	"time"
)

type Anl_method_set struct {
	Method_set_id      string     `json:"method_set_id"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Method_set_name    *string    `json:"method_set_name"`
	Method_set_type    *string    `json:"method_set_type"`
	Owner_ba_id        *string    `json:"owner_ba_id"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Preferred_ind      *string    `json:"preferred_ind"`
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
