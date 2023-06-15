package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetBaLicenseViolation(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ba_license_violation")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ba_license_violation

	for rows.Next() {
		var ba_license_violation dto.Ba_license_violation
		if err := rows.Scan(&ba_license_violation.Licensee_ba_id, &ba_license_violation.License_id, &ba_license_violation.Violation_id, &ba_license_violation.Active_ind, &ba_license_violation.Condition_id, &ba_license_violation.Effective_date, &ba_license_violation.Expiry_date, &ba_license_violation.Ppdm_guid, &ba_license_violation.Remark, &ba_license_violation.Resolve_date, &ba_license_violation.Resolve_desc, &ba_license_violation.Resolve_type, &ba_license_violation.Source, &ba_license_violation.Violation_date, &ba_license_violation.Violation_desc, &ba_license_violation.Violation_type, &ba_license_violation.Row_changed_by, &ba_license_violation.Row_changed_date, &ba_license_violation.Row_created_by, &ba_license_violation.Row_created_date, &ba_license_violation.Row_effective_date, &ba_license_violation.Row_expiry_date, &ba_license_violation.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ba_license_violation to result
		result = append(result, ba_license_violation)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetBaLicenseViolation(c *fiber.Ctx) error {
	var ba_license_violation dto.Ba_license_violation

	setDefaults(&ba_license_violation)

	if err := json.Unmarshal(c.Body(), &ba_license_violation); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ba_license_violation values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	ba_license_violation.Row_created_date = formatDate(ba_license_violation.Row_created_date)
	ba_license_violation.Row_changed_date = formatDate(ba_license_violation.Row_changed_date)
	ba_license_violation.Effective_date = formatDateString(ba_license_violation.Effective_date)
	ba_license_violation.Expiry_date = formatDateString(ba_license_violation.Expiry_date)
	ba_license_violation.Resolve_date = formatDateString(ba_license_violation.Resolve_date)
	ba_license_violation.Violation_date = formatDateString(ba_license_violation.Violation_date)
	ba_license_violation.Row_effective_date = formatDateString(ba_license_violation.Row_effective_date)
	ba_license_violation.Row_expiry_date = formatDateString(ba_license_violation.Row_expiry_date)






	rows, err := stmt.Exec(ba_license_violation.Licensee_ba_id, ba_license_violation.License_id, ba_license_violation.Violation_id, ba_license_violation.Active_ind, ba_license_violation.Condition_id, ba_license_violation.Effective_date, ba_license_violation.Expiry_date, ba_license_violation.Ppdm_guid, ba_license_violation.Remark, ba_license_violation.Resolve_date, ba_license_violation.Resolve_desc, ba_license_violation.Resolve_type, ba_license_violation.Source, ba_license_violation.Violation_date, ba_license_violation.Violation_desc, ba_license_violation.Violation_type, ba_license_violation.Row_changed_by, ba_license_violation.Row_changed_date, ba_license_violation.Row_created_by, ba_license_violation.Row_created_date, ba_license_violation.Row_effective_date, ba_license_violation.Row_expiry_date, ba_license_violation.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateBaLicenseViolation(c *fiber.Ctx) error {
	var ba_license_violation dto.Ba_license_violation

	setDefaults(&ba_license_violation)

	if err := json.Unmarshal(c.Body(), &ba_license_violation); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ba_license_violation.Licensee_ba_id = id

    if ba_license_violation.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ba_license_violation set license_id = :1, violation_id = :2, active_ind = :3, condition_id = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, resolve_date = :9, resolve_desc = :10, resolve_type = :11, source = :12, violation_date = :13, violation_desc = :14, violation_type = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where licensee_ba_id = :23")
	if err != nil {
		return err
	}

	ba_license_violation.Row_changed_date = formatDate(ba_license_violation.Row_changed_date)
	ba_license_violation.Effective_date = formatDateString(ba_license_violation.Effective_date)
	ba_license_violation.Expiry_date = formatDateString(ba_license_violation.Expiry_date)
	ba_license_violation.Resolve_date = formatDateString(ba_license_violation.Resolve_date)
	ba_license_violation.Violation_date = formatDateString(ba_license_violation.Violation_date)
	ba_license_violation.Row_effective_date = formatDateString(ba_license_violation.Row_effective_date)
	ba_license_violation.Row_expiry_date = formatDateString(ba_license_violation.Row_expiry_date)






	rows, err := stmt.Exec(ba_license_violation.License_id, ba_license_violation.Violation_id, ba_license_violation.Active_ind, ba_license_violation.Condition_id, ba_license_violation.Effective_date, ba_license_violation.Expiry_date, ba_license_violation.Ppdm_guid, ba_license_violation.Remark, ba_license_violation.Resolve_date, ba_license_violation.Resolve_desc, ba_license_violation.Resolve_type, ba_license_violation.Source, ba_license_violation.Violation_date, ba_license_violation.Violation_desc, ba_license_violation.Violation_type, ba_license_violation.Row_changed_by, ba_license_violation.Row_changed_date, ba_license_violation.Row_created_by, ba_license_violation.Row_effective_date, ba_license_violation.Row_expiry_date, ba_license_violation.Row_quality, ba_license_violation.Licensee_ba_id)
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

func PatchBaLicenseViolation(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ba_license_violation set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "resolve_date" || key == "violation_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteBaLicenseViolation(c *fiber.Ctx) error {
	id := c.Params("id")
	var ba_license_violation dto.Ba_license_violation
	ba_license_violation.Licensee_ba_id = id

	stmt, err := db.Prepare("delete from ba_license_violation where licensee_ba_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ba_license_violation.Licensee_ba_id)
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


