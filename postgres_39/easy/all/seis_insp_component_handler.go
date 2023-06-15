package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisInspComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_insp_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_insp_component

	for rows.Next() {
		var seis_insp_component dto.Seis_insp_component
		if err := rows.Scan(&seis_insp_component.Inspection_id, &seis_insp_component.Inspection_component_id, &seis_insp_component.Active_ind, &seis_insp_component.Activity_obs_no, &seis_insp_component.Analysis_id, &seis_insp_component.Application_id, &seis_insp_component.Area_id, &seis_insp_component.Area_type, &seis_insp_component.Authority_id, &seis_insp_component.Authorize_id, &seis_insp_component.Ba_organization_id, &seis_insp_component.Ba_organization_seq_no, &seis_insp_component.Business_associate_id, &seis_insp_component.Catalogue_additive_id, &seis_insp_component.Catalogue_equip_id, &seis_insp_component.Circ_id, &seis_insp_component.Classification_system_id, &seis_insp_component.Class_level_id, &seis_insp_component.Component_type, &seis_insp_component.Consent_id, &seis_insp_component.Consult_id, &seis_insp_component.Contest_id, &seis_insp_component.Contract_id, &seis_insp_component.Ecozone_id, &seis_insp_component.Effective_date, &seis_insp_component.Employee_ba_id, &seis_insp_component.Employee_obs_no, &seis_insp_component.Employer_ba_id, &seis_insp_component.Entitlement_id, &seis_insp_component.Equipment_id, &seis_insp_component.Expiry_date, &seis_insp_component.Facility_id, &seis_insp_component.Facility_type, &seis_insp_component.Field_id, &seis_insp_component.Field_station_id, &seis_insp_component.Finance_id, &seis_insp_component.Fossil_id, &seis_insp_component.Incident_id, &seis_insp_component.Incident_set_id, &seis_insp_component.Incident_type, &seis_insp_component.Information_item_id, &seis_insp_component.Info_item_subtype, &seis_insp_component.Inspection_status, &seis_insp_component.Instrument_id, &seis_insp_component.Interest_set_id, &seis_insp_component.Interest_set_seq_no, &seis_insp_component.Jurisdiction, &seis_insp_component.Land_right_id, &seis_insp_component.Land_right_subtype, &seis_insp_component.Land_sale_number, &seis_insp_component.Language, &seis_insp_component.Lithology_log_id, &seis_insp_component.Lith_log_source, &seis_insp_component.Notification_id, &seis_insp_component.Obligation_id, &seis_insp_component.Obligation_seq_no, &seis_insp_component.Paleo_summary_id, &seis_insp_component.Pden_id, &seis_insp_component.Pden_source, &seis_insp_component.Pden_subtype, &seis_insp_component.Physical_item_id, &seis_insp_component.Pool_id, &seis_insp_component.Ppdm_guid, &seis_insp_component.Ppdm_system_id, &seis_insp_component.Ppdm_table_name, &seis_insp_component.Product_type, &seis_insp_component.Prod_string_id, &seis_insp_component.Prod_string_source, &seis_insp_component.Project_id, &seis_insp_component.Rate_schedule_id, &seis_insp_component.Referenced_guid, &seis_insp_component.Referenced_system_id, &seis_insp_component.Referenced_table_name, &seis_insp_component.Remark, &seis_insp_component.Resent_id, &seis_insp_component.Reserve_class_id, &seis_insp_component.Sample_anal_source, &seis_insp_component.Seis_set_id, &seis_insp_component.Seis_set_subtype, &seis_insp_component.Sf_subtype, &seis_insp_component.Source, &seis_insp_component.Spatial_description_id, &seis_insp_component.Spatial_obs_no, &seis_insp_component.Store_id, &seis_insp_component.Store_structure_obs_no, &seis_insp_component.Strat_name_set_id, &seis_insp_component.Strat_unit_id, &seis_insp_component.Support_facility_id, &seis_insp_component.Sw_application_id, &seis_insp_component.Thesaurus_id, &seis_insp_component.Thesaurus_word_id, &seis_insp_component.Uwi, &seis_insp_component.Well_activity_set_id, &seis_insp_component.Well_activity_source, &seis_insp_component.Well_activity_type_id, &seis_insp_component.Well_set_id, &seis_insp_component.Work_order_id, &seis_insp_component.Row_changed_by, &seis_insp_component.Row_changed_date, &seis_insp_component.Row_created_by, &seis_insp_component.Row_created_date, &seis_insp_component.Row_effective_date, &seis_insp_component.Row_expiry_date, &seis_insp_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_insp_component to result
		result = append(result, seis_insp_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisInspComponent(c *fiber.Ctx) error {
	var seis_insp_component dto.Seis_insp_component

	setDefaults(&seis_insp_component)

	if err := json.Unmarshal(c.Body(), &seis_insp_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_insp_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103, :104)")
	if err != nil {
		return err
	}
	seis_insp_component.Row_created_date = formatDate(seis_insp_component.Row_created_date)
	seis_insp_component.Row_changed_date = formatDate(seis_insp_component.Row_changed_date)
	seis_insp_component.Effective_date = formatDateString(seis_insp_component.Effective_date)
	seis_insp_component.Expiry_date = formatDateString(seis_insp_component.Expiry_date)
	seis_insp_component.Row_effective_date = formatDateString(seis_insp_component.Row_effective_date)
	seis_insp_component.Row_expiry_date = formatDateString(seis_insp_component.Row_expiry_date)






	rows, err := stmt.Exec(seis_insp_component.Inspection_id, seis_insp_component.Inspection_component_id, seis_insp_component.Active_ind, seis_insp_component.Activity_obs_no, seis_insp_component.Analysis_id, seis_insp_component.Application_id, seis_insp_component.Area_id, seis_insp_component.Area_type, seis_insp_component.Authority_id, seis_insp_component.Authorize_id, seis_insp_component.Ba_organization_id, seis_insp_component.Ba_organization_seq_no, seis_insp_component.Business_associate_id, seis_insp_component.Catalogue_additive_id, seis_insp_component.Catalogue_equip_id, seis_insp_component.Circ_id, seis_insp_component.Classification_system_id, seis_insp_component.Class_level_id, seis_insp_component.Component_type, seis_insp_component.Consent_id, seis_insp_component.Consult_id, seis_insp_component.Contest_id, seis_insp_component.Contract_id, seis_insp_component.Ecozone_id, seis_insp_component.Effective_date, seis_insp_component.Employee_ba_id, seis_insp_component.Employee_obs_no, seis_insp_component.Employer_ba_id, seis_insp_component.Entitlement_id, seis_insp_component.Equipment_id, seis_insp_component.Expiry_date, seis_insp_component.Facility_id, seis_insp_component.Facility_type, seis_insp_component.Field_id, seis_insp_component.Field_station_id, seis_insp_component.Finance_id, seis_insp_component.Fossil_id, seis_insp_component.Incident_id, seis_insp_component.Incident_set_id, seis_insp_component.Incident_type, seis_insp_component.Information_item_id, seis_insp_component.Info_item_subtype, seis_insp_component.Inspection_status, seis_insp_component.Instrument_id, seis_insp_component.Interest_set_id, seis_insp_component.Interest_set_seq_no, seis_insp_component.Jurisdiction, seis_insp_component.Land_right_id, seis_insp_component.Land_right_subtype, seis_insp_component.Land_sale_number, seis_insp_component.Language, seis_insp_component.Lithology_log_id, seis_insp_component.Lith_log_source, seis_insp_component.Notification_id, seis_insp_component.Obligation_id, seis_insp_component.Obligation_seq_no, seis_insp_component.Paleo_summary_id, seis_insp_component.Pden_id, seis_insp_component.Pden_source, seis_insp_component.Pden_subtype, seis_insp_component.Physical_item_id, seis_insp_component.Pool_id, seis_insp_component.Ppdm_guid, seis_insp_component.Ppdm_system_id, seis_insp_component.Ppdm_table_name, seis_insp_component.Product_type, seis_insp_component.Prod_string_id, seis_insp_component.Prod_string_source, seis_insp_component.Project_id, seis_insp_component.Rate_schedule_id, seis_insp_component.Referenced_guid, seis_insp_component.Referenced_system_id, seis_insp_component.Referenced_table_name, seis_insp_component.Remark, seis_insp_component.Resent_id, seis_insp_component.Reserve_class_id, seis_insp_component.Sample_anal_source, seis_insp_component.Seis_set_id, seis_insp_component.Seis_set_subtype, seis_insp_component.Sf_subtype, seis_insp_component.Source, seis_insp_component.Spatial_description_id, seis_insp_component.Spatial_obs_no, seis_insp_component.Store_id, seis_insp_component.Store_structure_obs_no, seis_insp_component.Strat_name_set_id, seis_insp_component.Strat_unit_id, seis_insp_component.Support_facility_id, seis_insp_component.Sw_application_id, seis_insp_component.Thesaurus_id, seis_insp_component.Thesaurus_word_id, seis_insp_component.Uwi, seis_insp_component.Well_activity_set_id, seis_insp_component.Well_activity_source, seis_insp_component.Well_activity_type_id, seis_insp_component.Well_set_id, seis_insp_component.Work_order_id, seis_insp_component.Row_changed_by, seis_insp_component.Row_changed_date, seis_insp_component.Row_created_by, seis_insp_component.Row_created_date, seis_insp_component.Row_effective_date, seis_insp_component.Row_expiry_date, seis_insp_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisInspComponent(c *fiber.Ctx) error {
	var seis_insp_component dto.Seis_insp_component

	setDefaults(&seis_insp_component)

	if err := json.Unmarshal(c.Body(), &seis_insp_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_insp_component.Inspection_id = id

    if seis_insp_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_insp_component set inspection_component_id = :1, active_ind = :2, activity_obs_no = :3, analysis_id = :4, application_id = :5, area_id = :6, area_type = :7, authority_id = :8, authorize_id = :9, ba_organization_id = :10, ba_organization_seq_no = :11, business_associate_id = :12, catalogue_additive_id = :13, catalogue_equip_id = :14, circ_id = :15, classification_system_id = :16, class_level_id = :17, component_type = :18, consent_id = :19, consult_id = :20, contest_id = :21, contract_id = :22, ecozone_id = :23, effective_date = :24, employee_ba_id = :25, employee_obs_no = :26, employer_ba_id = :27, entitlement_id = :28, equipment_id = :29, expiry_date = :30, facility_id = :31, facility_type = :32, field_id = :33, field_station_id = :34, finance_id = :35, fossil_id = :36, incident_id = :37, incident_set_id = :38, incident_type = :39, information_item_id = :40, info_item_subtype = :41, inspection_status = :42, instrument_id = :43, interest_set_id = :44, interest_set_seq_no = :45, jurisdiction = :46, land_right_id = :47, land_right_subtype = :48, land_sale_number = :49, language = :50, lithology_log_id = :51, lith_log_source = :52, notification_id = :53, obligation_id = :54, obligation_seq_no = :55, paleo_summary_id = :56, pden_id = :57, pden_source = :58, pden_subtype = :59, physical_item_id = :60, pool_id = :61, ppdm_guid = :62, ppdm_system_id = :63, ppdm_table_name = :64, product_type = :65, prod_string_id = :66, prod_string_source = :67, project_id = :68, rate_schedule_id = :69, referenced_guid = :70, referenced_system_id = :71, referenced_table_name = :72, remark = :73, resent_id = :74, reserve_class_id = :75, sample_anal_source = :76, seis_set_id = :77, seis_set_subtype = :78, sf_subtype = :79, source = :80, spatial_description_id = :81, spatial_obs_no = :82, store_id = :83, store_structure_obs_no = :84, strat_name_set_id = :85, strat_unit_id = :86, support_facility_id = :87, sw_application_id = :88, thesaurus_id = :89, thesaurus_word_id = :90, uwi = :91, well_activity_set_id = :92, well_activity_source = :93, well_activity_type_id = :94, well_set_id = :95, work_order_id = :96, row_changed_by = :97, row_changed_date = :98, row_created_by = :99, row_effective_date = :100, row_expiry_date = :101, row_quality = :102 where inspection_id = :104")
	if err != nil {
		return err
	}

	seis_insp_component.Row_changed_date = formatDate(seis_insp_component.Row_changed_date)
	seis_insp_component.Effective_date = formatDateString(seis_insp_component.Effective_date)
	seis_insp_component.Expiry_date = formatDateString(seis_insp_component.Expiry_date)
	seis_insp_component.Row_effective_date = formatDateString(seis_insp_component.Row_effective_date)
	seis_insp_component.Row_expiry_date = formatDateString(seis_insp_component.Row_expiry_date)






	rows, err := stmt.Exec(seis_insp_component.Inspection_component_id, seis_insp_component.Active_ind, seis_insp_component.Activity_obs_no, seis_insp_component.Analysis_id, seis_insp_component.Application_id, seis_insp_component.Area_id, seis_insp_component.Area_type, seis_insp_component.Authority_id, seis_insp_component.Authorize_id, seis_insp_component.Ba_organization_id, seis_insp_component.Ba_organization_seq_no, seis_insp_component.Business_associate_id, seis_insp_component.Catalogue_additive_id, seis_insp_component.Catalogue_equip_id, seis_insp_component.Circ_id, seis_insp_component.Classification_system_id, seis_insp_component.Class_level_id, seis_insp_component.Component_type, seis_insp_component.Consent_id, seis_insp_component.Consult_id, seis_insp_component.Contest_id, seis_insp_component.Contract_id, seis_insp_component.Ecozone_id, seis_insp_component.Effective_date, seis_insp_component.Employee_ba_id, seis_insp_component.Employee_obs_no, seis_insp_component.Employer_ba_id, seis_insp_component.Entitlement_id, seis_insp_component.Equipment_id, seis_insp_component.Expiry_date, seis_insp_component.Facility_id, seis_insp_component.Facility_type, seis_insp_component.Field_id, seis_insp_component.Field_station_id, seis_insp_component.Finance_id, seis_insp_component.Fossil_id, seis_insp_component.Incident_id, seis_insp_component.Incident_set_id, seis_insp_component.Incident_type, seis_insp_component.Information_item_id, seis_insp_component.Info_item_subtype, seis_insp_component.Inspection_status, seis_insp_component.Instrument_id, seis_insp_component.Interest_set_id, seis_insp_component.Interest_set_seq_no, seis_insp_component.Jurisdiction, seis_insp_component.Land_right_id, seis_insp_component.Land_right_subtype, seis_insp_component.Land_sale_number, seis_insp_component.Language, seis_insp_component.Lithology_log_id, seis_insp_component.Lith_log_source, seis_insp_component.Notification_id, seis_insp_component.Obligation_id, seis_insp_component.Obligation_seq_no, seis_insp_component.Paleo_summary_id, seis_insp_component.Pden_id, seis_insp_component.Pden_source, seis_insp_component.Pden_subtype, seis_insp_component.Physical_item_id, seis_insp_component.Pool_id, seis_insp_component.Ppdm_guid, seis_insp_component.Ppdm_system_id, seis_insp_component.Ppdm_table_name, seis_insp_component.Product_type, seis_insp_component.Prod_string_id, seis_insp_component.Prod_string_source, seis_insp_component.Project_id, seis_insp_component.Rate_schedule_id, seis_insp_component.Referenced_guid, seis_insp_component.Referenced_system_id, seis_insp_component.Referenced_table_name, seis_insp_component.Remark, seis_insp_component.Resent_id, seis_insp_component.Reserve_class_id, seis_insp_component.Sample_anal_source, seis_insp_component.Seis_set_id, seis_insp_component.Seis_set_subtype, seis_insp_component.Sf_subtype, seis_insp_component.Source, seis_insp_component.Spatial_description_id, seis_insp_component.Spatial_obs_no, seis_insp_component.Store_id, seis_insp_component.Store_structure_obs_no, seis_insp_component.Strat_name_set_id, seis_insp_component.Strat_unit_id, seis_insp_component.Support_facility_id, seis_insp_component.Sw_application_id, seis_insp_component.Thesaurus_id, seis_insp_component.Thesaurus_word_id, seis_insp_component.Uwi, seis_insp_component.Well_activity_set_id, seis_insp_component.Well_activity_source, seis_insp_component.Well_activity_type_id, seis_insp_component.Well_set_id, seis_insp_component.Work_order_id, seis_insp_component.Row_changed_by, seis_insp_component.Row_changed_date, seis_insp_component.Row_created_by, seis_insp_component.Row_effective_date, seis_insp_component.Row_expiry_date, seis_insp_component.Row_quality, seis_insp_component.Inspection_id)
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

func PatchSeisInspComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_insp_component set "
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
	query += " where inspection_id = :id"

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

func DeleteSeisInspComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_insp_component dto.Seis_insp_component
	seis_insp_component.Inspection_id = id

	stmt, err := db.Prepare("delete from seis_insp_component where inspection_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_insp_component.Inspection_id)
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


