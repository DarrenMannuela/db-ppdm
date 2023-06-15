package dto

import (
	"time"
)

type Anl_valid_ba struct {
	Method_set_id      string     `json:"method_set_id"`
	Method_id          string     `json:"method_id"`
	Valid_ba_id        string     `json:"valid_ba_id"`
	Valid_ba_obs_no    int        `json:"valid_ba_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Ba_role_type       *string    `json:"ba_role_type"`
	Confidence_type    *string    `json:"confidence_type"`
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
