package dto

import (
	"time"
)

type Ba_contact_info struct {
	Business_associate_id  string     `json:"business_associate_id"`
	Location_id            string     `json:"location_id"`
	Active_ind             *string    `json:"active_ind"`
	Address_obs_no         *int       `json:"address_obs_no"`
	Address_source         *string    `json:"address_source"`
	Ba_organization_id     *string    `json:"ba_organization_id"`
	Ba_organization_seq_no *int       `json:"ba_organization_seq_no"`
	Contact_loc_type       *string    `json:"contact_loc_type"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Location_name          *string    `json:"location_name"`
	Order_level            *int       `json:"order_level"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Preferred_ind          *string    `json:"preferred_ind"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
