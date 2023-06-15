package dto

import (
	"time"
)

type Cont_account_proc struct {
	Contract_id          string     `json:"contract_id"`
	Account_proc_type    string     `json:"account_proc_type"`
	Account_proc_obs_no  int        `json:"account_proc_obs_no"`
	Active_ind           *string    `json:"active_ind"`
	Advance_percent      *float64   `json:"advance_percent"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Inventory_period     *float64   `json:"inventory_period"`
	Inventory_period_uom *string    `json:"inventory_period_uom"`
	Material_sale_limit  *float64   `json:"material_sale_limit"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Quorum_count         *int       `json:"quorum_count"`
	Quorum_percent       *float64   `json:"quorum_percent"`
	Remark               *string    `json:"remark"`
	Source               *string    `json:"source"`
	Source_document_id   *string    `json:"source_document_id"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
