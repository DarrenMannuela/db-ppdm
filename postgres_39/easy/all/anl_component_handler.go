package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_component

	for rows.Next() {
		var anl_component dto.Anl_component
		if err := rows.Scan(&anl_component.Analysis_id, &anl_component.Anl_source, &anl_component.Component_obs_no, &anl_component.Active_ind, &anl_component.Activity_obs_no, &anl_component.Application_id, &anl_component.Area_id, &anl_component.Area_type, &anl_component.Authority_id, &anl_component.Ba_organization_id, &anl_component.Ba_organization_seq_no, &anl_component.Business_associate_id, &anl_component.Catalogue_additive_id, &anl_component.Catalogue_equip_id, &anl_component.Classification_system_id, &anl_component.Class_level_id, &anl_component.Component_type, &anl_component.Consent_id, &anl_component.Consult_id, &anl_component.Contest_id, &anl_component.Contract_id, &anl_component.Core_id, &anl_component.Core_type, &anl_component.Description_obs_no, &anl_component.Ecozone_id, &anl_component.Effective_date, &anl_component.Employee_ba_id, &anl_component.Employee_obs_no, &anl_component.Employer_ba_id, &anl_component.Entitlement_id, &anl_component.Equipment_id, &anl_component.Expiry_date, &anl_component.Facility_id, &anl_component.Facility_type, &anl_component.Field_id, &anl_component.Field_station_id, &anl_component.Finance_id, &anl_component.Fossil_id, &anl_component.Incident_id, &anl_component.Incident_set_id, &anl_component.Incident_type, &anl_component.Information_item_id, &anl_component.Info_item_subtype, &anl_component.Instrument_id, &anl_component.Interest_set_id, &anl_component.Interest_set_seq_no, &anl_component.Jurisdiction, &anl_component.Land_right_id, &anl_component.Land_right_subtype, &anl_component.Land_sale_number, &anl_component.Language, &anl_component.Lithology_log_id, &anl_component.Lith_log_source, &anl_component.Notification_id, &anl_component.Obligation_id, &anl_component.Obligation_seq_no, &anl_component.Paleo_summary_id, &anl_component.Pden_id, &anl_component.Pden_source, &anl_component.Pden_subtype, &anl_component.Physical_item_id, &anl_component.Pool_id, &anl_component.Ppdm_guid, &anl_component.Ppdm_system_id, &anl_component.Ppdm_table_name, &anl_component.Product_type, &anl_component.Prod_string_source, &anl_component.Project_id, &anl_component.Pr_str_form_obs_no, &anl_component.Rate_schedule_id, &anl_component.Recovery_obs_no, &anl_component.Referenced_guid, &anl_component.Referenced_system_id, &anl_component.Referenced_table_name, &anl_component.Remark, &anl_component.Reported_ind, &anl_component.Resent_id, &anl_component.Reserve_class_id, &anl_component.Run_num, &anl_component.Sample_date, &anl_component.Sample_id, &anl_component.Sample_source, &anl_component.Seis_set_id, &anl_component.Seis_set_subtype, &anl_component.Sf_subtype, &anl_component.Source, &anl_component.Spatial_description_id, &anl_component.Spatial_obs_no, &anl_component.Step_seq_no, &anl_component.Store_id, &anl_component.Store_structure_obs_no, &anl_component.Strat_name_set_id, &anl_component.Strat_unit_id, &anl_component.String_id, &anl_component.Support_facility_id, &anl_component.Test_num, &anl_component.Test_source, &anl_component.Test_type, &anl_component.Thesaurus_id, &anl_component.Thesaurus_word_id, &anl_component.Uwi, &anl_component.Well_activity_id, &anl_component.Well_activity_set_id, &anl_component.Well_activity_source, &anl_component.Well_activity_type_id, &anl_component.Well_core_desc_id, &anl_component.Well_test_id, &anl_component.Well_test_recovery_id, &anl_component.Well_test_source, &anl_component.Well_test_type, &anl_component.Work_order_id, &anl_component.Row_changed_by, &anl_component.Row_changed_date, &anl_component.Row_created_by, &anl_component.Row_created_date, &anl_component.Row_effective_date, &anl_component.Row_expiry_date, &anl_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_component to result
		result = append(result, anl_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlComponent(c *fiber.Ctx) error {
	var anl_component dto.Anl_component

	setDefaults(&anl_component)

	if err := json.Unmarshal(c.Body(), &anl_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103, :104, :105, :106, :107, :108, :109, :110, :111, :112, :113, :114, :115, :116, :117, :118)")
	if err != nil {
		return err
	}
	anl_component.Row_created_date = formatDate(anl_component.Row_created_date)
	anl_component.Row_changed_date = formatDate(anl_component.Row_changed_date)
	anl_component.Effective_date = formatDateString(anl_component.Effective_date)
	anl_component.Expiry_date = formatDateString(anl_component.Expiry_date)
	anl_component.Sample_date = formatDateString(anl_component.Sample_date)
	anl_component.Row_effective_date = formatDateString(anl_component.Row_effective_date)
	anl_component.Row_expiry_date = formatDateString(anl_component.Row_expiry_date)






	rows, err := stmt.Exec(anl_component.Analysis_id, anl_component.Anl_source, anl_component.Component_obs_no, anl_component.Active_ind, anl_component.Activity_obs_no, anl_component.Application_id, anl_component.Area_id, anl_component.Area_type, anl_component.Authority_id, anl_component.Ba_organization_id, anl_component.Ba_organization_seq_no, anl_component.Business_associate_id, anl_component.Catalogue_additive_id, anl_component.Catalogue_equip_id, anl_component.Classification_system_id, anl_component.Class_level_id, anl_component.Component_type, anl_component.Consent_id, anl_component.Consult_id, anl_component.Contest_id, anl_component.Contract_id, anl_component.Core_id, anl_component.Core_type, anl_component.Description_obs_no, anl_component.Ecozone_id, anl_component.Effective_date, anl_component.Employee_ba_id, anl_component.Employee_obs_no, anl_component.Employer_ba_id, anl_component.Entitlement_id, anl_component.Equipment_id, anl_component.Expiry_date, anl_component.Facility_id, anl_component.Facility_type, anl_component.Field_id, anl_component.Field_station_id, anl_component.Finance_id, anl_component.Fossil_id, anl_component.Incident_id, anl_component.Incident_set_id, anl_component.Incident_type, anl_component.Information_item_id, anl_component.Info_item_subtype, anl_component.Instrument_id, anl_component.Interest_set_id, anl_component.Interest_set_seq_no, anl_component.Jurisdiction, anl_component.Land_right_id, anl_component.Land_right_subtype, anl_component.Land_sale_number, anl_component.Language, anl_component.Lithology_log_id, anl_component.Lith_log_source, anl_component.Notification_id, anl_component.Obligation_id, anl_component.Obligation_seq_no, anl_component.Paleo_summary_id, anl_component.Pden_id, anl_component.Pden_source, anl_component.Pden_subtype, anl_component.Physical_item_id, anl_component.Pool_id, anl_component.Ppdm_guid, anl_component.Ppdm_system_id, anl_component.Ppdm_table_name, anl_component.Product_type, anl_component.Prod_string_source, anl_component.Project_id, anl_component.Pr_str_form_obs_no, anl_component.Rate_schedule_id, anl_component.Recovery_obs_no, anl_component.Referenced_guid, anl_component.Referenced_system_id, anl_component.Referenced_table_name, anl_component.Remark, anl_component.Reported_ind, anl_component.Resent_id, anl_component.Reserve_class_id, anl_component.Run_num, anl_component.Sample_date, anl_component.Sample_id, anl_component.Sample_source, anl_component.Seis_set_id, anl_component.Seis_set_subtype, anl_component.Sf_subtype, anl_component.Source, anl_component.Spatial_description_id, anl_component.Spatial_obs_no, anl_component.Step_seq_no, anl_component.Store_id, anl_component.Store_structure_obs_no, anl_component.Strat_name_set_id, anl_component.Strat_unit_id, anl_component.String_id, anl_component.Support_facility_id, anl_component.Test_num, anl_component.Test_source, anl_component.Test_type, anl_component.Thesaurus_id, anl_component.Thesaurus_word_id, anl_component.Uwi, anl_component.Well_activity_id, anl_component.Well_activity_set_id, anl_component.Well_activity_source, anl_component.Well_activity_type_id, anl_component.Well_core_desc_id, anl_component.Well_test_id, anl_component.Well_test_recovery_id, anl_component.Well_test_source, anl_component.Well_test_type, anl_component.Work_order_id, anl_component.Row_changed_by, anl_component.Row_changed_date, anl_component.Row_created_by, anl_component.Row_created_date, anl_component.Row_effective_date, anl_component.Row_expiry_date, anl_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlComponent(c *fiber.Ctx) error {
	var anl_component dto.Anl_component

	setDefaults(&anl_component)

	if err := json.Unmarshal(c.Body(), &anl_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_component.Analysis_id = id

    if anl_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_component set anl_source = :1, component_obs_no = :2, active_ind = :3, activity_obs_no = :4, application_id = :5, area_id = :6, area_type = :7, authority_id = :8, ba_organization_id = :9, ba_organization_seq_no = :10, business_associate_id = :11, catalogue_additive_id = :12, catalogue_equip_id = :13, classification_system_id = :14, class_level_id = :15, component_type = :16, consent_id = :17, consult_id = :18, contest_id = :19, contract_id = :20, core_id = :21, core_type = :22, description_obs_no = :23, ecozone_id = :24, effective_date = :25, employee_ba_id = :26, employee_obs_no = :27, employer_ba_id = :28, entitlement_id = :29, equipment_id = :30, expiry_date = :31, facility_id = :32, facility_type = :33, field_id = :34, field_station_id = :35, finance_id = :36, fossil_id = :37, incident_id = :38, incident_set_id = :39, incident_type = :40, information_item_id = :41, info_item_subtype = :42, instrument_id = :43, interest_set_id = :44, interest_set_seq_no = :45, jurisdiction = :46, land_right_id = :47, land_right_subtype = :48, land_sale_number = :49, language = :50, lithology_log_id = :51, lith_log_source = :52, notification_id = :53, obligation_id = :54, obligation_seq_no = :55, paleo_summary_id = :56, pden_id = :57, pden_source = :58, pden_subtype = :59, physical_item_id = :60, pool_id = :61, ppdm_guid = :62, ppdm_system_id = :63, ppdm_table_name = :64, product_type = :65, prod_string_source = :66, project_id = :67, pr_str_form_obs_no = :68, rate_schedule_id = :69, recovery_obs_no = :70, referenced_guid = :71, referenced_system_id = :72, referenced_table_name = :73, remark = :74, reported_ind = :75, resent_id = :76, reserve_class_id = :77, run_num = :78, sample_date = :79, sample_id = :80, sample_source = :81, seis_set_id = :82, seis_set_subtype = :83, sf_subtype = :84, source = :85, spatial_description_id = :86, spatial_obs_no = :87, step_seq_no = :88, store_id = :89, store_structure_obs_no = :90, strat_name_set_id = :91, strat_unit_id = :92, string_id = :93, support_facility_id = :94, test_num = :95, test_source = :96, test_type = :97, thesaurus_id = :98, thesaurus_word_id = :99, uwi = :100, well_activity_id = :101, well_activity_set_id = :102, well_activity_source = :103, well_activity_type_id = :104, well_core_desc_id = :105, well_test_id = :106, well_test_recovery_id = :107, well_test_source = :108, well_test_type = :109, work_order_id = :110, row_changed_by = :111, row_changed_date = :112, row_created_by = :113, row_effective_date = :114, row_expiry_date = :115, row_quality = :116 where analysis_id = :118")
	if err != nil {
		return err
	}

	anl_component.Row_changed_date = formatDate(anl_component.Row_changed_date)
	anl_component.Effective_date = formatDateString(anl_component.Effective_date)
	anl_component.Expiry_date = formatDateString(anl_component.Expiry_date)
	anl_component.Sample_date = formatDateString(anl_component.Sample_date)
	anl_component.Row_effective_date = formatDateString(anl_component.Row_effective_date)
	anl_component.Row_expiry_date = formatDateString(anl_component.Row_expiry_date)






	rows, err := stmt.Exec(anl_component.Anl_source, anl_component.Component_obs_no, anl_component.Active_ind, anl_component.Activity_obs_no, anl_component.Application_id, anl_component.Area_id, anl_component.Area_type, anl_component.Authority_id, anl_component.Ba_organization_id, anl_component.Ba_organization_seq_no, anl_component.Business_associate_id, anl_component.Catalogue_additive_id, anl_component.Catalogue_equip_id, anl_component.Classification_system_id, anl_component.Class_level_id, anl_component.Component_type, anl_component.Consent_id, anl_component.Consult_id, anl_component.Contest_id, anl_component.Contract_id, anl_component.Core_id, anl_component.Core_type, anl_component.Description_obs_no, anl_component.Ecozone_id, anl_component.Effective_date, anl_component.Employee_ba_id, anl_component.Employee_obs_no, anl_component.Employer_ba_id, anl_component.Entitlement_id, anl_component.Equipment_id, anl_component.Expiry_date, anl_component.Facility_id, anl_component.Facility_type, anl_component.Field_id, anl_component.Field_station_id, anl_component.Finance_id, anl_component.Fossil_id, anl_component.Incident_id, anl_component.Incident_set_id, anl_component.Incident_type, anl_component.Information_item_id, anl_component.Info_item_subtype, anl_component.Instrument_id, anl_component.Interest_set_id, anl_component.Interest_set_seq_no, anl_component.Jurisdiction, anl_component.Land_right_id, anl_component.Land_right_subtype, anl_component.Land_sale_number, anl_component.Language, anl_component.Lithology_log_id, anl_component.Lith_log_source, anl_component.Notification_id, anl_component.Obligation_id, anl_component.Obligation_seq_no, anl_component.Paleo_summary_id, anl_component.Pden_id, anl_component.Pden_source, anl_component.Pden_subtype, anl_component.Physical_item_id, anl_component.Pool_id, anl_component.Ppdm_guid, anl_component.Ppdm_system_id, anl_component.Ppdm_table_name, anl_component.Product_type, anl_component.Prod_string_source, anl_component.Project_id, anl_component.Pr_str_form_obs_no, anl_component.Rate_schedule_id, anl_component.Recovery_obs_no, anl_component.Referenced_guid, anl_component.Referenced_system_id, anl_component.Referenced_table_name, anl_component.Remark, anl_component.Reported_ind, anl_component.Resent_id, anl_component.Reserve_class_id, anl_component.Run_num, anl_component.Sample_date, anl_component.Sample_id, anl_component.Sample_source, anl_component.Seis_set_id, anl_component.Seis_set_subtype, anl_component.Sf_subtype, anl_component.Source, anl_component.Spatial_description_id, anl_component.Spatial_obs_no, anl_component.Step_seq_no, anl_component.Store_id, anl_component.Store_structure_obs_no, anl_component.Strat_name_set_id, anl_component.Strat_unit_id, anl_component.String_id, anl_component.Support_facility_id, anl_component.Test_num, anl_component.Test_source, anl_component.Test_type, anl_component.Thesaurus_id, anl_component.Thesaurus_word_id, anl_component.Uwi, anl_component.Well_activity_id, anl_component.Well_activity_set_id, anl_component.Well_activity_source, anl_component.Well_activity_type_id, anl_component.Well_core_desc_id, anl_component.Well_test_id, anl_component.Well_test_recovery_id, anl_component.Well_test_source, anl_component.Well_test_type, anl_component.Work_order_id, anl_component.Row_changed_by, anl_component.Row_changed_date, anl_component.Row_created_by, anl_component.Row_effective_date, anl_component.Row_expiry_date, anl_component.Row_quality, anl_component.Analysis_id)
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

func PatchAnlComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_component set "
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
	query += " where analysis_id = :id"

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

func DeleteAnlComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_component dto.Anl_component
	anl_component.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_component where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_component.Analysis_id)
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


