package dto

type Paleo_fossil_list struct {
	Paleo_summary_id       string   `json:"paleo_summary_id" default:"source"`
	Fossil_detail_id       string   `json:"fossil_detail_id" default:"source"`
	Fossil_id              string   `json:"fossil_id" default:"source"`
	Abundance_count        *float64 `json:"abundance_count" default:""`
	Abundance_qualifier_id *string  `json:"abundance_qualifier_id" default:""`
	Active_ind             *string  `json:"active_ind" default:""`
	Analysis_id            *string  `json:"analysis_id" default:""`
	Anl_source             *string  `json:"anl_source" default:""`
	Confidence_id          *string  `json:"confidence_id" default:""`
	Depth                  *float64 `json:"depth" default:""`
	Depth_ouom             *string  `json:"depth_ouom" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Fossil_color           *string  `json:"fossil_color" default:""`
	Lith_sample_id         *string  `json:"lith_sample_id" default:""`
	Maturation_obs_no      *int     `json:"maturation_obs_no" default:""`
	Ontogeny_type          *string  `json:"ontogeny_type" default:""`
	Paleo_analyst_ba_id    *string  `json:"paleo_analyst_ba_id" default:""`
	Physical_item_id       *string  `json:"physical_item_id" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Preferred_ind          *string  `json:"preferred_ind" default:""`
	Preservation_quality   *string  `json:"preservation_quality" default:""`
	Preservation_type      *string  `json:"preservation_type" default:""`
	Remark                 *string  `json:"remark" default:""`
	Slide_loc_x            *string  `json:"slide_loc_x" default:""`
	Slide_loc_y            *string  `json:"slide_loc_y" default:""`
	Slide_num              *string  `json:"slide_num" default:""`
	Source                 *string  `json:"source" default:""`
	Tai_color              *string  `json:"tai_color" default:""`
	Type_fossil_type       *string  `json:"type_fossil_type" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
