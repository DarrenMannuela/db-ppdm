package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandRightComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_right_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_right_component

	for rows.Next() {
		var land_right_component dto.Land_right_component
		if err := rows.Scan(&land_right_component.Land_right_subtype, &land_right_component.Land_right_id, &land_right_component.Component_obs_no, &land_right_component.Active_ind, &land_right_component.Activity_obs_no, &land_right_component.Analysis_id, &land_right_component.Application_id, &land_right_component.Area_id, &land_right_component.Area_type, &land_right_component.Authority_id, &land_right_component.Ba_organization_id, &land_right_component.Ba_organization_seq_no, &land_right_component.Business_associate_id, &land_right_component.Catalogue_additive_id, &land_right_component.Catalogue_equip_id, &land_right_component.Classification_system_id, &land_right_component.Class_level_id, &land_right_component.Consent_id, &land_right_component.Consult_id, &land_right_component.Contest_id, &land_right_component.Contract_id, &land_right_component.Ecozone_id, &land_right_component.Effective_date, &land_right_component.Employee_ba_id, &land_right_component.Employee_obs_no, &land_right_component.Employer_ba_id, &land_right_component.Entitlement_id, &land_right_component.Equipment_id, &land_right_component.Expiry_date, &land_right_component.Facility_id, &land_right_component.Facility_type, &land_right_component.Field_id, &land_right_component.Field_station_id, &land_right_component.Finance_id, &land_right_component.Fossil_id, &land_right_component.Incident_id, &land_right_component.Incident_set_id, &land_right_component.Incident_type, &land_right_component.Information_item_id, &land_right_component.Info_item_subtype, &land_right_component.Instrument_id, &land_right_component.Interest_set_id, &land_right_component.Interest_set_seq_no, &land_right_component.Jurisdiction, &land_right_component.Land_component_type, &land_right_component.Land_sale_number, &land_right_component.Language, &land_right_component.Lithology_log_id, &land_right_component.Lith_log_source, &land_right_component.Notification_id, &land_right_component.Obligation_id, &land_right_component.Obligation_seq_no, &land_right_component.Paleo_summary_id, &land_right_component.Pden_id, &land_right_component.Pden_source, &land_right_component.Pden_subtype, &land_right_component.Physical_item_id, &land_right_component.Pool_id, &land_right_component.Ppdm_guid, &land_right_component.Ppdm_system_id, &land_right_component.Ppdm_table_name, &land_right_component.Product_type, &land_right_component.Prod_string_id, &land_right_component.Prod_string_source, &land_right_component.Project_id, &land_right_component.Rate_schedule_id, &land_right_component.Referenced_guid, &land_right_component.Referenced_system_id, &land_right_component.Referenced_table_name, &land_right_component.Remark, &land_right_component.Resent_id, &land_right_component.Reserve_class_id, &land_right_component.Sample_anal_source, &land_right_component.Seis_set_id, &land_right_component.Seis_set_subtype, &land_right_component.Sf_subtype, &land_right_component.Source, &land_right_component.Spatial_description_id, &land_right_component.Spatial_obs_no, &land_right_component.Store_id, &land_right_component.Store_structure_obs_no, &land_right_component.Strat_name_set_id, &land_right_component.Strat_unit_id, &land_right_component.Support_facility_id, &land_right_component.Sw_application_id, &land_right_component.Thesaurus_id, &land_right_component.Thesaurus_word_id, &land_right_component.Uwi, &land_right_component.Well_activity_set_id, &land_right_component.Well_activity_source, &land_right_component.Well_activity_type_id, &land_right_component.Well_set_id, &land_right_component.Work_order_id, &land_right_component.Row_changed_by, &land_right_component.Row_changed_date, &land_right_component.Row_created_by, &land_right_component.Row_created_date, &land_right_component.Row_effective_date, &land_right_component.Row_expiry_date, &land_right_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_right_component to result
		result = append(result, land_right_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandRightComponent(c *fiber.Ctx) error {
	var land_right_component dto.Land_right_component

	setDefaults(&land_right_component)

	if err := json.Unmarshal(c.Body(), &land_right_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_right_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100)")
	if err != nil {
		return err
	}
	land_right_component.Row_created_date = formatDate(land_right_component.Row_created_date)
	land_right_component.Row_changed_date = formatDate(land_right_component.Row_changed_date)
	land_right_component.Effective_date = formatDateString(land_right_component.Effective_date)
	land_right_component.Expiry_date = formatDateString(land_right_component.Expiry_date)
	land_right_component.Row_effective_date = formatDateString(land_right_component.Row_effective_date)
	land_right_component.Row_expiry_date = formatDateString(land_right_component.Row_expiry_date)






	rows, err := stmt.Exec(land_right_component.Land_right_subtype, land_right_component.Land_right_id, land_right_component.Component_obs_no, land_right_component.Active_ind, land_right_component.Activity_obs_no, land_right_component.Analysis_id, land_right_component.Application_id, land_right_component.Area_id, land_right_component.Area_type, land_right_component.Authority_id, land_right_component.Ba_organization_id, land_right_component.Ba_organization_seq_no, land_right_component.Business_associate_id, land_right_component.Catalogue_additive_id, land_right_component.Catalogue_equip_id, land_right_component.Classification_system_id, land_right_component.Class_level_id, land_right_component.Consent_id, land_right_component.Consult_id, land_right_component.Contest_id, land_right_component.Contract_id, land_right_component.Ecozone_id, land_right_component.Effective_date, land_right_component.Employee_ba_id, land_right_component.Employee_obs_no, land_right_component.Employer_ba_id, land_right_component.Entitlement_id, land_right_component.Equipment_id, land_right_component.Expiry_date, land_right_component.Facility_id, land_right_component.Facility_type, land_right_component.Field_id, land_right_component.Field_station_id, land_right_component.Finance_id, land_right_component.Fossil_id, land_right_component.Incident_id, land_right_component.Incident_set_id, land_right_component.Incident_type, land_right_component.Information_item_id, land_right_component.Info_item_subtype, land_right_component.Instrument_id, land_right_component.Interest_set_id, land_right_component.Interest_set_seq_no, land_right_component.Jurisdiction, land_right_component.Land_component_type, land_right_component.Land_sale_number, land_right_component.Language, land_right_component.Lithology_log_id, land_right_component.Lith_log_source, land_right_component.Notification_id, land_right_component.Obligation_id, land_right_component.Obligation_seq_no, land_right_component.Paleo_summary_id, land_right_component.Pden_id, land_right_component.Pden_source, land_right_component.Pden_subtype, land_right_component.Physical_item_id, land_right_component.Pool_id, land_right_component.Ppdm_guid, land_right_component.Ppdm_system_id, land_right_component.Ppdm_table_name, land_right_component.Product_type, land_right_component.Prod_string_id, land_right_component.Prod_string_source, land_right_component.Project_id, land_right_component.Rate_schedule_id, land_right_component.Referenced_guid, land_right_component.Referenced_system_id, land_right_component.Referenced_table_name, land_right_component.Remark, land_right_component.Resent_id, land_right_component.Reserve_class_id, land_right_component.Sample_anal_source, land_right_component.Seis_set_id, land_right_component.Seis_set_subtype, land_right_component.Sf_subtype, land_right_component.Source, land_right_component.Spatial_description_id, land_right_component.Spatial_obs_no, land_right_component.Store_id, land_right_component.Store_structure_obs_no, land_right_component.Strat_name_set_id, land_right_component.Strat_unit_id, land_right_component.Support_facility_id, land_right_component.Sw_application_id, land_right_component.Thesaurus_id, land_right_component.Thesaurus_word_id, land_right_component.Uwi, land_right_component.Well_activity_set_id, land_right_component.Well_activity_source, land_right_component.Well_activity_type_id, land_right_component.Well_set_id, land_right_component.Work_order_id, land_right_component.Row_changed_by, land_right_component.Row_changed_date, land_right_component.Row_created_by, land_right_component.Row_created_date, land_right_component.Row_effective_date, land_right_component.Row_expiry_date, land_right_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandRightComponent(c *fiber.Ctx) error {
	var land_right_component dto.Land_right_component

	setDefaults(&land_right_component)

	if err := json.Unmarshal(c.Body(), &land_right_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_right_component.Land_right_subtype = id

    if land_right_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_right_component set land_right_id = :1, component_obs_no = :2, active_ind = :3, activity_obs_no = :4, analysis_id = :5, application_id = :6, area_id = :7, area_type = :8, authority_id = :9, ba_organization_id = :10, ba_organization_seq_no = :11, business_associate_id = :12, catalogue_additive_id = :13, catalogue_equip_id = :14, classification_system_id = :15, class_level_id = :16, consent_id = :17, consult_id = :18, contest_id = :19, contract_id = :20, ecozone_id = :21, effective_date = :22, employee_ba_id = :23, employee_obs_no = :24, employer_ba_id = :25, entitlement_id = :26, equipment_id = :27, expiry_date = :28, facility_id = :29, facility_type = :30, field_id = :31, field_station_id = :32, finance_id = :33, fossil_id = :34, incident_id = :35, incident_set_id = :36, incident_type = :37, information_item_id = :38, info_item_subtype = :39, instrument_id = :40, interest_set_id = :41, interest_set_seq_no = :42, jurisdiction = :43, land_component_type = :44, land_sale_number = :45, language = :46, lithology_log_id = :47, lith_log_source = :48, notification_id = :49, obligation_id = :50, obligation_seq_no = :51, paleo_summary_id = :52, pden_id = :53, pden_source = :54, pden_subtype = :55, physical_item_id = :56, pool_id = :57, ppdm_guid = :58, ppdm_system_id = :59, ppdm_table_name = :60, product_type = :61, prod_string_id = :62, prod_string_source = :63, project_id = :64, rate_schedule_id = :65, referenced_guid = :66, referenced_system_id = :67, referenced_table_name = :68, remark = :69, resent_id = :70, reserve_class_id = :71, sample_anal_source = :72, seis_set_id = :73, seis_set_subtype = :74, sf_subtype = :75, source = :76, spatial_description_id = :77, spatial_obs_no = :78, store_id = :79, store_structure_obs_no = :80, strat_name_set_id = :81, strat_unit_id = :82, support_facility_id = :83, sw_application_id = :84, thesaurus_id = :85, thesaurus_word_id = :86, uwi = :87, well_activity_set_id = :88, well_activity_source = :89, well_activity_type_id = :90, well_set_id = :91, work_order_id = :92, row_changed_by = :93, row_changed_date = :94, row_created_by = :95, row_effective_date = :96, row_expiry_date = :97, row_quality = :98 where land_right_subtype = :100")
	if err != nil {
		return err
	}

	land_right_component.Row_changed_date = formatDate(land_right_component.Row_changed_date)
	land_right_component.Effective_date = formatDateString(land_right_component.Effective_date)
	land_right_component.Expiry_date = formatDateString(land_right_component.Expiry_date)
	land_right_component.Row_effective_date = formatDateString(land_right_component.Row_effective_date)
	land_right_component.Row_expiry_date = formatDateString(land_right_component.Row_expiry_date)






	rows, err := stmt.Exec(land_right_component.Land_right_id, land_right_component.Component_obs_no, land_right_component.Active_ind, land_right_component.Activity_obs_no, land_right_component.Analysis_id, land_right_component.Application_id, land_right_component.Area_id, land_right_component.Area_type, land_right_component.Authority_id, land_right_component.Ba_organization_id, land_right_component.Ba_organization_seq_no, land_right_component.Business_associate_id, land_right_component.Catalogue_additive_id, land_right_component.Catalogue_equip_id, land_right_component.Classification_system_id, land_right_component.Class_level_id, land_right_component.Consent_id, land_right_component.Consult_id, land_right_component.Contest_id, land_right_component.Contract_id, land_right_component.Ecozone_id, land_right_component.Effective_date, land_right_component.Employee_ba_id, land_right_component.Employee_obs_no, land_right_component.Employer_ba_id, land_right_component.Entitlement_id, land_right_component.Equipment_id, land_right_component.Expiry_date, land_right_component.Facility_id, land_right_component.Facility_type, land_right_component.Field_id, land_right_component.Field_station_id, land_right_component.Finance_id, land_right_component.Fossil_id, land_right_component.Incident_id, land_right_component.Incident_set_id, land_right_component.Incident_type, land_right_component.Information_item_id, land_right_component.Info_item_subtype, land_right_component.Instrument_id, land_right_component.Interest_set_id, land_right_component.Interest_set_seq_no, land_right_component.Jurisdiction, land_right_component.Land_component_type, land_right_component.Land_sale_number, land_right_component.Language, land_right_component.Lithology_log_id, land_right_component.Lith_log_source, land_right_component.Notification_id, land_right_component.Obligation_id, land_right_component.Obligation_seq_no, land_right_component.Paleo_summary_id, land_right_component.Pden_id, land_right_component.Pden_source, land_right_component.Pden_subtype, land_right_component.Physical_item_id, land_right_component.Pool_id, land_right_component.Ppdm_guid, land_right_component.Ppdm_system_id, land_right_component.Ppdm_table_name, land_right_component.Product_type, land_right_component.Prod_string_id, land_right_component.Prod_string_source, land_right_component.Project_id, land_right_component.Rate_schedule_id, land_right_component.Referenced_guid, land_right_component.Referenced_system_id, land_right_component.Referenced_table_name, land_right_component.Remark, land_right_component.Resent_id, land_right_component.Reserve_class_id, land_right_component.Sample_anal_source, land_right_component.Seis_set_id, land_right_component.Seis_set_subtype, land_right_component.Sf_subtype, land_right_component.Source, land_right_component.Spatial_description_id, land_right_component.Spatial_obs_no, land_right_component.Store_id, land_right_component.Store_structure_obs_no, land_right_component.Strat_name_set_id, land_right_component.Strat_unit_id, land_right_component.Support_facility_id, land_right_component.Sw_application_id, land_right_component.Thesaurus_id, land_right_component.Thesaurus_word_id, land_right_component.Uwi, land_right_component.Well_activity_set_id, land_right_component.Well_activity_source, land_right_component.Well_activity_type_id, land_right_component.Well_set_id, land_right_component.Work_order_id, land_right_component.Row_changed_by, land_right_component.Row_changed_date, land_right_component.Row_created_by, land_right_component.Row_effective_date, land_right_component.Row_expiry_date, land_right_component.Row_quality, land_right_component.Land_right_subtype)
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

func PatchLandRightComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_right_component set "
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
	query += " where land_right_subtype = :id"

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

func DeleteLandRightComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_right_component dto.Land_right_component
	land_right_component.Land_right_subtype = id

	stmt, err := db.Prepare("delete from land_right_component where land_right_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_right_component.Land_right_subtype)
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


