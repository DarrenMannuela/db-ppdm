package dto

import (
	"time"
)

type Substance_ba struct {
	Substance_id           string     `json:"substance_id"`
	Anl_source             string     `json:"anl_source"`
	Ba_obs_no              int        `json:"ba_obs_no"`
	Active_ind             *string    `json:"active_ind"`
	Ba_role_type           *string    `json:"ba_role_type"`
	Cas_number             *string    `json:"cas_number"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Location               *string    `json:"location"`
	Manufacturer           *string    `json:"manufacturer"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Price                  *float64   `json:"price"`
	Price_ouom             *string    `json:"price_ouom"`
	Problem_ind            *string    `json:"problem_ind"`
	Purchase_quantity      *float64   `json:"purchase_quantity"`
	Purchase_quantity_ouom *string    `json:"purchase_quantity_ouom"`
	Purchase_quantity_type *string    `json:"purchase_quantity_type"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Step_seq_no            *int       `json:"step_seq_no"`
	Substance_ba_id        *string    `json:"substance_ba_id"`
	Supplier               *string    `json:"supplier"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
