package dto

import (
	"time"
)

type Sp_zone_definition struct {
	Zone_definition_id     string     `json:"zone_definition_id"`
	Active_ind             *string    `json:"active_ind"`
	Base_depth             *float64   `json:"base_depth"`
	Base_depth_ouom        *string    `json:"base_depth_ouom"`
	Base_qualifier         *string    `json:"base_qualifier"`
	Base_strat_unit_id     *string    `json:"base_strat_unit_id"`
	Defined_strat_unit_id  *string    `json:"defined_strat_unit_id"`
	Description            *string    `json:"description"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Owner_ba_id            *string    `json:"owner_ba_id"`
	Owner_ref_num          *string    `json:"owner_ref_num"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Strat_name_set_id      *string    `json:"strat_name_set_id"`
	Top_depth              *float64   `json:"top_depth"`
	Top_depth_ouom         *string    `json:"top_depth_ouom"`
	Top_qualifier          *string    `json:"top_qualifier"`
	Top_strat_unit_id      *string    `json:"top_strat_unit_id"`
	Uwi                    *string    `json:"uwi"`
	Zone_derivation_method *string    `json:"zone_derivation_method"`
	Zone_type              *string    `json:"zone_type"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
