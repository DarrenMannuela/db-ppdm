package dto

import (
	"time"
)

type Rest_contact struct {
	Restriction_id        string     `json:"restriction_id"`
	Restriction_version   int        `json:"restriction_version"`
	Business_associate_id string     `json:"business_associate_id"`
	Contact_obs_no        int        `json:"contact_obs_no"`
	Active_ind            *string    `json:"active_ind"`
	Address_obs_no        *int       `json:"address_obs_no"`
	Address_source        *string    `json:"address_source"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Phone_num             *string    `json:"phone_num"`
	Phone_num_id          *string    `json:"phone_num_id"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Primary_contact_ind   *string    `json:"primary_contact_ind"`
	Remark                *string    `json:"remark"`
	Source                *string    `json:"source"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
