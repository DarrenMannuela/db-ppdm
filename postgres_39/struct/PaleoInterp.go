package dto

import (
	"time"
)

type Paleo_interp struct {
	Paleo_summary_id           string     `json:"paleo_summary_id"`
	Detail_id                  string     `json:"detail_id"`
	Active_ind                 *string    `json:"active_ind"`
	Analysis_id                *string    `json:"analysis_id"`
	Anl_source                 *string    `json:"anl_source"`
	Base_md                    *float64   `json:"base_md"`
	Base_md_ouom               *string    `json:"base_md_ouom"`
	Climate_id                 *string    `json:"climate_id"`
	Description                *string    `json:"description"`
	Ecozone_confidence_id      *string    `json:"ecozone_confidence_id"`
	Ecozone_id                 *string    `json:"ecozone_id"`
	Ecozone_qualifier_id       *string    `json:"ecozone_qualifier_id"`
	Effective_date             *time.Time `json:"effective_date"`
	Expiry_date                *time.Time `json:"expiry_date"`
	From_rel_strat_name_set_id *string    `json:"from_rel_strat_name_set_id"`
	From_rel_strat_unit_id     *string    `json:"from_rel_strat_unit_id"`
	From_strat_name_set_id     *string    `json:"from_strat_name_set_id"`
	From_strat_unit_id         *string    `json:"from_strat_unit_id"`
	Interp_type                *string    `json:"interp_type"`
	Lithology_type             *string    `json:"lithology_type"`
	Lith_structure_id          *string    `json:"lith_structure_id"`
	Maturation_obs_no          *int       `json:"maturation_obs_no"`
	Paleo_confidence_id        *string    `json:"paleo_confidence_id"`
	Paleo_qualifier_id         *string    `json:"paleo_qualifier_id"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Preferred_ind              *string    `json:"preferred_ind"`
	Remark                     *string    `json:"remark"`
	Source                     *string    `json:"source"`
	Tai_color                  *string    `json:"tai_color"`
	Top_md                     *float64   `json:"top_md"`
	Top_md_ouom                *string    `json:"top_md_ouom"`
	To_rel_strat_name_set_id   *string    `json:"to_rel_strat_name_set_id"`
	To_rel_strat_unit_id       *string    `json:"to_rel_strat_unit_id"`
	To_strat_name_set_id       *string    `json:"to_strat_name_set_id"`
	To_strat_unit_id           *string    `json:"to_strat_unit_id"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
