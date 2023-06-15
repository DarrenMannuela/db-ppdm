package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetProjectComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from project_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Project_component

	for rows.Next() {
		var project_component dto.Project_component
		if err := rows.Scan(&project_component.Project_id, &project_component.Component_obs_no, &project_component.Active_ind, &project_component.Activity_obs_no, &project_component.Analysis_id, &project_component.Application_id, &project_component.Area_id, &project_component.Area_type, &project_component.Authority_id, &project_component.Ba_licensee_ba_id, &project_component.Ba_license_id, &project_component.Ba_organization_id, &project_component.Ba_organization_seq_no, &project_component.Business_associate_id, &project_component.Catalogue_additive_id, &project_component.Catalogue_equip_id, &project_component.Classification_system_id, &project_component.Class_level_id, &project_component.Component_reason, &project_component.Component_type, &project_component.Condition_id, &project_component.Consent_id, &project_component.Consult_id, &project_component.Contest_id, &project_component.Contract_id, &project_component.Description, &project_component.Ecozone_id, &project_component.Effective_date, &project_component.Employee_ba_id, &project_component.Employee_obs_no, &project_component.Employer_ba_id, &project_component.Entitlement_id, &project_component.Equipment_id, &project_component.Expiry_date, &project_component.Facility_id, &project_component.Facility_license_id, &project_component.Facility_type, &project_component.Field_id, &project_component.Field_source, &project_component.Field_station_id, &project_component.Field_strat_interp_id, &project_component.Field_strat_unit_id, &project_component.Finance_id, &project_component.Fossil_id, &project_component.Hse_incident_id, &project_component.Incident_set_id, &project_component.Incident_type, &project_component.Information_item_id, &project_component.Info_item_subtype, &project_component.Input_ind, &project_component.Instrument_id, &project_component.Interest_set_id, &project_component.Interest_set_seq_no, &project_component.Jurisdiction, &project_component.Land_right_id, &project_component.Land_right_subtype, &project_component.Land_sale_number, &project_component.Language, &project_component.Lithology_log_id, &project_component.Lith_log_source, &project_component.Lith_sample_id, &project_component.Notification_id, &project_component.Obligation_id, &project_component.Obligation_seq_no, &project_component.Other_project_id, &project_component.Output_ind, &project_component.Paleo_summary_id, &project_component.Pden_id, &project_component.Pden_source, &project_component.Pden_subtype, &project_component.Physical_item_id, &project_component.Pool_id, &project_component.Ppdm_guid, &project_component.Ppdm_system_id, &project_component.Ppdm_table_name, &project_component.Product_type, &project_component.Prod_string_id, &project_component.Prod_string_source, &project_component.Project_plan_id, &project_component.Project_step_id, &project_component.Provision_id, &project_component.Rate_schedule_id, &project_component.Referenced_guid, &project_component.Referenced_system_id, &project_component.Referenced_table_name, &project_component.Remark, &project_component.Resent_id, &project_component.Reserve_class_id, &project_component.Reserve_product_type, &project_component.Reserve_revision_id, &project_component.Reserve_revision_obs_no, &project_component.Reserve_summary_id, &project_component.Reserve_summary_obs_no, &project_component.Sample_anal_source, &project_component.Seis_license_id, &project_component.Seis_proc_plan_id, &project_component.Seis_set_id, &project_component.Seis_set_subtype, &project_component.Sf_subtype, &project_component.Source, &project_component.Spatial_description_id, &project_component.Spatial_obs_no, &project_component.Store_id, &project_component.Store_structure_obs_no, &project_component.Strat_column_id, &project_component.Strat_column_source, &project_component.Strat_name_set_id, &project_component.Strat_unit_id, &project_component.Support_facility_id, &project_component.Sw_application_id, &project_component.Thesaurus_id, &project_component.Thesaurus_word_id, &project_component.Uwi, &project_component.Well_activity_set_id, &project_component.Well_activity_source, &project_component.Well_activity_type_id, &project_component.Well_int_interp_id, &project_component.Well_int_source, &project_component.Well_int_strat_unit_id, &project_component.Well_int_uwi, &project_component.Well_license_id, &project_component.Well_license_source, &project_component.Well_set_id, &project_component.Work_order_id, &project_component.Row_changed_by, &project_component.Row_changed_date, &project_component.Row_created_by, &project_component.Row_created_date, &project_component.Row_effective_date, &project_component.Row_expiry_date, &project_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append project_component to result
		result = append(result, project_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetProjectComponent(c *fiber.Ctx) error {
	var project_component dto.Project_component

	setDefaults(&project_component)

	if err := json.Unmarshal(c.Body(), &project_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into project_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103, :104, :105, :106, :107, :108, :109, :110, :111, :112, :113, :114, :115, :116, :117, :118, :119, :120, :121, :122, :123, :124, :125, :126, :127, :128, :129, :130, :131)")
	if err != nil {
		return err
	}
	project_component.Row_created_date = formatDate(project_component.Row_created_date)
	project_component.Row_changed_date = formatDate(project_component.Row_changed_date)
	project_component.Effective_date = formatDateString(project_component.Effective_date)
	project_component.Expiry_date = formatDateString(project_component.Expiry_date)
	project_component.Row_effective_date = formatDateString(project_component.Row_effective_date)
	project_component.Row_expiry_date = formatDateString(project_component.Row_expiry_date)






	rows, err := stmt.Exec(project_component.Project_id, project_component.Component_obs_no, project_component.Active_ind, project_component.Activity_obs_no, project_component.Analysis_id, project_component.Application_id, project_component.Area_id, project_component.Area_type, project_component.Authority_id, project_component.Ba_licensee_ba_id, project_component.Ba_license_id, project_component.Ba_organization_id, project_component.Ba_organization_seq_no, project_component.Business_associate_id, project_component.Catalogue_additive_id, project_component.Catalogue_equip_id, project_component.Classification_system_id, project_component.Class_level_id, project_component.Component_reason, project_component.Component_type, project_component.Condition_id, project_component.Consent_id, project_component.Consult_id, project_component.Contest_id, project_component.Contract_id, project_component.Description, project_component.Ecozone_id, project_component.Effective_date, project_component.Employee_ba_id, project_component.Employee_obs_no, project_component.Employer_ba_id, project_component.Entitlement_id, project_component.Equipment_id, project_component.Expiry_date, project_component.Facility_id, project_component.Facility_license_id, project_component.Facility_type, project_component.Field_id, project_component.Field_source, project_component.Field_station_id, project_component.Field_strat_interp_id, project_component.Field_strat_unit_id, project_component.Finance_id, project_component.Fossil_id, project_component.Hse_incident_id, project_component.Incident_set_id, project_component.Incident_type, project_component.Information_item_id, project_component.Info_item_subtype, project_component.Input_ind, project_component.Instrument_id, project_component.Interest_set_id, project_component.Interest_set_seq_no, project_component.Jurisdiction, project_component.Land_right_id, project_component.Land_right_subtype, project_component.Land_sale_number, project_component.Language, project_component.Lithology_log_id, project_component.Lith_log_source, project_component.Lith_sample_id, project_component.Notification_id, project_component.Obligation_id, project_component.Obligation_seq_no, project_component.Other_project_id, project_component.Output_ind, project_component.Paleo_summary_id, project_component.Pden_id, project_component.Pden_source, project_component.Pden_subtype, project_component.Physical_item_id, project_component.Pool_id, project_component.Ppdm_guid, project_component.Ppdm_system_id, project_component.Ppdm_table_name, project_component.Product_type, project_component.Prod_string_id, project_component.Prod_string_source, project_component.Project_plan_id, project_component.Project_step_id, project_component.Provision_id, project_component.Rate_schedule_id, project_component.Referenced_guid, project_component.Referenced_system_id, project_component.Referenced_table_name, project_component.Remark, project_component.Resent_id, project_component.Reserve_class_id, project_component.Reserve_product_type, project_component.Reserve_revision_id, project_component.Reserve_revision_obs_no, project_component.Reserve_summary_id, project_component.Reserve_summary_obs_no, project_component.Sample_anal_source, project_component.Seis_license_id, project_component.Seis_proc_plan_id, project_component.Seis_set_id, project_component.Seis_set_subtype, project_component.Sf_subtype, project_component.Source, project_component.Spatial_description_id, project_component.Spatial_obs_no, project_component.Store_id, project_component.Store_structure_obs_no, project_component.Strat_column_id, project_component.Strat_column_source, project_component.Strat_name_set_id, project_component.Strat_unit_id, project_component.Support_facility_id, project_component.Sw_application_id, project_component.Thesaurus_id, project_component.Thesaurus_word_id, project_component.Uwi, project_component.Well_activity_set_id, project_component.Well_activity_source, project_component.Well_activity_type_id, project_component.Well_int_interp_id, project_component.Well_int_source, project_component.Well_int_strat_unit_id, project_component.Well_int_uwi, project_component.Well_license_id, project_component.Well_license_source, project_component.Well_set_id, project_component.Work_order_id, project_component.Row_changed_by, project_component.Row_changed_date, project_component.Row_created_by, project_component.Row_created_date, project_component.Row_effective_date, project_component.Row_expiry_date, project_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateProjectComponent(c *fiber.Ctx) error {
	var project_component dto.Project_component

	setDefaults(&project_component)

	if err := json.Unmarshal(c.Body(), &project_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	project_component.Project_id = id

    if project_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update project_component set component_obs_no = :1, active_ind = :2, activity_obs_no = :3, analysis_id = :4, application_id = :5, area_id = :6, area_type = :7, authority_id = :8, ba_licensee_ba_id = :9, ba_license_id = :10, ba_organization_id = :11, ba_organization_seq_no = :12, business_associate_id = :13, catalogue_additive_id = :14, catalogue_equip_id = :15, classification_system_id = :16, class_level_id = :17, component_reason = :18, component_type = :19, condition_id = :20, consent_id = :21, consult_id = :22, contest_id = :23, contract_id = :24, description = :25, ecozone_id = :26, effective_date = :27, employee_ba_id = :28, employee_obs_no = :29, employer_ba_id = :30, entitlement_id = :31, equipment_id = :32, expiry_date = :33, facility_id = :34, facility_license_id = :35, facility_type = :36, field_id = :37, field_source = :38, field_station_id = :39, field_strat_interp_id = :40, field_strat_unit_id = :41, finance_id = :42, fossil_id = :43, hse_incident_id = :44, incident_set_id = :45, incident_type = :46, information_item_id = :47, info_item_subtype = :48, input_ind = :49, instrument_id = :50, interest_set_id = :51, interest_set_seq_no = :52, jurisdiction = :53, land_right_id = :54, land_right_subtype = :55, land_sale_number = :56, language = :57, lithology_log_id = :58, lith_log_source = :59, lith_sample_id = :60, notification_id = :61, obligation_id = :62, obligation_seq_no = :63, other_project_id = :64, output_ind = :65, paleo_summary_id = :66, pden_id = :67, pden_source = :68, pden_subtype = :69, physical_item_id = :70, pool_id = :71, ppdm_guid = :72, ppdm_system_id = :73, ppdm_table_name = :74, product_type = :75, prod_string_id = :76, prod_string_source = :77, project_plan_id = :78, project_step_id = :79, provision_id = :80, rate_schedule_id = :81, referenced_guid = :82, referenced_system_id = :83, referenced_table_name = :84, remark = :85, resent_id = :86, reserve_class_id = :87, reserve_product_type = :88, reserve_revision_id = :89, reserve_revision_obs_no = :90, reserve_summary_id = :91, reserve_summary_obs_no = :92, sample_anal_source = :93, seis_license_id = :94, seis_proc_plan_id = :95, seis_set_id = :96, seis_set_subtype = :97, sf_subtype = :98, source = :99, spatial_description_id = :100, spatial_obs_no = :101, store_id = :102, store_structure_obs_no = :103, strat_column_id = :104, strat_column_source = :105, strat_name_set_id = :106, strat_unit_id = :107, support_facility_id = :108, sw_application_id = :109, thesaurus_id = :110, thesaurus_word_id = :111, uwi = :112, well_activity_set_id = :113, well_activity_source = :114, well_activity_type_id = :115, well_int_interp_id = :116, well_int_source = :117, well_int_strat_unit_id = :118, well_int_uwi = :119, well_license_id = :120, well_license_source = :121, well_set_id = :122, work_order_id = :123, row_changed_by = :124, row_changed_date = :125, row_created_by = :126, row_effective_date = :127, row_expiry_date = :128, row_quality = :129 where project_id = :131")
	if err != nil {
		return err
	}

	project_component.Row_changed_date = formatDate(project_component.Row_changed_date)
	project_component.Effective_date = formatDateString(project_component.Effective_date)
	project_component.Expiry_date = formatDateString(project_component.Expiry_date)
	project_component.Row_effective_date = formatDateString(project_component.Row_effective_date)
	project_component.Row_expiry_date = formatDateString(project_component.Row_expiry_date)






	rows, err := stmt.Exec(project_component.Component_obs_no, project_component.Active_ind, project_component.Activity_obs_no, project_component.Analysis_id, project_component.Application_id, project_component.Area_id, project_component.Area_type, project_component.Authority_id, project_component.Ba_licensee_ba_id, project_component.Ba_license_id, project_component.Ba_organization_id, project_component.Ba_organization_seq_no, project_component.Business_associate_id, project_component.Catalogue_additive_id, project_component.Catalogue_equip_id, project_component.Classification_system_id, project_component.Class_level_id, project_component.Component_reason, project_component.Component_type, project_component.Condition_id, project_component.Consent_id, project_component.Consult_id, project_component.Contest_id, project_component.Contract_id, project_component.Description, project_component.Ecozone_id, project_component.Effective_date, project_component.Employee_ba_id, project_component.Employee_obs_no, project_component.Employer_ba_id, project_component.Entitlement_id, project_component.Equipment_id, project_component.Expiry_date, project_component.Facility_id, project_component.Facility_license_id, project_component.Facility_type, project_component.Field_id, project_component.Field_source, project_component.Field_station_id, project_component.Field_strat_interp_id, project_component.Field_strat_unit_id, project_component.Finance_id, project_component.Fossil_id, project_component.Hse_incident_id, project_component.Incident_set_id, project_component.Incident_type, project_component.Information_item_id, project_component.Info_item_subtype, project_component.Input_ind, project_component.Instrument_id, project_component.Interest_set_id, project_component.Interest_set_seq_no, project_component.Jurisdiction, project_component.Land_right_id, project_component.Land_right_subtype, project_component.Land_sale_number, project_component.Language, project_component.Lithology_log_id, project_component.Lith_log_source, project_component.Lith_sample_id, project_component.Notification_id, project_component.Obligation_id, project_component.Obligation_seq_no, project_component.Other_project_id, project_component.Output_ind, project_component.Paleo_summary_id, project_component.Pden_id, project_component.Pden_source, project_component.Pden_subtype, project_component.Physical_item_id, project_component.Pool_id, project_component.Ppdm_guid, project_component.Ppdm_system_id, project_component.Ppdm_table_name, project_component.Product_type, project_component.Prod_string_id, project_component.Prod_string_source, project_component.Project_plan_id, project_component.Project_step_id, project_component.Provision_id, project_component.Rate_schedule_id, project_component.Referenced_guid, project_component.Referenced_system_id, project_component.Referenced_table_name, project_component.Remark, project_component.Resent_id, project_component.Reserve_class_id, project_component.Reserve_product_type, project_component.Reserve_revision_id, project_component.Reserve_revision_obs_no, project_component.Reserve_summary_id, project_component.Reserve_summary_obs_no, project_component.Sample_anal_source, project_component.Seis_license_id, project_component.Seis_proc_plan_id, project_component.Seis_set_id, project_component.Seis_set_subtype, project_component.Sf_subtype, project_component.Source, project_component.Spatial_description_id, project_component.Spatial_obs_no, project_component.Store_id, project_component.Store_structure_obs_no, project_component.Strat_column_id, project_component.Strat_column_source, project_component.Strat_name_set_id, project_component.Strat_unit_id, project_component.Support_facility_id, project_component.Sw_application_id, project_component.Thesaurus_id, project_component.Thesaurus_word_id, project_component.Uwi, project_component.Well_activity_set_id, project_component.Well_activity_source, project_component.Well_activity_type_id, project_component.Well_int_interp_id, project_component.Well_int_source, project_component.Well_int_strat_unit_id, project_component.Well_int_uwi, project_component.Well_license_id, project_component.Well_license_source, project_component.Well_set_id, project_component.Work_order_id, project_component.Row_changed_by, project_component.Row_changed_date, project_component.Row_created_by, project_component.Row_effective_date, project_component.Row_expiry_date, project_component.Row_quality, project_component.Project_id)
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

func PatchProjectComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update project_component set "
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
	query += " where project_id = :id"

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

func DeleteProjectComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var project_component dto.Project_component
	project_component.Project_id = id

	stmt, err := db.Prepare("delete from project_component where project_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(project_component.Project_id)
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


