package dto

import (
	"time"
)

type Rm_phys_item_maint struct {
	Physical_item_id   string     `json:"physical_item_id"`
	Maint_id           string     `json:"maint_id"`
	Active_ind         *string    `json:"active_ind"`
	Actual_date        *time.Time `json:"actual_date"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Maintain_process   *string    `json:"maintain_process"`
	Maint_ba_id        *string    `json:"maint_ba_id"`
	Maint_by_name      *string    `json:"maint_by_name"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Scheduled_date     *time.Time `json:"scheduled_date"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
