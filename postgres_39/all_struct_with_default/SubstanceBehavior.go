package dto

type Substance_behavior struct {
	Substance_id         string  `json:"substance_id" default:"source"`
	Active_ind           *string `json:"active_ind" default:""`
	Composite_ind        *string `json:"composite_ind" default:""`
	Drilling_mud_ind     *string `json:"drilling_mud_ind" default:""`
	Effective_date       *string `json:"effective_date" default:""`
	Element_ind          *string `json:"element_ind" default:""`
	Expiry_date          *string `json:"expiry_date" default:""`
	Hydrocarbon_ind      *string `json:"hydrocarbon_ind" default:""`
	Ion_ind              *string `json:"ion_ind" default:""`
	Isomer_ind           *string `json:"isomer_ind" default:""`
	Isotope_ind          *string `json:"isotope_ind" default:""`
	Mineral_ind          *string `json:"mineral_ind" default:""`
	Organic_matter_ind   *string `json:"organic_matter_ind" default:""`
	Ppdm_guid            *string `json:"ppdm_guid" default:""`
	Preferred_ind        *string `json:"preferred_ind" default:""`
	Preferred_long_name  *string `json:"preferred_long_name" default:""`
	Preferred_short_name *string `json:"preferred_short_name" default:""`
	Production_stuff_ind *string `json:"production_stuff_ind" default:""`
	Property_set_id      *string `json:"property_set_id" default:""`
	Remark               *string `json:"remark" default:""`
	Solvent_ind          *string `json:"solvent_ind" default:""`
	Source               *string `json:"source" default:""`
	Source_document_id   *string `json:"source_document_id" default:""`
	Row_changed_by       *string `json:"row_changed_by" default:""`
	Row_changed_date     *string `json:"row_changed_date" default:""`
	Row_created_by       *string `json:"row_created_by" default:""`
	Row_created_date     *string `json:"row_created_date" default:""`
	Row_effective_date   *string `json:"row_effective_date" default:""`
	Row_expiry_date      *string `json:"row_expiry_date" default:""`
	Row_quality          *string `json:"row_quality" default:""`
}
