package dto

import (
	"time"
)

type Reserve_entity struct {
	Resent_id            string     `json:"resent_id"`
	Active_ind           *string    `json:"active_ind"`
	Created_by_ba_id     *string    `json:"created_by_ba_id"`
	Description          *string    `json:"description"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Group_type           *string    `json:"group_type"`
	Last_approve_ba_id   *string    `json:"last_approve_ba_id"`
	Last_update_ba_id    *string    `json:"last_update_ba_id"`
	Last_update_date     *time.Time `json:"last_update_date"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Primary_product_type *string    `json:"primary_product_type"`
	Remark               *string    `json:"remark"`
	Report_ind           *string    `json:"report_ind"`
	Resent_long_name     *string    `json:"resent_long_name"`
	Resent_short_name    *string    `json:"resent_short_name"`
	Source               *string    `json:"source"`
	Update_schedule      *string    `json:"update_schedule"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
