package dto

type Cont_mktg_elect_subst struct {
	Contract_id                string   `json:"contract_id" default:"source"`
	Operating_procedure_id     string   `json:"operating_procedure_id" default:"source"`
	Substance_id               string   `json:"substance_id" default:"source"`
	Substance_obs_no           int      `json:"substance_obs_no" default:"1"`
	Active_ind                 *string  `json:"active_ind" default:""`
	Both_ind                   *string  `json:"both_ind" default:""`
	Effective_date             *string  `json:"effective_date" default:""`
	Election_cost              *float64 `json:"election_cost" default:""`
	Election_cost_qualifier    *string  `json:"election_cost_qualifier" default:""`
	Election_percent           *float64 `json:"election_percent" default:""`
	Election_percent_qualifier *string  `json:"election_percent_qualifier" default:""`
	Expiry_date                *string  `json:"expiry_date" default:""`
	Ppdm_guid                  *string  `json:"ppdm_guid" default:""`
	Remark                     *string  `json:"remark" default:""`
	Source                     *string  `json:"source" default:""`
	Substance_ouom             *string  `json:"substance_ouom" default:""`
	Substance_uom              *string  `json:"substance_uom" default:""`
	Row_changed_by             *string  `json:"row_changed_by" default:""`
	Row_changed_date           *string  `json:"row_changed_date" default:""`
	Row_created_by             *string  `json:"row_created_by" default:""`
	Row_created_date           *string  `json:"row_created_date" default:""`
	Row_effective_date         *string  `json:"row_effective_date" default:""`
	Row_expiry_date            *string  `json:"row_expiry_date" default:""`
	Row_quality                *string  `json:"row_quality" default:""`
}
