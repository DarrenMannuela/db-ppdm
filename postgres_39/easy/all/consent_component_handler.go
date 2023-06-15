package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetConsentComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from consent_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Consent_component

	for rows.Next() {
		var consent_component dto.Consent_component
		if err := rows.Scan(&consent_component.Consent_id, &consent_component.Component_id, &consent_component.Active_ind, &consent_component.Activity_obs_no, &consent_component.Analysis_id, &consent_component.Application_id, &consent_component.Area_id, &consent_component.Area_type, &consent_component.Authority_id, &consent_component.Ba_licensee_ba_id, &consent_component.Ba_license_id, &consent_component.Ba_organization_id, &consent_component.Ba_organization_seq_no, &consent_component.Business_associate_id, &consent_component.Catalogue_additive_id, &consent_component.Catalogue_equip_id, &consent_component.Classification_system_id, &consent_component.Class_level_id, &consent_component.Consent_component_type, &consent_component.Consult_id, &consent_component.Contest_id, &consent_component.Contract_id, &consent_component.Ecozone_id, &consent_component.Effective_date, &consent_component.Employee_ba_id, &consent_component.Employee_obs_no, &consent_component.Employer_ba_id, &consent_component.Entitlement_id, &consent_component.Equipment_id, &consent_component.Expiry_date, &consent_component.Facility_id, &consent_component.Facility_license_id, &consent_component.Facility_type, &consent_component.Field_id, &consent_component.Field_station_id, &consent_component.Finance_id, &consent_component.Fossil_id, &consent_component.Incident_id, &consent_component.Incident_set_id, &consent_component.Incident_type, &consent_component.Information_item_id, &consent_component.Info_item_subtype, &consent_component.Instrument_id, &consent_component.Interest_set_id, &consent_component.Interest_set_seq_no, &consent_component.Jurisdiction, &consent_component.Land_right_id, &consent_component.Land_right_subtype, &consent_component.Land_sale_number, &consent_component.Language, &consent_component.Lithology_log_id, &consent_component.Lith_log_source, &consent_component.Notification_id, &consent_component.Obligation_id, &consent_component.Obligation_seq_no, &consent_component.Paleo_summary_id, &consent_component.Pden_id, &consent_component.Pden_source, &consent_component.Pden_subtype, &consent_component.Physical_item_id, &consent_component.Pool_id, &consent_component.Ppdm_guid, &consent_component.Ppdm_system_id, &consent_component.Ppdm_table_name, &consent_component.Product_type, &consent_component.Prod_string_id, &consent_component.Prod_string_source, &consent_component.Project_id, &consent_component.Rate_schedule_id, &consent_component.Referenced_guid, &consent_component.Referenced_system_id, &consent_component.Referenced_table_name, &consent_component.Remark, &consent_component.Resent_id, &consent_component.Reserve_class_id, &consent_component.Restriction_id, &consent_component.Restriction_version, &consent_component.Sample_anal_source, &consent_component.Seis_license_id, &consent_component.Seis_set_id, &consent_component.Seis_set_subtype, &consent_component.Sf_subtype, &consent_component.Source, &consent_component.Spatial_description_id, &consent_component.Spatial_obs_no, &consent_component.Store_id, &consent_component.Store_structure_obs_no, &consent_component.Strat_name_set_id, &consent_component.Strat_unit_id, &consent_component.Support_facility_id, &consent_component.Sw_application_id, &consent_component.Thesaurus_id, &consent_component.Thesaurus_word_id, &consent_component.Uwi, &consent_component.Well_activity_set_id, &consent_component.Well_activity_source, &consent_component.Well_activity_type_id, &consent_component.Well_license_id, &consent_component.Well_license_source, &consent_component.Well_set_id, &consent_component.Work_order_id, &consent_component.Row_changed_by, &consent_component.Row_changed_date, &consent_component.Row_created_by, &consent_component.Row_created_date, &consent_component.Row_effective_date, &consent_component.Row_expiry_date, &consent_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append consent_component to result
		result = append(result, consent_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetConsentComponent(c *fiber.Ctx) error {
	var consent_component dto.Consent_component

	setDefaults(&consent_component)

	if err := json.Unmarshal(c.Body(), &consent_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into consent_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103, :104, :105, :106, :107, :108)")
	if err != nil {
		return err
	}
	consent_component.Row_created_date = formatDate(consent_component.Row_created_date)
	consent_component.Row_changed_date = formatDate(consent_component.Row_changed_date)
	consent_component.Effective_date = formatDateString(consent_component.Effective_date)
	consent_component.Expiry_date = formatDateString(consent_component.Expiry_date)
	consent_component.Row_effective_date = formatDateString(consent_component.Row_effective_date)
	consent_component.Row_expiry_date = formatDateString(consent_component.Row_expiry_date)






	rows, err := stmt.Exec(consent_component.Consent_id, consent_component.Component_id, consent_component.Active_ind, consent_component.Activity_obs_no, consent_component.Analysis_id, consent_component.Application_id, consent_component.Area_id, consent_component.Area_type, consent_component.Authority_id, consent_component.Ba_licensee_ba_id, consent_component.Ba_license_id, consent_component.Ba_organization_id, consent_component.Ba_organization_seq_no, consent_component.Business_associate_id, consent_component.Catalogue_additive_id, consent_component.Catalogue_equip_id, consent_component.Classification_system_id, consent_component.Class_level_id, consent_component.Consent_component_type, consent_component.Consult_id, consent_component.Contest_id, consent_component.Contract_id, consent_component.Ecozone_id, consent_component.Effective_date, consent_component.Employee_ba_id, consent_component.Employee_obs_no, consent_component.Employer_ba_id, consent_component.Entitlement_id, consent_component.Equipment_id, consent_component.Expiry_date, consent_component.Facility_id, consent_component.Facility_license_id, consent_component.Facility_type, consent_component.Field_id, consent_component.Field_station_id, consent_component.Finance_id, consent_component.Fossil_id, consent_component.Incident_id, consent_component.Incident_set_id, consent_component.Incident_type, consent_component.Information_item_id, consent_component.Info_item_subtype, consent_component.Instrument_id, consent_component.Interest_set_id, consent_component.Interest_set_seq_no, consent_component.Jurisdiction, consent_component.Land_right_id, consent_component.Land_right_subtype, consent_component.Land_sale_number, consent_component.Language, consent_component.Lithology_log_id, consent_component.Lith_log_source, consent_component.Notification_id, consent_component.Obligation_id, consent_component.Obligation_seq_no, consent_component.Paleo_summary_id, consent_component.Pden_id, consent_component.Pden_source, consent_component.Pden_subtype, consent_component.Physical_item_id, consent_component.Pool_id, consent_component.Ppdm_guid, consent_component.Ppdm_system_id, consent_component.Ppdm_table_name, consent_component.Product_type, consent_component.Prod_string_id, consent_component.Prod_string_source, consent_component.Project_id, consent_component.Rate_schedule_id, consent_component.Referenced_guid, consent_component.Referenced_system_id, consent_component.Referenced_table_name, consent_component.Remark, consent_component.Resent_id, consent_component.Reserve_class_id, consent_component.Restriction_id, consent_component.Restriction_version, consent_component.Sample_anal_source, consent_component.Seis_license_id, consent_component.Seis_set_id, consent_component.Seis_set_subtype, consent_component.Sf_subtype, consent_component.Source, consent_component.Spatial_description_id, consent_component.Spatial_obs_no, consent_component.Store_id, consent_component.Store_structure_obs_no, consent_component.Strat_name_set_id, consent_component.Strat_unit_id, consent_component.Support_facility_id, consent_component.Sw_application_id, consent_component.Thesaurus_id, consent_component.Thesaurus_word_id, consent_component.Uwi, consent_component.Well_activity_set_id, consent_component.Well_activity_source, consent_component.Well_activity_type_id, consent_component.Well_license_id, consent_component.Well_license_source, consent_component.Well_set_id, consent_component.Work_order_id, consent_component.Row_changed_by, consent_component.Row_changed_date, consent_component.Row_created_by, consent_component.Row_created_date, consent_component.Row_effective_date, consent_component.Row_expiry_date, consent_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateConsentComponent(c *fiber.Ctx) error {
	var consent_component dto.Consent_component

	setDefaults(&consent_component)

	if err := json.Unmarshal(c.Body(), &consent_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	consent_component.Consent_id = id

    if consent_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update consent_component set component_id = :1, active_ind = :2, activity_obs_no = :3, analysis_id = :4, application_id = :5, area_id = :6, area_type = :7, authority_id = :8, ba_licensee_ba_id = :9, ba_license_id = :10, ba_organization_id = :11, ba_organization_seq_no = :12, business_associate_id = :13, catalogue_additive_id = :14, catalogue_equip_id = :15, classification_system_id = :16, class_level_id = :17, consent_component_type = :18, consult_id = :19, contest_id = :20, contract_id = :21, ecozone_id = :22, effective_date = :23, employee_ba_id = :24, employee_obs_no = :25, employer_ba_id = :26, entitlement_id = :27, equipment_id = :28, expiry_date = :29, facility_id = :30, facility_license_id = :31, facility_type = :32, field_id = :33, field_station_id = :34, finance_id = :35, fossil_id = :36, incident_id = :37, incident_set_id = :38, incident_type = :39, information_item_id = :40, info_item_subtype = :41, instrument_id = :42, interest_set_id = :43, interest_set_seq_no = :44, jurisdiction = :45, land_right_id = :46, land_right_subtype = :47, land_sale_number = :48, language = :49, lithology_log_id = :50, lith_log_source = :51, notification_id = :52, obligation_id = :53, obligation_seq_no = :54, paleo_summary_id = :55, pden_id = :56, pden_source = :57, pden_subtype = :58, physical_item_id = :59, pool_id = :60, ppdm_guid = :61, ppdm_system_id = :62, ppdm_table_name = :63, product_type = :64, prod_string_id = :65, prod_string_source = :66, project_id = :67, rate_schedule_id = :68, referenced_guid = :69, referenced_system_id = :70, referenced_table_name = :71, remark = :72, resent_id = :73, reserve_class_id = :74, restriction_id = :75, restriction_version = :76, sample_anal_source = :77, seis_license_id = :78, seis_set_id = :79, seis_set_subtype = :80, sf_subtype = :81, source = :82, spatial_description_id = :83, spatial_obs_no = :84, store_id = :85, store_structure_obs_no = :86, strat_name_set_id = :87, strat_unit_id = :88, support_facility_id = :89, sw_application_id = :90, thesaurus_id = :91, thesaurus_word_id = :92, uwi = :93, well_activity_set_id = :94, well_activity_source = :95, well_activity_type_id = :96, well_license_id = :97, well_license_source = :98, well_set_id = :99, work_order_id = :100, row_changed_by = :101, row_changed_date = :102, row_created_by = :103, row_effective_date = :104, row_expiry_date = :105, row_quality = :106 where consent_id = :108")
	if err != nil {
		return err
	}

	consent_component.Row_changed_date = formatDate(consent_component.Row_changed_date)
	consent_component.Effective_date = formatDateString(consent_component.Effective_date)
	consent_component.Expiry_date = formatDateString(consent_component.Expiry_date)
	consent_component.Row_effective_date = formatDateString(consent_component.Row_effective_date)
	consent_component.Row_expiry_date = formatDateString(consent_component.Row_expiry_date)






	rows, err := stmt.Exec(consent_component.Component_id, consent_component.Active_ind, consent_component.Activity_obs_no, consent_component.Analysis_id, consent_component.Application_id, consent_component.Area_id, consent_component.Area_type, consent_component.Authority_id, consent_component.Ba_licensee_ba_id, consent_component.Ba_license_id, consent_component.Ba_organization_id, consent_component.Ba_organization_seq_no, consent_component.Business_associate_id, consent_component.Catalogue_additive_id, consent_component.Catalogue_equip_id, consent_component.Classification_system_id, consent_component.Class_level_id, consent_component.Consent_component_type, consent_component.Consult_id, consent_component.Contest_id, consent_component.Contract_id, consent_component.Ecozone_id, consent_component.Effective_date, consent_component.Employee_ba_id, consent_component.Employee_obs_no, consent_component.Employer_ba_id, consent_component.Entitlement_id, consent_component.Equipment_id, consent_component.Expiry_date, consent_component.Facility_id, consent_component.Facility_license_id, consent_component.Facility_type, consent_component.Field_id, consent_component.Field_station_id, consent_component.Finance_id, consent_component.Fossil_id, consent_component.Incident_id, consent_component.Incident_set_id, consent_component.Incident_type, consent_component.Information_item_id, consent_component.Info_item_subtype, consent_component.Instrument_id, consent_component.Interest_set_id, consent_component.Interest_set_seq_no, consent_component.Jurisdiction, consent_component.Land_right_id, consent_component.Land_right_subtype, consent_component.Land_sale_number, consent_component.Language, consent_component.Lithology_log_id, consent_component.Lith_log_source, consent_component.Notification_id, consent_component.Obligation_id, consent_component.Obligation_seq_no, consent_component.Paleo_summary_id, consent_component.Pden_id, consent_component.Pden_source, consent_component.Pden_subtype, consent_component.Physical_item_id, consent_component.Pool_id, consent_component.Ppdm_guid, consent_component.Ppdm_system_id, consent_component.Ppdm_table_name, consent_component.Product_type, consent_component.Prod_string_id, consent_component.Prod_string_source, consent_component.Project_id, consent_component.Rate_schedule_id, consent_component.Referenced_guid, consent_component.Referenced_system_id, consent_component.Referenced_table_name, consent_component.Remark, consent_component.Resent_id, consent_component.Reserve_class_id, consent_component.Restriction_id, consent_component.Restriction_version, consent_component.Sample_anal_source, consent_component.Seis_license_id, consent_component.Seis_set_id, consent_component.Seis_set_subtype, consent_component.Sf_subtype, consent_component.Source, consent_component.Spatial_description_id, consent_component.Spatial_obs_no, consent_component.Store_id, consent_component.Store_structure_obs_no, consent_component.Strat_name_set_id, consent_component.Strat_unit_id, consent_component.Support_facility_id, consent_component.Sw_application_id, consent_component.Thesaurus_id, consent_component.Thesaurus_word_id, consent_component.Uwi, consent_component.Well_activity_set_id, consent_component.Well_activity_source, consent_component.Well_activity_type_id, consent_component.Well_license_id, consent_component.Well_license_source, consent_component.Well_set_id, consent_component.Work_order_id, consent_component.Row_changed_by, consent_component.Row_changed_date, consent_component.Row_created_by, consent_component.Row_effective_date, consent_component.Row_expiry_date, consent_component.Row_quality, consent_component.Consent_id)
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

func PatchConsentComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update consent_component set "
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
	query += " where consent_id = :id"

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

func DeleteConsentComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var consent_component dto.Consent_component
	consent_component.Consent_id = id

	stmt, err := db.Prepare("delete from consent_component where consent_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(consent_component.Consent_id)
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


