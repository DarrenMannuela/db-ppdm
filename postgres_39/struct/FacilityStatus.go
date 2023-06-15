package dto

import (
	"time"
)

type Facility_status struct {
	Facility_id        string     `json:"facility_id"`
	Facility_type      string     `json:"facility_type"`
	Status_id          string     `json:"status_id"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	End_time           *time.Time `json:"end_time"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Percent_capability *float64   `json:"percent_capability"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Start_time         *time.Time `json:"start_time"`
	Status             *string    `json:"status"`
	Status_type        *string    `json:"status_type"`
	Timezone           *string    `json:"timezone"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
