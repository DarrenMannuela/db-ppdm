package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFinComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from fin_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Fin_component

	for rows.Next() {
		var fin_component dto.Fin_component
		if err := rows.Scan(&fin_component.Finance_id, &fin_component.Finance_component_id, &fin_component.Acquisition_cost_id, &fin_component.Active_ind, &fin_component.Activity_obs_no, &fin_component.Analysis_id, &fin_component.Application_id, &fin_component.Area_id, &fin_component.Area_type, &fin_component.Authority_id, &fin_component.Ba_licensee_ba_id, &fin_component.Ba_license_id, &fin_component.Ba_organization_id, &fin_component.Ba_organization_seq_no, &fin_component.Business_associate_id, &fin_component.Catalogue_additive_id, &fin_component.Catalogue_equip_id, &fin_component.Classification_system_id, &fin_component.Class_level_id, &fin_component.Consent_id, &fin_component.Consult_id, &fin_component.Contest_id, &fin_component.Contract_id, &fin_component.Ecozone_id, &fin_component.Effective_date, &fin_component.Employee_ba_id, &fin_component.Employee_obs_no, &fin_component.Employer_ba_id, &fin_component.Entitlement_id, &fin_component.Equipment_id, &fin_component.Equip_ba_obs_no, &fin_component.Equip_maint_id, &fin_component.Expiry_date, &fin_component.Facility_id, &fin_component.Facility_license_id, &fin_component.Facility_type, &fin_component.Field_id, &fin_component.Field_station_id, &fin_component.Finance_id2, &fin_component.Fin_component_table, &fin_component.Fin_component_type, &fin_component.Fossil_id, &fin_component.Incident_id, &fin_component.Incident_set_id, &fin_component.Incident_type, &fin_component.Information_item_id, &fin_component.Info_item_subtype, &fin_component.Instrument_id, &fin_component.Interest_set_id, &fin_component.Interest_set_seq_no, &fin_component.Land_right_id, &fin_component.Land_right_subtype, &fin_component.Land_sale_bid_set_id, &fin_component.Land_sale_jurisdiction, &fin_component.Land_sale_number, &fin_component.Land_sale_offering_bid, &fin_component.Land_sale_offering_id, &fin_component.Language, &fin_component.Lithology_log_id, &fin_component.Lith_log_source, &fin_component.Notification_id, &fin_component.Obligation_id, &fin_component.Obligation_seq_no, &fin_component.Paleo_summary_id, &fin_component.Pden_id, &fin_component.Pden_source, &fin_component.Pden_subtype, &fin_component.Physical_item_id, &fin_component.Pool_id, &fin_component.Ppdm_guid, &fin_component.Ppdm_system_id, &fin_component.Ppdm_table_name, &fin_component.Product_type, &fin_component.Prod_string_id, &fin_component.Prod_string_source, &fin_component.Project_id, &fin_component.Pr_str_form_obs_no, &fin_component.Rate_schedule_id, &fin_component.Referenced_guid, &fin_component.Referenced_system_id, &fin_component.Referenced_table_name, &fin_component.Remark, &fin_component.Resent_id, &fin_component.Reserve_class_id, &fin_component.Restriction_id, &fin_component.Restriction_version, &fin_component.Revision_id, &fin_component.Revision_obs_no, &fin_component.Sample_anal_source, &fin_component.Seis_license_id, &fin_component.Seis_set_id, &fin_component.Seis_set_subtype, &fin_component.Seis_transaction_id, &fin_component.Seis_transaction_type, &fin_component.Sf_subtype, &fin_component.Source, &fin_component.Spatial_description_id, &fin_component.Spatial_obs_no, &fin_component.Store_id, &fin_component.Store_structure_obs_no, &fin_component.Strat_name_set_id, &fin_component.Strat_unit_id, &fin_component.Support_facility_id, &fin_component.Sw_application_id, &fin_component.Thesaurus_id, &fin_component.Thesaurus_word_id, &fin_component.Uwi, &fin_component.Well_activity_set_id, &fin_component.Well_activity_source, &fin_component.Well_activity_type_id, &fin_component.Well_drill_period_obs_no, &fin_component.Well_license_id, &fin_component.Well_license_source, &fin_component.Well_service_provided_by, &fin_component.Well_service_seg_obs_no, &fin_component.Well_service_seq_no, &fin_component.Well_set_id, &fin_component.Work_order_id, &fin_component.Row_changed_by, &fin_component.Row_changed_date, &fin_component.Row_created_by, &fin_component.Row_created_date, &fin_component.Row_effective_date, &fin_component.Row_expiry_date, &fin_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append fin_component to result
		result = append(result, fin_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFinComponent(c *fiber.Ctx) error {
	var fin_component dto.Fin_component

	setDefaults(&fin_component)

	if err := json.Unmarshal(c.Body(), &fin_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into fin_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103, :104, :105, :106, :107, :108, :109, :110, :111, :112, :113, :114, :115, :116, :117, :118, :119, :120, :121, :122, :123, :124, :125)")
	if err != nil {
		return err
	}
	fin_component.Row_created_date = formatDate(fin_component.Row_created_date)
	fin_component.Row_changed_date = formatDate(fin_component.Row_changed_date)
	fin_component.Effective_date = formatDateString(fin_component.Effective_date)
	fin_component.Expiry_date = formatDateString(fin_component.Expiry_date)
	fin_component.Row_effective_date = formatDateString(fin_component.Row_effective_date)
	fin_component.Row_expiry_date = formatDateString(fin_component.Row_expiry_date)






	rows, err := stmt.Exec(fin_component.Finance_id, fin_component.Finance_component_id, fin_component.Acquisition_cost_id, fin_component.Active_ind, fin_component.Activity_obs_no, fin_component.Analysis_id, fin_component.Application_id, fin_component.Area_id, fin_component.Area_type, fin_component.Authority_id, fin_component.Ba_licensee_ba_id, fin_component.Ba_license_id, fin_component.Ba_organization_id, fin_component.Ba_organization_seq_no, fin_component.Business_associate_id, fin_component.Catalogue_additive_id, fin_component.Catalogue_equip_id, fin_component.Classification_system_id, fin_component.Class_level_id, fin_component.Consent_id, fin_component.Consult_id, fin_component.Contest_id, fin_component.Contract_id, fin_component.Ecozone_id, fin_component.Effective_date, fin_component.Employee_ba_id, fin_component.Employee_obs_no, fin_component.Employer_ba_id, fin_component.Entitlement_id, fin_component.Equipment_id, fin_component.Equip_ba_obs_no, fin_component.Equip_maint_id, fin_component.Expiry_date, fin_component.Facility_id, fin_component.Facility_license_id, fin_component.Facility_type, fin_component.Field_id, fin_component.Field_station_id, fin_component.Finance_id2, fin_component.Fin_component_table, fin_component.Fin_component_type, fin_component.Fossil_id, fin_component.Incident_id, fin_component.Incident_set_id, fin_component.Incident_type, fin_component.Information_item_id, fin_component.Info_item_subtype, fin_component.Instrument_id, fin_component.Interest_set_id, fin_component.Interest_set_seq_no, fin_component.Land_right_id, fin_component.Land_right_subtype, fin_component.Land_sale_bid_set_id, fin_component.Land_sale_jurisdiction, fin_component.Land_sale_number, fin_component.Land_sale_offering_bid, fin_component.Land_sale_offering_id, fin_component.Language, fin_component.Lithology_log_id, fin_component.Lith_log_source, fin_component.Notification_id, fin_component.Obligation_id, fin_component.Obligation_seq_no, fin_component.Paleo_summary_id, fin_component.Pden_id, fin_component.Pden_source, fin_component.Pden_subtype, fin_component.Physical_item_id, fin_component.Pool_id, fin_component.Ppdm_guid, fin_component.Ppdm_system_id, fin_component.Ppdm_table_name, fin_component.Product_type, fin_component.Prod_string_id, fin_component.Prod_string_source, fin_component.Project_id, fin_component.Pr_str_form_obs_no, fin_component.Rate_schedule_id, fin_component.Referenced_guid, fin_component.Referenced_system_id, fin_component.Referenced_table_name, fin_component.Remark, fin_component.Resent_id, fin_component.Reserve_class_id, fin_component.Restriction_id, fin_component.Restriction_version, fin_component.Revision_id, fin_component.Revision_obs_no, fin_component.Sample_anal_source, fin_component.Seis_license_id, fin_component.Seis_set_id, fin_component.Seis_set_subtype, fin_component.Seis_transaction_id, fin_component.Seis_transaction_type, fin_component.Sf_subtype, fin_component.Source, fin_component.Spatial_description_id, fin_component.Spatial_obs_no, fin_component.Store_id, fin_component.Store_structure_obs_no, fin_component.Strat_name_set_id, fin_component.Strat_unit_id, fin_component.Support_facility_id, fin_component.Sw_application_id, fin_component.Thesaurus_id, fin_component.Thesaurus_word_id, fin_component.Uwi, fin_component.Well_activity_set_id, fin_component.Well_activity_source, fin_component.Well_activity_type_id, fin_component.Well_drill_period_obs_no, fin_component.Well_license_id, fin_component.Well_license_source, fin_component.Well_service_provided_by, fin_component.Well_service_seg_obs_no, fin_component.Well_service_seq_no, fin_component.Well_set_id, fin_component.Work_order_id, fin_component.Row_changed_by, fin_component.Row_changed_date, fin_component.Row_created_by, fin_component.Row_created_date, fin_component.Row_effective_date, fin_component.Row_expiry_date, fin_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFinComponent(c *fiber.Ctx) error {
	var fin_component dto.Fin_component

	setDefaults(&fin_component)

	if err := json.Unmarshal(c.Body(), &fin_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	fin_component.Finance_id = id

    if fin_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update fin_component set finance_component_id = :1, acquisition_cost_id = :2, active_ind = :3, activity_obs_no = :4, analysis_id = :5, application_id = :6, area_id = :7, area_type = :8, authority_id = :9, ba_licensee_ba_id = :10, ba_license_id = :11, ba_organization_id = :12, ba_organization_seq_no = :13, business_associate_id = :14, catalogue_additive_id = :15, catalogue_equip_id = :16, classification_system_id = :17, class_level_id = :18, consent_id = :19, consult_id = :20, contest_id = :21, contract_id = :22, ecozone_id = :23, effective_date = :24, employee_ba_id = :25, employee_obs_no = :26, employer_ba_id = :27, entitlement_id = :28, equipment_id = :29, equip_ba_obs_no = :30, equip_maint_id = :31, expiry_date = :32, facility_id = :33, facility_license_id = :34, facility_type = :35, field_id = :36, field_station_id = :37, finance_id2 = :38, fin_component_table = :39, fin_component_type = :40, fossil_id = :41, incident_id = :42, incident_set_id = :43, incident_type = :44, information_item_id = :45, info_item_subtype = :46, instrument_id = :47, interest_set_id = :48, interest_set_seq_no = :49, land_right_id = :50, land_right_subtype = :51, land_sale_bid_set_id = :52, land_sale_jurisdiction = :53, land_sale_number = :54, land_sale_offering_bid = :55, land_sale_offering_id = :56, language = :57, lithology_log_id = :58, lith_log_source = :59, notification_id = :60, obligation_id = :61, obligation_seq_no = :62, paleo_summary_id = :63, pden_id = :64, pden_source = :65, pden_subtype = :66, physical_item_id = :67, pool_id = :68, ppdm_guid = :69, ppdm_system_id = :70, ppdm_table_name = :71, product_type = :72, prod_string_id = :73, prod_string_source = :74, project_id = :75, pr_str_form_obs_no = :76, rate_schedule_id = :77, referenced_guid = :78, referenced_system_id = :79, referenced_table_name = :80, remark = :81, resent_id = :82, reserve_class_id = :83, restriction_id = :84, restriction_version = :85, revision_id = :86, revision_obs_no = :87, sample_anal_source = :88, seis_license_id = :89, seis_set_id = :90, seis_set_subtype = :91, seis_transaction_id = :92, seis_transaction_type = :93, sf_subtype = :94, source = :95, spatial_description_id = :96, spatial_obs_no = :97, store_id = :98, store_structure_obs_no = :99, strat_name_set_id = :100, strat_unit_id = :101, support_facility_id = :102, sw_application_id = :103, thesaurus_id = :104, thesaurus_word_id = :105, uwi = :106, well_activity_set_id = :107, well_activity_source = :108, well_activity_type_id = :109, well_drill_period_obs_no = :110, well_license_id = :111, well_license_source = :112, well_service_provided_by = :113, well_service_seg_obs_no = :114, well_service_seq_no = :115, well_set_id = :116, work_order_id = :117, row_changed_by = :118, row_changed_date = :119, row_created_by = :120, row_effective_date = :121, row_expiry_date = :122, row_quality = :123 where finance_id = :125")
	if err != nil {
		return err
	}

	fin_component.Row_changed_date = formatDate(fin_component.Row_changed_date)
	fin_component.Effective_date = formatDateString(fin_component.Effective_date)
	fin_component.Expiry_date = formatDateString(fin_component.Expiry_date)
	fin_component.Row_effective_date = formatDateString(fin_component.Row_effective_date)
	fin_component.Row_expiry_date = formatDateString(fin_component.Row_expiry_date)






	rows, err := stmt.Exec(fin_component.Finance_component_id, fin_component.Acquisition_cost_id, fin_component.Active_ind, fin_component.Activity_obs_no, fin_component.Analysis_id, fin_component.Application_id, fin_component.Area_id, fin_component.Area_type, fin_component.Authority_id, fin_component.Ba_licensee_ba_id, fin_component.Ba_license_id, fin_component.Ba_organization_id, fin_component.Ba_organization_seq_no, fin_component.Business_associate_id, fin_component.Catalogue_additive_id, fin_component.Catalogue_equip_id, fin_component.Classification_system_id, fin_component.Class_level_id, fin_component.Consent_id, fin_component.Consult_id, fin_component.Contest_id, fin_component.Contract_id, fin_component.Ecozone_id, fin_component.Effective_date, fin_component.Employee_ba_id, fin_component.Employee_obs_no, fin_component.Employer_ba_id, fin_component.Entitlement_id, fin_component.Equipment_id, fin_component.Equip_ba_obs_no, fin_component.Equip_maint_id, fin_component.Expiry_date, fin_component.Facility_id, fin_component.Facility_license_id, fin_component.Facility_type, fin_component.Field_id, fin_component.Field_station_id, fin_component.Finance_id2, fin_component.Fin_component_table, fin_component.Fin_component_type, fin_component.Fossil_id, fin_component.Incident_id, fin_component.Incident_set_id, fin_component.Incident_type, fin_component.Information_item_id, fin_component.Info_item_subtype, fin_component.Instrument_id, fin_component.Interest_set_id, fin_component.Interest_set_seq_no, fin_component.Land_right_id, fin_component.Land_right_subtype, fin_component.Land_sale_bid_set_id, fin_component.Land_sale_jurisdiction, fin_component.Land_sale_number, fin_component.Land_sale_offering_bid, fin_component.Land_sale_offering_id, fin_component.Language, fin_component.Lithology_log_id, fin_component.Lith_log_source, fin_component.Notification_id, fin_component.Obligation_id, fin_component.Obligation_seq_no, fin_component.Paleo_summary_id, fin_component.Pden_id, fin_component.Pden_source, fin_component.Pden_subtype, fin_component.Physical_item_id, fin_component.Pool_id, fin_component.Ppdm_guid, fin_component.Ppdm_system_id, fin_component.Ppdm_table_name, fin_component.Product_type, fin_component.Prod_string_id, fin_component.Prod_string_source, fin_component.Project_id, fin_component.Pr_str_form_obs_no, fin_component.Rate_schedule_id, fin_component.Referenced_guid, fin_component.Referenced_system_id, fin_component.Referenced_table_name, fin_component.Remark, fin_component.Resent_id, fin_component.Reserve_class_id, fin_component.Restriction_id, fin_component.Restriction_version, fin_component.Revision_id, fin_component.Revision_obs_no, fin_component.Sample_anal_source, fin_component.Seis_license_id, fin_component.Seis_set_id, fin_component.Seis_set_subtype, fin_component.Seis_transaction_id, fin_component.Seis_transaction_type, fin_component.Sf_subtype, fin_component.Source, fin_component.Spatial_description_id, fin_component.Spatial_obs_no, fin_component.Store_id, fin_component.Store_structure_obs_no, fin_component.Strat_name_set_id, fin_component.Strat_unit_id, fin_component.Support_facility_id, fin_component.Sw_application_id, fin_component.Thesaurus_id, fin_component.Thesaurus_word_id, fin_component.Uwi, fin_component.Well_activity_set_id, fin_component.Well_activity_source, fin_component.Well_activity_type_id, fin_component.Well_drill_period_obs_no, fin_component.Well_license_id, fin_component.Well_license_source, fin_component.Well_service_provided_by, fin_component.Well_service_seg_obs_no, fin_component.Well_service_seq_no, fin_component.Well_set_id, fin_component.Work_order_id, fin_component.Row_changed_by, fin_component.Row_changed_date, fin_component.Row_created_by, fin_component.Row_effective_date, fin_component.Row_expiry_date, fin_component.Row_quality, fin_component.Finance_id)
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

func PatchFinComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update fin_component set "
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
	query += " where finance_id = :id"

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

func DeleteFinComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var fin_component dto.Fin_component
	fin_component.Finance_id = id

	stmt, err := db.Prepare("delete from fin_component where finance_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(fin_component.Finance_id)
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


