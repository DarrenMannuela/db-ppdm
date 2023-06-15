package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetObligationComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from obligation_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Obligation_component

	for rows.Next() {
		var obligation_component dto.Obligation_component
		if err := rows.Scan(&obligation_component.Obligation_id, &obligation_component.Obligation_seq_no, &obligation_component.Component_obs_no, &obligation_component.Active_ind, &obligation_component.Activity_obs_no, &obligation_component.Analysis_id, &obligation_component.Application_id, &obligation_component.Area_id, &obligation_component.Area_type, &obligation_component.Authority_id, &obligation_component.Ba_licensee_ba_id, &obligation_component.Ba_license_id, &obligation_component.Ba_organization_id, &obligation_component.Ba_organization_seq_no, &obligation_component.Ba_service_ba_id, &obligation_component.Ba_service_provider, &obligation_component.Ba_service_seq_no, &obligation_component.Business_associate_id, &obligation_component.Calculation_method, &obligation_component.Catalogue_additive_id, &obligation_component.Catalogue_equip_id, &obligation_component.Classification_system_id, &obligation_component.Class_level_id, &obligation_component.Component_reason, &obligation_component.Consent_id, &obligation_component.Consult_id, &obligation_component.Contest_id, &obligation_component.Contract_extension_id, &obligation_component.Contract_id, &obligation_component.Contract_provision_id, &obligation_component.Cont_service_ba_id, &obligation_component.Cont_service_seq_no, &obligation_component.Critical_date, &obligation_component.Description, &obligation_component.Ecozone_id, &obligation_component.Effective_date, &obligation_component.Employee_ba_id, &obligation_component.Employee_obs_no, &obligation_component.Employer_ba_id, &obligation_component.Entitlement_id, &obligation_component.Equipment_id, &obligation_component.Expiry_date, &obligation_component.Facility_id, &obligation_component.Facility_license_id, &obligation_component.Facility_service_ba_id, &obligation_component.Facility_service_seq_no, &obligation_component.Facility_type, &obligation_component.Field_id, &obligation_component.Field_station_id, &obligation_component.Finance_id, &obligation_component.Fossil_id, &obligation_component.Fulfilled_date, &obligation_component.Fulfilled_ind, &obligation_component.Fulfilled_remark, &obligation_component.Hse_incident_id, &obligation_component.Incident_set_id, &obligation_component.Incident_type, &obligation_component.Information_item_id, &obligation_component.Info_item_subtype, &obligation_component.Instrument_id, &obligation_component.Interest_set_id, &obligation_component.Interest_set_seq_no, &obligation_component.Land_cost_id, &obligation_component.Land_offering_bid_id, &obligation_component.Land_request_id, &obligation_component.Land_right_id, &obligation_component.Land_right_subtype, &obligation_component.Land_sale_bid_set_id, &obligation_component.Land_sale_jurisdiction, &obligation_component.Land_sale_number, &obligation_component.Land_sale_offering_id, &obligation_component.Land_service_provided_by, &obligation_component.Land_service_seq_no, &obligation_component.Language, &obligation_component.Lithology_log_id, &obligation_component.Lith_log_source, &obligation_component.Lr_termination_id, &obligation_component.Lr_termination_seq_no, &obligation_component.Notification_id, &obligation_component.Obligation_component_type, &obligation_component.Paleo_summary_id, &obligation_component.Pden_id, &obligation_component.Pden_source, &obligation_component.Pden_subtype, &obligation_component.Physical_item_id, &obligation_component.Pool_id, &obligation_component.Ppdm_guid, &obligation_component.Ppdm_system_id, &obligation_component.Ppdm_table_name, &obligation_component.Product_type, &obligation_component.Prod_string_id, &obligation_component.Prod_string_source, &obligation_component.Project_id, &obligation_component.Rate_schedule_id, &obligation_component.Referenced_guid, &obligation_component.Referenced_system_id, &obligation_component.Referenced_table_name, &obligation_component.Remark, &obligation_component.Resent_id, &obligation_component.Reserve_class_id, &obligation_component.Sample_anal_source, &obligation_component.Seis_license_id, &obligation_component.Seis_point_flow_id, &obligation_component.Seis_point_id, &obligation_component.Seis_service_ba_id, &obligation_component.Seis_service_seq_no, &obligation_component.Seis_set_id, &obligation_component.Seis_set_subtype, &obligation_component.Seis_transaction_id, &obligation_component.Seis_transaction_type, &obligation_component.Sf_subtype, &obligation_component.Source, &obligation_component.Spatial_description_id, &obligation_component.Spatial_obs_no, &obligation_component.Store_id, &obligation_component.Store_structure_obs_no, &obligation_component.Strat_name_set_id, &obligation_component.Strat_unit_id, &obligation_component.Support_facility_id, &obligation_component.Sw_application_id, &obligation_component.Thesaurus_id, &obligation_component.Thesaurus_word_id, &obligation_component.Uwi, &obligation_component.Well_activity_set_id, &obligation_component.Well_activity_source, &obligation_component.Well_activity_type_id, &obligation_component.Well_license_id, &obligation_component.Well_license_source, &obligation_component.Well_service_ba_id, &obligation_component.Well_service_seq_no, &obligation_component.Well_set_id, &obligation_component.Work_order_delivery_id, &obligation_component.Work_order_id, &obligation_component.Row_changed_by, &obligation_component.Row_changed_date, &obligation_component.Row_created_by, &obligation_component.Row_created_date, &obligation_component.Row_effective_date, &obligation_component.Row_expiry_date, &obligation_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append obligation_component to result
		result = append(result, obligation_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetObligationComponent(c *fiber.Ctx) error {
	var obligation_component dto.Obligation_component

	setDefaults(&obligation_component)

	if err := json.Unmarshal(c.Body(), &obligation_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into obligation_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103, :104, :105, :106, :107, :108, :109, :110, :111, :112, :113, :114, :115, :116, :117, :118, :119, :120, :121, :122, :123, :124, :125, :126, :127, :128, :129, :130, :131, :132, :133, :134, :135, :136, :137, :138, :139, :140)")
	if err != nil {
		return err
	}
	obligation_component.Row_created_date = formatDate(obligation_component.Row_created_date)
	obligation_component.Row_changed_date = formatDate(obligation_component.Row_changed_date)
	obligation_component.Critical_date = formatDateString(obligation_component.Critical_date)
	obligation_component.Effective_date = formatDateString(obligation_component.Effective_date)
	obligation_component.Expiry_date = formatDateString(obligation_component.Expiry_date)
	obligation_component.Fulfilled_date = formatDateString(obligation_component.Fulfilled_date)
	obligation_component.Row_effective_date = formatDateString(obligation_component.Row_effective_date)
	obligation_component.Row_expiry_date = formatDateString(obligation_component.Row_expiry_date)






	rows, err := stmt.Exec(obligation_component.Obligation_id, obligation_component.Obligation_seq_no, obligation_component.Component_obs_no, obligation_component.Active_ind, obligation_component.Activity_obs_no, obligation_component.Analysis_id, obligation_component.Application_id, obligation_component.Area_id, obligation_component.Area_type, obligation_component.Authority_id, obligation_component.Ba_licensee_ba_id, obligation_component.Ba_license_id, obligation_component.Ba_organization_id, obligation_component.Ba_organization_seq_no, obligation_component.Ba_service_ba_id, obligation_component.Ba_service_provider, obligation_component.Ba_service_seq_no, obligation_component.Business_associate_id, obligation_component.Calculation_method, obligation_component.Catalogue_additive_id, obligation_component.Catalogue_equip_id, obligation_component.Classification_system_id, obligation_component.Class_level_id, obligation_component.Component_reason, obligation_component.Consent_id, obligation_component.Consult_id, obligation_component.Contest_id, obligation_component.Contract_extension_id, obligation_component.Contract_id, obligation_component.Contract_provision_id, obligation_component.Cont_service_ba_id, obligation_component.Cont_service_seq_no, obligation_component.Critical_date, obligation_component.Description, obligation_component.Ecozone_id, obligation_component.Effective_date, obligation_component.Employee_ba_id, obligation_component.Employee_obs_no, obligation_component.Employer_ba_id, obligation_component.Entitlement_id, obligation_component.Equipment_id, obligation_component.Expiry_date, obligation_component.Facility_id, obligation_component.Facility_license_id, obligation_component.Facility_service_ba_id, obligation_component.Facility_service_seq_no, obligation_component.Facility_type, obligation_component.Field_id, obligation_component.Field_station_id, obligation_component.Finance_id, obligation_component.Fossil_id, obligation_component.Fulfilled_date, obligation_component.Fulfilled_ind, obligation_component.Fulfilled_remark, obligation_component.Hse_incident_id, obligation_component.Incident_set_id, obligation_component.Incident_type, obligation_component.Information_item_id, obligation_component.Info_item_subtype, obligation_component.Instrument_id, obligation_component.Interest_set_id, obligation_component.Interest_set_seq_no, obligation_component.Land_cost_id, obligation_component.Land_offering_bid_id, obligation_component.Land_request_id, obligation_component.Land_right_id, obligation_component.Land_right_subtype, obligation_component.Land_sale_bid_set_id, obligation_component.Land_sale_jurisdiction, obligation_component.Land_sale_number, obligation_component.Land_sale_offering_id, obligation_component.Land_service_provided_by, obligation_component.Land_service_seq_no, obligation_component.Language, obligation_component.Lithology_log_id, obligation_component.Lith_log_source, obligation_component.Lr_termination_id, obligation_component.Lr_termination_seq_no, obligation_component.Notification_id, obligation_component.Obligation_component_type, obligation_component.Paleo_summary_id, obligation_component.Pden_id, obligation_component.Pden_source, obligation_component.Pden_subtype, obligation_component.Physical_item_id, obligation_component.Pool_id, obligation_component.Ppdm_guid, obligation_component.Ppdm_system_id, obligation_component.Ppdm_table_name, obligation_component.Product_type, obligation_component.Prod_string_id, obligation_component.Prod_string_source, obligation_component.Project_id, obligation_component.Rate_schedule_id, obligation_component.Referenced_guid, obligation_component.Referenced_system_id, obligation_component.Referenced_table_name, obligation_component.Remark, obligation_component.Resent_id, obligation_component.Reserve_class_id, obligation_component.Sample_anal_source, obligation_component.Seis_license_id, obligation_component.Seis_point_flow_id, obligation_component.Seis_point_id, obligation_component.Seis_service_ba_id, obligation_component.Seis_service_seq_no, obligation_component.Seis_set_id, obligation_component.Seis_set_subtype, obligation_component.Seis_transaction_id, obligation_component.Seis_transaction_type, obligation_component.Sf_subtype, obligation_component.Source, obligation_component.Spatial_description_id, obligation_component.Spatial_obs_no, obligation_component.Store_id, obligation_component.Store_structure_obs_no, obligation_component.Strat_name_set_id, obligation_component.Strat_unit_id, obligation_component.Support_facility_id, obligation_component.Sw_application_id, obligation_component.Thesaurus_id, obligation_component.Thesaurus_word_id, obligation_component.Uwi, obligation_component.Well_activity_set_id, obligation_component.Well_activity_source, obligation_component.Well_activity_type_id, obligation_component.Well_license_id, obligation_component.Well_license_source, obligation_component.Well_service_ba_id, obligation_component.Well_service_seq_no, obligation_component.Well_set_id, obligation_component.Work_order_delivery_id, obligation_component.Work_order_id, obligation_component.Row_changed_by, obligation_component.Row_changed_date, obligation_component.Row_created_by, obligation_component.Row_created_date, obligation_component.Row_effective_date, obligation_component.Row_expiry_date, obligation_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateObligationComponent(c *fiber.Ctx) error {
	var obligation_component dto.Obligation_component

	setDefaults(&obligation_component)

	if err := json.Unmarshal(c.Body(), &obligation_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	obligation_component.Obligation_id = id

    if obligation_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update obligation_component set obligation_seq_no = :1, component_obs_no = :2, active_ind = :3, activity_obs_no = :4, analysis_id = :5, application_id = :6, area_id = :7, area_type = :8, authority_id = :9, ba_licensee_ba_id = :10, ba_license_id = :11, ba_organization_id = :12, ba_organization_seq_no = :13, ba_service_ba_id = :14, ba_service_provider = :15, ba_service_seq_no = :16, business_associate_id = :17, calculation_method = :18, catalogue_additive_id = :19, catalogue_equip_id = :20, classification_system_id = :21, class_level_id = :22, component_reason = :23, consent_id = :24, consult_id = :25, contest_id = :26, contract_extension_id = :27, contract_id = :28, contract_provision_id = :29, cont_service_ba_id = :30, cont_service_seq_no = :31, critical_date = :32, description = :33, ecozone_id = :34, effective_date = :35, employee_ba_id = :36, employee_obs_no = :37, employer_ba_id = :38, entitlement_id = :39, equipment_id = :40, expiry_date = :41, facility_id = :42, facility_license_id = :43, facility_service_ba_id = :44, facility_service_seq_no = :45, facility_type = :46, field_id = :47, field_station_id = :48, finance_id = :49, fossil_id = :50, fulfilled_date = :51, fulfilled_ind = :52, fulfilled_remark = :53, hse_incident_id = :54, incident_set_id = :55, incident_type = :56, information_item_id = :57, info_item_subtype = :58, instrument_id = :59, interest_set_id = :60, interest_set_seq_no = :61, land_cost_id = :62, land_offering_bid_id = :63, land_request_id = :64, land_right_id = :65, land_right_subtype = :66, land_sale_bid_set_id = :67, land_sale_jurisdiction = :68, land_sale_number = :69, land_sale_offering_id = :70, land_service_provided_by = :71, land_service_seq_no = :72, language = :73, lithology_log_id = :74, lith_log_source = :75, lr_termination_id = :76, lr_termination_seq_no = :77, notification_id = :78, obligation_component_type = :79, paleo_summary_id = :80, pden_id = :81, pden_source = :82, pden_subtype = :83, physical_item_id = :84, pool_id = :85, ppdm_guid = :86, ppdm_system_id = :87, ppdm_table_name = :88, product_type = :89, prod_string_id = :90, prod_string_source = :91, project_id = :92, rate_schedule_id = :93, referenced_guid = :94, referenced_system_id = :95, referenced_table_name = :96, remark = :97, resent_id = :98, reserve_class_id = :99, sample_anal_source = :100, seis_license_id = :101, seis_point_flow_id = :102, seis_point_id = :103, seis_service_ba_id = :104, seis_service_seq_no = :105, seis_set_id = :106, seis_set_subtype = :107, seis_transaction_id = :108, seis_transaction_type = :109, sf_subtype = :110, source = :111, spatial_description_id = :112, spatial_obs_no = :113, store_id = :114, store_structure_obs_no = :115, strat_name_set_id = :116, strat_unit_id = :117, support_facility_id = :118, sw_application_id = :119, thesaurus_id = :120, thesaurus_word_id = :121, uwi = :122, well_activity_set_id = :123, well_activity_source = :124, well_activity_type_id = :125, well_license_id = :126, well_license_source = :127, well_service_ba_id = :128, well_service_seq_no = :129, well_set_id = :130, work_order_delivery_id = :131, work_order_id = :132, row_changed_by = :133, row_changed_date = :134, row_created_by = :135, row_effective_date = :136, row_expiry_date = :137, row_quality = :138 where obligation_id = :140")
	if err != nil {
		return err
	}

	obligation_component.Row_changed_date = formatDate(obligation_component.Row_changed_date)
	obligation_component.Critical_date = formatDateString(obligation_component.Critical_date)
	obligation_component.Effective_date = formatDateString(obligation_component.Effective_date)
	obligation_component.Expiry_date = formatDateString(obligation_component.Expiry_date)
	obligation_component.Fulfilled_date = formatDateString(obligation_component.Fulfilled_date)
	obligation_component.Row_effective_date = formatDateString(obligation_component.Row_effective_date)
	obligation_component.Row_expiry_date = formatDateString(obligation_component.Row_expiry_date)






	rows, err := stmt.Exec(obligation_component.Obligation_seq_no, obligation_component.Component_obs_no, obligation_component.Active_ind, obligation_component.Activity_obs_no, obligation_component.Analysis_id, obligation_component.Application_id, obligation_component.Area_id, obligation_component.Area_type, obligation_component.Authority_id, obligation_component.Ba_licensee_ba_id, obligation_component.Ba_license_id, obligation_component.Ba_organization_id, obligation_component.Ba_organization_seq_no, obligation_component.Ba_service_ba_id, obligation_component.Ba_service_provider, obligation_component.Ba_service_seq_no, obligation_component.Business_associate_id, obligation_component.Calculation_method, obligation_component.Catalogue_additive_id, obligation_component.Catalogue_equip_id, obligation_component.Classification_system_id, obligation_component.Class_level_id, obligation_component.Component_reason, obligation_component.Consent_id, obligation_component.Consult_id, obligation_component.Contest_id, obligation_component.Contract_extension_id, obligation_component.Contract_id, obligation_component.Contract_provision_id, obligation_component.Cont_service_ba_id, obligation_component.Cont_service_seq_no, obligation_component.Critical_date, obligation_component.Description, obligation_component.Ecozone_id, obligation_component.Effective_date, obligation_component.Employee_ba_id, obligation_component.Employee_obs_no, obligation_component.Employer_ba_id, obligation_component.Entitlement_id, obligation_component.Equipment_id, obligation_component.Expiry_date, obligation_component.Facility_id, obligation_component.Facility_license_id, obligation_component.Facility_service_ba_id, obligation_component.Facility_service_seq_no, obligation_component.Facility_type, obligation_component.Field_id, obligation_component.Field_station_id, obligation_component.Finance_id, obligation_component.Fossil_id, obligation_component.Fulfilled_date, obligation_component.Fulfilled_ind, obligation_component.Fulfilled_remark, obligation_component.Hse_incident_id, obligation_component.Incident_set_id, obligation_component.Incident_type, obligation_component.Information_item_id, obligation_component.Info_item_subtype, obligation_component.Instrument_id, obligation_component.Interest_set_id, obligation_component.Interest_set_seq_no, obligation_component.Land_cost_id, obligation_component.Land_offering_bid_id, obligation_component.Land_request_id, obligation_component.Land_right_id, obligation_component.Land_right_subtype, obligation_component.Land_sale_bid_set_id, obligation_component.Land_sale_jurisdiction, obligation_component.Land_sale_number, obligation_component.Land_sale_offering_id, obligation_component.Land_service_provided_by, obligation_component.Land_service_seq_no, obligation_component.Language, obligation_component.Lithology_log_id, obligation_component.Lith_log_source, obligation_component.Lr_termination_id, obligation_component.Lr_termination_seq_no, obligation_component.Notification_id, obligation_component.Obligation_component_type, obligation_component.Paleo_summary_id, obligation_component.Pden_id, obligation_component.Pden_source, obligation_component.Pden_subtype, obligation_component.Physical_item_id, obligation_component.Pool_id, obligation_component.Ppdm_guid, obligation_component.Ppdm_system_id, obligation_component.Ppdm_table_name, obligation_component.Product_type, obligation_component.Prod_string_id, obligation_component.Prod_string_source, obligation_component.Project_id, obligation_component.Rate_schedule_id, obligation_component.Referenced_guid, obligation_component.Referenced_system_id, obligation_component.Referenced_table_name, obligation_component.Remark, obligation_component.Resent_id, obligation_component.Reserve_class_id, obligation_component.Sample_anal_source, obligation_component.Seis_license_id, obligation_component.Seis_point_flow_id, obligation_component.Seis_point_id, obligation_component.Seis_service_ba_id, obligation_component.Seis_service_seq_no, obligation_component.Seis_set_id, obligation_component.Seis_set_subtype, obligation_component.Seis_transaction_id, obligation_component.Seis_transaction_type, obligation_component.Sf_subtype, obligation_component.Source, obligation_component.Spatial_description_id, obligation_component.Spatial_obs_no, obligation_component.Store_id, obligation_component.Store_structure_obs_no, obligation_component.Strat_name_set_id, obligation_component.Strat_unit_id, obligation_component.Support_facility_id, obligation_component.Sw_application_id, obligation_component.Thesaurus_id, obligation_component.Thesaurus_word_id, obligation_component.Uwi, obligation_component.Well_activity_set_id, obligation_component.Well_activity_source, obligation_component.Well_activity_type_id, obligation_component.Well_license_id, obligation_component.Well_license_source, obligation_component.Well_service_ba_id, obligation_component.Well_service_seq_no, obligation_component.Well_set_id, obligation_component.Work_order_delivery_id, obligation_component.Work_order_id, obligation_component.Row_changed_by, obligation_component.Row_changed_date, obligation_component.Row_created_by, obligation_component.Row_effective_date, obligation_component.Row_expiry_date, obligation_component.Row_quality, obligation_component.Obligation_id)
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

func PatchObligationComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update obligation_component set "
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
		} else if key == "critical_date" || key == "effective_date" || key == "expiry_date" || key == "fulfilled_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where obligation_id = :id"

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

func DeleteObligationComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var obligation_component dto.Obligation_component
	obligation_component.Obligation_id = id

	stmt, err := db.Prepare("delete from obligation_component where obligation_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(obligation_component.Obligation_id)
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


