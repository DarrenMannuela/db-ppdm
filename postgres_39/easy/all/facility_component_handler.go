package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFacilityComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from facility_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Facility_component

	for rows.Next() {
		var facility_component dto.Facility_component
		if err := rows.Scan(&facility_component.Facility_id, &facility_component.Facility_type, &facility_component.Component_obs_no, &facility_component.Active_ind, &facility_component.Activity_obs_no, &facility_component.Analysis_id, &facility_component.Application_id, &facility_component.Area_id, &facility_component.Area_type, &facility_component.Authority_id, &facility_component.Ba_organization_id, &facility_component.Ba_organization_seq_no, &facility_component.Business_associate_id, &facility_component.Catalogue_additive_id, &facility_component.Catalogue_equip_id, &facility_component.Classification_system_id, &facility_component.Class_level_id, &facility_component.Consent_id, &facility_component.Consult_id, &facility_component.Contest_id, &facility_component.Contract_id, &facility_component.Ecozone_id, &facility_component.Effective_date, &facility_component.Employee_ba_id, &facility_component.Employee_obs_no, &facility_component.Employer_ba_id, &facility_component.Entitlement_id, &facility_component.Equipment_id, &facility_component.Expiry_date, &facility_component.Facility_component_type, &facility_component.Field_id, &facility_component.Field_station_id, &facility_component.Finance_id, &facility_component.Fossil_id, &facility_component.Incident_id, &facility_component.Incident_set_id, &facility_component.Incident_type, &facility_component.Information_item_id, &facility_component.Info_item_subtype, &facility_component.Instrument_id, &facility_component.Interest_set_id, &facility_component.Interest_set_seq_no, &facility_component.Jurisdiction, &facility_component.Land_right_id, &facility_component.Land_right_subtype, &facility_component.Land_sale_number, &facility_component.Language, &facility_component.Lithology_log_id, &facility_component.Lith_log_source, &facility_component.Notification_id, &facility_component.Obligation_id, &facility_component.Obligation_seq_no, &facility_component.Paleo_summary_id, &facility_component.Pden_id, &facility_component.Pden_source, &facility_component.Pden_subtype, &facility_component.Physical_item_id, &facility_component.Pool_id, &facility_component.Ppdm_guid, &facility_component.Ppdm_system_id, &facility_component.Ppdm_table_name, &facility_component.Product_type, &facility_component.Project_id, &facility_component.Pr_str_source, &facility_component.Pr_str_uwi, &facility_component.Rate_schedule_id, &facility_component.Referenced_guid, &facility_component.Referenced_system_id, &facility_component.Referenced_table_name, &facility_component.Remark, &facility_component.Resent_id, &facility_component.Reserve_class_id, &facility_component.Sample_anal_source, &facility_component.Seis_set_id, &facility_component.Seis_set_subtype, &facility_component.Sf_subtype, &facility_component.Source, &facility_component.Spatial_description_id, &facility_component.Spatial_obs_no, &facility_component.Store_id, &facility_component.Store_structure_obs_no, &facility_component.Strat_name_set_id, &facility_component.Strat_unit_id, &facility_component.String_id, &facility_component.Support_facility_id, &facility_component.Sw_application_id, &facility_component.Thesaurus_id, &facility_component.Thesaurus_word_id, &facility_component.Uwi, &facility_component.Well_activity_set_id, &facility_component.Well_activity_source, &facility_component.Well_activity_type_id, &facility_component.Well_activity_uwi, &facility_component.Well_set_id, &facility_component.Work_order_id, &facility_component.Row_changed_by, &facility_component.Row_changed_date, &facility_component.Row_created_by, &facility_component.Row_created_date, &facility_component.Row_effective_date, &facility_component.Row_expiry_date, &facility_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append facility_component to result
		result = append(result, facility_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFacilityComponent(c *fiber.Ctx) error {
	var facility_component dto.Facility_component

	setDefaults(&facility_component)

	if err := json.Unmarshal(c.Body(), &facility_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into facility_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102)")
	if err != nil {
		return err
	}
	facility_component.Row_created_date = formatDate(facility_component.Row_created_date)
	facility_component.Row_changed_date = formatDate(facility_component.Row_changed_date)
	facility_component.Effective_date = formatDateString(facility_component.Effective_date)
	facility_component.Expiry_date = formatDateString(facility_component.Expiry_date)
	facility_component.Row_effective_date = formatDateString(facility_component.Row_effective_date)
	facility_component.Row_expiry_date = formatDateString(facility_component.Row_expiry_date)






	rows, err := stmt.Exec(facility_component.Facility_id, facility_component.Facility_type, facility_component.Component_obs_no, facility_component.Active_ind, facility_component.Activity_obs_no, facility_component.Analysis_id, facility_component.Application_id, facility_component.Area_id, facility_component.Area_type, facility_component.Authority_id, facility_component.Ba_organization_id, facility_component.Ba_organization_seq_no, facility_component.Business_associate_id, facility_component.Catalogue_additive_id, facility_component.Catalogue_equip_id, facility_component.Classification_system_id, facility_component.Class_level_id, facility_component.Consent_id, facility_component.Consult_id, facility_component.Contest_id, facility_component.Contract_id, facility_component.Ecozone_id, facility_component.Effective_date, facility_component.Employee_ba_id, facility_component.Employee_obs_no, facility_component.Employer_ba_id, facility_component.Entitlement_id, facility_component.Equipment_id, facility_component.Expiry_date, facility_component.Facility_component_type, facility_component.Field_id, facility_component.Field_station_id, facility_component.Finance_id, facility_component.Fossil_id, facility_component.Incident_id, facility_component.Incident_set_id, facility_component.Incident_type, facility_component.Information_item_id, facility_component.Info_item_subtype, facility_component.Instrument_id, facility_component.Interest_set_id, facility_component.Interest_set_seq_no, facility_component.Jurisdiction, facility_component.Land_right_id, facility_component.Land_right_subtype, facility_component.Land_sale_number, facility_component.Language, facility_component.Lithology_log_id, facility_component.Lith_log_source, facility_component.Notification_id, facility_component.Obligation_id, facility_component.Obligation_seq_no, facility_component.Paleo_summary_id, facility_component.Pden_id, facility_component.Pden_source, facility_component.Pden_subtype, facility_component.Physical_item_id, facility_component.Pool_id, facility_component.Ppdm_guid, facility_component.Ppdm_system_id, facility_component.Ppdm_table_name, facility_component.Product_type, facility_component.Project_id, facility_component.Pr_str_source, facility_component.Pr_str_uwi, facility_component.Rate_schedule_id, facility_component.Referenced_guid, facility_component.Referenced_system_id, facility_component.Referenced_table_name, facility_component.Remark, facility_component.Resent_id, facility_component.Reserve_class_id, facility_component.Sample_anal_source, facility_component.Seis_set_id, facility_component.Seis_set_subtype, facility_component.Sf_subtype, facility_component.Source, facility_component.Spatial_description_id, facility_component.Spatial_obs_no, facility_component.Store_id, facility_component.Store_structure_obs_no, facility_component.Strat_name_set_id, facility_component.Strat_unit_id, facility_component.String_id, facility_component.Support_facility_id, facility_component.Sw_application_id, facility_component.Thesaurus_id, facility_component.Thesaurus_word_id, facility_component.Uwi, facility_component.Well_activity_set_id, facility_component.Well_activity_source, facility_component.Well_activity_type_id, facility_component.Well_activity_uwi, facility_component.Well_set_id, facility_component.Work_order_id, facility_component.Row_changed_by, facility_component.Row_changed_date, facility_component.Row_created_by, facility_component.Row_created_date, facility_component.Row_effective_date, facility_component.Row_expiry_date, facility_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFacilityComponent(c *fiber.Ctx) error {
	var facility_component dto.Facility_component

	setDefaults(&facility_component)

	if err := json.Unmarshal(c.Body(), &facility_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	facility_component.Facility_id = id

    if facility_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update facility_component set facility_type = :1, component_obs_no = :2, active_ind = :3, activity_obs_no = :4, analysis_id = :5, application_id = :6, area_id = :7, area_type = :8, authority_id = :9, ba_organization_id = :10, ba_organization_seq_no = :11, business_associate_id = :12, catalogue_additive_id = :13, catalogue_equip_id = :14, classification_system_id = :15, class_level_id = :16, consent_id = :17, consult_id = :18, contest_id = :19, contract_id = :20, ecozone_id = :21, effective_date = :22, employee_ba_id = :23, employee_obs_no = :24, employer_ba_id = :25, entitlement_id = :26, equipment_id = :27, expiry_date = :28, facility_component_type = :29, field_id = :30, field_station_id = :31, finance_id = :32, fossil_id = :33, incident_id = :34, incident_set_id = :35, incident_type = :36, information_item_id = :37, info_item_subtype = :38, instrument_id = :39, interest_set_id = :40, interest_set_seq_no = :41, jurisdiction = :42, land_right_id = :43, land_right_subtype = :44, land_sale_number = :45, language = :46, lithology_log_id = :47, lith_log_source = :48, notification_id = :49, obligation_id = :50, obligation_seq_no = :51, paleo_summary_id = :52, pden_id = :53, pden_source = :54, pden_subtype = :55, physical_item_id = :56, pool_id = :57, ppdm_guid = :58, ppdm_system_id = :59, ppdm_table_name = :60, product_type = :61, project_id = :62, pr_str_source = :63, pr_str_uwi = :64, rate_schedule_id = :65, referenced_guid = :66, referenced_system_id = :67, referenced_table_name = :68, remark = :69, resent_id = :70, reserve_class_id = :71, sample_anal_source = :72, seis_set_id = :73, seis_set_subtype = :74, sf_subtype = :75, source = :76, spatial_description_id = :77, spatial_obs_no = :78, store_id = :79, store_structure_obs_no = :80, strat_name_set_id = :81, strat_unit_id = :82, string_id = :83, support_facility_id = :84, sw_application_id = :85, thesaurus_id = :86, thesaurus_word_id = :87, uwi = :88, well_activity_set_id = :89, well_activity_source = :90, well_activity_type_id = :91, well_activity_uwi = :92, well_set_id = :93, work_order_id = :94, row_changed_by = :95, row_changed_date = :96, row_created_by = :97, row_effective_date = :98, row_expiry_date = :99, row_quality = :100 where facility_id = :102")
	if err != nil {
		return err
	}

	facility_component.Row_changed_date = formatDate(facility_component.Row_changed_date)
	facility_component.Effective_date = formatDateString(facility_component.Effective_date)
	facility_component.Expiry_date = formatDateString(facility_component.Expiry_date)
	facility_component.Row_effective_date = formatDateString(facility_component.Row_effective_date)
	facility_component.Row_expiry_date = formatDateString(facility_component.Row_expiry_date)






	rows, err := stmt.Exec(facility_component.Facility_type, facility_component.Component_obs_no, facility_component.Active_ind, facility_component.Activity_obs_no, facility_component.Analysis_id, facility_component.Application_id, facility_component.Area_id, facility_component.Area_type, facility_component.Authority_id, facility_component.Ba_organization_id, facility_component.Ba_organization_seq_no, facility_component.Business_associate_id, facility_component.Catalogue_additive_id, facility_component.Catalogue_equip_id, facility_component.Classification_system_id, facility_component.Class_level_id, facility_component.Consent_id, facility_component.Consult_id, facility_component.Contest_id, facility_component.Contract_id, facility_component.Ecozone_id, facility_component.Effective_date, facility_component.Employee_ba_id, facility_component.Employee_obs_no, facility_component.Employer_ba_id, facility_component.Entitlement_id, facility_component.Equipment_id, facility_component.Expiry_date, facility_component.Facility_component_type, facility_component.Field_id, facility_component.Field_station_id, facility_component.Finance_id, facility_component.Fossil_id, facility_component.Incident_id, facility_component.Incident_set_id, facility_component.Incident_type, facility_component.Information_item_id, facility_component.Info_item_subtype, facility_component.Instrument_id, facility_component.Interest_set_id, facility_component.Interest_set_seq_no, facility_component.Jurisdiction, facility_component.Land_right_id, facility_component.Land_right_subtype, facility_component.Land_sale_number, facility_component.Language, facility_component.Lithology_log_id, facility_component.Lith_log_source, facility_component.Notification_id, facility_component.Obligation_id, facility_component.Obligation_seq_no, facility_component.Paleo_summary_id, facility_component.Pden_id, facility_component.Pden_source, facility_component.Pden_subtype, facility_component.Physical_item_id, facility_component.Pool_id, facility_component.Ppdm_guid, facility_component.Ppdm_system_id, facility_component.Ppdm_table_name, facility_component.Product_type, facility_component.Project_id, facility_component.Pr_str_source, facility_component.Pr_str_uwi, facility_component.Rate_schedule_id, facility_component.Referenced_guid, facility_component.Referenced_system_id, facility_component.Referenced_table_name, facility_component.Remark, facility_component.Resent_id, facility_component.Reserve_class_id, facility_component.Sample_anal_source, facility_component.Seis_set_id, facility_component.Seis_set_subtype, facility_component.Sf_subtype, facility_component.Source, facility_component.Spatial_description_id, facility_component.Spatial_obs_no, facility_component.Store_id, facility_component.Store_structure_obs_no, facility_component.Strat_name_set_id, facility_component.Strat_unit_id, facility_component.String_id, facility_component.Support_facility_id, facility_component.Sw_application_id, facility_component.Thesaurus_id, facility_component.Thesaurus_word_id, facility_component.Uwi, facility_component.Well_activity_set_id, facility_component.Well_activity_source, facility_component.Well_activity_type_id, facility_component.Well_activity_uwi, facility_component.Well_set_id, facility_component.Work_order_id, facility_component.Row_changed_by, facility_component.Row_changed_date, facility_component.Row_created_by, facility_component.Row_effective_date, facility_component.Row_expiry_date, facility_component.Row_quality, facility_component.Facility_id)
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

func PatchFacilityComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update facility_component set "
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
	query += " where facility_id = :id"

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

func DeleteFacilityComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var facility_component dto.Facility_component
	facility_component.Facility_id = id

	stmt, err := db.Prepare("delete from facility_component where facility_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(facility_component.Facility_id)
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


