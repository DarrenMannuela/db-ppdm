package dto

import (
	"time"
)

type Land_sale_bid_set struct {
	Land_sale_bid_set_id string     `json:"land_sale_bid_set_id"`
	Active_ind           *string    `json:"active_ind"`
	Bid_status           *string    `json:"bid_status"`
	Confidential_ind     *string    `json:"confidential_ind"`
	Contingency_desc     *string    `json:"contingency_desc"`
	Contingency_ind      *string    `json:"contingency_ind"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Owner_ba_id          *string    `json:"owner_ba_id"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Set_status_date      *time.Time `json:"set_status_date"`
	Source               *string    `json:"source"`
	Successful_ind       *string    `json:"successful_ind"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
