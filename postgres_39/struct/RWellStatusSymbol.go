package dto

import (
	"time"
)

type R_well_status_symbol struct {
	Plot_symbol        string     `json:"plot_symbol"`
	Plot_symbol_obs_no int        `json:"plot_symbol_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Qualifier_value    *string    `json:"qualifier_value"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Status             *string    `json:"status"`
	Status_facet_count *int       `json:"status_facet_count"`
	Status_qualifier   *string    `json:"status_qualifier"`
	Status_type        *string    `json:"status_type"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
