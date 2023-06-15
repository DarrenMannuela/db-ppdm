package dto

type Oblig_calc struct {
	Obligation_id       string  `json:"obligation_id" default:"source"`
	Obligation_seq_no   int     `json:"obligation_seq_no" default:"1"`
	Calculation_obs_no  int     `json:"calculation_obs_no" default:"1"`
	Active_ind          *string `json:"active_ind" default:""`
	Calculation_formula *string `json:"calculation_formula" default:""`
	Calculation_method  *string `json:"calculation_method" default:""`
	Calculation_point   *string `json:"calculation_point" default:""`
	Effective_date      *string `json:"effective_date" default:""`
	Expiry_date         *string `json:"expiry_date" default:""`
	Ppdm_guid           *string `json:"ppdm_guid" default:""`
	Remark              *string `json:"remark" default:""`
	Source              *string `json:"source" default:""`
	Substance_id        *string `json:"substance_id" default:""`
	Row_changed_by      *string `json:"row_changed_by" default:""`
	Row_changed_date    *string `json:"row_changed_date" default:""`
	Row_created_by      *string `json:"row_created_by" default:""`
	Row_created_date    *string `json:"row_created_date" default:""`
	Row_effective_date  *string `json:"row_effective_date" default:""`
	Row_expiry_date     *string `json:"row_expiry_date" default:""`
	Row_quality         *string `json:"row_quality" default:""`
}
