package dto

import (
	"time"
)

type Well_log_curve_proc struct {
	Uwi                string     `json:"uwi"`
	Curve_id           string     `json:"curve_id"`
	Process_obs_no     int        `json:"process_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Dictionary_id      *string    `json:"dictionary_id"`
	Dict_process_id    *string    `json:"dict_process_id"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Reported_mnemonic  *string    `json:"reported_mnemonic"`
	Source             *string    `json:"source"`
	Splice_obs_no      *int       `json:"splice_obs_no"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
