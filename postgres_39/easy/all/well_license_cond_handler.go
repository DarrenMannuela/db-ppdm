package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellLicenseCond(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_license_cond")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_license_cond

	for rows.Next() {
		var well_license_cond dto.Well_license_cond
		if err := rows.Scan(&well_license_cond.Uwi, &well_license_cond.License_id, &well_license_cond.Well_lic_source, &well_license_cond.Condition_id, &well_license_cond.Active_ind, &well_license_cond.Condition_code, &well_license_cond.Condition_type, &well_license_cond.Condition_value, &well_license_cond.Condition_value_ouom, &well_license_cond.Condition_value_uom, &well_license_cond.Contact_ba_id, &well_license_cond.Date_format_desc, &well_license_cond.Description, &well_license_cond.Due_condition, &well_license_cond.Due_date, &well_license_cond.Due_frequency, &well_license_cond.Due_term, &well_license_cond.Due_term_uom, &well_license_cond.Effective_date, &well_license_cond.Exempt_ind, &well_license_cond.Expiry_date, &well_license_cond.Fulfilled_by_ba_id, &well_license_cond.Fulfilled_date, &well_license_cond.Fulfilled_ind, &well_license_cond.Ppdm_guid, &well_license_cond.Product_type, &well_license_cond.Remark, &well_license_cond.Restriction_id, &well_license_cond.Restriction_version, &well_license_cond.Source, &well_license_cond.Strat_name_set_id, &well_license_cond.Strat_unit_id, &well_license_cond.Row_changed_by, &well_license_cond.Row_changed_date, &well_license_cond.Row_created_by, &well_license_cond.Row_created_date, &well_license_cond.Row_effective_date, &well_license_cond.Row_expiry_date, &well_license_cond.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_license_cond to result
		result = append(result, well_license_cond)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellLicenseCond(c *fiber.Ctx) error {
	var well_license_cond dto.Well_license_cond

	setDefaults(&well_license_cond)

	if err := json.Unmarshal(c.Body(), &well_license_cond); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_license_cond values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39)")
	if err != nil {
		return err
	}
	well_license_cond.Row_created_date = formatDate(well_license_cond.Row_created_date)
	well_license_cond.Row_changed_date = formatDate(well_license_cond.Row_changed_date)
	well_license_cond.Date_format_desc = formatDateString(well_license_cond.Date_format_desc)
	well_license_cond.Due_date = formatDateString(well_license_cond.Due_date)
	well_license_cond.Effective_date = formatDateString(well_license_cond.Effective_date)
	well_license_cond.Expiry_date = formatDateString(well_license_cond.Expiry_date)
	well_license_cond.Fulfilled_date = formatDateString(well_license_cond.Fulfilled_date)
	well_license_cond.Row_effective_date = formatDateString(well_license_cond.Row_effective_date)
	well_license_cond.Row_expiry_date = formatDateString(well_license_cond.Row_expiry_date)






	rows, err := stmt.Exec(well_license_cond.Uwi, well_license_cond.License_id, well_license_cond.Well_lic_source, well_license_cond.Condition_id, well_license_cond.Active_ind, well_license_cond.Condition_code, well_license_cond.Condition_type, well_license_cond.Condition_value, well_license_cond.Condition_value_ouom, well_license_cond.Condition_value_uom, well_license_cond.Contact_ba_id, well_license_cond.Date_format_desc, well_license_cond.Description, well_license_cond.Due_condition, well_license_cond.Due_date, well_license_cond.Due_frequency, well_license_cond.Due_term, well_license_cond.Due_term_uom, well_license_cond.Effective_date, well_license_cond.Exempt_ind, well_license_cond.Expiry_date, well_license_cond.Fulfilled_by_ba_id, well_license_cond.Fulfilled_date, well_license_cond.Fulfilled_ind, well_license_cond.Ppdm_guid, well_license_cond.Product_type, well_license_cond.Remark, well_license_cond.Restriction_id, well_license_cond.Restriction_version, well_license_cond.Source, well_license_cond.Strat_name_set_id, well_license_cond.Strat_unit_id, well_license_cond.Row_changed_by, well_license_cond.Row_changed_date, well_license_cond.Row_created_by, well_license_cond.Row_created_date, well_license_cond.Row_effective_date, well_license_cond.Row_expiry_date, well_license_cond.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellLicenseCond(c *fiber.Ctx) error {
	var well_license_cond dto.Well_license_cond

	setDefaults(&well_license_cond)

	if err := json.Unmarshal(c.Body(), &well_license_cond); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_license_cond.Uwi = id

    if well_license_cond.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_license_cond set license_id = :1, well_lic_source = :2, condition_id = :3, active_ind = :4, condition_code = :5, condition_type = :6, condition_value = :7, condition_value_ouom = :8, condition_value_uom = :9, contact_ba_id = :10, date_format_desc = :11, description = :12, due_condition = :13, due_date = :14, due_frequency = :15, due_term = :16, due_term_uom = :17, effective_date = :18, exempt_ind = :19, expiry_date = :20, fulfilled_by_ba_id = :21, fulfilled_date = :22, fulfilled_ind = :23, ppdm_guid = :24, product_type = :25, remark = :26, restriction_id = :27, restriction_version = :28, source = :29, strat_name_set_id = :30, strat_unit_id = :31, row_changed_by = :32, row_changed_date = :33, row_created_by = :34, row_effective_date = :35, row_expiry_date = :36, row_quality = :37 where uwi = :39")
	if err != nil {
		return err
	}

	well_license_cond.Row_changed_date = formatDate(well_license_cond.Row_changed_date)
	well_license_cond.Date_format_desc = formatDateString(well_license_cond.Date_format_desc)
	well_license_cond.Due_date = formatDateString(well_license_cond.Due_date)
	well_license_cond.Effective_date = formatDateString(well_license_cond.Effective_date)
	well_license_cond.Expiry_date = formatDateString(well_license_cond.Expiry_date)
	well_license_cond.Fulfilled_date = formatDateString(well_license_cond.Fulfilled_date)
	well_license_cond.Row_effective_date = formatDateString(well_license_cond.Row_effective_date)
	well_license_cond.Row_expiry_date = formatDateString(well_license_cond.Row_expiry_date)






	rows, err := stmt.Exec(well_license_cond.License_id, well_license_cond.Well_lic_source, well_license_cond.Condition_id, well_license_cond.Active_ind, well_license_cond.Condition_code, well_license_cond.Condition_type, well_license_cond.Condition_value, well_license_cond.Condition_value_ouom, well_license_cond.Condition_value_uom, well_license_cond.Contact_ba_id, well_license_cond.Date_format_desc, well_license_cond.Description, well_license_cond.Due_condition, well_license_cond.Due_date, well_license_cond.Due_frequency, well_license_cond.Due_term, well_license_cond.Due_term_uom, well_license_cond.Effective_date, well_license_cond.Exempt_ind, well_license_cond.Expiry_date, well_license_cond.Fulfilled_by_ba_id, well_license_cond.Fulfilled_date, well_license_cond.Fulfilled_ind, well_license_cond.Ppdm_guid, well_license_cond.Product_type, well_license_cond.Remark, well_license_cond.Restriction_id, well_license_cond.Restriction_version, well_license_cond.Source, well_license_cond.Strat_name_set_id, well_license_cond.Strat_unit_id, well_license_cond.Row_changed_by, well_license_cond.Row_changed_date, well_license_cond.Row_created_by, well_license_cond.Row_effective_date, well_license_cond.Row_expiry_date, well_license_cond.Row_quality, well_license_cond.Uwi)
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

func PatchWellLicenseCond(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_license_cond set "
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
		} else if key == "date_format_desc" || key == "due_date" || key == "effective_date" || key == "expiry_date" || key == "fulfilled_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where uwi = :id"

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

func DeleteWellLicenseCond(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_license_cond dto.Well_license_cond
	well_license_cond.Uwi = id

	stmt, err := db.Prepare("delete from well_license_cond where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_license_cond.Uwi)
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


