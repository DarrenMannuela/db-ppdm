package dto

import (
	"time"
)

type Hse_incident_detail struct {
	Incident_id            string     `json:"incident_id"`
	Detail_obs_no          int        `json:"detail_obs_no"`
	Active_ind             *string    `json:"active_ind"`
	Detail_type            *string    `json:"detail_type"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Incident_date          *time.Time `json:"incident_date"`
	Incident_duration      *float64   `json:"incident_duration"`
	Incident_duration_ouom *string    `json:"incident_duration_ouom"`
	Incident_duration_uom  *string    `json:"incident_duration_uom"`
	Incident_set_id        *string    `json:"incident_set_id"`
	Incident_severity_id   *string    `json:"incident_severity_id"`
	Incident_type_id       *string    `json:"incident_type_id"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
