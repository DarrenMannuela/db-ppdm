package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetBaOrganizationComp(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ba_organization_comp")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ba_organization_comp

	for rows.Next() {
		var ba_organization_comp dto.Ba_organization_comp
		if err := rows.Scan(&ba_organization_comp.Business_associate_id, &ba_organization_comp.Organization_id, &ba_organization_comp.Organization_seq_no, &ba_organization_comp.Component_obs_no, &ba_organization_comp.Active_ind, &ba_organization_comp.Activity_obs_no, &ba_organization_comp.Analysis_id, &ba_organization_comp.Application_id, &ba_organization_comp.Area_id, &ba_organization_comp.Area_type, &ba_organization_comp.Authority_id, &ba_organization_comp.Ba_organization_component_type, &ba_organization_comp.Catalogue_additive_id, &ba_organization_comp.Catalogue_equip_id, &ba_organization_comp.Classification_system_id, &ba_organization_comp.Class_level_id, &ba_organization_comp.Consent_id, &ba_organization_comp.Consult_id, &ba_organization_comp.Contest_id, &ba_organization_comp.Contract_id, &ba_organization_comp.Description, &ba_organization_comp.Ecozone_id, &ba_organization_comp.Effective_date, &ba_organization_comp.Employee_ba_id, &ba_organization_comp.Employee_obs_no, &ba_organization_comp.Employer_ba_id, &ba_organization_comp.Entitlement_id, &ba_organization_comp.Equipment_id, &ba_organization_comp.Expiry_date, &ba_organization_comp.Facility_id, &ba_organization_comp.Facility_type, &ba_organization_comp.Field_id, &ba_organization_comp.Field_station_id, &ba_organization_comp.Finance_id, &ba_organization_comp.Fossil_id, &ba_organization_comp.Incident_id, &ba_organization_comp.Incident_set_id, &ba_organization_comp.Incident_type, &ba_organization_comp.Information_item_id, &ba_organization_comp.Info_item_subtype, &ba_organization_comp.Instrument_id, &ba_organization_comp.Interest_set_id, &ba_organization_comp.Interest_set_seq_no, &ba_organization_comp.Jurisdiction, &ba_organization_comp.Land_right_id, &ba_organization_comp.Land_right_subtype, &ba_organization_comp.Land_sale_number, &ba_organization_comp.Lithology_log_id, &ba_organization_comp.Lith_log_source, &ba_organization_comp.Notification_id, &ba_organization_comp.Obligation_id, &ba_organization_comp.Obligation_seq_no, &ba_organization_comp.Organization_id2, &ba_organization_comp.Organization_seq_no2, &ba_organization_comp.Paleo_summary_id, &ba_organization_comp.Pden_id, &ba_organization_comp.Pden_source, &ba_organization_comp.Pden_subtype, &ba_organization_comp.Physical_item_id, &ba_organization_comp.Pool_id, &ba_organization_comp.Ppdm_guid, &ba_organization_comp.Ppdm_system_id, &ba_organization_comp.Ppdm_table_name, &ba_organization_comp.Product_type, &ba_organization_comp.Prod_string_id, &ba_organization_comp.Prod_string_source, &ba_organization_comp.Project_id, &ba_organization_comp.Pr_str_form_obs_no, &ba_organization_comp.Rate_schedule_id, &ba_organization_comp.Referenced_guid, &ba_organization_comp.Referenced_system_id, &ba_organization_comp.Referenced_table_name, &ba_organization_comp.Remark, &ba_organization_comp.Resent_id, &ba_organization_comp.Reserve_class_id, &ba_organization_comp.Sample_anal_source, &ba_organization_comp.Seis_set_id, &ba_organization_comp.Seis_set_subtype, &ba_organization_comp.Sf_subtype, &ba_organization_comp.Source, &ba_organization_comp.Source_document_id, &ba_organization_comp.Spatial_description_id, &ba_organization_comp.Spatial_obs_no, &ba_organization_comp.Store_id, &ba_organization_comp.Store_structure_obs_no, &ba_organization_comp.Strat_name_set_id, &ba_organization_comp.Strat_unit_id, &ba_organization_comp.Support_facility_id, &ba_organization_comp.Sw_application_id, &ba_organization_comp.Thesaurus_id, &ba_organization_comp.Thesaurus_word_id, &ba_organization_comp.Uwi, &ba_organization_comp.Well_activity_set_id, &ba_organization_comp.Well_activity_source, &ba_organization_comp.Well_activity_type_id, &ba_organization_comp.Well_set_id, &ba_organization_comp.Work_order_id, &ba_organization_comp.Row_changed_by, &ba_organization_comp.Row_changed_date, &ba_organization_comp.Row_created_by, &ba_organization_comp.Row_created_date, &ba_organization_comp.Row_effective_date, &ba_organization_comp.Row_expiry_date, &ba_organization_comp.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ba_organization_comp to result
		result = append(result, ba_organization_comp)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetBaOrganizationComp(c *fiber.Ctx) error {
	var ba_organization_comp dto.Ba_organization_comp

	setDefaults(&ba_organization_comp)

	if err := json.Unmarshal(c.Body(), &ba_organization_comp); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ba_organization_comp values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103, :104)")
	if err != nil {
		return err
	}
	ba_organization_comp.Row_created_date = formatDate(ba_organization_comp.Row_created_date)
	ba_organization_comp.Row_changed_date = formatDate(ba_organization_comp.Row_changed_date)
	ba_organization_comp.Effective_date = formatDateString(ba_organization_comp.Effective_date)
	ba_organization_comp.Expiry_date = formatDateString(ba_organization_comp.Expiry_date)
	ba_organization_comp.Row_effective_date = formatDateString(ba_organization_comp.Row_effective_date)
	ba_organization_comp.Row_expiry_date = formatDateString(ba_organization_comp.Row_expiry_date)






	rows, err := stmt.Exec(ba_organization_comp.Business_associate_id, ba_organization_comp.Organization_id, ba_organization_comp.Organization_seq_no, ba_organization_comp.Component_obs_no, ba_organization_comp.Active_ind, ba_organization_comp.Activity_obs_no, ba_organization_comp.Analysis_id, ba_organization_comp.Application_id, ba_organization_comp.Area_id, ba_organization_comp.Area_type, ba_organization_comp.Authority_id, ba_organization_comp.Ba_organization_component_type, ba_organization_comp.Catalogue_additive_id, ba_organization_comp.Catalogue_equip_id, ba_organization_comp.Classification_system_id, ba_organization_comp.Class_level_id, ba_organization_comp.Consent_id, ba_organization_comp.Consult_id, ba_organization_comp.Contest_id, ba_organization_comp.Contract_id, ba_organization_comp.Description, ba_organization_comp.Ecozone_id, ba_organization_comp.Effective_date, ba_organization_comp.Employee_ba_id, ba_organization_comp.Employee_obs_no, ba_organization_comp.Employer_ba_id, ba_organization_comp.Entitlement_id, ba_organization_comp.Equipment_id, ba_organization_comp.Expiry_date, ba_organization_comp.Facility_id, ba_organization_comp.Facility_type, ba_organization_comp.Field_id, ba_organization_comp.Field_station_id, ba_organization_comp.Finance_id, ba_organization_comp.Fossil_id, ba_organization_comp.Incident_id, ba_organization_comp.Incident_set_id, ba_organization_comp.Incident_type, ba_organization_comp.Information_item_id, ba_organization_comp.Info_item_subtype, ba_organization_comp.Instrument_id, ba_organization_comp.Interest_set_id, ba_organization_comp.Interest_set_seq_no, ba_organization_comp.Jurisdiction, ba_organization_comp.Land_right_id, ba_organization_comp.Land_right_subtype, ba_organization_comp.Land_sale_number, ba_organization_comp.Lithology_log_id, ba_organization_comp.Lith_log_source, ba_organization_comp.Notification_id, ba_organization_comp.Obligation_id, ba_organization_comp.Obligation_seq_no, ba_organization_comp.Organization_id2, ba_organization_comp.Organization_seq_no2, ba_organization_comp.Paleo_summary_id, ba_organization_comp.Pden_id, ba_organization_comp.Pden_source, ba_organization_comp.Pden_subtype, ba_organization_comp.Physical_item_id, ba_organization_comp.Pool_id, ba_organization_comp.Ppdm_guid, ba_organization_comp.Ppdm_system_id, ba_organization_comp.Ppdm_table_name, ba_organization_comp.Product_type, ba_organization_comp.Prod_string_id, ba_organization_comp.Prod_string_source, ba_organization_comp.Project_id, ba_organization_comp.Pr_str_form_obs_no, ba_organization_comp.Rate_schedule_id, ba_organization_comp.Referenced_guid, ba_organization_comp.Referenced_system_id, ba_organization_comp.Referenced_table_name, ba_organization_comp.Remark, ba_organization_comp.Resent_id, ba_organization_comp.Reserve_class_id, ba_organization_comp.Sample_anal_source, ba_organization_comp.Seis_set_id, ba_organization_comp.Seis_set_subtype, ba_organization_comp.Sf_subtype, ba_organization_comp.Source, ba_organization_comp.Source_document_id, ba_organization_comp.Spatial_description_id, ba_organization_comp.Spatial_obs_no, ba_organization_comp.Store_id, ba_organization_comp.Store_structure_obs_no, ba_organization_comp.Strat_name_set_id, ba_organization_comp.Strat_unit_id, ba_organization_comp.Support_facility_id, ba_organization_comp.Sw_application_id, ba_organization_comp.Thesaurus_id, ba_organization_comp.Thesaurus_word_id, ba_organization_comp.Uwi, ba_organization_comp.Well_activity_set_id, ba_organization_comp.Well_activity_source, ba_organization_comp.Well_activity_type_id, ba_organization_comp.Well_set_id, ba_organization_comp.Work_order_id, ba_organization_comp.Row_changed_by, ba_organization_comp.Row_changed_date, ba_organization_comp.Row_created_by, ba_organization_comp.Row_created_date, ba_organization_comp.Row_effective_date, ba_organization_comp.Row_expiry_date, ba_organization_comp.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateBaOrganizationComp(c *fiber.Ctx) error {
	var ba_organization_comp dto.Ba_organization_comp

	setDefaults(&ba_organization_comp)

	if err := json.Unmarshal(c.Body(), &ba_organization_comp); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ba_organization_comp.Business_associate_id = id

    if ba_organization_comp.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ba_organization_comp set organization_id = :1, organization_seq_no = :2, component_obs_no = :3, active_ind = :4, activity_obs_no = :5, analysis_id = :6, application_id = :7, area_id = :8, area_type = :9, authority_id = :10, ba_organization_component_type = :11, catalogue_additive_id = :12, catalogue_equip_id = :13, classification_system_id = :14, class_level_id = :15, consent_id = :16, consult_id = :17, contest_id = :18, contract_id = :19, description = :20, ecozone_id = :21, effective_date = :22, employee_ba_id = :23, employee_obs_no = :24, employer_ba_id = :25, entitlement_id = :26, equipment_id = :27, expiry_date = :28, facility_id = :29, facility_type = :30, field_id = :31, field_station_id = :32, finance_id = :33, fossil_id = :34, incident_id = :35, incident_set_id = :36, incident_type = :37, information_item_id = :38, info_item_subtype = :39, instrument_id = :40, interest_set_id = :41, interest_set_seq_no = :42, jurisdiction = :43, land_right_id = :44, land_right_subtype = :45, land_sale_number = :46, lithology_log_id = :47, lith_log_source = :48, notification_id = :49, obligation_id = :50, obligation_seq_no = :51, organization_id2 = :52, organization_seq_no2 = :53, paleo_summary_id = :54, pden_id = :55, pden_source = :56, pden_subtype = :57, physical_item_id = :58, pool_id = :59, ppdm_guid = :60, ppdm_system_id = :61, ppdm_table_name = :62, product_type = :63, prod_string_id = :64, prod_string_source = :65, project_id = :66, pr_str_form_obs_no = :67, rate_schedule_id = :68, referenced_guid = :69, referenced_system_id = :70, referenced_table_name = :71, remark = :72, resent_id = :73, reserve_class_id = :74, sample_anal_source = :75, seis_set_id = :76, seis_set_subtype = :77, sf_subtype = :78, source = :79, source_document_id = :80, spatial_description_id = :81, spatial_obs_no = :82, store_id = :83, store_structure_obs_no = :84, strat_name_set_id = :85, strat_unit_id = :86, support_facility_id = :87, sw_application_id = :88, thesaurus_id = :89, thesaurus_word_id = :90, uwi = :91, well_activity_set_id = :92, well_activity_source = :93, well_activity_type_id = :94, well_set_id = :95, work_order_id = :96, row_changed_by = :97, row_changed_date = :98, row_created_by = :99, row_effective_date = :100, row_expiry_date = :101, row_quality = :102 where business_associate_id = :104")
	if err != nil {
		return err
	}

	ba_organization_comp.Row_changed_date = formatDate(ba_organization_comp.Row_changed_date)
	ba_organization_comp.Effective_date = formatDateString(ba_organization_comp.Effective_date)
	ba_organization_comp.Expiry_date = formatDateString(ba_organization_comp.Expiry_date)
	ba_organization_comp.Row_effective_date = formatDateString(ba_organization_comp.Row_effective_date)
	ba_organization_comp.Row_expiry_date = formatDateString(ba_organization_comp.Row_expiry_date)






	rows, err := stmt.Exec(ba_organization_comp.Organization_id, ba_organization_comp.Organization_seq_no, ba_organization_comp.Component_obs_no, ba_organization_comp.Active_ind, ba_organization_comp.Activity_obs_no, ba_organization_comp.Analysis_id, ba_organization_comp.Application_id, ba_organization_comp.Area_id, ba_organization_comp.Area_type, ba_organization_comp.Authority_id, ba_organization_comp.Ba_organization_component_type, ba_organization_comp.Catalogue_additive_id, ba_organization_comp.Catalogue_equip_id, ba_organization_comp.Classification_system_id, ba_organization_comp.Class_level_id, ba_organization_comp.Consent_id, ba_organization_comp.Consult_id, ba_organization_comp.Contest_id, ba_organization_comp.Contract_id, ba_organization_comp.Description, ba_organization_comp.Ecozone_id, ba_organization_comp.Effective_date, ba_organization_comp.Employee_ba_id, ba_organization_comp.Employee_obs_no, ba_organization_comp.Employer_ba_id, ba_organization_comp.Entitlement_id, ba_organization_comp.Equipment_id, ba_organization_comp.Expiry_date, ba_organization_comp.Facility_id, ba_organization_comp.Facility_type, ba_organization_comp.Field_id, ba_organization_comp.Field_station_id, ba_organization_comp.Finance_id, ba_organization_comp.Fossil_id, ba_organization_comp.Incident_id, ba_organization_comp.Incident_set_id, ba_organization_comp.Incident_type, ba_organization_comp.Information_item_id, ba_organization_comp.Info_item_subtype, ba_organization_comp.Instrument_id, ba_organization_comp.Interest_set_id, ba_organization_comp.Interest_set_seq_no, ba_organization_comp.Jurisdiction, ba_organization_comp.Land_right_id, ba_organization_comp.Land_right_subtype, ba_organization_comp.Land_sale_number, ba_organization_comp.Lithology_log_id, ba_organization_comp.Lith_log_source, ba_organization_comp.Notification_id, ba_organization_comp.Obligation_id, ba_organization_comp.Obligation_seq_no, ba_organization_comp.Organization_id2, ba_organization_comp.Organization_seq_no2, ba_organization_comp.Paleo_summary_id, ba_organization_comp.Pden_id, ba_organization_comp.Pden_source, ba_organization_comp.Pden_subtype, ba_organization_comp.Physical_item_id, ba_organization_comp.Pool_id, ba_organization_comp.Ppdm_guid, ba_organization_comp.Ppdm_system_id, ba_organization_comp.Ppdm_table_name, ba_organization_comp.Product_type, ba_organization_comp.Prod_string_id, ba_organization_comp.Prod_string_source, ba_organization_comp.Project_id, ba_organization_comp.Pr_str_form_obs_no, ba_organization_comp.Rate_schedule_id, ba_organization_comp.Referenced_guid, ba_organization_comp.Referenced_system_id, ba_organization_comp.Referenced_table_name, ba_organization_comp.Remark, ba_organization_comp.Resent_id, ba_organization_comp.Reserve_class_id, ba_organization_comp.Sample_anal_source, ba_organization_comp.Seis_set_id, ba_organization_comp.Seis_set_subtype, ba_organization_comp.Sf_subtype, ba_organization_comp.Source, ba_organization_comp.Source_document_id, ba_organization_comp.Spatial_description_id, ba_organization_comp.Spatial_obs_no, ba_organization_comp.Store_id, ba_organization_comp.Store_structure_obs_no, ba_organization_comp.Strat_name_set_id, ba_organization_comp.Strat_unit_id, ba_organization_comp.Support_facility_id, ba_organization_comp.Sw_application_id, ba_organization_comp.Thesaurus_id, ba_organization_comp.Thesaurus_word_id, ba_organization_comp.Uwi, ba_organization_comp.Well_activity_set_id, ba_organization_comp.Well_activity_source, ba_organization_comp.Well_activity_type_id, ba_organization_comp.Well_set_id, ba_organization_comp.Work_order_id, ba_organization_comp.Row_changed_by, ba_organization_comp.Row_changed_date, ba_organization_comp.Row_created_by, ba_organization_comp.Row_effective_date, ba_organization_comp.Row_expiry_date, ba_organization_comp.Row_quality, ba_organization_comp.Business_associate_id)
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

func PatchBaOrganizationComp(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ba_organization_comp set "
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
	query += " where business_associate_id = :id"

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

func DeleteBaOrganizationComp(c *fiber.Ctx) error {
	id := c.Params("id")
	var ba_organization_comp dto.Ba_organization_comp
	ba_organization_comp.Business_associate_id = id

	stmt, err := db.Prepare("delete from ba_organization_comp where business_associate_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ba_organization_comp.Business_associate_id)
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


