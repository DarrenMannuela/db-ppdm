package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPden(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden

	for rows.Next() {
		var pden dto.Pden
		if err := rows.Scan(&pden.Pden_subtype, &pden.Pden_id, &pden.Source, &pden.Active_ind, &pden.Area_id, &pden.Area_type, &pden.Current_operator, &pden.Current_prod_str_name, &pden.Current_status_date, &pden.Current_well_str_number, &pden.Effective_date, &pden.Enhanced_recovery_type, &pden.Expiry_date, &pden.Field_id, &pden.Last_injection_date, &pden.Last_production_date, &pden.Last_reported_date, &pden.Location_desc, &pden.Location_desc_type, &pden.On_injection_date, &pden.On_production_date, &pden.Pden_long_name, &pden.Pden_short_name, &pden.Pden_status, &pden.Pden_status_type, &pden.Plot_name, &pden.Plot_symbol, &pden.Pool_id, &pden.Ppdm_guid, &pden.Primary_product, &pden.Production_method, &pden.Proprietary_ind, &pden.Remark, &pden.State_or_federal_waters, &pden.Strat_name_set_id, &pden.Strat_unit_id, &pden.String_serial_number, &pden.Row_changed_by, &pden.Row_changed_date, &pden.Row_created_by, &pden.Row_created_date, &pden.Row_effective_date, &pden.Row_expiry_date, &pden.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden to result
		result = append(result, pden)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPden(c *fiber.Ctx) error {
	var pden dto.Pden

	setDefaults(&pden)

	if err := json.Unmarshal(c.Body(), &pden); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44)")
	if err != nil {
		return err
	}
	pden.Row_created_date = formatDate(pden.Row_created_date)
	pden.Row_changed_date = formatDate(pden.Row_changed_date)
	pden.Current_status_date = formatDateString(pden.Current_status_date)
	pden.Effective_date = formatDateString(pden.Effective_date)
	pden.Expiry_date = formatDateString(pden.Expiry_date)
	pden.Last_injection_date = formatDateString(pden.Last_injection_date)
	pden.Last_production_date = formatDateString(pden.Last_production_date)
	pden.Last_reported_date = formatDateString(pden.Last_reported_date)
	pden.On_injection_date = formatDateString(pden.On_injection_date)
	pden.On_production_date = formatDateString(pden.On_production_date)
	pden.Row_effective_date = formatDateString(pden.Row_effective_date)
	pden.Row_expiry_date = formatDateString(pden.Row_expiry_date)






	rows, err := stmt.Exec(pden.Pden_subtype, pden.Pden_id, pden.Source, pden.Active_ind, pden.Area_id, pden.Area_type, pden.Current_operator, pden.Current_prod_str_name, pden.Current_status_date, pden.Current_well_str_number, pden.Effective_date, pden.Enhanced_recovery_type, pden.Expiry_date, pden.Field_id, pden.Last_injection_date, pden.Last_production_date, pden.Last_reported_date, pden.Location_desc, pden.Location_desc_type, pden.On_injection_date, pden.On_production_date, pden.Pden_long_name, pden.Pden_short_name, pden.Pden_status, pden.Pden_status_type, pden.Plot_name, pden.Plot_symbol, pden.Pool_id, pden.Ppdm_guid, pden.Primary_product, pden.Production_method, pden.Proprietary_ind, pden.Remark, pden.State_or_federal_waters, pden.Strat_name_set_id, pden.Strat_unit_id, pden.String_serial_number, pden.Row_changed_by, pden.Row_changed_date, pden.Row_created_by, pden.Row_created_date, pden.Row_effective_date, pden.Row_expiry_date, pden.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePden(c *fiber.Ctx) error {
	var pden dto.Pden

	setDefaults(&pden)

	if err := json.Unmarshal(c.Body(), &pden); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden.Pden_subtype = id

    if pden.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden set pden_id = :1, source = :2, active_ind = :3, area_id = :4, area_type = :5, current_operator = :6, current_prod_str_name = :7, current_status_date = :8, current_well_str_number = :9, effective_date = :10, enhanced_recovery_type = :11, expiry_date = :12, field_id = :13, last_injection_date = :14, last_production_date = :15, last_reported_date = :16, location_desc = :17, location_desc_type = :18, on_injection_date = :19, on_production_date = :20, pden_long_name = :21, pden_short_name = :22, pden_status = :23, pden_status_type = :24, plot_name = :25, plot_symbol = :26, pool_id = :27, ppdm_guid = :28, primary_product = :29, production_method = :30, proprietary_ind = :31, remark = :32, state_or_federal_waters = :33, strat_name_set_id = :34, strat_unit_id = :35, string_serial_number = :36, row_changed_by = :37, row_changed_date = :38, row_created_by = :39, row_effective_date = :40, row_expiry_date = :41, row_quality = :42 where pden_subtype = :44")
	if err != nil {
		return err
	}

	pden.Row_changed_date = formatDate(pden.Row_changed_date)
	pden.Current_status_date = formatDateString(pden.Current_status_date)
	pden.Effective_date = formatDateString(pden.Effective_date)
	pden.Expiry_date = formatDateString(pden.Expiry_date)
	pden.Last_injection_date = formatDateString(pden.Last_injection_date)
	pden.Last_production_date = formatDateString(pden.Last_production_date)
	pden.Last_reported_date = formatDateString(pden.Last_reported_date)
	pden.On_injection_date = formatDateString(pden.On_injection_date)
	pden.On_production_date = formatDateString(pden.On_production_date)
	pden.Row_effective_date = formatDateString(pden.Row_effective_date)
	pden.Row_expiry_date = formatDateString(pden.Row_expiry_date)






	rows, err := stmt.Exec(pden.Pden_id, pden.Source, pden.Active_ind, pden.Area_id, pden.Area_type, pden.Current_operator, pden.Current_prod_str_name, pden.Current_status_date, pden.Current_well_str_number, pden.Effective_date, pden.Enhanced_recovery_type, pden.Expiry_date, pden.Field_id, pden.Last_injection_date, pden.Last_production_date, pden.Last_reported_date, pden.Location_desc, pden.Location_desc_type, pden.On_injection_date, pden.On_production_date, pden.Pden_long_name, pden.Pden_short_name, pden.Pden_status, pden.Pden_status_type, pden.Plot_name, pden.Plot_symbol, pden.Pool_id, pden.Ppdm_guid, pden.Primary_product, pden.Production_method, pden.Proprietary_ind, pden.Remark, pden.State_or_federal_waters, pden.Strat_name_set_id, pden.Strat_unit_id, pden.String_serial_number, pden.Row_changed_by, pden.Row_changed_date, pden.Row_created_by, pden.Row_effective_date, pden.Row_expiry_date, pden.Row_quality, pden.Pden_subtype)
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

func PatchPden(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden set "
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
		} else if key == "current_status_date" || key == "effective_date" || key == "expiry_date" || key == "last_injection_date" || key == "last_production_date" || key == "last_reported_date" || key == "on_injection_date" || key == "on_production_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where pden_subtype = :id"

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

func DeletePden(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden dto.Pden
	pden.Pden_subtype = id

	stmt, err := db.Prepare("delete from pden where pden_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden.Pden_subtype)
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


