package dto

type Fossil_name_set struct {
	Fossil_name_set_id   string  `json:"fossil_name_set_id" default:"source"`
	Active_ind           *string `json:"active_ind" default:""`
	Area_id              *string `json:"area_id" default:""`
	Area_type            *string `json:"area_type" default:""`
	Certified_ind        *string `json:"certified_ind" default:""`
	Effective_date       *string `json:"effective_date" default:""`
	Expiry_date          *string `json:"expiry_date" default:""`
	Fossil_name_set_type *string `json:"fossil_name_set_type" default:""`
	Name_set_name        *string `json:"name_set_name" default:""`
	Owner_ba_id          *string `json:"owner_ba_id" default:""`
	Ppdm_guid            *string `json:"ppdm_guid" default:""`
	Preferred_ind        *string `json:"preferred_ind" default:""`
	Remark               *string `json:"remark" default:""`
	Source               *string `json:"source" default:""`
	Row_changed_by       *string `json:"row_changed_by" default:""`
	Row_changed_date     *string `json:"row_changed_date" default:""`
	Row_created_by       *string `json:"row_created_by" default:""`
	Row_created_date     *string `json:"row_created_date" default:""`
	Row_effective_date   *string `json:"row_effective_date" default:""`
	Row_expiry_date      *string `json:"row_expiry_date" default:""`
	Row_quality          *string `json:"row_quality" default:""`
}
