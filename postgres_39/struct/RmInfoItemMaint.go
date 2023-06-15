package dto

import (
	"time"
)

type Rm_info_item_maint struct {
	Info_item_subtype   string     `json:"info_item_subtype"`
	Information_item_id string     `json:"information_item_id"`
	Maintain_id         string     `json:"maintain_id"`
	Active_ind          *string    `json:"active_ind"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Instruction         *string    `json:"instruction"`
	Maint_ba_id         *string    `json:"maint_ba_id"`
	Maint_complete_date *time.Time `json:"maint_complete_date"`
	Maint_due_date      *time.Time `json:"maint_due_date"`
	Maint_period        *float64   `json:"maint_period"`
	Maint_period_type   *string    `json:"maint_period_type"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Scheduled_ind       *string    `json:"scheduled_ind"`
	Source              *string    `json:"source"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
