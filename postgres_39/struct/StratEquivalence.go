package dto

import (
	"time"
)

type Strat_equivalence struct {
	Strat_name_set_id         string     `json:"strat_name_set_id"`
	Strat_unit_id             string     `json:"strat_unit_id"`
	Equiv_strat_name_set_id   string     `json:"equiv_strat_name_set_id"`
	Equiv_strat_unit_id       string     `json:"equiv_strat_unit_id"`
	Equiv_id                  string     `json:"equiv_id"`
	Source                    string     `json:"source"`
	Active_ind                *string    `json:"active_ind"`
	Area_id                   *string    `json:"area_id"`
	Area_id2                  *string    `json:"area_id_2"`
	Area_type                 *string    `json:"area_type"`
	Area_type2                *string    `json:"area_type_2"`
	Effective_date            *time.Time `json:"effective_date"`
	Equiv_direction           *string    `json:"equiv_direction"`
	Equiv_interp_id           *string    `json:"equiv_interp_id"`
	Equiv_strat_column_id     *string    `json:"equiv_strat_column_id"`
	Equiv_strat_column_source *string    `json:"equiv_strat_column_source"`
	Expiry_date               *time.Time `json:"expiry_date"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Remark                    *string    `json:"remark"`
	Source_document_id        *string    `json:"source_document_id"`
	Strat_column_id           *string    `json:"strat_column_id"`
	Strat_column_source       *string    `json:"strat_column_source"`
	Strat_equivalence_type    *string    `json:"strat_equivalence_type"`
	Strat_interp_id           *string    `json:"strat_interp_id"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
