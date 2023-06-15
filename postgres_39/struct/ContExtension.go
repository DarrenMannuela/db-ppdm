package dto

import (
	"time"
)

type Cont_extension struct {
	Contract_id        string     `json:"contract_id"`
	Extension_id       string     `json:"extension_id"`
	Active_ind         *string    `json:"active_ind"`
	Description        *string    `json:"description"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Extension_type     *string    `json:"extension_type"`
	Issued_by          *string    `json:"issued_by"`
	Issued_date        *time.Time `json:"issued_date"`
	Land_right_id      *string    `json:"land_right_id"`
	Land_right_subtype *string    `json:"land_right_subtype"`
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
