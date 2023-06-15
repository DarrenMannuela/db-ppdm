package dto

import (
	"time"
)

type Land_right struct {
	Land_right_subtype          string     `json:"land_right_subtype"`
	Land_right_id               string     `json:"land_right_id"`
	Acqtn_date                  *time.Time `json:"acqtn_date"`
	Active_ind                  *string    `json:"active_ind"`
	Area_id                     *string    `json:"area_id"`
	Area_type                   *string    `json:"area_type"`
	Calculated_interest         *float64   `json:"calculated_interest"`
	Case_serial_num             *string    `json:"case_serial_num"`
	Confidential_ind            *string    `json:"confidential_ind"`
	Cost_center_num             *string    `json:"cost_center_num"`
	Effective_date              *time.Time `json:"effective_date"`
	Expiry_date                 *time.Time `json:"expiry_date"`
	First_platform_approve_date *time.Time `json:"first_platform_approve_date"`
	First_platform_install_date *time.Time `json:"first_platform_install_date"`
	First_production_date       *time.Time `json:"first_production_date"`
	Fractional_interest         *string    `json:"fractional_interest"`
	Granted_right_type          *string    `json:"granted_right_type"`
	Gross_size                  *float64   `json:"gross_size"`
	Gross_size_ouom             *string    `json:"gross_size_ouom"`
	High_water_depth            *float64   `json:"high_water_depth"`
	High_water_depth_ouom       *string    `json:"high_water_depth_ouom"`
	Inactivation_date           *time.Time `json:"inactivation_date"`
	Incent_cert_no              *int       `json:"incent_cert_no"`
	Issue_date                  *time.Time `json:"issue_date"`
	Jurisdiction                *string    `json:"jurisdiction"`
	Land_acqtn_method           *string    `json:"land_acqtn_method"`
	Land_case_action            *string    `json:"land_case_action"`
	Land_case_type              *string    `json:"land_case_type"`
	Land_property_type          *string    `json:"land_property_type"`
	Land_right_category         *string    `json:"land_right_category"`
	Land_sale_number            *string    `json:"land_sale_number"`
	Land_sale_offering_id       *string    `json:"land_sale_offering_id"`
	Last_action_date            *time.Time `json:"last_action_date"`
	Lessor_name                 *string    `json:"lessor_name"`
	Lessor_num                  *string    `json:"lessor_num"`
	Lessor_type                 *string    `json:"lessor_type"`
	Low_water_depth             *float64   `json:"low_water_depth"`
	Low_water_depth_ouom        *string    `json:"low_water_depth_ouom"`
	Municipality                *string    `json:"municipality"`
	Occupant_name               *string    `json:"occupant_name"`
	Offshore_distance           *float64   `json:"offshore_distance"`
	Offshore_distance_ouom      *string    `json:"offshore_distance_ouom"`
	Oil_sand_deposit            *string    `json:"oil_sand_deposit"`
	Overlap_ind                 *string    `json:"overlap_ind"`
	Platform_count              *int       `json:"platform_count"`
	Ppdm_guid                   string     `json:"ppdm_guid"`
	Primary_term                *float64   `json:"primary_term"`
	Primary_term_ouom           *string    `json:"primary_term_ouom"`
	Producing_ind               *string    `json:"producing_ind"`
	Proprietary_num             *string    `json:"proprietary_num"`
	Qualification_date          *time.Time `json:"qualification_date"`
	Reclamation_cert_num        *string    `json:"reclamation_cert_num"`
	Reclamation_end_date        *time.Time `json:"reclamation_end_date"`
	Reclamation_start_date      *time.Time `json:"reclamation_start_date"`
	Relinquishable_ind          *string    `json:"relinquishable_ind"`
	Remark                      *string    `json:"remark"`
	Rental_allocation_ind       *string    `json:"rental_allocation_ind"`
	Report_acreage_ind          *string    `json:"report_acreage_ind"`
	Scheme_approval_num         *string    `json:"scheme_approval_num"`
	Scheme_expiry_date          *time.Time `json:"scheme_expiry_date"`
	Secondary_term              *float64   `json:"secondary_term"`
	Secondary_term_ouom         *string    `json:"secondary_term_ouom"`
	Source                      *string    `json:"source"`
	Subsurface_ind              *string    `json:"subsurface_ind"`
	Surface_ind                 *string    `json:"surface_ind"`
	Termin_notice_period        *float64   `json:"termin_notice_period"`
	Termin_notice_period_ouom   *string    `json:"termin_notice_period_ouom"`
	Term_exiry_date             *time.Time `json:"term_exiry_date"`
	Unit_operated_ind           *string    `json:"unit_operated_ind"`
	Well_qualific_type          *string    `json:"well_qualific_type"`
	Row_changed_by              *string    `json:"row_changed_by"`
	Row_changed_date            *time.Time `json:"row_changed_date"`
	Row_created_by              *string    `json:"row_created_by"`
	Row_created_date            *time.Time `json:"row_created_date"`
	Row_effective_date          *time.Time `json:"row_effective_date"`
	Row_expiry_date             *time.Time `json:"row_expiry_date"`
	Row_quality                 *string    `json:"row_quality"`
}
