package dto

type Sf_vessel struct {
	Sf_subtype                  string   `json:"sf_subtype" default:"source"`
	Vessel_id                   string   `json:"vessel_id" default:"source"`
	Active_ind                  *string  `json:"active_ind" default:""`
	Area_id                     *string  `json:"area_id" default:""`
	Area_type                   *string  `json:"area_type" default:""`
	Backup_antenna_location     *string  `json:"backup_antenna_location" default:""`
	Backup_antenna_type         *string  `json:"backup_antenna_type" default:""`
	Current_operator            *string  `json:"current_operator" default:""`
	Drill_hole_position         *string  `json:"drill_hole_position" default:""`
	Effective_date              *string  `json:"effective_date" default:""`
	Expiry_date                 *string  `json:"expiry_date" default:""`
	Positioning_method          *string  `json:"positioning_method" default:""`
	Ppdm_guid                   *string  `json:"ppdm_guid" default:""`
	Primary_antenna_location    *string  `json:"primary_antenna_location" default:""`
	Primary_antenna_type        *string  `json:"primary_antenna_type" default:""`
	Remark                      *string  `json:"remark" default:""`
	Secondary_antennal_location *string  `json:"secondary_antennal_location" default:""`
	Secondary_antenna_type      *string  `json:"secondary_antenna_type" default:""`
	Source                      *string  `json:"source" default:""`
	Vessel_beam                 *float64 `json:"vessel_beam" default:""`
	Vessel_displacement         *float64 `json:"vessel_displacement" default:""`
	Vessel_draft                *float64 `json:"vessel_draft" default:""`
	Vessel_length               *float64 `json:"vessel_length" default:""`
	Vessel_name                 *string  `json:"vessel_name" default:""`
	Vessel_size                 *string  `json:"vessel_size" default:""`
	Vessel_type                 *string  `json:"vessel_type" default:""`
	Row_changed_by              *string  `json:"row_changed_by" default:""`
	Row_changed_date            *string  `json:"row_changed_date" default:""`
	Row_created_by              *string  `json:"row_created_by" default:""`
	Row_created_date            *string  `json:"row_created_date" default:""`
	Row_effective_date          *string  `json:"row_effective_date" default:""`
	Row_expiry_date             *string  `json:"row_expiry_date" default:""`
	Row_quality                 *string  `json:"row_quality" default:""`
}
