package dto

type Application struct {
	Application_id          string  `json:"application_id" default:"source"`
	Active_ind              *string `json:"active_ind" default:""`
	Application_type        *string `json:"application_type" default:""`
	Contract_id             *string `json:"contract_id" default:""`
	Current_status          *string `json:"current_status" default:""`
	Decision                *string `json:"decision" default:""`
	Decision_date           *string `json:"decision_date" default:""`
	Effective_date          *string `json:"effective_date" default:""`
	Expiry_date             *string `json:"expiry_date" default:""`
	Extension_id            *string `json:"extension_id" default:""`
	Fees_desc               *string `json:"fees_desc" default:""`
	Fees_paid_ind           *string `json:"fees_paid_ind" default:""`
	Ppdm_guid               *string `json:"ppdm_guid" default:""`
	Previous_application_id *string `json:"previous_application_id" default:""`
	Rate_schedule_id        *string `json:"rate_schedule_id" default:""`
	Received_date           *string `json:"received_date" default:""`
	Reference_num           *string `json:"reference_num" default:""`
	Remark                  *string `json:"remark" default:""`
	Resubmission_ind        *string `json:"resubmission_ind" default:""`
	Section_of_act          *string `json:"section_of_act" default:""`
	Section_of_act_date     *string `json:"section_of_act_date" default:""`
	Source                  *string `json:"source" default:""`
	Submission_complete_ind *string `json:"submission_complete_ind" default:""`
	Submission_desc         *string `json:"submission_desc" default:""`
	Submitted_by            *string `json:"submitted_by" default:""`
	Submitted_date          *string `json:"submitted_date" default:""`
	Submitted_to            *string `json:"submitted_to" default:""`
	Row_changed_by          *string `json:"row_changed_by" default:""`
	Row_changed_date        *string `json:"row_changed_date" default:""`
	Row_created_by          *string `json:"row_created_by" default:""`
	Row_created_date        *string `json:"row_created_date" default:""`
	Row_effective_date      *string `json:"row_effective_date" default:""`
	Row_expiry_date         *string `json:"row_expiry_date" default:""`
	Row_quality             *string `json:"row_quality" default:""`
}
