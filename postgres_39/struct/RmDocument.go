package dto

import (
	"time"
)

type Rm_document struct {
	Info_item_subtype   string     `json:"info_item_subtype"`
	Information_item_id string     `json:"information_item_id"`
	Active_ind          *string    `json:"active_ind"`
	Document_status     *string    `json:"document_status"`
	Document_type       *string    `json:"document_type"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Source              *string    `json:"source"`
	Source_document_id  *string    `json:"source_document_id"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
