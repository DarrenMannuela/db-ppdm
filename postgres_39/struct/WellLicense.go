package dto

import (
	"time"
)

type Well_license struct {
	Uwi                        string     `json:"uwi"`
	Source                     string     `json:"source"`
	License_id                 string     `json:"license_id"`
	Active_ind                 *string    `json:"active_ind"`
	Agent                      *string    `json:"agent"`
	Application_id             *string    `json:"application_id"`
	Authorized_strat_unit_id   *string    `json:"authorized_strat_unit_id"`
	Bidding_round_num          *string    `json:"bidding_round_num"`
	Contractor                 *string    `json:"contractor"`
	Direction_to_loc           *float64   `json:"direction_to_loc"`
	Direction_to_loc_ouom      *string    `json:"direction_to_loc_ouom"`
	Distance_ref_point         *string    `json:"distance_ref_point"`
	Distance_to_loc            *float64   `json:"distance_to_loc"`
	Distance_to_loc_ouom       *string    `json:"distance_to_loc_ouom"`
	Drill_rig_num              *string    `json:"drill_rig_num"`
	Drill_slot_no              *int       `json:"drill_slot_no"`
	Drill_tool                 *string    `json:"drill_tool"`
	Effective_date             *time.Time `json:"effective_date"`
	Exception_granted          *string    `json:"exception_granted"`
	Exception_requested        *string    `json:"exception_requested"`
	Expired_ind                *string    `json:"expired_ind"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Fees_paid_ind              *string    `json:"fees_paid_ind"`
	Licensee                   *string    `json:"licensee"`
	Licensee_contact_id        *string    `json:"licensee_contact_id"`
	License_date               *time.Time `json:"license_date"`
	License_num                *string    `json:"license_num"`
	No_of_wells                *float64   `json:"no_of_wells"`
	Offshore_completion_type   *string    `json:"offshore_completion_type"`
	Permit_reference_num       *string    `json:"permit_reference_num"`
	Permit_reissue_date        *time.Time `json:"permit_reissue_date"`
	Permit_type                *string    `json:"permit_type"`
	Platform_name              *string    `json:"platform_name"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Projected_depth            *float64   `json:"projected_depth"`
	Projected_depth_ouom       *string    `json:"projected_depth_ouom"`
	Projected_strat_unit_id    *string    `json:"projected_strat_unit_id"`
	Projected_tvd              *float64   `json:"projected_tvd"`
	Projected_tvd_ouom         *string    `json:"projected_tvd_ouom"`
	Proposed_spud_date         *time.Time `json:"proposed_spud_date"`
	Purpose                    *string    `json:"purpose"`
	Rate_schedule_id           *string    `json:"rate_schedule_id"`
	Regulation_text            *string    `json:"regulation_text"`
	Regulation_text_source_doc *string    `json:"regulation_text_source_doc"`
	Regulatory_agency          *string    `json:"regulatory_agency"`
	Regulatory_contact_id      *string    `json:"regulatory_contact_id"`
	Remark                     *string    `json:"remark"`
	Rig_code                   *string    `json:"rig_code"`
	Rig_sf_subtype             *string    `json:"rig_sf_subtype"`
	Rig_substr_height          *float64   `json:"rig_substr_height"`
	Rig_substr_height_ouom     *string    `json:"rig_substr_height_ouom"`
	Rig_support_facility_id    *string    `json:"rig_support_facility_id"`
	Rig_type                   *string    `json:"rig_type"`
	Section_of_regulation      *string    `json:"section_of_regulation"`
	Strat_name_set_id          *string    `json:"strat_name_set_id"`
	Surveyor                   *string    `json:"surveyor"`
	Target_objective_fluid     *string    `json:"target_objective_fluid"`
	Violation_ind              *string    `json:"violation_ind"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
