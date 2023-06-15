package dto

type Notification struct {
	Notification_id       string  `json:"notification_id" default:"source"`
	Active_ind            *string `json:"active_ind" default:""`
	Business_associate_id *string `json:"business_associate_id" default:""`
	Contract_ind          *string `json:"contract_ind" default:""`
	Critical_date         *string `json:"critical_date" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	Land_right_id         *string `json:"land_right_id" default:""`
	Land_right_ind        *string `json:"land_right_ind" default:""`
	Land_right_subtype    *string `json:"land_right_subtype" default:""`
	Notification_type     *string `json:"notification_type" default:""`
	Obligation_id         *string `json:"obligation_id" default:""`
	Obligation_ind        *string `json:"obligation_ind" default:""`
	Obligation_seq_no     *int    `json:"obligation_seq_no" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
	Remark                *string `json:"remark" default:""`
	Served_ind            *string `json:"served_ind" default:""`
	Source                *string `json:"source" default:""`
	Row_changed_by        *string `json:"row_changed_by" default:""`
	Row_changed_date      *string `json:"row_changed_date" default:""`
	Row_created_by        *string `json:"row_created_by" default:""`
	Row_created_date      *string `json:"row_created_date" default:""`
	Row_effective_date    *string `json:"row_effective_date" default:""`
	Row_expiry_date       *string `json:"row_expiry_date" default:""`
	Row_quality           *string `json:"row_quality" default:""`
}
