package dto

import (
	"time"
)

type Work_order_delivery struct {
	Work_order_id              string     `json:"work_order_id"`
	Delivery_id                string     `json:"delivery_id"`
	Active_ind                 *string    `json:"active_ind"`
	Ba_obs_no                  *int       `json:"ba_obs_no"`
	Business_associate_id      *string    `json:"business_associate_id"`
	Delivery_ba_address_obs_no *int       `json:"delivery_ba_address_obs_no"`
	Delivery_ba_address_source *string    `json:"delivery_ba_address_source"`
	Delivery_ba_id             *string    `json:"delivery_ba_id"`
	Delivery_date              *time.Time `json:"delivery_date"`
	Delivery_type              *string    `json:"delivery_type"`
	Effective_date             *time.Time `json:"effective_date"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Remark                     *string    `json:"remark"`
	Source                     *string    `json:"source"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
