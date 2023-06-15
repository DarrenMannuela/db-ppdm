package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellActivityCause(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_activity_cause")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_activity_cause

	for rows.Next() {
		var well_activity_cause dto.Well_activity_cause
		if err := rows.Scan(&well_activity_cause.Uwi, &well_activity_cause.Source, &well_activity_cause.Activity_obs_no, &well_activity_cause.Cause_type, &well_activity_cause.Active_ind, &well_activity_cause.Description, &well_activity_cause.Effective_date, &well_activity_cause.Equipment_id, &well_activity_cause.Expiry_date, &well_activity_cause.Period_obs_no, &well_activity_cause.Ppdm_guid, &well_activity_cause.Remark, &well_activity_cause.Row_changed_by, &well_activity_cause.Row_changed_date, &well_activity_cause.Row_created_by, &well_activity_cause.Row_created_date, &well_activity_cause.Row_effective_date, &well_activity_cause.Row_expiry_date, &well_activity_cause.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_activity_cause to result
		result = append(result, well_activity_cause)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellActivityCause(c *fiber.Ctx) error {
	var well_activity_cause dto.Well_activity_cause

	setDefaults(&well_activity_cause)

	if err := json.Unmarshal(c.Body(), &well_activity_cause); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_activity_cause values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	well_activity_cause.Row_created_date = formatDate(well_activity_cause.Row_created_date)
	well_activity_cause.Row_changed_date = formatDate(well_activity_cause.Row_changed_date)
	well_activity_cause.Effective_date = formatDateString(well_activity_cause.Effective_date)
	well_activity_cause.Expiry_date = formatDateString(well_activity_cause.Expiry_date)
	well_activity_cause.Row_effective_date = formatDateString(well_activity_cause.Row_effective_date)
	well_activity_cause.Row_expiry_date = formatDateString(well_activity_cause.Row_expiry_date)






	rows, err := stmt.Exec(well_activity_cause.Uwi, well_activity_cause.Source, well_activity_cause.Activity_obs_no, well_activity_cause.Cause_type, well_activity_cause.Active_ind, well_activity_cause.Description, well_activity_cause.Effective_date, well_activity_cause.Equipment_id, well_activity_cause.Expiry_date, well_activity_cause.Period_obs_no, well_activity_cause.Ppdm_guid, well_activity_cause.Remark, well_activity_cause.Row_changed_by, well_activity_cause.Row_changed_date, well_activity_cause.Row_created_by, well_activity_cause.Row_created_date, well_activity_cause.Row_effective_date, well_activity_cause.Row_expiry_date, well_activity_cause.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellActivityCause(c *fiber.Ctx) error {
	var well_activity_cause dto.Well_activity_cause

	setDefaults(&well_activity_cause)

	if err := json.Unmarshal(c.Body(), &well_activity_cause); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_activity_cause.Uwi = id

    if well_activity_cause.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_activity_cause set source = :1, activity_obs_no = :2, cause_type = :3, active_ind = :4, description = :5, effective_date = :6, equipment_id = :7, expiry_date = :8, period_obs_no = :9, ppdm_guid = :10, remark = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where uwi = :19")
	if err != nil {
		return err
	}

	well_activity_cause.Row_changed_date = formatDate(well_activity_cause.Row_changed_date)
	well_activity_cause.Effective_date = formatDateString(well_activity_cause.Effective_date)
	well_activity_cause.Expiry_date = formatDateString(well_activity_cause.Expiry_date)
	well_activity_cause.Row_effective_date = formatDateString(well_activity_cause.Row_effective_date)
	well_activity_cause.Row_expiry_date = formatDateString(well_activity_cause.Row_expiry_date)






	rows, err := stmt.Exec(well_activity_cause.Source, well_activity_cause.Activity_obs_no, well_activity_cause.Cause_type, well_activity_cause.Active_ind, well_activity_cause.Description, well_activity_cause.Effective_date, well_activity_cause.Equipment_id, well_activity_cause.Expiry_date, well_activity_cause.Period_obs_no, well_activity_cause.Ppdm_guid, well_activity_cause.Remark, well_activity_cause.Row_changed_by, well_activity_cause.Row_changed_date, well_activity_cause.Row_created_by, well_activity_cause.Row_effective_date, well_activity_cause.Row_expiry_date, well_activity_cause.Row_quality, well_activity_cause.Uwi)
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

func PatchWellActivityCause(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_activity_cause set "
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

func DeleteWellActivityCause(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_activity_cause dto.Well_activity_cause
	well_activity_cause.Uwi = id

	stmt, err := db.Prepare("delete from well_activity_cause where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_activity_cause.Uwi)
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


