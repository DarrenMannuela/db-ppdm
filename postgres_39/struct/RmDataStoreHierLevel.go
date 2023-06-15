package dto

import (
	"time"
)

type Rm_data_store_hier_level struct {
	Data_store_hier_id   string     `json:"data_store_hier_id"`
	Hier_level_id        string     `json:"hier_level_id"`
	Active_ind           *string    `json:"active_ind"`
	Data_store_type      *string    `json:"data_store_type"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Level_name           *string    `json:"level_name"`
	Level_seq_no         *int       `json:"level_seq_no"`
	Parent_hier_level_id *string    `json:"parent_hier_level_id"`
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
