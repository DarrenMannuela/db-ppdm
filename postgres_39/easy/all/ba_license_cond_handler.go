package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetBaLicenseCond(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ba_license_cond")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ba_license_cond

	for rows.Next() {
		var ba_license_cond dto.Ba_license_cond
		if err := rows.Scan(&ba_license_cond.Licensee_ba_id, &ba_license_cond.License_id, &ba_license_cond.Condition_id, &ba_license_cond.Active_ind, &ba_license_cond.Condition_code, &ba_license_cond.Condition_desc, &ba_license_cond.Condition_type, &ba_license_cond.Condition_value, &ba_license_cond.Condition_value_ouom, &ba_license_cond.Condition_value_uom, &ba_license_cond.Contact_ba_id, &ba_license_cond.Currency_conversion, &ba_license_cond.Date_format_desc, &ba_license_cond.Description, &ba_license_cond.Due_condition, &ba_license_cond.Due_date, &ba_license_cond.Due_frequency, &ba_license_cond.Due_term, &ba_license_cond.Due_term_uom, &ba_license_cond.Effective_date, &ba_license_cond.Exempt_ind, &ba_license_cond.Expiry_date, &ba_license_cond.Fulfilled_by_ba_id, &ba_license_cond.Fulfilled_date, &ba_license_cond.Fulfilled_ind, &ba_license_cond.Granted_by_ba_id, &ba_license_cond.License_type, &ba_license_cond.Max_value, &ba_license_cond.Max_value_ouom, &ba_license_cond.Max_value_uom, &ba_license_cond.Min_value, &ba_license_cond.Min_value_ouom, &ba_license_cond.Min_value_uom, &ba_license_cond.Ppdm_guid, &ba_license_cond.Remark, &ba_license_cond.Restriction_id, &ba_license_cond.Restriction_version, &ba_license_cond.Source, &ba_license_cond.Row_changed_by, &ba_license_cond.Row_changed_date, &ba_license_cond.Row_created_by, &ba_license_cond.Row_created_date, &ba_license_cond.Row_effective_date, &ba_license_cond.Row_expiry_date, &ba_license_cond.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ba_license_cond to result
		result = append(result, ba_license_cond)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetBaLicenseCond(c *fiber.Ctx) error {
	var ba_license_cond dto.Ba_license_cond

	setDefaults(&ba_license_cond)

	if err := json.Unmarshal(c.Body(), &ba_license_cond); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ba_license_cond values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45)")
	if err != nil {
		return err
	}
	ba_license_cond.Row_created_date = formatDate(ba_license_cond.Row_created_date)
	ba_license_cond.Row_changed_date = formatDate(ba_license_cond.Row_changed_date)
	ba_license_cond.Date_format_desc = formatDateString(ba_license_cond.Date_format_desc)
	ba_license_cond.Due_date = formatDateString(ba_license_cond.Due_date)
	ba_license_cond.Effective_date = formatDateString(ba_license_cond.Effective_date)
	ba_license_cond.Expiry_date = formatDateString(ba_license_cond.Expiry_date)
	ba_license_cond.Fulfilled_date = formatDateString(ba_license_cond.Fulfilled_date)
	ba_license_cond.Row_effective_date = formatDateString(ba_license_cond.Row_effective_date)
	ba_license_cond.Row_expiry_date = formatDateString(ba_license_cond.Row_expiry_date)






	rows, err := stmt.Exec(ba_license_cond.Licensee_ba_id, ba_license_cond.License_id, ba_license_cond.Condition_id, ba_license_cond.Active_ind, ba_license_cond.Condition_code, ba_license_cond.Condition_desc, ba_license_cond.Condition_type, ba_license_cond.Condition_value, ba_license_cond.Condition_value_ouom, ba_license_cond.Condition_value_uom, ba_license_cond.Contact_ba_id, ba_license_cond.Currency_conversion, ba_license_cond.Date_format_desc, ba_license_cond.Description, ba_license_cond.Due_condition, ba_license_cond.Due_date, ba_license_cond.Due_frequency, ba_license_cond.Due_term, ba_license_cond.Due_term_uom, ba_license_cond.Effective_date, ba_license_cond.Exempt_ind, ba_license_cond.Expiry_date, ba_license_cond.Fulfilled_by_ba_id, ba_license_cond.Fulfilled_date, ba_license_cond.Fulfilled_ind, ba_license_cond.Granted_by_ba_id, ba_license_cond.License_type, ba_license_cond.Max_value, ba_license_cond.Max_value_ouom, ba_license_cond.Max_value_uom, ba_license_cond.Min_value, ba_license_cond.Min_value_ouom, ba_license_cond.Min_value_uom, ba_license_cond.Ppdm_guid, ba_license_cond.Remark, ba_license_cond.Restriction_id, ba_license_cond.Restriction_version, ba_license_cond.Source, ba_license_cond.Row_changed_by, ba_license_cond.Row_changed_date, ba_license_cond.Row_created_by, ba_license_cond.Row_created_date, ba_license_cond.Row_effective_date, ba_license_cond.Row_expiry_date, ba_license_cond.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateBaLicenseCond(c *fiber.Ctx) error {
	var ba_license_cond dto.Ba_license_cond

	setDefaults(&ba_license_cond)

	if err := json.Unmarshal(c.Body(), &ba_license_cond); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ba_license_cond.Licensee_ba_id = id

    if ba_license_cond.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ba_license_cond set license_id = :1, condition_id = :2, active_ind = :3, condition_code = :4, condition_desc = :5, condition_type = :6, condition_value = :7, condition_value_ouom = :8, condition_value_uom = :9, contact_ba_id = :10, currency_conversion = :11, date_format_desc = :12, description = :13, due_condition = :14, due_date = :15, due_frequency = :16, due_term = :17, due_term_uom = :18, effective_date = :19, exempt_ind = :20, expiry_date = :21, fulfilled_by_ba_id = :22, fulfilled_date = :23, fulfilled_ind = :24, granted_by_ba_id = :25, license_type = :26, max_value = :27, max_value_ouom = :28, max_value_uom = :29, min_value = :30, min_value_ouom = :31, min_value_uom = :32, ppdm_guid = :33, remark = :34, restriction_id = :35, restriction_version = :36, source = :37, row_changed_by = :38, row_changed_date = :39, row_created_by = :40, row_effective_date = :41, row_expiry_date = :42, row_quality = :43 where licensee_ba_id = :45")
	if err != nil {
		return err
	}

	ba_license_cond.Row_changed_date = formatDate(ba_license_cond.Row_changed_date)
	ba_license_cond.Date_format_desc = formatDateString(ba_license_cond.Date_format_desc)
	ba_license_cond.Due_date = formatDateString(ba_license_cond.Due_date)
	ba_license_cond.Effective_date = formatDateString(ba_license_cond.Effective_date)
	ba_license_cond.Expiry_date = formatDateString(ba_license_cond.Expiry_date)
	ba_license_cond.Fulfilled_date = formatDateString(ba_license_cond.Fulfilled_date)
	ba_license_cond.Row_effective_date = formatDateString(ba_license_cond.Row_effective_date)
	ba_license_cond.Row_expiry_date = formatDateString(ba_license_cond.Row_expiry_date)






	rows, err := stmt.Exec(ba_license_cond.License_id, ba_license_cond.Condition_id, ba_license_cond.Active_ind, ba_license_cond.Condition_code, ba_license_cond.Condition_desc, ba_license_cond.Condition_type, ba_license_cond.Condition_value, ba_license_cond.Condition_value_ouom, ba_license_cond.Condition_value_uom, ba_license_cond.Contact_ba_id, ba_license_cond.Currency_conversion, ba_license_cond.Date_format_desc, ba_license_cond.Description, ba_license_cond.Due_condition, ba_license_cond.Due_date, ba_license_cond.Due_frequency, ba_license_cond.Due_term, ba_license_cond.Due_term_uom, ba_license_cond.Effective_date, ba_license_cond.Exempt_ind, ba_license_cond.Expiry_date, ba_license_cond.Fulfilled_by_ba_id, ba_license_cond.Fulfilled_date, ba_license_cond.Fulfilled_ind, ba_license_cond.Granted_by_ba_id, ba_license_cond.License_type, ba_license_cond.Max_value, ba_license_cond.Max_value_ouom, ba_license_cond.Max_value_uom, ba_license_cond.Min_value, ba_license_cond.Min_value_ouom, ba_license_cond.Min_value_uom, ba_license_cond.Ppdm_guid, ba_license_cond.Remark, ba_license_cond.Restriction_id, ba_license_cond.Restriction_version, ba_license_cond.Source, ba_license_cond.Row_changed_by, ba_license_cond.Row_changed_date, ba_license_cond.Row_created_by, ba_license_cond.Row_effective_date, ba_license_cond.Row_expiry_date, ba_license_cond.Row_quality, ba_license_cond.Licensee_ba_id)
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

func PatchBaLicenseCond(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ba_license_cond set "
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
	query += " where licensee_ba_id = :id"

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

func DeleteBaLicenseCond(c *fiber.Ctx) error {
	id := c.Params("id")
	var ba_license_cond dto.Ba_license_cond
	ba_license_cond.Licensee_ba_id = id

	stmt, err := db.Prepare("delete from ba_license_cond where licensee_ba_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ba_license_cond.Licensee_ba_id)
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


