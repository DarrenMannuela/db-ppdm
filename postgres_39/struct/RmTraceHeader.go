package dto

import (
	"time"
)

type Rm_trace_header struct {
	Info_item_subtype   string     `json:"info_item_subtype"`
	Information_item_id string     `json:"information_item_id"`
	Header_id           string     `json:"header_id"`
	Active_ind          *string    `json:"active_ind"`
	Description         *string    `json:"description"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Header_format       *string    `json:"header_format"`
	Header_offset       *float64   `json:"header_offset"`
	Header_word         *string    `json:"header_word"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Source              *string    `json:"source"`
	Word_length         *float64   `json:"word_length"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
