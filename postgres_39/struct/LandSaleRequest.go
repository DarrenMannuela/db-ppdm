package dto

import (
	"time"
)

type Land_sale_request struct {
	Land_request_id       string     `json:"land_request_id"`
	Active_ind            *string    `json:"active_ind"`
	Advance_booking_ind   *string    `json:"advance_booking_ind"`
	Ami_ind               *string    `json:"ami_ind"`
	Area_id               *string    `json:"area_id"`
	Area_type             *string    `json:"area_type"`
	Bid_ind               *string    `json:"bid_ind"`
	Broker_ba_id          *string    `json:"broker_ba_id"`
	Contact_ba_id         *string    `json:"contact_ba_id"`
	Critical_date         *time.Time `json:"critical_date"`
	Drilling_priority_ind *string    `json:"drilling_priority_ind"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Granted_right_type    *string    `json:"granted_right_type"`
	Gross_size            *float64   `json:"gross_size"`
	Gross_size_ouom       *string    `json:"gross_size_ouom"`
	Investigator          *string    `json:"investigator"`
	Jurisdiction          *string    `json:"jurisdiction"`
	Land_req_status       *string    `json:"land_req_status"`
	Land_right_id         *string    `json:"land_right_id"`
	Land_right_subtype    *string    `json:"land_right_subtype"`
	Land_sale_number      *string    `json:"land_sale_number"`
	Land_sale_offering_id *string    `json:"land_sale_offering_id"`
	Lessor_refused_date   *time.Time `json:"lessor_refused_date"`
	Originator_ba_id      *string    `json:"originator_ba_id"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Previous_request_id   *string    `json:"previous_request_id"`
	Purpose               *string    `json:"purpose"`
	Reference_num         *string    `json:"reference_num"`
	Remark                *string    `json:"remark"`
	Requested_sale_date   *time.Time `json:"requested_sale_date"`
	Requested_sale_name   *string    `json:"requested_sale_name"`
	Requester             *string    `json:"requester"`
	Request_cancel_date   *time.Time `json:"request_cancel_date"`
	Request_description   *string    `json:"request_description"`
	Request_received_date *time.Time `json:"request_received_date"`
	Request_size          *float64   `json:"request_size"`
	Request_size_ouom     *string    `json:"request_size_ouom"`
	Request_type          *string    `json:"request_type"`
	Response_date         *time.Time `json:"response_date"`
	Response_desc         *string    `json:"response_desc"`
	Source                *string    `json:"source"`
	Submitted_date        *time.Time `json:"submitted_date"`
	Term_duration         *float64   `json:"term_duration"`
	Term_duration_ouom    *string    `json:"term_duration_ouom"`
	Uwi                   *string    `json:"uwi"`
	Well_license_id       *string    `json:"well_license_id"`
	Well_source           *string    `json:"well_source"`
	Withdrawn_date        *time.Time `json:"withdrawn_date"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
