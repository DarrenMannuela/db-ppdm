package dto

type Land_right struct {
	Land_right_subtype          string   `json:"land_right_subtype" default:"source"`
	Land_right_id               string   `json:"land_right_id" default:"source"`
	Acqtn_date                  *string  `json:"acqtn_date" default:""`
	Active_ind                  *string  `json:"active_ind" default:""`
	Area_id                     *string  `json:"area_id" default:""`
	Area_type                   *string  `json:"area_type" default:""`
	Calculated_interest         *float64 `json:"calculated_interest" default:""`
	Case_serial_num             *string  `json:"case_serial_num" default:""`
	Confidential_ind            *string  `json:"confidential_ind" default:""`
	Cost_center_num             *string  `json:"cost_center_num" default:""`
	Effective_date              *string  `json:"effective_date" default:""`
	Expiry_date                 *string  `json:"expiry_date" default:""`
	First_platform_approve_date *string  `json:"first_platform_approve_date" default:""`
	First_platform_install_date *string  `json:"first_platform_install_date" default:""`
	First_production_date       *string  `json:"first_production_date" default:""`
	Fractional_interest         *string  `json:"fractional_interest" default:""`
	Granted_right_type          *string  `json:"granted_right_type" default:""`
	Gross_size                  *float64 `json:"gross_size" default:""`
	Gross_size_ouom             *string  `json:"gross_size_ouom" default:""`
	High_water_depth            *float64 `json:"high_water_depth" default:""`
	High_water_depth_ouom       *string  `json:"high_water_depth_ouom" default:""`
	Inactivation_date           *string  `json:"inactivation_date" default:""`
	Incent_cert_no              *int     `json:"incent_cert_no" default:""`
	Issue_date                  *string  `json:"issue_date" default:""`
	Jurisdiction                *string  `json:"jurisdiction" default:""`
	Land_acqtn_method           *string  `json:"land_acqtn_method" default:""`
	Land_case_action            *string  `json:"land_case_action" default:""`
	Land_case_type              *string  `json:"land_case_type" default:""`
	Land_property_type          *string  `json:"land_property_type" default:""`
	Land_right_category         *string  `json:"land_right_category" default:""`
	Land_sale_number            *string  `json:"land_sale_number" default:""`
	Land_sale_offering_id       *string  `json:"land_sale_offering_id" default:""`
	Last_action_date            *string  `json:"last_action_date" default:""`
	Lessor_name                 *string  `json:"lessor_name" default:""`
	Lessor_num                  *string  `json:"lessor_num" default:""`
	Lessor_type                 *string  `json:"lessor_type" default:""`
	Low_water_depth             *float64 `json:"low_water_depth" default:""`
	Low_water_depth_ouom        *string  `json:"low_water_depth_ouom" default:""`
	Municipality                *string  `json:"municipality" default:""`
	Occupant_name               *string  `json:"occupant_name" default:""`
	Offshore_distance           *float64 `json:"offshore_distance" default:""`
	Offshore_distance_ouom      *string  `json:"offshore_distance_ouom" default:""`
	Oil_sand_deposit            *string  `json:"oil_sand_deposit" default:""`
	Overlap_ind                 *string  `json:"overlap_ind" default:""`
	Platform_count              *int     `json:"platform_count" default:""`
	Ppdm_guid                   *string  `json:"ppdm_guid" default:""`
	Primary_term                *float64 `json:"primary_term" default:""`
	Primary_term_ouom           *string  `json:"primary_term_ouom" default:""`
	Producing_ind               *string  `json:"producing_ind" default:""`
	Proprietary_num             *string  `json:"proprietary_num" default:""`
	Qualification_date          *string  `json:"qualification_date" default:""`
	Reclamation_cert_num        *string  `json:"reclamation_cert_num" default:""`
	Reclamation_end_date        *string  `json:"reclamation_end_date" default:""`
	Reclamation_start_date      *string  `json:"reclamation_start_date" default:""`
	Relinquishable_ind          *string  `json:"relinquishable_ind" default:""`
	Remark                      *string  `json:"remark" default:""`
	Rental_allocation_ind       *string  `json:"rental_allocation_ind" default:""`
	Report_acreage_ind          *string  `json:"report_acreage_ind" default:""`
	Scheme_approval_num         *string  `json:"scheme_approval_num" default:""`
	Scheme_expiry_date          *string  `json:"scheme_expiry_date" default:""`
	Secondary_term              *float64 `json:"secondary_term" default:""`
	Secondary_term_ouom         *string  `json:"secondary_term_ouom" default:""`
	Source                      *string  `json:"source" default:""`
	Subsurface_ind              *string  `json:"subsurface_ind" default:""`
	Surface_ind                 *string  `json:"surface_ind" default:""`
	Termin_notice_period        *float64 `json:"termin_notice_period" default:""`
	Termin_notice_period_ouom   *string  `json:"termin_notice_period_ouom" default:""`
	Term_exiry_date             *string  `json:"term_exiry_date" default:""`
	Unit_operated_ind           *string  `json:"unit_operated_ind" default:""`
	Well_qualific_type          *string  `json:"well_qualific_type" default:""`
	Row_changed_by              *string  `json:"row_changed_by" default:""`
	Row_changed_date            *string  `json:"row_changed_date" default:""`
	Row_created_by              *string  `json:"row_created_by" default:""`
	Row_created_date            *string  `json:"row_created_date" default:""`
	Row_effective_date          *string  `json:"row_effective_date" default:""`
	Row_expiry_date             *string  `json:"row_expiry_date" default:""`
	Row_quality                 *string  `json:"row_quality" default:""`
}
