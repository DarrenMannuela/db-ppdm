package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWorkOrderComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from work_order_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Work_order_component

	for rows.Next() {
		var work_order_component dto.Work_order_component
		if err := rows.Scan(&work_order_component.Work_order_id, &work_order_component.Component_id, &work_order_component.Active_ind, &work_order_component.Activity_obs_no, &work_order_component.Analysis_id, &work_order_component.Application_id, &work_order_component.Area_id, &work_order_component.Area_type, &work_order_component.Authority_id, &work_order_component.Ba_organization_id, &work_order_component.Ba_organization_seq_no, &work_order_component.Business_associate_id, &work_order_component.Catalogue_additive_id, &work_order_component.Catalogue_equip_id, &work_order_component.Circ_id, &work_order_component.Classification_system_id, &work_order_component.Class_level_id, &work_order_component.Component_type, &work_order_component.Consent_id, &work_order_component.Consult_id, &work_order_component.Contest_id, &work_order_component.Contract_id, &work_order_component.Ecozone_id, &work_order_component.Effective_date, &work_order_component.Employee_ba_id, &work_order_component.Employee_obs_no, &work_order_component.Employer_ba_id, &work_order_component.Entitlement_id, &work_order_component.Equipment_id, &work_order_component.Equip_maint_id, &work_order_component.Expiry_date, &work_order_component.Facility_id, &work_order_component.Facility_maintain_id, &work_order_component.Facility_type, &work_order_component.Field_id, &work_order_component.Field_station_id, &work_order_component.Finance_id, &work_order_component.Fossil_id, &work_order_component.Incident_id, &work_order_component.Incident_set_id, &work_order_component.Incident_type, &work_order_component.Information_item_id, &work_order_component.Info_item_subtype, &work_order_component.Inspection_id, &work_order_component.Instrument_id, &work_order_component.Interest_set_id, &work_order_component.Interest_set_seq_no, &work_order_component.Jurisdiction, &work_order_component.Land_right_id, &work_order_component.Land_right_subtype, &work_order_component.Land_sale_number, &work_order_component.Language, &work_order_component.Lithology_log_id, &work_order_component.Lith_log_source, &work_order_component.Notification_id, &work_order_component.Obligation_id, &work_order_component.Obligation_seq_no, &work_order_component.Origin_seq_no, &work_order_component.Paleo_summary_id, &work_order_component.Pden_id, &work_order_component.Pden_source, &work_order_component.Pden_subtype, &work_order_component.Physical_item_id, &work_order_component.Pool_id, &work_order_component.Ppdm_guid, &work_order_component.Ppdm_system_id, &work_order_component.Ppdm_table_name, &work_order_component.Product_type, &work_order_component.Prod_string_id, &work_order_component.Prod_string_source, &work_order_component.Project_id, &work_order_component.Rate_schedule_id, &work_order_component.Referenced_guid, &work_order_component.Referenced_system_id, &work_order_component.Referenced_table_name, &work_order_component.Remark, &work_order_component.Resent_id, &work_order_component.Reserve_class_id, &work_order_component.Sample_anal_source, &work_order_component.Seis_set_id, &work_order_component.Seis_set_subtype, &work_order_component.Seis_transaction_id, &work_order_component.Sf_subtype, &work_order_component.Source, &work_order_component.Spatial_description_id, &work_order_component.Spatial_obs_no, &work_order_component.Store_id, &work_order_component.Store_structure_obs_no, &work_order_component.Strat_name_set_id, &work_order_component.Strat_unit_id, &work_order_component.Support_facility_id, &work_order_component.Sw_application_id, &work_order_component.Thesaurus_id, &work_order_component.Thesaurus_word_id, &work_order_component.Transaction_type, &work_order_component.Uwi, &work_order_component.Well_activity_set_id, &work_order_component.Well_activity_source, &work_order_component.Well_activity_type_id, &work_order_component.Well_set_id, &work_order_component.Row_changed_by, &work_order_component.Row_changed_date, &work_order_component.Row_created_by, &work_order_component.Row_created_date, &work_order_component.Row_effective_date, &work_order_component.Row_expiry_date, &work_order_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append work_order_component to result
		result = append(result, work_order_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWorkOrderComponent(c *fiber.Ctx) error {
	var work_order_component dto.Work_order_component

	setDefaults(&work_order_component)

	if err := json.Unmarshal(c.Body(), &work_order_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into work_order_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103, :104, :105, :106, :107)")
	if err != nil {
		return err
	}
	work_order_component.Row_created_date = formatDate(work_order_component.Row_created_date)
	work_order_component.Row_changed_date = formatDate(work_order_component.Row_changed_date)
	work_order_component.Effective_date = formatDateString(work_order_component.Effective_date)
	work_order_component.Expiry_date = formatDateString(work_order_component.Expiry_date)
	work_order_component.Row_effective_date = formatDateString(work_order_component.Row_effective_date)
	work_order_component.Row_expiry_date = formatDateString(work_order_component.Row_expiry_date)






	rows, err := stmt.Exec(work_order_component.Work_order_id, work_order_component.Component_id, work_order_component.Active_ind, work_order_component.Activity_obs_no, work_order_component.Analysis_id, work_order_component.Application_id, work_order_component.Area_id, work_order_component.Area_type, work_order_component.Authority_id, work_order_component.Ba_organization_id, work_order_component.Ba_organization_seq_no, work_order_component.Business_associate_id, work_order_component.Catalogue_additive_id, work_order_component.Catalogue_equip_id, work_order_component.Circ_id, work_order_component.Classification_system_id, work_order_component.Class_level_id, work_order_component.Component_type, work_order_component.Consent_id, work_order_component.Consult_id, work_order_component.Contest_id, work_order_component.Contract_id, work_order_component.Ecozone_id, work_order_component.Effective_date, work_order_component.Employee_ba_id, work_order_component.Employee_obs_no, work_order_component.Employer_ba_id, work_order_component.Entitlement_id, work_order_component.Equipment_id, work_order_component.Equip_maint_id, work_order_component.Expiry_date, work_order_component.Facility_id, work_order_component.Facility_maintain_id, work_order_component.Facility_type, work_order_component.Field_id, work_order_component.Field_station_id, work_order_component.Finance_id, work_order_component.Fossil_id, work_order_component.Incident_id, work_order_component.Incident_set_id, work_order_component.Incident_type, work_order_component.Information_item_id, work_order_component.Info_item_subtype, work_order_component.Inspection_id, work_order_component.Instrument_id, work_order_component.Interest_set_id, work_order_component.Interest_set_seq_no, work_order_component.Jurisdiction, work_order_component.Land_right_id, work_order_component.Land_right_subtype, work_order_component.Land_sale_number, work_order_component.Language, work_order_component.Lithology_log_id, work_order_component.Lith_log_source, work_order_component.Notification_id, work_order_component.Obligation_id, work_order_component.Obligation_seq_no, work_order_component.Origin_seq_no, work_order_component.Paleo_summary_id, work_order_component.Pden_id, work_order_component.Pden_source, work_order_component.Pden_subtype, work_order_component.Physical_item_id, work_order_component.Pool_id, work_order_component.Ppdm_guid, work_order_component.Ppdm_system_id, work_order_component.Ppdm_table_name, work_order_component.Product_type, work_order_component.Prod_string_id, work_order_component.Prod_string_source, work_order_component.Project_id, work_order_component.Rate_schedule_id, work_order_component.Referenced_guid, work_order_component.Referenced_system_id, work_order_component.Referenced_table_name, work_order_component.Remark, work_order_component.Resent_id, work_order_component.Reserve_class_id, work_order_component.Sample_anal_source, work_order_component.Seis_set_id, work_order_component.Seis_set_subtype, work_order_component.Seis_transaction_id, work_order_component.Sf_subtype, work_order_component.Source, work_order_component.Spatial_description_id, work_order_component.Spatial_obs_no, work_order_component.Store_id, work_order_component.Store_structure_obs_no, work_order_component.Strat_name_set_id, work_order_component.Strat_unit_id, work_order_component.Support_facility_id, work_order_component.Sw_application_id, work_order_component.Thesaurus_id, work_order_component.Thesaurus_word_id, work_order_component.Transaction_type, work_order_component.Uwi, work_order_component.Well_activity_set_id, work_order_component.Well_activity_source, work_order_component.Well_activity_type_id, work_order_component.Well_set_id, work_order_component.Row_changed_by, work_order_component.Row_changed_date, work_order_component.Row_created_by, work_order_component.Row_created_date, work_order_component.Row_effective_date, work_order_component.Row_expiry_date, work_order_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWorkOrderComponent(c *fiber.Ctx) error {
	var work_order_component dto.Work_order_component

	setDefaults(&work_order_component)

	if err := json.Unmarshal(c.Body(), &work_order_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	work_order_component.Work_order_id = id

    if work_order_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update work_order_component set component_id = :1, active_ind = :2, activity_obs_no = :3, analysis_id = :4, application_id = :5, area_id = :6, area_type = :7, authority_id = :8, ba_organization_id = :9, ba_organization_seq_no = :10, business_associate_id = :11, catalogue_additive_id = :12, catalogue_equip_id = :13, circ_id = :14, classification_system_id = :15, class_level_id = :16, component_type = :17, consent_id = :18, consult_id = :19, contest_id = :20, contract_id = :21, ecozone_id = :22, effective_date = :23, employee_ba_id = :24, employee_obs_no = :25, employer_ba_id = :26, entitlement_id = :27, equipment_id = :28, equip_maint_id = :29, expiry_date = :30, facility_id = :31, facility_maintain_id = :32, facility_type = :33, field_id = :34, field_station_id = :35, finance_id = :36, fossil_id = :37, incident_id = :38, incident_set_id = :39, incident_type = :40, information_item_id = :41, info_item_subtype = :42, inspection_id = :43, instrument_id = :44, interest_set_id = :45, interest_set_seq_no = :46, jurisdiction = :47, land_right_id = :48, land_right_subtype = :49, land_sale_number = :50, language = :51, lithology_log_id = :52, lith_log_source = :53, notification_id = :54, obligation_id = :55, obligation_seq_no = :56, origin_seq_no = :57, paleo_summary_id = :58, pden_id = :59, pden_source = :60, pden_subtype = :61, physical_item_id = :62, pool_id = :63, ppdm_guid = :64, ppdm_system_id = :65, ppdm_table_name = :66, product_type = :67, prod_string_id = :68, prod_string_source = :69, project_id = :70, rate_schedule_id = :71, referenced_guid = :72, referenced_system_id = :73, referenced_table_name = :74, remark = :75, resent_id = :76, reserve_class_id = :77, sample_anal_source = :78, seis_set_id = :79, seis_set_subtype = :80, seis_transaction_id = :81, sf_subtype = :82, source = :83, spatial_description_id = :84, spatial_obs_no = :85, store_id = :86, store_structure_obs_no = :87, strat_name_set_id = :88, strat_unit_id = :89, support_facility_id = :90, sw_application_id = :91, thesaurus_id = :92, thesaurus_word_id = :93, transaction_type = :94, uwi = :95, well_activity_set_id = :96, well_activity_source = :97, well_activity_type_id = :98, well_set_id = :99, row_changed_by = :100, row_changed_date = :101, row_created_by = :102, row_effective_date = :103, row_expiry_date = :104, row_quality = :105 where work_order_id = :107")
	if err != nil {
		return err
	}

	work_order_component.Row_changed_date = formatDate(work_order_component.Row_changed_date)
	work_order_component.Effective_date = formatDateString(work_order_component.Effective_date)
	work_order_component.Expiry_date = formatDateString(work_order_component.Expiry_date)
	work_order_component.Row_effective_date = formatDateString(work_order_component.Row_effective_date)
	work_order_component.Row_expiry_date = formatDateString(work_order_component.Row_expiry_date)






	rows, err := stmt.Exec(work_order_component.Component_id, work_order_component.Active_ind, work_order_component.Activity_obs_no, work_order_component.Analysis_id, work_order_component.Application_id, work_order_component.Area_id, work_order_component.Area_type, work_order_component.Authority_id, work_order_component.Ba_organization_id, work_order_component.Ba_organization_seq_no, work_order_component.Business_associate_id, work_order_component.Catalogue_additive_id, work_order_component.Catalogue_equip_id, work_order_component.Circ_id, work_order_component.Classification_system_id, work_order_component.Class_level_id, work_order_component.Component_type, work_order_component.Consent_id, work_order_component.Consult_id, work_order_component.Contest_id, work_order_component.Contract_id, work_order_component.Ecozone_id, work_order_component.Effective_date, work_order_component.Employee_ba_id, work_order_component.Employee_obs_no, work_order_component.Employer_ba_id, work_order_component.Entitlement_id, work_order_component.Equipment_id, work_order_component.Equip_maint_id, work_order_component.Expiry_date, work_order_component.Facility_id, work_order_component.Facility_maintain_id, work_order_component.Facility_type, work_order_component.Field_id, work_order_component.Field_station_id, work_order_component.Finance_id, work_order_component.Fossil_id, work_order_component.Incident_id, work_order_component.Incident_set_id, work_order_component.Incident_type, work_order_component.Information_item_id, work_order_component.Info_item_subtype, work_order_component.Inspection_id, work_order_component.Instrument_id, work_order_component.Interest_set_id, work_order_component.Interest_set_seq_no, work_order_component.Jurisdiction, work_order_component.Land_right_id, work_order_component.Land_right_subtype, work_order_component.Land_sale_number, work_order_component.Language, work_order_component.Lithology_log_id, work_order_component.Lith_log_source, work_order_component.Notification_id, work_order_component.Obligation_id, work_order_component.Obligation_seq_no, work_order_component.Origin_seq_no, work_order_component.Paleo_summary_id, work_order_component.Pden_id, work_order_component.Pden_source, work_order_component.Pden_subtype, work_order_component.Physical_item_id, work_order_component.Pool_id, work_order_component.Ppdm_guid, work_order_component.Ppdm_system_id, work_order_component.Ppdm_table_name, work_order_component.Product_type, work_order_component.Prod_string_id, work_order_component.Prod_string_source, work_order_component.Project_id, work_order_component.Rate_schedule_id, work_order_component.Referenced_guid, work_order_component.Referenced_system_id, work_order_component.Referenced_table_name, work_order_component.Remark, work_order_component.Resent_id, work_order_component.Reserve_class_id, work_order_component.Sample_anal_source, work_order_component.Seis_set_id, work_order_component.Seis_set_subtype, work_order_component.Seis_transaction_id, work_order_component.Sf_subtype, work_order_component.Source, work_order_component.Spatial_description_id, work_order_component.Spatial_obs_no, work_order_component.Store_id, work_order_component.Store_structure_obs_no, work_order_component.Strat_name_set_id, work_order_component.Strat_unit_id, work_order_component.Support_facility_id, work_order_component.Sw_application_id, work_order_component.Thesaurus_id, work_order_component.Thesaurus_word_id, work_order_component.Transaction_type, work_order_component.Uwi, work_order_component.Well_activity_set_id, work_order_component.Well_activity_source, work_order_component.Well_activity_type_id, work_order_component.Well_set_id, work_order_component.Row_changed_by, work_order_component.Row_changed_date, work_order_component.Row_created_by, work_order_component.Row_effective_date, work_order_component.Row_expiry_date, work_order_component.Row_quality, work_order_component.Work_order_id)
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

func PatchWorkOrderComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update work_order_component set "
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
	query += " where work_order_id = :id"

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

func DeleteWorkOrderComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var work_order_component dto.Work_order_component
	work_order_component.Work_order_id = id

	stmt, err := db.Prepare("delete from work_order_component where work_order_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(work_order_component.Work_order_id)
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


