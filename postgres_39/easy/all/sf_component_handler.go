package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_component

	for rows.Next() {
		var sf_component dto.Sf_component
		if err := rows.Scan(&sf_component.Sf_subtype, &sf_component.Support_facility_id, &sf_component.Use_id, &sf_component.Component_obs_no, &sf_component.Active_ind, &sf_component.Area_id, &sf_component.Area_type, &sf_component.Ba_licensee_ba_id, &sf_component.Ba_license_id, &sf_component.Business_associate_id, &sf_component.Consent_id, &sf_component.Consult_id, &sf_component.Contest_id, &sf_component.Contract_id, &sf_component.Effective_date, &sf_component.Expiry_date, &sf_component.Facility_id, &sf_component.Facility_license_id, &sf_component.Facility_type, &sf_component.Field_id, &sf_component.Field_station_id, &sf_component.Finance_id, &sf_component.Hse_incident_id, &sf_component.Information_item_id, &sf_component.Info_item_subtype, &sf_component.Instrument_id, &sf_component.Interest_set_id, &sf_component.Interest_set_seq_no, &sf_component.Land_right_id, &sf_component.Land_right_subtype, &sf_component.Land_sale_jurisdiction, &sf_component.Land_sale_number, &sf_component.Obligation_id, &sf_component.Obligation_seq_no, &sf_component.Pden_id, &sf_component.Pden_source, &sf_component.Pden_subtype, &sf_component.Pool_id, &sf_component.Ppdm_guid, &sf_component.Ppdm_system_id, &sf_component.Ppdm_table_name, &sf_component.Prod_string_id, &sf_component.Prod_string_source, &sf_component.Project_id, &sf_component.Pr_str_form_obs_no, &sf_component.Rate_schedule_id, &sf_component.Remark, &sf_component.Seis_license_id, &sf_component.Seis_set_id, &sf_component.Seis_set_subtype, &sf_component.Sf_component_type, &sf_component.Source, &sf_component.Source_document_id, &sf_component.Sw_application_id, &sf_component.Uwi, &sf_component.Well_license_id, &sf_component.Well_license_source, &sf_component.Well_set_id, &sf_component.Work_order_id, &sf_component.Row_changed_by, &sf_component.Row_changed_date, &sf_component.Row_created_by, &sf_component.Row_created_date, &sf_component.Row_effective_date, &sf_component.Row_expiry_date, &sf_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_component to result
		result = append(result, sf_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfComponent(c *fiber.Ctx) error {
	var sf_component dto.Sf_component

	setDefaults(&sf_component)

	if err := json.Unmarshal(c.Body(), &sf_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66)")
	if err != nil {
		return err
	}
	sf_component.Row_created_date = formatDate(sf_component.Row_created_date)
	sf_component.Row_changed_date = formatDate(sf_component.Row_changed_date)
	sf_component.Effective_date = formatDateString(sf_component.Effective_date)
	sf_component.Expiry_date = formatDateString(sf_component.Expiry_date)
	sf_component.Row_effective_date = formatDateString(sf_component.Row_effective_date)
	sf_component.Row_expiry_date = formatDateString(sf_component.Row_expiry_date)






	rows, err := stmt.Exec(sf_component.Sf_subtype, sf_component.Support_facility_id, sf_component.Use_id, sf_component.Component_obs_no, sf_component.Active_ind, sf_component.Area_id, sf_component.Area_type, sf_component.Ba_licensee_ba_id, sf_component.Ba_license_id, sf_component.Business_associate_id, sf_component.Consent_id, sf_component.Consult_id, sf_component.Contest_id, sf_component.Contract_id, sf_component.Effective_date, sf_component.Expiry_date, sf_component.Facility_id, sf_component.Facility_license_id, sf_component.Facility_type, sf_component.Field_id, sf_component.Field_station_id, sf_component.Finance_id, sf_component.Hse_incident_id, sf_component.Information_item_id, sf_component.Info_item_subtype, sf_component.Instrument_id, sf_component.Interest_set_id, sf_component.Interest_set_seq_no, sf_component.Land_right_id, sf_component.Land_right_subtype, sf_component.Land_sale_jurisdiction, sf_component.Land_sale_number, sf_component.Obligation_id, sf_component.Obligation_seq_no, sf_component.Pden_id, sf_component.Pden_source, sf_component.Pden_subtype, sf_component.Pool_id, sf_component.Ppdm_guid, sf_component.Ppdm_system_id, sf_component.Ppdm_table_name, sf_component.Prod_string_id, sf_component.Prod_string_source, sf_component.Project_id, sf_component.Pr_str_form_obs_no, sf_component.Rate_schedule_id, sf_component.Remark, sf_component.Seis_license_id, sf_component.Seis_set_id, sf_component.Seis_set_subtype, sf_component.Sf_component_type, sf_component.Source, sf_component.Source_document_id, sf_component.Sw_application_id, sf_component.Uwi, sf_component.Well_license_id, sf_component.Well_license_source, sf_component.Well_set_id, sf_component.Work_order_id, sf_component.Row_changed_by, sf_component.Row_changed_date, sf_component.Row_created_by, sf_component.Row_created_date, sf_component.Row_effective_date, sf_component.Row_expiry_date, sf_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfComponent(c *fiber.Ctx) error {
	var sf_component dto.Sf_component

	setDefaults(&sf_component)

	if err := json.Unmarshal(c.Body(), &sf_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_component.Sf_subtype = id

    if sf_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_component set support_facility_id = :1, use_id = :2, component_obs_no = :3, active_ind = :4, area_id = :5, area_type = :6, ba_licensee_ba_id = :7, ba_license_id = :8, business_associate_id = :9, consent_id = :10, consult_id = :11, contest_id = :12, contract_id = :13, effective_date = :14, expiry_date = :15, facility_id = :16, facility_license_id = :17, facility_type = :18, field_id = :19, field_station_id = :20, finance_id = :21, hse_incident_id = :22, information_item_id = :23, info_item_subtype = :24, instrument_id = :25, interest_set_id = :26, interest_set_seq_no = :27, land_right_id = :28, land_right_subtype = :29, land_sale_jurisdiction = :30, land_sale_number = :31, obligation_id = :32, obligation_seq_no = :33, pden_id = :34, pden_source = :35, pden_subtype = :36, pool_id = :37, ppdm_guid = :38, ppdm_system_id = :39, ppdm_table_name = :40, prod_string_id = :41, prod_string_source = :42, project_id = :43, pr_str_form_obs_no = :44, rate_schedule_id = :45, remark = :46, seis_license_id = :47, seis_set_id = :48, seis_set_subtype = :49, sf_component_type = :50, source = :51, source_document_id = :52, sw_application_id = :53, uwi = :54, well_license_id = :55, well_license_source = :56, well_set_id = :57, work_order_id = :58, row_changed_by = :59, row_changed_date = :60, row_created_by = :61, row_effective_date = :62, row_expiry_date = :63, row_quality = :64 where sf_subtype = :66")
	if err != nil {
		return err
	}

	sf_component.Row_changed_date = formatDate(sf_component.Row_changed_date)
	sf_component.Effective_date = formatDateString(sf_component.Effective_date)
	sf_component.Expiry_date = formatDateString(sf_component.Expiry_date)
	sf_component.Row_effective_date = formatDateString(sf_component.Row_effective_date)
	sf_component.Row_expiry_date = formatDateString(sf_component.Row_expiry_date)






	rows, err := stmt.Exec(sf_component.Support_facility_id, sf_component.Use_id, sf_component.Component_obs_no, sf_component.Active_ind, sf_component.Area_id, sf_component.Area_type, sf_component.Ba_licensee_ba_id, sf_component.Ba_license_id, sf_component.Business_associate_id, sf_component.Consent_id, sf_component.Consult_id, sf_component.Contest_id, sf_component.Contract_id, sf_component.Effective_date, sf_component.Expiry_date, sf_component.Facility_id, sf_component.Facility_license_id, sf_component.Facility_type, sf_component.Field_id, sf_component.Field_station_id, sf_component.Finance_id, sf_component.Hse_incident_id, sf_component.Information_item_id, sf_component.Info_item_subtype, sf_component.Instrument_id, sf_component.Interest_set_id, sf_component.Interest_set_seq_no, sf_component.Land_right_id, sf_component.Land_right_subtype, sf_component.Land_sale_jurisdiction, sf_component.Land_sale_number, sf_component.Obligation_id, sf_component.Obligation_seq_no, sf_component.Pden_id, sf_component.Pden_source, sf_component.Pden_subtype, sf_component.Pool_id, sf_component.Ppdm_guid, sf_component.Ppdm_system_id, sf_component.Ppdm_table_name, sf_component.Prod_string_id, sf_component.Prod_string_source, sf_component.Project_id, sf_component.Pr_str_form_obs_no, sf_component.Rate_schedule_id, sf_component.Remark, sf_component.Seis_license_id, sf_component.Seis_set_id, sf_component.Seis_set_subtype, sf_component.Sf_component_type, sf_component.Source, sf_component.Source_document_id, sf_component.Sw_application_id, sf_component.Uwi, sf_component.Well_license_id, sf_component.Well_license_source, sf_component.Well_set_id, sf_component.Work_order_id, sf_component.Row_changed_by, sf_component.Row_changed_date, sf_component.Row_created_by, sf_component.Row_effective_date, sf_component.Row_expiry_date, sf_component.Row_quality, sf_component.Sf_subtype)
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

func PatchSfComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_component set "
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
	query += " where sf_subtype = :id"

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

func DeleteSfComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_component dto.Sf_component
	sf_component.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_component where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_component.Sf_subtype)
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


