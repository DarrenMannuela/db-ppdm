package dto

import (
	"time"
)

type Well_drill_report struct {
	Uwi                      string     `json:"uwi"`
	Report_id                string     `json:"report_id"`
	Active_ind               *string    `json:"active_ind"`
	Contractor_ba_id         *string    `json:"contractor_ba_id"`
	Effective_date           *time.Time `json:"effective_date"`
	End_date                 *time.Time `json:"end_date"`
	End_time                 *time.Time `json:"end_time"`
	End_timezone             *string    `json:"end_timezone"`
	Expiry_date              *time.Time `json:"expiry_date"`
	File_date                *time.Time `json:"file_date"`
	File_time                *time.Time `json:"file_time"`
	File_timezone            *string    `json:"file_timezone"`
	Period_count             *int       `json:"period_count"`
	Period_duration          *float64   `json:"period_duration"`
	Period_duration_uom      *string    `json:"period_duration_uom"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Record_count             *int       `json:"record_count"`
	Reference_num            *string    `json:"reference_num"`
	Remark                   *string    `json:"remark"`
	Reported_contractor_name *string    `json:"reported_contractor_name"`
	Reported_rig_num         *string    `json:"reported_rig_num"`
	Source                   *string    `json:"source"`
	Standard_digital_format  *string    `json:"standard_digital_format"`
	Start_date               *time.Time `json:"start_date"`
	Start_time               *time.Time `json:"start_time"`
	Start_timezone           *string    `json:"start_timezone"`
	Sw_application_id        *string    `json:"sw_application_id"`
	Vendor_digital_format    *string    `json:"vendor_digital_format"`
	Version_seq_no           *int       `json:"version_seq_no"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
