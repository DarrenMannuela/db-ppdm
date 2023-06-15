package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFacilityLicViolation(c *fiber.Ctx) error {
	rows, err := db.Query("select * from facility_lic_violation")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Facility_lic_violation

	for rows.Next() {
		var facility_lic_violation dto.Facility_lic_violation
		if err := rows.Scan(&facility_lic_violation.License_id, &facility_lic_violation.Facility_type, &facility_lic_violation.Facility_id, &facility_lic_violation.Violation_id, &facility_lic_violation.Active_ind, &facility_lic_violation.Condition_id, &facility_lic_violation.Effective_date, &facility_lic_violation.Expiry_date, &facility_lic_violation.Ppdm_guid, &facility_lic_violation.Remark, &facility_lic_violation.Resolve_date, &facility_lic_violation.Resolve_desc, &facility_lic_violation.Resolve_type, &facility_lic_violation.Source, &facility_lic_violation.Violation_date, &facility_lic_violation.Violation_desc, &facility_lic_violation.Violation_type, &facility_lic_violation.Row_changed_by, &facility_lic_violation.Row_changed_date, &facility_lic_violation.Row_created_by, &facility_lic_violation.Row_created_date, &facility_lic_violation.Row_effective_date, &facility_lic_violation.Row_expiry_date, &facility_lic_violation.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append facility_lic_violation to result
		result = append(result, facility_lic_violation)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFacilityLicViolation(c *fiber.Ctx) error {
	var facility_lic_violation dto.Facility_lic_violation

	setDefaults(&facility_lic_violation)

	if err := json.Unmarshal(c.Body(), &facility_lic_violation); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into facility_lic_violation values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	facility_lic_violation.Row_created_date = formatDate(facility_lic_violation.Row_created_date)
	facility_lic_violation.Row_changed_date = formatDate(facility_lic_violation.Row_changed_date)
	facility_lic_violation.Effective_date = formatDateString(facility_lic_violation.Effective_date)
	facility_lic_violation.Expiry_date = formatDateString(facility_lic_violation.Expiry_date)
	facility_lic_violation.Resolve_date = formatDateString(facility_lic_violation.Resolve_date)
	facility_lic_violation.Violation_date = formatDateString(facility_lic_violation.Violation_date)
	facility_lic_violation.Row_effective_date = formatDateString(facility_lic_violation.Row_effective_date)
	facility_lic_violation.Row_expiry_date = formatDateString(facility_lic_violation.Row_expiry_date)






	rows, err := stmt.Exec(facility_lic_violation.License_id, facility_lic_violation.Facility_type, facility_lic_violation.Facility_id, facility_lic_violation.Violation_id, facility_lic_violation.Active_ind, facility_lic_violation.Condition_id, facility_lic_violation.Effective_date, facility_lic_violation.Expiry_date, facility_lic_violation.Ppdm_guid, facility_lic_violation.Remark, facility_lic_violation.Resolve_date, facility_lic_violation.Resolve_desc, facility_lic_violation.Resolve_type, facility_lic_violation.Source, facility_lic_violation.Violation_date, facility_lic_violation.Violation_desc, facility_lic_violation.Violation_type, facility_lic_violation.Row_changed_by, facility_lic_violation.Row_changed_date, facility_lic_violation.Row_created_by, facility_lic_violation.Row_created_date, facility_lic_violation.Row_effective_date, facility_lic_violation.Row_expiry_date, facility_lic_violation.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFacilityLicViolation(c *fiber.Ctx) error {
	var facility_lic_violation dto.Facility_lic_violation

	setDefaults(&facility_lic_violation)

	if err := json.Unmarshal(c.Body(), &facility_lic_violation); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	facility_lic_violation.License_id = id

    if facility_lic_violation.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update facility_lic_violation set facility_type = :1, facility_id = :2, violation_id = :3, active_ind = :4, condition_id = :5, effective_date = :6, expiry_date = :7, ppdm_guid = :8, remark = :9, resolve_date = :10, resolve_desc = :11, resolve_type = :12, source = :13, violation_date = :14, violation_desc = :15, violation_type = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where license_id = :24")
	if err != nil {
		return err
	}

	facility_lic_violation.Row_changed_date = formatDate(facility_lic_violation.Row_changed_date)
	facility_lic_violation.Effective_date = formatDateString(facility_lic_violation.Effective_date)
	facility_lic_violation.Expiry_date = formatDateString(facility_lic_violation.Expiry_date)
	facility_lic_violation.Resolve_date = formatDateString(facility_lic_violation.Resolve_date)
	facility_lic_violation.Violation_date = formatDateString(facility_lic_violation.Violation_date)
	facility_lic_violation.Row_effective_date = formatDateString(facility_lic_violation.Row_effective_date)
	facility_lic_violation.Row_expiry_date = formatDateString(facility_lic_violation.Row_expiry_date)






	rows, err := stmt.Exec(facility_lic_violation.Facility_type, facility_lic_violation.Facility_id, facility_lic_violation.Violation_id, facility_lic_violation.Active_ind, facility_lic_violation.Condition_id, facility_lic_violation.Effective_date, facility_lic_violation.Expiry_date, facility_lic_violation.Ppdm_guid, facility_lic_violation.Remark, facility_lic_violation.Resolve_date, facility_lic_violation.Resolve_desc, facility_lic_violation.Resolve_type, facility_lic_violation.Source, facility_lic_violation.Violation_date, facility_lic_violation.Violation_desc, facility_lic_violation.Violation_type, facility_lic_violation.Row_changed_by, facility_lic_violation.Row_changed_date, facility_lic_violation.Row_created_by, facility_lic_violation.Row_effective_date, facility_lic_violation.Row_expiry_date, facility_lic_violation.Row_quality, facility_lic_violation.License_id)
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

func PatchFacilityLicViolation(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update facility_lic_violation set "
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
	query += " where license_id = :id"

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

func DeleteFacilityLicViolation(c *fiber.Ctx) error {
	id := c.Params("id")
	var facility_lic_violation dto.Facility_lic_violation
	facility_lic_violation.License_id = id

	stmt, err := db.Prepare("delete from facility_lic_violation where license_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(facility_lic_violation.License_id)
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


