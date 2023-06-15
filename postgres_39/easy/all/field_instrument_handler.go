package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFieldInstrument(c *fiber.Ctx) error {
	rows, err := db.Query("select * from field_instrument")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Field_instrument

	for rows.Next() {
		var field_instrument dto.Field_instrument
		if err := rows.Scan(&field_instrument.Field_id, &field_instrument.Instrument_id, &field_instrument.Active_ind, &field_instrument.Effective_date, &field_instrument.Expiry_date, &field_instrument.Ppdm_guid, &field_instrument.Remark, &field_instrument.Source, &field_instrument.Row_changed_by, &field_instrument.Row_changed_date, &field_instrument.Row_created_by, &field_instrument.Row_created_date, &field_instrument.Row_effective_date, &field_instrument.Row_expiry_date, &field_instrument.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append field_instrument to result
		result = append(result, field_instrument)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFieldInstrument(c *fiber.Ctx) error {
	var field_instrument dto.Field_instrument

	setDefaults(&field_instrument)

	if err := json.Unmarshal(c.Body(), &field_instrument); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into field_instrument values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15)")
	if err != nil {
		return err
	}
	field_instrument.Row_created_date = formatDate(field_instrument.Row_created_date)
	field_instrument.Row_changed_date = formatDate(field_instrument.Row_changed_date)
	field_instrument.Effective_date = formatDateString(field_instrument.Effective_date)
	field_instrument.Expiry_date = formatDateString(field_instrument.Expiry_date)
	field_instrument.Row_effective_date = formatDateString(field_instrument.Row_effective_date)
	field_instrument.Row_expiry_date = formatDateString(field_instrument.Row_expiry_date)






	rows, err := stmt.Exec(field_instrument.Field_id, field_instrument.Instrument_id, field_instrument.Active_ind, field_instrument.Effective_date, field_instrument.Expiry_date, field_instrument.Ppdm_guid, field_instrument.Remark, field_instrument.Source, field_instrument.Row_changed_by, field_instrument.Row_changed_date, field_instrument.Row_created_by, field_instrument.Row_created_date, field_instrument.Row_effective_date, field_instrument.Row_expiry_date, field_instrument.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFieldInstrument(c *fiber.Ctx) error {
	var field_instrument dto.Field_instrument

	setDefaults(&field_instrument)

	if err := json.Unmarshal(c.Body(), &field_instrument); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	field_instrument.Field_id = id

    if field_instrument.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update field_instrument set instrument_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, remark = :6, source = :7, row_changed_by = :8, row_changed_date = :9, row_created_by = :10, row_effective_date = :11, row_expiry_date = :12, row_quality = :13 where field_id = :15")
	if err != nil {
		return err
	}

	field_instrument.Row_changed_date = formatDate(field_instrument.Row_changed_date)
	field_instrument.Effective_date = formatDateString(field_instrument.Effective_date)
	field_instrument.Expiry_date = formatDateString(field_instrument.Expiry_date)
	field_instrument.Row_effective_date = formatDateString(field_instrument.Row_effective_date)
	field_instrument.Row_expiry_date = formatDateString(field_instrument.Row_expiry_date)






	rows, err := stmt.Exec(field_instrument.Instrument_id, field_instrument.Active_ind, field_instrument.Effective_date, field_instrument.Expiry_date, field_instrument.Ppdm_guid, field_instrument.Remark, field_instrument.Source, field_instrument.Row_changed_by, field_instrument.Row_changed_date, field_instrument.Row_created_by, field_instrument.Row_effective_date, field_instrument.Row_expiry_date, field_instrument.Row_quality, field_instrument.Field_id)
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

func PatchFieldInstrument(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update field_instrument set "
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
	query += " where field_id = :id"

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

func DeleteFieldInstrument(c *fiber.Ctx) error {
	id := c.Params("id")
	var field_instrument dto.Field_instrument
	field_instrument.Field_id = id

	stmt, err := db.Prepare("delete from field_instrument where field_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(field_instrument.Field_id)
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


