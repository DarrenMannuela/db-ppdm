package dto

type Ba_license struct {
	Licensee_ba_id        string   `json:"licensee_ba_id" default:"source"`
	License_id            string   `json:"license_id" default:"source"`
	Active_ind            *string  `json:"active_ind" default:""`
	Application_id        *string  `json:"application_id" default:""`
	Description           *string  `json:"description" default:""`
	Effective_date        *string  `json:"effective_date" default:""`
	Expiry_date           *string  `json:"expiry_date" default:""`
	Extended_date         *string  `json:"extended_date" default:""`
	Fees_paid_ind         *string  `json:"fees_paid_ind" default:""`
	Granted_by_ba_id      *string  `json:"granted_by_ba_id" default:""`
	Granted_by_contact_id *string  `json:"granted_by_contact_id" default:""`
	Granted_date          *string  `json:"granted_date" default:""`
	Granted_to_ba_id      *string  `json:"granted_to_ba_id" default:""`
	Granted_to_contact_id *string  `json:"granted_to_contact_id" default:""`
	License_num           *string  `json:"license_num" default:""`
	License_term          *float64 `json:"license_term" default:""`
	License_term_uom      *string  `json:"license_term_uom" default:""`
	License_type          *string  `json:"license_type" default:""`
	Ppdm_guid             *string  `json:"ppdm_guid" default:""`
	Rate_schedule_id      *string  `json:"rate_schedule_id" default:""`
	Rel_licensee_ba_id    *string  `json:"rel_licensee_ba_id" default:""`
	Rel_license_id        *string  `json:"rel_license_id" default:""`
	Remark                *string  `json:"remark" default:""`
	Secondary_term        *float64 `json:"secondary_term" default:""`
	Secondary_term_uom    *string  `json:"secondary_term_uom" default:""`
	Source                *string  `json:"source" default:""`
	Violation_ind         *string  `json:"violation_ind" default:""`
	Row_changed_by        *string  `json:"row_changed_by" default:""`
	Row_changed_date      *string  `json:"row_changed_date" default:""`
	Row_created_by        *string  `json:"row_created_by" default:""`
	Row_created_date      *string  `json:"row_created_date" default:""`
	Row_effective_date    *string  `json:"row_effective_date" default:""`
	Row_expiry_date       *string  `json:"row_expiry_date" default:""`
	Row_quality           *string  `json:"row_quality" default:""`
}
