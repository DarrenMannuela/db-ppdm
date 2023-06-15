package dto

type Int_set_status struct {
	Interest_set_id       string  `json:"interest_set_id" default:"source"`
	Interest_set_seq_no   int     `json:"interest_set_seq_no" default:"1"`
	Status_obs_no         int     `json:"status_obs_no" default:"1"`
	Active_ind            *string `json:"active_ind" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	Effective_term        *string `json:"effective_term" default:""`
	Effective_term_ouom   *string `json:"effective_term_ouom" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
	Reason                *string `json:"reason" default:""`
	Remark                *string `json:"remark" default:""`
	Source                *string `json:"source" default:""`
	Status                *string `json:"status" default:""`
	Status_type           *string `json:"status_type" default:""`
	Undetermined_term_ind *string `json:"undetermined_term_ind" default:""`
	Row_changed_by        *string `json:"row_changed_by" default:""`
	Row_changed_date      *string `json:"row_changed_date" default:""`
	Row_created_by        *string `json:"row_created_by" default:""`
	Row_created_date      *string `json:"row_created_date" default:""`
	Row_effective_date    *string `json:"row_effective_date" default:""`
	Row_expiry_date       *string `json:"row_expiry_date" default:""`
	Row_quality           *string `json:"row_quality" default:""`
}
