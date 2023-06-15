package dto

import (
	"time"
)

type Ppdm_group_object struct {
	Group_id                 string     `json:"group_id"`
	Object_obs_no            int        `json:"object_obs_no"`
	Active_ind               *string    `json:"active_ind"`
	Code_version_obs_no      *int       `json:"code_version_obs_no"`
	Code_version_source      *string    `json:"code_version_source"`
	Column_alias_id          *string    `json:"column_alias_id"`
	Column_name              *string    `json:"column_name"`
	Constraint_name          *string    `json:"constraint_name"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Group_use                *string    `json:"group_use"`
	Index_id                 *string    `json:"index_id"`
	Object_type              *string    `json:"object_type"`
	Output_font              *string    `json:"output_font"`
	Output_font_backgr_color *string    `json:"output_font_backgr_color"`
	Output_font_color        *string    `json:"output_font_color"`
	Output_font_effect       *string    `json:"output_font_effect"`
	Output_font_size         *float64   `json:"output_font_size"`
	Output_font_size_uom     *string    `json:"output_font_size_uom"`
	Output_heading           *string    `json:"output_heading"`
	Output_length            *int       `json:"output_length"`
	Output_length_uom        *string    `json:"output_length_uom"`
	Output_precision         *float64   `json:"output_precision"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Preferred_uom            *string    `json:"preferred_uom"`
	Procedure_id             *string    `json:"procedure_id"`
	Property_set_id          *string    `json:"property_set_id"`
	Remark                   *string    `json:"remark"`
	Rule_id                  *string    `json:"rule_id"`
	Source                   *string    `json:"source"`
	Sw_application_id        *string    `json:"sw_application_id"`
	System_id                *string    `json:"system_id"`
	Table_alias              *string    `json:"table_alias"`
	Table_name               *string    `json:"table_name"`
	Use_rule_description     *string    `json:"use_rule_description"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
