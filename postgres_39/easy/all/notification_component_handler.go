package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetNotificationComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from notification_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Notification_component

	for rows.Next() {
		var notification_component dto.Notification_component
		if err := rows.Scan(&notification_component.Notification_id, &notification_component.Component_obs_no, &notification_component.Active_ind, &notification_component.Activity_obs_no, &notification_component.Analysis_id, &notification_component.Application_id, &notification_component.Area_id, &notification_component.Area_type, &notification_component.Authority_id, &notification_component.Ba_organization_id, &notification_component.Ba_organization_seq_no, &notification_component.Business_associate_id, &notification_component.Catalogue_additive_id, &notification_component.Catalogue_equip_id, &notification_component.Classification_system_id, &notification_component.Class_level_id, &notification_component.Consent_id, &notification_component.Consult_id, &notification_component.Contest_id, &notification_component.Contract_id, &notification_component.Ecozone_id, &notification_component.Effective_date, &notification_component.Employee_ba_id, &notification_component.Employee_obs_no, &notification_component.Employer_ba_id, &notification_component.Entitlement_id, &notification_component.Equipment_id, &notification_component.Expiry_date, &notification_component.Facility_id, &notification_component.Facility_type, &notification_component.Field_id, &notification_component.Field_station_id, &notification_component.Finance_id, &notification_component.Fossil_id, &notification_component.Incident_id, &notification_component.Incident_set_id, &notification_component.Incident_type, &notification_component.Information_item_id, &notification_component.Info_item_subtype, &notification_component.Instrument_id, &notification_component.Interest_set_id, &notification_component.Interest_set_seq_no, &notification_component.Jurisdiction, &notification_component.Land_right_id, &notification_component.Land_right_subtype, &notification_component.Land_sale_number, &notification_component.Language, &notification_component.Lithology_log_id, &notification_component.Lith_log_source, &notification_component.Notification_component_type, &notification_component.Obligation_id, &notification_component.Obligation_seq_no, &notification_component.Paleo_summary_id, &notification_component.Pden_id, &notification_component.Pden_source, &notification_component.Pden_subtype, &notification_component.Physical_item_id, &notification_component.Pool_id, &notification_component.Ppdm_guid, &notification_component.Ppdm_system_id, &notification_component.Ppdm_table_name, &notification_component.Product_type, &notification_component.Project_id, &notification_component.Pr_str_source, &notification_component.Pr_str_uwi, &notification_component.Rate_schedule_id, &notification_component.Referenced_guid, &notification_component.Referenced_system_id, &notification_component.Referenced_table_name, &notification_component.Remark, &notification_component.Resent_id, &notification_component.Reserve_class_id, &notification_component.Sample_anal_source, &notification_component.Seis_set_id, &notification_component.Seis_set_subtype, &notification_component.Sf_subtype, &notification_component.Source, &notification_component.Spatial_description_id, &notification_component.Spatial_obs_no, &notification_component.Store_id, &notification_component.Store_structure_obs_no, &notification_component.Strat_name_set_id, &notification_component.Strat_unit_id, &notification_component.String_id, &notification_component.Support_facility_id, &notification_component.Sw_application_id, &notification_component.Thesaurus_id, &notification_component.Thesaurus_word_id, &notification_component.Uwi, &notification_component.Well_activity_set_id, &notification_component.Well_activity_source, &notification_component.Well_activity_type_id, &notification_component.Well_activity_uwi, &notification_component.Well_set_id, &notification_component.Work_order_id, &notification_component.Row_changed_by, &notification_component.Row_changed_date, &notification_component.Row_created_by, &notification_component.Row_created_date, &notification_component.Row_effective_date, &notification_component.Row_expiry_date, &notification_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append notification_component to result
		result = append(result, notification_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetNotificationComponent(c *fiber.Ctx) error {
	var notification_component dto.Notification_component

	setDefaults(&notification_component)

	if err := json.Unmarshal(c.Body(), &notification_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into notification_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102)")
	if err != nil {
		return err
	}
	notification_component.Row_created_date = formatDate(notification_component.Row_created_date)
	notification_component.Row_changed_date = formatDate(notification_component.Row_changed_date)
	notification_component.Effective_date = formatDateString(notification_component.Effective_date)
	notification_component.Expiry_date = formatDateString(notification_component.Expiry_date)
	notification_component.Row_effective_date = formatDateString(notification_component.Row_effective_date)
	notification_component.Row_expiry_date = formatDateString(notification_component.Row_expiry_date)






	rows, err := stmt.Exec(notification_component.Notification_id, notification_component.Component_obs_no, notification_component.Active_ind, notification_component.Activity_obs_no, notification_component.Analysis_id, notification_component.Application_id, notification_component.Area_id, notification_component.Area_type, notification_component.Authority_id, notification_component.Ba_organization_id, notification_component.Ba_organization_seq_no, notification_component.Business_associate_id, notification_component.Catalogue_additive_id, notification_component.Catalogue_equip_id, notification_component.Classification_system_id, notification_component.Class_level_id, notification_component.Consent_id, notification_component.Consult_id, notification_component.Contest_id, notification_component.Contract_id, notification_component.Ecozone_id, notification_component.Effective_date, notification_component.Employee_ba_id, notification_component.Employee_obs_no, notification_component.Employer_ba_id, notification_component.Entitlement_id, notification_component.Equipment_id, notification_component.Expiry_date, notification_component.Facility_id, notification_component.Facility_type, notification_component.Field_id, notification_component.Field_station_id, notification_component.Finance_id, notification_component.Fossil_id, notification_component.Incident_id, notification_component.Incident_set_id, notification_component.Incident_type, notification_component.Information_item_id, notification_component.Info_item_subtype, notification_component.Instrument_id, notification_component.Interest_set_id, notification_component.Interest_set_seq_no, notification_component.Jurisdiction, notification_component.Land_right_id, notification_component.Land_right_subtype, notification_component.Land_sale_number, notification_component.Language, notification_component.Lithology_log_id, notification_component.Lith_log_source, notification_component.Notification_component_type, notification_component.Obligation_id, notification_component.Obligation_seq_no, notification_component.Paleo_summary_id, notification_component.Pden_id, notification_component.Pden_source, notification_component.Pden_subtype, notification_component.Physical_item_id, notification_component.Pool_id, notification_component.Ppdm_guid, notification_component.Ppdm_system_id, notification_component.Ppdm_table_name, notification_component.Product_type, notification_component.Project_id, notification_component.Pr_str_source, notification_component.Pr_str_uwi, notification_component.Rate_schedule_id, notification_component.Referenced_guid, notification_component.Referenced_system_id, notification_component.Referenced_table_name, notification_component.Remark, notification_component.Resent_id, notification_component.Reserve_class_id, notification_component.Sample_anal_source, notification_component.Seis_set_id, notification_component.Seis_set_subtype, notification_component.Sf_subtype, notification_component.Source, notification_component.Spatial_description_id, notification_component.Spatial_obs_no, notification_component.Store_id, notification_component.Store_structure_obs_no, notification_component.Strat_name_set_id, notification_component.Strat_unit_id, notification_component.String_id, notification_component.Support_facility_id, notification_component.Sw_application_id, notification_component.Thesaurus_id, notification_component.Thesaurus_word_id, notification_component.Uwi, notification_component.Well_activity_set_id, notification_component.Well_activity_source, notification_component.Well_activity_type_id, notification_component.Well_activity_uwi, notification_component.Well_set_id, notification_component.Work_order_id, notification_component.Row_changed_by, notification_component.Row_changed_date, notification_component.Row_created_by, notification_component.Row_created_date, notification_component.Row_effective_date, notification_component.Row_expiry_date, notification_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateNotificationComponent(c *fiber.Ctx) error {
	var notification_component dto.Notification_component

	setDefaults(&notification_component)

	if err := json.Unmarshal(c.Body(), &notification_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	notification_component.Notification_id = id

    if notification_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update notification_component set component_obs_no = :1, active_ind = :2, activity_obs_no = :3, analysis_id = :4, application_id = :5, area_id = :6, area_type = :7, authority_id = :8, ba_organization_id = :9, ba_organization_seq_no = :10, business_associate_id = :11, catalogue_additive_id = :12, catalogue_equip_id = :13, classification_system_id = :14, class_level_id = :15, consent_id = :16, consult_id = :17, contest_id = :18, contract_id = :19, ecozone_id = :20, effective_date = :21, employee_ba_id = :22, employee_obs_no = :23, employer_ba_id = :24, entitlement_id = :25, equipment_id = :26, expiry_date = :27, facility_id = :28, facility_type = :29, field_id = :30, field_station_id = :31, finance_id = :32, fossil_id = :33, incident_id = :34, incident_set_id = :35, incident_type = :36, information_item_id = :37, info_item_subtype = :38, instrument_id = :39, interest_set_id = :40, interest_set_seq_no = :41, jurisdiction = :42, land_right_id = :43, land_right_subtype = :44, land_sale_number = :45, language = :46, lithology_log_id = :47, lith_log_source = :48, notification_component_type = :49, obligation_id = :50, obligation_seq_no = :51, paleo_summary_id = :52, pden_id = :53, pden_source = :54, pden_subtype = :55, physical_item_id = :56, pool_id = :57, ppdm_guid = :58, ppdm_system_id = :59, ppdm_table_name = :60, product_type = :61, project_id = :62, pr_str_source = :63, pr_str_uwi = :64, rate_schedule_id = :65, referenced_guid = :66, referenced_system_id = :67, referenced_table_name = :68, remark = :69, resent_id = :70, reserve_class_id = :71, sample_anal_source = :72, seis_set_id = :73, seis_set_subtype = :74, sf_subtype = :75, source = :76, spatial_description_id = :77, spatial_obs_no = :78, store_id = :79, store_structure_obs_no = :80, strat_name_set_id = :81, strat_unit_id = :82, string_id = :83, support_facility_id = :84, sw_application_id = :85, thesaurus_id = :86, thesaurus_word_id = :87, uwi = :88, well_activity_set_id = :89, well_activity_source = :90, well_activity_type_id = :91, well_activity_uwi = :92, well_set_id = :93, work_order_id = :94, row_changed_by = :95, row_changed_date = :96, row_created_by = :97, row_effective_date = :98, row_expiry_date = :99, row_quality = :100 where notification_id = :102")
	if err != nil {
		return err
	}

	notification_component.Row_changed_date = formatDate(notification_component.Row_changed_date)
	notification_component.Effective_date = formatDateString(notification_component.Effective_date)
	notification_component.Expiry_date = formatDateString(notification_component.Expiry_date)
	notification_component.Row_effective_date = formatDateString(notification_component.Row_effective_date)
	notification_component.Row_expiry_date = formatDateString(notification_component.Row_expiry_date)






	rows, err := stmt.Exec(notification_component.Component_obs_no, notification_component.Active_ind, notification_component.Activity_obs_no, notification_component.Analysis_id, notification_component.Application_id, notification_component.Area_id, notification_component.Area_type, notification_component.Authority_id, notification_component.Ba_organization_id, notification_component.Ba_organization_seq_no, notification_component.Business_associate_id, notification_component.Catalogue_additive_id, notification_component.Catalogue_equip_id, notification_component.Classification_system_id, notification_component.Class_level_id, notification_component.Consent_id, notification_component.Consult_id, notification_component.Contest_id, notification_component.Contract_id, notification_component.Ecozone_id, notification_component.Effective_date, notification_component.Employee_ba_id, notification_component.Employee_obs_no, notification_component.Employer_ba_id, notification_component.Entitlement_id, notification_component.Equipment_id, notification_component.Expiry_date, notification_component.Facility_id, notification_component.Facility_type, notification_component.Field_id, notification_component.Field_station_id, notification_component.Finance_id, notification_component.Fossil_id, notification_component.Incident_id, notification_component.Incident_set_id, notification_component.Incident_type, notification_component.Information_item_id, notification_component.Info_item_subtype, notification_component.Instrument_id, notification_component.Interest_set_id, notification_component.Interest_set_seq_no, notification_component.Jurisdiction, notification_component.Land_right_id, notification_component.Land_right_subtype, notification_component.Land_sale_number, notification_component.Language, notification_component.Lithology_log_id, notification_component.Lith_log_source, notification_component.Notification_component_type, notification_component.Obligation_id, notification_component.Obligation_seq_no, notification_component.Paleo_summary_id, notification_component.Pden_id, notification_component.Pden_source, notification_component.Pden_subtype, notification_component.Physical_item_id, notification_component.Pool_id, notification_component.Ppdm_guid, notification_component.Ppdm_system_id, notification_component.Ppdm_table_name, notification_component.Product_type, notification_component.Project_id, notification_component.Pr_str_source, notification_component.Pr_str_uwi, notification_component.Rate_schedule_id, notification_component.Referenced_guid, notification_component.Referenced_system_id, notification_component.Referenced_table_name, notification_component.Remark, notification_component.Resent_id, notification_component.Reserve_class_id, notification_component.Sample_anal_source, notification_component.Seis_set_id, notification_component.Seis_set_subtype, notification_component.Sf_subtype, notification_component.Source, notification_component.Spatial_description_id, notification_component.Spatial_obs_no, notification_component.Store_id, notification_component.Store_structure_obs_no, notification_component.Strat_name_set_id, notification_component.Strat_unit_id, notification_component.String_id, notification_component.Support_facility_id, notification_component.Sw_application_id, notification_component.Thesaurus_id, notification_component.Thesaurus_word_id, notification_component.Uwi, notification_component.Well_activity_set_id, notification_component.Well_activity_source, notification_component.Well_activity_type_id, notification_component.Well_activity_uwi, notification_component.Well_set_id, notification_component.Work_order_id, notification_component.Row_changed_by, notification_component.Row_changed_date, notification_component.Row_created_by, notification_component.Row_effective_date, notification_component.Row_expiry_date, notification_component.Row_quality, notification_component.Notification_id)
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

func PatchNotificationComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update notification_component set "
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
	query += " where notification_id = :id"

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

func DeleteNotificationComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var notification_component dto.Notification_component
	notification_component.Notification_id = id

	stmt, err := db.Prepare("delete from notification_component where notification_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(notification_component.Notification_id)
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


