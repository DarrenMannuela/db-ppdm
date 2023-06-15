package dto

import (
	"time"
)

type Well_log_parm struct {
	Uwi                  string     `json:"uwi"`
	Well_log_id          string     `json:"well_log_id"`
	Well_log_source      string     `json:"well_log_source"`
	Parameter_seq_no     int        `json:"parameter_seq_no"`
	Active_ind           *string    `json:"active_ind"`
	Array_ind            *string    `json:"array_ind"`
	Description          *string    `json:"description"`
	Dictionary_id        *string    `json:"dictionary_id"`
	Dict_parameter_id    *string    `json:"dict_parameter_id"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Parameter_output     *string    `json:"parameter_output"`
	Parameter_text_value *string    `json:"parameter_text_value"`
	Parameter_value      *float64   `json:"parameter_value"`
	Parameter_value_ouom *string    `json:"parameter_value_ouom"`
	Parameter_value_uom  *string    `json:"parameter_value_uom"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Reported_desc        *string    `json:"reported_desc"`
	Reported_mnemonic    *string    `json:"reported_mnemonic"`
	Source               *string    `json:"source"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
