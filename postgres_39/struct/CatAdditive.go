package dto

import (
	"time"
)

type Cat_additive struct {
	Catalogue_additive_id  string     `json:"catalogue_additive_id"`
	Active_ind             *string    `json:"active_ind"`
	Additive_group         *string    `json:"additive_group"`
	Additive_name          *string    `json:"additive_name"`
	Additive_type          *string    `json:"additive_type"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Manufacturer           *string    `json:"manufacturer"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Purchase_quantity      *float64   `json:"purchase_quantity"`
	Purchase_quantity_type *string    `json:"purchase_quantity_type"`
	Purchase_quantity_uom  *string    `json:"purchase_quantity_uom"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Supplier               *string    `json:"supplier"`
	Upc_code               *string    `json:"upc_code"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
