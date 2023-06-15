package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetApplicationComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from application_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Application_component

	for rows.Next() {
		var application_component dto.Application_component
		if err := rows.Scan(&application_component.Application_id, &application_component.Component_obs_no, &application_component.Active_ind, &application_component.Activity_obs_no, &application_component.Analysis_id, &application_component.Application_component_type, &application_component.App_application_id, &application_component.Area_id, &application_component.Area_type, &application_component.Authority_id, &application_component.Ba_organization_id, &application_component.Ba_organization_seq_no, &application_component.Business_associate_id, &application_component.Catalogue_additive_id, &application_component.Catalogue_equip_id, &application_component.Classification_system_id, &application_component.Class_level_id, &application_component.Consent_id, &application_component.Consult_id, &application_component.Contest_id, &application_component.Contract_id, &application_component.Ecozone_id, &application_component.Effective_date, &application_component.Employee_ba_id, &application_component.Employee_obs_no, &application_component.Employer_ba_id, &application_component.Entitlement_id, &application_component.Equipment_id, &application_component.Expiry_date, &application_component.Facility_id, &application_component.Facility_type, &application_component.Field_id, &application_component.Field_station_id, &application_component.Finance_id, &application_component.Fossil_id, &application_component.Incident_id, &application_component.Incident_set_id, &application_component.Incident_type, &application_component.Information_item_id, &application_component.Info_item_subtype, &application_component.Instrument_id, &application_component.Interest_set_id, &application_component.Interest_set_seq_no, &application_component.Jurisdiction, &application_component.Land_right_id, &application_component.Land_right_subtype, &application_component.Land_sale_number, &application_component.Language, &application_component.Lithology_log_id, &application_component.Lith_log_source, &application_component.Notification_id, &application_component.Obligation_id, &application_component.Obligation_seq_no, &application_component.Paleo_summary_id, &application_component.Pden_id, &application_component.Pden_source, &application_component.Pden_subtype, &application_component.Physical_item_id, &application_component.Pool_id, &application_component.Ppdm_guid, &application_component.Ppdm_system_id, &application_component.Ppdm_table_name, &application_component.Product_type, &application_component.Project_id, &application_component.Pr_str_source, &application_component.Pr_str_uwi, &application_component.Rate_schedule_id, &application_component.Referenced_guid, &application_component.Referenced_system_id, &application_component.Referenced_table_name, &application_component.Remark, &application_component.Resent_id, &application_component.Sample_anal_source, &application_component.Seis_set_id, &application_component.Seis_set_subtype, &application_component.Sf_subtype, &application_component.Source, &application_component.Spatial_description_id, &application_component.Spatial_obs_no, &application_component.Store_id, &application_component.Store_structure_obs_no, &application_component.Strat_name_set_id, &application_component.Strat_unit_id, &application_component.String_id, &application_component.Support_facility_id, &application_component.Sw_application_id, &application_component.Thesaurus_id, &application_component.Thesaurus_word_id, &application_component.Uwi, &application_component.Well_activity_set_id, &application_component.Well_activity_source, &application_component.Well_activity_type_id, &application_component.Well_activity_uwi, &application_component.Well_set_id, &application_component.Work_order_id, &application_component.Row_changed_by, &application_component.Row_changed_date, &application_component.Row_created_by, &application_component.Row_created_date, &application_component.Row_effective_date, &application_component.Row_expiry_date, &application_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append application_component to result
		result = append(result, application_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetApplicationComponent(c *fiber.Ctx) error {
	var application_component dto.Application_component

	setDefaults(&application_component)

	if err := json.Unmarshal(c.Body(), &application_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into application_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102)")
	if err != nil {
		return err
	}
	application_component.Row_created_date = formatDate(application_component.Row_created_date)
	application_component.Row_changed_date = formatDate(application_component.Row_changed_date)
	application_component.Effective_date = formatDateString(application_component.Effective_date)
	application_component.Expiry_date = formatDateString(application_component.Expiry_date)
	application_component.Row_effective_date = formatDateString(application_component.Row_effective_date)
	application_component.Row_expiry_date = formatDateString(application_component.Row_expiry_date)






	rows, err := stmt.Exec(application_component.Application_id, application_component.Component_obs_no, application_component.Active_ind, application_component.Activity_obs_no, application_component.Analysis_id, application_component.Application_component_type, application_component.App_application_id, application_component.Area_id, application_component.Area_type, application_component.Authority_id, application_component.Ba_organization_id, application_component.Ba_organization_seq_no, application_component.Business_associate_id, application_component.Catalogue_additive_id, application_component.Catalogue_equip_id, application_component.Classification_system_id, application_component.Class_level_id, application_component.Consent_id, application_component.Consult_id, application_component.Contest_id, application_component.Contract_id, application_component.Ecozone_id, application_component.Effective_date, application_component.Employee_ba_id, application_component.Employee_obs_no, application_component.Employer_ba_id, application_component.Entitlement_id, application_component.Equipment_id, application_component.Expiry_date, application_component.Facility_id, application_component.Facility_type, application_component.Field_id, application_component.Field_station_id, application_component.Finance_id, application_component.Fossil_id, application_component.Incident_id, application_component.Incident_set_id, application_component.Incident_type, application_component.Information_item_id, application_component.Info_item_subtype, application_component.Instrument_id, application_component.Interest_set_id, application_component.Interest_set_seq_no, application_component.Jurisdiction, application_component.Land_right_id, application_component.Land_right_subtype, application_component.Land_sale_number, application_component.Language, application_component.Lithology_log_id, application_component.Lith_log_source, application_component.Notification_id, application_component.Obligation_id, application_component.Obligation_seq_no, application_component.Paleo_summary_id, application_component.Pden_id, application_component.Pden_source, application_component.Pden_subtype, application_component.Physical_item_id, application_component.Pool_id, application_component.Ppdm_guid, application_component.Ppdm_system_id, application_component.Ppdm_table_name, application_component.Product_type, application_component.Project_id, application_component.Pr_str_source, application_component.Pr_str_uwi, application_component.Rate_schedule_id, application_component.Referenced_guid, application_component.Referenced_system_id, application_component.Referenced_table_name, application_component.Remark, application_component.Resent_id, application_component.Sample_anal_source, application_component.Seis_set_id, application_component.Seis_set_subtype, application_component.Sf_subtype, application_component.Source, application_component.Spatial_description_id, application_component.Spatial_obs_no, application_component.Store_id, application_component.Store_structure_obs_no, application_component.Strat_name_set_id, application_component.Strat_unit_id, application_component.String_id, application_component.Support_facility_id, application_component.Sw_application_id, application_component.Thesaurus_id, application_component.Thesaurus_word_id, application_component.Uwi, application_component.Well_activity_set_id, application_component.Well_activity_source, application_component.Well_activity_type_id, application_component.Well_activity_uwi, application_component.Well_set_id, application_component.Work_order_id, application_component.Row_changed_by, application_component.Row_changed_date, application_component.Row_created_by, application_component.Row_created_date, application_component.Row_effective_date, application_component.Row_expiry_date, application_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateApplicationComponent(c *fiber.Ctx) error {
	var application_component dto.Application_component

	setDefaults(&application_component)

	if err := json.Unmarshal(c.Body(), &application_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	application_component.Application_id = id

    if application_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update application_component set component_obs_no = :1, active_ind = :2, activity_obs_no = :3, analysis_id = :4, application_component_type = :5, app_application_id = :6, area_id = :7, area_type = :8, authority_id = :9, ba_organization_id = :10, ba_organization_seq_no = :11, business_associate_id = :12, catalogue_additive_id = :13, catalogue_equip_id = :14, classification_system_id = :15, class_level_id = :16, consent_id = :17, consult_id = :18, contest_id = :19, contract_id = :20, ecozone_id = :21, effective_date = :22, employee_ba_id = :23, employee_obs_no = :24, employer_ba_id = :25, entitlement_id = :26, equipment_id = :27, expiry_date = :28, facility_id = :29, facility_type = :30, field_id = :31, field_station_id = :32, finance_id = :33, fossil_id = :34, incident_id = :35, incident_set_id = :36, incident_type = :37, information_item_id = :38, info_item_subtype = :39, instrument_id = :40, interest_set_id = :41, interest_set_seq_no = :42, jurisdiction = :43, land_right_id = :44, land_right_subtype = :45, land_sale_number = :46, language = :47, lithology_log_id = :48, lith_log_source = :49, notification_id = :50, obligation_id = :51, obligation_seq_no = :52, paleo_summary_id = :53, pden_id = :54, pden_source = :55, pden_subtype = :56, physical_item_id = :57, pool_id = :58, ppdm_guid = :59, ppdm_system_id = :60, ppdm_table_name = :61, product_type = :62, project_id = :63, pr_str_source = :64, pr_str_uwi = :65, rate_schedule_id = :66, referenced_guid = :67, referenced_system_id = :68, referenced_table_name = :69, remark = :70, resent_id = :71, sample_anal_source = :72, seis_set_id = :73, seis_set_subtype = :74, sf_subtype = :75, source = :76, spatial_description_id = :77, spatial_obs_no = :78, store_id = :79, store_structure_obs_no = :80, strat_name_set_id = :81, strat_unit_id = :82, string_id = :83, support_facility_id = :84, sw_application_id = :85, thesaurus_id = :86, thesaurus_word_id = :87, uwi = :88, well_activity_set_id = :89, well_activity_source = :90, well_activity_type_id = :91, well_activity_uwi = :92, well_set_id = :93, work_order_id = :94, row_changed_by = :95, row_changed_date = :96, row_created_by = :97, row_effective_date = :98, row_expiry_date = :99, row_quality = :100 where application_id = :102")
	if err != nil {
		return err
	}

	application_component.Row_changed_date = formatDate(application_component.Row_changed_date)
	application_component.Effective_date = formatDateString(application_component.Effective_date)
	application_component.Expiry_date = formatDateString(application_component.Expiry_date)
	application_component.Row_effective_date = formatDateString(application_component.Row_effective_date)
	application_component.Row_expiry_date = formatDateString(application_component.Row_expiry_date)






	rows, err := stmt.Exec(application_component.Component_obs_no, application_component.Active_ind, application_component.Activity_obs_no, application_component.Analysis_id, application_component.Application_component_type, application_component.App_application_id, application_component.Area_id, application_component.Area_type, application_component.Authority_id, application_component.Ba_organization_id, application_component.Ba_organization_seq_no, application_component.Business_associate_id, application_component.Catalogue_additive_id, application_component.Catalogue_equip_id, application_component.Classification_system_id, application_component.Class_level_id, application_component.Consent_id, application_component.Consult_id, application_component.Contest_id, application_component.Contract_id, application_component.Ecozone_id, application_component.Effective_date, application_component.Employee_ba_id, application_component.Employee_obs_no, application_component.Employer_ba_id, application_component.Entitlement_id, application_component.Equipment_id, application_component.Expiry_date, application_component.Facility_id, application_component.Facility_type, application_component.Field_id, application_component.Field_station_id, application_component.Finance_id, application_component.Fossil_id, application_component.Incident_id, application_component.Incident_set_id, application_component.Incident_type, application_component.Information_item_id, application_component.Info_item_subtype, application_component.Instrument_id, application_component.Interest_set_id, application_component.Interest_set_seq_no, application_component.Jurisdiction, application_component.Land_right_id, application_component.Land_right_subtype, application_component.Land_sale_number, application_component.Language, application_component.Lithology_log_id, application_component.Lith_log_source, application_component.Notification_id, application_component.Obligation_id, application_component.Obligation_seq_no, application_component.Paleo_summary_id, application_component.Pden_id, application_component.Pden_source, application_component.Pden_subtype, application_component.Physical_item_id, application_component.Pool_id, application_component.Ppdm_guid, application_component.Ppdm_system_id, application_component.Ppdm_table_name, application_component.Product_type, application_component.Project_id, application_component.Pr_str_source, application_component.Pr_str_uwi, application_component.Rate_schedule_id, application_component.Referenced_guid, application_component.Referenced_system_id, application_component.Referenced_table_name, application_component.Remark, application_component.Resent_id, application_component.Sample_anal_source, application_component.Seis_set_id, application_component.Seis_set_subtype, application_component.Sf_subtype, application_component.Source, application_component.Spatial_description_id, application_component.Spatial_obs_no, application_component.Store_id, application_component.Store_structure_obs_no, application_component.Strat_name_set_id, application_component.Strat_unit_id, application_component.String_id, application_component.Support_facility_id, application_component.Sw_application_id, application_component.Thesaurus_id, application_component.Thesaurus_word_id, application_component.Uwi, application_component.Well_activity_set_id, application_component.Well_activity_source, application_component.Well_activity_type_id, application_component.Well_activity_uwi, application_component.Well_set_id, application_component.Work_order_id, application_component.Row_changed_by, application_component.Row_changed_date, application_component.Row_created_by, application_component.Row_effective_date, application_component.Row_expiry_date, application_component.Row_quality, application_component.Application_id)
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

func PatchApplicationComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update application_component set "
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
	query += " where application_id = :id"

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

func DeleteApplicationComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var application_component dto.Application_component
	application_component.Application_id = id

	stmt, err := db.Prepare("delete from application_component where application_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(application_component.Application_id)
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


