package dto

type Fossil struct {
	Fossil_id          string  `json:"fossil_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Climate_type       *string `json:"climate_type" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Life_habit         *string `json:"life_habit" default:""`
	Lower_ecozone_id   *string `json:"lower_ecozone_id" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Preferred_ind      *string `json:"preferred_ind" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Taxon_group        *string `json:"taxon_group" default:""`
	Taxon_leaf_id      *string `json:"taxon_leaf_id" default:""`
	Upper_ecozone_id   *string `json:"upper_ecozone_id" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
