package dto

import (
	"time"
)

type Rm_encoding struct {
	Physical_item_id   string     `json:"physical_item_id"`
	Encoding_seq_no    int        `json:"encoding_seq_no"`
	Active_ind         *string    `json:"active_ind"`
	Decrypt_key_id     *string    `json:"decrypt_key_id"`
	Effective_date     *time.Time `json:"effective_date"`
	Encoding_type      *string    `json:"encoding_type"`
	Encoding_version   *string    `json:"encoding_version"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Sw_application_id  *string    `json:"sw_application_id"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
