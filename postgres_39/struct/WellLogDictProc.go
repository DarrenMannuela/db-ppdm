package dto

import (
	"time"
)

type Well_log_dict_proc struct {
	Dictionary_id      string     `json:"dictionary_id"`
	Process_id         string     `json:"process_id"`
	Active_ind         *string    `json:"active_ind"`
	Description        *string    `json:"description"`
	Dic_proc_name      *string    `json:"dic_proc_name"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Mnemonic           *string    `json:"mnemonic"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
