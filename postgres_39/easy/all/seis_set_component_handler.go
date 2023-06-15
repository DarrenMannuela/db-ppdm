package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisSetComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_set_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_set_component

	for rows.Next() {
		var seis_set_component dto.Seis_set_component
		if err := rows.Scan(&seis_set_component.Seis_set_subtype, &seis_set_component.Seis_set_id, &seis_set_component.Component_obs_no, &seis_set_component.Active_ind, &seis_set_component.Analysis_id, &seis_set_component.Application_id, &seis_set_component.Area_id, &seis_set_component.Area_type, &seis_set_component.Authority_id, &seis_set_component.Ba_organization_id, &seis_set_component.Ba_organization_seq_no, &seis_set_component.Business_associate_id, &seis_set_component.Catalogue_additive_id, &seis_set_component.Catalogue_equip_id, &seis_set_component.Classification_system_id, &seis_set_component.Class_level_id, &seis_set_component.Consent_id, &seis_set_component.Consult_id, &seis_set_component.Contest_id, &seis_set_component.Contract_id, &seis_set_component.Ecozone_id, &seis_set_component.Effective_date, &seis_set_component.Employee_ba_id, &seis_set_component.Employee_obs_no, &seis_set_component.Employer_ba_id, &seis_set_component.Entitlement_id, &seis_set_component.Equipment_id, &seis_set_component.Expiry_date, &seis_set_component.Facility_id, &seis_set_component.Facility_type, &seis_set_component.Field_id, &seis_set_component.Field_station_id, &seis_set_component.Finance_id, &seis_set_component.Fossil_id, &seis_set_component.Incident_id, &seis_set_component.Incident_set_id, &seis_set_component.Incident_type, &seis_set_component.Information_item_id, &seis_set_component.Info_item_subtype, &seis_set_component.Instrument_id, &seis_set_component.Interest_set_id, &seis_set_component.Interest_set_seq_no, &seis_set_component.Jurisdiction, &seis_set_component.Land_right_id, &seis_set_component.Land_right_subtype, &seis_set_component.Land_sale_number, &seis_set_component.Language, &seis_set_component.Lithology_log_id, &seis_set_component.Lith_log_source, &seis_set_component.Notification_id, &seis_set_component.Obligation_id, &seis_set_component.Obligation_seq_no, &seis_set_component.Paleo_summary_id, &seis_set_component.Pden_id, &seis_set_component.Pden_source, &seis_set_component.Pden_subtype, &seis_set_component.Physical_item_id, &seis_set_component.Pool_id, &seis_set_component.Ppdm_guid, &seis_set_component.Ppdm_system_id, &seis_set_component.Ppdm_table_name, &seis_set_component.Product_type, &seis_set_component.Project_id, &seis_set_component.Pr_str_source, &seis_set_component.Pr_str_uwi, &seis_set_component.Rate_schedule_id, &seis_set_component.Referenced_guid, &seis_set_component.Referenced_system_id, &seis_set_component.Referenced_table_name, &seis_set_component.Remark, &seis_set_component.Resent_id, &seis_set_component.Reserve_class_id, &seis_set_component.Sample_anal_source, &seis_set_component.Seis_act_obs_no, &seis_set_component.Seis_act_seis_set_id, &seis_set_component.Seis_act_seis_set_subtype, &seis_set_component.Seis_set_component_type, &seis_set_component.Sf_subtype, &seis_set_component.Source, &seis_set_component.Spatial_description_id, &seis_set_component.Spatial_obs_no, &seis_set_component.Store_id, &seis_set_component.Store_structure_obs_no, &seis_set_component.Strat_name_set_id, &seis_set_component.Strat_unit_id, &seis_set_component.String_id, &seis_set_component.Support_facility_id, &seis_set_component.Sw_application_id, &seis_set_component.Thesaurus_id, &seis_set_component.Thesaurus_word_id, &seis_set_component.Uwi, &seis_set_component.Well_act_obs_no, &seis_set_component.Well_act_set_id, &seis_set_component.Well_act_source, &seis_set_component.Well_act_type, &seis_set_component.Well_act_uwi, &seis_set_component.Well_set_id, &seis_set_component.Work_order_id, &seis_set_component.Row_changed_by, &seis_set_component.Row_changed_date, &seis_set_component.Row_created_by, &seis_set_component.Row_created_date, &seis_set_component.Row_effective_date, &seis_set_component.Row_expiry_date, &seis_set_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_set_component to result
		result = append(result, seis_set_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisSetComponent(c *fiber.Ctx) error {
	var seis_set_component dto.Seis_set_component

	setDefaults(&seis_set_component)

	if err := json.Unmarshal(c.Body(), &seis_set_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_set_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103, :104, :105)")
	if err != nil {
		return err
	}
	seis_set_component.Row_created_date = formatDate(seis_set_component.Row_created_date)
	seis_set_component.Row_changed_date = formatDate(seis_set_component.Row_changed_date)
	seis_set_component.Effective_date = formatDateString(seis_set_component.Effective_date)
	seis_set_component.Expiry_date = formatDateString(seis_set_component.Expiry_date)
	seis_set_component.Row_effective_date = formatDateString(seis_set_component.Row_effective_date)
	seis_set_component.Row_expiry_date = formatDateString(seis_set_component.Row_expiry_date)






	rows, err := stmt.Exec(seis_set_component.Seis_set_subtype, seis_set_component.Seis_set_id, seis_set_component.Component_obs_no, seis_set_component.Active_ind, seis_set_component.Analysis_id, seis_set_component.Application_id, seis_set_component.Area_id, seis_set_component.Area_type, seis_set_component.Authority_id, seis_set_component.Ba_organization_id, seis_set_component.Ba_organization_seq_no, seis_set_component.Business_associate_id, seis_set_component.Catalogue_additive_id, seis_set_component.Catalogue_equip_id, seis_set_component.Classification_system_id, seis_set_component.Class_level_id, seis_set_component.Consent_id, seis_set_component.Consult_id, seis_set_component.Contest_id, seis_set_component.Contract_id, seis_set_component.Ecozone_id, seis_set_component.Effective_date, seis_set_component.Employee_ba_id, seis_set_component.Employee_obs_no, seis_set_component.Employer_ba_id, seis_set_component.Entitlement_id, seis_set_component.Equipment_id, seis_set_component.Expiry_date, seis_set_component.Facility_id, seis_set_component.Facility_type, seis_set_component.Field_id, seis_set_component.Field_station_id, seis_set_component.Finance_id, seis_set_component.Fossil_id, seis_set_component.Incident_id, seis_set_component.Incident_set_id, seis_set_component.Incident_type, seis_set_component.Information_item_id, seis_set_component.Info_item_subtype, seis_set_component.Instrument_id, seis_set_component.Interest_set_id, seis_set_component.Interest_set_seq_no, seis_set_component.Jurisdiction, seis_set_component.Land_right_id, seis_set_component.Land_right_subtype, seis_set_component.Land_sale_number, seis_set_component.Language, seis_set_component.Lithology_log_id, seis_set_component.Lith_log_source, seis_set_component.Notification_id, seis_set_component.Obligation_id, seis_set_component.Obligation_seq_no, seis_set_component.Paleo_summary_id, seis_set_component.Pden_id, seis_set_component.Pden_source, seis_set_component.Pden_subtype, seis_set_component.Physical_item_id, seis_set_component.Pool_id, seis_set_component.Ppdm_guid, seis_set_component.Ppdm_system_id, seis_set_component.Ppdm_table_name, seis_set_component.Product_type, seis_set_component.Project_id, seis_set_component.Pr_str_source, seis_set_component.Pr_str_uwi, seis_set_component.Rate_schedule_id, seis_set_component.Referenced_guid, seis_set_component.Referenced_system_id, seis_set_component.Referenced_table_name, seis_set_component.Remark, seis_set_component.Resent_id, seis_set_component.Reserve_class_id, seis_set_component.Sample_anal_source, seis_set_component.Seis_act_obs_no, seis_set_component.Seis_act_seis_set_id, seis_set_component.Seis_act_seis_set_subtype, seis_set_component.Seis_set_component_type, seis_set_component.Sf_subtype, seis_set_component.Source, seis_set_component.Spatial_description_id, seis_set_component.Spatial_obs_no, seis_set_component.Store_id, seis_set_component.Store_structure_obs_no, seis_set_component.Strat_name_set_id, seis_set_component.Strat_unit_id, seis_set_component.String_id, seis_set_component.Support_facility_id, seis_set_component.Sw_application_id, seis_set_component.Thesaurus_id, seis_set_component.Thesaurus_word_id, seis_set_component.Uwi, seis_set_component.Well_act_obs_no, seis_set_component.Well_act_set_id, seis_set_component.Well_act_source, seis_set_component.Well_act_type, seis_set_component.Well_act_uwi, seis_set_component.Well_set_id, seis_set_component.Work_order_id, seis_set_component.Row_changed_by, seis_set_component.Row_changed_date, seis_set_component.Row_created_by, seis_set_component.Row_created_date, seis_set_component.Row_effective_date, seis_set_component.Row_expiry_date, seis_set_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisSetComponent(c *fiber.Ctx) error {
	var seis_set_component dto.Seis_set_component

	setDefaults(&seis_set_component)

	if err := json.Unmarshal(c.Body(), &seis_set_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_set_component.Seis_set_subtype = id

    if seis_set_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_set_component set seis_set_id = :1, component_obs_no = :2, active_ind = :3, analysis_id = :4, application_id = :5, area_id = :6, area_type = :7, authority_id = :8, ba_organization_id = :9, ba_organization_seq_no = :10, business_associate_id = :11, catalogue_additive_id = :12, catalogue_equip_id = :13, classification_system_id = :14, class_level_id = :15, consent_id = :16, consult_id = :17, contest_id = :18, contract_id = :19, ecozone_id = :20, effective_date = :21, employee_ba_id = :22, employee_obs_no = :23, employer_ba_id = :24, entitlement_id = :25, equipment_id = :26, expiry_date = :27, facility_id = :28, facility_type = :29, field_id = :30, field_station_id = :31, finance_id = :32, fossil_id = :33, incident_id = :34, incident_set_id = :35, incident_type = :36, information_item_id = :37, info_item_subtype = :38, instrument_id = :39, interest_set_id = :40, interest_set_seq_no = :41, jurisdiction = :42, land_right_id = :43, land_right_subtype = :44, land_sale_number = :45, language = :46, lithology_log_id = :47, lith_log_source = :48, notification_id = :49, obligation_id = :50, obligation_seq_no = :51, paleo_summary_id = :52, pden_id = :53, pden_source = :54, pden_subtype = :55, physical_item_id = :56, pool_id = :57, ppdm_guid = :58, ppdm_system_id = :59, ppdm_table_name = :60, product_type = :61, project_id = :62, pr_str_source = :63, pr_str_uwi = :64, rate_schedule_id = :65, referenced_guid = :66, referenced_system_id = :67, referenced_table_name = :68, remark = :69, resent_id = :70, reserve_class_id = :71, sample_anal_source = :72, seis_act_obs_no = :73, seis_act_seis_set_id = :74, seis_act_seis_set_subtype = :75, seis_set_component_type = :76, sf_subtype = :77, source = :78, spatial_description_id = :79, spatial_obs_no = :80, store_id = :81, store_structure_obs_no = :82, strat_name_set_id = :83, strat_unit_id = :84, string_id = :85, support_facility_id = :86, sw_application_id = :87, thesaurus_id = :88, thesaurus_word_id = :89, uwi = :90, well_act_obs_no = :91, well_act_set_id = :92, well_act_source = :93, well_act_type = :94, well_act_uwi = :95, well_set_id = :96, work_order_id = :97, row_changed_by = :98, row_changed_date = :99, row_created_by = :100, row_effective_date = :101, row_expiry_date = :102, row_quality = :103 where seis_set_subtype = :105")
	if err != nil {
		return err
	}

	seis_set_component.Row_changed_date = formatDate(seis_set_component.Row_changed_date)
	seis_set_component.Effective_date = formatDateString(seis_set_component.Effective_date)
	seis_set_component.Expiry_date = formatDateString(seis_set_component.Expiry_date)
	seis_set_component.Row_effective_date = formatDateString(seis_set_component.Row_effective_date)
	seis_set_component.Row_expiry_date = formatDateString(seis_set_component.Row_expiry_date)






	rows, err := stmt.Exec(seis_set_component.Seis_set_id, seis_set_component.Component_obs_no, seis_set_component.Active_ind, seis_set_component.Analysis_id, seis_set_component.Application_id, seis_set_component.Area_id, seis_set_component.Area_type, seis_set_component.Authority_id, seis_set_component.Ba_organization_id, seis_set_component.Ba_organization_seq_no, seis_set_component.Business_associate_id, seis_set_component.Catalogue_additive_id, seis_set_component.Catalogue_equip_id, seis_set_component.Classification_system_id, seis_set_component.Class_level_id, seis_set_component.Consent_id, seis_set_component.Consult_id, seis_set_component.Contest_id, seis_set_component.Contract_id, seis_set_component.Ecozone_id, seis_set_component.Effective_date, seis_set_component.Employee_ba_id, seis_set_component.Employee_obs_no, seis_set_component.Employer_ba_id, seis_set_component.Entitlement_id, seis_set_component.Equipment_id, seis_set_component.Expiry_date, seis_set_component.Facility_id, seis_set_component.Facility_type, seis_set_component.Field_id, seis_set_component.Field_station_id, seis_set_component.Finance_id, seis_set_component.Fossil_id, seis_set_component.Incident_id, seis_set_component.Incident_set_id, seis_set_component.Incident_type, seis_set_component.Information_item_id, seis_set_component.Info_item_subtype, seis_set_component.Instrument_id, seis_set_component.Interest_set_id, seis_set_component.Interest_set_seq_no, seis_set_component.Jurisdiction, seis_set_component.Land_right_id, seis_set_component.Land_right_subtype, seis_set_component.Land_sale_number, seis_set_component.Language, seis_set_component.Lithology_log_id, seis_set_component.Lith_log_source, seis_set_component.Notification_id, seis_set_component.Obligation_id, seis_set_component.Obligation_seq_no, seis_set_component.Paleo_summary_id, seis_set_component.Pden_id, seis_set_component.Pden_source, seis_set_component.Pden_subtype, seis_set_component.Physical_item_id, seis_set_component.Pool_id, seis_set_component.Ppdm_guid, seis_set_component.Ppdm_system_id, seis_set_component.Ppdm_table_name, seis_set_component.Product_type, seis_set_component.Project_id, seis_set_component.Pr_str_source, seis_set_component.Pr_str_uwi, seis_set_component.Rate_schedule_id, seis_set_component.Referenced_guid, seis_set_component.Referenced_system_id, seis_set_component.Referenced_table_name, seis_set_component.Remark, seis_set_component.Resent_id, seis_set_component.Reserve_class_id, seis_set_component.Sample_anal_source, seis_set_component.Seis_act_obs_no, seis_set_component.Seis_act_seis_set_id, seis_set_component.Seis_act_seis_set_subtype, seis_set_component.Seis_set_component_type, seis_set_component.Sf_subtype, seis_set_component.Source, seis_set_component.Spatial_description_id, seis_set_component.Spatial_obs_no, seis_set_component.Store_id, seis_set_component.Store_structure_obs_no, seis_set_component.Strat_name_set_id, seis_set_component.Strat_unit_id, seis_set_component.String_id, seis_set_component.Support_facility_id, seis_set_component.Sw_application_id, seis_set_component.Thesaurus_id, seis_set_component.Thesaurus_word_id, seis_set_component.Uwi, seis_set_component.Well_act_obs_no, seis_set_component.Well_act_set_id, seis_set_component.Well_act_source, seis_set_component.Well_act_type, seis_set_component.Well_act_uwi, seis_set_component.Well_set_id, seis_set_component.Work_order_id, seis_set_component.Row_changed_by, seis_set_component.Row_changed_date, seis_set_component.Row_created_by, seis_set_component.Row_effective_date, seis_set_component.Row_expiry_date, seis_set_component.Row_quality, seis_set_component.Seis_set_subtype)
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

func PatchSeisSetComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_set_component set "
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
	query += " where seis_set_subtype = :id"

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

func DeleteSeisSetComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_set_component dto.Seis_set_component
	seis_set_component.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_set_component where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_set_component.Seis_set_subtype)
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


