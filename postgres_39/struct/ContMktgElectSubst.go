package dto

import (
	"time"
)

type Cont_mktg_elect_subst struct {
	Contract_id                string     `json:"contract_id"`
	Operating_procedure_id     string     `json:"operating_procedure_id"`
	Substance_id               string     `json:"substance_id"`
	Substance_obs_no           int        `json:"substance_obs_no"`
	Active_ind                 *string    `json:"active_ind"`
	Both_ind                   *string    `json:"both_ind"`
	Effective_date             *time.Time `json:"effective_date"`
	Election_cost              *float64   `json:"election_cost"`
	Election_cost_qualifier    *string    `json:"election_cost_qualifier"`
	Election_percent           *float64   `json:"election_percent"`
	Election_percent_qualifier *string    `json:"election_percent_qualifier"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Remark                     *string    `json:"remark"`
	Source                     *string    `json:"source"`
	Substance_ouom             *string    `json:"substance_ouom"`
	Substance_uom              *string    `json:"substance_uom"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
