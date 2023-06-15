package dto

import (
	"time"
)

type Land_termination struct {
	Lr_termination_id      string     `json:"lr_termination_id"`
	Lr_termination_seq_no  int        `json:"lr_termination_seq_no"`
	Active_ind             *string    `json:"active_ind"`
	Business_associate_id  *string    `json:"business_associate_id"`
	Description            *string    `json:"description"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Fulfilled_date         *time.Time `json:"fulfilled_date"`
	Fulfilled_user         *string    `json:"fulfilled_user"`
	Jurisdiction           *string    `json:"jurisdiction"`
	Land_request_id        *string    `json:"land_request_id"`
	Land_right_id          *string    `json:"land_right_id"`
	Land_right_subtype     *string    `json:"land_right_subtype"`
	Land_sale_number       *string    `json:"land_sale_number"`
	Land_sale_offering_id  *string    `json:"land_sale_offering_id"`
	Mineral_zone_id        *string    `json:"mineral_zone_id"`
	Notification_id        *string    `json:"notification_id"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Spatial_description_id *string    `json:"spatial_description_id"`
	Spatial_obs_no         *int       `json:"spatial_obs_no"`
	Termination_reqmt      *string    `json:"termination_reqmt"`
	Termination_type       *string    `json:"termination_type"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
