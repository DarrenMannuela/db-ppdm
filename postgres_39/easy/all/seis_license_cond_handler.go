package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisLicenseCond(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_license_cond")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_license_cond

	for rows.Next() {
		var seis_license_cond dto.Seis_license_cond
		if err := rows.Scan(&seis_license_cond.Seis_set_subtype, &seis_license_cond.Seis_set_id, &seis_license_cond.License_id, &seis_license_cond.Condition_id, &seis_license_cond.Active_ind, &seis_license_cond.Condition_code, &seis_license_cond.Condition_desc, &seis_license_cond.Condition_type, &seis_license_cond.Condition_value, &seis_license_cond.Condition_value_uom, &seis_license_cond.Contact_ba_id, &seis_license_cond.Date_format_desc, &seis_license_cond.Description, &seis_license_cond.Due_condition, &seis_license_cond.Due_date, &seis_license_cond.Due_frequency, &seis_license_cond.Due_term, &seis_license_cond.Due_term_uom, &seis_license_cond.Effective_date, &seis_license_cond.Exempt_ind, &seis_license_cond.Expiry_date, &seis_license_cond.Fulfilled_by_ba_id, &seis_license_cond.Fulfilled_date, &seis_license_cond.Fulfilled_ind, &seis_license_cond.Ppdm_guid, &seis_license_cond.Remark, &seis_license_cond.Restriction_id, &seis_license_cond.Restriction_version, &seis_license_cond.Source, &seis_license_cond.Row_changed_by, &seis_license_cond.Row_changed_date, &seis_license_cond.Row_created_by, &seis_license_cond.Row_created_date, &seis_license_cond.Row_effective_date, &seis_license_cond.Row_expiry_date, &seis_license_cond.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_license_cond to result
		result = append(result, seis_license_cond)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisLicenseCond(c *fiber.Ctx) error {
	var seis_license_cond dto.Seis_license_cond

	setDefaults(&seis_license_cond)

	if err := json.Unmarshal(c.Body(), &seis_license_cond); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_license_cond values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36)")
	if err != nil {
		return err
	}
	seis_license_cond.Row_created_date = formatDate(seis_license_cond.Row_created_date)
	seis_license_cond.Row_changed_date = formatDate(seis_license_cond.Row_changed_date)
	seis_license_cond.Date_format_desc = formatDateString(seis_license_cond.Date_format_desc)
	seis_license_cond.Due_date = formatDateString(seis_license_cond.Due_date)
	seis_license_cond.Effective_date = formatDateString(seis_license_cond.Effective_date)
	seis_license_cond.Expiry_date = formatDateString(seis_license_cond.Expiry_date)
	seis_license_cond.Fulfilled_date = formatDateString(seis_license_cond.Fulfilled_date)
	seis_license_cond.Row_effective_date = formatDateString(seis_license_cond.Row_effective_date)
	seis_license_cond.Row_expiry_date = formatDateString(seis_license_cond.Row_expiry_date)






	rows, err := stmt.Exec(seis_license_cond.Seis_set_subtype, seis_license_cond.Seis_set_id, seis_license_cond.License_id, seis_license_cond.Condition_id, seis_license_cond.Active_ind, seis_license_cond.Condition_code, seis_license_cond.Condition_desc, seis_license_cond.Condition_type, seis_license_cond.Condition_value, seis_license_cond.Condition_value_uom, seis_license_cond.Contact_ba_id, seis_license_cond.Date_format_desc, seis_license_cond.Description, seis_license_cond.Due_condition, seis_license_cond.Due_date, seis_license_cond.Due_frequency, seis_license_cond.Due_term, seis_license_cond.Due_term_uom, seis_license_cond.Effective_date, seis_license_cond.Exempt_ind, seis_license_cond.Expiry_date, seis_license_cond.Fulfilled_by_ba_id, seis_license_cond.Fulfilled_date, seis_license_cond.Fulfilled_ind, seis_license_cond.Ppdm_guid, seis_license_cond.Remark, seis_license_cond.Restriction_id, seis_license_cond.Restriction_version, seis_license_cond.Source, seis_license_cond.Row_changed_by, seis_license_cond.Row_changed_date, seis_license_cond.Row_created_by, seis_license_cond.Row_created_date, seis_license_cond.Row_effective_date, seis_license_cond.Row_expiry_date, seis_license_cond.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisLicenseCond(c *fiber.Ctx) error {
	var seis_license_cond dto.Seis_license_cond

	setDefaults(&seis_license_cond)

	if err := json.Unmarshal(c.Body(), &seis_license_cond); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_license_cond.Seis_set_subtype = id

    if seis_license_cond.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_license_cond set seis_set_id = :1, license_id = :2, condition_id = :3, active_ind = :4, condition_code = :5, condition_desc = :6, condition_type = :7, condition_value = :8, condition_value_uom = :9, contact_ba_id = :10, date_format_desc = :11, description = :12, due_condition = :13, due_date = :14, due_frequency = :15, due_term = :16, due_term_uom = :17, effective_date = :18, exempt_ind = :19, expiry_date = :20, fulfilled_by_ba_id = :21, fulfilled_date = :22, fulfilled_ind = :23, ppdm_guid = :24, remark = :25, restriction_id = :26, restriction_version = :27, source = :28, row_changed_by = :29, row_changed_date = :30, row_created_by = :31, row_effective_date = :32, row_expiry_date = :33, row_quality = :34 where seis_set_subtype = :36")
	if err != nil {
		return err
	}

	seis_license_cond.Row_changed_date = formatDate(seis_license_cond.Row_changed_date)
	seis_license_cond.Date_format_desc = formatDateString(seis_license_cond.Date_format_desc)
	seis_license_cond.Due_date = formatDateString(seis_license_cond.Due_date)
	seis_license_cond.Effective_date = formatDateString(seis_license_cond.Effective_date)
	seis_license_cond.Expiry_date = formatDateString(seis_license_cond.Expiry_date)
	seis_license_cond.Fulfilled_date = formatDateString(seis_license_cond.Fulfilled_date)
	seis_license_cond.Row_effective_date = formatDateString(seis_license_cond.Row_effective_date)
	seis_license_cond.Row_expiry_date = formatDateString(seis_license_cond.Row_expiry_date)






	rows, err := stmt.Exec(seis_license_cond.Seis_set_id, seis_license_cond.License_id, seis_license_cond.Condition_id, seis_license_cond.Active_ind, seis_license_cond.Condition_code, seis_license_cond.Condition_desc, seis_license_cond.Condition_type, seis_license_cond.Condition_value, seis_license_cond.Condition_value_uom, seis_license_cond.Contact_ba_id, seis_license_cond.Date_format_desc, seis_license_cond.Description, seis_license_cond.Due_condition, seis_license_cond.Due_date, seis_license_cond.Due_frequency, seis_license_cond.Due_term, seis_license_cond.Due_term_uom, seis_license_cond.Effective_date, seis_license_cond.Exempt_ind, seis_license_cond.Expiry_date, seis_license_cond.Fulfilled_by_ba_id, seis_license_cond.Fulfilled_date, seis_license_cond.Fulfilled_ind, seis_license_cond.Ppdm_guid, seis_license_cond.Remark, seis_license_cond.Restriction_id, seis_license_cond.Restriction_version, seis_license_cond.Source, seis_license_cond.Row_changed_by, seis_license_cond.Row_changed_date, seis_license_cond.Row_created_by, seis_license_cond.Row_effective_date, seis_license_cond.Row_expiry_date, seis_license_cond.Row_quality, seis_license_cond.Seis_set_subtype)
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

func PatchSeisLicenseCond(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_license_cond set "
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

func DeleteSeisLicenseCond(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_license_cond dto.Seis_license_cond
	seis_license_cond.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_license_cond where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_license_cond.Seis_set_subtype)
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


