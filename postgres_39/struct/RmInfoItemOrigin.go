package dto

import (
	"time"
)

type Rm_info_item_origin struct {
	Info_item_subtype          string     `json:"info_item_subtype"`
	Information_item_id        string     `json:"information_item_id"`
	Origin_seq_no              int        `json:"origin_seq_no"`
	Active_ind                 *string    `json:"active_ind"`
	Effective_date             *time.Time `json:"effective_date"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Information_process        *string    `json:"information_process"`
	Parent_information_item_id *string    `json:"parent_information_item_id"`
	Parent_info_item_subtype   *string    `json:"parent_info_item_subtype"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Processor                  *string    `json:"processor"`
	Process_date               *time.Time `json:"process_date"`
	Remark                     *string    `json:"remark"`
	Source                     *string    `json:"source"`
	Version                    *string    `json:"version"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
