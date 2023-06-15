package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_component

	for rows.Next() {
		var sp_component dto.Sp_component
		if err := rows.Scan(&sp_component.Spatial_description_id, &sp_component.Spatial_obs_no, &sp_component.Component_id, &sp_component.Active_ind, &sp_component.Activity_obs_no, &sp_component.Analysis_id, &sp_component.Application_id, &sp_component.Area_id, &sp_component.Area_type, &sp_component.Authority_id, &sp_component.Ba_address_obs_no, &sp_component.Ba_address_source, &sp_component.Ba_licensee_ba_id, &sp_component.Ba_license_id, &sp_component.Ba_organization_id, &sp_component.Ba_organization_seq_no, &sp_component.Business_associate_id, &sp_component.Catalogue_additive_id, &sp_component.Catalogue_equip_id, &sp_component.Classification_system_id, &sp_component.Class_level_id, &sp_component.Consent_id, &sp_component.Consult_id, &sp_component.Contest_id, &sp_component.Contract_id, &sp_component.Ecozone_id, &sp_component.Effective_date, &sp_component.Employee_ba_id, &sp_component.Employee_obs_no, &sp_component.Employer_ba_id, &sp_component.Entitlement_id, &sp_component.Equipment_id, &sp_component.Expiry_date, &sp_component.Facility_id, &sp_component.Facility_type, &sp_component.Field_id, &sp_component.Field_node_id, &sp_component.Field_node_source, &sp_component.Field_station_id, &sp_component.Finance_id, &sp_component.Fossil_id, &sp_component.Hse_incident_id, &sp_component.Incident_set_id, &sp_component.Incident_type, &sp_component.Information_item_id, &sp_component.Info_item_subtype, &sp_component.Instrument_id, &sp_component.Interest_set_id, &sp_component.Interest_set_seq_no, &sp_component.Jurisdiction, &sp_component.Land_request_id, &sp_component.Land_right_id, &sp_component.Land_right_status, &sp_component.Land_right_status_seq_no, &sp_component.Land_right_status_type, &sp_component.Land_right_subtype, &sp_component.Land_sale_number, &sp_component.Land_sale_offering_id, &sp_component.Language, &sp_component.Lithology_log_id, &sp_component.Lith_log_source, &sp_component.Lith_sample_id, &sp_component.Node_id, &sp_component.Notification_id, &sp_component.Obligation_id, &sp_component.Obligation_seq_no, &sp_component.Paleo_summary_id, &sp_component.Pden_id, &sp_component.Pden_source, &sp_component.Pden_subtype, &sp_component.Physical_item_id, &sp_component.Pool_id, &sp_component.Ppdm_guid, &sp_component.Ppdm_system_id, &sp_component.Ppdm_table_name, &sp_component.Product_type, &sp_component.Prod_string_id, &sp_component.Prod_string_source, &sp_component.Project_id, &sp_component.Rate_schedule_id, &sp_component.Referenced_guid, &sp_component.Referenced_system_id, &sp_component.Referenced_table_name, &sp_component.Remark, &sp_component.Resent_id, &sp_component.Resent_product_type, &sp_component.Resent_reserve_class_id, &sp_component.Reserve_class_id, &sp_component.Restriction_id, &sp_component.Restriction_version, &sp_component.Sample_anal_source, &sp_component.Seis_bin_grid_id, &sp_component.Seis_bin_source, &sp_component.Seis_set_id, &sp_component.Seis_set_subtype, &sp_component.Sf_subtype, &sp_component.Size_seq_no, &sp_component.Size_type, &sp_component.Source, &sp_component.Spacing_unit_id, &sp_component.Sp_component_type, &sp_component.Store_id, &sp_component.Store_structure_obs_no, &sp_component.Strat_name_set_id, &sp_component.Strat_unit_id, &sp_component.Support_facility_id, &sp_component.Support_facility_subtype, &sp_component.Sw_application_id, &sp_component.Thesaurus_id, &sp_component.Thesaurus_word_id, &sp_component.Uwi, &sp_component.Well_activity_set_id, &sp_component.Well_activity_source, &sp_component.Well_activity_type_id, &sp_component.Well_set_id, &sp_component.Work_order_id, &sp_component.Row_changed_by, &sp_component.Row_changed_date, &sp_component.Row_created_by, &sp_component.Row_created_date, &sp_component.Row_effective_date, &sp_component.Row_expiry_date, &sp_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_component to result
		result = append(result, sp_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpComponent(c *fiber.Ctx) error {
	var sp_component dto.Sp_component

	setDefaults(&sp_component)

	if err := json.Unmarshal(c.Body(), &sp_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103, :104, :105, :106, :107, :108, :109, :110, :111, :112, :113, :114, :115, :116, :117, :118, :119, :120, :121, :122, :123)")
	if err != nil {
		return err
	}
	sp_component.Row_created_date = formatDate(sp_component.Row_created_date)
	sp_component.Row_changed_date = formatDate(sp_component.Row_changed_date)
	sp_component.Effective_date = formatDateString(sp_component.Effective_date)
	sp_component.Expiry_date = formatDateString(sp_component.Expiry_date)
	sp_component.Row_effective_date = formatDateString(sp_component.Row_effective_date)
	sp_component.Row_expiry_date = formatDateString(sp_component.Row_expiry_date)






	rows, err := stmt.Exec(sp_component.Spatial_description_id, sp_component.Spatial_obs_no, sp_component.Component_id, sp_component.Active_ind, sp_component.Activity_obs_no, sp_component.Analysis_id, sp_component.Application_id, sp_component.Area_id, sp_component.Area_type, sp_component.Authority_id, sp_component.Ba_address_obs_no, sp_component.Ba_address_source, sp_component.Ba_licensee_ba_id, sp_component.Ba_license_id, sp_component.Ba_organization_id, sp_component.Ba_organization_seq_no, sp_component.Business_associate_id, sp_component.Catalogue_additive_id, sp_component.Catalogue_equip_id, sp_component.Classification_system_id, sp_component.Class_level_id, sp_component.Consent_id, sp_component.Consult_id, sp_component.Contest_id, sp_component.Contract_id, sp_component.Ecozone_id, sp_component.Effective_date, sp_component.Employee_ba_id, sp_component.Employee_obs_no, sp_component.Employer_ba_id, sp_component.Entitlement_id, sp_component.Equipment_id, sp_component.Expiry_date, sp_component.Facility_id, sp_component.Facility_type, sp_component.Field_id, sp_component.Field_node_id, sp_component.Field_node_source, sp_component.Field_station_id, sp_component.Finance_id, sp_component.Fossil_id, sp_component.Hse_incident_id, sp_component.Incident_set_id, sp_component.Incident_type, sp_component.Information_item_id, sp_component.Info_item_subtype, sp_component.Instrument_id, sp_component.Interest_set_id, sp_component.Interest_set_seq_no, sp_component.Jurisdiction, sp_component.Land_request_id, sp_component.Land_right_id, sp_component.Land_right_status, sp_component.Land_right_status_seq_no, sp_component.Land_right_status_type, sp_component.Land_right_subtype, sp_component.Land_sale_number, sp_component.Land_sale_offering_id, sp_component.Language, sp_component.Lithology_log_id, sp_component.Lith_log_source, sp_component.Lith_sample_id, sp_component.Node_id, sp_component.Notification_id, sp_component.Obligation_id, sp_component.Obligation_seq_no, sp_component.Paleo_summary_id, sp_component.Pden_id, sp_component.Pden_source, sp_component.Pden_subtype, sp_component.Physical_item_id, sp_component.Pool_id, sp_component.Ppdm_guid, sp_component.Ppdm_system_id, sp_component.Ppdm_table_name, sp_component.Product_type, sp_component.Prod_string_id, sp_component.Prod_string_source, sp_component.Project_id, sp_component.Rate_schedule_id, sp_component.Referenced_guid, sp_component.Referenced_system_id, sp_component.Referenced_table_name, sp_component.Remark, sp_component.Resent_id, sp_component.Resent_product_type, sp_component.Resent_reserve_class_id, sp_component.Reserve_class_id, sp_component.Restriction_id, sp_component.Restriction_version, sp_component.Sample_anal_source, sp_component.Seis_bin_grid_id, sp_component.Seis_bin_source, sp_component.Seis_set_id, sp_component.Seis_set_subtype, sp_component.Sf_subtype, sp_component.Size_seq_no, sp_component.Size_type, sp_component.Source, sp_component.Spacing_unit_id, sp_component.Sp_component_type, sp_component.Store_id, sp_component.Store_structure_obs_no, sp_component.Strat_name_set_id, sp_component.Strat_unit_id, sp_component.Support_facility_id, sp_component.Support_facility_subtype, sp_component.Sw_application_id, sp_component.Thesaurus_id, sp_component.Thesaurus_word_id, sp_component.Uwi, sp_component.Well_activity_set_id, sp_component.Well_activity_source, sp_component.Well_activity_type_id, sp_component.Well_set_id, sp_component.Work_order_id, sp_component.Row_changed_by, sp_component.Row_changed_date, sp_component.Row_created_by, sp_component.Row_created_date, sp_component.Row_effective_date, sp_component.Row_expiry_date, sp_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpComponent(c *fiber.Ctx) error {
	var sp_component dto.Sp_component

	setDefaults(&sp_component)

	if err := json.Unmarshal(c.Body(), &sp_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_component.Spatial_description_id = id

    if sp_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_component set spatial_obs_no = :1, component_id = :2, active_ind = :3, activity_obs_no = :4, analysis_id = :5, application_id = :6, area_id = :7, area_type = :8, authority_id = :9, ba_address_obs_no = :10, ba_address_source = :11, ba_licensee_ba_id = :12, ba_license_id = :13, ba_organization_id = :14, ba_organization_seq_no = :15, business_associate_id = :16, catalogue_additive_id = :17, catalogue_equip_id = :18, classification_system_id = :19, class_level_id = :20, consent_id = :21, consult_id = :22, contest_id = :23, contract_id = :24, ecozone_id = :25, effective_date = :26, employee_ba_id = :27, employee_obs_no = :28, employer_ba_id = :29, entitlement_id = :30, equipment_id = :31, expiry_date = :32, facility_id = :33, facility_type = :34, field_id = :35, field_node_id = :36, field_node_source = :37, field_station_id = :38, finance_id = :39, fossil_id = :40, hse_incident_id = :41, incident_set_id = :42, incident_type = :43, information_item_id = :44, info_item_subtype = :45, instrument_id = :46, interest_set_id = :47, interest_set_seq_no = :48, jurisdiction = :49, land_request_id = :50, land_right_id = :51, land_right_status = :52, land_right_status_seq_no = :53, land_right_status_type = :54, land_right_subtype = :55, land_sale_number = :56, land_sale_offering_id = :57, language = :58, lithology_log_id = :59, lith_log_source = :60, lith_sample_id = :61, node_id = :62, notification_id = :63, obligation_id = :64, obligation_seq_no = :65, paleo_summary_id = :66, pden_id = :67, pden_source = :68, pden_subtype = :69, physical_item_id = :70, pool_id = :71, ppdm_guid = :72, ppdm_system_id = :73, ppdm_table_name = :74, product_type = :75, prod_string_id = :76, prod_string_source = :77, project_id = :78, rate_schedule_id = :79, referenced_guid = :80, referenced_system_id = :81, referenced_table_name = :82, remark = :83, resent_id = :84, resent_product_type = :85, resent_reserve_class_id = :86, reserve_class_id = :87, restriction_id = :88, restriction_version = :89, sample_anal_source = :90, seis_bin_grid_id = :91, seis_bin_source = :92, seis_set_id = :93, seis_set_subtype = :94, sf_subtype = :95, size_seq_no = :96, size_type = :97, source = :98, spacing_unit_id = :99, sp_component_type = :100, store_id = :101, store_structure_obs_no = :102, strat_name_set_id = :103, strat_unit_id = :104, support_facility_id = :105, support_facility_subtype = :106, sw_application_id = :107, thesaurus_id = :108, thesaurus_word_id = :109, uwi = :110, well_activity_set_id = :111, well_activity_source = :112, well_activity_type_id = :113, well_set_id = :114, work_order_id = :115, row_changed_by = :116, row_changed_date = :117, row_created_by = :118, row_effective_date = :119, row_expiry_date = :120, row_quality = :121 where spatial_description_id = :123")
	if err != nil {
		return err
	}

	sp_component.Row_changed_date = formatDate(sp_component.Row_changed_date)
	sp_component.Effective_date = formatDateString(sp_component.Effective_date)
	sp_component.Expiry_date = formatDateString(sp_component.Expiry_date)
	sp_component.Row_effective_date = formatDateString(sp_component.Row_effective_date)
	sp_component.Row_expiry_date = formatDateString(sp_component.Row_expiry_date)






	rows, err := stmt.Exec(sp_component.Spatial_obs_no, sp_component.Component_id, sp_component.Active_ind, sp_component.Activity_obs_no, sp_component.Analysis_id, sp_component.Application_id, sp_component.Area_id, sp_component.Area_type, sp_component.Authority_id, sp_component.Ba_address_obs_no, sp_component.Ba_address_source, sp_component.Ba_licensee_ba_id, sp_component.Ba_license_id, sp_component.Ba_organization_id, sp_component.Ba_organization_seq_no, sp_component.Business_associate_id, sp_component.Catalogue_additive_id, sp_component.Catalogue_equip_id, sp_component.Classification_system_id, sp_component.Class_level_id, sp_component.Consent_id, sp_component.Consult_id, sp_component.Contest_id, sp_component.Contract_id, sp_component.Ecozone_id, sp_component.Effective_date, sp_component.Employee_ba_id, sp_component.Employee_obs_no, sp_component.Employer_ba_id, sp_component.Entitlement_id, sp_component.Equipment_id, sp_component.Expiry_date, sp_component.Facility_id, sp_component.Facility_type, sp_component.Field_id, sp_component.Field_node_id, sp_component.Field_node_source, sp_component.Field_station_id, sp_component.Finance_id, sp_component.Fossil_id, sp_component.Hse_incident_id, sp_component.Incident_set_id, sp_component.Incident_type, sp_component.Information_item_id, sp_component.Info_item_subtype, sp_component.Instrument_id, sp_component.Interest_set_id, sp_component.Interest_set_seq_no, sp_component.Jurisdiction, sp_component.Land_request_id, sp_component.Land_right_id, sp_component.Land_right_status, sp_component.Land_right_status_seq_no, sp_component.Land_right_status_type, sp_component.Land_right_subtype, sp_component.Land_sale_number, sp_component.Land_sale_offering_id, sp_component.Language, sp_component.Lithology_log_id, sp_component.Lith_log_source, sp_component.Lith_sample_id, sp_component.Node_id, sp_component.Notification_id, sp_component.Obligation_id, sp_component.Obligation_seq_no, sp_component.Paleo_summary_id, sp_component.Pden_id, sp_component.Pden_source, sp_component.Pden_subtype, sp_component.Physical_item_id, sp_component.Pool_id, sp_component.Ppdm_guid, sp_component.Ppdm_system_id, sp_component.Ppdm_table_name, sp_component.Product_type, sp_component.Prod_string_id, sp_component.Prod_string_source, sp_component.Project_id, sp_component.Rate_schedule_id, sp_component.Referenced_guid, sp_component.Referenced_system_id, sp_component.Referenced_table_name, sp_component.Remark, sp_component.Resent_id, sp_component.Resent_product_type, sp_component.Resent_reserve_class_id, sp_component.Reserve_class_id, sp_component.Restriction_id, sp_component.Restriction_version, sp_component.Sample_anal_source, sp_component.Seis_bin_grid_id, sp_component.Seis_bin_source, sp_component.Seis_set_id, sp_component.Seis_set_subtype, sp_component.Sf_subtype, sp_component.Size_seq_no, sp_component.Size_type, sp_component.Source, sp_component.Spacing_unit_id, sp_component.Sp_component_type, sp_component.Store_id, sp_component.Store_structure_obs_no, sp_component.Strat_name_set_id, sp_component.Strat_unit_id, sp_component.Support_facility_id, sp_component.Support_facility_subtype, sp_component.Sw_application_id, sp_component.Thesaurus_id, sp_component.Thesaurus_word_id, sp_component.Uwi, sp_component.Well_activity_set_id, sp_component.Well_activity_source, sp_component.Well_activity_type_id, sp_component.Well_set_id, sp_component.Work_order_id, sp_component.Row_changed_by, sp_component.Row_changed_date, sp_component.Row_created_by, sp_component.Row_effective_date, sp_component.Row_expiry_date, sp_component.Row_quality, sp_component.Spatial_description_id)
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

func PatchSpComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_component set "
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
	query += " where spatial_description_id = :id"

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

func DeleteSpComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_component dto.Sp_component
	sp_component.Spatial_description_id = id

	stmt, err := db.Prepare("delete from sp_component where spatial_description_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_component.Spatial_description_id)
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


