package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetIntSetComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from int_set_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Int_set_component

	for rows.Next() {
		var int_set_component dto.Int_set_component
		if err := rows.Scan(&int_set_component.Interest_set_id, &int_set_component.Interest_set_seq_no, &int_set_component.Component_obs_no, &int_set_component.Active_ind, &int_set_component.Activity_obs_no, &int_set_component.Analysis_id, &int_set_component.Application_id, &int_set_component.Area_id, &int_set_component.Area_type, &int_set_component.Authority_id, &int_set_component.Ba_organization_id, &int_set_component.Ba_organization_seq_no, &int_set_component.Business_associate_id, &int_set_component.Catalogue_additive_id, &int_set_component.Catalogue_equip_id, &int_set_component.Classification_system_id, &int_set_component.Class_level_id, &int_set_component.Component_type, &int_set_component.Consent_id, &int_set_component.Consult_id, &int_set_component.Contest_id, &int_set_component.Contract_id, &int_set_component.Ecozone_id, &int_set_component.Effective_date, &int_set_component.Employee_ba_id, &int_set_component.Employee_obs_no, &int_set_component.Employer_ba_id, &int_set_component.Entitlement_id, &int_set_component.Equipment_id, &int_set_component.Expiry_date, &int_set_component.Facility_id, &int_set_component.Facility_type, &int_set_component.Field_id, &int_set_component.Field_station_id, &int_set_component.Finance_id, &int_set_component.Fossil_id, &int_set_component.Incident_id, &int_set_component.Incident_set_id, &int_set_component.Incident_type, &int_set_component.Information_item_id, &int_set_component.Info_item_subtype, &int_set_component.Instrument_id, &int_set_component.Jurisdiction, &int_set_component.Land_offering_bid, &int_set_component.Land_right_id, &int_set_component.Land_right_subtype, &int_set_component.Land_sale_number, &int_set_component.Land_sale_offering_id, &int_set_component.Language, &int_set_component.Lithology_log_id, &int_set_component.Lith_log_source, &int_set_component.Notification_id, &int_set_component.Obligation_id, &int_set_component.Obligation_seq_no, &int_set_component.Paleo_summary_id, &int_set_component.Pden_activity_type, &int_set_component.Pden_amendment_seq_no, &int_set_component.Pden_disposition_obs_no, &int_set_component.Pden_id, &int_set_component.Pden_period_id, &int_set_component.Pden_period_type, &int_set_component.Pden_product_type, &int_set_component.Pden_source, &int_set_component.Pden_subtype, &int_set_component.Pden_volume_date, &int_set_component.Pden_volume_date_desc, &int_set_component.Pden_volume_method, &int_set_component.Physical_item_id, &int_set_component.Pool_id, &int_set_component.Ppdm_guid, &int_set_component.Ppdm_system_id, &int_set_component.Ppdm_table_name, &int_set_component.Product_type, &int_set_component.Prod_string_id, &int_set_component.Prod_string_source, &int_set_component.Project_id, &int_set_component.Provision_id, &int_set_component.Pr_str_form_obs_no, &int_set_component.Rate_schedule_id, &int_set_component.Referenced_guid, &int_set_component.Referenced_system_id, &int_set_component.Referenced_table_name, &int_set_component.Remark, &int_set_component.Resent_id, &int_set_component.Reserve_class_id, &int_set_component.Sample_anal_source, &int_set_component.Seis_set_id, &int_set_component.Seis_set_subtype, &int_set_component.Sf_subtype, &int_set_component.Source, &int_set_component.Spatial_description_id, &int_set_component.Spatial_obs_no, &int_set_component.Store_id, &int_set_component.Store_structure_obs_no, &int_set_component.Strat_name_set_id, &int_set_component.Strat_unit_id, &int_set_component.Support_facility_id, &int_set_component.Sw_application_id, &int_set_component.Thesaurus_id, &int_set_component.Thesaurus_word_id, &int_set_component.Trigger_desc, &int_set_component.Trigger_event, &int_set_component.Uwi, &int_set_component.Well_activity_set_id, &int_set_component.Well_activity_source, &int_set_component.Well_activity_type_id, &int_set_component.Well_log_curve_id, &int_set_component.Well_log_id, &int_set_component.Well_log_source, &int_set_component.Well_set_id, &int_set_component.Work_order_id, &int_set_component.Row_changed_by, &int_set_component.Row_changed_date, &int_set_component.Row_created_by, &int_set_component.Row_created_date, &int_set_component.Row_effective_date, &int_set_component.Row_expiry_date, &int_set_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append int_set_component to result
		result = append(result, int_set_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetIntSetComponent(c *fiber.Ctx) error {
	var int_set_component dto.Int_set_component

	setDefaults(&int_set_component)

	if err := json.Unmarshal(c.Body(), &int_set_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into int_set_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103, :104, :105, :106, :107, :108, :109, :110, :111, :112, :113, :114, :115, :116, :117, :118)")
	if err != nil {
		return err
	}
	int_set_component.Row_created_date = formatDate(int_set_component.Row_created_date)
	int_set_component.Row_changed_date = formatDate(int_set_component.Row_changed_date)
	int_set_component.Effective_date = formatDateString(int_set_component.Effective_date)
	int_set_component.Expiry_date = formatDateString(int_set_component.Expiry_date)
	int_set_component.Pden_volume_date = formatDateString(int_set_component.Pden_volume_date)
	int_set_component.Row_effective_date = formatDateString(int_set_component.Row_effective_date)
	int_set_component.Row_expiry_date = formatDateString(int_set_component.Row_expiry_date)






	rows, err := stmt.Exec(int_set_component.Interest_set_id, int_set_component.Interest_set_seq_no, int_set_component.Component_obs_no, int_set_component.Active_ind, int_set_component.Activity_obs_no, int_set_component.Analysis_id, int_set_component.Application_id, int_set_component.Area_id, int_set_component.Area_type, int_set_component.Authority_id, int_set_component.Ba_organization_id, int_set_component.Ba_organization_seq_no, int_set_component.Business_associate_id, int_set_component.Catalogue_additive_id, int_set_component.Catalogue_equip_id, int_set_component.Classification_system_id, int_set_component.Class_level_id, int_set_component.Component_type, int_set_component.Consent_id, int_set_component.Consult_id, int_set_component.Contest_id, int_set_component.Contract_id, int_set_component.Ecozone_id, int_set_component.Effective_date, int_set_component.Employee_ba_id, int_set_component.Employee_obs_no, int_set_component.Employer_ba_id, int_set_component.Entitlement_id, int_set_component.Equipment_id, int_set_component.Expiry_date, int_set_component.Facility_id, int_set_component.Facility_type, int_set_component.Field_id, int_set_component.Field_station_id, int_set_component.Finance_id, int_set_component.Fossil_id, int_set_component.Incident_id, int_set_component.Incident_set_id, int_set_component.Incident_type, int_set_component.Information_item_id, int_set_component.Info_item_subtype, int_set_component.Instrument_id, int_set_component.Jurisdiction, int_set_component.Land_offering_bid, int_set_component.Land_right_id, int_set_component.Land_right_subtype, int_set_component.Land_sale_number, int_set_component.Land_sale_offering_id, int_set_component.Language, int_set_component.Lithology_log_id, int_set_component.Lith_log_source, int_set_component.Notification_id, int_set_component.Obligation_id, int_set_component.Obligation_seq_no, int_set_component.Paleo_summary_id, int_set_component.Pden_activity_type, int_set_component.Pden_amendment_seq_no, int_set_component.Pden_disposition_obs_no, int_set_component.Pden_id, int_set_component.Pden_period_id, int_set_component.Pden_period_type, int_set_component.Pden_product_type, int_set_component.Pden_source, int_set_component.Pden_subtype, int_set_component.Pden_volume_date, int_set_component.Pden_volume_date_desc, int_set_component.Pden_volume_method, int_set_component.Physical_item_id, int_set_component.Pool_id, int_set_component.Ppdm_guid, int_set_component.Ppdm_system_id, int_set_component.Ppdm_table_name, int_set_component.Product_type, int_set_component.Prod_string_id, int_set_component.Prod_string_source, int_set_component.Project_id, int_set_component.Provision_id, int_set_component.Pr_str_form_obs_no, int_set_component.Rate_schedule_id, int_set_component.Referenced_guid, int_set_component.Referenced_system_id, int_set_component.Referenced_table_name, int_set_component.Remark, int_set_component.Resent_id, int_set_component.Reserve_class_id, int_set_component.Sample_anal_source, int_set_component.Seis_set_id, int_set_component.Seis_set_subtype, int_set_component.Sf_subtype, int_set_component.Source, int_set_component.Spatial_description_id, int_set_component.Spatial_obs_no, int_set_component.Store_id, int_set_component.Store_structure_obs_no, int_set_component.Strat_name_set_id, int_set_component.Strat_unit_id, int_set_component.Support_facility_id, int_set_component.Sw_application_id, int_set_component.Thesaurus_id, int_set_component.Thesaurus_word_id, int_set_component.Trigger_desc, int_set_component.Trigger_event, int_set_component.Uwi, int_set_component.Well_activity_set_id, int_set_component.Well_activity_source, int_set_component.Well_activity_type_id, int_set_component.Well_log_curve_id, int_set_component.Well_log_id, int_set_component.Well_log_source, int_set_component.Well_set_id, int_set_component.Work_order_id, int_set_component.Row_changed_by, int_set_component.Row_changed_date, int_set_component.Row_created_by, int_set_component.Row_created_date, int_set_component.Row_effective_date, int_set_component.Row_expiry_date, int_set_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateIntSetComponent(c *fiber.Ctx) error {
	var int_set_component dto.Int_set_component

	setDefaults(&int_set_component)

	if err := json.Unmarshal(c.Body(), &int_set_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	int_set_component.Interest_set_id = id

    if int_set_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update int_set_component set interest_set_seq_no = :1, component_obs_no = :2, active_ind = :3, activity_obs_no = :4, analysis_id = :5, application_id = :6, area_id = :7, area_type = :8, authority_id = :9, ba_organization_id = :10, ba_organization_seq_no = :11, business_associate_id = :12, catalogue_additive_id = :13, catalogue_equip_id = :14, classification_system_id = :15, class_level_id = :16, component_type = :17, consent_id = :18, consult_id = :19, contest_id = :20, contract_id = :21, ecozone_id = :22, effective_date = :23, employee_ba_id = :24, employee_obs_no = :25, employer_ba_id = :26, entitlement_id = :27, equipment_id = :28, expiry_date = :29, facility_id = :30, facility_type = :31, field_id = :32, field_station_id = :33, finance_id = :34, fossil_id = :35, incident_id = :36, incident_set_id = :37, incident_type = :38, information_item_id = :39, info_item_subtype = :40, instrument_id = :41, jurisdiction = :42, land_offering_bid = :43, land_right_id = :44, land_right_subtype = :45, land_sale_number = :46, land_sale_offering_id = :47, language = :48, lithology_log_id = :49, lith_log_source = :50, notification_id = :51, obligation_id = :52, obligation_seq_no = :53, paleo_summary_id = :54, pden_activity_type = :55, pden_amendment_seq_no = :56, pden_disposition_obs_no = :57, pden_id = :58, pden_period_id = :59, pden_period_type = :60, pden_product_type = :61, pden_source = :62, pden_subtype = :63, pden_volume_date = :64, pden_volume_date_desc = :65, pden_volume_method = :66, physical_item_id = :67, pool_id = :68, ppdm_guid = :69, ppdm_system_id = :70, ppdm_table_name = :71, product_type = :72, prod_string_id = :73, prod_string_source = :74, project_id = :75, provision_id = :76, pr_str_form_obs_no = :77, rate_schedule_id = :78, referenced_guid = :79, referenced_system_id = :80, referenced_table_name = :81, remark = :82, resent_id = :83, reserve_class_id = :84, sample_anal_source = :85, seis_set_id = :86, seis_set_subtype = :87, sf_subtype = :88, source = :89, spatial_description_id = :90, spatial_obs_no = :91, store_id = :92, store_structure_obs_no = :93, strat_name_set_id = :94, strat_unit_id = :95, support_facility_id = :96, sw_application_id = :97, thesaurus_id = :98, thesaurus_word_id = :99, trigger_desc = :100, trigger_event = :101, uwi = :102, well_activity_set_id = :103, well_activity_source = :104, well_activity_type_id = :105, well_log_curve_id = :106, well_log_id = :107, well_log_source = :108, well_set_id = :109, work_order_id = :110, row_changed_by = :111, row_changed_date = :112, row_created_by = :113, row_effective_date = :114, row_expiry_date = :115, row_quality = :116 where interest_set_id = :118")
	if err != nil {
		return err
	}

	int_set_component.Row_changed_date = formatDate(int_set_component.Row_changed_date)
	int_set_component.Effective_date = formatDateString(int_set_component.Effective_date)
	int_set_component.Expiry_date = formatDateString(int_set_component.Expiry_date)
	int_set_component.Pden_volume_date = formatDateString(int_set_component.Pden_volume_date)
	int_set_component.Row_effective_date = formatDateString(int_set_component.Row_effective_date)
	int_set_component.Row_expiry_date = formatDateString(int_set_component.Row_expiry_date)






	rows, err := stmt.Exec(int_set_component.Interest_set_seq_no, int_set_component.Component_obs_no, int_set_component.Active_ind, int_set_component.Activity_obs_no, int_set_component.Analysis_id, int_set_component.Application_id, int_set_component.Area_id, int_set_component.Area_type, int_set_component.Authority_id, int_set_component.Ba_organization_id, int_set_component.Ba_organization_seq_no, int_set_component.Business_associate_id, int_set_component.Catalogue_additive_id, int_set_component.Catalogue_equip_id, int_set_component.Classification_system_id, int_set_component.Class_level_id, int_set_component.Component_type, int_set_component.Consent_id, int_set_component.Consult_id, int_set_component.Contest_id, int_set_component.Contract_id, int_set_component.Ecozone_id, int_set_component.Effective_date, int_set_component.Employee_ba_id, int_set_component.Employee_obs_no, int_set_component.Employer_ba_id, int_set_component.Entitlement_id, int_set_component.Equipment_id, int_set_component.Expiry_date, int_set_component.Facility_id, int_set_component.Facility_type, int_set_component.Field_id, int_set_component.Field_station_id, int_set_component.Finance_id, int_set_component.Fossil_id, int_set_component.Incident_id, int_set_component.Incident_set_id, int_set_component.Incident_type, int_set_component.Information_item_id, int_set_component.Info_item_subtype, int_set_component.Instrument_id, int_set_component.Jurisdiction, int_set_component.Land_offering_bid, int_set_component.Land_right_id, int_set_component.Land_right_subtype, int_set_component.Land_sale_number, int_set_component.Land_sale_offering_id, int_set_component.Language, int_set_component.Lithology_log_id, int_set_component.Lith_log_source, int_set_component.Notification_id, int_set_component.Obligation_id, int_set_component.Obligation_seq_no, int_set_component.Paleo_summary_id, int_set_component.Pden_activity_type, int_set_component.Pden_amendment_seq_no, int_set_component.Pden_disposition_obs_no, int_set_component.Pden_id, int_set_component.Pden_period_id, int_set_component.Pden_period_type, int_set_component.Pden_product_type, int_set_component.Pden_source, int_set_component.Pden_subtype, int_set_component.Pden_volume_date, int_set_component.Pden_volume_date_desc, int_set_component.Pden_volume_method, int_set_component.Physical_item_id, int_set_component.Pool_id, int_set_component.Ppdm_guid, int_set_component.Ppdm_system_id, int_set_component.Ppdm_table_name, int_set_component.Product_type, int_set_component.Prod_string_id, int_set_component.Prod_string_source, int_set_component.Project_id, int_set_component.Provision_id, int_set_component.Pr_str_form_obs_no, int_set_component.Rate_schedule_id, int_set_component.Referenced_guid, int_set_component.Referenced_system_id, int_set_component.Referenced_table_name, int_set_component.Remark, int_set_component.Resent_id, int_set_component.Reserve_class_id, int_set_component.Sample_anal_source, int_set_component.Seis_set_id, int_set_component.Seis_set_subtype, int_set_component.Sf_subtype, int_set_component.Source, int_set_component.Spatial_description_id, int_set_component.Spatial_obs_no, int_set_component.Store_id, int_set_component.Store_structure_obs_no, int_set_component.Strat_name_set_id, int_set_component.Strat_unit_id, int_set_component.Support_facility_id, int_set_component.Sw_application_id, int_set_component.Thesaurus_id, int_set_component.Thesaurus_word_id, int_set_component.Trigger_desc, int_set_component.Trigger_event, int_set_component.Uwi, int_set_component.Well_activity_set_id, int_set_component.Well_activity_source, int_set_component.Well_activity_type_id, int_set_component.Well_log_curve_id, int_set_component.Well_log_id, int_set_component.Well_log_source, int_set_component.Well_set_id, int_set_component.Work_order_id, int_set_component.Row_changed_by, int_set_component.Row_changed_date, int_set_component.Row_created_by, int_set_component.Row_effective_date, int_set_component.Row_expiry_date, int_set_component.Row_quality, int_set_component.Interest_set_id)
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

func PatchIntSetComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update int_set_component set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "pden_volume_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where interest_set_id = :id"

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

func DeleteIntSetComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var int_set_component dto.Int_set_component
	int_set_component.Interest_set_id = id

	stmt, err := db.Prepare("delete from int_set_component where interest_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(int_set_component.Interest_set_id)
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


