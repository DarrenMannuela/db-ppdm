package dto

import (
	"time"
)

type Cont_provision_xref struct {
	Contract_id         string     `json:"contract_id"`
	Provision_id        string     `json:"provision_id"`
	Contract_id_2       string     `json:"contract_id_2"`
	Provision_id_2      string     `json:"provision_id_2"`
	Active_ind          *string    `json:"active_ind"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Provision_xref_type *string    `json:"provision_xref_type"`
	Remark              *string    `json:"remark"`
	Source              *string    `json:"source"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
