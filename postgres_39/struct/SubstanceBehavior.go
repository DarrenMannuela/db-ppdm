package dto

import (
	"time"
)

type Substance_behavior struct {
	Substance_id         string     `json:"substance_id"`
	Active_ind           *string    `json:"active_ind"`
	Composite_ind        *string    `json:"composite_ind"`
	Drilling_mud_ind     *string    `json:"drilling_mud_ind"`
	Effective_date       *time.Time `json:"effective_date"`
	Element_ind          *string    `json:"element_ind"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Hydrocarbon_ind      *string    `json:"hydrocarbon_ind"`
	Ion_ind              *string    `json:"ion_ind"`
	Isomer_ind           *string    `json:"isomer_ind"`
	Isotope_ind          *string    `json:"isotope_ind"`
	Mineral_ind          *string    `json:"mineral_ind"`
	Organic_matter_ind   *string    `json:"organic_matter_ind"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Preferred_ind        *string    `json:"preferred_ind"`
	Preferred_long_name  *string    `json:"preferred_long_name"`
	Preferred_short_name *string    `json:"preferred_short_name"`
	Production_stuff_ind *string    `json:"production_stuff_ind"`
	Property_set_id      *string    `json:"property_set_id"`
	Remark               *string    `json:"remark"`
	Solvent_ind          *string    `json:"solvent_ind"`
	Source               *string    `json:"source"`
	Source_document_id   *string    `json:"source_document_id"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
