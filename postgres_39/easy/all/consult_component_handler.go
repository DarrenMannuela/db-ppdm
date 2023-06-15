package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetConsultComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from consult_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Consult_component

	for rows.Next() {
		var consult_component dto.Consult_component
		if err := rows.Scan(&consult_component.Consult_id, &consult_component.Component_id, &consult_component.Active_ind, &consult_component.Activity_obs_no, &consult_component.Analysis_id, &consult_component.Application_id, &consult_component.Area_id, &consult_component.Area_type, &consult_component.Authority_id, &consult_component.Ba_licensee_ba_id, &consult_component.Ba_license_id, &consult_component.Ba_organization_id, &consult_component.Ba_organization_seq_no, &consult_component.Business_associate_id, &consult_component.Catalogue_additive_id, &consult_component.Catalogue_equip_id, &consult_component.Classification_system_id, &consult_component.Class_level_id, &consult_component.Component_type, &consult_component.Consent_id, &consult_component.Contest_id, &consult_component.Contract_id, &consult_component.Currency_conversion, &consult_component.Currency_ouom, &consult_component.Ecozone_id, &consult_component.Effective_date, &consult_component.Employee_ba_id, &consult_component.Employee_obs_no, &consult_component.Employer_ba_id, &consult_component.Entitlement_id, &consult_component.Equipment_id, &consult_component.Expiry_date, &consult_component.Facility_id, &consult_component.Facility_license_id, &consult_component.Facility_type, &consult_component.Field_id, &consult_component.Field_station_id, &consult_component.Finance_id, &consult_component.Fossil_id, &consult_component.Hse_incident_id, &consult_component.Incident_set_id, &consult_component.Incident_type, &consult_component.Information_item_id, &consult_component.Info_item_subtype, &consult_component.Instrument_id, &consult_component.Interest_set_id, &consult_component.Interest_set_seq_no, &consult_component.Land_right_id, &consult_component.Land_right_subtype, &consult_component.Land_sale_jurisdiction, &consult_component.Land_sale_number, &consult_component.Lithology_log_id, &consult_component.Lith_log_source, &consult_component.Notification_id, &consult_component.Obligation_id, &consult_component.Obligation_seq_no, &consult_component.Paleo_summary_id, &consult_component.Pden_id, &consult_component.Pden_source, &consult_component.Pden_subtype, &consult_component.Physical_item_id, &consult_component.Platform_id, &consult_component.Platform_source, &consult_component.Pool_id, &consult_component.Ppdm_guid, &consult_component.Ppdm_system_id, &consult_component.Ppdm_table_name, &consult_component.Product_type, &consult_component.Prod_string_id, &consult_component.Prod_string_source, &consult_component.Project_id, &consult_component.Provision_id, &consult_component.Rate_schedule_id, &consult_component.Referenced_guid, &consult_component.Referenced_system_id, &consult_component.Referenced_table_name, &consult_component.Remark, &consult_component.Resent_id, &consult_component.Reserve_class_id, &consult_component.Restriction_id, &consult_component.Restriction_version, &consult_component.Sample_anal_source, &consult_component.Seis_license_id, &consult_component.Seis_set_id, &consult_component.Seis_set_subtype, &consult_component.Sf_subtype, &consult_component.Source, &consult_component.Spatial_description_id, &consult_component.Spatial_obs_no, &consult_component.Store_id, &consult_component.Store_structure_obs_no, &consult_component.Strat_name_set_id, &consult_component.Strat_unit_id, &consult_component.Support_facility_id, &consult_component.Sw_application_id, &consult_component.Thesaurus_id, &consult_component.Thesaurus_word_id, &consult_component.Uwi, &consult_component.Well_activity_set_id, &consult_component.Well_activity_source, &consult_component.Well_activity_type_id, &consult_component.Well_license_id, &consult_component.Well_license_source, &consult_component.Well_set_id, &consult_component.Work_order_id, &consult_component.Row_changed_by, &consult_component.Row_changed_date, &consult_component.Row_created_by, &consult_component.Row_created_date, &consult_component.Row_effective_date, &consult_component.Row_expiry_date, &consult_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append consult_component to result
		result = append(result, consult_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetConsultComponent(c *fiber.Ctx) error {
	var consult_component dto.Consult_component

	setDefaults(&consult_component)

	if err := json.Unmarshal(c.Body(), &consult_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into consult_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103, :104, :105, :106, :107, :108, :109, :110, :111, :112)")
	if err != nil {
		return err
	}
	consult_component.Row_created_date = formatDate(consult_component.Row_created_date)
	consult_component.Row_changed_date = formatDate(consult_component.Row_changed_date)
	consult_component.Effective_date = formatDateString(consult_component.Effective_date)
	consult_component.Expiry_date = formatDateString(consult_component.Expiry_date)
	consult_component.Row_effective_date = formatDateString(consult_component.Row_effective_date)
	consult_component.Row_expiry_date = formatDateString(consult_component.Row_expiry_date)






	rows, err := stmt.Exec(consult_component.Consult_id, consult_component.Component_id, consult_component.Active_ind, consult_component.Activity_obs_no, consult_component.Analysis_id, consult_component.Application_id, consult_component.Area_id, consult_component.Area_type, consult_component.Authority_id, consult_component.Ba_licensee_ba_id, consult_component.Ba_license_id, consult_component.Ba_organization_id, consult_component.Ba_organization_seq_no, consult_component.Business_associate_id, consult_component.Catalogue_additive_id, consult_component.Catalogue_equip_id, consult_component.Classification_system_id, consult_component.Class_level_id, consult_component.Component_type, consult_component.Consent_id, consult_component.Contest_id, consult_component.Contract_id, consult_component.Currency_conversion, consult_component.Currency_ouom, consult_component.Ecozone_id, consult_component.Effective_date, consult_component.Employee_ba_id, consult_component.Employee_obs_no, consult_component.Employer_ba_id, consult_component.Entitlement_id, consult_component.Equipment_id, consult_component.Expiry_date, consult_component.Facility_id, consult_component.Facility_license_id, consult_component.Facility_type, consult_component.Field_id, consult_component.Field_station_id, consult_component.Finance_id, consult_component.Fossil_id, consult_component.Hse_incident_id, consult_component.Incident_set_id, consult_component.Incident_type, consult_component.Information_item_id, consult_component.Info_item_subtype, consult_component.Instrument_id, consult_component.Interest_set_id, consult_component.Interest_set_seq_no, consult_component.Land_right_id, consult_component.Land_right_subtype, consult_component.Land_sale_jurisdiction, consult_component.Land_sale_number, consult_component.Lithology_log_id, consult_component.Lith_log_source, consult_component.Notification_id, consult_component.Obligation_id, consult_component.Obligation_seq_no, consult_component.Paleo_summary_id, consult_component.Pden_id, consult_component.Pden_source, consult_component.Pden_subtype, consult_component.Physical_item_id, consult_component.Platform_id, consult_component.Platform_source, consult_component.Pool_id, consult_component.Ppdm_guid, consult_component.Ppdm_system_id, consult_component.Ppdm_table_name, consult_component.Product_type, consult_component.Prod_string_id, consult_component.Prod_string_source, consult_component.Project_id, consult_component.Provision_id, consult_component.Rate_schedule_id, consult_component.Referenced_guid, consult_component.Referenced_system_id, consult_component.Referenced_table_name, consult_component.Remark, consult_component.Resent_id, consult_component.Reserve_class_id, consult_component.Restriction_id, consult_component.Restriction_version, consult_component.Sample_anal_source, consult_component.Seis_license_id, consult_component.Seis_set_id, consult_component.Seis_set_subtype, consult_component.Sf_subtype, consult_component.Source, consult_component.Spatial_description_id, consult_component.Spatial_obs_no, consult_component.Store_id, consult_component.Store_structure_obs_no, consult_component.Strat_name_set_id, consult_component.Strat_unit_id, consult_component.Support_facility_id, consult_component.Sw_application_id, consult_component.Thesaurus_id, consult_component.Thesaurus_word_id, consult_component.Uwi, consult_component.Well_activity_set_id, consult_component.Well_activity_source, consult_component.Well_activity_type_id, consult_component.Well_license_id, consult_component.Well_license_source, consult_component.Well_set_id, consult_component.Work_order_id, consult_component.Row_changed_by, consult_component.Row_changed_date, consult_component.Row_created_by, consult_component.Row_created_date, consult_component.Row_effective_date, consult_component.Row_expiry_date, consult_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateConsultComponent(c *fiber.Ctx) error {
	var consult_component dto.Consult_component

	setDefaults(&consult_component)

	if err := json.Unmarshal(c.Body(), &consult_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	consult_component.Consult_id = id

    if consult_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update consult_component set component_id = :1, active_ind = :2, activity_obs_no = :3, analysis_id = :4, application_id = :5, area_id = :6, area_type = :7, authority_id = :8, ba_licensee_ba_id = :9, ba_license_id = :10, ba_organization_id = :11, ba_organization_seq_no = :12, business_associate_id = :13, catalogue_additive_id = :14, catalogue_equip_id = :15, classification_system_id = :16, class_level_id = :17, component_type = :18, consent_id = :19, contest_id = :20, contract_id = :21, currency_conversion = :22, currency_ouom = :23, ecozone_id = :24, effective_date = :25, employee_ba_id = :26, employee_obs_no = :27, employer_ba_id = :28, entitlement_id = :29, equipment_id = :30, expiry_date = :31, facility_id = :32, facility_license_id = :33, facility_type = :34, field_id = :35, field_station_id = :36, finance_id = :37, fossil_id = :38, hse_incident_id = :39, incident_set_id = :40, incident_type = :41, information_item_id = :42, info_item_subtype = :43, instrument_id = :44, interest_set_id = :45, interest_set_seq_no = :46, land_right_id = :47, land_right_subtype = :48, land_sale_jurisdiction = :49, land_sale_number = :50, lithology_log_id = :51, lith_log_source = :52, notification_id = :53, obligation_id = :54, obligation_seq_no = :55, paleo_summary_id = :56, pden_id = :57, pden_source = :58, pden_subtype = :59, physical_item_id = :60, platform_id = :61, platform_source = :62, pool_id = :63, ppdm_guid = :64, ppdm_system_id = :65, ppdm_table_name = :66, product_type = :67, prod_string_id = :68, prod_string_source = :69, project_id = :70, provision_id = :71, rate_schedule_id = :72, referenced_guid = :73, referenced_system_id = :74, referenced_table_name = :75, remark = :76, resent_id = :77, reserve_class_id = :78, restriction_id = :79, restriction_version = :80, sample_anal_source = :81, seis_license_id = :82, seis_set_id = :83, seis_set_subtype = :84, sf_subtype = :85, source = :86, spatial_description_id = :87, spatial_obs_no = :88, store_id = :89, store_structure_obs_no = :90, strat_name_set_id = :91, strat_unit_id = :92, support_facility_id = :93, sw_application_id = :94, thesaurus_id = :95, thesaurus_word_id = :96, uwi = :97, well_activity_set_id = :98, well_activity_source = :99, well_activity_type_id = :100, well_license_id = :101, well_license_source = :102, well_set_id = :103, work_order_id = :104, row_changed_by = :105, row_changed_date = :106, row_created_by = :107, row_effective_date = :108, row_expiry_date = :109, row_quality = :110 where consult_id = :112")
	if err != nil {
		return err
	}

	consult_component.Row_changed_date = formatDate(consult_component.Row_changed_date)
	consult_component.Effective_date = formatDateString(consult_component.Effective_date)
	consult_component.Expiry_date = formatDateString(consult_component.Expiry_date)
	consult_component.Row_effective_date = formatDateString(consult_component.Row_effective_date)
	consult_component.Row_expiry_date = formatDateString(consult_component.Row_expiry_date)






	rows, err := stmt.Exec(consult_component.Component_id, consult_component.Active_ind, consult_component.Activity_obs_no, consult_component.Analysis_id, consult_component.Application_id, consult_component.Area_id, consult_component.Area_type, consult_component.Authority_id, consult_component.Ba_licensee_ba_id, consult_component.Ba_license_id, consult_component.Ba_organization_id, consult_component.Ba_organization_seq_no, consult_component.Business_associate_id, consult_component.Catalogue_additive_id, consult_component.Catalogue_equip_id, consult_component.Classification_system_id, consult_component.Class_level_id, consult_component.Component_type, consult_component.Consent_id, consult_component.Contest_id, consult_component.Contract_id, consult_component.Currency_conversion, consult_component.Currency_ouom, consult_component.Ecozone_id, consult_component.Effective_date, consult_component.Employee_ba_id, consult_component.Employee_obs_no, consult_component.Employer_ba_id, consult_component.Entitlement_id, consult_component.Equipment_id, consult_component.Expiry_date, consult_component.Facility_id, consult_component.Facility_license_id, consult_component.Facility_type, consult_component.Field_id, consult_component.Field_station_id, consult_component.Finance_id, consult_component.Fossil_id, consult_component.Hse_incident_id, consult_component.Incident_set_id, consult_component.Incident_type, consult_component.Information_item_id, consult_component.Info_item_subtype, consult_component.Instrument_id, consult_component.Interest_set_id, consult_component.Interest_set_seq_no, consult_component.Land_right_id, consult_component.Land_right_subtype, consult_component.Land_sale_jurisdiction, consult_component.Land_sale_number, consult_component.Lithology_log_id, consult_component.Lith_log_source, consult_component.Notification_id, consult_component.Obligation_id, consult_component.Obligation_seq_no, consult_component.Paleo_summary_id, consult_component.Pden_id, consult_component.Pden_source, consult_component.Pden_subtype, consult_component.Physical_item_id, consult_component.Platform_id, consult_component.Platform_source, consult_component.Pool_id, consult_component.Ppdm_guid, consult_component.Ppdm_system_id, consult_component.Ppdm_table_name, consult_component.Product_type, consult_component.Prod_string_id, consult_component.Prod_string_source, consult_component.Project_id, consult_component.Provision_id, consult_component.Rate_schedule_id, consult_component.Referenced_guid, consult_component.Referenced_system_id, consult_component.Referenced_table_name, consult_component.Remark, consult_component.Resent_id, consult_component.Reserve_class_id, consult_component.Restriction_id, consult_component.Restriction_version, consult_component.Sample_anal_source, consult_component.Seis_license_id, consult_component.Seis_set_id, consult_component.Seis_set_subtype, consult_component.Sf_subtype, consult_component.Source, consult_component.Spatial_description_id, consult_component.Spatial_obs_no, consult_component.Store_id, consult_component.Store_structure_obs_no, consult_component.Strat_name_set_id, consult_component.Strat_unit_id, consult_component.Support_facility_id, consult_component.Sw_application_id, consult_component.Thesaurus_id, consult_component.Thesaurus_word_id, consult_component.Uwi, consult_component.Well_activity_set_id, consult_component.Well_activity_source, consult_component.Well_activity_type_id, consult_component.Well_license_id, consult_component.Well_license_source, consult_component.Well_set_id, consult_component.Work_order_id, consult_component.Row_changed_by, consult_component.Row_changed_date, consult_component.Row_created_by, consult_component.Row_effective_date, consult_component.Row_expiry_date, consult_component.Row_quality, consult_component.Consult_id)
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

func PatchConsultComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update consult_component set "
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
	query += " where consult_id = :id"

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

func DeleteConsultComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var consult_component dto.Consult_component
	consult_component.Consult_id = id

	stmt, err := db.Prepare("delete from consult_component where consult_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(consult_component.Consult_id)
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


