package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellActivityComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_activity_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_activity_component

	for rows.Next() {
		var well_activity_component dto.Well_activity_component
		if err := rows.Scan(&well_activity_component.Uwi, &well_activity_component.Activity_source, &well_activity_component.Activity_obs_no, &well_activity_component.Component_obs_no, &well_activity_component.Active_ind, &well_activity_component.Analysis_id, &well_activity_component.Application_id, &well_activity_component.Area_id, &well_activity_component.Area_type, &well_activity_component.Authority_id, &well_activity_component.Ba_organization_id, &well_activity_component.Ba_organization_seq_no, &well_activity_component.Ba_service_provided_by, &well_activity_component.Ba_service_seq_no, &well_activity_component.Bit_interval_obs_no, &well_activity_component.Bit_source, &well_activity_component.Business_associate_id, &well_activity_component.Catalogue_additive_id, &well_activity_component.Catalogue_equip_id, &well_activity_component.Cement_obs_no, &well_activity_component.Classification_system_id, &well_activity_component.Class_level_id, &well_activity_component.Component_type, &well_activity_component.Consent_id, &well_activity_component.Consult_id, &well_activity_component.Contest_id, &well_activity_component.Contract_id, &well_activity_component.Crew_company_ba_id, &well_activity_component.Crew_id, &well_activity_component.Dir_survey_id, &well_activity_component.Dir_survey_source, &well_activity_component.Drill_bit_period_obs_no, &well_activity_component.Drill_check_obs_no, &well_activity_component.Ecozone_id, &well_activity_component.Effective_date, &well_activity_component.Employee_ba_id, &well_activity_component.Employee_obs_no, &well_activity_component.Employer_ba_id, &well_activity_component.Entitlement_id, &well_activity_component.Equipment_obs_no, &well_activity_component.Expiry_date, &well_activity_component.Facility_id, &well_activity_component.Facility_type, &well_activity_component.Field_id, &well_activity_component.Field_station_id, &well_activity_component.Finance_id, &well_activity_component.Fossil_id, &well_activity_component.Incident_id, &well_activity_component.Incident_set_id, &well_activity_component.Incident_type, &well_activity_component.Information_item_id, &well_activity_component.Info_item_subtype, &well_activity_component.Instrument_id, &well_activity_component.Interest_set_id, &well_activity_component.Interest_set_seq_no, &well_activity_component.Jurisdiction, &well_activity_component.Land_right_id, &well_activity_component.Land_right_subtype, &well_activity_component.Land_sale_number, &well_activity_component.Language, &well_activity_component.Lithology_log_id, &well_activity_component.Lith_log_source, &well_activity_component.Maint_equipment_id, &well_activity_component.Maint_equip_maint_id, &well_activity_component.Notification_id, &well_activity_component.Obligation_id, &well_activity_component.Obligation_seq_no, &well_activity_component.Paleo_summary_id, &well_activity_component.Pden_activity_type, &well_activity_component.Pden_amendment_seq_no, &well_activity_component.Pden_id, &well_activity_component.Pden_period_id, &well_activity_component.Pden_period_type, &well_activity_component.Pden_source, &well_activity_component.Pden_subtype, &well_activity_component.Pden_volume_date, &well_activity_component.Pden_volume_date_desc, &well_activity_component.Pden_volume_method, &well_activity_component.Perforation_obs_no, &well_activity_component.Perforation_source, &well_activity_component.Period_obs_no, &well_activity_component.Physical_item_id, &well_activity_component.Plugback_obs_no, &well_activity_component.Plugback_source, &well_activity_component.Pool_id, &well_activity_component.Ppdm_guid, &well_activity_component.Ppdm_system_id, &well_activity_component.Ppdm_table_name, &well_activity_component.Product_type, &well_activity_component.Prod_string_id, &well_activity_component.Prod_string_source, &well_activity_component.Project_id, &well_activity_component.Pr_str_form_obs_no, &well_activity_component.Rate_schedule_id, &well_activity_component.Referenced_guid, &well_activity_component.Referenced_system_id, &well_activity_component.Referenced_table_name, &well_activity_component.Remark, &well_activity_component.Resent_id, &well_activity_component.Reserve_class_id, &well_activity_component.Sample_anal_source, &well_activity_component.Seis_set_id, &well_activity_component.Seis_set_subtype, &well_activity_component.Sf_subtype, &well_activity_component.Source, &well_activity_component.Spatial_description_id, &well_activity_component.Spatial_obs_no, &well_activity_component.Store_id, &well_activity_component.Store_structure_obs_no, &well_activity_component.Strat_name_set_id, &well_activity_component.Strat_unit_id, &well_activity_component.Support_facility_id, &well_activity_component.Sw_application_id, &well_activity_component.Thesaurus_id, &well_activity_component.Thesaurus_word_id, &well_activity_component.Tubing_obs_no, &well_activity_component.Tubing_source, &well_activity_component.Tubing_type, &well_activity_component.Well_activity_set_id, &well_activity_component.Well_activity_type_id, &well_activity_component.Well_core_id, &well_activity_component.Well_core_source, &well_activity_component.Well_log_id, &well_activity_component.Well_log_source, &well_activity_component.Well_set_id, &well_activity_component.Well_test_num, &well_activity_component.Well_test_run_num, &well_activity_component.Well_test_source, &well_activity_component.Well_test_type, &well_activity_component.Work_order_id, &well_activity_component.Row_changed_by, &well_activity_component.Row_changed_date, &well_activity_component.Row_created_by, &well_activity_component.Row_created_date, &well_activity_component.Row_effective_date, &well_activity_component.Row_expiry_date, &well_activity_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_activity_component to result
		result = append(result, well_activity_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellActivityComponent(c *fiber.Ctx) error {
	var well_activity_component dto.Well_activity_component

	setDefaults(&well_activity_component)

	if err := json.Unmarshal(c.Body(), &well_activity_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_activity_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103, :104, :105, :106, :107, :108, :109, :110, :111, :112, :113, :114, :115, :116, :117, :118, :119, :120, :121, :122, :123, :124, :125, :126, :127, :128, :129, :130, :131, :132, :133, :134, :135, :136, :137)")
	if err != nil {
		return err
	}
	well_activity_component.Row_created_date = formatDate(well_activity_component.Row_created_date)
	well_activity_component.Row_changed_date = formatDate(well_activity_component.Row_changed_date)
	well_activity_component.Effective_date = formatDateString(well_activity_component.Effective_date)
	well_activity_component.Expiry_date = formatDateString(well_activity_component.Expiry_date)
	well_activity_component.Pden_volume_date = formatDateString(well_activity_component.Pden_volume_date)
	well_activity_component.Row_effective_date = formatDateString(well_activity_component.Row_effective_date)
	well_activity_component.Row_expiry_date = formatDateString(well_activity_component.Row_expiry_date)






	rows, err := stmt.Exec(well_activity_component.Uwi, well_activity_component.Activity_source, well_activity_component.Activity_obs_no, well_activity_component.Component_obs_no, well_activity_component.Active_ind, well_activity_component.Analysis_id, well_activity_component.Application_id, well_activity_component.Area_id, well_activity_component.Area_type, well_activity_component.Authority_id, well_activity_component.Ba_organization_id, well_activity_component.Ba_organization_seq_no, well_activity_component.Ba_service_provided_by, well_activity_component.Ba_service_seq_no, well_activity_component.Bit_interval_obs_no, well_activity_component.Bit_source, well_activity_component.Business_associate_id, well_activity_component.Catalogue_additive_id, well_activity_component.Catalogue_equip_id, well_activity_component.Cement_obs_no, well_activity_component.Classification_system_id, well_activity_component.Class_level_id, well_activity_component.Component_type, well_activity_component.Consent_id, well_activity_component.Consult_id, well_activity_component.Contest_id, well_activity_component.Contract_id, well_activity_component.Crew_company_ba_id, well_activity_component.Crew_id, well_activity_component.Dir_survey_id, well_activity_component.Dir_survey_source, well_activity_component.Drill_bit_period_obs_no, well_activity_component.Drill_check_obs_no, well_activity_component.Ecozone_id, well_activity_component.Effective_date, well_activity_component.Employee_ba_id, well_activity_component.Employee_obs_no, well_activity_component.Employer_ba_id, well_activity_component.Entitlement_id, well_activity_component.Equipment_obs_no, well_activity_component.Expiry_date, well_activity_component.Facility_id, well_activity_component.Facility_type, well_activity_component.Field_id, well_activity_component.Field_station_id, well_activity_component.Finance_id, well_activity_component.Fossil_id, well_activity_component.Incident_id, well_activity_component.Incident_set_id, well_activity_component.Incident_type, well_activity_component.Information_item_id, well_activity_component.Info_item_subtype, well_activity_component.Instrument_id, well_activity_component.Interest_set_id, well_activity_component.Interest_set_seq_no, well_activity_component.Jurisdiction, well_activity_component.Land_right_id, well_activity_component.Land_right_subtype, well_activity_component.Land_sale_number, well_activity_component.Language, well_activity_component.Lithology_log_id, well_activity_component.Lith_log_source, well_activity_component.Maint_equipment_id, well_activity_component.Maint_equip_maint_id, well_activity_component.Notification_id, well_activity_component.Obligation_id, well_activity_component.Obligation_seq_no, well_activity_component.Paleo_summary_id, well_activity_component.Pden_activity_type, well_activity_component.Pden_amendment_seq_no, well_activity_component.Pden_id, well_activity_component.Pden_period_id, well_activity_component.Pden_period_type, well_activity_component.Pden_source, well_activity_component.Pden_subtype, well_activity_component.Pden_volume_date, well_activity_component.Pden_volume_date_desc, well_activity_component.Pden_volume_method, well_activity_component.Perforation_obs_no, well_activity_component.Perforation_source, well_activity_component.Period_obs_no, well_activity_component.Physical_item_id, well_activity_component.Plugback_obs_no, well_activity_component.Plugback_source, well_activity_component.Pool_id, well_activity_component.Ppdm_guid, well_activity_component.Ppdm_system_id, well_activity_component.Ppdm_table_name, well_activity_component.Product_type, well_activity_component.Prod_string_id, well_activity_component.Prod_string_source, well_activity_component.Project_id, well_activity_component.Pr_str_form_obs_no, well_activity_component.Rate_schedule_id, well_activity_component.Referenced_guid, well_activity_component.Referenced_system_id, well_activity_component.Referenced_table_name, well_activity_component.Remark, well_activity_component.Resent_id, well_activity_component.Reserve_class_id, well_activity_component.Sample_anal_source, well_activity_component.Seis_set_id, well_activity_component.Seis_set_subtype, well_activity_component.Sf_subtype, well_activity_component.Source, well_activity_component.Spatial_description_id, well_activity_component.Spatial_obs_no, well_activity_component.Store_id, well_activity_component.Store_structure_obs_no, well_activity_component.Strat_name_set_id, well_activity_component.Strat_unit_id, well_activity_component.Support_facility_id, well_activity_component.Sw_application_id, well_activity_component.Thesaurus_id, well_activity_component.Thesaurus_word_id, well_activity_component.Tubing_obs_no, well_activity_component.Tubing_source, well_activity_component.Tubing_type, well_activity_component.Well_activity_set_id, well_activity_component.Well_activity_type_id, well_activity_component.Well_core_id, well_activity_component.Well_core_source, well_activity_component.Well_log_id, well_activity_component.Well_log_source, well_activity_component.Well_set_id, well_activity_component.Well_test_num, well_activity_component.Well_test_run_num, well_activity_component.Well_test_source, well_activity_component.Well_test_type, well_activity_component.Work_order_id, well_activity_component.Row_changed_by, well_activity_component.Row_changed_date, well_activity_component.Row_created_by, well_activity_component.Row_created_date, well_activity_component.Row_effective_date, well_activity_component.Row_expiry_date, well_activity_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellActivityComponent(c *fiber.Ctx) error {
	var well_activity_component dto.Well_activity_component

	setDefaults(&well_activity_component)

	if err := json.Unmarshal(c.Body(), &well_activity_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_activity_component.Uwi = id

    if well_activity_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_activity_component set activity_source = :1, activity_obs_no = :2, component_obs_no = :3, active_ind = :4, analysis_id = :5, application_id = :6, area_id = :7, area_type = :8, authority_id = :9, ba_organization_id = :10, ba_organization_seq_no = :11, ba_service_provided_by = :12, ba_service_seq_no = :13, bit_interval_obs_no = :14, bit_source = :15, business_associate_id = :16, catalogue_additive_id = :17, catalogue_equip_id = :18, cement_obs_no = :19, classification_system_id = :20, class_level_id = :21, component_type = :22, consent_id = :23, consult_id = :24, contest_id = :25, contract_id = :26, crew_company_ba_id = :27, crew_id = :28, dir_survey_id = :29, dir_survey_source = :30, drill_bit_period_obs_no = :31, drill_check_obs_no = :32, ecozone_id = :33, effective_date = :34, employee_ba_id = :35, employee_obs_no = :36, employer_ba_id = :37, entitlement_id = :38, equipment_obs_no = :39, expiry_date = :40, facility_id = :41, facility_type = :42, field_id = :43, field_station_id = :44, finance_id = :45, fossil_id = :46, incident_id = :47, incident_set_id = :48, incident_type = :49, information_item_id = :50, info_item_subtype = :51, instrument_id = :52, interest_set_id = :53, interest_set_seq_no = :54, jurisdiction = :55, land_right_id = :56, land_right_subtype = :57, land_sale_number = :58, language = :59, lithology_log_id = :60, lith_log_source = :61, maint_equipment_id = :62, maint_equip_maint_id = :63, notification_id = :64, obligation_id = :65, obligation_seq_no = :66, paleo_summary_id = :67, pden_activity_type = :68, pden_amendment_seq_no = :69, pden_id = :70, pden_period_id = :71, pden_period_type = :72, pden_source = :73, pden_subtype = :74, pden_volume_date = :75, pden_volume_date_desc = :76, pden_volume_method = :77, perforation_obs_no = :78, perforation_source = :79, period_obs_no = :80, physical_item_id = :81, plugback_obs_no = :82, plugback_source = :83, pool_id = :84, ppdm_guid = :85, ppdm_system_id = :86, ppdm_table_name = :87, product_type = :88, prod_string_id = :89, prod_string_source = :90, project_id = :91, pr_str_form_obs_no = :92, rate_schedule_id = :93, referenced_guid = :94, referenced_system_id = :95, referenced_table_name = :96, remark = :97, resent_id = :98, reserve_class_id = :99, sample_anal_source = :100, seis_set_id = :101, seis_set_subtype = :102, sf_subtype = :103, source = :104, spatial_description_id = :105, spatial_obs_no = :106, store_id = :107, store_structure_obs_no = :108, strat_name_set_id = :109, strat_unit_id = :110, support_facility_id = :111, sw_application_id = :112, thesaurus_id = :113, thesaurus_word_id = :114, tubing_obs_no = :115, tubing_source = :116, tubing_type = :117, well_activity_set_id = :118, well_activity_type_id = :119, well_core_id = :120, well_core_source = :121, well_log_id = :122, well_log_source = :123, well_set_id = :124, well_test_num = :125, well_test_run_num = :126, well_test_source = :127, well_test_type = :128, work_order_id = :129, row_changed_by = :130, row_changed_date = :131, row_created_by = :132, row_effective_date = :133, row_expiry_date = :134, row_quality = :135 where uwi = :137")
	if err != nil {
		return err
	}

	well_activity_component.Row_changed_date = formatDate(well_activity_component.Row_changed_date)
	well_activity_component.Effective_date = formatDateString(well_activity_component.Effective_date)
	well_activity_component.Expiry_date = formatDateString(well_activity_component.Expiry_date)
	well_activity_component.Pden_volume_date = formatDateString(well_activity_component.Pden_volume_date)
	well_activity_component.Row_effective_date = formatDateString(well_activity_component.Row_effective_date)
	well_activity_component.Row_expiry_date = formatDateString(well_activity_component.Row_expiry_date)






	rows, err := stmt.Exec(well_activity_component.Activity_source, well_activity_component.Activity_obs_no, well_activity_component.Component_obs_no, well_activity_component.Active_ind, well_activity_component.Analysis_id, well_activity_component.Application_id, well_activity_component.Area_id, well_activity_component.Area_type, well_activity_component.Authority_id, well_activity_component.Ba_organization_id, well_activity_component.Ba_organization_seq_no, well_activity_component.Ba_service_provided_by, well_activity_component.Ba_service_seq_no, well_activity_component.Bit_interval_obs_no, well_activity_component.Bit_source, well_activity_component.Business_associate_id, well_activity_component.Catalogue_additive_id, well_activity_component.Catalogue_equip_id, well_activity_component.Cement_obs_no, well_activity_component.Classification_system_id, well_activity_component.Class_level_id, well_activity_component.Component_type, well_activity_component.Consent_id, well_activity_component.Consult_id, well_activity_component.Contest_id, well_activity_component.Contract_id, well_activity_component.Crew_company_ba_id, well_activity_component.Crew_id, well_activity_component.Dir_survey_id, well_activity_component.Dir_survey_source, well_activity_component.Drill_bit_period_obs_no, well_activity_component.Drill_check_obs_no, well_activity_component.Ecozone_id, well_activity_component.Effective_date, well_activity_component.Employee_ba_id, well_activity_component.Employee_obs_no, well_activity_component.Employer_ba_id, well_activity_component.Entitlement_id, well_activity_component.Equipment_obs_no, well_activity_component.Expiry_date, well_activity_component.Facility_id, well_activity_component.Facility_type, well_activity_component.Field_id, well_activity_component.Field_station_id, well_activity_component.Finance_id, well_activity_component.Fossil_id, well_activity_component.Incident_id, well_activity_component.Incident_set_id, well_activity_component.Incident_type, well_activity_component.Information_item_id, well_activity_component.Info_item_subtype, well_activity_component.Instrument_id, well_activity_component.Interest_set_id, well_activity_component.Interest_set_seq_no, well_activity_component.Jurisdiction, well_activity_component.Land_right_id, well_activity_component.Land_right_subtype, well_activity_component.Land_sale_number, well_activity_component.Language, well_activity_component.Lithology_log_id, well_activity_component.Lith_log_source, well_activity_component.Maint_equipment_id, well_activity_component.Maint_equip_maint_id, well_activity_component.Notification_id, well_activity_component.Obligation_id, well_activity_component.Obligation_seq_no, well_activity_component.Paleo_summary_id, well_activity_component.Pden_activity_type, well_activity_component.Pden_amendment_seq_no, well_activity_component.Pden_id, well_activity_component.Pden_period_id, well_activity_component.Pden_period_type, well_activity_component.Pden_source, well_activity_component.Pden_subtype, well_activity_component.Pden_volume_date, well_activity_component.Pden_volume_date_desc, well_activity_component.Pden_volume_method, well_activity_component.Perforation_obs_no, well_activity_component.Perforation_source, well_activity_component.Period_obs_no, well_activity_component.Physical_item_id, well_activity_component.Plugback_obs_no, well_activity_component.Plugback_source, well_activity_component.Pool_id, well_activity_component.Ppdm_guid, well_activity_component.Ppdm_system_id, well_activity_component.Ppdm_table_name, well_activity_component.Product_type, well_activity_component.Prod_string_id, well_activity_component.Prod_string_source, well_activity_component.Project_id, well_activity_component.Pr_str_form_obs_no, well_activity_component.Rate_schedule_id, well_activity_component.Referenced_guid, well_activity_component.Referenced_system_id, well_activity_component.Referenced_table_name, well_activity_component.Remark, well_activity_component.Resent_id, well_activity_component.Reserve_class_id, well_activity_component.Sample_anal_source, well_activity_component.Seis_set_id, well_activity_component.Seis_set_subtype, well_activity_component.Sf_subtype, well_activity_component.Source, well_activity_component.Spatial_description_id, well_activity_component.Spatial_obs_no, well_activity_component.Store_id, well_activity_component.Store_structure_obs_no, well_activity_component.Strat_name_set_id, well_activity_component.Strat_unit_id, well_activity_component.Support_facility_id, well_activity_component.Sw_application_id, well_activity_component.Thesaurus_id, well_activity_component.Thesaurus_word_id, well_activity_component.Tubing_obs_no, well_activity_component.Tubing_source, well_activity_component.Tubing_type, well_activity_component.Well_activity_set_id, well_activity_component.Well_activity_type_id, well_activity_component.Well_core_id, well_activity_component.Well_core_source, well_activity_component.Well_log_id, well_activity_component.Well_log_source, well_activity_component.Well_set_id, well_activity_component.Well_test_num, well_activity_component.Well_test_run_num, well_activity_component.Well_test_source, well_activity_component.Well_test_type, well_activity_component.Work_order_id, well_activity_component.Row_changed_by, well_activity_component.Row_changed_date, well_activity_component.Row_created_by, well_activity_component.Row_effective_date, well_activity_component.Row_expiry_date, well_activity_component.Row_quality, well_activity_component.Uwi)
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

func PatchWellActivityComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_activity_component set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "pden_volume_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where uwi = :id"

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

func DeleteWellActivityComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_activity_component dto.Well_activity_component
	well_activity_component.Uwi = id

	stmt, err := db.Prepare("delete from well_activity_component where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_activity_component.Uwi)
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


