package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetStratUnitComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from strat_unit_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Strat_unit_component

	for rows.Next() {
		var strat_unit_component dto.Strat_unit_component
		if err := rows.Scan(&strat_unit_component.Strat_name_set_id, &strat_unit_component.Strat_unit_id, &strat_unit_component.Component_obs_no, &strat_unit_component.Active_ind, &strat_unit_component.Activity_obs_no, &strat_unit_component.Analysis_id, &strat_unit_component.Application_id, &strat_unit_component.Area_id, &strat_unit_component.Area_type, &strat_unit_component.Authority_id, &strat_unit_component.Ba_organization_id, &strat_unit_component.Ba_organization_seq_no, &strat_unit_component.Business_associate_id, &strat_unit_component.Catalogue_additive_id, &strat_unit_component.Catalogue_equip_id, &strat_unit_component.Classification_system_id, &strat_unit_component.Class_level_id, &strat_unit_component.Consent_id, &strat_unit_component.Consult_id, &strat_unit_component.Contest_id, &strat_unit_component.Contract_id, &strat_unit_component.Ecozone_id, &strat_unit_component.Effective_date, &strat_unit_component.Employee_ba_id, &strat_unit_component.Employee_obs_no, &strat_unit_component.Employer_ba_id, &strat_unit_component.Entitlement_id, &strat_unit_component.Equipment_id, &strat_unit_component.Expiry_date, &strat_unit_component.Facility_id, &strat_unit_component.Facility_type, &strat_unit_component.Field_id, &strat_unit_component.Field_station_id, &strat_unit_component.Finance_id, &strat_unit_component.Fossil_id, &strat_unit_component.Incident_id, &strat_unit_component.Incident_set_id, &strat_unit_component.Incident_type, &strat_unit_component.Information_item_id, &strat_unit_component.Info_item_subtype, &strat_unit_component.Instrument_id, &strat_unit_component.Interest_set_id, &strat_unit_component.Interest_set_seq_no, &strat_unit_component.Jurisdiction, &strat_unit_component.Land_right_id, &strat_unit_component.Land_right_subtype, &strat_unit_component.Land_sale_number, &strat_unit_component.Language, &strat_unit_component.Lithology_log_id, &strat_unit_component.Lith_log_source, &strat_unit_component.Notification_id, &strat_unit_component.Obligation_id, &strat_unit_component.Obligation_seq_no, &strat_unit_component.Paleo_summary_id, &strat_unit_component.Pden_id, &strat_unit_component.Pden_source, &strat_unit_component.Pden_subtype, &strat_unit_component.Physical_item_id, &strat_unit_component.Pool_id, &strat_unit_component.Ppdm_guid, &strat_unit_component.Ppdm_system_id, &strat_unit_component.Ppdm_table_name, &strat_unit_component.Product_type, &strat_unit_component.Project_id, &strat_unit_component.Pr_str_source, &strat_unit_component.Pr_str_uwi, &strat_unit_component.Rate_schedule_id, &strat_unit_component.Referenced_guid, &strat_unit_component.Referenced_system_id, &strat_unit_component.Referenced_table_name, &strat_unit_component.Remark, &strat_unit_component.Resent_id, &strat_unit_component.Reserve_class_id, &strat_unit_component.Sample_anal_source, &strat_unit_component.Seis_set_id, &strat_unit_component.Seis_set_subtype, &strat_unit_component.Sf_subtype, &strat_unit_component.Source, &strat_unit_component.Spatial_description_id, &strat_unit_component.Spatial_obs_no, &strat_unit_component.Store_id, &strat_unit_component.Store_structure_obs_no, &strat_unit_component.Strat_unit_component_type, &strat_unit_component.String_id, &strat_unit_component.Support_facility_id, &strat_unit_component.Sw_application_id, &strat_unit_component.Thesaurus_id, &strat_unit_component.Thesaurus_word_id, &strat_unit_component.Uwi, &strat_unit_component.Well_activity_set_id, &strat_unit_component.Well_activity_source, &strat_unit_component.Well_activity_type_id, &strat_unit_component.Well_activity_uwi, &strat_unit_component.Well_set_id, &strat_unit_component.Work_order_id, &strat_unit_component.Row_changed_by, &strat_unit_component.Row_changed_date, &strat_unit_component.Row_created_by, &strat_unit_component.Row_created_date, &strat_unit_component.Row_effective_date, &strat_unit_component.Row_expiry_date, &strat_unit_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append strat_unit_component to result
		result = append(result, strat_unit_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetStratUnitComponent(c *fiber.Ctx) error {
	var strat_unit_component dto.Strat_unit_component

	setDefaults(&strat_unit_component)

	if err := json.Unmarshal(c.Body(), &strat_unit_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into strat_unit_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102)")
	if err != nil {
		return err
	}
	strat_unit_component.Row_created_date = formatDate(strat_unit_component.Row_created_date)
	strat_unit_component.Row_changed_date = formatDate(strat_unit_component.Row_changed_date)
	strat_unit_component.Effective_date = formatDateString(strat_unit_component.Effective_date)
	strat_unit_component.Expiry_date = formatDateString(strat_unit_component.Expiry_date)
	strat_unit_component.Row_effective_date = formatDateString(strat_unit_component.Row_effective_date)
	strat_unit_component.Row_expiry_date = formatDateString(strat_unit_component.Row_expiry_date)






	rows, err := stmt.Exec(strat_unit_component.Strat_name_set_id, strat_unit_component.Strat_unit_id, strat_unit_component.Component_obs_no, strat_unit_component.Active_ind, strat_unit_component.Activity_obs_no, strat_unit_component.Analysis_id, strat_unit_component.Application_id, strat_unit_component.Area_id, strat_unit_component.Area_type, strat_unit_component.Authority_id, strat_unit_component.Ba_organization_id, strat_unit_component.Ba_organization_seq_no, strat_unit_component.Business_associate_id, strat_unit_component.Catalogue_additive_id, strat_unit_component.Catalogue_equip_id, strat_unit_component.Classification_system_id, strat_unit_component.Class_level_id, strat_unit_component.Consent_id, strat_unit_component.Consult_id, strat_unit_component.Contest_id, strat_unit_component.Contract_id, strat_unit_component.Ecozone_id, strat_unit_component.Effective_date, strat_unit_component.Employee_ba_id, strat_unit_component.Employee_obs_no, strat_unit_component.Employer_ba_id, strat_unit_component.Entitlement_id, strat_unit_component.Equipment_id, strat_unit_component.Expiry_date, strat_unit_component.Facility_id, strat_unit_component.Facility_type, strat_unit_component.Field_id, strat_unit_component.Field_station_id, strat_unit_component.Finance_id, strat_unit_component.Fossil_id, strat_unit_component.Incident_id, strat_unit_component.Incident_set_id, strat_unit_component.Incident_type, strat_unit_component.Information_item_id, strat_unit_component.Info_item_subtype, strat_unit_component.Instrument_id, strat_unit_component.Interest_set_id, strat_unit_component.Interest_set_seq_no, strat_unit_component.Jurisdiction, strat_unit_component.Land_right_id, strat_unit_component.Land_right_subtype, strat_unit_component.Land_sale_number, strat_unit_component.Language, strat_unit_component.Lithology_log_id, strat_unit_component.Lith_log_source, strat_unit_component.Notification_id, strat_unit_component.Obligation_id, strat_unit_component.Obligation_seq_no, strat_unit_component.Paleo_summary_id, strat_unit_component.Pden_id, strat_unit_component.Pden_source, strat_unit_component.Pden_subtype, strat_unit_component.Physical_item_id, strat_unit_component.Pool_id, strat_unit_component.Ppdm_guid, strat_unit_component.Ppdm_system_id, strat_unit_component.Ppdm_table_name, strat_unit_component.Product_type, strat_unit_component.Project_id, strat_unit_component.Pr_str_source, strat_unit_component.Pr_str_uwi, strat_unit_component.Rate_schedule_id, strat_unit_component.Referenced_guid, strat_unit_component.Referenced_system_id, strat_unit_component.Referenced_table_name, strat_unit_component.Remark, strat_unit_component.Resent_id, strat_unit_component.Reserve_class_id, strat_unit_component.Sample_anal_source, strat_unit_component.Seis_set_id, strat_unit_component.Seis_set_subtype, strat_unit_component.Sf_subtype, strat_unit_component.Source, strat_unit_component.Spatial_description_id, strat_unit_component.Spatial_obs_no, strat_unit_component.Store_id, strat_unit_component.Store_structure_obs_no, strat_unit_component.Strat_unit_component_type, strat_unit_component.String_id, strat_unit_component.Support_facility_id, strat_unit_component.Sw_application_id, strat_unit_component.Thesaurus_id, strat_unit_component.Thesaurus_word_id, strat_unit_component.Uwi, strat_unit_component.Well_activity_set_id, strat_unit_component.Well_activity_source, strat_unit_component.Well_activity_type_id, strat_unit_component.Well_activity_uwi, strat_unit_component.Well_set_id, strat_unit_component.Work_order_id, strat_unit_component.Row_changed_by, strat_unit_component.Row_changed_date, strat_unit_component.Row_created_by, strat_unit_component.Row_created_date, strat_unit_component.Row_effective_date, strat_unit_component.Row_expiry_date, strat_unit_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateStratUnitComponent(c *fiber.Ctx) error {
	var strat_unit_component dto.Strat_unit_component

	setDefaults(&strat_unit_component)

	if err := json.Unmarshal(c.Body(), &strat_unit_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	strat_unit_component.Strat_name_set_id = id

    if strat_unit_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update strat_unit_component set strat_unit_id = :1, component_obs_no = :2, active_ind = :3, activity_obs_no = :4, analysis_id = :5, application_id = :6, area_id = :7, area_type = :8, authority_id = :9, ba_organization_id = :10, ba_organization_seq_no = :11, business_associate_id = :12, catalogue_additive_id = :13, catalogue_equip_id = :14, classification_system_id = :15, class_level_id = :16, consent_id = :17, consult_id = :18, contest_id = :19, contract_id = :20, ecozone_id = :21, effective_date = :22, employee_ba_id = :23, employee_obs_no = :24, employer_ba_id = :25, entitlement_id = :26, equipment_id = :27, expiry_date = :28, facility_id = :29, facility_type = :30, field_id = :31, field_station_id = :32, finance_id = :33, fossil_id = :34, incident_id = :35, incident_set_id = :36, incident_type = :37, information_item_id = :38, info_item_subtype = :39, instrument_id = :40, interest_set_id = :41, interest_set_seq_no = :42, jurisdiction = :43, land_right_id = :44, land_right_subtype = :45, land_sale_number = :46, language = :47, lithology_log_id = :48, lith_log_source = :49, notification_id = :50, obligation_id = :51, obligation_seq_no = :52, paleo_summary_id = :53, pden_id = :54, pden_source = :55, pden_subtype = :56, physical_item_id = :57, pool_id = :58, ppdm_guid = :59, ppdm_system_id = :60, ppdm_table_name = :61, product_type = :62, project_id = :63, pr_str_source = :64, pr_str_uwi = :65, rate_schedule_id = :66, referenced_guid = :67, referenced_system_id = :68, referenced_table_name = :69, remark = :70, resent_id = :71, reserve_class_id = :72, sample_anal_source = :73, seis_set_id = :74, seis_set_subtype = :75, sf_subtype = :76, source = :77, spatial_description_id = :78, spatial_obs_no = :79, store_id = :80, store_structure_obs_no = :81, strat_unit_component_type = :82, string_id = :83, support_facility_id = :84, sw_application_id = :85, thesaurus_id = :86, thesaurus_word_id = :87, uwi = :88, well_activity_set_id = :89, well_activity_source = :90, well_activity_type_id = :91, well_activity_uwi = :92, well_set_id = :93, work_order_id = :94, row_changed_by = :95, row_changed_date = :96, row_created_by = :97, row_effective_date = :98, row_expiry_date = :99, row_quality = :100 where strat_name_set_id = :102")
	if err != nil {
		return err
	}

	strat_unit_component.Row_changed_date = formatDate(strat_unit_component.Row_changed_date)
	strat_unit_component.Effective_date = formatDateString(strat_unit_component.Effective_date)
	strat_unit_component.Expiry_date = formatDateString(strat_unit_component.Expiry_date)
	strat_unit_component.Row_effective_date = formatDateString(strat_unit_component.Row_effective_date)
	strat_unit_component.Row_expiry_date = formatDateString(strat_unit_component.Row_expiry_date)






	rows, err := stmt.Exec(strat_unit_component.Strat_unit_id, strat_unit_component.Component_obs_no, strat_unit_component.Active_ind, strat_unit_component.Activity_obs_no, strat_unit_component.Analysis_id, strat_unit_component.Application_id, strat_unit_component.Area_id, strat_unit_component.Area_type, strat_unit_component.Authority_id, strat_unit_component.Ba_organization_id, strat_unit_component.Ba_organization_seq_no, strat_unit_component.Business_associate_id, strat_unit_component.Catalogue_additive_id, strat_unit_component.Catalogue_equip_id, strat_unit_component.Classification_system_id, strat_unit_component.Class_level_id, strat_unit_component.Consent_id, strat_unit_component.Consult_id, strat_unit_component.Contest_id, strat_unit_component.Contract_id, strat_unit_component.Ecozone_id, strat_unit_component.Effective_date, strat_unit_component.Employee_ba_id, strat_unit_component.Employee_obs_no, strat_unit_component.Employer_ba_id, strat_unit_component.Entitlement_id, strat_unit_component.Equipment_id, strat_unit_component.Expiry_date, strat_unit_component.Facility_id, strat_unit_component.Facility_type, strat_unit_component.Field_id, strat_unit_component.Field_station_id, strat_unit_component.Finance_id, strat_unit_component.Fossil_id, strat_unit_component.Incident_id, strat_unit_component.Incident_set_id, strat_unit_component.Incident_type, strat_unit_component.Information_item_id, strat_unit_component.Info_item_subtype, strat_unit_component.Instrument_id, strat_unit_component.Interest_set_id, strat_unit_component.Interest_set_seq_no, strat_unit_component.Jurisdiction, strat_unit_component.Land_right_id, strat_unit_component.Land_right_subtype, strat_unit_component.Land_sale_number, strat_unit_component.Language, strat_unit_component.Lithology_log_id, strat_unit_component.Lith_log_source, strat_unit_component.Notification_id, strat_unit_component.Obligation_id, strat_unit_component.Obligation_seq_no, strat_unit_component.Paleo_summary_id, strat_unit_component.Pden_id, strat_unit_component.Pden_source, strat_unit_component.Pden_subtype, strat_unit_component.Physical_item_id, strat_unit_component.Pool_id, strat_unit_component.Ppdm_guid, strat_unit_component.Ppdm_system_id, strat_unit_component.Ppdm_table_name, strat_unit_component.Product_type, strat_unit_component.Project_id, strat_unit_component.Pr_str_source, strat_unit_component.Pr_str_uwi, strat_unit_component.Rate_schedule_id, strat_unit_component.Referenced_guid, strat_unit_component.Referenced_system_id, strat_unit_component.Referenced_table_name, strat_unit_component.Remark, strat_unit_component.Resent_id, strat_unit_component.Reserve_class_id, strat_unit_component.Sample_anal_source, strat_unit_component.Seis_set_id, strat_unit_component.Seis_set_subtype, strat_unit_component.Sf_subtype, strat_unit_component.Source, strat_unit_component.Spatial_description_id, strat_unit_component.Spatial_obs_no, strat_unit_component.Store_id, strat_unit_component.Store_structure_obs_no, strat_unit_component.Strat_unit_component_type, strat_unit_component.String_id, strat_unit_component.Support_facility_id, strat_unit_component.Sw_application_id, strat_unit_component.Thesaurus_id, strat_unit_component.Thesaurus_word_id, strat_unit_component.Uwi, strat_unit_component.Well_activity_set_id, strat_unit_component.Well_activity_source, strat_unit_component.Well_activity_type_id, strat_unit_component.Well_activity_uwi, strat_unit_component.Well_set_id, strat_unit_component.Work_order_id, strat_unit_component.Row_changed_by, strat_unit_component.Row_changed_date, strat_unit_component.Row_created_by, strat_unit_component.Row_effective_date, strat_unit_component.Row_expiry_date, strat_unit_component.Row_quality, strat_unit_component.Strat_name_set_id)
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

func PatchStratUnitComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update strat_unit_component set "
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
	query += " where strat_name_set_id = :id"

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

func DeleteStratUnitComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var strat_unit_component dto.Strat_unit_component
	strat_unit_component.Strat_name_set_id = id

	stmt, err := db.Prepare("delete from strat_unit_component where strat_name_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(strat_unit_component.Strat_name_set_id)
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


