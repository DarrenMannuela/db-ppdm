package dto

import (
	"time"
)

type Rm_circ_process struct {
	Circ_id                string     `json:"circ_id"`
	Data_circ_process_id   string     `json:"data_circ_process_id"`
	Active_ind             *string    `json:"active_ind"`
	Data_circ_process      *string    `json:"data_circ_process"`
	Data_circ_process_date *time.Time `json:"data_circ_process_date"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
