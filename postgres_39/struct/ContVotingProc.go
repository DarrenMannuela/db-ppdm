package dto

import (
	"time"
)

type Cont_voting_proc struct {
	Contract_id         string     `json:"contract_id"`
	Voting_procedure_id string     `json:"voting_procedure_id"`
	Active_ind          *string    `json:"active_ind"`
	Defeat_count        *int       `json:"defeat_count"`
	Defeat_percent      *float64   `json:"defeat_percent"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Interest_percent    *float64   `json:"interest_percent"`
	No_vote_response    *string    `json:"no_vote_response"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Quorum_count        *int       `json:"quorum_count"`
	Remark              *string    `json:"remark"`
	Response_period     *float64   `json:"response_period"`
	Response_period_uom *string    `json:"response_period_uom"`
	Source              *string    `json:"source"`
	Source_document_id  *string    `json:"source_document_id"`
	Vote_procedure_type *string    `json:"vote_procedure_type"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
