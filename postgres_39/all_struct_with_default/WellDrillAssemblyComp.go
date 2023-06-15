package dto

type Well_drill_assembly_comp struct {
	Uwi                       string   `json:"uwi" default:"source"`
	Assembly_id               string   `json:"assembly_id" default:"source"`
	Component_id              string   `json:"component_id" default:"source"`
	Active_ind                *string  `json:"active_ind" default:""`
	Catalogue_equip_id        *string  `json:"catalogue_equip_id" default:""`
	Component_count           *int     `json:"component_count" default:""`
	Component_desc            *string  `json:"component_desc" default:""`
	Component_length          *float64 `json:"component_length" default:""`
	Component_length_ouom     *string  `json:"component_length_ouom" default:""`
	Component_seq_no          *int     `json:"component_seq_no" default:""`
	Component_type            *string  `json:"component_type" default:""`
	Component_weight          *float64 `json:"component_weight" default:""`
	Component_weight_ouom     *string  `json:"component_weight_ouom" default:""`
	Effective_date            *string  `json:"effective_date" default:""`
	Expiry_date               *string  `json:"expiry_date" default:""`
	Max_inside_diameter       *float64 `json:"max_inside_diameter" default:""`
	Max_inside_diameter_ouom  *string  `json:"max_inside_diameter_ouom" default:""`
	Max_outside_diameter      *float64 `json:"max_outside_diameter" default:""`
	Max_outside_diameter_ouom *string  `json:"max_outside_diameter_ouom" default:""`
	Min_inside_diameter       *float64 `json:"min_inside_diameter" default:""`
	Min_inside_diameter_ouom  *string  `json:"min_inside_diameter_ouom" default:""`
	Min_outside_diameter      *float64 `json:"min_outside_diameter" default:""`
	Min_outside_diameter_ouom *string  `json:"min_outside_diameter_ouom" default:""`
	Ppdm_guid                 *string  `json:"ppdm_guid" default:""`
	Remark                    *string  `json:"remark" default:""`
	Serial_number             *string  `json:"serial_number" default:""`
	Source                    *string  `json:"source" default:""`
	Row_changed_by            *string  `json:"row_changed_by" default:""`
	Row_changed_date          *string  `json:"row_changed_date" default:""`
	Row_created_by            *string  `json:"row_created_by" default:""`
	Row_created_date          *string  `json:"row_created_date" default:""`
	Row_effective_date        *string  `json:"row_effective_date" default:""`
	Row_expiry_date           *string  `json:"row_expiry_date" default:""`
	Row_quality               *string  `json:"row_quality" default:""`
}
