package dto

import (
	"time"
)

type Rm_image_comp struct {
	Physical_item_id   string     `json:"physical_item_id"`
	Log_section_id_1   string     `json:"log_section_id_1"`
	Log_section_id_2   string     `json:"log_section_id_2"`
	Active_ind         *string    `json:"active_ind"`
	Composite_seq_no   *int       `json:"composite_seq_no"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
