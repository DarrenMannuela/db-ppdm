package dto

type Sf_bridge struct {
	Sf_subtype           string   `json:"sf_subtype" default:"source"`
	Bridge_id            string   `json:"bridge_id" default:"source"`
	Active_ind           *string  `json:"active_ind" default:""`
	Bridge_capacity      *float64 `json:"bridge_capacity" default:""`
	Bridge_capacity_ouom *string  `json:"bridge_capacity_ouom" default:""`
	Bridge_height        *float64 `json:"bridge_height" default:""`
	Bridge_height_ouom   *string  `json:"bridge_height_ouom" default:""`
	Bridge_length        *float64 `json:"bridge_length" default:""`
	Bridge_length_ouom   *string  `json:"bridge_length_ouom" default:""`
	Bridge_type          *string  `json:"bridge_type" default:""`
	Bridge_width         *float64 `json:"bridge_width" default:""`
	Bridge_width_ouom    *string  `json:"bridge_width_ouom" default:""`
	Effective_date       *string  `json:"effective_date" default:""`
	Expiry_date          *string  `json:"expiry_date" default:""`
	Ppdm_guid            *string  `json:"ppdm_guid" default:""`
	Remark               *string  `json:"remark" default:""`
	Source               *string  `json:"source" default:""`
	Surface_type         *string  `json:"surface_type" default:""`
	Row_changed_by       *string  `json:"row_changed_by" default:""`
	Row_changed_date     *string  `json:"row_changed_date" default:""`
	Row_created_by       *string  `json:"row_created_by" default:""`
	Row_created_date     *string  `json:"row_created_date" default:""`
	Row_effective_date   *string  `json:"row_effective_date" default:""`
	Row_expiry_date      *string  `json:"row_expiry_date" default:""`
	Row_quality          *string  `json:"row_quality" default:""`
}
