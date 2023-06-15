package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetResentComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from resent_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Resent_component

	for rows.Next() {
		var resent_component dto.Resent_component
		if err := rows.Scan(&resent_component.Resent_id, &resent_component.Component_id, &resent_component.Active_ind, &resent_component.Activity_obs_no, &resent_component.Analysis_id, &resent_component.Application_id, &resent_component.Area_id, &resent_component.Area_type, &resent_component.Authority_id, &resent_component.Ba_organization_id, &resent_component.Ba_organization_seq_no, &resent_component.Business_associate_id, &resent_component.Catalogue_additive_id, &resent_component.Catalogue_equip_id, &resent_component.Classification_system_id, &resent_component.Class_level_id, &resent_component.Component_resent_id, &resent_component.Consent_id, &resent_component.Consult_id, &resent_component.Contest_id, &resent_component.Contract_id, &resent_component.Ecozone_id, &resent_component.Effective_date, &resent_component.Employee_ba_id, &resent_component.Employee_obs_no, &resent_component.Employer_ba_id, &resent_component.Entitlement_id, &resent_component.Equipment_id, &resent_component.Expiry_date, &resent_component.Facility_id, &resent_component.Facility_type, &resent_component.Field_id, &resent_component.Field_station_id, &resent_component.Finance_id, &resent_component.Fossil_id, &resent_component.Incident_id, &resent_component.Incident_set_id, &resent_component.Incident_type, &resent_component.Information_item_id, &resent_component.Info_item_subtype, &resent_component.Instrument_id, &resent_component.Interest_set_id, &resent_component.Interest_set_seq_no, &resent_component.Jurisdiction, &resent_component.Land_right_id, &resent_component.Land_right_subtype, &resent_component.Land_sale_number, &resent_component.Language, &resent_component.Lithology_log_id, &resent_component.Lith_log_source, &resent_component.Notification_id, &resent_component.Obligation_id, &resent_component.Obligation_seq_no, &resent_component.Paleo_summary_id, &resent_component.Pden_id, &resent_component.Pden_source, &resent_component.Pden_subtype, &resent_component.Percent_contribution, &resent_component.Physical_item_id, &resent_component.Pool_id, &resent_component.Ppdm_guid, &resent_component.Ppdm_system_id, &resent_component.Ppdm_table_name, &resent_component.Product_type, &resent_component.Prod_string_id, &resent_component.Prod_string_source, &resent_component.Project_id, &resent_component.Pr_str_form_obs_no, &resent_component.Rate_schedule_id, &resent_component.Referenced_guid, &resent_component.Referenced_system_id, &resent_component.Referenced_table_name, &resent_component.Remark, &resent_component.Report_ind, &resent_component.Resent_component_type, &resent_component.Reserve_class_id, &resent_component.Sample_anal_source, &resent_component.Seis_set_id, &resent_component.Seis_set_subtype, &resent_component.Sf_subtype, &resent_component.Source, &resent_component.Spatial_description_id, &resent_component.Spatial_obs_no, &resent_component.Store_id, &resent_component.Store_structure_obs_no, &resent_component.Strat_name_set_id, &resent_component.Strat_unit_id, &resent_component.Support_facility_id, &resent_component.Sw_application_id, &resent_component.Thesaurus_id, &resent_component.Thesaurus_word_id, &resent_component.Uwi, &resent_component.Well_activity_set_id, &resent_component.Well_activity_source, &resent_component.Well_activity_type_id, &resent_component.Well_set_id, &resent_component.Work_order_id, &resent_component.Row_changed_by, &resent_component.Row_changed_date, &resent_component.Row_created_by, &resent_component.Row_created_date, &resent_component.Row_effective_date, &resent_component.Row_expiry_date, &resent_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append resent_component to result
		result = append(result, resent_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetResentComponent(c *fiber.Ctx) error {
	var resent_component dto.Resent_component

	setDefaults(&resent_component)

	if err := json.Unmarshal(c.Body(), &resent_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into resent_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103, :104)")
	if err != nil {
		return err
	}
	resent_component.Row_created_date = formatDate(resent_component.Row_created_date)
	resent_component.Row_changed_date = formatDate(resent_component.Row_changed_date)
	resent_component.Effective_date = formatDateString(resent_component.Effective_date)
	resent_component.Expiry_date = formatDateString(resent_component.Expiry_date)
	resent_component.Row_effective_date = formatDateString(resent_component.Row_effective_date)
	resent_component.Row_expiry_date = formatDateString(resent_component.Row_expiry_date)






	rows, err := stmt.Exec(resent_component.Resent_id, resent_component.Component_id, resent_component.Active_ind, resent_component.Activity_obs_no, resent_component.Analysis_id, resent_component.Application_id, resent_component.Area_id, resent_component.Area_type, resent_component.Authority_id, resent_component.Ba_organization_id, resent_component.Ba_organization_seq_no, resent_component.Business_associate_id, resent_component.Catalogue_additive_id, resent_component.Catalogue_equip_id, resent_component.Classification_system_id, resent_component.Class_level_id, resent_component.Component_resent_id, resent_component.Consent_id, resent_component.Consult_id, resent_component.Contest_id, resent_component.Contract_id, resent_component.Ecozone_id, resent_component.Effective_date, resent_component.Employee_ba_id, resent_component.Employee_obs_no, resent_component.Employer_ba_id, resent_component.Entitlement_id, resent_component.Equipment_id, resent_component.Expiry_date, resent_component.Facility_id, resent_component.Facility_type, resent_component.Field_id, resent_component.Field_station_id, resent_component.Finance_id, resent_component.Fossil_id, resent_component.Incident_id, resent_component.Incident_set_id, resent_component.Incident_type, resent_component.Information_item_id, resent_component.Info_item_subtype, resent_component.Instrument_id, resent_component.Interest_set_id, resent_component.Interest_set_seq_no, resent_component.Jurisdiction, resent_component.Land_right_id, resent_component.Land_right_subtype, resent_component.Land_sale_number, resent_component.Language, resent_component.Lithology_log_id, resent_component.Lith_log_source, resent_component.Notification_id, resent_component.Obligation_id, resent_component.Obligation_seq_no, resent_component.Paleo_summary_id, resent_component.Pden_id, resent_component.Pden_source, resent_component.Pden_subtype, resent_component.Percent_contribution, resent_component.Physical_item_id, resent_component.Pool_id, resent_component.Ppdm_guid, resent_component.Ppdm_system_id, resent_component.Ppdm_table_name, resent_component.Product_type, resent_component.Prod_string_id, resent_component.Prod_string_source, resent_component.Project_id, resent_component.Pr_str_form_obs_no, resent_component.Rate_schedule_id, resent_component.Referenced_guid, resent_component.Referenced_system_id, resent_component.Referenced_table_name, resent_component.Remark, resent_component.Report_ind, resent_component.Resent_component_type, resent_component.Reserve_class_id, resent_component.Sample_anal_source, resent_component.Seis_set_id, resent_component.Seis_set_subtype, resent_component.Sf_subtype, resent_component.Source, resent_component.Spatial_description_id, resent_component.Spatial_obs_no, resent_component.Store_id, resent_component.Store_structure_obs_no, resent_component.Strat_name_set_id, resent_component.Strat_unit_id, resent_component.Support_facility_id, resent_component.Sw_application_id, resent_component.Thesaurus_id, resent_component.Thesaurus_word_id, resent_component.Uwi, resent_component.Well_activity_set_id, resent_component.Well_activity_source, resent_component.Well_activity_type_id, resent_component.Well_set_id, resent_component.Work_order_id, resent_component.Row_changed_by, resent_component.Row_changed_date, resent_component.Row_created_by, resent_component.Row_created_date, resent_component.Row_effective_date, resent_component.Row_expiry_date, resent_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateResentComponent(c *fiber.Ctx) error {
	var resent_component dto.Resent_component

	setDefaults(&resent_component)

	if err := json.Unmarshal(c.Body(), &resent_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	resent_component.Resent_id = id

    if resent_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update resent_component set component_id = :1, active_ind = :2, activity_obs_no = :3, analysis_id = :4, application_id = :5, area_id = :6, area_type = :7, authority_id = :8, ba_organization_id = :9, ba_organization_seq_no = :10, business_associate_id = :11, catalogue_additive_id = :12, catalogue_equip_id = :13, classification_system_id = :14, class_level_id = :15, component_resent_id = :16, consent_id = :17, consult_id = :18, contest_id = :19, contract_id = :20, ecozone_id = :21, effective_date = :22, employee_ba_id = :23, employee_obs_no = :24, employer_ba_id = :25, entitlement_id = :26, equipment_id = :27, expiry_date = :28, facility_id = :29, facility_type = :30, field_id = :31, field_station_id = :32, finance_id = :33, fossil_id = :34, incident_id = :35, incident_set_id = :36, incident_type = :37, information_item_id = :38, info_item_subtype = :39, instrument_id = :40, interest_set_id = :41, interest_set_seq_no = :42, jurisdiction = :43, land_right_id = :44, land_right_subtype = :45, land_sale_number = :46, language = :47, lithology_log_id = :48, lith_log_source = :49, notification_id = :50, obligation_id = :51, obligation_seq_no = :52, paleo_summary_id = :53, pden_id = :54, pden_source = :55, pden_subtype = :56, percent_contribution = :57, physical_item_id = :58, pool_id = :59, ppdm_guid = :60, ppdm_system_id = :61, ppdm_table_name = :62, product_type = :63, prod_string_id = :64, prod_string_source = :65, project_id = :66, pr_str_form_obs_no = :67, rate_schedule_id = :68, referenced_guid = :69, referenced_system_id = :70, referenced_table_name = :71, remark = :72, report_ind = :73, resent_component_type = :74, reserve_class_id = :75, sample_anal_source = :76, seis_set_id = :77, seis_set_subtype = :78, sf_subtype = :79, source = :80, spatial_description_id = :81, spatial_obs_no = :82, store_id = :83, store_structure_obs_no = :84, strat_name_set_id = :85, strat_unit_id = :86, support_facility_id = :87, sw_application_id = :88, thesaurus_id = :89, thesaurus_word_id = :90, uwi = :91, well_activity_set_id = :92, well_activity_source = :93, well_activity_type_id = :94, well_set_id = :95, work_order_id = :96, row_changed_by = :97, row_changed_date = :98, row_created_by = :99, row_effective_date = :100, row_expiry_date = :101, row_quality = :102 where resent_id = :104")
	if err != nil {
		return err
	}

	resent_component.Row_changed_date = formatDate(resent_component.Row_changed_date)
	resent_component.Effective_date = formatDateString(resent_component.Effective_date)
	resent_component.Expiry_date = formatDateString(resent_component.Expiry_date)
	resent_component.Row_effective_date = formatDateString(resent_component.Row_effective_date)
	resent_component.Row_expiry_date = formatDateString(resent_component.Row_expiry_date)






	rows, err := stmt.Exec(resent_component.Component_id, resent_component.Active_ind, resent_component.Activity_obs_no, resent_component.Analysis_id, resent_component.Application_id, resent_component.Area_id, resent_component.Area_type, resent_component.Authority_id, resent_component.Ba_organization_id, resent_component.Ba_organization_seq_no, resent_component.Business_associate_id, resent_component.Catalogue_additive_id, resent_component.Catalogue_equip_id, resent_component.Classification_system_id, resent_component.Class_level_id, resent_component.Component_resent_id, resent_component.Consent_id, resent_component.Consult_id, resent_component.Contest_id, resent_component.Contract_id, resent_component.Ecozone_id, resent_component.Effective_date, resent_component.Employee_ba_id, resent_component.Employee_obs_no, resent_component.Employer_ba_id, resent_component.Entitlement_id, resent_component.Equipment_id, resent_component.Expiry_date, resent_component.Facility_id, resent_component.Facility_type, resent_component.Field_id, resent_component.Field_station_id, resent_component.Finance_id, resent_component.Fossil_id, resent_component.Incident_id, resent_component.Incident_set_id, resent_component.Incident_type, resent_component.Information_item_id, resent_component.Info_item_subtype, resent_component.Instrument_id, resent_component.Interest_set_id, resent_component.Interest_set_seq_no, resent_component.Jurisdiction, resent_component.Land_right_id, resent_component.Land_right_subtype, resent_component.Land_sale_number, resent_component.Language, resent_component.Lithology_log_id, resent_component.Lith_log_source, resent_component.Notification_id, resent_component.Obligation_id, resent_component.Obligation_seq_no, resent_component.Paleo_summary_id, resent_component.Pden_id, resent_component.Pden_source, resent_component.Pden_subtype, resent_component.Percent_contribution, resent_component.Physical_item_id, resent_component.Pool_id, resent_component.Ppdm_guid, resent_component.Ppdm_system_id, resent_component.Ppdm_table_name, resent_component.Product_type, resent_component.Prod_string_id, resent_component.Prod_string_source, resent_component.Project_id, resent_component.Pr_str_form_obs_no, resent_component.Rate_schedule_id, resent_component.Referenced_guid, resent_component.Referenced_system_id, resent_component.Referenced_table_name, resent_component.Remark, resent_component.Report_ind, resent_component.Resent_component_type, resent_component.Reserve_class_id, resent_component.Sample_anal_source, resent_component.Seis_set_id, resent_component.Seis_set_subtype, resent_component.Sf_subtype, resent_component.Source, resent_component.Spatial_description_id, resent_component.Spatial_obs_no, resent_component.Store_id, resent_component.Store_structure_obs_no, resent_component.Strat_name_set_id, resent_component.Strat_unit_id, resent_component.Support_facility_id, resent_component.Sw_application_id, resent_component.Thesaurus_id, resent_component.Thesaurus_word_id, resent_component.Uwi, resent_component.Well_activity_set_id, resent_component.Well_activity_source, resent_component.Well_activity_type_id, resent_component.Well_set_id, resent_component.Work_order_id, resent_component.Row_changed_by, resent_component.Row_changed_date, resent_component.Row_created_by, resent_component.Row_effective_date, resent_component.Row_expiry_date, resent_component.Row_quality, resent_component.Resent_id)
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

func PatchResentComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update resent_component set "
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
	query += " where resent_id = :id"

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

func DeleteResentComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var resent_component dto.Resent_component
	resent_component.Resent_id = id

	stmt, err := db.Prepare("delete from resent_component where resent_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(resent_component.Resent_id)
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


