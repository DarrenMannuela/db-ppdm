package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisTransComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_trans_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_trans_component

	for rows.Next() {
		var seis_trans_component dto.Seis_trans_component
		if err := rows.Scan(&seis_trans_component.Seis_transaction_id, &seis_trans_component.Transaction_type, &seis_trans_component.Transaction_component_id, &seis_trans_component.Active_ind, &seis_trans_component.Activity_obs_no, &seis_trans_component.Analysis_id, &seis_trans_component.Application_id, &seis_trans_component.Area_id, &seis_trans_component.Area_type, &seis_trans_component.Authority_id, &seis_trans_component.Authorize_id, &seis_trans_component.Ba_organization_id, &seis_trans_component.Ba_organization_seq_no, &seis_trans_component.Business_associate_id, &seis_trans_component.Catalogue_additive_id, &seis_trans_component.Catalogue_equip_id, &seis_trans_component.Circ_id, &seis_trans_component.Classification_system_id, &seis_trans_component.Class_level_id, &seis_trans_component.Component_type, &seis_trans_component.Consent_id, &seis_trans_component.Consult_id, &seis_trans_component.Contest_id, &seis_trans_component.Contract_id, &seis_trans_component.Currency_conversion, &seis_trans_component.Currency_ouom, &seis_trans_component.Discount_rate, &seis_trans_component.Ecozone_id, &seis_trans_component.Effective_date, &seis_trans_component.Employee_ba_id, &seis_trans_component.Employee_obs_no, &seis_trans_component.Employer_ba_id, &seis_trans_component.Entitlement_id, &seis_trans_component.Equipment_id, &seis_trans_component.Expiry_date, &seis_trans_component.Facility_id, &seis_trans_component.Facility_type, &seis_trans_component.Field_id, &seis_trans_component.Field_station_id, &seis_trans_component.Finance_id, &seis_trans_component.Fossil_id, &seis_trans_component.Incident_id, &seis_trans_component.Incident_set_id, &seis_trans_component.Incident_type, &seis_trans_component.Information_item_id, &seis_trans_component.Info_item_subtype, &seis_trans_component.Inspection_id, &seis_trans_component.Instrument_id, &seis_trans_component.Interest_set_id, &seis_trans_component.Interest_set_seq_no, &seis_trans_component.Jurisdiction, &seis_trans_component.Land_right_id, &seis_trans_component.Land_right_subtype, &seis_trans_component.Land_sale_number, &seis_trans_component.Language, &seis_trans_component.Length, &seis_trans_component.Length_ouom, &seis_trans_component.Lithology_log_id, &seis_trans_component.Lith_log_source, &seis_trans_component.Notification_id, &seis_trans_component.Obligation_id, &seis_trans_component.Obligation_seq_no, &seis_trans_component.Paleo_summary_id, &seis_trans_component.Pden_id, &seis_trans_component.Pden_source, &seis_trans_component.Pden_subtype, &seis_trans_component.Physical_item_id, &seis_trans_component.Pool_id, &seis_trans_component.Ppdm_guid, &seis_trans_component.Ppdm_system_id, &seis_trans_component.Ppdm_table_name, &seis_trans_component.Price_per_length, &seis_trans_component.Product_type, &seis_trans_component.Prod_string_id, &seis_trans_component.Prod_string_source, &seis_trans_component.Project_id, &seis_trans_component.Rate_schedule_id, &seis_trans_component.Referenced_guid, &seis_trans_component.Referenced_system_id, &seis_trans_component.Referenced_table_name, &seis_trans_component.Remark, &seis_trans_component.Resent_id, &seis_trans_component.Reserve_class_id, &seis_trans_component.Sample_anal_source, &seis_trans_component.Seis_set_id, &seis_trans_component.Seis_set_subtype, &seis_trans_component.Sf_subtype, &seis_trans_component.Source, &seis_trans_component.Spatial_description_id, &seis_trans_component.Spatial_obs_no, &seis_trans_component.Store_id, &seis_trans_component.Store_structure_obs_no, &seis_trans_component.Strat_name_set_id, &seis_trans_component.Strat_unit_id, &seis_trans_component.Support_facility_id, &seis_trans_component.Sw_application_id, &seis_trans_component.Thesaurus_id, &seis_trans_component.Thesaurus_word_id, &seis_trans_component.Total_price, &seis_trans_component.Transaction_status, &seis_trans_component.Uwi, &seis_trans_component.Well_activity_set_id, &seis_trans_component.Well_activity_source, &seis_trans_component.Well_activity_type_id, &seis_trans_component.Well_set_id, &seis_trans_component.Work_order_id, &seis_trans_component.Row_changed_by, &seis_trans_component.Row_changed_date, &seis_trans_component.Row_created_by, &seis_trans_component.Row_created_date, &seis_trans_component.Row_effective_date, &seis_trans_component.Row_expiry_date, &seis_trans_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_trans_component to result
		result = append(result, seis_trans_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisTransComponent(c *fiber.Ctx) error {
	var seis_trans_component dto.Seis_trans_component

	setDefaults(&seis_trans_component)

	if err := json.Unmarshal(c.Body(), &seis_trans_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_trans_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103, :104, :105, :106, :107, :108, :109, :110, :111, :112, :113)")
	if err != nil {
		return err
	}
	seis_trans_component.Row_created_date = formatDate(seis_trans_component.Row_created_date)
	seis_trans_component.Row_changed_date = formatDate(seis_trans_component.Row_changed_date)
	seis_trans_component.Effective_date = formatDateString(seis_trans_component.Effective_date)
	seis_trans_component.Expiry_date = formatDateString(seis_trans_component.Expiry_date)
	seis_trans_component.Row_effective_date = formatDateString(seis_trans_component.Row_effective_date)
	seis_trans_component.Row_expiry_date = formatDateString(seis_trans_component.Row_expiry_date)






	rows, err := stmt.Exec(seis_trans_component.Seis_transaction_id, seis_trans_component.Transaction_type, seis_trans_component.Transaction_component_id, seis_trans_component.Active_ind, seis_trans_component.Activity_obs_no, seis_trans_component.Analysis_id, seis_trans_component.Application_id, seis_trans_component.Area_id, seis_trans_component.Area_type, seis_trans_component.Authority_id, seis_trans_component.Authorize_id, seis_trans_component.Ba_organization_id, seis_trans_component.Ba_organization_seq_no, seis_trans_component.Business_associate_id, seis_trans_component.Catalogue_additive_id, seis_trans_component.Catalogue_equip_id, seis_trans_component.Circ_id, seis_trans_component.Classification_system_id, seis_trans_component.Class_level_id, seis_trans_component.Component_type, seis_trans_component.Consent_id, seis_trans_component.Consult_id, seis_trans_component.Contest_id, seis_trans_component.Contract_id, seis_trans_component.Currency_conversion, seis_trans_component.Currency_ouom, seis_trans_component.Discount_rate, seis_trans_component.Ecozone_id, seis_trans_component.Effective_date, seis_trans_component.Employee_ba_id, seis_trans_component.Employee_obs_no, seis_trans_component.Employer_ba_id, seis_trans_component.Entitlement_id, seis_trans_component.Equipment_id, seis_trans_component.Expiry_date, seis_trans_component.Facility_id, seis_trans_component.Facility_type, seis_trans_component.Field_id, seis_trans_component.Field_station_id, seis_trans_component.Finance_id, seis_trans_component.Fossil_id, seis_trans_component.Incident_id, seis_trans_component.Incident_set_id, seis_trans_component.Incident_type, seis_trans_component.Information_item_id, seis_trans_component.Info_item_subtype, seis_trans_component.Inspection_id, seis_trans_component.Instrument_id, seis_trans_component.Interest_set_id, seis_trans_component.Interest_set_seq_no, seis_trans_component.Jurisdiction, seis_trans_component.Land_right_id, seis_trans_component.Land_right_subtype, seis_trans_component.Land_sale_number, seis_trans_component.Language, seis_trans_component.Length, seis_trans_component.Length_ouom, seis_trans_component.Lithology_log_id, seis_trans_component.Lith_log_source, seis_trans_component.Notification_id, seis_trans_component.Obligation_id, seis_trans_component.Obligation_seq_no, seis_trans_component.Paleo_summary_id, seis_trans_component.Pden_id, seis_trans_component.Pden_source, seis_trans_component.Pden_subtype, seis_trans_component.Physical_item_id, seis_trans_component.Pool_id, seis_trans_component.Ppdm_guid, seis_trans_component.Ppdm_system_id, seis_trans_component.Ppdm_table_name, seis_trans_component.Price_per_length, seis_trans_component.Product_type, seis_trans_component.Prod_string_id, seis_trans_component.Prod_string_source, seis_trans_component.Project_id, seis_trans_component.Rate_schedule_id, seis_trans_component.Referenced_guid, seis_trans_component.Referenced_system_id, seis_trans_component.Referenced_table_name, seis_trans_component.Remark, seis_trans_component.Resent_id, seis_trans_component.Reserve_class_id, seis_trans_component.Sample_anal_source, seis_trans_component.Seis_set_id, seis_trans_component.Seis_set_subtype, seis_trans_component.Sf_subtype, seis_trans_component.Source, seis_trans_component.Spatial_description_id, seis_trans_component.Spatial_obs_no, seis_trans_component.Store_id, seis_trans_component.Store_structure_obs_no, seis_trans_component.Strat_name_set_id, seis_trans_component.Strat_unit_id, seis_trans_component.Support_facility_id, seis_trans_component.Sw_application_id, seis_trans_component.Thesaurus_id, seis_trans_component.Thesaurus_word_id, seis_trans_component.Total_price, seis_trans_component.Transaction_status, seis_trans_component.Uwi, seis_trans_component.Well_activity_set_id, seis_trans_component.Well_activity_source, seis_trans_component.Well_activity_type_id, seis_trans_component.Well_set_id, seis_trans_component.Work_order_id, seis_trans_component.Row_changed_by, seis_trans_component.Row_changed_date, seis_trans_component.Row_created_by, seis_trans_component.Row_created_date, seis_trans_component.Row_effective_date, seis_trans_component.Row_expiry_date, seis_trans_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisTransComponent(c *fiber.Ctx) error {
	var seis_trans_component dto.Seis_trans_component

	setDefaults(&seis_trans_component)

	if err := json.Unmarshal(c.Body(), &seis_trans_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_trans_component.Seis_transaction_id = id

    if seis_trans_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_trans_component set transaction_type = :1, transaction_component_id = :2, active_ind = :3, activity_obs_no = :4, analysis_id = :5, application_id = :6, area_id = :7, area_type = :8, authority_id = :9, authorize_id = :10, ba_organization_id = :11, ba_organization_seq_no = :12, business_associate_id = :13, catalogue_additive_id = :14, catalogue_equip_id = :15, circ_id = :16, classification_system_id = :17, class_level_id = :18, component_type = :19, consent_id = :20, consult_id = :21, contest_id = :22, contract_id = :23, currency_conversion = :24, currency_ouom = :25, discount_rate = :26, ecozone_id = :27, effective_date = :28, employee_ba_id = :29, employee_obs_no = :30, employer_ba_id = :31, entitlement_id = :32, equipment_id = :33, expiry_date = :34, facility_id = :35, facility_type = :36, field_id = :37, field_station_id = :38, finance_id = :39, fossil_id = :40, incident_id = :41, incident_set_id = :42, incident_type = :43, information_item_id = :44, info_item_subtype = :45, inspection_id = :46, instrument_id = :47, interest_set_id = :48, interest_set_seq_no = :49, jurisdiction = :50, land_right_id = :51, land_right_subtype = :52, land_sale_number = :53, language = :54, length = :55, length_ouom = :56, lithology_log_id = :57, lith_log_source = :58, notification_id = :59, obligation_id = :60, obligation_seq_no = :61, paleo_summary_id = :62, pden_id = :63, pden_source = :64, pden_subtype = :65, physical_item_id = :66, pool_id = :67, ppdm_guid = :68, ppdm_system_id = :69, ppdm_table_name = :70, price_per_length = :71, product_type = :72, prod_string_id = :73, prod_string_source = :74, project_id = :75, rate_schedule_id = :76, referenced_guid = :77, referenced_system_id = :78, referenced_table_name = :79, remark = :80, resent_id = :81, reserve_class_id = :82, sample_anal_source = :83, seis_set_id = :84, seis_set_subtype = :85, sf_subtype = :86, source = :87, spatial_description_id = :88, spatial_obs_no = :89, store_id = :90, store_structure_obs_no = :91, strat_name_set_id = :92, strat_unit_id = :93, support_facility_id = :94, sw_application_id = :95, thesaurus_id = :96, thesaurus_word_id = :97, total_price = :98, transaction_status = :99, uwi = :100, well_activity_set_id = :101, well_activity_source = :102, well_activity_type_id = :103, well_set_id = :104, work_order_id = :105, row_changed_by = :106, row_changed_date = :107, row_created_by = :108, row_effective_date = :109, row_expiry_date = :110, row_quality = :111 where seis_transaction_id = :113")
	if err != nil {
		return err
	}

	seis_trans_component.Row_changed_date = formatDate(seis_trans_component.Row_changed_date)
	seis_trans_component.Effective_date = formatDateString(seis_trans_component.Effective_date)
	seis_trans_component.Expiry_date = formatDateString(seis_trans_component.Expiry_date)
	seis_trans_component.Row_effective_date = formatDateString(seis_trans_component.Row_effective_date)
	seis_trans_component.Row_expiry_date = formatDateString(seis_trans_component.Row_expiry_date)






	rows, err := stmt.Exec(seis_trans_component.Transaction_type, seis_trans_component.Transaction_component_id, seis_trans_component.Active_ind, seis_trans_component.Activity_obs_no, seis_trans_component.Analysis_id, seis_trans_component.Application_id, seis_trans_component.Area_id, seis_trans_component.Area_type, seis_trans_component.Authority_id, seis_trans_component.Authorize_id, seis_trans_component.Ba_organization_id, seis_trans_component.Ba_organization_seq_no, seis_trans_component.Business_associate_id, seis_trans_component.Catalogue_additive_id, seis_trans_component.Catalogue_equip_id, seis_trans_component.Circ_id, seis_trans_component.Classification_system_id, seis_trans_component.Class_level_id, seis_trans_component.Component_type, seis_trans_component.Consent_id, seis_trans_component.Consult_id, seis_trans_component.Contest_id, seis_trans_component.Contract_id, seis_trans_component.Currency_conversion, seis_trans_component.Currency_ouom, seis_trans_component.Discount_rate, seis_trans_component.Ecozone_id, seis_trans_component.Effective_date, seis_trans_component.Employee_ba_id, seis_trans_component.Employee_obs_no, seis_trans_component.Employer_ba_id, seis_trans_component.Entitlement_id, seis_trans_component.Equipment_id, seis_trans_component.Expiry_date, seis_trans_component.Facility_id, seis_trans_component.Facility_type, seis_trans_component.Field_id, seis_trans_component.Field_station_id, seis_trans_component.Finance_id, seis_trans_component.Fossil_id, seis_trans_component.Incident_id, seis_trans_component.Incident_set_id, seis_trans_component.Incident_type, seis_trans_component.Information_item_id, seis_trans_component.Info_item_subtype, seis_trans_component.Inspection_id, seis_trans_component.Instrument_id, seis_trans_component.Interest_set_id, seis_trans_component.Interest_set_seq_no, seis_trans_component.Jurisdiction, seis_trans_component.Land_right_id, seis_trans_component.Land_right_subtype, seis_trans_component.Land_sale_number, seis_trans_component.Language, seis_trans_component.Length, seis_trans_component.Length_ouom, seis_trans_component.Lithology_log_id, seis_trans_component.Lith_log_source, seis_trans_component.Notification_id, seis_trans_component.Obligation_id, seis_trans_component.Obligation_seq_no, seis_trans_component.Paleo_summary_id, seis_trans_component.Pden_id, seis_trans_component.Pden_source, seis_trans_component.Pden_subtype, seis_trans_component.Physical_item_id, seis_trans_component.Pool_id, seis_trans_component.Ppdm_guid, seis_trans_component.Ppdm_system_id, seis_trans_component.Ppdm_table_name, seis_trans_component.Price_per_length, seis_trans_component.Product_type, seis_trans_component.Prod_string_id, seis_trans_component.Prod_string_source, seis_trans_component.Project_id, seis_trans_component.Rate_schedule_id, seis_trans_component.Referenced_guid, seis_trans_component.Referenced_system_id, seis_trans_component.Referenced_table_name, seis_trans_component.Remark, seis_trans_component.Resent_id, seis_trans_component.Reserve_class_id, seis_trans_component.Sample_anal_source, seis_trans_component.Seis_set_id, seis_trans_component.Seis_set_subtype, seis_trans_component.Sf_subtype, seis_trans_component.Source, seis_trans_component.Spatial_description_id, seis_trans_component.Spatial_obs_no, seis_trans_component.Store_id, seis_trans_component.Store_structure_obs_no, seis_trans_component.Strat_name_set_id, seis_trans_component.Strat_unit_id, seis_trans_component.Support_facility_id, seis_trans_component.Sw_application_id, seis_trans_component.Thesaurus_id, seis_trans_component.Thesaurus_word_id, seis_trans_component.Total_price, seis_trans_component.Transaction_status, seis_trans_component.Uwi, seis_trans_component.Well_activity_set_id, seis_trans_component.Well_activity_source, seis_trans_component.Well_activity_type_id, seis_trans_component.Well_set_id, seis_trans_component.Work_order_id, seis_trans_component.Row_changed_by, seis_trans_component.Row_changed_date, seis_trans_component.Row_created_by, seis_trans_component.Row_effective_date, seis_trans_component.Row_expiry_date, seis_trans_component.Row_quality, seis_trans_component.Seis_transaction_id)
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

func PatchSeisTransComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_trans_component set "
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
	query += " where seis_transaction_id = :id"

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

func DeleteSeisTransComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_trans_component dto.Seis_trans_component
	seis_trans_component.Seis_transaction_id = id

	stmt, err := db.Prepare("delete from seis_trans_component where seis_transaction_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_trans_component.Seis_transaction_id)
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


