package dto

import (
	"time"
)

type R_well_status_xref struct {
	Status_xref_id     string     `json:"status_xref_id"`
	Status_xref_obs_no int        `json:"status_xref_obs_no"`
	Active_ind         *string    `json:"active_ind"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Mapping_count      *int       `json:"mapping_count"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Qualifier_value1   *string    `json:"qualifier_value_1"`
	Qualifier_value2   *string    `json:"qualifier_value_2"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Status1            *string    `json:"status_1"`
	Status2            *string    `json:"status_2"`
	Status_qualifier1  *string    `json:"status_qualifier_1"`
	Status_qualifier2  *string    `json:"status_qualifier_2"`
	Status_type1       *string    `json:"status_type_1"`
	Status_type2       *string    `json:"status_type_2"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
