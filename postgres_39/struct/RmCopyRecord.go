package dto

import (
	"time"
)

type Rm_copy_record struct {
	Old_physical_item_id string     `json:"old_physical_item_id"`
	New_physical_item_id string     `json:"new_physical_item_id"`
	Copy_seq_no          int        `json:"copy_seq_no"`
	Active_ind           *string    `json:"active_ind"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	New_record_no        *int       `json:"new_record_no"`
	Old_record_no        *int       `json:"old_record_no"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Remark               *string    `json:"remark"`
	Source               *string    `json:"source"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
