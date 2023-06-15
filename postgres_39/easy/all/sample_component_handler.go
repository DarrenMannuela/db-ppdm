package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSampleComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sample_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sample_component

	for rows.Next() {
		var sample_component dto.Sample_component
		if err := rows.Scan(&sample_component.Sample_id, &sample_component.Component_obs_no, &sample_component.Active_ind, &sample_component.Activity_obs_no, &sample_component.Analysis_id, &sample_component.Anl_source, &sample_component.Application_id, &sample_component.Area_id, &sample_component.Area_type, &sample_component.Authority_id, &sample_component.Ba_organization_id, &sample_component.Ba_organization_seq_no, &sample_component.Business_associate_id, &sample_component.Catalogue_additive_id, &sample_component.Catalogue_equip_id, &sample_component.Classification_system_id, &sample_component.Class_level_id, &sample_component.Component_type, &sample_component.Consent_id, &sample_component.Consult_id, &sample_component.Contest_id, &sample_component.Contract_id, &sample_component.Ecozone_id, &sample_component.Effective_date, &sample_component.Employee_ba_id, &sample_component.Employee_obs_no, &sample_component.Employer_ba_id, &sample_component.Entitlement_id, &sample_component.Equipment_id, &sample_component.Expiry_date, &sample_component.Facility_id, &sample_component.Facility_type, &sample_component.Field_id, &sample_component.Field_station_id, &sample_component.Finance_id, &sample_component.Fossil_id, &sample_component.Incident_id, &sample_component.Incident_set_id, &sample_component.Incident_type, &sample_component.Information_item_id, &sample_component.Info_item_subtype, &sample_component.Instrument_id, &sample_component.Interest_set_id, &sample_component.Interest_set_seq_no, &sample_component.Jurisdiction, &sample_component.Land_right_id, &sample_component.Land_right_subtype, &sample_component.Land_sale_number, &sample_component.Language, &sample_component.Lithology_log_id, &sample_component.Lith_log_source, &sample_component.Notification_id, &sample_component.Obligation_id, &sample_component.Obligation_seq_no, &sample_component.Paleo_summary_id, &sample_component.Pden_id, &sample_component.Pden_source, &sample_component.Pden_subtype, &sample_component.Physical_item_id, &sample_component.Pool_id, &sample_component.Ppdm_guid, &sample_component.Ppdm_system_id, &sample_component.Ppdm_table_name, &sample_component.Product_type, &sample_component.Project_id, &sample_component.Pr_str_source, &sample_component.Pr_str_uwi, &sample_component.Rate_schedule_id, &sample_component.Referenced_guid, &sample_component.Referenced_system_id, &sample_component.Referenced_table_name, &sample_component.Remark, &sample_component.Resent_id, &sample_component.Reserve_class_id, &sample_component.Seis_set_id, &sample_component.Seis_set_subtype, &sample_component.Sf_subtype, &sample_component.Source, &sample_component.Spatial_description_id, &sample_component.Spatial_obs_no, &sample_component.Store_id, &sample_component.Store_structure_obs_no, &sample_component.Strat_name_set_id, &sample_component.Strat_unit_id, &sample_component.String_id, &sample_component.Support_facility_id, &sample_component.Sw_application_id, &sample_component.Thesaurus_id, &sample_component.Thesaurus_word_id, &sample_component.Uwi, &sample_component.Well_activity_set_id, &sample_component.Well_activity_source, &sample_component.Well_activity_type_id, &sample_component.Well_activity_uwi, &sample_component.Well_core_id, &sample_component.Well_core_source, &sample_component.Well_core_uwi, &sample_component.Well_set_id, &sample_component.Work_order_id, &sample_component.Row_changed_by, &sample_component.Row_changed_date, &sample_component.Row_created_by, &sample_component.Row_created_date, &sample_component.Row_effective_date, &sample_component.Row_expiry_date, &sample_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sample_component to result
		result = append(result, sample_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSampleComponent(c *fiber.Ctx) error {
	var sample_component dto.Sample_component

	setDefaults(&sample_component)

	if err := json.Unmarshal(c.Body(), &sample_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sample_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103, :104, :105, :106)")
	if err != nil {
		return err
	}
	sample_component.Row_created_date = formatDate(sample_component.Row_created_date)
	sample_component.Row_changed_date = formatDate(sample_component.Row_changed_date)
	sample_component.Effective_date = formatDateString(sample_component.Effective_date)
	sample_component.Expiry_date = formatDateString(sample_component.Expiry_date)
	sample_component.Row_effective_date = formatDateString(sample_component.Row_effective_date)
	sample_component.Row_expiry_date = formatDateString(sample_component.Row_expiry_date)






	rows, err := stmt.Exec(sample_component.Sample_id, sample_component.Component_obs_no, sample_component.Active_ind, sample_component.Activity_obs_no, sample_component.Analysis_id, sample_component.Anl_source, sample_component.Application_id, sample_component.Area_id, sample_component.Area_type, sample_component.Authority_id, sample_component.Ba_organization_id, sample_component.Ba_organization_seq_no, sample_component.Business_associate_id, sample_component.Catalogue_additive_id, sample_component.Catalogue_equip_id, sample_component.Classification_system_id, sample_component.Class_level_id, sample_component.Component_type, sample_component.Consent_id, sample_component.Consult_id, sample_component.Contest_id, sample_component.Contract_id, sample_component.Ecozone_id, sample_component.Effective_date, sample_component.Employee_ba_id, sample_component.Employee_obs_no, sample_component.Employer_ba_id, sample_component.Entitlement_id, sample_component.Equipment_id, sample_component.Expiry_date, sample_component.Facility_id, sample_component.Facility_type, sample_component.Field_id, sample_component.Field_station_id, sample_component.Finance_id, sample_component.Fossil_id, sample_component.Incident_id, sample_component.Incident_set_id, sample_component.Incident_type, sample_component.Information_item_id, sample_component.Info_item_subtype, sample_component.Instrument_id, sample_component.Interest_set_id, sample_component.Interest_set_seq_no, sample_component.Jurisdiction, sample_component.Land_right_id, sample_component.Land_right_subtype, sample_component.Land_sale_number, sample_component.Language, sample_component.Lithology_log_id, sample_component.Lith_log_source, sample_component.Notification_id, sample_component.Obligation_id, sample_component.Obligation_seq_no, sample_component.Paleo_summary_id, sample_component.Pden_id, sample_component.Pden_source, sample_component.Pden_subtype, sample_component.Physical_item_id, sample_component.Pool_id, sample_component.Ppdm_guid, sample_component.Ppdm_system_id, sample_component.Ppdm_table_name, sample_component.Product_type, sample_component.Project_id, sample_component.Pr_str_source, sample_component.Pr_str_uwi, sample_component.Rate_schedule_id, sample_component.Referenced_guid, sample_component.Referenced_system_id, sample_component.Referenced_table_name, sample_component.Remark, sample_component.Resent_id, sample_component.Reserve_class_id, sample_component.Seis_set_id, sample_component.Seis_set_subtype, sample_component.Sf_subtype, sample_component.Source, sample_component.Spatial_description_id, sample_component.Spatial_obs_no, sample_component.Store_id, sample_component.Store_structure_obs_no, sample_component.Strat_name_set_id, sample_component.Strat_unit_id, sample_component.String_id, sample_component.Support_facility_id, sample_component.Sw_application_id, sample_component.Thesaurus_id, sample_component.Thesaurus_word_id, sample_component.Uwi, sample_component.Well_activity_set_id, sample_component.Well_activity_source, sample_component.Well_activity_type_id, sample_component.Well_activity_uwi, sample_component.Well_core_id, sample_component.Well_core_source, sample_component.Well_core_uwi, sample_component.Well_set_id, sample_component.Work_order_id, sample_component.Row_changed_by, sample_component.Row_changed_date, sample_component.Row_created_by, sample_component.Row_created_date, sample_component.Row_effective_date, sample_component.Row_expiry_date, sample_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSampleComponent(c *fiber.Ctx) error {
	var sample_component dto.Sample_component

	setDefaults(&sample_component)

	if err := json.Unmarshal(c.Body(), &sample_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sample_component.Sample_id = id

    if sample_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sample_component set component_obs_no = :1, active_ind = :2, activity_obs_no = :3, analysis_id = :4, anl_source = :5, application_id = :6, area_id = :7, area_type = :8, authority_id = :9, ba_organization_id = :10, ba_organization_seq_no = :11, business_associate_id = :12, catalogue_additive_id = :13, catalogue_equip_id = :14, classification_system_id = :15, class_level_id = :16, component_type = :17, consent_id = :18, consult_id = :19, contest_id = :20, contract_id = :21, ecozone_id = :22, effective_date = :23, employee_ba_id = :24, employee_obs_no = :25, employer_ba_id = :26, entitlement_id = :27, equipment_id = :28, expiry_date = :29, facility_id = :30, facility_type = :31, field_id = :32, field_station_id = :33, finance_id = :34, fossil_id = :35, incident_id = :36, incident_set_id = :37, incident_type = :38, information_item_id = :39, info_item_subtype = :40, instrument_id = :41, interest_set_id = :42, interest_set_seq_no = :43, jurisdiction = :44, land_right_id = :45, land_right_subtype = :46, land_sale_number = :47, language = :48, lithology_log_id = :49, lith_log_source = :50, notification_id = :51, obligation_id = :52, obligation_seq_no = :53, paleo_summary_id = :54, pden_id = :55, pden_source = :56, pden_subtype = :57, physical_item_id = :58, pool_id = :59, ppdm_guid = :60, ppdm_system_id = :61, ppdm_table_name = :62, product_type = :63, project_id = :64, pr_str_source = :65, pr_str_uwi = :66, rate_schedule_id = :67, referenced_guid = :68, referenced_system_id = :69, referenced_table_name = :70, remark = :71, resent_id = :72, reserve_class_id = :73, seis_set_id = :74, seis_set_subtype = :75, sf_subtype = :76, source = :77, spatial_description_id = :78, spatial_obs_no = :79, store_id = :80, store_structure_obs_no = :81, strat_name_set_id = :82, strat_unit_id = :83, string_id = :84, support_facility_id = :85, sw_application_id = :86, thesaurus_id = :87, thesaurus_word_id = :88, uwi = :89, well_activity_set_id = :90, well_activity_source = :91, well_activity_type_id = :92, well_activity_uwi = :93, well_core_id = :94, well_core_source = :95, well_core_uwi = :96, well_set_id = :97, work_order_id = :98, row_changed_by = :99, row_changed_date = :100, row_created_by = :101, row_effective_date = :102, row_expiry_date = :103, row_quality = :104 where sample_id = :106")
	if err != nil {
		return err
	}

	sample_component.Row_changed_date = formatDate(sample_component.Row_changed_date)
	sample_component.Effective_date = formatDateString(sample_component.Effective_date)
	sample_component.Expiry_date = formatDateString(sample_component.Expiry_date)
	sample_component.Row_effective_date = formatDateString(sample_component.Row_effective_date)
	sample_component.Row_expiry_date = formatDateString(sample_component.Row_expiry_date)






	rows, err := stmt.Exec(sample_component.Component_obs_no, sample_component.Active_ind, sample_component.Activity_obs_no, sample_component.Analysis_id, sample_component.Anl_source, sample_component.Application_id, sample_component.Area_id, sample_component.Area_type, sample_component.Authority_id, sample_component.Ba_organization_id, sample_component.Ba_organization_seq_no, sample_component.Business_associate_id, sample_component.Catalogue_additive_id, sample_component.Catalogue_equip_id, sample_component.Classification_system_id, sample_component.Class_level_id, sample_component.Component_type, sample_component.Consent_id, sample_component.Consult_id, sample_component.Contest_id, sample_component.Contract_id, sample_component.Ecozone_id, sample_component.Effective_date, sample_component.Employee_ba_id, sample_component.Employee_obs_no, sample_component.Employer_ba_id, sample_component.Entitlement_id, sample_component.Equipment_id, sample_component.Expiry_date, sample_component.Facility_id, sample_component.Facility_type, sample_component.Field_id, sample_component.Field_station_id, sample_component.Finance_id, sample_component.Fossil_id, sample_component.Incident_id, sample_component.Incident_set_id, sample_component.Incident_type, sample_component.Information_item_id, sample_component.Info_item_subtype, sample_component.Instrument_id, sample_component.Interest_set_id, sample_component.Interest_set_seq_no, sample_component.Jurisdiction, sample_component.Land_right_id, sample_component.Land_right_subtype, sample_component.Land_sale_number, sample_component.Language, sample_component.Lithology_log_id, sample_component.Lith_log_source, sample_component.Notification_id, sample_component.Obligation_id, sample_component.Obligation_seq_no, sample_component.Paleo_summary_id, sample_component.Pden_id, sample_component.Pden_source, sample_component.Pden_subtype, sample_component.Physical_item_id, sample_component.Pool_id, sample_component.Ppdm_guid, sample_component.Ppdm_system_id, sample_component.Ppdm_table_name, sample_component.Product_type, sample_component.Project_id, sample_component.Pr_str_source, sample_component.Pr_str_uwi, sample_component.Rate_schedule_id, sample_component.Referenced_guid, sample_component.Referenced_system_id, sample_component.Referenced_table_name, sample_component.Remark, sample_component.Resent_id, sample_component.Reserve_class_id, sample_component.Seis_set_id, sample_component.Seis_set_subtype, sample_component.Sf_subtype, sample_component.Source, sample_component.Spatial_description_id, sample_component.Spatial_obs_no, sample_component.Store_id, sample_component.Store_structure_obs_no, sample_component.Strat_name_set_id, sample_component.Strat_unit_id, sample_component.String_id, sample_component.Support_facility_id, sample_component.Sw_application_id, sample_component.Thesaurus_id, sample_component.Thesaurus_word_id, sample_component.Uwi, sample_component.Well_activity_set_id, sample_component.Well_activity_source, sample_component.Well_activity_type_id, sample_component.Well_activity_uwi, sample_component.Well_core_id, sample_component.Well_core_source, sample_component.Well_core_uwi, sample_component.Well_set_id, sample_component.Work_order_id, sample_component.Row_changed_by, sample_component.Row_changed_date, sample_component.Row_created_by, sample_component.Row_effective_date, sample_component.Row_expiry_date, sample_component.Row_quality, sample_component.Sample_id)
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

func PatchSampleComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sample_component set "
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
	query += " where sample_id = :id"

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

func DeleteSampleComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var sample_component dto.Sample_component
	sample_component.Sample_id = id

	stmt, err := db.Prepare("delete from sample_component where sample_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sample_component.Sample_id)
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


