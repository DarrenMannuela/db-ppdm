package dto

import (
	"time"
)

type Report_hier_level struct {
	Report_hierarchy_id    string     `json:"report_hierarchy_id"`
	Component_id           string     `json:"component_id"`
	Active_ind             *string    `json:"active_ind"`
	Area_id                *string    `json:"area_id"`
	Area_type              *string    `json:"area_type"`
	Ba_organization_id     *string    `json:"ba_organization_id"`
	Ba_organization_seq_no *int       `json:"ba_organization_seq_no"`
	Business_associate_id  *string    `json:"business_associate_id"`
	Component_type         *string    `json:"component_type"`
	Contract_id            *string    `json:"contract_id"`
	Effective_date         *time.Time `json:"effective_date"`
	Equipment_id           *string    `json:"equipment_id"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Facility_id            *string    `json:"facility_id"`
	Facility_type          *string    `json:"facility_type"`
	Field_id               *string    `json:"field_id"`
	Finance_id             *string    `json:"finance_id"`
	Hierarchy_type_id      *string    `json:"hierarchy_type_id"`
	Land_right_id          *string    `json:"land_right_id"`
	Land_right_subtype     *string    `json:"land_right_subtype"`
	Level_seq_no           *int       `json:"level_seq_no"`
	Parent_component_id    *string    `json:"parent_component_id"`
	Pool_id                *string    `json:"pool_id"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Project_id             *string    `json:"project_id"`
	Remark                 *string    `json:"remark"`
	Sf_subtype             *string    `json:"sf_subtype"`
	Source                 *string    `json:"source"`
	Strat_name_set_id      *string    `json:"strat_name_set_id"`
	Strat_unit_id          *string    `json:"strat_unit_id"`
	Support_facility_id    *string    `json:"support_facility_id"`
	Uwi                    *string    `json:"uwi"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
