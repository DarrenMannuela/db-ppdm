package dto

import (
	"time"
)

type Seis_patch struct {
	Patch_id           string     `json:"patch_id"`
	Active_ind         *string    `json:"active_ind"`
	Channel_count      *int       `json:"channel_count"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Gap_count          *int       `json:"gap_count"`
	Patch_layout       *string    `json:"patch_layout"`
	Patch_type         *string    `json:"patch_type"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Roll_along_method  *string    `json:"roll_along_method"`
	Shot_gap_ind       *string    `json:"shot_gap_ind"`
	Source             *string    `json:"source"`
	Symmetric_ind      *string    `json:"symmetric_ind"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
