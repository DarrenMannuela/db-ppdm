package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetProdStringComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from prod_string_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Prod_string_component

	for rows.Next() {
		var prod_string_component dto.Prod_string_component
		if err := rows.Scan(&prod_string_component.Uwi, &prod_string_component.Pr_str_source, &prod_string_component.String_id, &prod_string_component.Component_obs_no, &prod_string_component.Active_ind, &prod_string_component.Activity_obs_no, &prod_string_component.Analysis_id, &prod_string_component.Application_id, &prod_string_component.Area_id, &prod_string_component.Area_type, &prod_string_component.Authority_id, &prod_string_component.Ba_organization_id, &prod_string_component.Ba_organization_seq_no, &prod_string_component.Business_associate_id, &prod_string_component.Catalogue_additive_id, &prod_string_component.Catalogue_equip_id, &prod_string_component.Classification_system_id, &prod_string_component.Class_level_id, &prod_string_component.Consent_id, &prod_string_component.Consult_id, &prod_string_component.Contest_id, &prod_string_component.Contract_id, &prod_string_component.Ecozone_id, &prod_string_component.Effective_date, &prod_string_component.Employee_ba_id, &prod_string_component.Employee_obs_no, &prod_string_component.Employer_ba_id, &prod_string_component.Entitlement_id, &prod_string_component.Equipment_id, &prod_string_component.Expiry_date, &prod_string_component.Facility_id, &prod_string_component.Facility_type, &prod_string_component.Field_id, &prod_string_component.Field_station_id, &prod_string_component.Finance_id, &prod_string_component.Fossil_id, &prod_string_component.Incident_id, &prod_string_component.Incident_set_id, &prod_string_component.Incident_type, &prod_string_component.Information_item_id, &prod_string_component.Info_item_subtype, &prod_string_component.Instrument_id, &prod_string_component.Interest_set_id, &prod_string_component.Interest_set_seq_no, &prod_string_component.Jurisdiction, &prod_string_component.Land_right_id, &prod_string_component.Land_right_subtype, &prod_string_component.Land_sale_number, &prod_string_component.Language, &prod_string_component.Lithology_log_id, &prod_string_component.Lith_log_source, &prod_string_component.Notification_id, &prod_string_component.Obligation_id, &prod_string_component.Obligation_seq_no, &prod_string_component.Paleo_summary_id, &prod_string_component.Pden_id, &prod_string_component.Pden_source, &prod_string_component.Pden_subtype, &prod_string_component.Physical_item_id, &prod_string_component.Pool_id, &prod_string_component.Ppdm_guid, &prod_string_component.Ppdm_system_id, &prod_string_component.Ppdm_table_name, &prod_string_component.Product_type, &prod_string_component.Prod_string_component_type, &prod_string_component.Project_id, &prod_string_component.Rate_schedule_id, &prod_string_component.Referenced_guid, &prod_string_component.Referenced_system_id, &prod_string_component.Referenced_table_name, &prod_string_component.Remark, &prod_string_component.Resent_id, &prod_string_component.Reserve_class_id, &prod_string_component.Sample_anal_source, &prod_string_component.Seis_set_id, &prod_string_component.Seis_set_subtype, &prod_string_component.Sf_subtype, &prod_string_component.Source, &prod_string_component.Spatial_description_id, &prod_string_component.Spatial_obs_no, &prod_string_component.Store_id, &prod_string_component.Store_structure_obs_no, &prod_string_component.Strat_name_set_id, &prod_string_component.Strat_unit_id, &prod_string_component.Support_facility_id, &prod_string_component.Sw_application_id, &prod_string_component.Thesaurus_id, &prod_string_component.Thesaurus_word_id, &prod_string_component.Well_activity_set_id, &prod_string_component.Well_activity_source, &prod_string_component.Well_activity_type_id, &prod_string_component.Well_activity_uwi, &prod_string_component.Well_set_id, &prod_string_component.Work_order_id, &prod_string_component.Row_changed_by, &prod_string_component.Row_changed_date, &prod_string_component.Row_created_by, &prod_string_component.Row_created_date, &prod_string_component.Row_effective_date, &prod_string_component.Row_expiry_date, &prod_string_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append prod_string_component to result
		result = append(result, prod_string_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetProdStringComponent(c *fiber.Ctx) error {
	var prod_string_component dto.Prod_string_component

	setDefaults(&prod_string_component)

	if err := json.Unmarshal(c.Body(), &prod_string_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into prod_string_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101)")
	if err != nil {
		return err
	}
	prod_string_component.Row_created_date = formatDate(prod_string_component.Row_created_date)
	prod_string_component.Row_changed_date = formatDate(prod_string_component.Row_changed_date)
	prod_string_component.Effective_date = formatDateString(prod_string_component.Effective_date)
	prod_string_component.Expiry_date = formatDateString(prod_string_component.Expiry_date)
	prod_string_component.Row_effective_date = formatDateString(prod_string_component.Row_effective_date)
	prod_string_component.Row_expiry_date = formatDateString(prod_string_component.Row_expiry_date)






	rows, err := stmt.Exec(prod_string_component.Uwi, prod_string_component.Pr_str_source, prod_string_component.String_id, prod_string_component.Component_obs_no, prod_string_component.Active_ind, prod_string_component.Activity_obs_no, prod_string_component.Analysis_id, prod_string_component.Application_id, prod_string_component.Area_id, prod_string_component.Area_type, prod_string_component.Authority_id, prod_string_component.Ba_organization_id, prod_string_component.Ba_organization_seq_no, prod_string_component.Business_associate_id, prod_string_component.Catalogue_additive_id, prod_string_component.Catalogue_equip_id, prod_string_component.Classification_system_id, prod_string_component.Class_level_id, prod_string_component.Consent_id, prod_string_component.Consult_id, prod_string_component.Contest_id, prod_string_component.Contract_id, prod_string_component.Ecozone_id, prod_string_component.Effective_date, prod_string_component.Employee_ba_id, prod_string_component.Employee_obs_no, prod_string_component.Employer_ba_id, prod_string_component.Entitlement_id, prod_string_component.Equipment_id, prod_string_component.Expiry_date, prod_string_component.Facility_id, prod_string_component.Facility_type, prod_string_component.Field_id, prod_string_component.Field_station_id, prod_string_component.Finance_id, prod_string_component.Fossil_id, prod_string_component.Incident_id, prod_string_component.Incident_set_id, prod_string_component.Incident_type, prod_string_component.Information_item_id, prod_string_component.Info_item_subtype, prod_string_component.Instrument_id, prod_string_component.Interest_set_id, prod_string_component.Interest_set_seq_no, prod_string_component.Jurisdiction, prod_string_component.Land_right_id, prod_string_component.Land_right_subtype, prod_string_component.Land_sale_number, prod_string_component.Language, prod_string_component.Lithology_log_id, prod_string_component.Lith_log_source, prod_string_component.Notification_id, prod_string_component.Obligation_id, prod_string_component.Obligation_seq_no, prod_string_component.Paleo_summary_id, prod_string_component.Pden_id, prod_string_component.Pden_source, prod_string_component.Pden_subtype, prod_string_component.Physical_item_id, prod_string_component.Pool_id, prod_string_component.Ppdm_guid, prod_string_component.Ppdm_system_id, prod_string_component.Ppdm_table_name, prod_string_component.Product_type, prod_string_component.Prod_string_component_type, prod_string_component.Project_id, prod_string_component.Rate_schedule_id, prod_string_component.Referenced_guid, prod_string_component.Referenced_system_id, prod_string_component.Referenced_table_name, prod_string_component.Remark, prod_string_component.Resent_id, prod_string_component.Reserve_class_id, prod_string_component.Sample_anal_source, prod_string_component.Seis_set_id, prod_string_component.Seis_set_subtype, prod_string_component.Sf_subtype, prod_string_component.Source, prod_string_component.Spatial_description_id, prod_string_component.Spatial_obs_no, prod_string_component.Store_id, prod_string_component.Store_structure_obs_no, prod_string_component.Strat_name_set_id, prod_string_component.Strat_unit_id, prod_string_component.Support_facility_id, prod_string_component.Sw_application_id, prod_string_component.Thesaurus_id, prod_string_component.Thesaurus_word_id, prod_string_component.Well_activity_set_id, prod_string_component.Well_activity_source, prod_string_component.Well_activity_type_id, prod_string_component.Well_activity_uwi, prod_string_component.Well_set_id, prod_string_component.Work_order_id, prod_string_component.Row_changed_by, prod_string_component.Row_changed_date, prod_string_component.Row_created_by, prod_string_component.Row_created_date, prod_string_component.Row_effective_date, prod_string_component.Row_expiry_date, prod_string_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateProdStringComponent(c *fiber.Ctx) error {
	var prod_string_component dto.Prod_string_component

	setDefaults(&prod_string_component)

	if err := json.Unmarshal(c.Body(), &prod_string_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	prod_string_component.Uwi = id

    if prod_string_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update prod_string_component set pr_str_source = :1, string_id = :2, component_obs_no = :3, active_ind = :4, activity_obs_no = :5, analysis_id = :6, application_id = :7, area_id = :8, area_type = :9, authority_id = :10, ba_organization_id = :11, ba_organization_seq_no = :12, business_associate_id = :13, catalogue_additive_id = :14, catalogue_equip_id = :15, classification_system_id = :16, class_level_id = :17, consent_id = :18, consult_id = :19, contest_id = :20, contract_id = :21, ecozone_id = :22, effective_date = :23, employee_ba_id = :24, employee_obs_no = :25, employer_ba_id = :26, entitlement_id = :27, equipment_id = :28, expiry_date = :29, facility_id = :30, facility_type = :31, field_id = :32, field_station_id = :33, finance_id = :34, fossil_id = :35, incident_id = :36, incident_set_id = :37, incident_type = :38, information_item_id = :39, info_item_subtype = :40, instrument_id = :41, interest_set_id = :42, interest_set_seq_no = :43, jurisdiction = :44, land_right_id = :45, land_right_subtype = :46, land_sale_number = :47, language = :48, lithology_log_id = :49, lith_log_source = :50, notification_id = :51, obligation_id = :52, obligation_seq_no = :53, paleo_summary_id = :54, pden_id = :55, pden_source = :56, pden_subtype = :57, physical_item_id = :58, pool_id = :59, ppdm_guid = :60, ppdm_system_id = :61, ppdm_table_name = :62, product_type = :63, prod_string_component_type = :64, project_id = :65, rate_schedule_id = :66, referenced_guid = :67, referenced_system_id = :68, referenced_table_name = :69, remark = :70, resent_id = :71, reserve_class_id = :72, sample_anal_source = :73, seis_set_id = :74, seis_set_subtype = :75, sf_subtype = :76, source = :77, spatial_description_id = :78, spatial_obs_no = :79, store_id = :80, store_structure_obs_no = :81, strat_name_set_id = :82, strat_unit_id = :83, support_facility_id = :84, sw_application_id = :85, thesaurus_id = :86, thesaurus_word_id = :87, well_activity_set_id = :88, well_activity_source = :89, well_activity_type_id = :90, well_activity_uwi = :91, well_set_id = :92, work_order_id = :93, row_changed_by = :94, row_changed_date = :95, row_created_by = :96, row_effective_date = :97, row_expiry_date = :98, row_quality = :99 where uwi = :101")
	if err != nil {
		return err
	}

	prod_string_component.Row_changed_date = formatDate(prod_string_component.Row_changed_date)
	prod_string_component.Effective_date = formatDateString(prod_string_component.Effective_date)
	prod_string_component.Expiry_date = formatDateString(prod_string_component.Expiry_date)
	prod_string_component.Row_effective_date = formatDateString(prod_string_component.Row_effective_date)
	prod_string_component.Row_expiry_date = formatDateString(prod_string_component.Row_expiry_date)






	rows, err := stmt.Exec(prod_string_component.Pr_str_source, prod_string_component.String_id, prod_string_component.Component_obs_no, prod_string_component.Active_ind, prod_string_component.Activity_obs_no, prod_string_component.Analysis_id, prod_string_component.Application_id, prod_string_component.Area_id, prod_string_component.Area_type, prod_string_component.Authority_id, prod_string_component.Ba_organization_id, prod_string_component.Ba_organization_seq_no, prod_string_component.Business_associate_id, prod_string_component.Catalogue_additive_id, prod_string_component.Catalogue_equip_id, prod_string_component.Classification_system_id, prod_string_component.Class_level_id, prod_string_component.Consent_id, prod_string_component.Consult_id, prod_string_component.Contest_id, prod_string_component.Contract_id, prod_string_component.Ecozone_id, prod_string_component.Effective_date, prod_string_component.Employee_ba_id, prod_string_component.Employee_obs_no, prod_string_component.Employer_ba_id, prod_string_component.Entitlement_id, prod_string_component.Equipment_id, prod_string_component.Expiry_date, prod_string_component.Facility_id, prod_string_component.Facility_type, prod_string_component.Field_id, prod_string_component.Field_station_id, prod_string_component.Finance_id, prod_string_component.Fossil_id, prod_string_component.Incident_id, prod_string_component.Incident_set_id, prod_string_component.Incident_type, prod_string_component.Information_item_id, prod_string_component.Info_item_subtype, prod_string_component.Instrument_id, prod_string_component.Interest_set_id, prod_string_component.Interest_set_seq_no, prod_string_component.Jurisdiction, prod_string_component.Land_right_id, prod_string_component.Land_right_subtype, prod_string_component.Land_sale_number, prod_string_component.Language, prod_string_component.Lithology_log_id, prod_string_component.Lith_log_source, prod_string_component.Notification_id, prod_string_component.Obligation_id, prod_string_component.Obligation_seq_no, prod_string_component.Paleo_summary_id, prod_string_component.Pden_id, prod_string_component.Pden_source, prod_string_component.Pden_subtype, prod_string_component.Physical_item_id, prod_string_component.Pool_id, prod_string_component.Ppdm_guid, prod_string_component.Ppdm_system_id, prod_string_component.Ppdm_table_name, prod_string_component.Product_type, prod_string_component.Prod_string_component_type, prod_string_component.Project_id, prod_string_component.Rate_schedule_id, prod_string_component.Referenced_guid, prod_string_component.Referenced_system_id, prod_string_component.Referenced_table_name, prod_string_component.Remark, prod_string_component.Resent_id, prod_string_component.Reserve_class_id, prod_string_component.Sample_anal_source, prod_string_component.Seis_set_id, prod_string_component.Seis_set_subtype, prod_string_component.Sf_subtype, prod_string_component.Source, prod_string_component.Spatial_description_id, prod_string_component.Spatial_obs_no, prod_string_component.Store_id, prod_string_component.Store_structure_obs_no, prod_string_component.Strat_name_set_id, prod_string_component.Strat_unit_id, prod_string_component.Support_facility_id, prod_string_component.Sw_application_id, prod_string_component.Thesaurus_id, prod_string_component.Thesaurus_word_id, prod_string_component.Well_activity_set_id, prod_string_component.Well_activity_source, prod_string_component.Well_activity_type_id, prod_string_component.Well_activity_uwi, prod_string_component.Well_set_id, prod_string_component.Work_order_id, prod_string_component.Row_changed_by, prod_string_component.Row_changed_date, prod_string_component.Row_created_by, prod_string_component.Row_effective_date, prod_string_component.Row_expiry_date, prod_string_component.Row_quality, prod_string_component.Uwi)
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

func PatchProdStringComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update prod_string_component set "
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
	query += " where uwi = :id"

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

func DeleteProdStringComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var prod_string_component dto.Prod_string_component
	prod_string_component.Uwi = id

	stmt, err := db.Prepare("delete from prod_string_component where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(prod_string_component.Uwi)
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


