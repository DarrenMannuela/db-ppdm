package dto

import (
	"time"
)

type Rm_file_content struct {
	File_id             string     `json:"file_id"`
	Active_ind          *string    `json:"active_ind"`
	Digital_format      *string    `json:"digital_format"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	File_content        *string    `json:"file_content"`
	File_size           *float64   `json:"file_size"`
	File_size_uom       *string    `json:"file_size_uom"`
	Information_item_id *string    `json:"information_item_id"`
	Info_item_subtype   *string    `json:"info_item_subtype"`
	Physical_item_id    *string    `json:"physical_item_id"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Source              *string    `json:"source"`
	Sw_application_id   *string    `json:"sw_application_id"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
