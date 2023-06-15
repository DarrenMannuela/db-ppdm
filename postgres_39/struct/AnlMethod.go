package dto

import (
	"time"
)

type Anl_method struct {
	Method_set_id      string     `json:"method_set_id"`
	Method_id          string     `json:"method_id"`
	Abbreviation       *string    `json:"abbreviation"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Long_name          *string    `json:"long_name"`
	Method_desc        *string    `json:"method_desc"`
	Method_seq_no      *int       `json:"method_seq_no"`
	Method_type        *string    `json:"method_type"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Preparation_class  *string    `json:"preparation_class"`
	Remark             *string    `json:"remark"`
	Short_name         *string    `json:"short_name"`
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
