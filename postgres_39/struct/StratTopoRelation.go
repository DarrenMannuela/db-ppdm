package dto

import (
	"time"
)

type Strat_topo_relation struct {
	Strat_name_set_id_1   string     `json:"strat_name_set_id_1"`
	Strat_unit_id_1       string     `json:"strat_unit_id_1"`
	Strat_name_set_id_2   string     `json:"strat_name_set_id_2"`
	Strat_unit_id_2       string     `json:"strat_unit_id_2"`
	Source                string     `json:"source"`
	Topo_relation_id      string     `json:"topo_relation_id"`
	Active_ind            *string    `json:"active_ind"`
	Area_id               *string    `json:"area_id"`
	Area_type             *string    `json:"area_type"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Interp_id_1           *string    `json:"interp_id_1"`
	Interp_id_2           *string    `json:"interp_id_2"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Strat_column_id_1     *string    `json:"strat_column_id_1"`
	Strat_column_id_2     *string    `json:"strat_column_id_2"`
	Strat_column_source_1 *string    `json:"strat_column_source_1"`
	Strat_column_source_2 *string    `json:"strat_column_source_2"`
	Topo_relation_type    *string    `json:"topo_relation_type"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
