package dto

import (
	"time"
)

type Hse_incident struct {
	Incident_id            string     `json:"incident_id"`
	Active_ind             *string    `json:"active_ind"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Incident_class_id      *string    `json:"incident_class_id"`
	Incident_date          *time.Time `json:"incident_date"`
	Incident_duration      *float64   `json:"incident_duration"`
	Incident_duration_ouom *string    `json:"incident_duration_ouom"`
	Incident_duration_uom  *string    `json:"incident_duration_uom"`
	Lost_time_ind          *string    `json:"lost_time_ind"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Recorded_time          *time.Time `json:"recorded_time"`
	Recorded_timezone      *string    `json:"recorded_timezone"`
	Remark                 *string    `json:"remark"`
	Reported_by_ba_id      *string    `json:"reported_by_ba_id"`
	Reported_by_name       *string    `json:"reported_by_name"`
	Reported_ind           *string    `json:"reported_ind"`
	Source                 *string    `json:"source"`
	Work_related_ind       *string    `json:"work_related_ind"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
