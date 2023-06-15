package dto

import (
	"time"
)

type Rm_phys_item_store struct {
	Store_id               string     `json:"store_id"`
	Physical_item_id       string     `json:"physical_item_id"`
	Store_seq_no           int        `json:"store_seq_no"`
	Active_ind             *string    `json:"active_ind"`
	Create_date            *time.Time `json:"create_date"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Physical_item_status   *string    `json:"physical_item_status"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Preferred_location_ind *string    `json:"preferred_location_ind"`
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
