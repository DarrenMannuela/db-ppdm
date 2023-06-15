package dto

import (
	"time"
)

type Ppdm_property_set struct {
	Property_set_id    string     `json:"property_set_id"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Property_set_name  *string    `json:"property_set_name"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Use_system_id      *string    `json:"use_system_id"`
	Use_table_name     *string    `json:"use_table_name"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
