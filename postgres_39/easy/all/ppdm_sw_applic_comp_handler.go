package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmSwApplicComp(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_sw_applic_comp")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_sw_applic_comp

	for rows.Next() {
		var ppdm_sw_applic_comp dto.Ppdm_sw_applic_comp
		if err := rows.Scan(&ppdm_sw_applic_comp.Sw_application_id, &ppdm_sw_applic_comp.Component_obs_no, &ppdm_sw_applic_comp.Active_ind, &ppdm_sw_applic_comp.Activity_obs_no, &ppdm_sw_applic_comp.Analysis_id, &ppdm_sw_applic_comp.Anl_source, &ppdm_sw_applic_comp.Application_id, &ppdm_sw_applic_comp.Area_id, &ppdm_sw_applic_comp.Area_type, &ppdm_sw_applic_comp.Authority_id, &ppdm_sw_applic_comp.Ba_organization_id, &ppdm_sw_applic_comp.Ba_organization_seq_no, &ppdm_sw_applic_comp.Business_associate_id, &ppdm_sw_applic_comp.Catalogue_additive_id, &ppdm_sw_applic_comp.Catalogue_equip_id, &ppdm_sw_applic_comp.Classification_system_id, &ppdm_sw_applic_comp.Class_level_id, &ppdm_sw_applic_comp.Component_type, &ppdm_sw_applic_comp.Consent_id, &ppdm_sw_applic_comp.Consult_id, &ppdm_sw_applic_comp.Contest_id, &ppdm_sw_applic_comp.Contract_id, &ppdm_sw_applic_comp.Core_id, &ppdm_sw_applic_comp.Core_type, &ppdm_sw_applic_comp.Description_obs_no, &ppdm_sw_applic_comp.Ecozone_id, &ppdm_sw_applic_comp.Effective_date, &ppdm_sw_applic_comp.Employee_ba_id, &ppdm_sw_applic_comp.Employee_obs_no, &ppdm_sw_applic_comp.Employer_ba_id, &ppdm_sw_applic_comp.Entitlement_id, &ppdm_sw_applic_comp.Equipment_id, &ppdm_sw_applic_comp.Expiry_date, &ppdm_sw_applic_comp.Facility_id, &ppdm_sw_applic_comp.Facility_type, &ppdm_sw_applic_comp.Field_id, &ppdm_sw_applic_comp.Field_station_id, &ppdm_sw_applic_comp.Finance_id, &ppdm_sw_applic_comp.Fossil_id, &ppdm_sw_applic_comp.Incident_id, &ppdm_sw_applic_comp.Incident_set_id, &ppdm_sw_applic_comp.Incident_type, &ppdm_sw_applic_comp.Information_item_id, &ppdm_sw_applic_comp.Info_item_subtype, &ppdm_sw_applic_comp.Instrument_id, &ppdm_sw_applic_comp.Interest_set_id, &ppdm_sw_applic_comp.Interest_set_seq_no, &ppdm_sw_applic_comp.Jurisdiction, &ppdm_sw_applic_comp.Land_right_id, &ppdm_sw_applic_comp.Land_right_subtype, &ppdm_sw_applic_comp.Land_sale_number, &ppdm_sw_applic_comp.Language, &ppdm_sw_applic_comp.Lithology_log_id, &ppdm_sw_applic_comp.Lith_log_source, &ppdm_sw_applic_comp.Notification_id, &ppdm_sw_applic_comp.Obligation_id, &ppdm_sw_applic_comp.Obligation_seq_no, &ppdm_sw_applic_comp.Paleo_summary_id, &ppdm_sw_applic_comp.Pden_id, &ppdm_sw_applic_comp.Pden_source, &ppdm_sw_applic_comp.Pden_subtype, &ppdm_sw_applic_comp.Physical_item_id, &ppdm_sw_applic_comp.Pool_id, &ppdm_sw_applic_comp.Ppdm_guid, &ppdm_sw_applic_comp.Ppdm_system_id, &ppdm_sw_applic_comp.Ppdm_table_name, &ppdm_sw_applic_comp.Product_type, &ppdm_sw_applic_comp.Prod_string_source, &ppdm_sw_applic_comp.Project_id, &ppdm_sw_applic_comp.Pr_str_form_obs_no, &ppdm_sw_applic_comp.Rate_schedule_id, &ppdm_sw_applic_comp.Recovery_obs_no, &ppdm_sw_applic_comp.Referenced_guid, &ppdm_sw_applic_comp.Referenced_system_id, &ppdm_sw_applic_comp.Referenced_table_name, &ppdm_sw_applic_comp.Remark, &ppdm_sw_applic_comp.Reported_ind, &ppdm_sw_applic_comp.Resent_id, &ppdm_sw_applic_comp.Reserve_class_id, &ppdm_sw_applic_comp.Run_num, &ppdm_sw_applic_comp.Sample_date, &ppdm_sw_applic_comp.Sample_id, &ppdm_sw_applic_comp.Sample_source, &ppdm_sw_applic_comp.Seis_set_id, &ppdm_sw_applic_comp.Seis_set_subtype, &ppdm_sw_applic_comp.Sf_subtype, &ppdm_sw_applic_comp.Source, &ppdm_sw_applic_comp.Spatial_description_id, &ppdm_sw_applic_comp.Spatial_obs_no, &ppdm_sw_applic_comp.Step_seq_no, &ppdm_sw_applic_comp.Store_id, &ppdm_sw_applic_comp.Store_structure_obs_no, &ppdm_sw_applic_comp.Strat_name_set_id, &ppdm_sw_applic_comp.Strat_unit_id, &ppdm_sw_applic_comp.String_id, &ppdm_sw_applic_comp.Support_facility_id, &ppdm_sw_applic_comp.Test_num, &ppdm_sw_applic_comp.Test_source, &ppdm_sw_applic_comp.Test_type, &ppdm_sw_applic_comp.Thesaurus_id, &ppdm_sw_applic_comp.Thesaurus_word_id, &ppdm_sw_applic_comp.Uwi, &ppdm_sw_applic_comp.Well_activity_id, &ppdm_sw_applic_comp.Well_activity_set_id, &ppdm_sw_applic_comp.Well_activity_source, &ppdm_sw_applic_comp.Well_activity_type_id, &ppdm_sw_applic_comp.Well_core_desc_id, &ppdm_sw_applic_comp.Well_test_id, &ppdm_sw_applic_comp.Well_test_recovery_id, &ppdm_sw_applic_comp.Well_test_source, &ppdm_sw_applic_comp.Well_test_type, &ppdm_sw_applic_comp.Work_order_id, &ppdm_sw_applic_comp.Row_changed_by, &ppdm_sw_applic_comp.Row_changed_date, &ppdm_sw_applic_comp.Row_created_by, &ppdm_sw_applic_comp.Row_created_date, &ppdm_sw_applic_comp.Row_effective_date, &ppdm_sw_applic_comp.Row_expiry_date, &ppdm_sw_applic_comp.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_sw_applic_comp to result
		result = append(result, ppdm_sw_applic_comp)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmSwApplicComp(c *fiber.Ctx) error {
	var ppdm_sw_applic_comp dto.Ppdm_sw_applic_comp

	setDefaults(&ppdm_sw_applic_comp)

	if err := json.Unmarshal(c.Body(), &ppdm_sw_applic_comp); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_sw_applic_comp values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103, :104, :105, :106, :107, :108, :109, :110, :111, :112, :113, :114, :115, :116, :117, :118, :119)")
	if err != nil {
		return err
	}
	ppdm_sw_applic_comp.Row_created_date = formatDate(ppdm_sw_applic_comp.Row_created_date)
	ppdm_sw_applic_comp.Row_changed_date = formatDate(ppdm_sw_applic_comp.Row_changed_date)
	ppdm_sw_applic_comp.Effective_date = formatDateString(ppdm_sw_applic_comp.Effective_date)
	ppdm_sw_applic_comp.Expiry_date = formatDateString(ppdm_sw_applic_comp.Expiry_date)
	ppdm_sw_applic_comp.Sample_date = formatDateString(ppdm_sw_applic_comp.Sample_date)
	ppdm_sw_applic_comp.Row_effective_date = formatDateString(ppdm_sw_applic_comp.Row_effective_date)
	ppdm_sw_applic_comp.Row_expiry_date = formatDateString(ppdm_sw_applic_comp.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_sw_applic_comp.Sw_application_id, ppdm_sw_applic_comp.Component_obs_no, ppdm_sw_applic_comp.Active_ind, ppdm_sw_applic_comp.Activity_obs_no, ppdm_sw_applic_comp.Analysis_id, ppdm_sw_applic_comp.Anl_source, ppdm_sw_applic_comp.Application_id, ppdm_sw_applic_comp.Area_id, ppdm_sw_applic_comp.Area_type, ppdm_sw_applic_comp.Authority_id, ppdm_sw_applic_comp.Ba_organization_id, ppdm_sw_applic_comp.Ba_organization_seq_no, ppdm_sw_applic_comp.Business_associate_id, ppdm_sw_applic_comp.Catalogue_additive_id, ppdm_sw_applic_comp.Catalogue_equip_id, ppdm_sw_applic_comp.Classification_system_id, ppdm_sw_applic_comp.Class_level_id, ppdm_sw_applic_comp.Component_type, ppdm_sw_applic_comp.Consent_id, ppdm_sw_applic_comp.Consult_id, ppdm_sw_applic_comp.Contest_id, ppdm_sw_applic_comp.Contract_id, ppdm_sw_applic_comp.Core_id, ppdm_sw_applic_comp.Core_type, ppdm_sw_applic_comp.Description_obs_no, ppdm_sw_applic_comp.Ecozone_id, ppdm_sw_applic_comp.Effective_date, ppdm_sw_applic_comp.Employee_ba_id, ppdm_sw_applic_comp.Employee_obs_no, ppdm_sw_applic_comp.Employer_ba_id, ppdm_sw_applic_comp.Entitlement_id, ppdm_sw_applic_comp.Equipment_id, ppdm_sw_applic_comp.Expiry_date, ppdm_sw_applic_comp.Facility_id, ppdm_sw_applic_comp.Facility_type, ppdm_sw_applic_comp.Field_id, ppdm_sw_applic_comp.Field_station_id, ppdm_sw_applic_comp.Finance_id, ppdm_sw_applic_comp.Fossil_id, ppdm_sw_applic_comp.Incident_id, ppdm_sw_applic_comp.Incident_set_id, ppdm_sw_applic_comp.Incident_type, ppdm_sw_applic_comp.Information_item_id, ppdm_sw_applic_comp.Info_item_subtype, ppdm_sw_applic_comp.Instrument_id, ppdm_sw_applic_comp.Interest_set_id, ppdm_sw_applic_comp.Interest_set_seq_no, ppdm_sw_applic_comp.Jurisdiction, ppdm_sw_applic_comp.Land_right_id, ppdm_sw_applic_comp.Land_right_subtype, ppdm_sw_applic_comp.Land_sale_number, ppdm_sw_applic_comp.Language, ppdm_sw_applic_comp.Lithology_log_id, ppdm_sw_applic_comp.Lith_log_source, ppdm_sw_applic_comp.Notification_id, ppdm_sw_applic_comp.Obligation_id, ppdm_sw_applic_comp.Obligation_seq_no, ppdm_sw_applic_comp.Paleo_summary_id, ppdm_sw_applic_comp.Pden_id, ppdm_sw_applic_comp.Pden_source, ppdm_sw_applic_comp.Pden_subtype, ppdm_sw_applic_comp.Physical_item_id, ppdm_sw_applic_comp.Pool_id, ppdm_sw_applic_comp.Ppdm_guid, ppdm_sw_applic_comp.Ppdm_system_id, ppdm_sw_applic_comp.Ppdm_table_name, ppdm_sw_applic_comp.Product_type, ppdm_sw_applic_comp.Prod_string_source, ppdm_sw_applic_comp.Project_id, ppdm_sw_applic_comp.Pr_str_form_obs_no, ppdm_sw_applic_comp.Rate_schedule_id, ppdm_sw_applic_comp.Recovery_obs_no, ppdm_sw_applic_comp.Referenced_guid, ppdm_sw_applic_comp.Referenced_system_id, ppdm_sw_applic_comp.Referenced_table_name, ppdm_sw_applic_comp.Remark, ppdm_sw_applic_comp.Reported_ind, ppdm_sw_applic_comp.Resent_id, ppdm_sw_applic_comp.Reserve_class_id, ppdm_sw_applic_comp.Run_num, ppdm_sw_applic_comp.Sample_date, ppdm_sw_applic_comp.Sample_id, ppdm_sw_applic_comp.Sample_source, ppdm_sw_applic_comp.Seis_set_id, ppdm_sw_applic_comp.Seis_set_subtype, ppdm_sw_applic_comp.Sf_subtype, ppdm_sw_applic_comp.Source, ppdm_sw_applic_comp.Spatial_description_id, ppdm_sw_applic_comp.Spatial_obs_no, ppdm_sw_applic_comp.Step_seq_no, ppdm_sw_applic_comp.Store_id, ppdm_sw_applic_comp.Store_structure_obs_no, ppdm_sw_applic_comp.Strat_name_set_id, ppdm_sw_applic_comp.Strat_unit_id, ppdm_sw_applic_comp.String_id, ppdm_sw_applic_comp.Support_facility_id, ppdm_sw_applic_comp.Test_num, ppdm_sw_applic_comp.Test_source, ppdm_sw_applic_comp.Test_type, ppdm_sw_applic_comp.Thesaurus_id, ppdm_sw_applic_comp.Thesaurus_word_id, ppdm_sw_applic_comp.Uwi, ppdm_sw_applic_comp.Well_activity_id, ppdm_sw_applic_comp.Well_activity_set_id, ppdm_sw_applic_comp.Well_activity_source, ppdm_sw_applic_comp.Well_activity_type_id, ppdm_sw_applic_comp.Well_core_desc_id, ppdm_sw_applic_comp.Well_test_id, ppdm_sw_applic_comp.Well_test_recovery_id, ppdm_sw_applic_comp.Well_test_source, ppdm_sw_applic_comp.Well_test_type, ppdm_sw_applic_comp.Work_order_id, ppdm_sw_applic_comp.Row_changed_by, ppdm_sw_applic_comp.Row_changed_date, ppdm_sw_applic_comp.Row_created_by, ppdm_sw_applic_comp.Row_created_date, ppdm_sw_applic_comp.Row_effective_date, ppdm_sw_applic_comp.Row_expiry_date, ppdm_sw_applic_comp.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmSwApplicComp(c *fiber.Ctx) error {
	var ppdm_sw_applic_comp dto.Ppdm_sw_applic_comp

	setDefaults(&ppdm_sw_applic_comp)

	if err := json.Unmarshal(c.Body(), &ppdm_sw_applic_comp); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_sw_applic_comp.Sw_application_id = id

    if ppdm_sw_applic_comp.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_sw_applic_comp set component_obs_no = :1, active_ind = :2, activity_obs_no = :3, analysis_id = :4, anl_source = :5, application_id = :6, area_id = :7, area_type = :8, authority_id = :9, ba_organization_id = :10, ba_organization_seq_no = :11, business_associate_id = :12, catalogue_additive_id = :13, catalogue_equip_id = :14, classification_system_id = :15, class_level_id = :16, component_type = :17, consent_id = :18, consult_id = :19, contest_id = :20, contract_id = :21, core_id = :22, core_type = :23, description_obs_no = :24, ecozone_id = :25, effective_date = :26, employee_ba_id = :27, employee_obs_no = :28, employer_ba_id = :29, entitlement_id = :30, equipment_id = :31, expiry_date = :32, facility_id = :33, facility_type = :34, field_id = :35, field_station_id = :36, finance_id = :37, fossil_id = :38, incident_id = :39, incident_set_id = :40, incident_type = :41, information_item_id = :42, info_item_subtype = :43, instrument_id = :44, interest_set_id = :45, interest_set_seq_no = :46, jurisdiction = :47, land_right_id = :48, land_right_subtype = :49, land_sale_number = :50, language = :51, lithology_log_id = :52, lith_log_source = :53, notification_id = :54, obligation_id = :55, obligation_seq_no = :56, paleo_summary_id = :57, pden_id = :58, pden_source = :59, pden_subtype = :60, physical_item_id = :61, pool_id = :62, ppdm_guid = :63, ppdm_system_id = :64, ppdm_table_name = :65, product_type = :66, prod_string_source = :67, project_id = :68, pr_str_form_obs_no = :69, rate_schedule_id = :70, recovery_obs_no = :71, referenced_guid = :72, referenced_system_id = :73, referenced_table_name = :74, remark = :75, reported_ind = :76, resent_id = :77, reserve_class_id = :78, run_num = :79, sample_date = :80, sample_id = :81, sample_source = :82, seis_set_id = :83, seis_set_subtype = :84, sf_subtype = :85, source = :86, spatial_description_id = :87, spatial_obs_no = :88, step_seq_no = :89, store_id = :90, store_structure_obs_no = :91, strat_name_set_id = :92, strat_unit_id = :93, string_id = :94, support_facility_id = :95, test_num = :96, test_source = :97, test_type = :98, thesaurus_id = :99, thesaurus_word_id = :100, uwi = :101, well_activity_id = :102, well_activity_set_id = :103, well_activity_source = :104, well_activity_type_id = :105, well_core_desc_id = :106, well_test_id = :107, well_test_recovery_id = :108, well_test_source = :109, well_test_type = :110, work_order_id = :111, row_changed_by = :112, row_changed_date = :113, row_created_by = :114, row_effective_date = :115, row_expiry_date = :116, row_quality = :117 where sw_application_id = :119")
	if err != nil {
		return err
	}

	ppdm_sw_applic_comp.Row_changed_date = formatDate(ppdm_sw_applic_comp.Row_changed_date)
	ppdm_sw_applic_comp.Effective_date = formatDateString(ppdm_sw_applic_comp.Effective_date)
	ppdm_sw_applic_comp.Expiry_date = formatDateString(ppdm_sw_applic_comp.Expiry_date)
	ppdm_sw_applic_comp.Sample_date = formatDateString(ppdm_sw_applic_comp.Sample_date)
	ppdm_sw_applic_comp.Row_effective_date = formatDateString(ppdm_sw_applic_comp.Row_effective_date)
	ppdm_sw_applic_comp.Row_expiry_date = formatDateString(ppdm_sw_applic_comp.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_sw_applic_comp.Component_obs_no, ppdm_sw_applic_comp.Active_ind, ppdm_sw_applic_comp.Activity_obs_no, ppdm_sw_applic_comp.Analysis_id, ppdm_sw_applic_comp.Anl_source, ppdm_sw_applic_comp.Application_id, ppdm_sw_applic_comp.Area_id, ppdm_sw_applic_comp.Area_type, ppdm_sw_applic_comp.Authority_id, ppdm_sw_applic_comp.Ba_organization_id, ppdm_sw_applic_comp.Ba_organization_seq_no, ppdm_sw_applic_comp.Business_associate_id, ppdm_sw_applic_comp.Catalogue_additive_id, ppdm_sw_applic_comp.Catalogue_equip_id, ppdm_sw_applic_comp.Classification_system_id, ppdm_sw_applic_comp.Class_level_id, ppdm_sw_applic_comp.Component_type, ppdm_sw_applic_comp.Consent_id, ppdm_sw_applic_comp.Consult_id, ppdm_sw_applic_comp.Contest_id, ppdm_sw_applic_comp.Contract_id, ppdm_sw_applic_comp.Core_id, ppdm_sw_applic_comp.Core_type, ppdm_sw_applic_comp.Description_obs_no, ppdm_sw_applic_comp.Ecozone_id, ppdm_sw_applic_comp.Effective_date, ppdm_sw_applic_comp.Employee_ba_id, ppdm_sw_applic_comp.Employee_obs_no, ppdm_sw_applic_comp.Employer_ba_id, ppdm_sw_applic_comp.Entitlement_id, ppdm_sw_applic_comp.Equipment_id, ppdm_sw_applic_comp.Expiry_date, ppdm_sw_applic_comp.Facility_id, ppdm_sw_applic_comp.Facility_type, ppdm_sw_applic_comp.Field_id, ppdm_sw_applic_comp.Field_station_id, ppdm_sw_applic_comp.Finance_id, ppdm_sw_applic_comp.Fossil_id, ppdm_sw_applic_comp.Incident_id, ppdm_sw_applic_comp.Incident_set_id, ppdm_sw_applic_comp.Incident_type, ppdm_sw_applic_comp.Information_item_id, ppdm_sw_applic_comp.Info_item_subtype, ppdm_sw_applic_comp.Instrument_id, ppdm_sw_applic_comp.Interest_set_id, ppdm_sw_applic_comp.Interest_set_seq_no, ppdm_sw_applic_comp.Jurisdiction, ppdm_sw_applic_comp.Land_right_id, ppdm_sw_applic_comp.Land_right_subtype, ppdm_sw_applic_comp.Land_sale_number, ppdm_sw_applic_comp.Language, ppdm_sw_applic_comp.Lithology_log_id, ppdm_sw_applic_comp.Lith_log_source, ppdm_sw_applic_comp.Notification_id, ppdm_sw_applic_comp.Obligation_id, ppdm_sw_applic_comp.Obligation_seq_no, ppdm_sw_applic_comp.Paleo_summary_id, ppdm_sw_applic_comp.Pden_id, ppdm_sw_applic_comp.Pden_source, ppdm_sw_applic_comp.Pden_subtype, ppdm_sw_applic_comp.Physical_item_id, ppdm_sw_applic_comp.Pool_id, ppdm_sw_applic_comp.Ppdm_guid, ppdm_sw_applic_comp.Ppdm_system_id, ppdm_sw_applic_comp.Ppdm_table_name, ppdm_sw_applic_comp.Product_type, ppdm_sw_applic_comp.Prod_string_source, ppdm_sw_applic_comp.Project_id, ppdm_sw_applic_comp.Pr_str_form_obs_no, ppdm_sw_applic_comp.Rate_schedule_id, ppdm_sw_applic_comp.Recovery_obs_no, ppdm_sw_applic_comp.Referenced_guid, ppdm_sw_applic_comp.Referenced_system_id, ppdm_sw_applic_comp.Referenced_table_name, ppdm_sw_applic_comp.Remark, ppdm_sw_applic_comp.Reported_ind, ppdm_sw_applic_comp.Resent_id, ppdm_sw_applic_comp.Reserve_class_id, ppdm_sw_applic_comp.Run_num, ppdm_sw_applic_comp.Sample_date, ppdm_sw_applic_comp.Sample_id, ppdm_sw_applic_comp.Sample_source, ppdm_sw_applic_comp.Seis_set_id, ppdm_sw_applic_comp.Seis_set_subtype, ppdm_sw_applic_comp.Sf_subtype, ppdm_sw_applic_comp.Source, ppdm_sw_applic_comp.Spatial_description_id, ppdm_sw_applic_comp.Spatial_obs_no, ppdm_sw_applic_comp.Step_seq_no, ppdm_sw_applic_comp.Store_id, ppdm_sw_applic_comp.Store_structure_obs_no, ppdm_sw_applic_comp.Strat_name_set_id, ppdm_sw_applic_comp.Strat_unit_id, ppdm_sw_applic_comp.String_id, ppdm_sw_applic_comp.Support_facility_id, ppdm_sw_applic_comp.Test_num, ppdm_sw_applic_comp.Test_source, ppdm_sw_applic_comp.Test_type, ppdm_sw_applic_comp.Thesaurus_id, ppdm_sw_applic_comp.Thesaurus_word_id, ppdm_sw_applic_comp.Uwi, ppdm_sw_applic_comp.Well_activity_id, ppdm_sw_applic_comp.Well_activity_set_id, ppdm_sw_applic_comp.Well_activity_source, ppdm_sw_applic_comp.Well_activity_type_id, ppdm_sw_applic_comp.Well_core_desc_id, ppdm_sw_applic_comp.Well_test_id, ppdm_sw_applic_comp.Well_test_recovery_id, ppdm_sw_applic_comp.Well_test_source, ppdm_sw_applic_comp.Well_test_type, ppdm_sw_applic_comp.Work_order_id, ppdm_sw_applic_comp.Row_changed_by, ppdm_sw_applic_comp.Row_changed_date, ppdm_sw_applic_comp.Row_created_by, ppdm_sw_applic_comp.Row_effective_date, ppdm_sw_applic_comp.Row_expiry_date, ppdm_sw_applic_comp.Row_quality, ppdm_sw_applic_comp.Sw_application_id)
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

func PatchPpdmSwApplicComp(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_sw_applic_comp set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "sample_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where sw_application_id = :id"

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

func DeletePpdmSwApplicComp(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_sw_applic_comp dto.Ppdm_sw_applic_comp
	ppdm_sw_applic_comp.Sw_application_id = id

	stmt, err := db.Prepare("delete from ppdm_sw_applic_comp where sw_application_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_sw_applic_comp.Sw_application_id)
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


