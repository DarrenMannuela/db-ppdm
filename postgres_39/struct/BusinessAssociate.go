package dto

import (
	"time"
)

type Business_associate struct {
	Business_associate_id string     `json:"business_associate_id"`
	Active_ind            *string    `json:"active_ind"`
	Ba_abbreviation       *string    `json:"ba_abbreviation"`
	Ba_category           *string    `json:"ba_category"`
	Ba_code               *string    `json:"ba_code"`
	Ba_long_name          *string    `json:"ba_long_name"`
	Ba_short_name         *string    `json:"ba_short_name"`
	Ba_type               *string    `json:"ba_type"`
	Credit_check_date     *time.Time `json:"credit_check_date"`
	Credit_check_ind      *string    `json:"credit_check_ind"`
	Credit_check_source   *string    `json:"credit_check_source"`
	Credit_rating         *string    `json:"credit_rating"`
	Credit_rating_source  *string    `json:"credit_rating_source"`
	Current_status        *string    `json:"current_status"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	First_name            *string    `json:"first_name"`
	Last_name             *string    `json:"last_name"`
	Middle_initial        *string    `json:"middle_initial"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Service_period        *string    `json:"service_period"`
	Source                *string    `json:"source"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
