package dto

type Land_status struct {
	Land_right_subtype    string  `json:"land_right_subtype" default:"source"`
	Land_right_id         string  `json:"land_right_id" default:"source"`
	Status_type           string  `json:"status_type" default:"source"`
	Land_right_status     string  `json:"land_right_status" default:"source"`
	Status_seq_no         int     `json:"status_seq_no" default:"1"`
	Active_ind            *string `json:"active_ind" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	Effective_term        *string `json:"effective_term" default:""`
	Effective_term_ouom   *string `json:"effective_term_ouom" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
	Reason                *string `json:"reason" default:""`
	Remark                *string `json:"remark" default:""`
	Section_of_act        *string `json:"section_of_act" default:""`
	Section_of_act_date   *string `json:"section_of_act_date" default:""`
	Source                *string `json:"source" default:""`
	Undetermined_term_ind *string `json:"undetermined_term_ind" default:""`
	Row_changed_by        *string `json:"row_changed_by" default:""`
	Row_changed_date      *string `json:"row_changed_date" default:""`
	Row_created_by        *string `json:"row_created_by" default:""`
	Row_created_date      *string `json:"row_created_date" default:""`
	Row_effective_date    *string `json:"row_effective_date" default:""`
	Row_expiry_date       *string `json:"row_expiry_date" default:""`
	Row_quality           *string `json:"row_quality" default:""`
}
