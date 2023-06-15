package dto

type Int_set_xref struct {
	Interest_set_id       string  `json:"interest_set_id" default:"source"`
	Interest_set_seq_no   int     `json:"interest_set_seq_no" default:"1"`
	Interest_set_id_2     string  `json:"interest_set_id_2" default:"source"`
	Interest_set_seq_no_2 int     `json:"interest_set_seq_no_2" default:"1"`
	Int_set_xref_obs_no   int     `json:"int_set_xref_obs_no" default:"1"`
	Active_ind            *string `json:"active_ind" default:""`
	Contract_id           *string `json:"contract_id" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	Int_set_xref_type     *string `json:"int_set_xref_type" default:""`
	Partner_ba_id         *string `json:"partner_ba_id" default:""`
	Partner_ba_id_2       *string `json:"partner_ba_id_2" default:""`
	Partner_obs_no        *int    `json:"partner_obs_no" default:""`
	Partner_obs_no_2      *int    `json:"partner_obs_no_2" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
	Provision_id          *string `json:"provision_id" default:""`
	Remark                *string `json:"remark" default:""`
	Source                *string `json:"source" default:""`
	Row_changed_by        *string `json:"row_changed_by" default:""`
	Row_changed_date      *string `json:"row_changed_date" default:""`
	Row_created_by        *string `json:"row_created_by" default:""`
	Row_created_date      *string `json:"row_created_date" default:""`
	Row_effective_date    *string `json:"row_effective_date" default:""`
	Row_expiry_date       *string `json:"row_expiry_date" default:""`
	Row_quality           *string `json:"row_quality" default:""`
}
