package dto

type Oblig_allow_deduction struct {
	Obligation_id        string  `json:"obligation_id" default:"source"`
	Obligation_seq_no    int     `json:"obligation_seq_no" default:"1"`
	Allow_deduction_id   string  `json:"allow_deduction_id" default:"source"`
	Active_ind           *string `json:"active_ind" default:""`
	Allow_deduction_type *string `json:"allow_deduction_type" default:""`
	Allow_deduct_remark  *string `json:"allow_deduct_remark" default:""`
	Contract_id          *string `json:"contract_id" default:""`
	Effective_date       *string `json:"effective_date" default:""`
	Expiry_date          *string `json:"expiry_date" default:""`
	Ppdm_guid            *string `json:"ppdm_guid" default:""`
	Remark               *string `json:"remark" default:""`
	Source               *string `json:"source" default:""`
	Row_changed_by       *string `json:"row_changed_by" default:""`
	Row_changed_date     *string `json:"row_changed_date" default:""`
	Row_created_by       *string `json:"row_created_by" default:""`
	Row_created_date     *string `json:"row_created_date" default:""`
	Row_effective_date   *string `json:"row_effective_date" default:""`
	Row_expiry_date      *string `json:"row_expiry_date" default:""`
	Row_quality          *string `json:"row_quality" default:""`
}
