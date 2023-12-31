package dto

type Int_set_partner struct {
	Interest_set_id        string   `json:"interest_set_id" default:"source"`
	Interest_set_seq_no    int      `json:"interest_set_seq_no" default:"1"`
	Partner_ba_id          string   `json:"partner_ba_id" default:"source"`
	Partner_obs_no         int      `json:"partner_obs_no" default:"1"`
	Active_ind             *string  `json:"active_ind" default:""`
	Address_obs_no         *int     `json:"address_obs_no" default:""`
	Address_source         *string  `json:"address_source" default:""`
	Breach_desc            *string  `json:"breach_desc" default:""`
	Breach_ind             *string  `json:"breach_ind" default:""`
	Change_reason          *string  `json:"change_reason" default:""`
	Confidential_ind       *string  `json:"confidential_ind" default:""`
	Contract_id            *string  `json:"contract_id" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Escrow_desc            *string  `json:"escrow_desc" default:""`
	Escrow_ind             *string  `json:"escrow_ind" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Gross_fract_interest   *string  `json:"gross_fract_interest" default:""`
	Gross_percent_interest *float64 `json:"gross_percent_interest" default:""`
	Inactive_date          *string  `json:"inactive_date" default:""`
	Instrument_id          *string  `json:"instrument_id" default:""`
	Interest_set_role      *string  `json:"interest_set_role" default:""`
	Net_fract_interest     *string  `json:"net_fract_interest" default:""`
	Net_percent_interest   *float64 `json:"net_percent_interest" default:""`
	Penalty_ind            *string  `json:"penalty_ind" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Provision_id           *string  `json:"provision_id" default:""`
	Remark                 *string  `json:"remark" default:""`
	Right_to_earn_desc     *string  `json:"right_to_earn_desc" default:""`
	Right_to_earn_ind      *string  `json:"right_to_earn_ind" default:""`
	Source                 *string  `json:"source" default:""`
	Title_own_type         *string  `json:"title_own_type" default:""`
	Unit_operated_ind      *string  `json:"unit_operated_ind" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
