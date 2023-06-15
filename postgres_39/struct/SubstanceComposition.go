package dto

import (
	"time"
)

type Substance_composition struct {
	Substance_id             string     `json:"substance_id"`
	Sub_substance_id         string     `json:"sub_substance_id"`
	Composition_obs_no       int        `json:"composition_obs_no"`
	Active_ind               *string    `json:"active_ind"`
	Defined_by_ba_id         *string    `json:"defined_by_ba_id"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Formula                  *string    `json:"formula"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Remark                   *string    `json:"remark"`
	Source                   *string    `json:"source"`
	Source_document_id       *string    `json:"source_document_id"`
	Substance_component_type *string    `json:"substance_component_type"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
