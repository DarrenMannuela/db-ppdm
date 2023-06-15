package dto

import (
	"time"
)

type Facility_license struct {
	Facility_id             string     `json:"facility_id"`
	Facility_type           string     `json:"facility_type"`
	License_id              string     `json:"license_id"`
	Active_ind              *string    `json:"active_ind"`
	Adjust_13_month_desc    *string    `json:"adjust_13_month_desc"`
	Adjust_13_month_ind     *string    `json:"adjust_13_month_ind"`
	Application_id          *string    `json:"application_id"`
	Approved_facility_class *string    `json:"approved_facility_class"`
	Assigned_field_id       *string    `json:"assigned_field_id"`
	Description             *string    `json:"description"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Facility_license_type   *string    `json:"facility_license_type"`
	Fees_paid_ind           *string    `json:"fees_paid_ind"`
	Granted_by_ba_id        *string    `json:"granted_by_ba_id"`
	Granted_by_contact_id   *string    `json:"granted_by_contact_id"`
	Granted_date            *time.Time `json:"granted_date"`
	Granted_to_ba_id        *string    `json:"granted_to_ba_id"`
	Granted_to_contact_id   *string    `json:"granted_to_contact_id"`
	License_extension_cond  *string    `json:"license_extension_cond"`
	License_latitude        *float64   `json:"license_latitude"`
	License_location        *string    `json:"license_location"`
	License_longitude       *float64   `json:"license_longitude"`
	License_num             *string    `json:"license_num"`
	License_term            *float64   `json:"license_term"`
	License_term_uom        *string    `json:"license_term_uom"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Rate_schedule_id        *string    `json:"rate_schedule_id"`
	Rel_license_id          *string    `json:"rel_license_id"`
	Remark                  *string    `json:"remark"`
	Renewal_condition       *string    `json:"renewal_condition"`
	Secondary_term          *float64   `json:"secondary_term"`
	Secondary_term_uom      *string    `json:"secondary_term_uom"`
	Source                  *string    `json:"source"`
	Violation_ind           *string    `json:"violation_ind"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
