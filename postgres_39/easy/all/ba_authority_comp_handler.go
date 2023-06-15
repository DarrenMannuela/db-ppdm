package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetBaAuthorityComp(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ba_authority_comp")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ba_authority_comp

	for rows.Next() {
		var ba_authority_comp dto.Ba_authority_comp
		if err := rows.Scan(&ba_authority_comp.Business_associate_id, &ba_authority_comp.Authority_id, &ba_authority_comp.Component_obs_no, &ba_authority_comp.Active_ind, &ba_authority_comp.Activity_obs_no, &ba_authority_comp.Analysis_id, &ba_authority_comp.Application_id, &ba_authority_comp.Area_id, &ba_authority_comp.Area_type, &ba_authority_comp.Authority_type, &ba_authority_comp.Ba_authority_component_type, &ba_authority_comp.Ba_organization_id, &ba_authority_comp.Ba_organization_seq_no, &ba_authority_comp.Catalogue_additive_id, &ba_authority_comp.Catalogue_equip_id, &ba_authority_comp.Classification_system_id, &ba_authority_comp.Class_level_id, &ba_authority_comp.Consent_id, &ba_authority_comp.Consult_id, &ba_authority_comp.Contest_id, &ba_authority_comp.Contract_id, &ba_authority_comp.Ecozone_id, &ba_authority_comp.Effective_date, &ba_authority_comp.Employee_ba_id, &ba_authority_comp.Employee_obs_no, &ba_authority_comp.Employer_ba_id, &ba_authority_comp.Entitlement_id, &ba_authority_comp.Equipment_id, &ba_authority_comp.Expiry_date, &ba_authority_comp.Facility_id, &ba_authority_comp.Facility_type, &ba_authority_comp.Field_id, &ba_authority_comp.Finance_id, &ba_authority_comp.Fossil_id, &ba_authority_comp.Incident_id, &ba_authority_comp.Incident_set_id, &ba_authority_comp.Incident_type, &ba_authority_comp.Information_item_id, &ba_authority_comp.Info_item_subtype, &ba_authority_comp.Instrument_id, &ba_authority_comp.Interest_set_id, &ba_authority_comp.Interest_set_seq_no, &ba_authority_comp.Jurisdiction, &ba_authority_comp.Land_right_id, &ba_authority_comp.Land_right_subtype, &ba_authority_comp.Land_sale_number, &ba_authority_comp.Lithology_log_id, &ba_authority_comp.Lith_log_source, &ba_authority_comp.Notification_id, &ba_authority_comp.Obligation_id, &ba_authority_comp.Obligation_seq_no, &ba_authority_comp.Paleo_summary_id, &ba_authority_comp.Pden_id, &ba_authority_comp.Pden_source, &ba_authority_comp.Pden_subtype, &ba_authority_comp.Physical_item_id, &ba_authority_comp.Pool_id, &ba_authority_comp.Ppdm_guid, &ba_authority_comp.Ppdm_system_id, &ba_authority_comp.Ppdm_table_name, &ba_authority_comp.Product_type, &ba_authority_comp.Prod_string_id, &ba_authority_comp.Prod_string_source, &ba_authority_comp.Project_id, &ba_authority_comp.Pr_str_form_obs_no, &ba_authority_comp.Rate_schedule_id, &ba_authority_comp.Referenced_guid, &ba_authority_comp.Referenced_system_id, &ba_authority_comp.Referenced_table_name, &ba_authority_comp.Remark, &ba_authority_comp.Resent_id, &ba_authority_comp.Reserve_class_id, &ba_authority_comp.Sample_anal_source, &ba_authority_comp.Seis_set_id, &ba_authority_comp.Seis_set_subtype, &ba_authority_comp.Sf_subtype, &ba_authority_comp.Source, &ba_authority_comp.Spatial_description_id, &ba_authority_comp.Spatial_obs_no, &ba_authority_comp.Store_id, &ba_authority_comp.Store_structure_obs_no, &ba_authority_comp.Strat_name_set_id, &ba_authority_comp.Strat_unit_id, &ba_authority_comp.Support_facility_id, &ba_authority_comp.Sw_application_id, &ba_authority_comp.Thesaurus_id, &ba_authority_comp.Thesaurus_word_id, &ba_authority_comp.Uwi, &ba_authority_comp.Well_activity_set_id, &ba_authority_comp.Well_activity_source, &ba_authority_comp.Well_activity_type_id, &ba_authority_comp.Well_set_id, &ba_authority_comp.Work_order_id, &ba_authority_comp.Row_changed_by, &ba_authority_comp.Row_changed_date, &ba_authority_comp.Row_created_by, &ba_authority_comp.Row_created_date, &ba_authority_comp.Row_effective_date, &ba_authority_comp.Row_expiry_date, &ba_authority_comp.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ba_authority_comp to result
		result = append(result, ba_authority_comp)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetBaAuthorityComp(c *fiber.Ctx) error {
	var ba_authority_comp dto.Ba_authority_comp

	setDefaults(&ba_authority_comp)

	if err := json.Unmarshal(c.Body(), &ba_authority_comp); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ba_authority_comp values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100)")
	if err != nil {
		return err
	}
	ba_authority_comp.Row_created_date = formatDate(ba_authority_comp.Row_created_date)
	ba_authority_comp.Row_changed_date = formatDate(ba_authority_comp.Row_changed_date)
	ba_authority_comp.Effective_date = formatDateString(ba_authority_comp.Effective_date)
	ba_authority_comp.Expiry_date = formatDateString(ba_authority_comp.Expiry_date)
	ba_authority_comp.Row_effective_date = formatDateString(ba_authority_comp.Row_effective_date)
	ba_authority_comp.Row_expiry_date = formatDateString(ba_authority_comp.Row_expiry_date)






	rows, err := stmt.Exec(ba_authority_comp.Business_associate_id, ba_authority_comp.Authority_id, ba_authority_comp.Component_obs_no, ba_authority_comp.Active_ind, ba_authority_comp.Activity_obs_no, ba_authority_comp.Analysis_id, ba_authority_comp.Application_id, ba_authority_comp.Area_id, ba_authority_comp.Area_type, ba_authority_comp.Authority_type, ba_authority_comp.Ba_authority_component_type, ba_authority_comp.Ba_organization_id, ba_authority_comp.Ba_organization_seq_no, ba_authority_comp.Catalogue_additive_id, ba_authority_comp.Catalogue_equip_id, ba_authority_comp.Classification_system_id, ba_authority_comp.Class_level_id, ba_authority_comp.Consent_id, ba_authority_comp.Consult_id, ba_authority_comp.Contest_id, ba_authority_comp.Contract_id, ba_authority_comp.Ecozone_id, ba_authority_comp.Effective_date, ba_authority_comp.Employee_ba_id, ba_authority_comp.Employee_obs_no, ba_authority_comp.Employer_ba_id, ba_authority_comp.Entitlement_id, ba_authority_comp.Equipment_id, ba_authority_comp.Expiry_date, ba_authority_comp.Facility_id, ba_authority_comp.Facility_type, ba_authority_comp.Field_id, ba_authority_comp.Finance_id, ba_authority_comp.Fossil_id, ba_authority_comp.Incident_id, ba_authority_comp.Incident_set_id, ba_authority_comp.Incident_type, ba_authority_comp.Information_item_id, ba_authority_comp.Info_item_subtype, ba_authority_comp.Instrument_id, ba_authority_comp.Interest_set_id, ba_authority_comp.Interest_set_seq_no, ba_authority_comp.Jurisdiction, ba_authority_comp.Land_right_id, ba_authority_comp.Land_right_subtype, ba_authority_comp.Land_sale_number, ba_authority_comp.Lithology_log_id, ba_authority_comp.Lith_log_source, ba_authority_comp.Notification_id, ba_authority_comp.Obligation_id, ba_authority_comp.Obligation_seq_no, ba_authority_comp.Paleo_summary_id, ba_authority_comp.Pden_id, ba_authority_comp.Pden_source, ba_authority_comp.Pden_subtype, ba_authority_comp.Physical_item_id, ba_authority_comp.Pool_id, ba_authority_comp.Ppdm_guid, ba_authority_comp.Ppdm_system_id, ba_authority_comp.Ppdm_table_name, ba_authority_comp.Product_type, ba_authority_comp.Prod_string_id, ba_authority_comp.Prod_string_source, ba_authority_comp.Project_id, ba_authority_comp.Pr_str_form_obs_no, ba_authority_comp.Rate_schedule_id, ba_authority_comp.Referenced_guid, ba_authority_comp.Referenced_system_id, ba_authority_comp.Referenced_table_name, ba_authority_comp.Remark, ba_authority_comp.Resent_id, ba_authority_comp.Reserve_class_id, ba_authority_comp.Sample_anal_source, ba_authority_comp.Seis_set_id, ba_authority_comp.Seis_set_subtype, ba_authority_comp.Sf_subtype, ba_authority_comp.Source, ba_authority_comp.Spatial_description_id, ba_authority_comp.Spatial_obs_no, ba_authority_comp.Store_id, ba_authority_comp.Store_structure_obs_no, ba_authority_comp.Strat_name_set_id, ba_authority_comp.Strat_unit_id, ba_authority_comp.Support_facility_id, ba_authority_comp.Sw_application_id, ba_authority_comp.Thesaurus_id, ba_authority_comp.Thesaurus_word_id, ba_authority_comp.Uwi, ba_authority_comp.Well_activity_set_id, ba_authority_comp.Well_activity_source, ba_authority_comp.Well_activity_type_id, ba_authority_comp.Well_set_id, ba_authority_comp.Work_order_id, ba_authority_comp.Row_changed_by, ba_authority_comp.Row_changed_date, ba_authority_comp.Row_created_by, ba_authority_comp.Row_created_date, ba_authority_comp.Row_effective_date, ba_authority_comp.Row_expiry_date, ba_authority_comp.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateBaAuthorityComp(c *fiber.Ctx) error {
	var ba_authority_comp dto.Ba_authority_comp

	setDefaults(&ba_authority_comp)

	if err := json.Unmarshal(c.Body(), &ba_authority_comp); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ba_authority_comp.Business_associate_id = id

    if ba_authority_comp.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ba_authority_comp set authority_id = :1, component_obs_no = :2, active_ind = :3, activity_obs_no = :4, analysis_id = :5, application_id = :6, area_id = :7, area_type = :8, authority_type = :9, ba_authority_component_type = :10, ba_organization_id = :11, ba_organization_seq_no = :12, catalogue_additive_id = :13, catalogue_equip_id = :14, classification_system_id = :15, class_level_id = :16, consent_id = :17, consult_id = :18, contest_id = :19, contract_id = :20, ecozone_id = :21, effective_date = :22, employee_ba_id = :23, employee_obs_no = :24, employer_ba_id = :25, entitlement_id = :26, equipment_id = :27, expiry_date = :28, facility_id = :29, facility_type = :30, field_id = :31, finance_id = :32, fossil_id = :33, incident_id = :34, incident_set_id = :35, incident_type = :36, information_item_id = :37, info_item_subtype = :38, instrument_id = :39, interest_set_id = :40, interest_set_seq_no = :41, jurisdiction = :42, land_right_id = :43, land_right_subtype = :44, land_sale_number = :45, lithology_log_id = :46, lith_log_source = :47, notification_id = :48, obligation_id = :49, obligation_seq_no = :50, paleo_summary_id = :51, pden_id = :52, pden_source = :53, pden_subtype = :54, physical_item_id = :55, pool_id = :56, ppdm_guid = :57, ppdm_system_id = :58, ppdm_table_name = :59, product_type = :60, prod_string_id = :61, prod_string_source = :62, project_id = :63, pr_str_form_obs_no = :64, rate_schedule_id = :65, referenced_guid = :66, referenced_system_id = :67, referenced_table_name = :68, remark = :69, resent_id = :70, reserve_class_id = :71, sample_anal_source = :72, seis_set_id = :73, seis_set_subtype = :74, sf_subtype = :75, source = :76, spatial_description_id = :77, spatial_obs_no = :78, store_id = :79, store_structure_obs_no = :80, strat_name_set_id = :81, strat_unit_id = :82, support_facility_id = :83, sw_application_id = :84, thesaurus_id = :85, thesaurus_word_id = :86, uwi = :87, well_activity_set_id = :88, well_activity_source = :89, well_activity_type_id = :90, well_set_id = :91, work_order_id = :92, row_changed_by = :93, row_changed_date = :94, row_created_by = :95, row_effective_date = :96, row_expiry_date = :97, row_quality = :98 where business_associate_id = :100")
	if err != nil {
		return err
	}

	ba_authority_comp.Row_changed_date = formatDate(ba_authority_comp.Row_changed_date)
	ba_authority_comp.Effective_date = formatDateString(ba_authority_comp.Effective_date)
	ba_authority_comp.Expiry_date = formatDateString(ba_authority_comp.Expiry_date)
	ba_authority_comp.Row_effective_date = formatDateString(ba_authority_comp.Row_effective_date)
	ba_authority_comp.Row_expiry_date = formatDateString(ba_authority_comp.Row_expiry_date)






	rows, err := stmt.Exec(ba_authority_comp.Authority_id, ba_authority_comp.Component_obs_no, ba_authority_comp.Active_ind, ba_authority_comp.Activity_obs_no, ba_authority_comp.Analysis_id, ba_authority_comp.Application_id, ba_authority_comp.Area_id, ba_authority_comp.Area_type, ba_authority_comp.Authority_type, ba_authority_comp.Ba_authority_component_type, ba_authority_comp.Ba_organization_id, ba_authority_comp.Ba_organization_seq_no, ba_authority_comp.Catalogue_additive_id, ba_authority_comp.Catalogue_equip_id, ba_authority_comp.Classification_system_id, ba_authority_comp.Class_level_id, ba_authority_comp.Consent_id, ba_authority_comp.Consult_id, ba_authority_comp.Contest_id, ba_authority_comp.Contract_id, ba_authority_comp.Ecozone_id, ba_authority_comp.Effective_date, ba_authority_comp.Employee_ba_id, ba_authority_comp.Employee_obs_no, ba_authority_comp.Employer_ba_id, ba_authority_comp.Entitlement_id, ba_authority_comp.Equipment_id, ba_authority_comp.Expiry_date, ba_authority_comp.Facility_id, ba_authority_comp.Facility_type, ba_authority_comp.Field_id, ba_authority_comp.Finance_id, ba_authority_comp.Fossil_id, ba_authority_comp.Incident_id, ba_authority_comp.Incident_set_id, ba_authority_comp.Incident_type, ba_authority_comp.Information_item_id, ba_authority_comp.Info_item_subtype, ba_authority_comp.Instrument_id, ba_authority_comp.Interest_set_id, ba_authority_comp.Interest_set_seq_no, ba_authority_comp.Jurisdiction, ba_authority_comp.Land_right_id, ba_authority_comp.Land_right_subtype, ba_authority_comp.Land_sale_number, ba_authority_comp.Lithology_log_id, ba_authority_comp.Lith_log_source, ba_authority_comp.Notification_id, ba_authority_comp.Obligation_id, ba_authority_comp.Obligation_seq_no, ba_authority_comp.Paleo_summary_id, ba_authority_comp.Pden_id, ba_authority_comp.Pden_source, ba_authority_comp.Pden_subtype, ba_authority_comp.Physical_item_id, ba_authority_comp.Pool_id, ba_authority_comp.Ppdm_guid, ba_authority_comp.Ppdm_system_id, ba_authority_comp.Ppdm_table_name, ba_authority_comp.Product_type, ba_authority_comp.Prod_string_id, ba_authority_comp.Prod_string_source, ba_authority_comp.Project_id, ba_authority_comp.Pr_str_form_obs_no, ba_authority_comp.Rate_schedule_id, ba_authority_comp.Referenced_guid, ba_authority_comp.Referenced_system_id, ba_authority_comp.Referenced_table_name, ba_authority_comp.Remark, ba_authority_comp.Resent_id, ba_authority_comp.Reserve_class_id, ba_authority_comp.Sample_anal_source, ba_authority_comp.Seis_set_id, ba_authority_comp.Seis_set_subtype, ba_authority_comp.Sf_subtype, ba_authority_comp.Source, ba_authority_comp.Spatial_description_id, ba_authority_comp.Spatial_obs_no, ba_authority_comp.Store_id, ba_authority_comp.Store_structure_obs_no, ba_authority_comp.Strat_name_set_id, ba_authority_comp.Strat_unit_id, ba_authority_comp.Support_facility_id, ba_authority_comp.Sw_application_id, ba_authority_comp.Thesaurus_id, ba_authority_comp.Thesaurus_word_id, ba_authority_comp.Uwi, ba_authority_comp.Well_activity_set_id, ba_authority_comp.Well_activity_source, ba_authority_comp.Well_activity_type_id, ba_authority_comp.Well_set_id, ba_authority_comp.Work_order_id, ba_authority_comp.Row_changed_by, ba_authority_comp.Row_changed_date, ba_authority_comp.Row_created_by, ba_authority_comp.Row_effective_date, ba_authority_comp.Row_expiry_date, ba_authority_comp.Row_quality, ba_authority_comp.Business_associate_id)
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

func PatchBaAuthorityComp(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ba_authority_comp set "
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
	query += " where business_associate_id = :id"

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

func DeleteBaAuthorityComp(c *fiber.Ctx) error {
	id := c.Params("id")
	var ba_authority_comp dto.Ba_authority_comp
	ba_authority_comp.Business_associate_id = id

	stmt, err := db.Prepare("delete from ba_authority_comp where business_associate_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ba_authority_comp.Business_associate_id)
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


