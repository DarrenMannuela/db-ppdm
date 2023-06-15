package dto

import (
	"time"
)

type Land_sale_bid struct {
	Jurisdiction           string     `json:"jurisdiction"`
	Land_sale_number       string     `json:"land_sale_number"`
	Land_sale_offering_id  string     `json:"land_sale_offering_id"`
	Land_offering_bid_id   string     `json:"land_offering_bid_id"`
	Active_ind             *string    `json:"active_ind"`
	Address_for_service    *string    `json:"address_for_service"`
	Address_obs_no         *int       `json:"address_obs_no"`
	Address_source         *string    `json:"address_source"`
	Bidder                 *string    `json:"bidder"`
	Bidder_type            *string    `json:"bidder_type"`
	Bid_status             *string    `json:"bid_status"`
	Bid_status_date        *time.Time `json:"bid_status_date"`
	Bid_submit_date        *time.Time `json:"bid_submit_date"`
	Bid_type               *string    `json:"bid_type"`
	Cash_bid_type          *string    `json:"cash_bid_type"`
	Confidential_ind       *string    `json:"confidential_ind"`
	Contingency_desc       *string    `json:"contingency_desc"`
	Contingency_ind        *string    `json:"contingency_ind"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Payment_instruction_id *string    `json:"payment_instruction_id"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Priority_order         *int       `json:"priority_order"`
	Remark                 *string    `json:"remark"`
	Response               *string    `json:"response"`
	Response_date          *time.Time `json:"response_date"`
	Source                 *string    `json:"source"`
	Successful_ind         *string    `json:"successful_ind"`
	Work_bid_ind           *string    `json:"work_bid_ind"`
	Work_bid_remark        *string    `json:"work_bid_remark"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
