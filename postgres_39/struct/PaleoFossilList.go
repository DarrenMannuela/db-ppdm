package dto

import (
	"time"
)

type Paleo_fossil_list struct {
	Paleo_summary_id       string     `json:"paleo_summary_id"`
	Fossil_detail_id       string     `json:"fossil_detail_id"`
	Fossil_id              string     `json:"fossil_id"`
	Abundance_count        *float64   `json:"abundance_count"`
	Abundance_qualifier_id *string    `json:"abundance_qualifier_id"`
	Active_ind             *string    `json:"active_ind"`
	Analysis_id            *string    `json:"analysis_id"`
	Anl_source             *string    `json:"anl_source"`
	Confidence_id          *string    `json:"confidence_id"`
	Depth                  *float64   `json:"depth"`
	Depth_ouom             *string    `json:"depth_ouom"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Fossil_color           *string    `json:"fossil_color"`
	Lith_sample_id         *string    `json:"lith_sample_id"`
	Maturation_obs_no      *int       `json:"maturation_obs_no"`
	Ontogeny_type          *string    `json:"ontogeny_type"`
	Paleo_analyst_ba_id    *string    `json:"paleo_analyst_ba_id"`
	Physical_item_id       *string    `json:"physical_item_id"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Preferred_ind          *string    `json:"preferred_ind"`
	Preservation_quality   *string    `json:"preservation_quality"`
	Preservation_type      *string    `json:"preservation_type"`
	Remark                 *string    `json:"remark"`
	Slide_loc_x            *string    `json:"slide_loc_x"`
	Slide_loc_y            *string    `json:"slide_loc_y"`
	Slide_num              *string    `json:"slide_num"`
	Source                 *string    `json:"source"`
	Tai_color              *string    `json:"tai_color"`
	Type_fossil_type       *string    `json:"type_fossil_type"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
