package dto

import (
	"time"
)

type Application struct {
	Application_id          string     `json:"application_id"`
	Active_ind              *string    `json:"active_ind"`
	Application_type        *string    `json:"application_type"`
	Contract_id             *string    `json:"contract_id"`
	Current_status          *string    `json:"current_status"`
	Decision                *string    `json:"decision"`
	Decision_date           *time.Time `json:"decision_date"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Extension_id            *string    `json:"extension_id"`
	Fees_desc               *string    `json:"fees_desc"`
	Fees_paid_ind           *string    `json:"fees_paid_ind"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Previous_application_id *string    `json:"previous_application_id"`
	Rate_schedule_id        *string    `json:"rate_schedule_id"`
	Received_date           *time.Time `json:"received_date"`
	Reference_num           *string    `json:"reference_num"`
	Remark                  *string    `json:"remark"`
	Resubmission_ind        *string    `json:"resubmission_ind"`
	Section_of_act          *string    `json:"section_of_act"`
	Section_of_act_date     *time.Time `json:"section_of_act_date"`
	Source                  *string    `json:"source"`
	Submission_complete_ind *string    `json:"submission_complete_ind"`
	Submission_desc         *string    `json:"submission_desc"`
	Submitted_by            *string    `json:"submitted_by"`
	Submitted_date          *time.Time `json:"submitted_date"`
	Submitted_to            *string    `json:"submitted_to"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
