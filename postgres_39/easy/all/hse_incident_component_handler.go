package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetHseIncidentComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from hse_incident_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Hse_incident_component

	for rows.Next() {
		var hse_incident_component dto.Hse_incident_component
		if err := rows.Scan(&hse_incident_component.Incident_id, &hse_incident_component.Component_obs_no, &hse_incident_component.Active_ind, &hse_incident_component.Activity_obs_no, &hse_incident_component.Address_obs_no, &hse_incident_component.Address_source, &hse_incident_component.Analysis_id, &hse_incident_component.Application_id, &hse_incident_component.Area_id, &hse_incident_component.Area_type, &hse_incident_component.Authority_id, &hse_incident_component.Ba_organization_id, &hse_incident_component.Ba_organization_seq_no, &hse_incident_component.Business_associate_id, &hse_incident_component.Catalogue_additive_id, &hse_incident_component.Catalogue_equip_id, &hse_incident_component.Classification_system_id, &hse_incident_component.Class_level_id, &hse_incident_component.Component_role, &hse_incident_component.Component_type, &hse_incident_component.Consent_id, &hse_incident_component.Consult_id, &hse_incident_component.Contest_id, &hse_incident_component.Contract_id, &hse_incident_component.Crew_company_ba_id, &hse_incident_component.Crew_id, &hse_incident_component.Crew_member_ba_id, &hse_incident_component.Crew_member_obs_no, &hse_incident_component.Ecozone_id, &hse_incident_component.Effective_date, &hse_incident_component.Employee_ba_id, &hse_incident_component.Employee_obs_no, &hse_incident_component.Employer_ba_id, &hse_incident_component.Entitlement_id, &hse_incident_component.Equipment_id, &hse_incident_component.Expiry_date, &hse_incident_component.Facility_id, &hse_incident_component.Facility_type, &hse_incident_component.Field_id, &hse_incident_component.Field_station_id, &hse_incident_component.Finance_id, &hse_incident_component.Fossil_id, &hse_incident_component.Incident_set_id, &hse_incident_component.Incident_type, &hse_incident_component.Information_item_id, &hse_incident_component.Info_item_subtype, &hse_incident_component.Instrument_id, &hse_incident_component.Interest_set_id, &hse_incident_component.Interest_set_seq_no, &hse_incident_component.Jurisdiction, &hse_incident_component.Land_right_id, &hse_incident_component.Land_right_subtype, &hse_incident_component.Land_sale_number, &hse_incident_component.Language, &hse_incident_component.Lithology_log_id, &hse_incident_component.Lith_log_source, &hse_incident_component.Notification_id, &hse_incident_component.Obligation_id, &hse_incident_component.Obligation_seq_no, &hse_incident_component.Paleo_summary_id, &hse_incident_component.Pden_activity_type, &hse_incident_component.Pden_amendment_seq_no, &hse_incident_component.Pden_id, &hse_incident_component.Pden_period_id, &hse_incident_component.Pden_period_type, &hse_incident_component.Pden_source, &hse_incident_component.Pden_subtype, &hse_incident_component.Pden_volume_date, &hse_incident_component.Pden_volume_date_desc, &hse_incident_component.Pden_volume_method, &hse_incident_component.Physical_item_id, &hse_incident_component.Pool_id, &hse_incident_component.Ppdm_guid, &hse_incident_component.Ppdm_system_id, &hse_incident_component.Ppdm_table_name, &hse_incident_component.Product_type, &hse_incident_component.Prod_string_id, &hse_incident_component.Prod_string_source, &hse_incident_component.Project_id, &hse_incident_component.Project_plan_id, &hse_incident_component.Pr_str_form_obs_no, &hse_incident_component.Rate_schedule_id, &hse_incident_component.Referenced_guid, &hse_incident_component.Referenced_system_id, &hse_incident_component.Referenced_table_name, &hse_incident_component.Remark, &hse_incident_component.Resent_id, &hse_incident_component.Reserve_class_id, &hse_incident_component.Restriction_id, &hse_incident_component.Restriction_version, &hse_incident_component.Sample_anal_source, &hse_incident_component.Seis_set_id, &hse_incident_component.Seis_set_subtype, &hse_incident_component.Sf_subtype, &hse_incident_component.Source, &hse_incident_component.Spatial_description_id, &hse_incident_component.Spatial_obs_no, &hse_incident_component.Store_id, &hse_incident_component.Store_structure_obs_no, &hse_incident_component.Strat_name_set_id, &hse_incident_component.Strat_unit_id, &hse_incident_component.Support_facility_id, &hse_incident_component.Sw_application_id, &hse_incident_component.Thesaurus_id, &hse_incident_component.Thesaurus_word_id, &hse_incident_component.Total_incident_count, &hse_incident_component.Total_incident_period, &hse_incident_component.Total_incident_period_uom, &hse_incident_component.Uwi, &hse_incident_component.Well_activity_set_id, &hse_incident_component.Well_activity_source, &hse_incident_component.Well_activity_type_id, &hse_incident_component.Well_drill_period_obs_no, &hse_incident_component.Well_set_id, &hse_incident_component.Work_order_id, &hse_incident_component.Row_changed_by, &hse_incident_component.Row_changed_date, &hse_incident_component.Row_created_by, &hse_incident_component.Row_created_date, &hse_incident_component.Row_effective_date, &hse_incident_component.Row_expiry_date, &hse_incident_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append hse_incident_component to result
		result = append(result, hse_incident_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetHseIncidentComponent(c *fiber.Ctx) error {
	var hse_incident_component dto.Hse_incident_component

	setDefaults(&hse_incident_component)

	if err := json.Unmarshal(c.Body(), &hse_incident_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into hse_incident_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103, :104, :105, :106, :107, :108, :109, :110, :111, :112, :113, :114, :115, :116, :117, :118, :119, :120, :121, :122)")
	if err != nil {
		return err
	}
	hse_incident_component.Row_created_date = formatDate(hse_incident_component.Row_created_date)
	hse_incident_component.Row_changed_date = formatDate(hse_incident_component.Row_changed_date)
	hse_incident_component.Effective_date = formatDateString(hse_incident_component.Effective_date)
	hse_incident_component.Expiry_date = formatDateString(hse_incident_component.Expiry_date)
	hse_incident_component.Pden_volume_date = formatDateString(hse_incident_component.Pden_volume_date)
	hse_incident_component.Row_effective_date = formatDateString(hse_incident_component.Row_effective_date)
	hse_incident_component.Row_expiry_date = formatDateString(hse_incident_component.Row_expiry_date)






	rows, err := stmt.Exec(hse_incident_component.Incident_id, hse_incident_component.Component_obs_no, hse_incident_component.Active_ind, hse_incident_component.Activity_obs_no, hse_incident_component.Address_obs_no, hse_incident_component.Address_source, hse_incident_component.Analysis_id, hse_incident_component.Application_id, hse_incident_component.Area_id, hse_incident_component.Area_type, hse_incident_component.Authority_id, hse_incident_component.Ba_organization_id, hse_incident_component.Ba_organization_seq_no, hse_incident_component.Business_associate_id, hse_incident_component.Catalogue_additive_id, hse_incident_component.Catalogue_equip_id, hse_incident_component.Classification_system_id, hse_incident_component.Class_level_id, hse_incident_component.Component_role, hse_incident_component.Component_type, hse_incident_component.Consent_id, hse_incident_component.Consult_id, hse_incident_component.Contest_id, hse_incident_component.Contract_id, hse_incident_component.Crew_company_ba_id, hse_incident_component.Crew_id, hse_incident_component.Crew_member_ba_id, hse_incident_component.Crew_member_obs_no, hse_incident_component.Ecozone_id, hse_incident_component.Effective_date, hse_incident_component.Employee_ba_id, hse_incident_component.Employee_obs_no, hse_incident_component.Employer_ba_id, hse_incident_component.Entitlement_id, hse_incident_component.Equipment_id, hse_incident_component.Expiry_date, hse_incident_component.Facility_id, hse_incident_component.Facility_type, hse_incident_component.Field_id, hse_incident_component.Field_station_id, hse_incident_component.Finance_id, hse_incident_component.Fossil_id, hse_incident_component.Incident_set_id, hse_incident_component.Incident_type, hse_incident_component.Information_item_id, hse_incident_component.Info_item_subtype, hse_incident_component.Instrument_id, hse_incident_component.Interest_set_id, hse_incident_component.Interest_set_seq_no, hse_incident_component.Jurisdiction, hse_incident_component.Land_right_id, hse_incident_component.Land_right_subtype, hse_incident_component.Land_sale_number, hse_incident_component.Language, hse_incident_component.Lithology_log_id, hse_incident_component.Lith_log_source, hse_incident_component.Notification_id, hse_incident_component.Obligation_id, hse_incident_component.Obligation_seq_no, hse_incident_component.Paleo_summary_id, hse_incident_component.Pden_activity_type, hse_incident_component.Pden_amendment_seq_no, hse_incident_component.Pden_id, hse_incident_component.Pden_period_id, hse_incident_component.Pden_period_type, hse_incident_component.Pden_source, hse_incident_component.Pden_subtype, hse_incident_component.Pden_volume_date, hse_incident_component.Pden_volume_date_desc, hse_incident_component.Pden_volume_method, hse_incident_component.Physical_item_id, hse_incident_component.Pool_id, hse_incident_component.Ppdm_guid, hse_incident_component.Ppdm_system_id, hse_incident_component.Ppdm_table_name, hse_incident_component.Product_type, hse_incident_component.Prod_string_id, hse_incident_component.Prod_string_source, hse_incident_component.Project_id, hse_incident_component.Project_plan_id, hse_incident_component.Pr_str_form_obs_no, hse_incident_component.Rate_schedule_id, hse_incident_component.Referenced_guid, hse_incident_component.Referenced_system_id, hse_incident_component.Referenced_table_name, hse_incident_component.Remark, hse_incident_component.Resent_id, hse_incident_component.Reserve_class_id, hse_incident_component.Restriction_id, hse_incident_component.Restriction_version, hse_incident_component.Sample_anal_source, hse_incident_component.Seis_set_id, hse_incident_component.Seis_set_subtype, hse_incident_component.Sf_subtype, hse_incident_component.Source, hse_incident_component.Spatial_description_id, hse_incident_component.Spatial_obs_no, hse_incident_component.Store_id, hse_incident_component.Store_structure_obs_no, hse_incident_component.Strat_name_set_id, hse_incident_component.Strat_unit_id, hse_incident_component.Support_facility_id, hse_incident_component.Sw_application_id, hse_incident_component.Thesaurus_id, hse_incident_component.Thesaurus_word_id, hse_incident_component.Total_incident_count, hse_incident_component.Total_incident_period, hse_incident_component.Total_incident_period_uom, hse_incident_component.Uwi, hse_incident_component.Well_activity_set_id, hse_incident_component.Well_activity_source, hse_incident_component.Well_activity_type_id, hse_incident_component.Well_drill_period_obs_no, hse_incident_component.Well_set_id, hse_incident_component.Work_order_id, hse_incident_component.Row_changed_by, hse_incident_component.Row_changed_date, hse_incident_component.Row_created_by, hse_incident_component.Row_created_date, hse_incident_component.Row_effective_date, hse_incident_component.Row_expiry_date, hse_incident_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateHseIncidentComponent(c *fiber.Ctx) error {
	var hse_incident_component dto.Hse_incident_component

	setDefaults(&hse_incident_component)

	if err := json.Unmarshal(c.Body(), &hse_incident_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	hse_incident_component.Incident_id = id

    if hse_incident_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update hse_incident_component set component_obs_no = :1, active_ind = :2, activity_obs_no = :3, address_obs_no = :4, address_source = :5, analysis_id = :6, application_id = :7, area_id = :8, area_type = :9, authority_id = :10, ba_organization_id = :11, ba_organization_seq_no = :12, business_associate_id = :13, catalogue_additive_id = :14, catalogue_equip_id = :15, classification_system_id = :16, class_level_id = :17, component_role = :18, component_type = :19, consent_id = :20, consult_id = :21, contest_id = :22, contract_id = :23, crew_company_ba_id = :24, crew_id = :25, crew_member_ba_id = :26, crew_member_obs_no = :27, ecozone_id = :28, effective_date = :29, employee_ba_id = :30, employee_obs_no = :31, employer_ba_id = :32, entitlement_id = :33, equipment_id = :34, expiry_date = :35, facility_id = :36, facility_type = :37, field_id = :38, field_station_id = :39, finance_id = :40, fossil_id = :41, incident_set_id = :42, incident_type = :43, information_item_id = :44, info_item_subtype = :45, instrument_id = :46, interest_set_id = :47, interest_set_seq_no = :48, jurisdiction = :49, land_right_id = :50, land_right_subtype = :51, land_sale_number = :52, language = :53, lithology_log_id = :54, lith_log_source = :55, notification_id = :56, obligation_id = :57, obligation_seq_no = :58, paleo_summary_id = :59, pden_activity_type = :60, pden_amendment_seq_no = :61, pden_id = :62, pden_period_id = :63, pden_period_type = :64, pden_source = :65, pden_subtype = :66, pden_volume_date = :67, pden_volume_date_desc = :68, pden_volume_method = :69, physical_item_id = :70, pool_id = :71, ppdm_guid = :72, ppdm_system_id = :73, ppdm_table_name = :74, product_type = :75, prod_string_id = :76, prod_string_source = :77, project_id = :78, project_plan_id = :79, pr_str_form_obs_no = :80, rate_schedule_id = :81, referenced_guid = :82, referenced_system_id = :83, referenced_table_name = :84, remark = :85, resent_id = :86, reserve_class_id = :87, restriction_id = :88, restriction_version = :89, sample_anal_source = :90, seis_set_id = :91, seis_set_subtype = :92, sf_subtype = :93, source = :94, spatial_description_id = :95, spatial_obs_no = :96, store_id = :97, store_structure_obs_no = :98, strat_name_set_id = :99, strat_unit_id = :100, support_facility_id = :101, sw_application_id = :102, thesaurus_id = :103, thesaurus_word_id = :104, total_incident_count = :105, total_incident_period = :106, total_incident_period_uom = :107, uwi = :108, well_activity_set_id = :109, well_activity_source = :110, well_activity_type_id = :111, well_drill_period_obs_no = :112, well_set_id = :113, work_order_id = :114, row_changed_by = :115, row_changed_date = :116, row_created_by = :117, row_effective_date = :118, row_expiry_date = :119, row_quality = :120 where incident_id = :122")
	if err != nil {
		return err
	}

	hse_incident_component.Row_changed_date = formatDate(hse_incident_component.Row_changed_date)
	hse_incident_component.Effective_date = formatDateString(hse_incident_component.Effective_date)
	hse_incident_component.Expiry_date = formatDateString(hse_incident_component.Expiry_date)
	hse_incident_component.Pden_volume_date = formatDateString(hse_incident_component.Pden_volume_date)
	hse_incident_component.Row_effective_date = formatDateString(hse_incident_component.Row_effective_date)
	hse_incident_component.Row_expiry_date = formatDateString(hse_incident_component.Row_expiry_date)






	rows, err := stmt.Exec(hse_incident_component.Component_obs_no, hse_incident_component.Active_ind, hse_incident_component.Activity_obs_no, hse_incident_component.Address_obs_no, hse_incident_component.Address_source, hse_incident_component.Analysis_id, hse_incident_component.Application_id, hse_incident_component.Area_id, hse_incident_component.Area_type, hse_incident_component.Authority_id, hse_incident_component.Ba_organization_id, hse_incident_component.Ba_organization_seq_no, hse_incident_component.Business_associate_id, hse_incident_component.Catalogue_additive_id, hse_incident_component.Catalogue_equip_id, hse_incident_component.Classification_system_id, hse_incident_component.Class_level_id, hse_incident_component.Component_role, hse_incident_component.Component_type, hse_incident_component.Consent_id, hse_incident_component.Consult_id, hse_incident_component.Contest_id, hse_incident_component.Contract_id, hse_incident_component.Crew_company_ba_id, hse_incident_component.Crew_id, hse_incident_component.Crew_member_ba_id, hse_incident_component.Crew_member_obs_no, hse_incident_component.Ecozone_id, hse_incident_component.Effective_date, hse_incident_component.Employee_ba_id, hse_incident_component.Employee_obs_no, hse_incident_component.Employer_ba_id, hse_incident_component.Entitlement_id, hse_incident_component.Equipment_id, hse_incident_component.Expiry_date, hse_incident_component.Facility_id, hse_incident_component.Facility_type, hse_incident_component.Field_id, hse_incident_component.Field_station_id, hse_incident_component.Finance_id, hse_incident_component.Fossil_id, hse_incident_component.Incident_set_id, hse_incident_component.Incident_type, hse_incident_component.Information_item_id, hse_incident_component.Info_item_subtype, hse_incident_component.Instrument_id, hse_incident_component.Interest_set_id, hse_incident_component.Interest_set_seq_no, hse_incident_component.Jurisdiction, hse_incident_component.Land_right_id, hse_incident_component.Land_right_subtype, hse_incident_component.Land_sale_number, hse_incident_component.Language, hse_incident_component.Lithology_log_id, hse_incident_component.Lith_log_source, hse_incident_component.Notification_id, hse_incident_component.Obligation_id, hse_incident_component.Obligation_seq_no, hse_incident_component.Paleo_summary_id, hse_incident_component.Pden_activity_type, hse_incident_component.Pden_amendment_seq_no, hse_incident_component.Pden_id, hse_incident_component.Pden_period_id, hse_incident_component.Pden_period_type, hse_incident_component.Pden_source, hse_incident_component.Pden_subtype, hse_incident_component.Pden_volume_date, hse_incident_component.Pden_volume_date_desc, hse_incident_component.Pden_volume_method, hse_incident_component.Physical_item_id, hse_incident_component.Pool_id, hse_incident_component.Ppdm_guid, hse_incident_component.Ppdm_system_id, hse_incident_component.Ppdm_table_name, hse_incident_component.Product_type, hse_incident_component.Prod_string_id, hse_incident_component.Prod_string_source, hse_incident_component.Project_id, hse_incident_component.Project_plan_id, hse_incident_component.Pr_str_form_obs_no, hse_incident_component.Rate_schedule_id, hse_incident_component.Referenced_guid, hse_incident_component.Referenced_system_id, hse_incident_component.Referenced_table_name, hse_incident_component.Remark, hse_incident_component.Resent_id, hse_incident_component.Reserve_class_id, hse_incident_component.Restriction_id, hse_incident_component.Restriction_version, hse_incident_component.Sample_anal_source, hse_incident_component.Seis_set_id, hse_incident_component.Seis_set_subtype, hse_incident_component.Sf_subtype, hse_incident_component.Source, hse_incident_component.Spatial_description_id, hse_incident_component.Spatial_obs_no, hse_incident_component.Store_id, hse_incident_component.Store_structure_obs_no, hse_incident_component.Strat_name_set_id, hse_incident_component.Strat_unit_id, hse_incident_component.Support_facility_id, hse_incident_component.Sw_application_id, hse_incident_component.Thesaurus_id, hse_incident_component.Thesaurus_word_id, hse_incident_component.Total_incident_count, hse_incident_component.Total_incident_period, hse_incident_component.Total_incident_period_uom, hse_incident_component.Uwi, hse_incident_component.Well_activity_set_id, hse_incident_component.Well_activity_source, hse_incident_component.Well_activity_type_id, hse_incident_component.Well_drill_period_obs_no, hse_incident_component.Well_set_id, hse_incident_component.Work_order_id, hse_incident_component.Row_changed_by, hse_incident_component.Row_changed_date, hse_incident_component.Row_created_by, hse_incident_component.Row_effective_date, hse_incident_component.Row_expiry_date, hse_incident_component.Row_quality, hse_incident_component.Incident_id)
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

func PatchHseIncidentComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update hse_incident_component set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "pden_volume_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where incident_id = :id"

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

func DeleteHseIncidentComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var hse_incident_component dto.Hse_incident_component
	hse_incident_component.Incident_id = id

	stmt, err := db.Prepare("delete from hse_incident_component where incident_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(hse_incident_component.Incident_id)
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


