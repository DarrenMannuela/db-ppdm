package dto

type Cont_voting_proc struct {
	Contract_id         string   `json:"contract_id" default:"source"`
	Voting_procedure_id string   `json:"voting_procedure_id" default:"source"`
	Active_ind          *string  `json:"active_ind" default:""`
	Defeat_count        *int     `json:"defeat_count" default:""`
	Defeat_percent      *float64 `json:"defeat_percent" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Interest_percent    *float64 `json:"interest_percent" default:""`
	No_vote_response    *string  `json:"no_vote_response" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Quorum_count        *int     `json:"quorum_count" default:""`
	Remark              *string  `json:"remark" default:""`
	Response_period     *float64 `json:"response_period" default:""`
	Response_period_uom *string  `json:"response_period_uom" default:""`
	Source              *string  `json:"source" default:""`
	Source_document_id  *string  `json:"source_document_id" default:""`
	Vote_procedure_type *string  `json:"vote_procedure_type" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}
