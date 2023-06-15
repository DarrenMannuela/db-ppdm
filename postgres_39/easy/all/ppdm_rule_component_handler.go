package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmRuleComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_rule_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_rule_component

	for rows.Next() {
		var ppdm_rule_component dto.Ppdm_rule_component
		if err := rows.Scan(&ppdm_rule_component.Rule_id, &ppdm_rule_component.Component_obs_no, &ppdm_rule_component.Active_ind, &ppdm_rule_component.Activity_obs_no, &ppdm_rule_component.Analysis_id, &ppdm_rule_component.Application_id, &ppdm_rule_component.Area_id, &ppdm_rule_component.Area_type, &ppdm_rule_component.Authority_id, &ppdm_rule_component.Ba_licensee_ba_id, &ppdm_rule_component.Ba_license_id, &ppdm_rule_component.Ba_organization_id, &ppdm_rule_component.Ba_organization_seq_no, &ppdm_rule_component.Business_associate_id, &ppdm_rule_component.Catalogue_additive_id, &ppdm_rule_component.Catalogue_equip_id, &ppdm_rule_component.Classification_system_id, &ppdm_rule_component.Class_level_id, &ppdm_rule_component.Consent_id, &ppdm_rule_component.Consult_id, &ppdm_rule_component.Contest_id, &ppdm_rule_component.Contract_id, &ppdm_rule_component.Ecozone_id, &ppdm_rule_component.Effective_date, &ppdm_rule_component.Employee_ba_id, &ppdm_rule_component.Employee_obs_no, &ppdm_rule_component.Employer_ba_id, &ppdm_rule_component.Entitlement_id, &ppdm_rule_component.Equipment_id, &ppdm_rule_component.Expiry_date, &ppdm_rule_component.Facility_id, &ppdm_rule_component.Facility_license_id, &ppdm_rule_component.Facility_type, &ppdm_rule_component.Field_id, &ppdm_rule_component.Field_station_id, &ppdm_rule_component.Finance_id, &ppdm_rule_component.Fossil_id, &ppdm_rule_component.Incident_id, &ppdm_rule_component.Incident_set_id, &ppdm_rule_component.Incident_type, &ppdm_rule_component.Information_item_id, &ppdm_rule_component.Info_item_subtype, &ppdm_rule_component.Instrument_id, &ppdm_rule_component.Interest_set_id, &ppdm_rule_component.Interest_set_seq_no, &ppdm_rule_component.Jurisdiction, &ppdm_rule_component.Land_right_id, &ppdm_rule_component.Land_right_subtype, &ppdm_rule_component.Land_sale_number, &ppdm_rule_component.Language, &ppdm_rule_component.Lithology_log_id, &ppdm_rule_component.Lith_log_source, &ppdm_rule_component.Notification_id, &ppdm_rule_component.Obligation_id, &ppdm_rule_component.Obligation_seq_no, &ppdm_rule_component.Paleo_summary_id, &ppdm_rule_component.Pden_id, &ppdm_rule_component.Pden_source, &ppdm_rule_component.Pden_subtype, &ppdm_rule_component.Physical_item_id, &ppdm_rule_component.Pool_id, &ppdm_rule_component.Ppdm_column_name, &ppdm_rule_component.Ppdm_guid, &ppdm_rule_component.Ppdm_system_id, &ppdm_rule_component.Ppdm_table_name, &ppdm_rule_component.Product_type, &ppdm_rule_component.Project_id, &ppdm_rule_component.Pr_str_source, &ppdm_rule_component.Pr_str_uwi, &ppdm_rule_component.Rate_schedule_id, &ppdm_rule_component.Referenced_guid, &ppdm_rule_component.Referenced_system_id, &ppdm_rule_component.Referenced_table_name, &ppdm_rule_component.Remark, &ppdm_rule_component.Resent_id, &ppdm_rule_component.Reserve_class_id, &ppdm_rule_component.Restriction_id, &ppdm_rule_component.Restriction_version, &ppdm_rule_component.Rule_component_type, &ppdm_rule_component.Sample_anal_source, &ppdm_rule_component.Schema_entity_id, &ppdm_rule_component.Seis_license_id, &ppdm_rule_component.Seis_set_id, &ppdm_rule_component.Seis_set_subtype, &ppdm_rule_component.Sf_subtype, &ppdm_rule_component.Source, &ppdm_rule_component.Spatial_description_id, &ppdm_rule_component.Spatial_obs_no, &ppdm_rule_component.Store_id, &ppdm_rule_component.Store_structure_obs_no, &ppdm_rule_component.Strat_name_set_id, &ppdm_rule_component.Strat_unit_id, &ppdm_rule_component.String_id, &ppdm_rule_component.Support_facility_id, &ppdm_rule_component.Sw_application_id, &ppdm_rule_component.System_id, &ppdm_rule_component.Thesaurus_id, &ppdm_rule_component.Thesaurus_word_id, &ppdm_rule_component.Uwi, &ppdm_rule_component.Well_activity_set_id, &ppdm_rule_component.Well_activity_source, &ppdm_rule_component.Well_activity_type_id, &ppdm_rule_component.Well_activity_uwi, &ppdm_rule_component.Well_license_id, &ppdm_rule_component.Well_license_source, &ppdm_rule_component.Well_set_id, &ppdm_rule_component.Work_order_id, &ppdm_rule_component.Row_changed_by, &ppdm_rule_component.Row_changed_date, &ppdm_rule_component.Row_created_by, &ppdm_rule_component.Row_created_date, &ppdm_rule_component.Row_effective_date, &ppdm_rule_component.Row_expiry_date, &ppdm_rule_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_rule_component to result
		result = append(result, ppdm_rule_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmRuleComponent(c *fiber.Ctx) error {
	var ppdm_rule_component dto.Ppdm_rule_component

	setDefaults(&ppdm_rule_component)

	if err := json.Unmarshal(c.Body(), &ppdm_rule_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_rule_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103, :104, :105, :106, :107, :108, :109, :110, :111, :112, :113, :114)")
	if err != nil {
		return err
	}
	ppdm_rule_component.Row_created_date = formatDate(ppdm_rule_component.Row_created_date)
	ppdm_rule_component.Row_changed_date = formatDate(ppdm_rule_component.Row_changed_date)
	ppdm_rule_component.Effective_date = formatDateString(ppdm_rule_component.Effective_date)
	ppdm_rule_component.Expiry_date = formatDateString(ppdm_rule_component.Expiry_date)
	ppdm_rule_component.Row_effective_date = formatDateString(ppdm_rule_component.Row_effective_date)
	ppdm_rule_component.Row_expiry_date = formatDateString(ppdm_rule_component.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_rule_component.Rule_id, ppdm_rule_component.Component_obs_no, ppdm_rule_component.Active_ind, ppdm_rule_component.Activity_obs_no, ppdm_rule_component.Analysis_id, ppdm_rule_component.Application_id, ppdm_rule_component.Area_id, ppdm_rule_component.Area_type, ppdm_rule_component.Authority_id, ppdm_rule_component.Ba_licensee_ba_id, ppdm_rule_component.Ba_license_id, ppdm_rule_component.Ba_organization_id, ppdm_rule_component.Ba_organization_seq_no, ppdm_rule_component.Business_associate_id, ppdm_rule_component.Catalogue_additive_id, ppdm_rule_component.Catalogue_equip_id, ppdm_rule_component.Classification_system_id, ppdm_rule_component.Class_level_id, ppdm_rule_component.Consent_id, ppdm_rule_component.Consult_id, ppdm_rule_component.Contest_id, ppdm_rule_component.Contract_id, ppdm_rule_component.Ecozone_id, ppdm_rule_component.Effective_date, ppdm_rule_component.Employee_ba_id, ppdm_rule_component.Employee_obs_no, ppdm_rule_component.Employer_ba_id, ppdm_rule_component.Entitlement_id, ppdm_rule_component.Equipment_id, ppdm_rule_component.Expiry_date, ppdm_rule_component.Facility_id, ppdm_rule_component.Facility_license_id, ppdm_rule_component.Facility_type, ppdm_rule_component.Field_id, ppdm_rule_component.Field_station_id, ppdm_rule_component.Finance_id, ppdm_rule_component.Fossil_id, ppdm_rule_component.Incident_id, ppdm_rule_component.Incident_set_id, ppdm_rule_component.Incident_type, ppdm_rule_component.Information_item_id, ppdm_rule_component.Info_item_subtype, ppdm_rule_component.Instrument_id, ppdm_rule_component.Interest_set_id, ppdm_rule_component.Interest_set_seq_no, ppdm_rule_component.Jurisdiction, ppdm_rule_component.Land_right_id, ppdm_rule_component.Land_right_subtype, ppdm_rule_component.Land_sale_number, ppdm_rule_component.Language, ppdm_rule_component.Lithology_log_id, ppdm_rule_component.Lith_log_source, ppdm_rule_component.Notification_id, ppdm_rule_component.Obligation_id, ppdm_rule_component.Obligation_seq_no, ppdm_rule_component.Paleo_summary_id, ppdm_rule_component.Pden_id, ppdm_rule_component.Pden_source, ppdm_rule_component.Pden_subtype, ppdm_rule_component.Physical_item_id, ppdm_rule_component.Pool_id, ppdm_rule_component.Ppdm_column_name, ppdm_rule_component.Ppdm_guid, ppdm_rule_component.Ppdm_system_id, ppdm_rule_component.Ppdm_table_name, ppdm_rule_component.Product_type, ppdm_rule_component.Project_id, ppdm_rule_component.Pr_str_source, ppdm_rule_component.Pr_str_uwi, ppdm_rule_component.Rate_schedule_id, ppdm_rule_component.Referenced_guid, ppdm_rule_component.Referenced_system_id, ppdm_rule_component.Referenced_table_name, ppdm_rule_component.Remark, ppdm_rule_component.Resent_id, ppdm_rule_component.Reserve_class_id, ppdm_rule_component.Restriction_id, ppdm_rule_component.Restriction_version, ppdm_rule_component.Rule_component_type, ppdm_rule_component.Sample_anal_source, ppdm_rule_component.Schema_entity_id, ppdm_rule_component.Seis_license_id, ppdm_rule_component.Seis_set_id, ppdm_rule_component.Seis_set_subtype, ppdm_rule_component.Sf_subtype, ppdm_rule_component.Source, ppdm_rule_component.Spatial_description_id, ppdm_rule_component.Spatial_obs_no, ppdm_rule_component.Store_id, ppdm_rule_component.Store_structure_obs_no, ppdm_rule_component.Strat_name_set_id, ppdm_rule_component.Strat_unit_id, ppdm_rule_component.String_id, ppdm_rule_component.Support_facility_id, ppdm_rule_component.Sw_application_id, ppdm_rule_component.System_id, ppdm_rule_component.Thesaurus_id, ppdm_rule_component.Thesaurus_word_id, ppdm_rule_component.Uwi, ppdm_rule_component.Well_activity_set_id, ppdm_rule_component.Well_activity_source, ppdm_rule_component.Well_activity_type_id, ppdm_rule_component.Well_activity_uwi, ppdm_rule_component.Well_license_id, ppdm_rule_component.Well_license_source, ppdm_rule_component.Well_set_id, ppdm_rule_component.Work_order_id, ppdm_rule_component.Row_changed_by, ppdm_rule_component.Row_changed_date, ppdm_rule_component.Row_created_by, ppdm_rule_component.Row_created_date, ppdm_rule_component.Row_effective_date, ppdm_rule_component.Row_expiry_date, ppdm_rule_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmRuleComponent(c *fiber.Ctx) error {
	var ppdm_rule_component dto.Ppdm_rule_component

	setDefaults(&ppdm_rule_component)

	if err := json.Unmarshal(c.Body(), &ppdm_rule_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_rule_component.Rule_id = id

    if ppdm_rule_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_rule_component set component_obs_no = :1, active_ind = :2, activity_obs_no = :3, analysis_id = :4, application_id = :5, area_id = :6, area_type = :7, authority_id = :8, ba_licensee_ba_id = :9, ba_license_id = :10, ba_organization_id = :11, ba_organization_seq_no = :12, business_associate_id = :13, catalogue_additive_id = :14, catalogue_equip_id = :15, classification_system_id = :16, class_level_id = :17, consent_id = :18, consult_id = :19, contest_id = :20, contract_id = :21, ecozone_id = :22, effective_date = :23, employee_ba_id = :24, employee_obs_no = :25, employer_ba_id = :26, entitlement_id = :27, equipment_id = :28, expiry_date = :29, facility_id = :30, facility_license_id = :31, facility_type = :32, field_id = :33, field_station_id = :34, finance_id = :35, fossil_id = :36, incident_id = :37, incident_set_id = :38, incident_type = :39, information_item_id = :40, info_item_subtype = :41, instrument_id = :42, interest_set_id = :43, interest_set_seq_no = :44, jurisdiction = :45, land_right_id = :46, land_right_subtype = :47, land_sale_number = :48, language = :49, lithology_log_id = :50, lith_log_source = :51, notification_id = :52, obligation_id = :53, obligation_seq_no = :54, paleo_summary_id = :55, pden_id = :56, pden_source = :57, pden_subtype = :58, physical_item_id = :59, pool_id = :60, ppdm_column_name = :61, ppdm_guid = :62, ppdm_system_id = :63, ppdm_table_name = :64, product_type = :65, project_id = :66, pr_str_source = :67, pr_str_uwi = :68, rate_schedule_id = :69, referenced_guid = :70, referenced_system_id = :71, referenced_table_name = :72, remark = :73, resent_id = :74, reserve_class_id = :75, restriction_id = :76, restriction_version = :77, rule_component_type = :78, sample_anal_source = :79, schema_entity_id = :80, seis_license_id = :81, seis_set_id = :82, seis_set_subtype = :83, sf_subtype = :84, source = :85, spatial_description_id = :86, spatial_obs_no = :87, store_id = :88, store_structure_obs_no = :89, strat_name_set_id = :90, strat_unit_id = :91, string_id = :92, support_facility_id = :93, sw_application_id = :94, system_id = :95, thesaurus_id = :96, thesaurus_word_id = :97, uwi = :98, well_activity_set_id = :99, well_activity_source = :100, well_activity_type_id = :101, well_activity_uwi = :102, well_license_id = :103, well_license_source = :104, well_set_id = :105, work_order_id = :106, row_changed_by = :107, row_changed_date = :108, row_created_by = :109, row_effective_date = :110, row_expiry_date = :111, row_quality = :112 where rule_id = :114")
	if err != nil {
		return err
	}

	ppdm_rule_component.Row_changed_date = formatDate(ppdm_rule_component.Row_changed_date)
	ppdm_rule_component.Effective_date = formatDateString(ppdm_rule_component.Effective_date)
	ppdm_rule_component.Expiry_date = formatDateString(ppdm_rule_component.Expiry_date)
	ppdm_rule_component.Row_effective_date = formatDateString(ppdm_rule_component.Row_effective_date)
	ppdm_rule_component.Row_expiry_date = formatDateString(ppdm_rule_component.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_rule_component.Component_obs_no, ppdm_rule_component.Active_ind, ppdm_rule_component.Activity_obs_no, ppdm_rule_component.Analysis_id, ppdm_rule_component.Application_id, ppdm_rule_component.Area_id, ppdm_rule_component.Area_type, ppdm_rule_component.Authority_id, ppdm_rule_component.Ba_licensee_ba_id, ppdm_rule_component.Ba_license_id, ppdm_rule_component.Ba_organization_id, ppdm_rule_component.Ba_organization_seq_no, ppdm_rule_component.Business_associate_id, ppdm_rule_component.Catalogue_additive_id, ppdm_rule_component.Catalogue_equip_id, ppdm_rule_component.Classification_system_id, ppdm_rule_component.Class_level_id, ppdm_rule_component.Consent_id, ppdm_rule_component.Consult_id, ppdm_rule_component.Contest_id, ppdm_rule_component.Contract_id, ppdm_rule_component.Ecozone_id, ppdm_rule_component.Effective_date, ppdm_rule_component.Employee_ba_id, ppdm_rule_component.Employee_obs_no, ppdm_rule_component.Employer_ba_id, ppdm_rule_component.Entitlement_id, ppdm_rule_component.Equipment_id, ppdm_rule_component.Expiry_date, ppdm_rule_component.Facility_id, ppdm_rule_component.Facility_license_id, ppdm_rule_component.Facility_type, ppdm_rule_component.Field_id, ppdm_rule_component.Field_station_id, ppdm_rule_component.Finance_id, ppdm_rule_component.Fossil_id, ppdm_rule_component.Incident_id, ppdm_rule_component.Incident_set_id, ppdm_rule_component.Incident_type, ppdm_rule_component.Information_item_id, ppdm_rule_component.Info_item_subtype, ppdm_rule_component.Instrument_id, ppdm_rule_component.Interest_set_id, ppdm_rule_component.Interest_set_seq_no, ppdm_rule_component.Jurisdiction, ppdm_rule_component.Land_right_id, ppdm_rule_component.Land_right_subtype, ppdm_rule_component.Land_sale_number, ppdm_rule_component.Language, ppdm_rule_component.Lithology_log_id, ppdm_rule_component.Lith_log_source, ppdm_rule_component.Notification_id, ppdm_rule_component.Obligation_id, ppdm_rule_component.Obligation_seq_no, ppdm_rule_component.Paleo_summary_id, ppdm_rule_component.Pden_id, ppdm_rule_component.Pden_source, ppdm_rule_component.Pden_subtype, ppdm_rule_component.Physical_item_id, ppdm_rule_component.Pool_id, ppdm_rule_component.Ppdm_column_name, ppdm_rule_component.Ppdm_guid, ppdm_rule_component.Ppdm_system_id, ppdm_rule_component.Ppdm_table_name, ppdm_rule_component.Product_type, ppdm_rule_component.Project_id, ppdm_rule_component.Pr_str_source, ppdm_rule_component.Pr_str_uwi, ppdm_rule_component.Rate_schedule_id, ppdm_rule_component.Referenced_guid, ppdm_rule_component.Referenced_system_id, ppdm_rule_component.Referenced_table_name, ppdm_rule_component.Remark, ppdm_rule_component.Resent_id, ppdm_rule_component.Reserve_class_id, ppdm_rule_component.Restriction_id, ppdm_rule_component.Restriction_version, ppdm_rule_component.Rule_component_type, ppdm_rule_component.Sample_anal_source, ppdm_rule_component.Schema_entity_id, ppdm_rule_component.Seis_license_id, ppdm_rule_component.Seis_set_id, ppdm_rule_component.Seis_set_subtype, ppdm_rule_component.Sf_subtype, ppdm_rule_component.Source, ppdm_rule_component.Spatial_description_id, ppdm_rule_component.Spatial_obs_no, ppdm_rule_component.Store_id, ppdm_rule_component.Store_structure_obs_no, ppdm_rule_component.Strat_name_set_id, ppdm_rule_component.Strat_unit_id, ppdm_rule_component.String_id, ppdm_rule_component.Support_facility_id, ppdm_rule_component.Sw_application_id, ppdm_rule_component.System_id, ppdm_rule_component.Thesaurus_id, ppdm_rule_component.Thesaurus_word_id, ppdm_rule_component.Uwi, ppdm_rule_component.Well_activity_set_id, ppdm_rule_component.Well_activity_source, ppdm_rule_component.Well_activity_type_id, ppdm_rule_component.Well_activity_uwi, ppdm_rule_component.Well_license_id, ppdm_rule_component.Well_license_source, ppdm_rule_component.Well_set_id, ppdm_rule_component.Work_order_id, ppdm_rule_component.Row_changed_by, ppdm_rule_component.Row_changed_date, ppdm_rule_component.Row_created_by, ppdm_rule_component.Row_effective_date, ppdm_rule_component.Row_expiry_date, ppdm_rule_component.Row_quality, ppdm_rule_component.Rule_id)
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

func PatchPpdmRuleComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_rule_component set "
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
	query += " where rule_id = :id"

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

func DeletePpdmRuleComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_rule_component dto.Ppdm_rule_component
	ppdm_rule_component.Rule_id = id

	stmt, err := db.Prepare("delete from ppdm_rule_component where rule_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_rule_component.Rule_id)
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


