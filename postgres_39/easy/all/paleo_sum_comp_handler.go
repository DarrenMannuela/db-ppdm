package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPaleoSumComp(c *fiber.Ctx) error {
	rows, err := db.Query("select * from paleo_sum_comp")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Paleo_sum_comp

	for rows.Next() {
		var paleo_sum_comp dto.Paleo_sum_comp
		if err := rows.Scan(&paleo_sum_comp.Paleo_summary_id, &paleo_sum_comp.Paleo_component_id, &paleo_sum_comp.Active_ind, &paleo_sum_comp.Activity_obs_no, &paleo_sum_comp.Analysis_id, &paleo_sum_comp.Application_id, &paleo_sum_comp.Area_id, &paleo_sum_comp.Area_type, &paleo_sum_comp.Authority_id, &paleo_sum_comp.Ba_organization_id, &paleo_sum_comp.Ba_organization_seq_no, &paleo_sum_comp.Business_associate_id, &paleo_sum_comp.Catalogue_additive_id, &paleo_sum_comp.Catalogue_equip_id, &paleo_sum_comp.Classification_system_id, &paleo_sum_comp.Class_level_id, &paleo_sum_comp.Consent_id, &paleo_sum_comp.Consult_id, &paleo_sum_comp.Contest_id, &paleo_sum_comp.Contract_id, &paleo_sum_comp.Ecozone_id, &paleo_sum_comp.Ecozone_set_id, &paleo_sum_comp.Effective_date, &paleo_sum_comp.Employee_ba_id, &paleo_sum_comp.Employee_obs_no, &paleo_sum_comp.Employer_ba_id, &paleo_sum_comp.Entitlement_id, &paleo_sum_comp.Equipment_id, &paleo_sum_comp.Expiry_date, &paleo_sum_comp.Facility_id, &paleo_sum_comp.Facility_type, &paleo_sum_comp.Field_id, &paleo_sum_comp.Field_station_id, &paleo_sum_comp.Finance_id, &paleo_sum_comp.Fossil_id, &paleo_sum_comp.Incident_id, &paleo_sum_comp.Incident_set_id, &paleo_sum_comp.Incident_type, &paleo_sum_comp.Information_item_id, &paleo_sum_comp.Info_item_subtype, &paleo_sum_comp.Instrument_id, &paleo_sum_comp.Interest_set_id, &paleo_sum_comp.Interest_set_seq_no, &paleo_sum_comp.Jurisdiction, &paleo_sum_comp.Land_right_id, &paleo_sum_comp.Land_right_subtype, &paleo_sum_comp.Land_sale_number, &paleo_sum_comp.Language, &paleo_sum_comp.Lithology_log_id, &paleo_sum_comp.Lith_log_source, &paleo_sum_comp.Meas_section_id, &paleo_sum_comp.Meas_section_source, &paleo_sum_comp.Notification_id, &paleo_sum_comp.Obligation_id, &paleo_sum_comp.Obligation_seq_no, &paleo_sum_comp.Paleo_summary_component_type, &paleo_sum_comp.Pden_id, &paleo_sum_comp.Pden_source, &paleo_sum_comp.Pden_subtype, &paleo_sum_comp.Physical_item_id, &paleo_sum_comp.Pool_id, &paleo_sum_comp.Ppdm_guid, &paleo_sum_comp.Ppdm_system_id, &paleo_sum_comp.Ppdm_table_name, &paleo_sum_comp.Product_type, &paleo_sum_comp.Prod_string_id, &paleo_sum_comp.Prod_string_source, &paleo_sum_comp.Project_id, &paleo_sum_comp.Rate_schedule_id, &paleo_sum_comp.Referenced_guid, &paleo_sum_comp.Referenced_system_id, &paleo_sum_comp.Referenced_table_name, &paleo_sum_comp.Remark, &paleo_sum_comp.Resent_id, &paleo_sum_comp.Reserve_class_id, &paleo_sum_comp.Sample_anal_source, &paleo_sum_comp.Seis_set_id, &paleo_sum_comp.Seis_set_subtype, &paleo_sum_comp.Sf_subtype, &paleo_sum_comp.Source, &paleo_sum_comp.Source_document_id, &paleo_sum_comp.Spatial_description_id, &paleo_sum_comp.Spatial_obs_no, &paleo_sum_comp.Store_id, &paleo_sum_comp.Store_structure_obs_no, &paleo_sum_comp.Strat_name_set_id, &paleo_sum_comp.Strat_unit_id, &paleo_sum_comp.Support_facility_id, &paleo_sum_comp.Sw_application_id, &paleo_sum_comp.Thesaurus_id, &paleo_sum_comp.Thesaurus_word_id, &paleo_sum_comp.Uwi, &paleo_sum_comp.Well_activity_set_id, &paleo_sum_comp.Well_activity_source, &paleo_sum_comp.Well_activity_type_id, &paleo_sum_comp.Well_set_id, &paleo_sum_comp.Work_order_id, &paleo_sum_comp.Row_changed_by, &paleo_sum_comp.Row_changed_date, &paleo_sum_comp.Row_created_by, &paleo_sum_comp.Row_created_date, &paleo_sum_comp.Row_effective_date, &paleo_sum_comp.Row_expiry_date, &paleo_sum_comp.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append paleo_sum_comp to result
		result = append(result, paleo_sum_comp)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPaleoSumComp(c *fiber.Ctx) error {
	var paleo_sum_comp dto.Paleo_sum_comp

	setDefaults(&paleo_sum_comp)

	if err := json.Unmarshal(c.Body(), &paleo_sum_comp); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into paleo_sum_comp values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103, :104)")
	if err != nil {
		return err
	}
	paleo_sum_comp.Row_created_date = formatDate(paleo_sum_comp.Row_created_date)
	paleo_sum_comp.Row_changed_date = formatDate(paleo_sum_comp.Row_changed_date)
	paleo_sum_comp.Effective_date = formatDateString(paleo_sum_comp.Effective_date)
	paleo_sum_comp.Expiry_date = formatDateString(paleo_sum_comp.Expiry_date)
	paleo_sum_comp.Row_effective_date = formatDateString(paleo_sum_comp.Row_effective_date)
	paleo_sum_comp.Row_expiry_date = formatDateString(paleo_sum_comp.Row_expiry_date)






	rows, err := stmt.Exec(paleo_sum_comp.Paleo_summary_id, paleo_sum_comp.Paleo_component_id, paleo_sum_comp.Active_ind, paleo_sum_comp.Activity_obs_no, paleo_sum_comp.Analysis_id, paleo_sum_comp.Application_id, paleo_sum_comp.Area_id, paleo_sum_comp.Area_type, paleo_sum_comp.Authority_id, paleo_sum_comp.Ba_organization_id, paleo_sum_comp.Ba_organization_seq_no, paleo_sum_comp.Business_associate_id, paleo_sum_comp.Catalogue_additive_id, paleo_sum_comp.Catalogue_equip_id, paleo_sum_comp.Classification_system_id, paleo_sum_comp.Class_level_id, paleo_sum_comp.Consent_id, paleo_sum_comp.Consult_id, paleo_sum_comp.Contest_id, paleo_sum_comp.Contract_id, paleo_sum_comp.Ecozone_id, paleo_sum_comp.Ecozone_set_id, paleo_sum_comp.Effective_date, paleo_sum_comp.Employee_ba_id, paleo_sum_comp.Employee_obs_no, paleo_sum_comp.Employer_ba_id, paleo_sum_comp.Entitlement_id, paleo_sum_comp.Equipment_id, paleo_sum_comp.Expiry_date, paleo_sum_comp.Facility_id, paleo_sum_comp.Facility_type, paleo_sum_comp.Field_id, paleo_sum_comp.Field_station_id, paleo_sum_comp.Finance_id, paleo_sum_comp.Fossil_id, paleo_sum_comp.Incident_id, paleo_sum_comp.Incident_set_id, paleo_sum_comp.Incident_type, paleo_sum_comp.Information_item_id, paleo_sum_comp.Info_item_subtype, paleo_sum_comp.Instrument_id, paleo_sum_comp.Interest_set_id, paleo_sum_comp.Interest_set_seq_no, paleo_sum_comp.Jurisdiction, paleo_sum_comp.Land_right_id, paleo_sum_comp.Land_right_subtype, paleo_sum_comp.Land_sale_number, paleo_sum_comp.Language, paleo_sum_comp.Lithology_log_id, paleo_sum_comp.Lith_log_source, paleo_sum_comp.Meas_section_id, paleo_sum_comp.Meas_section_source, paleo_sum_comp.Notification_id, paleo_sum_comp.Obligation_id, paleo_sum_comp.Obligation_seq_no, paleo_sum_comp.Paleo_summary_component_type, paleo_sum_comp.Pden_id, paleo_sum_comp.Pden_source, paleo_sum_comp.Pden_subtype, paleo_sum_comp.Physical_item_id, paleo_sum_comp.Pool_id, paleo_sum_comp.Ppdm_guid, paleo_sum_comp.Ppdm_system_id, paleo_sum_comp.Ppdm_table_name, paleo_sum_comp.Product_type, paleo_sum_comp.Prod_string_id, paleo_sum_comp.Prod_string_source, paleo_sum_comp.Project_id, paleo_sum_comp.Rate_schedule_id, paleo_sum_comp.Referenced_guid, paleo_sum_comp.Referenced_system_id, paleo_sum_comp.Referenced_table_name, paleo_sum_comp.Remark, paleo_sum_comp.Resent_id, paleo_sum_comp.Reserve_class_id, paleo_sum_comp.Sample_anal_source, paleo_sum_comp.Seis_set_id, paleo_sum_comp.Seis_set_subtype, paleo_sum_comp.Sf_subtype, paleo_sum_comp.Source, paleo_sum_comp.Source_document_id, paleo_sum_comp.Spatial_description_id, paleo_sum_comp.Spatial_obs_no, paleo_sum_comp.Store_id, paleo_sum_comp.Store_structure_obs_no, paleo_sum_comp.Strat_name_set_id, paleo_sum_comp.Strat_unit_id, paleo_sum_comp.Support_facility_id, paleo_sum_comp.Sw_application_id, paleo_sum_comp.Thesaurus_id, paleo_sum_comp.Thesaurus_word_id, paleo_sum_comp.Uwi, paleo_sum_comp.Well_activity_set_id, paleo_sum_comp.Well_activity_source, paleo_sum_comp.Well_activity_type_id, paleo_sum_comp.Well_set_id, paleo_sum_comp.Work_order_id, paleo_sum_comp.Row_changed_by, paleo_sum_comp.Row_changed_date, paleo_sum_comp.Row_created_by, paleo_sum_comp.Row_created_date, paleo_sum_comp.Row_effective_date, paleo_sum_comp.Row_expiry_date, paleo_sum_comp.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePaleoSumComp(c *fiber.Ctx) error {
	var paleo_sum_comp dto.Paleo_sum_comp

	setDefaults(&paleo_sum_comp)

	if err := json.Unmarshal(c.Body(), &paleo_sum_comp); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	paleo_sum_comp.Paleo_summary_id = id

    if paleo_sum_comp.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update paleo_sum_comp set paleo_component_id = :1, active_ind = :2, activity_obs_no = :3, analysis_id = :4, application_id = :5, area_id = :6, area_type = :7, authority_id = :8, ba_organization_id = :9, ba_organization_seq_no = :10, business_associate_id = :11, catalogue_additive_id = :12, catalogue_equip_id = :13, classification_system_id = :14, class_level_id = :15, consent_id = :16, consult_id = :17, contest_id = :18, contract_id = :19, ecozone_id = :20, ecozone_set_id = :21, effective_date = :22, employee_ba_id = :23, employee_obs_no = :24, employer_ba_id = :25, entitlement_id = :26, equipment_id = :27, expiry_date = :28, facility_id = :29, facility_type = :30, field_id = :31, field_station_id = :32, finance_id = :33, fossil_id = :34, incident_id = :35, incident_set_id = :36, incident_type = :37, information_item_id = :38, info_item_subtype = :39, instrument_id = :40, interest_set_id = :41, interest_set_seq_no = :42, jurisdiction = :43, land_right_id = :44, land_right_subtype = :45, land_sale_number = :46, language = :47, lithology_log_id = :48, lith_log_source = :49, meas_section_id = :50, meas_section_source = :51, notification_id = :52, obligation_id = :53, obligation_seq_no = :54, paleo_summary_component_type = :55, pden_id = :56, pden_source = :57, pden_subtype = :58, physical_item_id = :59, pool_id = :60, ppdm_guid = :61, ppdm_system_id = :62, ppdm_table_name = :63, product_type = :64, prod_string_id = :65, prod_string_source = :66, project_id = :67, rate_schedule_id = :68, referenced_guid = :69, referenced_system_id = :70, referenced_table_name = :71, remark = :72, resent_id = :73, reserve_class_id = :74, sample_anal_source = :75, seis_set_id = :76, seis_set_subtype = :77, sf_subtype = :78, source = :79, source_document_id = :80, spatial_description_id = :81, spatial_obs_no = :82, store_id = :83, store_structure_obs_no = :84, strat_name_set_id = :85, strat_unit_id = :86, support_facility_id = :87, sw_application_id = :88, thesaurus_id = :89, thesaurus_word_id = :90, uwi = :91, well_activity_set_id = :92, well_activity_source = :93, well_activity_type_id = :94, well_set_id = :95, work_order_id = :96, row_changed_by = :97, row_changed_date = :98, row_created_by = :99, row_effective_date = :100, row_expiry_date = :101, row_quality = :102 where paleo_summary_id = :104")
	if err != nil {
		return err
	}

	paleo_sum_comp.Row_changed_date = formatDate(paleo_sum_comp.Row_changed_date)
	paleo_sum_comp.Effective_date = formatDateString(paleo_sum_comp.Effective_date)
	paleo_sum_comp.Expiry_date = formatDateString(paleo_sum_comp.Expiry_date)
	paleo_sum_comp.Row_effective_date = formatDateString(paleo_sum_comp.Row_effective_date)
	paleo_sum_comp.Row_expiry_date = formatDateString(paleo_sum_comp.Row_expiry_date)






	rows, err := stmt.Exec(paleo_sum_comp.Paleo_component_id, paleo_sum_comp.Active_ind, paleo_sum_comp.Activity_obs_no, paleo_sum_comp.Analysis_id, paleo_sum_comp.Application_id, paleo_sum_comp.Area_id, paleo_sum_comp.Area_type, paleo_sum_comp.Authority_id, paleo_sum_comp.Ba_organization_id, paleo_sum_comp.Ba_organization_seq_no, paleo_sum_comp.Business_associate_id, paleo_sum_comp.Catalogue_additive_id, paleo_sum_comp.Catalogue_equip_id, paleo_sum_comp.Classification_system_id, paleo_sum_comp.Class_level_id, paleo_sum_comp.Consent_id, paleo_sum_comp.Consult_id, paleo_sum_comp.Contest_id, paleo_sum_comp.Contract_id, paleo_sum_comp.Ecozone_id, paleo_sum_comp.Ecozone_set_id, paleo_sum_comp.Effective_date, paleo_sum_comp.Employee_ba_id, paleo_sum_comp.Employee_obs_no, paleo_sum_comp.Employer_ba_id, paleo_sum_comp.Entitlement_id, paleo_sum_comp.Equipment_id, paleo_sum_comp.Expiry_date, paleo_sum_comp.Facility_id, paleo_sum_comp.Facility_type, paleo_sum_comp.Field_id, paleo_sum_comp.Field_station_id, paleo_sum_comp.Finance_id, paleo_sum_comp.Fossil_id, paleo_sum_comp.Incident_id, paleo_sum_comp.Incident_set_id, paleo_sum_comp.Incident_type, paleo_sum_comp.Information_item_id, paleo_sum_comp.Info_item_subtype, paleo_sum_comp.Instrument_id, paleo_sum_comp.Interest_set_id, paleo_sum_comp.Interest_set_seq_no, paleo_sum_comp.Jurisdiction, paleo_sum_comp.Land_right_id, paleo_sum_comp.Land_right_subtype, paleo_sum_comp.Land_sale_number, paleo_sum_comp.Language, paleo_sum_comp.Lithology_log_id, paleo_sum_comp.Lith_log_source, paleo_sum_comp.Meas_section_id, paleo_sum_comp.Meas_section_source, paleo_sum_comp.Notification_id, paleo_sum_comp.Obligation_id, paleo_sum_comp.Obligation_seq_no, paleo_sum_comp.Paleo_summary_component_type, paleo_sum_comp.Pden_id, paleo_sum_comp.Pden_source, paleo_sum_comp.Pden_subtype, paleo_sum_comp.Physical_item_id, paleo_sum_comp.Pool_id, paleo_sum_comp.Ppdm_guid, paleo_sum_comp.Ppdm_system_id, paleo_sum_comp.Ppdm_table_name, paleo_sum_comp.Product_type, paleo_sum_comp.Prod_string_id, paleo_sum_comp.Prod_string_source, paleo_sum_comp.Project_id, paleo_sum_comp.Rate_schedule_id, paleo_sum_comp.Referenced_guid, paleo_sum_comp.Referenced_system_id, paleo_sum_comp.Referenced_table_name, paleo_sum_comp.Remark, paleo_sum_comp.Resent_id, paleo_sum_comp.Reserve_class_id, paleo_sum_comp.Sample_anal_source, paleo_sum_comp.Seis_set_id, paleo_sum_comp.Seis_set_subtype, paleo_sum_comp.Sf_subtype, paleo_sum_comp.Source, paleo_sum_comp.Source_document_id, paleo_sum_comp.Spatial_description_id, paleo_sum_comp.Spatial_obs_no, paleo_sum_comp.Store_id, paleo_sum_comp.Store_structure_obs_no, paleo_sum_comp.Strat_name_set_id, paleo_sum_comp.Strat_unit_id, paleo_sum_comp.Support_facility_id, paleo_sum_comp.Sw_application_id, paleo_sum_comp.Thesaurus_id, paleo_sum_comp.Thesaurus_word_id, paleo_sum_comp.Uwi, paleo_sum_comp.Well_activity_set_id, paleo_sum_comp.Well_activity_source, paleo_sum_comp.Well_activity_type_id, paleo_sum_comp.Well_set_id, paleo_sum_comp.Work_order_id, paleo_sum_comp.Row_changed_by, paleo_sum_comp.Row_changed_date, paleo_sum_comp.Row_created_by, paleo_sum_comp.Row_effective_date, paleo_sum_comp.Row_expiry_date, paleo_sum_comp.Row_quality, paleo_sum_comp.Paleo_summary_id)
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

func PatchPaleoSumComp(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update paleo_sum_comp set "
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
	query += " where paleo_summary_id = :id"

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

func DeletePaleoSumComp(c *fiber.Ctx) error {
	id := c.Params("id")
	var paleo_sum_comp dto.Paleo_sum_comp
	paleo_sum_comp.Paleo_summary_id = id

	stmt, err := db.Prepare("delete from paleo_sum_comp where paleo_summary_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(paleo_sum_comp.Paleo_summary_id)
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


