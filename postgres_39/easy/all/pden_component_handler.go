package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_component

	for rows.Next() {
		var pden_component dto.Pden_component
		if err := rows.Scan(&pden_component.Pden_subtype, &pden_component.Pden_id, &pden_component.Pden_source, &pden_component.Component_obs_no, &pden_component.Active_ind, &pden_component.Activity_obs_no, &pden_component.Analysis_id, &pden_component.Application_id, &pden_component.Area_id, &pden_component.Area_type, &pden_component.Authority_id, &pden_component.Ba_organization_id, &pden_component.Ba_organization_seq_no, &pden_component.Business_associate_id, &pden_component.Catalogue_additive_id, &pden_component.Catalogue_equip_id, &pden_component.Classification_system_id, &pden_component.Class_level_id, &pden_component.Consent_id, &pden_component.Consult_id, &pden_component.Contest_id, &pden_component.Contract_id, &pden_component.Ecozone_id, &pden_component.Effective_date, &pden_component.Employee_ba_id, &pden_component.Employee_obs_no, &pden_component.Employer_ba_id, &pden_component.Entitlement_id, &pden_component.Equipment_id, &pden_component.Expiry_date, &pden_component.Facility_id, &pden_component.Facility_type, &pden_component.Field_id, &pden_component.Field_station_id, &pden_component.Finance_id, &pden_component.Fossil_id, &pden_component.Incident_id, &pden_component.Incident_set_id, &pden_component.Incident_type, &pden_component.Information_item_id, &pden_component.Info_item_subtype, &pden_component.Instrument_id, &pden_component.Interest_set_id, &pden_component.Interest_set_seq_no, &pden_component.Jurisdiction, &pden_component.Land_right_id, &pden_component.Land_right_subtype, &pden_component.Land_sale_number, &pden_component.Language, &pden_component.Lithology_log_id, &pden_component.Lith_log_source, &pden_component.Notification_id, &pden_component.Obligation_id, &pden_component.Obligation_seq_no, &pden_component.Paleo_summary_id, &pden_component.Pden_component_type, &pden_component.Physical_item_id, &pden_component.Pool_id, &pden_component.Ppdm_guid, &pden_component.Ppdm_system_id, &pden_component.Ppdm_table_name, &pden_component.Product_type, &pden_component.Project_id, &pden_component.Pr_str_form_obs_no, &pden_component.Pr_str_source, &pden_component.Pr_str_uwi, &pden_component.Rate_schedule_id, &pden_component.Referenced_guid, &pden_component.Referenced_system_id, &pden_component.Referenced_table_name, &pden_component.Remark, &pden_component.Resent_id, &pden_component.Reserve_class_id, &pden_component.Sample_anal_source, &pden_component.Seis_set_id, &pden_component.Seis_set_subtype, &pden_component.Sf_subtype, &pden_component.Source, &pden_component.Spatial_description_id, &pden_component.Spatial_obs_no, &pden_component.Store_id, &pden_component.Store_structure_obs_no, &pden_component.Strat_name_set_id, &pden_component.Strat_unit_id, &pden_component.String_id, &pden_component.Support_facility_id, &pden_component.Sw_application_id, &pden_component.Thesaurus_id, &pden_component.Thesaurus_word_id, &pden_component.Uwi, &pden_component.Well_activity_set_id, &pden_component.Well_activity_source, &pden_component.Well_activity_type_id, &pden_component.Well_activity_uwi, &pden_component.Well_set_id, &pden_component.Work_order_id, &pden_component.Row_changed_by, &pden_component.Row_changed_date, &pden_component.Row_created_by, &pden_component.Row_created_date, &pden_component.Row_effective_date, &pden_component.Row_expiry_date, &pden_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_component to result
		result = append(result, pden_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenComponent(c *fiber.Ctx) error {
	var pden_component dto.Pden_component

	setDefaults(&pden_component)

	if err := json.Unmarshal(c.Body(), &pden_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103)")
	if err != nil {
		return err
	}
	pden_component.Row_created_date = formatDate(pden_component.Row_created_date)
	pden_component.Row_changed_date = formatDate(pden_component.Row_changed_date)
	pden_component.Effective_date = formatDateString(pden_component.Effective_date)
	pden_component.Expiry_date = formatDateString(pden_component.Expiry_date)
	pden_component.Row_effective_date = formatDateString(pden_component.Row_effective_date)
	pden_component.Row_expiry_date = formatDateString(pden_component.Row_expiry_date)






	rows, err := stmt.Exec(pden_component.Pden_subtype, pden_component.Pden_id, pden_component.Pden_source, pden_component.Component_obs_no, pden_component.Active_ind, pden_component.Activity_obs_no, pden_component.Analysis_id, pden_component.Application_id, pden_component.Area_id, pden_component.Area_type, pden_component.Authority_id, pden_component.Ba_organization_id, pden_component.Ba_organization_seq_no, pden_component.Business_associate_id, pden_component.Catalogue_additive_id, pden_component.Catalogue_equip_id, pden_component.Classification_system_id, pden_component.Class_level_id, pden_component.Consent_id, pden_component.Consult_id, pden_component.Contest_id, pden_component.Contract_id, pden_component.Ecozone_id, pden_component.Effective_date, pden_component.Employee_ba_id, pden_component.Employee_obs_no, pden_component.Employer_ba_id, pden_component.Entitlement_id, pden_component.Equipment_id, pden_component.Expiry_date, pden_component.Facility_id, pden_component.Facility_type, pden_component.Field_id, pden_component.Field_station_id, pden_component.Finance_id, pden_component.Fossil_id, pden_component.Incident_id, pden_component.Incident_set_id, pden_component.Incident_type, pden_component.Information_item_id, pden_component.Info_item_subtype, pden_component.Instrument_id, pden_component.Interest_set_id, pden_component.Interest_set_seq_no, pden_component.Jurisdiction, pden_component.Land_right_id, pden_component.Land_right_subtype, pden_component.Land_sale_number, pden_component.Language, pden_component.Lithology_log_id, pden_component.Lith_log_source, pden_component.Notification_id, pden_component.Obligation_id, pden_component.Obligation_seq_no, pden_component.Paleo_summary_id, pden_component.Pden_component_type, pden_component.Physical_item_id, pden_component.Pool_id, pden_component.Ppdm_guid, pden_component.Ppdm_system_id, pden_component.Ppdm_table_name, pden_component.Product_type, pden_component.Project_id, pden_component.Pr_str_form_obs_no, pden_component.Pr_str_source, pden_component.Pr_str_uwi, pden_component.Rate_schedule_id, pden_component.Referenced_guid, pden_component.Referenced_system_id, pden_component.Referenced_table_name, pden_component.Remark, pden_component.Resent_id, pden_component.Reserve_class_id, pden_component.Sample_anal_source, pden_component.Seis_set_id, pden_component.Seis_set_subtype, pden_component.Sf_subtype, pden_component.Source, pden_component.Spatial_description_id, pden_component.Spatial_obs_no, pden_component.Store_id, pden_component.Store_structure_obs_no, pden_component.Strat_name_set_id, pden_component.Strat_unit_id, pden_component.String_id, pden_component.Support_facility_id, pden_component.Sw_application_id, pden_component.Thesaurus_id, pden_component.Thesaurus_word_id, pden_component.Uwi, pden_component.Well_activity_set_id, pden_component.Well_activity_source, pden_component.Well_activity_type_id, pden_component.Well_activity_uwi, pden_component.Well_set_id, pden_component.Work_order_id, pden_component.Row_changed_by, pden_component.Row_changed_date, pden_component.Row_created_by, pden_component.Row_created_date, pden_component.Row_effective_date, pden_component.Row_expiry_date, pden_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenComponent(c *fiber.Ctx) error {
	var pden_component dto.Pden_component

	setDefaults(&pden_component)

	if err := json.Unmarshal(c.Body(), &pden_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_component.Pden_subtype = id

    if pden_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_component set pden_id = :1, pden_source = :2, component_obs_no = :3, active_ind = :4, activity_obs_no = :5, analysis_id = :6, application_id = :7, area_id = :8, area_type = :9, authority_id = :10, ba_organization_id = :11, ba_organization_seq_no = :12, business_associate_id = :13, catalogue_additive_id = :14, catalogue_equip_id = :15, classification_system_id = :16, class_level_id = :17, consent_id = :18, consult_id = :19, contest_id = :20, contract_id = :21, ecozone_id = :22, effective_date = :23, employee_ba_id = :24, employee_obs_no = :25, employer_ba_id = :26, entitlement_id = :27, equipment_id = :28, expiry_date = :29, facility_id = :30, facility_type = :31, field_id = :32, field_station_id = :33, finance_id = :34, fossil_id = :35, incident_id = :36, incident_set_id = :37, incident_type = :38, information_item_id = :39, info_item_subtype = :40, instrument_id = :41, interest_set_id = :42, interest_set_seq_no = :43, jurisdiction = :44, land_right_id = :45, land_right_subtype = :46, land_sale_number = :47, language = :48, lithology_log_id = :49, lith_log_source = :50, notification_id = :51, obligation_id = :52, obligation_seq_no = :53, paleo_summary_id = :54, pden_component_type = :55, physical_item_id = :56, pool_id = :57, ppdm_guid = :58, ppdm_system_id = :59, ppdm_table_name = :60, product_type = :61, project_id = :62, pr_str_form_obs_no = :63, pr_str_source = :64, pr_str_uwi = :65, rate_schedule_id = :66, referenced_guid = :67, referenced_system_id = :68, referenced_table_name = :69, remark = :70, resent_id = :71, reserve_class_id = :72, sample_anal_source = :73, seis_set_id = :74, seis_set_subtype = :75, sf_subtype = :76, source = :77, spatial_description_id = :78, spatial_obs_no = :79, store_id = :80, store_structure_obs_no = :81, strat_name_set_id = :82, strat_unit_id = :83, string_id = :84, support_facility_id = :85, sw_application_id = :86, thesaurus_id = :87, thesaurus_word_id = :88, uwi = :89, well_activity_set_id = :90, well_activity_source = :91, well_activity_type_id = :92, well_activity_uwi = :93, well_set_id = :94, work_order_id = :95, row_changed_by = :96, row_changed_date = :97, row_created_by = :98, row_effective_date = :99, row_expiry_date = :100, row_quality = :101 where pden_subtype = :103")
	if err != nil {
		return err
	}

	pden_component.Row_changed_date = formatDate(pden_component.Row_changed_date)
	pden_component.Effective_date = formatDateString(pden_component.Effective_date)
	pden_component.Expiry_date = formatDateString(pden_component.Expiry_date)
	pden_component.Row_effective_date = formatDateString(pden_component.Row_effective_date)
	pden_component.Row_expiry_date = formatDateString(pden_component.Row_expiry_date)






	rows, err := stmt.Exec(pden_component.Pden_id, pden_component.Pden_source, pden_component.Component_obs_no, pden_component.Active_ind, pden_component.Activity_obs_no, pden_component.Analysis_id, pden_component.Application_id, pden_component.Area_id, pden_component.Area_type, pden_component.Authority_id, pden_component.Ba_organization_id, pden_component.Ba_organization_seq_no, pden_component.Business_associate_id, pden_component.Catalogue_additive_id, pden_component.Catalogue_equip_id, pden_component.Classification_system_id, pden_component.Class_level_id, pden_component.Consent_id, pden_component.Consult_id, pden_component.Contest_id, pden_component.Contract_id, pden_component.Ecozone_id, pden_component.Effective_date, pden_component.Employee_ba_id, pden_component.Employee_obs_no, pden_component.Employer_ba_id, pden_component.Entitlement_id, pden_component.Equipment_id, pden_component.Expiry_date, pden_component.Facility_id, pden_component.Facility_type, pden_component.Field_id, pden_component.Field_station_id, pden_component.Finance_id, pden_component.Fossil_id, pden_component.Incident_id, pden_component.Incident_set_id, pden_component.Incident_type, pden_component.Information_item_id, pden_component.Info_item_subtype, pden_component.Instrument_id, pden_component.Interest_set_id, pden_component.Interest_set_seq_no, pden_component.Jurisdiction, pden_component.Land_right_id, pden_component.Land_right_subtype, pden_component.Land_sale_number, pden_component.Language, pden_component.Lithology_log_id, pden_component.Lith_log_source, pden_component.Notification_id, pden_component.Obligation_id, pden_component.Obligation_seq_no, pden_component.Paleo_summary_id, pden_component.Pden_component_type, pden_component.Physical_item_id, pden_component.Pool_id, pden_component.Ppdm_guid, pden_component.Ppdm_system_id, pden_component.Ppdm_table_name, pden_component.Product_type, pden_component.Project_id, pden_component.Pr_str_form_obs_no, pden_component.Pr_str_source, pden_component.Pr_str_uwi, pden_component.Rate_schedule_id, pden_component.Referenced_guid, pden_component.Referenced_system_id, pden_component.Referenced_table_name, pden_component.Remark, pden_component.Resent_id, pden_component.Reserve_class_id, pden_component.Sample_anal_source, pden_component.Seis_set_id, pden_component.Seis_set_subtype, pden_component.Sf_subtype, pden_component.Source, pden_component.Spatial_description_id, pden_component.Spatial_obs_no, pden_component.Store_id, pden_component.Store_structure_obs_no, pden_component.Strat_name_set_id, pden_component.Strat_unit_id, pden_component.String_id, pden_component.Support_facility_id, pden_component.Sw_application_id, pden_component.Thesaurus_id, pden_component.Thesaurus_word_id, pden_component.Uwi, pden_component.Well_activity_set_id, pden_component.Well_activity_source, pden_component.Well_activity_type_id, pden_component.Well_activity_uwi, pden_component.Well_set_id, pden_component.Work_order_id, pden_component.Row_changed_by, pden_component.Row_changed_date, pden_component.Row_created_by, pden_component.Row_effective_date, pden_component.Row_expiry_date, pden_component.Row_quality, pden_component.Pden_subtype)
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

func PatchPdenComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_component set "
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
	query += " where pden_subtype = :id"

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

func DeletePdenComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_component dto.Pden_component
	pden_component.Pden_subtype = id

	stmt, err := db.Prepare("delete from pden_component where pden_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_component.Pden_subtype)
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


