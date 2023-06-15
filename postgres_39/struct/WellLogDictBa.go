package dto

import (
	"time"
)

type Well_log_dict_ba struct {
	Dictionary_id      string     `json:"dictionary_id"`
	Use_obs_no         int        `json:"use_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Description        *string    `json:"description"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Preferred_name     *string    `json:"preferred_name"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Used_by_ba_id      *string    `json:"used_by_ba_id"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
