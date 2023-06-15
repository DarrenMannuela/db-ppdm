package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmInfoItemContent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_info_item_content")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_info_item_content

	for rows.Next() {
		var rm_info_item_content dto.Rm_info_item_content
		if err := rows.Scan(&rm_info_item_content.Info_item_subtype, &rm_info_item_content.Information_item_id, &rm_info_item_content.Content_obs_no, &rm_info_item_content.Active_ind, &rm_info_item_content.Activity_obs_no, &rm_info_item_content.Analysis_id, &rm_info_item_content.Application_id, &rm_info_item_content.Applic_attachment_id, &rm_info_item_content.Area_id, &rm_info_item_content.Area_type, &rm_info_item_content.Authority_id, &rm_info_item_content.Ba_licensee_ba_id, &rm_info_item_content.Ba_license_id, &rm_info_item_content.Bin_grid_id, &rm_info_item_content.Bin_grid_source, &rm_info_item_content.Business_associate_id, &rm_info_item_content.Catalogue_equip_id, &rm_info_item_content.Column_name, &rm_info_item_content.Consent_id, &rm_info_item_content.Consult_id, &rm_info_item_content.Content_type, &rm_info_item_content.Contest_id, &rm_info_item_content.Contract_id, &rm_info_item_content.Discussion_id, &rm_info_item_content.Drill_report_id, &rm_info_item_content.Ecozone_id, &rm_info_item_content.Ecozone_set_id, &rm_info_item_content.Effective_date, &rm_info_item_content.Entitlement_id, &rm_info_item_content.Equipment_id, &rm_info_item_content.Expiry_date, &rm_info_item_content.Facility_id, &rm_info_item_content.Facility_license_id, &rm_info_item_content.Facility_type, &rm_info_item_content.Field_id, &rm_info_item_content.Finance_id, &rm_info_item_content.Fossil_desc_id, &rm_info_item_content.Fossil_id, &rm_info_item_content.Hse_incident_id, &rm_info_item_content.Instrument_id, &rm_info_item_content.Interest_set_id, &rm_info_item_content.Interest_set_seq_no, &rm_info_item_content.Interp_id, &rm_info_item_content.Jurisdiction, &rm_info_item_content.Land_request_id, &rm_info_item_content.Land_right_id, &rm_info_item_content.Land_right_subtype, &rm_info_item_content.Land_sale_number, &rm_info_item_content.Land_sale_offering_id, &rm_info_item_content.Lithology_log_id, &rm_info_item_content.Lith_log_source, &rm_info_item_content.Lith_sample_id, &rm_info_item_content.Notification_id, &rm_info_item_content.Obligation_id, &rm_info_item_content.Obligation_seq_no, &rm_info_item_content.Organization_id, &rm_info_item_content.Organization_seq_no, &rm_info_item_content.Paleo_summary_id, &rm_info_item_content.Pden_id, &rm_info_item_content.Pden_source, &rm_info_item_content.Pden_subtype, &rm_info_item_content.Physical_item_custody_id, &rm_info_item_content.Physical_item_id, &rm_info_item_content.Platform_id, &rm_info_item_content.Platform_sf_subtype, &rm_info_item_content.Pool_id, &rm_info_item_content.Ppdm_audit_row_guid, &rm_info_item_content.Ppdm_audit_seq_no, &rm_info_item_content.Ppdm_guid, &rm_info_item_content.Product_type, &rm_info_item_content.Prod_string_id, &rm_info_item_content.Prod_string_source, &rm_info_item_content.Project_id, &rm_info_item_content.Project_step_id, &rm_info_item_content.Rate_schedule_id, &rm_info_item_content.Remark, &rm_info_item_content.Resent_id, &rm_info_item_content.Reserve_class_id, &rm_info_item_content.Reserve_product, &rm_info_item_content.Reserve_revision_id, &rm_info_item_content.Reserve_revision_obs_no, &rm_info_item_content.Reserve_summary_id, &rm_info_item_content.Reserve_summary_obs_no, &rm_info_item_content.Restriction_id, &rm_info_item_content.Restriction_version, &rm_info_item_content.Sample_anal_source, &rm_info_item_content.Seis_license_id, &rm_info_item_content.Seis_patch_id, &rm_info_item_content.Seis_set_id, &rm_info_item_content.Seis_set_subtype, &rm_info_item_content.Source, &rm_info_item_content.Source_document_id, &rm_info_item_content.Spatial_description_id, &rm_info_item_content.Spatial_obs_no, &rm_info_item_content.Strat_column_id, &rm_info_item_content.Strat_column_source, &rm_info_item_content.Strat_field_station_id, &rm_info_item_content.Strat_name_set_id, &rm_info_item_content.Strat_unit_id, &rm_info_item_content.Support_facility_id, &rm_info_item_content.Support_facility_subtype, &rm_info_item_content.System_id, &rm_info_item_content.Table_name, &rm_info_item_content.Uwi, &rm_info_item_content.Velocity_volume_id, &rm_info_item_content.Wc_core_id, &rm_info_item_content.Wc_source, &rm_info_item_content.Wc_uwi, &rm_info_item_content.Wds_source, &rm_info_item_content.Wds_survey_id, &rm_info_item_content.Wds_uwi, &rm_info_item_content.Well_activity_source, &rm_info_item_content.Well_license_id, &rm_info_item_content.Well_license_source, &rm_info_item_content.Well_log_curve_id, &rm_info_item_content.Well_log_id, &rm_info_item_content.Well_log_source, &rm_info_item_content.Wl_dgtz_curve_id, &rm_info_item_content.Work_order_id, &rm_info_item_content.Wtf_fluid_type, &rm_info_item_content.Wtf_period_obs_no, &rm_info_item_content.Wtf_period_type, &rm_info_item_content.Wtf_run_num, &rm_info_item_content.Wtf_source, &rm_info_item_content.Wtf_test_num, &rm_info_item_content.Wtf_test_type, &rm_info_item_content.Wtp_period_obs_no, &rm_info_item_content.Wtp_period_type, &rm_info_item_content.Wtp_run_num, &rm_info_item_content.Wtp_source, &rm_info_item_content.Wtp_test_num, &rm_info_item_content.Wtp_test_type, &rm_info_item_content.Row_changed_by, &rm_info_item_content.Row_changed_date, &rm_info_item_content.Row_created_by, &rm_info_item_content.Row_created_date, &rm_info_item_content.Row_effective_date, &rm_info_item_content.Row_expiry_date, &rm_info_item_content.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_info_item_content to result
		result = append(result, rm_info_item_content)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmInfoItemContent(c *fiber.Ctx) error {
	var rm_info_item_content dto.Rm_info_item_content

	setDefaults(&rm_info_item_content)

	if err := json.Unmarshal(c.Body(), &rm_info_item_content); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_info_item_content values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103, :104, :105, :106, :107, :108, :109, :110, :111, :112, :113, :114, :115, :116, :117, :118, :119, :120, :121, :122, :123, :124, :125, :126, :127, :128, :129, :130, :131, :132, :133, :134, :135, :136, :137, :138, :139)")
	if err != nil {
		return err
	}
	rm_info_item_content.Row_created_date = formatDate(rm_info_item_content.Row_created_date)
	rm_info_item_content.Row_changed_date = formatDate(rm_info_item_content.Row_changed_date)
	rm_info_item_content.Effective_date = formatDateString(rm_info_item_content.Effective_date)
	rm_info_item_content.Expiry_date = formatDateString(rm_info_item_content.Expiry_date)
	rm_info_item_content.Row_effective_date = formatDateString(rm_info_item_content.Row_effective_date)
	rm_info_item_content.Row_expiry_date = formatDateString(rm_info_item_content.Row_expiry_date)






	rows, err := stmt.Exec(rm_info_item_content.Info_item_subtype, rm_info_item_content.Information_item_id, rm_info_item_content.Content_obs_no, rm_info_item_content.Active_ind, rm_info_item_content.Activity_obs_no, rm_info_item_content.Analysis_id, rm_info_item_content.Application_id, rm_info_item_content.Applic_attachment_id, rm_info_item_content.Area_id, rm_info_item_content.Area_type, rm_info_item_content.Authority_id, rm_info_item_content.Ba_licensee_ba_id, rm_info_item_content.Ba_license_id, rm_info_item_content.Bin_grid_id, rm_info_item_content.Bin_grid_source, rm_info_item_content.Business_associate_id, rm_info_item_content.Catalogue_equip_id, rm_info_item_content.Column_name, rm_info_item_content.Consent_id, rm_info_item_content.Consult_id, rm_info_item_content.Content_type, rm_info_item_content.Contest_id, rm_info_item_content.Contract_id, rm_info_item_content.Discussion_id, rm_info_item_content.Drill_report_id, rm_info_item_content.Ecozone_id, rm_info_item_content.Ecozone_set_id, rm_info_item_content.Effective_date, rm_info_item_content.Entitlement_id, rm_info_item_content.Equipment_id, rm_info_item_content.Expiry_date, rm_info_item_content.Facility_id, rm_info_item_content.Facility_license_id, rm_info_item_content.Facility_type, rm_info_item_content.Field_id, rm_info_item_content.Finance_id, rm_info_item_content.Fossil_desc_id, rm_info_item_content.Fossil_id, rm_info_item_content.Hse_incident_id, rm_info_item_content.Instrument_id, rm_info_item_content.Interest_set_id, rm_info_item_content.Interest_set_seq_no, rm_info_item_content.Interp_id, rm_info_item_content.Jurisdiction, rm_info_item_content.Land_request_id, rm_info_item_content.Land_right_id, rm_info_item_content.Land_right_subtype, rm_info_item_content.Land_sale_number, rm_info_item_content.Land_sale_offering_id, rm_info_item_content.Lithology_log_id, rm_info_item_content.Lith_log_source, rm_info_item_content.Lith_sample_id, rm_info_item_content.Notification_id, rm_info_item_content.Obligation_id, rm_info_item_content.Obligation_seq_no, rm_info_item_content.Organization_id, rm_info_item_content.Organization_seq_no, rm_info_item_content.Paleo_summary_id, rm_info_item_content.Pden_id, rm_info_item_content.Pden_source, rm_info_item_content.Pden_subtype, rm_info_item_content.Physical_item_custody_id, rm_info_item_content.Physical_item_id, rm_info_item_content.Platform_id, rm_info_item_content.Platform_sf_subtype, rm_info_item_content.Pool_id, rm_info_item_content.Ppdm_audit_row_guid, rm_info_item_content.Ppdm_audit_seq_no, rm_info_item_content.Ppdm_guid, rm_info_item_content.Product_type, rm_info_item_content.Prod_string_id, rm_info_item_content.Prod_string_source, rm_info_item_content.Project_id, rm_info_item_content.Project_step_id, rm_info_item_content.Rate_schedule_id, rm_info_item_content.Remark, rm_info_item_content.Resent_id, rm_info_item_content.Reserve_class_id, rm_info_item_content.Reserve_product, rm_info_item_content.Reserve_revision_id, rm_info_item_content.Reserve_revision_obs_no, rm_info_item_content.Reserve_summary_id, rm_info_item_content.Reserve_summary_obs_no, rm_info_item_content.Restriction_id, rm_info_item_content.Restriction_version, rm_info_item_content.Sample_anal_source, rm_info_item_content.Seis_license_id, rm_info_item_content.Seis_patch_id, rm_info_item_content.Seis_set_id, rm_info_item_content.Seis_set_subtype, rm_info_item_content.Source, rm_info_item_content.Source_document_id, rm_info_item_content.Spatial_description_id, rm_info_item_content.Spatial_obs_no, rm_info_item_content.Strat_column_id, rm_info_item_content.Strat_column_source, rm_info_item_content.Strat_field_station_id, rm_info_item_content.Strat_name_set_id, rm_info_item_content.Strat_unit_id, rm_info_item_content.Support_facility_id, rm_info_item_content.Support_facility_subtype, rm_info_item_content.System_id, rm_info_item_content.Table_name, rm_info_item_content.Uwi, rm_info_item_content.Velocity_volume_id, rm_info_item_content.Wc_core_id, rm_info_item_content.Wc_source, rm_info_item_content.Wc_uwi, rm_info_item_content.Wds_source, rm_info_item_content.Wds_survey_id, rm_info_item_content.Wds_uwi, rm_info_item_content.Well_activity_source, rm_info_item_content.Well_license_id, rm_info_item_content.Well_license_source, rm_info_item_content.Well_log_curve_id, rm_info_item_content.Well_log_id, rm_info_item_content.Well_log_source, rm_info_item_content.Wl_dgtz_curve_id, rm_info_item_content.Work_order_id, rm_info_item_content.Wtf_fluid_type, rm_info_item_content.Wtf_period_obs_no, rm_info_item_content.Wtf_period_type, rm_info_item_content.Wtf_run_num, rm_info_item_content.Wtf_source, rm_info_item_content.Wtf_test_num, rm_info_item_content.Wtf_test_type, rm_info_item_content.Wtp_period_obs_no, rm_info_item_content.Wtp_period_type, rm_info_item_content.Wtp_run_num, rm_info_item_content.Wtp_source, rm_info_item_content.Wtp_test_num, rm_info_item_content.Wtp_test_type, rm_info_item_content.Row_changed_by, rm_info_item_content.Row_changed_date, rm_info_item_content.Row_created_by, rm_info_item_content.Row_created_date, rm_info_item_content.Row_effective_date, rm_info_item_content.Row_expiry_date, rm_info_item_content.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmInfoItemContent(c *fiber.Ctx) error {
	var rm_info_item_content dto.Rm_info_item_content

	setDefaults(&rm_info_item_content)

	if err := json.Unmarshal(c.Body(), &rm_info_item_content); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_info_item_content.Info_item_subtype = id

    if rm_info_item_content.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_info_item_content set information_item_id = :1, content_obs_no = :2, active_ind = :3, activity_obs_no = :4, analysis_id = :5, application_id = :6, applic_attachment_id = :7, area_id = :8, area_type = :9, authority_id = :10, ba_licensee_ba_id = :11, ba_license_id = :12, bin_grid_id = :13, bin_grid_source = :14, business_associate_id = :15, catalogue_equip_id = :16, column_name = :17, consent_id = :18, consult_id = :19, content_type = :20, contest_id = :21, contract_id = :22, discussion_id = :23, drill_report_id = :24, ecozone_id = :25, ecozone_set_id = :26, effective_date = :27, entitlement_id = :28, equipment_id = :29, expiry_date = :30, facility_id = :31, facility_license_id = :32, facility_type = :33, field_id = :34, finance_id = :35, fossil_desc_id = :36, fossil_id = :37, hse_incident_id = :38, instrument_id = :39, interest_set_id = :40, interest_set_seq_no = :41, interp_id = :42, jurisdiction = :43, land_request_id = :44, land_right_id = :45, land_right_subtype = :46, land_sale_number = :47, land_sale_offering_id = :48, lithology_log_id = :49, lith_log_source = :50, lith_sample_id = :51, notification_id = :52, obligation_id = :53, obligation_seq_no = :54, organization_id = :55, organization_seq_no = :56, paleo_summary_id = :57, pden_id = :58, pden_source = :59, pden_subtype = :60, physical_item_custody_id = :61, physical_item_id = :62, platform_id = :63, platform_sf_subtype = :64, pool_id = :65, ppdm_audit_row_guid = :66, ppdm_audit_seq_no = :67, ppdm_guid = :68, product_type = :69, prod_string_id = :70, prod_string_source = :71, project_id = :72, project_step_id = :73, rate_schedule_id = :74, remark = :75, resent_id = :76, reserve_class_id = :77, reserve_product = :78, reserve_revision_id = :79, reserve_revision_obs_no = :80, reserve_summary_id = :81, reserve_summary_obs_no = :82, restriction_id = :83, restriction_version = :84, sample_anal_source = :85, seis_license_id = :86, seis_patch_id = :87, seis_set_id = :88, seis_set_subtype = :89, source = :90, source_document_id = :91, spatial_description_id = :92, spatial_obs_no = :93, strat_column_id = :94, strat_column_source = :95, strat_field_station_id = :96, strat_name_set_id = :97, strat_unit_id = :98, support_facility_id = :99, support_facility_subtype = :100, system_id = :101, table_name = :102, uwi = :103, velocity_volume_id = :104, wc_core_id = :105, wc_source = :106, wc_uwi = :107, wds_source = :108, wds_survey_id = :109, wds_uwi = :110, well_activity_source = :111, well_license_id = :112, well_license_source = :113, well_log_curve_id = :114, well_log_id = :115, well_log_source = :116, wl_dgtz_curve_id = :117, work_order_id = :118, wtf_fluid_type = :119, wtf_period_obs_no = :120, wtf_period_type = :121, wtf_run_num = :122, wtf_source = :123, wtf_test_num = :124, wtf_test_type = :125, wtp_period_obs_no = :126, wtp_period_type = :127, wtp_run_num = :128, wtp_source = :129, wtp_test_num = :130, wtp_test_type = :131, row_changed_by = :132, row_changed_date = :133, row_created_by = :134, row_effective_date = :135, row_expiry_date = :136, row_quality = :137 where info_item_subtype = :139")
	if err != nil {
		return err
	}

	rm_info_item_content.Row_changed_date = formatDate(rm_info_item_content.Row_changed_date)
	rm_info_item_content.Effective_date = formatDateString(rm_info_item_content.Effective_date)
	rm_info_item_content.Expiry_date = formatDateString(rm_info_item_content.Expiry_date)
	rm_info_item_content.Row_effective_date = formatDateString(rm_info_item_content.Row_effective_date)
	rm_info_item_content.Row_expiry_date = formatDateString(rm_info_item_content.Row_expiry_date)






	rows, err := stmt.Exec(rm_info_item_content.Information_item_id, rm_info_item_content.Content_obs_no, rm_info_item_content.Active_ind, rm_info_item_content.Activity_obs_no, rm_info_item_content.Analysis_id, rm_info_item_content.Application_id, rm_info_item_content.Applic_attachment_id, rm_info_item_content.Area_id, rm_info_item_content.Area_type, rm_info_item_content.Authority_id, rm_info_item_content.Ba_licensee_ba_id, rm_info_item_content.Ba_license_id, rm_info_item_content.Bin_grid_id, rm_info_item_content.Bin_grid_source, rm_info_item_content.Business_associate_id, rm_info_item_content.Catalogue_equip_id, rm_info_item_content.Column_name, rm_info_item_content.Consent_id, rm_info_item_content.Consult_id, rm_info_item_content.Content_type, rm_info_item_content.Contest_id, rm_info_item_content.Contract_id, rm_info_item_content.Discussion_id, rm_info_item_content.Drill_report_id, rm_info_item_content.Ecozone_id, rm_info_item_content.Ecozone_set_id, rm_info_item_content.Effective_date, rm_info_item_content.Entitlement_id, rm_info_item_content.Equipment_id, rm_info_item_content.Expiry_date, rm_info_item_content.Facility_id, rm_info_item_content.Facility_license_id, rm_info_item_content.Facility_type, rm_info_item_content.Field_id, rm_info_item_content.Finance_id, rm_info_item_content.Fossil_desc_id, rm_info_item_content.Fossil_id, rm_info_item_content.Hse_incident_id, rm_info_item_content.Instrument_id, rm_info_item_content.Interest_set_id, rm_info_item_content.Interest_set_seq_no, rm_info_item_content.Interp_id, rm_info_item_content.Jurisdiction, rm_info_item_content.Land_request_id, rm_info_item_content.Land_right_id, rm_info_item_content.Land_right_subtype, rm_info_item_content.Land_sale_number, rm_info_item_content.Land_sale_offering_id, rm_info_item_content.Lithology_log_id, rm_info_item_content.Lith_log_source, rm_info_item_content.Lith_sample_id, rm_info_item_content.Notification_id, rm_info_item_content.Obligation_id, rm_info_item_content.Obligation_seq_no, rm_info_item_content.Organization_id, rm_info_item_content.Organization_seq_no, rm_info_item_content.Paleo_summary_id, rm_info_item_content.Pden_id, rm_info_item_content.Pden_source, rm_info_item_content.Pden_subtype, rm_info_item_content.Physical_item_custody_id, rm_info_item_content.Physical_item_id, rm_info_item_content.Platform_id, rm_info_item_content.Platform_sf_subtype, rm_info_item_content.Pool_id, rm_info_item_content.Ppdm_audit_row_guid, rm_info_item_content.Ppdm_audit_seq_no, rm_info_item_content.Ppdm_guid, rm_info_item_content.Product_type, rm_info_item_content.Prod_string_id, rm_info_item_content.Prod_string_source, rm_info_item_content.Project_id, rm_info_item_content.Project_step_id, rm_info_item_content.Rate_schedule_id, rm_info_item_content.Remark, rm_info_item_content.Resent_id, rm_info_item_content.Reserve_class_id, rm_info_item_content.Reserve_product, rm_info_item_content.Reserve_revision_id, rm_info_item_content.Reserve_revision_obs_no, rm_info_item_content.Reserve_summary_id, rm_info_item_content.Reserve_summary_obs_no, rm_info_item_content.Restriction_id, rm_info_item_content.Restriction_version, rm_info_item_content.Sample_anal_source, rm_info_item_content.Seis_license_id, rm_info_item_content.Seis_patch_id, rm_info_item_content.Seis_set_id, rm_info_item_content.Seis_set_subtype, rm_info_item_content.Source, rm_info_item_content.Source_document_id, rm_info_item_content.Spatial_description_id, rm_info_item_content.Spatial_obs_no, rm_info_item_content.Strat_column_id, rm_info_item_content.Strat_column_source, rm_info_item_content.Strat_field_station_id, rm_info_item_content.Strat_name_set_id, rm_info_item_content.Strat_unit_id, rm_info_item_content.Support_facility_id, rm_info_item_content.Support_facility_subtype, rm_info_item_content.System_id, rm_info_item_content.Table_name, rm_info_item_content.Uwi, rm_info_item_content.Velocity_volume_id, rm_info_item_content.Wc_core_id, rm_info_item_content.Wc_source, rm_info_item_content.Wc_uwi, rm_info_item_content.Wds_source, rm_info_item_content.Wds_survey_id, rm_info_item_content.Wds_uwi, rm_info_item_content.Well_activity_source, rm_info_item_content.Well_license_id, rm_info_item_content.Well_license_source, rm_info_item_content.Well_log_curve_id, rm_info_item_content.Well_log_id, rm_info_item_content.Well_log_source, rm_info_item_content.Wl_dgtz_curve_id, rm_info_item_content.Work_order_id, rm_info_item_content.Wtf_fluid_type, rm_info_item_content.Wtf_period_obs_no, rm_info_item_content.Wtf_period_type, rm_info_item_content.Wtf_run_num, rm_info_item_content.Wtf_source, rm_info_item_content.Wtf_test_num, rm_info_item_content.Wtf_test_type, rm_info_item_content.Wtp_period_obs_no, rm_info_item_content.Wtp_period_type, rm_info_item_content.Wtp_run_num, rm_info_item_content.Wtp_source, rm_info_item_content.Wtp_test_num, rm_info_item_content.Wtp_test_type, rm_info_item_content.Row_changed_by, rm_info_item_content.Row_changed_date, rm_info_item_content.Row_created_by, rm_info_item_content.Row_effective_date, rm_info_item_content.Row_expiry_date, rm_info_item_content.Row_quality, rm_info_item_content.Info_item_subtype)
	if err != nil {
		return err
	}

	if count, err := rows.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.Status(201).JSON(rows)
}

func PatchRmInfoItemContent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_info_item_content set "
	values := []interface{}{}
	i := 1
	for key, value := range updatedFields {
		query += key + " = :" + strconv.Itoa(i)
		i++
		if i <= len(updatedFields) {
			query += ", "
		}
		if key == "row_changed_date" {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDate(&date)
				value = formattedDate
			}
		} else if key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where info_item_subtype = :id"

	stmt, err := db.Prepare(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	values = append(values, id)
	res, err := stmt.Exec(values...)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if count, err := res.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.JSON(res)
}

func DeleteRmInfoItemContent(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_info_item_content dto.Rm_info_item_content
	rm_info_item_content.Info_item_subtype = id

	stmt, err := db.Prepare("delete from rm_info_item_content where info_item_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_info_item_content.Info_item_subtype)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return c.Status(404).SendString("No matching record found")
	}

	return c.JSON(rows)
}


