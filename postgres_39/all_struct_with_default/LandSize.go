package dto

type Land_size struct {
	Land_right_subtype    string   `json:"land_right_subtype" default:"source"`
	Land_right_id         string   `json:"land_right_id" default:"source"`
	Size_type             string   `json:"size_type" default:"source"`
	Size_seq_no           int      `json:"size_seq_no" default:"1"`
	Active_ind            *string  `json:"active_ind" default:""`
	Business_associate_id *string  `json:"business_associate_id" default:""`
	Effective_date        *string  `json:"effective_date" default:""`
	Expiry_date           *string  `json:"expiry_date" default:""`
	Gross_size            *float64 `json:"gross_size" default:""`
	Interest_set_id       *string  `json:"interest_set_id" default:""`
	Interest_set_seq_no   *int     `json:"interest_set_seq_no" default:""`
	Net_size              *float64 `json:"net_size" default:""`
	Partner_obs_no        *int     `json:"partner_obs_no" default:""`
	Ppdm_guid             *string  `json:"ppdm_guid" default:""`
	Remark                *string  `json:"remark" default:""`
	Size_ouom             *string  `json:"size_ouom" default:""`
	Source                *string  `json:"source" default:""`
	Row_changed_by        *string  `json:"row_changed_by" default:""`
	Row_changed_date      *string  `json:"row_changed_date" default:""`
	Row_created_by        *string  `json:"row_created_by" default:""`
	Row_created_date      *string  `json:"row_created_date" default:""`
	Row_effective_date    *string  `json:"row_effective_date" default:""`
	Row_expiry_date       *string  `json:"row_expiry_date" default:""`
	Row_quality           *string  `json:"row_quality" default:""`
}
