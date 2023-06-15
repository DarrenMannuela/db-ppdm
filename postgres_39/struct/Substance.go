package dto

import (
	"time"
)

type Substance struct {
	Substance_id         string     `json:"substance_id"`
	Active_ind           *string    `json:"active_ind"`
	Api_gravity_max      *float64   `json:"api_gravity_max"`
	Api_gravity_min      *float64   `json:"api_gravity_min"`
	Atomic_mass          *float64   `json:"atomic_mass"`
	Atomic_mass_number   *string    `json:"atomic_mass_number"`
	Atomic_weight        *float64   `json:"atomic_weight"`
	Carbon_count         *int       `json:"carbon_count"`
	Carbon_count_max     *int       `json:"carbon_count_max"`
	Carbon_count_min     *int       `json:"carbon_count_min"`
	Composition_formula  *string    `json:"composition_formula"`
	Conversion_quantity  *string    `json:"conversion_quantity"`
	Effective_date       *time.Time `json:"effective_date"`
	Expiry_date          *time.Time `json:"expiry_date"`
	Ionic_charge         *float64   `json:"ionic_charge"`
	Molecular_mass       *float64   `json:"molecular_mass"`
	M_z_ion              *string    `json:"m_z_ion"`
	Phase_type           *string    `json:"phase_type"`
	Ph_value             *float64   `json:"ph_value"`
	Ppdm_guid            string     `json:"ppdm_guid"`
	Preferred_ind        *string    `json:"preferred_ind"`
	Preferred_long_name  *string    `json:"preferred_long_name"`
	Preferred_short_name *string    `json:"preferred_short_name"`
	Preferred_uom_id     *string    `json:"preferred_uom_id"`
	Property_set_id      *string    `json:"property_set_id"`
	Remark               *string    `json:"remark"`
	Source               *string    `json:"source"`
	Source_document_id   *string    `json:"source_document_id"`
	Substance_definition *string    `json:"substance_definition"`
	Valence_value        *float64   `json:"valence_value"`
	Row_changed_by       *string    `json:"row_changed_by"`
	Row_changed_date     *time.Time `json:"row_changed_date"`
	Row_created_by       *string    `json:"row_created_by"`
	Row_created_date     *time.Time `json:"row_created_date"`
	Row_effective_date   *time.Time `json:"row_effective_date"`
	Row_expiry_date      *time.Time `json:"row_expiry_date"`
	Row_quality          *string    `json:"row_quality"`
}
