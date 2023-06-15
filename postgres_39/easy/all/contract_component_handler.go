package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetContractComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from contract_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Contract_component

	for rows.Next() {
		var contract_component dto.Contract_component
		if err := rows.Scan(&contract_component.Contract_id, &contract_component.Component_id, &contract_component.Active_ind, &contract_component.Activity_obs_no, &contract_component.Analysis_id, &contract_component.Application_id, &contract_component.Area_id, &contract_component.Area_type, &contract_component.Authority_id, &contract_component.Ba_organization_id, &contract_component.Ba_organization_seq_no, &contract_component.Ba_service_seq_no, &contract_component.Business_associate_id, &contract_component.Catalogue_additive_id, &contract_component.Catalogue_equip_id, &contract_component.Classification_system_id, &contract_component.Class_level_id, &contract_component.Consent_id, &contract_component.Consult_id, &contract_component.Contest_id, &contract_component.Contract_component_type, &contract_component.Ecozone_id, &contract_component.Effective_date, &contract_component.Employee_ba_id, &contract_component.Employee_obs_no, &contract_component.Employer_ba_id, &contract_component.Entitlement_id, &contract_component.Equipment_id, &contract_component.Expiry_date, &contract_component.Facility_id, &contract_component.Facility_type, &contract_component.Field_id, &contract_component.Field_station_id, &contract_component.Finance_id, &contract_component.Fossil_id, &contract_component.Incident_id, &contract_component.Incident_set_id, &contract_component.Incident_type, &contract_component.Information_item_id, &contract_component.Info_item_subtype, &contract_component.Instrument_id, &contract_component.Interest_set_id, &contract_component.Interest_set_seq_no, &contract_component.Jurisdiction, &contract_component.Land_offering_bid_id, &contract_component.Land_request_id, &contract_component.Land_right_acqtn_id, &contract_component.Land_right_id, &contract_component.Land_right_subtype, &contract_component.Land_sale_number, &contract_component.Land_sale_offering_id, &contract_component.Language, &contract_component.Lithology_log_id, &contract_component.Lith_log_source, &contract_component.Notification_id, &contract_component.Obligation_id, &contract_component.Obligation_seq_no, &contract_component.Paleo_summary_id, &contract_component.Pden_id, &contract_component.Pden_source, &contract_component.Pden_subtype, &contract_component.Physical_item_id, &contract_component.Pool_id, &contract_component.Ppdm_guid, &contract_component.Ppdm_system_id, &contract_component.Ppdm_table_name, &contract_component.Product_type, &contract_component.Prod_string_id, &contract_component.Prod_string_source, &contract_component.Project_id, &contract_component.Provision_id, &contract_component.Rate_schedule_id, &contract_component.Reason_type, &contract_component.Referenced_guid, &contract_component.Referenced_system_id, &contract_component.Referenced_table_name, &contract_component.Remark, &contract_component.Resent_id, &contract_component.Reserve_class_id, &contract_component.Sample_anal_source, &contract_component.Seis_proc_plan_id, &contract_component.Seis_set_id, &contract_component.Seis_set_subtype, &contract_component.Seis_transaction_id, &contract_component.Seis_transaction_type, &contract_component.Source, &contract_component.Spatial_description_id, &contract_component.Spatial_obs_no, &contract_component.Store_id, &contract_component.Store_structure_obs_no, &contract_component.Strat_name_set_id, &contract_component.Strat_unit_id, &contract_component.Support_facility_id, &contract_component.Support_facility_subtype, &contract_component.Sw_application_id, &contract_component.Thesaurus_id, &contract_component.Thesaurus_word_id, &contract_component.Uwi, &contract_component.Well_activity_set_id, &contract_component.Well_activity_source, &contract_component.Well_activity_type_id, &contract_component.Well_set_id, &contract_component.Work_order_id, &contract_component.Row_changed_by, &contract_component.Row_changed_date, &contract_component.Row_created_by, &contract_component.Row_created_date, &contract_component.Row_effective_date, &contract_component.Row_expiry_date, &contract_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append contract_component to result
		result = append(result, contract_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetContractComponent(c *fiber.Ctx) error {
	var contract_component dto.Contract_component

	setDefaults(&contract_component)

	if err := json.Unmarshal(c.Body(), &contract_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into contract_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103, :104, :105, :106, :107, :108, :109, :110)")
	if err != nil {
		return err
	}
	contract_component.Row_created_date = formatDate(contract_component.Row_created_date)
	contract_component.Row_changed_date = formatDate(contract_component.Row_changed_date)
	contract_component.Effective_date = formatDateString(contract_component.Effective_date)
	contract_component.Expiry_date = formatDateString(contract_component.Expiry_date)
	contract_component.Row_effective_date = formatDateString(contract_component.Row_effective_date)
	contract_component.Row_expiry_date = formatDateString(contract_component.Row_expiry_date)






	rows, err := stmt.Exec(contract_component.Contract_id, contract_component.Component_id, contract_component.Active_ind, contract_component.Activity_obs_no, contract_component.Analysis_id, contract_component.Application_id, contract_component.Area_id, contract_component.Area_type, contract_component.Authority_id, contract_component.Ba_organization_id, contract_component.Ba_organization_seq_no, contract_component.Ba_service_seq_no, contract_component.Business_associate_id, contract_component.Catalogue_additive_id, contract_component.Catalogue_equip_id, contract_component.Classification_system_id, contract_component.Class_level_id, contract_component.Consent_id, contract_component.Consult_id, contract_component.Contest_id, contract_component.Contract_component_type, contract_component.Ecozone_id, contract_component.Effective_date, contract_component.Employee_ba_id, contract_component.Employee_obs_no, contract_component.Employer_ba_id, contract_component.Entitlement_id, contract_component.Equipment_id, contract_component.Expiry_date, contract_component.Facility_id, contract_component.Facility_type, contract_component.Field_id, contract_component.Field_station_id, contract_component.Finance_id, contract_component.Fossil_id, contract_component.Incident_id, contract_component.Incident_set_id, contract_component.Incident_type, contract_component.Information_item_id, contract_component.Info_item_subtype, contract_component.Instrument_id, contract_component.Interest_set_id, contract_component.Interest_set_seq_no, contract_component.Jurisdiction, contract_component.Land_offering_bid_id, contract_component.Land_request_id, contract_component.Land_right_acqtn_id, contract_component.Land_right_id, contract_component.Land_right_subtype, contract_component.Land_sale_number, contract_component.Land_sale_offering_id, contract_component.Language, contract_component.Lithology_log_id, contract_component.Lith_log_source, contract_component.Notification_id, contract_component.Obligation_id, contract_component.Obligation_seq_no, contract_component.Paleo_summary_id, contract_component.Pden_id, contract_component.Pden_source, contract_component.Pden_subtype, contract_component.Physical_item_id, contract_component.Pool_id, contract_component.Ppdm_guid, contract_component.Ppdm_system_id, contract_component.Ppdm_table_name, contract_component.Product_type, contract_component.Prod_string_id, contract_component.Prod_string_source, contract_component.Project_id, contract_component.Provision_id, contract_component.Rate_schedule_id, contract_component.Reason_type, contract_component.Referenced_guid, contract_component.Referenced_system_id, contract_component.Referenced_table_name, contract_component.Remark, contract_component.Resent_id, contract_component.Reserve_class_id, contract_component.Sample_anal_source, contract_component.Seis_proc_plan_id, contract_component.Seis_set_id, contract_component.Seis_set_subtype, contract_component.Seis_transaction_id, contract_component.Seis_transaction_type, contract_component.Source, contract_component.Spatial_description_id, contract_component.Spatial_obs_no, contract_component.Store_id, contract_component.Store_structure_obs_no, contract_component.Strat_name_set_id, contract_component.Strat_unit_id, contract_component.Support_facility_id, contract_component.Support_facility_subtype, contract_component.Sw_application_id, contract_component.Thesaurus_id, contract_component.Thesaurus_word_id, contract_component.Uwi, contract_component.Well_activity_set_id, contract_component.Well_activity_source, contract_component.Well_activity_type_id, contract_component.Well_set_id, contract_component.Work_order_id, contract_component.Row_changed_by, contract_component.Row_changed_date, contract_component.Row_created_by, contract_component.Row_created_date, contract_component.Row_effective_date, contract_component.Row_expiry_date, contract_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateContractComponent(c *fiber.Ctx) error {
	var contract_component dto.Contract_component

	setDefaults(&contract_component)

	if err := json.Unmarshal(c.Body(), &contract_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	contract_component.Contract_id = id

    if contract_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update contract_component set component_id = :1, active_ind = :2, activity_obs_no = :3, analysis_id = :4, application_id = :5, area_id = :6, area_type = :7, authority_id = :8, ba_organization_id = :9, ba_organization_seq_no = :10, ba_service_seq_no = :11, business_associate_id = :12, catalogue_additive_id = :13, catalogue_equip_id = :14, classification_system_id = :15, class_level_id = :16, consent_id = :17, consult_id = :18, contest_id = :19, contract_component_type = :20, ecozone_id = :21, effective_date = :22, employee_ba_id = :23, employee_obs_no = :24, employer_ba_id = :25, entitlement_id = :26, equipment_id = :27, expiry_date = :28, facility_id = :29, facility_type = :30, field_id = :31, field_station_id = :32, finance_id = :33, fossil_id = :34, incident_id = :35, incident_set_id = :36, incident_type = :37, information_item_id = :38, info_item_subtype = :39, instrument_id = :40, interest_set_id = :41, interest_set_seq_no = :42, jurisdiction = :43, land_offering_bid_id = :44, land_request_id = :45, land_right_acqtn_id = :46, land_right_id = :47, land_right_subtype = :48, land_sale_number = :49, land_sale_offering_id = :50, language = :51, lithology_log_id = :52, lith_log_source = :53, notification_id = :54, obligation_id = :55, obligation_seq_no = :56, paleo_summary_id = :57, pden_id = :58, pden_source = :59, pden_subtype = :60, physical_item_id = :61, pool_id = :62, ppdm_guid = :63, ppdm_system_id = :64, ppdm_table_name = :65, product_type = :66, prod_string_id = :67, prod_string_source = :68, project_id = :69, provision_id = :70, rate_schedule_id = :71, reason_type = :72, referenced_guid = :73, referenced_system_id = :74, referenced_table_name = :75, remark = :76, resent_id = :77, reserve_class_id = :78, sample_anal_source = :79, seis_proc_plan_id = :80, seis_set_id = :81, seis_set_subtype = :82, seis_transaction_id = :83, seis_transaction_type = :84, source = :85, spatial_description_id = :86, spatial_obs_no = :87, store_id = :88, store_structure_obs_no = :89, strat_name_set_id = :90, strat_unit_id = :91, support_facility_id = :92, support_facility_subtype = :93, sw_application_id = :94, thesaurus_id = :95, thesaurus_word_id = :96, uwi = :97, well_activity_set_id = :98, well_activity_source = :99, well_activity_type_id = :100, well_set_id = :101, work_order_id = :102, row_changed_by = :103, row_changed_date = :104, row_created_by = :105, row_effective_date = :106, row_expiry_date = :107, row_quality = :108 where contract_id = :110")
	if err != nil {
		return err
	}

	contract_component.Row_changed_date = formatDate(contract_component.Row_changed_date)
	contract_component.Effective_date = formatDateString(contract_component.Effective_date)
	contract_component.Expiry_date = formatDateString(contract_component.Expiry_date)
	contract_component.Row_effective_date = formatDateString(contract_component.Row_effective_date)
	contract_component.Row_expiry_date = formatDateString(contract_component.Row_expiry_date)






	rows, err := stmt.Exec(contract_component.Component_id, contract_component.Active_ind, contract_component.Activity_obs_no, contract_component.Analysis_id, contract_component.Application_id, contract_component.Area_id, contract_component.Area_type, contract_component.Authority_id, contract_component.Ba_organization_id, contract_component.Ba_organization_seq_no, contract_component.Ba_service_seq_no, contract_component.Business_associate_id, contract_component.Catalogue_additive_id, contract_component.Catalogue_equip_id, contract_component.Classification_system_id, contract_component.Class_level_id, contract_component.Consent_id, contract_component.Consult_id, contract_component.Contest_id, contract_component.Contract_component_type, contract_component.Ecozone_id, contract_component.Effective_date, contract_component.Employee_ba_id, contract_component.Employee_obs_no, contract_component.Employer_ba_id, contract_component.Entitlement_id, contract_component.Equipment_id, contract_component.Expiry_date, contract_component.Facility_id, contract_component.Facility_type, contract_component.Field_id, contract_component.Field_station_id, contract_component.Finance_id, contract_component.Fossil_id, contract_component.Incident_id, contract_component.Incident_set_id, contract_component.Incident_type, contract_component.Information_item_id, contract_component.Info_item_subtype, contract_component.Instrument_id, contract_component.Interest_set_id, contract_component.Interest_set_seq_no, contract_component.Jurisdiction, contract_component.Land_offering_bid_id, contract_component.Land_request_id, contract_component.Land_right_acqtn_id, contract_component.Land_right_id, contract_component.Land_right_subtype, contract_component.Land_sale_number, contract_component.Land_sale_offering_id, contract_component.Language, contract_component.Lithology_log_id, contract_component.Lith_log_source, contract_component.Notification_id, contract_component.Obligation_id, contract_component.Obligation_seq_no, contract_component.Paleo_summary_id, contract_component.Pden_id, contract_component.Pden_source, contract_component.Pden_subtype, contract_component.Physical_item_id, contract_component.Pool_id, contract_component.Ppdm_guid, contract_component.Ppdm_system_id, contract_component.Ppdm_table_name, contract_component.Product_type, contract_component.Prod_string_id, contract_component.Prod_string_source, contract_component.Project_id, contract_component.Provision_id, contract_component.Rate_schedule_id, contract_component.Reason_type, contract_component.Referenced_guid, contract_component.Referenced_system_id, contract_component.Referenced_table_name, contract_component.Remark, contract_component.Resent_id, contract_component.Reserve_class_id, contract_component.Sample_anal_source, contract_component.Seis_proc_plan_id, contract_component.Seis_set_id, contract_component.Seis_set_subtype, contract_component.Seis_transaction_id, contract_component.Seis_transaction_type, contract_component.Source, contract_component.Spatial_description_id, contract_component.Spatial_obs_no, contract_component.Store_id, contract_component.Store_structure_obs_no, contract_component.Strat_name_set_id, contract_component.Strat_unit_id, contract_component.Support_facility_id, contract_component.Support_facility_subtype, contract_component.Sw_application_id, contract_component.Thesaurus_id, contract_component.Thesaurus_word_id, contract_component.Uwi, contract_component.Well_activity_set_id, contract_component.Well_activity_source, contract_component.Well_activity_type_id, contract_component.Well_set_id, contract_component.Work_order_id, contract_component.Row_changed_by, contract_component.Row_changed_date, contract_component.Row_created_by, contract_component.Row_effective_date, contract_component.Row_expiry_date, contract_component.Row_quality, contract_component.Contract_id)
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

func PatchContractComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update contract_component set "
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
	query += " where contract_id = :id"

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

func DeleteContractComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var contract_component dto.Contract_component
	contract_component.Contract_id = id

	stmt, err := db.Prepare("delete from contract_component where contract_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(contract_component.Contract_id)
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


