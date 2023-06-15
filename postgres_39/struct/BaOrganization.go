package dto

import (
	"time"
)

type Ba_organization struct {
	Business_associate_id string     `json:"business_associate_id"`
	Organization_id       string     `json:"organization_id"`
	Organization_seq_no   int        `json:"organization_seq_no"`
	Active_ind            *string    `json:"active_ind"`
	Address_obs_no        *int       `json:"address_obs_no"`
	Address_source        *string    `json:"address_source"`
	Area_id               *string    `json:"area_id"`
	Area_type             *string    `json:"area_type"`
	Created_date          *time.Time `json:"created_date"`
	Description           *string    `json:"description"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Organization_name     *string    `json:"organization_name"`
	Organization_type     *string    `json:"organization_type"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Removed_date          *time.Time `json:"removed_date"`
	Source                *string    `json:"source"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
