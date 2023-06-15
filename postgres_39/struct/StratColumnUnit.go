package dto

import (
	"time"
)

type Strat_column_unit struct {
	Strat_column_id         string     `json:"strat_column_id"`
	Strat_column_source     string     `json:"strat_column_source"`
	Strat_name_set_id       string     `json:"strat_name_set_id"`
	Strat_unit_id           string     `json:"strat_unit_id"`
	Interp_id               string     `json:"interp_id"`
	Active_ind              *string    `json:"active_ind"`
	Business_associate_id   *string    `json:"business_associate_id"`
	Certified_ind           *string    `json:"certified_ind"`
	Conformity_relationship *string    `json:"conformity_relationship"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Interpretation_version  *string    `json:"interpretation_version"`
	Lithology               *string    `json:"lithology"`
	Occurrence_type         *string    `json:"occurrence_type"`
	Ordinal_seq_no          *int       `json:"ordinal_seq_no"`
	Overturned_ind          *string    `json:"overturned_ind"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Preferred_ind           *string    `json:"preferred_ind"`
	Remark                  *string    `json:"remark"`
	Repeat_strat_occur_no   *int       `json:"repeat_strat_occur_no"`
	Repeat_strat_type       *string    `json:"repeat_strat_type"`
	Source                  *string    `json:"source"`
	Source_document_id      *string    `json:"source_document_id"`
	Sour_gas_ind            *string    `json:"sour_gas_ind"`
	Strat_interpret_method  *string    `json:"strat_interpret_method"`
	Strat_topo_relation_ind *string    `json:"strat_topo_relation_ind"`
	Sw_application_id       *string    `json:"sw_application_id"`
	Version_obs_no          *int       `json:"version_obs_no"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
