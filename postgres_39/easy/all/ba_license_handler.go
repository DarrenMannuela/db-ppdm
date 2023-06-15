package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetBaLicense(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ba_license")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ba_license

	for rows.Next() {
		var ba_license dto.Ba_license
		if err := rows.Scan(&ba_license.Licensee_ba_id, &ba_license.License_id, &ba_license.Active_ind, &ba_license.Application_id, &ba_license.Description, &ba_license.Effective_date, &ba_license.Expiry_date, &ba_license.Extended_date, &ba_license.Fees_paid_ind, &ba_license.Granted_by_ba_id, &ba_license.Granted_by_contact_id, &ba_license.Granted_date, &ba_license.Granted_to_ba_id, &ba_license.Granted_to_contact_id, &ba_license.License_num, &ba_license.License_term, &ba_license.License_term_uom, &ba_license.License_type, &ba_license.Ppdm_guid, &ba_license.Rate_schedule_id, &ba_license.Rel_licensee_ba_id, &ba_license.Rel_license_id, &ba_license.Remark, &ba_license.Secondary_term, &ba_license.Secondary_term_uom, &ba_license.Source, &ba_license.Violation_ind, &ba_license.Row_changed_by, &ba_license.Row_changed_date, &ba_license.Row_created_by, &ba_license.Row_created_date, &ba_license.Row_effective_date, &ba_license.Row_expiry_date, &ba_license.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ba_license to result
		result = append(result, ba_license)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetBaLicense(c *fiber.Ctx) error {
	var ba_license dto.Ba_license

	setDefaults(&ba_license)

	if err := json.Unmarshal(c.Body(), &ba_license); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ba_license values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34)")
	if err != nil {
		return err
	}
	ba_license.Row_created_date = formatDate(ba_license.Row_created_date)
	ba_license.Row_changed_date = formatDate(ba_license.Row_changed_date)
	ba_license.Effective_date = formatDateString(ba_license.Effective_date)
	ba_license.Expiry_date = formatDateString(ba_license.Expiry_date)
	ba_license.Extended_date = formatDateString(ba_license.Extended_date)
	ba_license.Granted_date = formatDateString(ba_license.Granted_date)
	ba_license.Row_effective_date = formatDateString(ba_license.Row_effective_date)
	ba_license.Row_expiry_date = formatDateString(ba_license.Row_expiry_date)






	rows, err := stmt.Exec(ba_license.Licensee_ba_id, ba_license.License_id, ba_license.Active_ind, ba_license.Application_id, ba_license.Description, ba_license.Effective_date, ba_license.Expiry_date, ba_license.Extended_date, ba_license.Fees_paid_ind, ba_license.Granted_by_ba_id, ba_license.Granted_by_contact_id, ba_license.Granted_date, ba_license.Granted_to_ba_id, ba_license.Granted_to_contact_id, ba_license.License_num, ba_license.License_term, ba_license.License_term_uom, ba_license.License_type, ba_license.Ppdm_guid, ba_license.Rate_schedule_id, ba_license.Rel_licensee_ba_id, ba_license.Rel_license_id, ba_license.Remark, ba_license.Secondary_term, ba_license.Secondary_term_uom, ba_license.Source, ba_license.Violation_ind, ba_license.Row_changed_by, ba_license.Row_changed_date, ba_license.Row_created_by, ba_license.Row_created_date, ba_license.Row_effective_date, ba_license.Row_expiry_date, ba_license.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateBaLicense(c *fiber.Ctx) error {
	var ba_license dto.Ba_license

	setDefaults(&ba_license)

	if err := json.Unmarshal(c.Body(), &ba_license); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ba_license.Licensee_ba_id = id

    if ba_license.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ba_license set license_id = :1, active_ind = :2, application_id = :3, description = :4, effective_date = :5, expiry_date = :6, extended_date = :7, fees_paid_ind = :8, granted_by_ba_id = :9, granted_by_contact_id = :10, granted_date = :11, granted_to_ba_id = :12, granted_to_contact_id = :13, license_num = :14, license_term = :15, license_term_uom = :16, license_type = :17, ppdm_guid = :18, rate_schedule_id = :19, rel_licensee_ba_id = :20, rel_license_id = :21, remark = :22, secondary_term = :23, secondary_term_uom = :24, source = :25, violation_ind = :26, row_changed_by = :27, row_changed_date = :28, row_created_by = :29, row_effective_date = :30, row_expiry_date = :31, row_quality = :32 where licensee_ba_id = :34")
	if err != nil {
		return err
	}

	ba_license.Row_changed_date = formatDate(ba_license.Row_changed_date)
	ba_license.Effective_date = formatDateString(ba_license.Effective_date)
	ba_license.Expiry_date = formatDateString(ba_license.Expiry_date)
	ba_license.Extended_date = formatDateString(ba_license.Extended_date)
	ba_license.Granted_date = formatDateString(ba_license.Granted_date)
	ba_license.Row_effective_date = formatDateString(ba_license.Row_effective_date)
	ba_license.Row_expiry_date = formatDateString(ba_license.Row_expiry_date)






	rows, err := stmt.Exec(ba_license.License_id, ba_license.Active_ind, ba_license.Application_id, ba_license.Description, ba_license.Effective_date, ba_license.Expiry_date, ba_license.Extended_date, ba_license.Fees_paid_ind, ba_license.Granted_by_ba_id, ba_license.Granted_by_contact_id, ba_license.Granted_date, ba_license.Granted_to_ba_id, ba_license.Granted_to_contact_id, ba_license.License_num, ba_license.License_term, ba_license.License_term_uom, ba_license.License_type, ba_license.Ppdm_guid, ba_license.Rate_schedule_id, ba_license.Rel_licensee_ba_id, ba_license.Rel_license_id, ba_license.Remark, ba_license.Secondary_term, ba_license.Secondary_term_uom, ba_license.Source, ba_license.Violation_ind, ba_license.Row_changed_by, ba_license.Row_changed_date, ba_license.Row_created_by, ba_license.Row_effective_date, ba_license.Row_expiry_date, ba_license.Row_quality, ba_license.Licensee_ba_id)
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

func PatchBaLicense(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ba_license set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "extended_date" || key == "granted_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteBaLicense(c *fiber.Ctx) error {
	id := c.Params("id")
	var ba_license dto.Ba_license
	ba_license.Licensee_ba_id = id

	stmt, err := db.Prepare("delete from ba_license where licensee_ba_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ba_license.Licensee_ba_id)
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


