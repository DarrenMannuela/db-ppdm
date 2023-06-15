package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetEntComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ent_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ent_component

	for rows.Next() {
		var ent_component dto.Ent_component
		if err := rows.Scan(&ent_component.Entitlement_id, &ent_component.Component_id, &ent_component.Component_obs_no, &ent_component.Access_cond_code, &ent_component.Active_ind, &ent_component.Activity_obs_no, &ent_component.Analysis_id, &ent_component.Application_id, &ent_component.Area_id, &ent_component.Area_type, &ent_component.Authority_id, &ent_component.Business_associate_id, &ent_component.Catalogue_additive_id, &ent_component.Catalogue_equip_id, &ent_component.Classification_system_id, &ent_component.Class_level_id, &ent_component.Consent_id, &ent_component.Consult_id, &ent_component.Contest_id, &ent_component.Contract_id, &ent_component.Description, &ent_component.Ecozone_id, &ent_component.Effective_date, &ent_component.Employee_ba_id, &ent_component.Employee_obs_no, &ent_component.Employer_ba_id, &ent_component.Ent_comp_type, &ent_component.Ent_rule, &ent_component.Equipment_id, &ent_component.Expiry_action, &ent_component.Expiry_date, &ent_component.Facility_id, &ent_component.Facility_type, &ent_component.Field_id, &ent_component.Field_station_id, &ent_component.Finance_id, &ent_component.Fossil_id, &ent_component.Incident_id, &ent_component.Incident_set_id, &ent_component.Incident_type, &ent_component.Information_item_id, &ent_component.Info_item_subtype, &ent_component.Instrument_id, &ent_component.Interest_set_id, &ent_component.Interest_set_seq_no, &ent_component.Jurisdiction, &ent_component.Land_right_id, &ent_component.Land_right_subtype, &ent_component.Land_sale_number, &ent_component.Language, &ent_component.Lithology_log_id, &ent_component.Lith_log_source, &ent_component.Notification_id, &ent_component.Obligation_id, &ent_component.Obligation_seq_no, &ent_component.Organization_id, &ent_component.Organization_seq_no, &ent_component.Paleo_summary_id, &ent_component.Pden_id, &ent_component.Pden_source, &ent_component.Pden_subtype, &ent_component.Physical_item_id, &ent_component.Pool_id, &ent_component.Ppdm_guid, &ent_component.Ppdm_system_id, &ent_component.Ppdm_table_name, &ent_component.Product_type, &ent_component.Prod_string_id, &ent_component.Prod_string_source, &ent_component.Project_id, &ent_component.Rate_schedule_id, &ent_component.Referenced_guid, &ent_component.Referenced_system_id, &ent_component.Referenced_table_name, &ent_component.Remark, &ent_component.Report_hierarchy_id, &ent_component.Report_hier_component_id, &ent_component.Resent_id, &ent_component.Reserve_class_id, &ent_component.Sample_anal_source, &ent_component.Seis_set_id, &ent_component.Seis_set_subtype, &ent_component.Seis_transaction_id, &ent_component.Sf_subtype, &ent_component.Source, &ent_component.Spatial_description_id, &ent_component.Spatial_obs_no, &ent_component.Store_id, &ent_component.Store_structure_obs_no, &ent_component.Strat_name_set_id, &ent_component.Strat_unit_id, &ent_component.Support_facility_id, &ent_component.Sw_application_id, &ent_component.Thesaurus_id, &ent_component.Thesaurus_word_id, &ent_component.Transaction_type, &ent_component.Uwi, &ent_component.Well_activity_set_id, &ent_component.Well_activity_source, &ent_component.Well_activity_type_id, &ent_component.Well_log_curve_id, &ent_component.Well_log_id, &ent_component.Well_log_source, &ent_component.Well_set_id, &ent_component.Work_order_id, &ent_component.Row_changed_by, &ent_component.Row_changed_date, &ent_component.Row_created_by, &ent_component.Row_created_date, &ent_component.Row_effective_date, &ent_component.Row_expiry_date, &ent_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ent_component to result
		result = append(result, ent_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetEntComponent(c *fiber.Ctx) error {
	var ent_component dto.Ent_component

	setDefaults(&ent_component)

	if err := json.Unmarshal(c.Body(), &ent_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ent_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103, :104, :105, :106, :107, :108, :109, :110, :111, :112)")
	if err != nil {
		return err
	}
	ent_component.Row_created_date = formatDate(ent_component.Row_created_date)
	ent_component.Row_changed_date = formatDate(ent_component.Row_changed_date)
	ent_component.Effective_date = formatDateString(ent_component.Effective_date)
	ent_component.Expiry_date = formatDateString(ent_component.Expiry_date)
	ent_component.Row_effective_date = formatDateString(ent_component.Row_effective_date)
	ent_component.Row_expiry_date = formatDateString(ent_component.Row_expiry_date)






	rows, err := stmt.Exec(ent_component.Entitlement_id, ent_component.Component_id, ent_component.Component_obs_no, ent_component.Access_cond_code, ent_component.Active_ind, ent_component.Activity_obs_no, ent_component.Analysis_id, ent_component.Application_id, ent_component.Area_id, ent_component.Area_type, ent_component.Authority_id, ent_component.Business_associate_id, ent_component.Catalogue_additive_id, ent_component.Catalogue_equip_id, ent_component.Classification_system_id, ent_component.Class_level_id, ent_component.Consent_id, ent_component.Consult_id, ent_component.Contest_id, ent_component.Contract_id, ent_component.Description, ent_component.Ecozone_id, ent_component.Effective_date, ent_component.Employee_ba_id, ent_component.Employee_obs_no, ent_component.Employer_ba_id, ent_component.Ent_comp_type, ent_component.Ent_rule, ent_component.Equipment_id, ent_component.Expiry_action, ent_component.Expiry_date, ent_component.Facility_id, ent_component.Facility_type, ent_component.Field_id, ent_component.Field_station_id, ent_component.Finance_id, ent_component.Fossil_id, ent_component.Incident_id, ent_component.Incident_set_id, ent_component.Incident_type, ent_component.Information_item_id, ent_component.Info_item_subtype, ent_component.Instrument_id, ent_component.Interest_set_id, ent_component.Interest_set_seq_no, ent_component.Jurisdiction, ent_component.Land_right_id, ent_component.Land_right_subtype, ent_component.Land_sale_number, ent_component.Language, ent_component.Lithology_log_id, ent_component.Lith_log_source, ent_component.Notification_id, ent_component.Obligation_id, ent_component.Obligation_seq_no, ent_component.Organization_id, ent_component.Organization_seq_no, ent_component.Paleo_summary_id, ent_component.Pden_id, ent_component.Pden_source, ent_component.Pden_subtype, ent_component.Physical_item_id, ent_component.Pool_id, ent_component.Ppdm_guid, ent_component.Ppdm_system_id, ent_component.Ppdm_table_name, ent_component.Product_type, ent_component.Prod_string_id, ent_component.Prod_string_source, ent_component.Project_id, ent_component.Rate_schedule_id, ent_component.Referenced_guid, ent_component.Referenced_system_id, ent_component.Referenced_table_name, ent_component.Remark, ent_component.Report_hierarchy_id, ent_component.Report_hier_component_id, ent_component.Resent_id, ent_component.Reserve_class_id, ent_component.Sample_anal_source, ent_component.Seis_set_id, ent_component.Seis_set_subtype, ent_component.Seis_transaction_id, ent_component.Sf_subtype, ent_component.Source, ent_component.Spatial_description_id, ent_component.Spatial_obs_no, ent_component.Store_id, ent_component.Store_structure_obs_no, ent_component.Strat_name_set_id, ent_component.Strat_unit_id, ent_component.Support_facility_id, ent_component.Sw_application_id, ent_component.Thesaurus_id, ent_component.Thesaurus_word_id, ent_component.Transaction_type, ent_component.Uwi, ent_component.Well_activity_set_id, ent_component.Well_activity_source, ent_component.Well_activity_type_id, ent_component.Well_log_curve_id, ent_component.Well_log_id, ent_component.Well_log_source, ent_component.Well_set_id, ent_component.Work_order_id, ent_component.Row_changed_by, ent_component.Row_changed_date, ent_component.Row_created_by, ent_component.Row_created_date, ent_component.Row_effective_date, ent_component.Row_expiry_date, ent_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateEntComponent(c *fiber.Ctx) error {
	var ent_component dto.Ent_component

	setDefaults(&ent_component)

	if err := json.Unmarshal(c.Body(), &ent_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ent_component.Entitlement_id = id

    if ent_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ent_component set component_id = :1, component_obs_no = :2, access_cond_code = :3, active_ind = :4, activity_obs_no = :5, analysis_id = :6, application_id = :7, area_id = :8, area_type = :9, authority_id = :10, business_associate_id = :11, catalogue_additive_id = :12, catalogue_equip_id = :13, classification_system_id = :14, class_level_id = :15, consent_id = :16, consult_id = :17, contest_id = :18, contract_id = :19, description = :20, ecozone_id = :21, effective_date = :22, employee_ba_id = :23, employee_obs_no = :24, employer_ba_id = :25, ent_comp_type = :26, ent_rule = :27, equipment_id = :28, expiry_action = :29, expiry_date = :30, facility_id = :31, facility_type = :32, field_id = :33, field_station_id = :34, finance_id = :35, fossil_id = :36, incident_id = :37, incident_set_id = :38, incident_type = :39, information_item_id = :40, info_item_subtype = :41, instrument_id = :42, interest_set_id = :43, interest_set_seq_no = :44, jurisdiction = :45, land_right_id = :46, land_right_subtype = :47, land_sale_number = :48, language = :49, lithology_log_id = :50, lith_log_source = :51, notification_id = :52, obligation_id = :53, obligation_seq_no = :54, organization_id = :55, organization_seq_no = :56, paleo_summary_id = :57, pden_id = :58, pden_source = :59, pden_subtype = :60, physical_item_id = :61, pool_id = :62, ppdm_guid = :63, ppdm_system_id = :64, ppdm_table_name = :65, product_type = :66, prod_string_id = :67, prod_string_source = :68, project_id = :69, rate_schedule_id = :70, referenced_guid = :71, referenced_system_id = :72, referenced_table_name = :73, remark = :74, report_hierarchy_id = :75, report_hier_component_id = :76, resent_id = :77, reserve_class_id = :78, sample_anal_source = :79, seis_set_id = :80, seis_set_subtype = :81, seis_transaction_id = :82, sf_subtype = :83, source = :84, spatial_description_id = :85, spatial_obs_no = :86, store_id = :87, store_structure_obs_no = :88, strat_name_set_id = :89, strat_unit_id = :90, support_facility_id = :91, sw_application_id = :92, thesaurus_id = :93, thesaurus_word_id = :94, transaction_type = :95, uwi = :96, well_activity_set_id = :97, well_activity_source = :98, well_activity_type_id = :99, well_log_curve_id = :100, well_log_id = :101, well_log_source = :102, well_set_id = :103, work_order_id = :104, row_changed_by = :105, row_changed_date = :106, row_created_by = :107, row_effective_date = :108, row_expiry_date = :109, row_quality = :110 where entitlement_id = :112")
	if err != nil {
		return err
	}

	ent_component.Row_changed_date = formatDate(ent_component.Row_changed_date)
	ent_component.Effective_date = formatDateString(ent_component.Effective_date)
	ent_component.Expiry_date = formatDateString(ent_component.Expiry_date)
	ent_component.Row_effective_date = formatDateString(ent_component.Row_effective_date)
	ent_component.Row_expiry_date = formatDateString(ent_component.Row_expiry_date)






	rows, err := stmt.Exec(ent_component.Component_id, ent_component.Component_obs_no, ent_component.Access_cond_code, ent_component.Active_ind, ent_component.Activity_obs_no, ent_component.Analysis_id, ent_component.Application_id, ent_component.Area_id, ent_component.Area_type, ent_component.Authority_id, ent_component.Business_associate_id, ent_component.Catalogue_additive_id, ent_component.Catalogue_equip_id, ent_component.Classification_system_id, ent_component.Class_level_id, ent_component.Consent_id, ent_component.Consult_id, ent_component.Contest_id, ent_component.Contract_id, ent_component.Description, ent_component.Ecozone_id, ent_component.Effective_date, ent_component.Employee_ba_id, ent_component.Employee_obs_no, ent_component.Employer_ba_id, ent_component.Ent_comp_type, ent_component.Ent_rule, ent_component.Equipment_id, ent_component.Expiry_action, ent_component.Expiry_date, ent_component.Facility_id, ent_component.Facility_type, ent_component.Field_id, ent_component.Field_station_id, ent_component.Finance_id, ent_component.Fossil_id, ent_component.Incident_id, ent_component.Incident_set_id, ent_component.Incident_type, ent_component.Information_item_id, ent_component.Info_item_subtype, ent_component.Instrument_id, ent_component.Interest_set_id, ent_component.Interest_set_seq_no, ent_component.Jurisdiction, ent_component.Land_right_id, ent_component.Land_right_subtype, ent_component.Land_sale_number, ent_component.Language, ent_component.Lithology_log_id, ent_component.Lith_log_source, ent_component.Notification_id, ent_component.Obligation_id, ent_component.Obligation_seq_no, ent_component.Organization_id, ent_component.Organization_seq_no, ent_component.Paleo_summary_id, ent_component.Pden_id, ent_component.Pden_source, ent_component.Pden_subtype, ent_component.Physical_item_id, ent_component.Pool_id, ent_component.Ppdm_guid, ent_component.Ppdm_system_id, ent_component.Ppdm_table_name, ent_component.Product_type, ent_component.Prod_string_id, ent_component.Prod_string_source, ent_component.Project_id, ent_component.Rate_schedule_id, ent_component.Referenced_guid, ent_component.Referenced_system_id, ent_component.Referenced_table_name, ent_component.Remark, ent_component.Report_hierarchy_id, ent_component.Report_hier_component_id, ent_component.Resent_id, ent_component.Reserve_class_id, ent_component.Sample_anal_source, ent_component.Seis_set_id, ent_component.Seis_set_subtype, ent_component.Seis_transaction_id, ent_component.Sf_subtype, ent_component.Source, ent_component.Spatial_description_id, ent_component.Spatial_obs_no, ent_component.Store_id, ent_component.Store_structure_obs_no, ent_component.Strat_name_set_id, ent_component.Strat_unit_id, ent_component.Support_facility_id, ent_component.Sw_application_id, ent_component.Thesaurus_id, ent_component.Thesaurus_word_id, ent_component.Transaction_type, ent_component.Uwi, ent_component.Well_activity_set_id, ent_component.Well_activity_source, ent_component.Well_activity_type_id, ent_component.Well_log_curve_id, ent_component.Well_log_id, ent_component.Well_log_source, ent_component.Well_set_id, ent_component.Work_order_id, ent_component.Row_changed_by, ent_component.Row_changed_date, ent_component.Row_created_by, ent_component.Row_effective_date, ent_component.Row_expiry_date, ent_component.Row_quality, ent_component.Entitlement_id)
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

func PatchEntComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ent_component set "
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
	query += " where entitlement_id = :id"

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

func DeleteEntComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var ent_component dto.Ent_component
	ent_component.Entitlement_id = id

	stmt, err := db.Prepare("delete from ent_component where entitlement_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ent_component.Entitlement_id)
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


