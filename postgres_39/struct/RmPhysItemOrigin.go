package dto

import (
	"time"
)

type Rm_phys_item_origin struct {
	Physical_item_id        string     `json:"physical_item_id"`
	Origin_seq_no           int        `json:"origin_seq_no"`
	Active_ind              *string    `json:"active_ind"`
	Business_associate_id   *string    `json:"business_associate_id"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Parent_physical_item_id *string    `json:"parent_physical_item_id"`
	Physical_process        *string    `json:"physical_process"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Process_date            *time.Time `json:"process_date"`
	Remark                  *string    `json:"remark"`
	Source                  *string    `json:"source"`
	Version                 *string    `json:"version"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
