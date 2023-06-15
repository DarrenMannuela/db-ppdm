package dto

import (
	"time"
)

type Equipment_spec_set struct {
	Spec_set_id        string     `json:"spec_set_id"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Equipment_type     *string    `json:"equipment_type"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Owner_ba_id        *string    `json:"owner_ba_id"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Preferred_name     *string    `json:"preferred_name"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Spec_set_desc      *string    `json:"spec_set_desc"`
	Spec_set_type      *string    `json:"spec_set_type"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
