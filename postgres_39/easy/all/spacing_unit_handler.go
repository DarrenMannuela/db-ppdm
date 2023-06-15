package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpacingUnit(c *fiber.Ctx) error {
	rows, err := db.Query("select * from spacing_unit")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Spacing_unit

	for rows.Next() {
		var spacing_unit dto.Spacing_unit
		if err := rows.Scan(&spacing_unit.Spacing_unit_id, &spacing_unit.Active_ind, &spacing_unit.Effective_date, &spacing_unit.Expiry_date, &spacing_unit.Gross_size, &spacing_unit.Gross_size_ouom, &spacing_unit.Ppdm_guid, &spacing_unit.Remark, &spacing_unit.Source, &spacing_unit.Spacing_unit_desc, &spacing_unit.Spacing_unit_name, &spacing_unit.Spacing_unit_type, &spacing_unit.Row_changed_by, &spacing_unit.Row_changed_date, &spacing_unit.Row_created_by, &spacing_unit.Row_created_date, &spacing_unit.Row_effective_date, &spacing_unit.Row_expiry_date, &spacing_unit.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append spacing_unit to result
		result = append(result, spacing_unit)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpacingUnit(c *fiber.Ctx) error {
	var spacing_unit dto.Spacing_unit

	setDefaults(&spacing_unit)

	if err := json.Unmarshal(c.Body(), &spacing_unit); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into spacing_unit values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	spacing_unit.Row_created_date = formatDate(spacing_unit.Row_created_date)
	spacing_unit.Row_changed_date = formatDate(spacing_unit.Row_changed_date)
	spacing_unit.Effective_date = formatDateString(spacing_unit.Effective_date)
	spacing_unit.Expiry_date = formatDateString(spacing_unit.Expiry_date)
	spacing_unit.Row_effective_date = formatDateString(spacing_unit.Row_effective_date)
	spacing_unit.Row_expiry_date = formatDateString(spacing_unit.Row_expiry_date)






	rows, err := stmt.Exec(spacing_unit.Spacing_unit_id, spacing_unit.Active_ind, spacing_unit.Effective_date, spacing_unit.Expiry_date, spacing_unit.Gross_size, spacing_unit.Gross_size_ouom, spacing_unit.Ppdm_guid, spacing_unit.Remark, spacing_unit.Source, spacing_unit.Spacing_unit_desc, spacing_unit.Spacing_unit_name, spacing_unit.Spacing_unit_type, spacing_unit.Row_changed_by, spacing_unit.Row_changed_date, spacing_unit.Row_created_by, spacing_unit.Row_created_date, spacing_unit.Row_effective_date, spacing_unit.Row_expiry_date, spacing_unit.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpacingUnit(c *fiber.Ctx) error {
	var spacing_unit dto.Spacing_unit

	setDefaults(&spacing_unit)

	if err := json.Unmarshal(c.Body(), &spacing_unit); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	spacing_unit.Spacing_unit_id = id

    if spacing_unit.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update spacing_unit set active_ind = :1, effective_date = :2, expiry_date = :3, gross_size = :4, gross_size_ouom = :5, ppdm_guid = :6, remark = :7, source = :8, spacing_unit_desc = :9, spacing_unit_name = :10, spacing_unit_type = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where spacing_unit_id = :19")
	if err != nil {
		return err
	}

	spacing_unit.Row_changed_date = formatDate(spacing_unit.Row_changed_date)
	spacing_unit.Effective_date = formatDateString(spacing_unit.Effective_date)
	spacing_unit.Expiry_date = formatDateString(spacing_unit.Expiry_date)
	spacing_unit.Row_effective_date = formatDateString(spacing_unit.Row_effective_date)
	spacing_unit.Row_expiry_date = formatDateString(spacing_unit.Row_expiry_date)






	rows, err := stmt.Exec(spacing_unit.Active_ind, spacing_unit.Effective_date, spacing_unit.Expiry_date, spacing_unit.Gross_size, spacing_unit.Gross_size_ouom, spacing_unit.Ppdm_guid, spacing_unit.Remark, spacing_unit.Source, spacing_unit.Spacing_unit_desc, spacing_unit.Spacing_unit_name, spacing_unit.Spacing_unit_type, spacing_unit.Row_changed_by, spacing_unit.Row_changed_date, spacing_unit.Row_created_by, spacing_unit.Row_effective_date, spacing_unit.Row_expiry_date, spacing_unit.Row_quality, spacing_unit.Spacing_unit_id)
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

func PatchSpacingUnit(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update spacing_unit set "
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
	query += " where spacing_unit_id = :id"

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

func DeleteSpacingUnit(c *fiber.Ctx) error {
	id := c.Params("id")
	var spacing_unit dto.Spacing_unit
	spacing_unit.Spacing_unit_id = id

	stmt, err := db.Prepare("delete from spacing_unit where spacing_unit_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(spacing_unit.Spacing_unit_id)
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


