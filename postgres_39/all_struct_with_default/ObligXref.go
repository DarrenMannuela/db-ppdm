package dto

type Oblig_xref struct {
	Obligation_id       string  `json:"obligation_id" default:"source"`
	Obligation_seq_no   int     `json:"obligation_seq_no" default:"1"`
	Obligation_id_2     string  `json:"obligation_id_2" default:"source"`
	Obligation_seq_no_2 int     `json:"obligation_seq_no_2" default:"1"`
	Xref_seq_no         int     `json:"xref_seq_no" default:"1"`
	Active_ind          *string `json:"active_ind" default:""`
	Effective_date      *string `json:"effective_date" default:""`
	Expiry_date         *string `json:"expiry_date" default:""`
	Oblig_xref_type     *string `json:"oblig_xref_type" default:""`
	Ppdm_guid           *string `json:"ppdm_guid" default:""`
	Remark              *string `json:"remark" default:""`
	Source              *string `json:"source" default:""`
	Row_changed_by      *string `json:"row_changed_by" default:""`
	Row_changed_date    *string `json:"row_changed_date" default:""`
	Row_created_by      *string `json:"row_created_by" default:""`
	Row_created_date    *string `json:"row_created_date" default:""`
	Row_effective_date  *string `json:"row_effective_date" default:""`
	Row_expiry_date     *string `json:"row_expiry_date" default:""`
	Row_quality         *string `json:"row_quality" default:""`
}
