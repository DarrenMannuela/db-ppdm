package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpacingUnitInst(c *fiber.Ctx) error {
	rows, err := db.Query("select * from spacing_unit_inst")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Spacing_unit_inst

	for rows.Next() {
		var spacing_unit_inst dto.Spacing_unit_inst
		if err := rows.Scan(&spacing_unit_inst.Spacing_unit_id, &spacing_unit_inst.Instrument_id, &spacing_unit_inst.Active_ind, &spacing_unit_inst.Effective_date, &spacing_unit_inst.Expiry_date, &spacing_unit_inst.Ppdm_guid, &spacing_unit_inst.Remark, &spacing_unit_inst.Source, &spacing_unit_inst.Row_changed_by, &spacing_unit_inst.Row_changed_date, &spacing_unit_inst.Row_created_by, &spacing_unit_inst.Row_created_date, &spacing_unit_inst.Row_effective_date, &spacing_unit_inst.Row_expiry_date, &spacing_unit_inst.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append spacing_unit_inst to result
		result = append(result, spacing_unit_inst)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpacingUnitInst(c *fiber.Ctx) error {
	var spacing_unit_inst dto.Spacing_unit_inst

	setDefaults(&spacing_unit_inst)

	if err := json.Unmarshal(c.Body(), &spacing_unit_inst); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into spacing_unit_inst values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15)")
	if err != nil {
		return err
	}
	spacing_unit_inst.Row_created_date = formatDate(spacing_unit_inst.Row_created_date)
	spacing_unit_inst.Row_changed_date = formatDate(spacing_unit_inst.Row_changed_date)
	spacing_unit_inst.Effective_date = formatDateString(spacing_unit_inst.Effective_date)
	spacing_unit_inst.Expiry_date = formatDateString(spacing_unit_inst.Expiry_date)
	spacing_unit_inst.Row_effective_date = formatDateString(spacing_unit_inst.Row_effective_date)
	spacing_unit_inst.Row_expiry_date = formatDateString(spacing_unit_inst.Row_expiry_date)






	rows, err := stmt.Exec(spacing_unit_inst.Spacing_unit_id, spacing_unit_inst.Instrument_id, spacing_unit_inst.Active_ind, spacing_unit_inst.Effective_date, spacing_unit_inst.Expiry_date, spacing_unit_inst.Ppdm_guid, spacing_unit_inst.Remark, spacing_unit_inst.Source, spacing_unit_inst.Row_changed_by, spacing_unit_inst.Row_changed_date, spacing_unit_inst.Row_created_by, spacing_unit_inst.Row_created_date, spacing_unit_inst.Row_effective_date, spacing_unit_inst.Row_expiry_date, spacing_unit_inst.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpacingUnitInst(c *fiber.Ctx) error {
	var spacing_unit_inst dto.Spacing_unit_inst

	setDefaults(&spacing_unit_inst)

	if err := json.Unmarshal(c.Body(), &spacing_unit_inst); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	spacing_unit_inst.Spacing_unit_id = id

    if spacing_unit_inst.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update spacing_unit_inst set instrument_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, remark = :6, source = :7, row_changed_by = :8, row_changed_date = :9, row_created_by = :10, row_effective_date = :11, row_expiry_date = :12, row_quality = :13 where spacing_unit_id = :15")
	if err != nil {
		return err
	}

	spacing_unit_inst.Row_changed_date = formatDate(spacing_unit_inst.Row_changed_date)
	spacing_unit_inst.Effective_date = formatDateString(spacing_unit_inst.Effective_date)
	spacing_unit_inst.Expiry_date = formatDateString(spacing_unit_inst.Expiry_date)
	spacing_unit_inst.Row_effective_date = formatDateString(spacing_unit_inst.Row_effective_date)
	spacing_unit_inst.Row_expiry_date = formatDateString(spacing_unit_inst.Row_expiry_date)






	rows, err := stmt.Exec(spacing_unit_inst.Instrument_id, spacing_unit_inst.Active_ind, spacing_unit_inst.Effective_date, spacing_unit_inst.Expiry_date, spacing_unit_inst.Ppdm_guid, spacing_unit_inst.Remark, spacing_unit_inst.Source, spacing_unit_inst.Row_changed_by, spacing_unit_inst.Row_changed_date, spacing_unit_inst.Row_created_by, spacing_unit_inst.Row_effective_date, spacing_unit_inst.Row_expiry_date, spacing_unit_inst.Row_quality, spacing_unit_inst.Spacing_unit_id)
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

func PatchSpacingUnitInst(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update spacing_unit_inst set "
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

func DeleteSpacingUnitInst(c *fiber.Ctx) error {
	id := c.Params("id")
	var spacing_unit_inst dto.Spacing_unit_inst
	spacing_unit_inst.Spacing_unit_id = id

	stmt, err := db.Prepare("delete from spacing_unit_inst where spacing_unit_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(spacing_unit_inst.Spacing_unit_id)
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


