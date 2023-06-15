package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetClassLevelComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from class_level_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Class_level_component

	for rows.Next() {
		var class_level_component dto.Class_level_component
		if err := rows.Scan(&class_level_component.Classification_system_id, &class_level_component.Class_level_id, &class_level_component.Component_obs_no, &class_level_component.Active_ind, &class_level_component.Activity_obs_no, &class_level_component.Analysis_id, &class_level_component.Application_id, &class_level_component.Area_id, &class_level_component.Area_type, &class_level_component.Authority_id, &class_level_component.Ba_organization_id, &class_level_component.Ba_organization_seq_no, &class_level_component.Business_associate_id, &class_level_component.Catalogue_additive_id, &class_level_component.Catalogue_equip_id, &class_level_component.Class_level_component_type, &class_level_component.Consent_id, &class_level_component.Consult_id, &class_level_component.Contest_id, &class_level_component.Contract_id, &class_level_component.Ecozone_id, &class_level_component.Effective_date, &class_level_component.Employee_ba_id, &class_level_component.Employee_obs_no, &class_level_component.Employer_ba_id, &class_level_component.Entitlement_id, &class_level_component.Equipment_id, &class_level_component.Expiry_date, &class_level_component.Facility_id, &class_level_component.Facility_type, &class_level_component.Field_id, &class_level_component.Field_station_id, &class_level_component.Finance_id, &class_level_component.Fossil_id, &class_level_component.Incident_id, &class_level_component.Incident_set_id, &class_level_component.Incident_type, &class_level_component.Information_item_id, &class_level_component.Info_item_subtype, &class_level_component.Instrument_id, &class_level_component.Interest_set_id, &class_level_component.Interest_set_seq_no, &class_level_component.Jurisdiction, &class_level_component.Land_right_id, &class_level_component.Land_right_subtype, &class_level_component.Land_sale_number, &class_level_component.Language, &class_level_component.Lithology_log_id, &class_level_component.Lith_log_source, &class_level_component.Notification_id, &class_level_component.Obligation_id, &class_level_component.Obligation_seq_no, &class_level_component.Paleo_summary_id, &class_level_component.Pden_id, &class_level_component.Pden_source, &class_level_component.Pden_subtype, &class_level_component.Physical_item_id, &class_level_component.Pool_id, &class_level_component.Ppdm_guid, &class_level_component.Product_type, &class_level_component.Project_id, &class_level_component.Pr_str_source, &class_level_component.Pr_str_uwi, &class_level_component.Rate_schedule_id, &class_level_component.Referenced_guid, &class_level_component.Referenced_system_id, &class_level_component.Referenced_table_name, &class_level_component.Remark, &class_level_component.Resent_id, &class_level_component.Reserve_class_id, &class_level_component.Sample_anal_source, &class_level_component.Seis_set_id, &class_level_component.Seis_set_subtype, &class_level_component.Sf_subtype, &class_level_component.Source, &class_level_component.Spatial_description_id, &class_level_component.Spatial_obs_no, &class_level_component.Store_id, &class_level_component.Store_structure_obs_no, &class_level_component.Strat_name_set_id, &class_level_component.Strat_unit_id, &class_level_component.String_id, &class_level_component.Support_facility_id, &class_level_component.Sw_application_id, &class_level_component.Thesaurus_id, &class_level_component.Thesaurus_word_id, &class_level_component.Uwi, &class_level_component.Well_activity_set_id, &class_level_component.Well_activity_source, &class_level_component.Well_activity_type_id, &class_level_component.Well_activity_uwi, &class_level_component.Well_set_id, &class_level_component.Work_order_id, &class_level_component.Row_changed_by, &class_level_component.Row_changed_date, &class_level_component.Row_created_by, &class_level_component.Row_created_date, &class_level_component.Row_effective_date, &class_level_component.Row_expiry_date, &class_level_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append class_level_component to result
		result = append(result, class_level_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetClassLevelComponent(c *fiber.Ctx) error {
	var class_level_component dto.Class_level_component

	setDefaults(&class_level_component)

	if err := json.Unmarshal(c.Body(), &class_level_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into class_level_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100)")
	if err != nil {
		return err
	}
	class_level_component.Row_created_date = formatDate(class_level_component.Row_created_date)
	class_level_component.Row_changed_date = formatDate(class_level_component.Row_changed_date)
	class_level_component.Effective_date = formatDateString(class_level_component.Effective_date)
	class_level_component.Expiry_date = formatDateString(class_level_component.Expiry_date)
	class_level_component.Row_effective_date = formatDateString(class_level_component.Row_effective_date)
	class_level_component.Row_expiry_date = formatDateString(class_level_component.Row_expiry_date)






	rows, err := stmt.Exec(class_level_component.Classification_system_id, class_level_component.Class_level_id, class_level_component.Component_obs_no, class_level_component.Active_ind, class_level_component.Activity_obs_no, class_level_component.Analysis_id, class_level_component.Application_id, class_level_component.Area_id, class_level_component.Area_type, class_level_component.Authority_id, class_level_component.Ba_organization_id, class_level_component.Ba_organization_seq_no, class_level_component.Business_associate_id, class_level_component.Catalogue_additive_id, class_level_component.Catalogue_equip_id, class_level_component.Class_level_component_type, class_level_component.Consent_id, class_level_component.Consult_id, class_level_component.Contest_id, class_level_component.Contract_id, class_level_component.Ecozone_id, class_level_component.Effective_date, class_level_component.Employee_ba_id, class_level_component.Employee_obs_no, class_level_component.Employer_ba_id, class_level_component.Entitlement_id, class_level_component.Equipment_id, class_level_component.Expiry_date, class_level_component.Facility_id, class_level_component.Facility_type, class_level_component.Field_id, class_level_component.Field_station_id, class_level_component.Finance_id, class_level_component.Fossil_id, class_level_component.Incident_id, class_level_component.Incident_set_id, class_level_component.Incident_type, class_level_component.Information_item_id, class_level_component.Info_item_subtype, class_level_component.Instrument_id, class_level_component.Interest_set_id, class_level_component.Interest_set_seq_no, class_level_component.Jurisdiction, class_level_component.Land_right_id, class_level_component.Land_right_subtype, class_level_component.Land_sale_number, class_level_component.Language, class_level_component.Lithology_log_id, class_level_component.Lith_log_source, class_level_component.Notification_id, class_level_component.Obligation_id, class_level_component.Obligation_seq_no, class_level_component.Paleo_summary_id, class_level_component.Pden_id, class_level_component.Pden_source, class_level_component.Pden_subtype, class_level_component.Physical_item_id, class_level_component.Pool_id, class_level_component.Ppdm_guid, class_level_component.Product_type, class_level_component.Project_id, class_level_component.Pr_str_source, class_level_component.Pr_str_uwi, class_level_component.Rate_schedule_id, class_level_component.Referenced_guid, class_level_component.Referenced_system_id, class_level_component.Referenced_table_name, class_level_component.Remark, class_level_component.Resent_id, class_level_component.Reserve_class_id, class_level_component.Sample_anal_source, class_level_component.Seis_set_id, class_level_component.Seis_set_subtype, class_level_component.Sf_subtype, class_level_component.Source, class_level_component.Spatial_description_id, class_level_component.Spatial_obs_no, class_level_component.Store_id, class_level_component.Store_structure_obs_no, class_level_component.Strat_name_set_id, class_level_component.Strat_unit_id, class_level_component.String_id, class_level_component.Support_facility_id, class_level_component.Sw_application_id, class_level_component.Thesaurus_id, class_level_component.Thesaurus_word_id, class_level_component.Uwi, class_level_component.Well_activity_set_id, class_level_component.Well_activity_source, class_level_component.Well_activity_type_id, class_level_component.Well_activity_uwi, class_level_component.Well_set_id, class_level_component.Work_order_id, class_level_component.Row_changed_by, class_level_component.Row_changed_date, class_level_component.Row_created_by, class_level_component.Row_created_date, class_level_component.Row_effective_date, class_level_component.Row_expiry_date, class_level_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateClassLevelComponent(c *fiber.Ctx) error {
	var class_level_component dto.Class_level_component

	setDefaults(&class_level_component)

	if err := json.Unmarshal(c.Body(), &class_level_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	class_level_component.Classification_system_id = id

    if class_level_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update class_level_component set class_level_id = :1, component_obs_no = :2, active_ind = :3, activity_obs_no = :4, analysis_id = :5, application_id = :6, area_id = :7, area_type = :8, authority_id = :9, ba_organization_id = :10, ba_organization_seq_no = :11, business_associate_id = :12, catalogue_additive_id = :13, catalogue_equip_id = :14, class_level_component_type = :15, consent_id = :16, consult_id = :17, contest_id = :18, contract_id = :19, ecozone_id = :20, effective_date = :21, employee_ba_id = :22, employee_obs_no = :23, employer_ba_id = :24, entitlement_id = :25, equipment_id = :26, expiry_date = :27, facility_id = :28, facility_type = :29, field_id = :30, field_station_id = :31, finance_id = :32, fossil_id = :33, incident_id = :34, incident_set_id = :35, incident_type = :36, information_item_id = :37, info_item_subtype = :38, instrument_id = :39, interest_set_id = :40, interest_set_seq_no = :41, jurisdiction = :42, land_right_id = :43, land_right_subtype = :44, land_sale_number = :45, language = :46, lithology_log_id = :47, lith_log_source = :48, notification_id = :49, obligation_id = :50, obligation_seq_no = :51, paleo_summary_id = :52, pden_id = :53, pden_source = :54, pden_subtype = :55, physical_item_id = :56, pool_id = :57, ppdm_guid = :58, product_type = :59, project_id = :60, pr_str_source = :61, pr_str_uwi = :62, rate_schedule_id = :63, referenced_guid = :64, referenced_system_id = :65, referenced_table_name = :66, remark = :67, resent_id = :68, reserve_class_id = :69, sample_anal_source = :70, seis_set_id = :71, seis_set_subtype = :72, sf_subtype = :73, source = :74, spatial_description_id = :75, spatial_obs_no = :76, store_id = :77, store_structure_obs_no = :78, strat_name_set_id = :79, strat_unit_id = :80, string_id = :81, support_facility_id = :82, sw_application_id = :83, thesaurus_id = :84, thesaurus_word_id = :85, uwi = :86, well_activity_set_id = :87, well_activity_source = :88, well_activity_type_id = :89, well_activity_uwi = :90, well_set_id = :91, work_order_id = :92, row_changed_by = :93, row_changed_date = :94, row_created_by = :95, row_effective_date = :96, row_expiry_date = :97, row_quality = :98 where classification_system_id = :100")
	if err != nil {
		return err
	}

	class_level_component.Row_changed_date = formatDate(class_level_component.Row_changed_date)
	class_level_component.Effective_date = formatDateString(class_level_component.Effective_date)
	class_level_component.Expiry_date = formatDateString(class_level_component.Expiry_date)
	class_level_component.Row_effective_date = formatDateString(class_level_component.Row_effective_date)
	class_level_component.Row_expiry_date = formatDateString(class_level_component.Row_expiry_date)






	rows, err := stmt.Exec(class_level_component.Class_level_id, class_level_component.Component_obs_no, class_level_component.Active_ind, class_level_component.Activity_obs_no, class_level_component.Analysis_id, class_level_component.Application_id, class_level_component.Area_id, class_level_component.Area_type, class_level_component.Authority_id, class_level_component.Ba_organization_id, class_level_component.Ba_organization_seq_no, class_level_component.Business_associate_id, class_level_component.Catalogue_additive_id, class_level_component.Catalogue_equip_id, class_level_component.Class_level_component_type, class_level_component.Consent_id, class_level_component.Consult_id, class_level_component.Contest_id, class_level_component.Contract_id, class_level_component.Ecozone_id, class_level_component.Effective_date, class_level_component.Employee_ba_id, class_level_component.Employee_obs_no, class_level_component.Employer_ba_id, class_level_component.Entitlement_id, class_level_component.Equipment_id, class_level_component.Expiry_date, class_level_component.Facility_id, class_level_component.Facility_type, class_level_component.Field_id, class_level_component.Field_station_id, class_level_component.Finance_id, class_level_component.Fossil_id, class_level_component.Incident_id, class_level_component.Incident_set_id, class_level_component.Incident_type, class_level_component.Information_item_id, class_level_component.Info_item_subtype, class_level_component.Instrument_id, class_level_component.Interest_set_id, class_level_component.Interest_set_seq_no, class_level_component.Jurisdiction, class_level_component.Land_right_id, class_level_component.Land_right_subtype, class_level_component.Land_sale_number, class_level_component.Language, class_level_component.Lithology_log_id, class_level_component.Lith_log_source, class_level_component.Notification_id, class_level_component.Obligation_id, class_level_component.Obligation_seq_no, class_level_component.Paleo_summary_id, class_level_component.Pden_id, class_level_component.Pden_source, class_level_component.Pden_subtype, class_level_component.Physical_item_id, class_level_component.Pool_id, class_level_component.Ppdm_guid, class_level_component.Product_type, class_level_component.Project_id, class_level_component.Pr_str_source, class_level_component.Pr_str_uwi, class_level_component.Rate_schedule_id, class_level_component.Referenced_guid, class_level_component.Referenced_system_id, class_level_component.Referenced_table_name, class_level_component.Remark, class_level_component.Resent_id, class_level_component.Reserve_class_id, class_level_component.Sample_anal_source, class_level_component.Seis_set_id, class_level_component.Seis_set_subtype, class_level_component.Sf_subtype, class_level_component.Source, class_level_component.Spatial_description_id, class_level_component.Spatial_obs_no, class_level_component.Store_id, class_level_component.Store_structure_obs_no, class_level_component.Strat_name_set_id, class_level_component.Strat_unit_id, class_level_component.String_id, class_level_component.Support_facility_id, class_level_component.Sw_application_id, class_level_component.Thesaurus_id, class_level_component.Thesaurus_word_id, class_level_component.Uwi, class_level_component.Well_activity_set_id, class_level_component.Well_activity_source, class_level_component.Well_activity_type_id, class_level_component.Well_activity_uwi, class_level_component.Well_set_id, class_level_component.Work_order_id, class_level_component.Row_changed_by, class_level_component.Row_changed_date, class_level_component.Row_created_by, class_level_component.Row_effective_date, class_level_component.Row_expiry_date, class_level_component.Row_quality, class_level_component.Classification_system_id)
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

func PatchClassLevelComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update class_level_component set "
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
	query += " where classification_system_id = :id"

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

func DeleteClassLevelComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var class_level_component dto.Class_level_component
	class_level_component.Classification_system_id = id

	stmt, err := db.Prepare("delete from class_level_component where classification_system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(class_level_component.Classification_system_id)
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


