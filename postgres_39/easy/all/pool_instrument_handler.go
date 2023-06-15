package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPoolInstrument(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pool_instrument")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pool_instrument

	for rows.Next() {
		var pool_instrument dto.Pool_instrument
		if err := rows.Scan(&pool_instrument.Instrument_id, &pool_instrument.Pool_id, &pool_instrument.Active_ind, &pool_instrument.Effective_date, &pool_instrument.Expiry_date, &pool_instrument.Ppdm_guid, &pool_instrument.Remark, &pool_instrument.Source, &pool_instrument.Row_changed_by, &pool_instrument.Row_changed_date, &pool_instrument.Row_created_by, &pool_instrument.Row_created_date, &pool_instrument.Row_effective_date, &pool_instrument.Row_expiry_date, &pool_instrument.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pool_instrument to result
		result = append(result, pool_instrument)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPoolInstrument(c *fiber.Ctx) error {
	var pool_instrument dto.Pool_instrument

	setDefaults(&pool_instrument)

	if err := json.Unmarshal(c.Body(), &pool_instrument); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pool_instrument values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15)")
	if err != nil {
		return err
	}
	pool_instrument.Row_created_date = formatDate(pool_instrument.Row_created_date)
	pool_instrument.Row_changed_date = formatDate(pool_instrument.Row_changed_date)
	pool_instrument.Effective_date = formatDateString(pool_instrument.Effective_date)
	pool_instrument.Expiry_date = formatDateString(pool_instrument.Expiry_date)
	pool_instrument.Row_effective_date = formatDateString(pool_instrument.Row_effective_date)
	pool_instrument.Row_expiry_date = formatDateString(pool_instrument.Row_expiry_date)






	rows, err := stmt.Exec(pool_instrument.Instrument_id, pool_instrument.Pool_id, pool_instrument.Active_ind, pool_instrument.Effective_date, pool_instrument.Expiry_date, pool_instrument.Ppdm_guid, pool_instrument.Remark, pool_instrument.Source, pool_instrument.Row_changed_by, pool_instrument.Row_changed_date, pool_instrument.Row_created_by, pool_instrument.Row_created_date, pool_instrument.Row_effective_date, pool_instrument.Row_expiry_date, pool_instrument.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePoolInstrument(c *fiber.Ctx) error {
	var pool_instrument dto.Pool_instrument

	setDefaults(&pool_instrument)

	if err := json.Unmarshal(c.Body(), &pool_instrument); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pool_instrument.Instrument_id = id

    if pool_instrument.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pool_instrument set pool_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, remark = :6, source = :7, row_changed_by = :8, row_changed_date = :9, row_created_by = :10, row_effective_date = :11, row_expiry_date = :12, row_quality = :13 where instrument_id = :15")
	if err != nil {
		return err
	}

	pool_instrument.Row_changed_date = formatDate(pool_instrument.Row_changed_date)
	pool_instrument.Effective_date = formatDateString(pool_instrument.Effective_date)
	pool_instrument.Expiry_date = formatDateString(pool_instrument.Expiry_date)
	pool_instrument.Row_effective_date = formatDateString(pool_instrument.Row_effective_date)
	pool_instrument.Row_expiry_date = formatDateString(pool_instrument.Row_expiry_date)






	rows, err := stmt.Exec(pool_instrument.Pool_id, pool_instrument.Active_ind, pool_instrument.Effective_date, pool_instrument.Expiry_date, pool_instrument.Ppdm_guid, pool_instrument.Remark, pool_instrument.Source, pool_instrument.Row_changed_by, pool_instrument.Row_changed_date, pool_instrument.Row_created_by, pool_instrument.Row_effective_date, pool_instrument.Row_expiry_date, pool_instrument.Row_quality, pool_instrument.Instrument_id)
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

func PatchPoolInstrument(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pool_instrument set "
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
	query += " where instrument_id = :id"

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

func DeletePoolInstrument(c *fiber.Ctx) error {
	id := c.Params("id")
	var pool_instrument dto.Pool_instrument
	pool_instrument.Instrument_id = id

	stmt, err := db.Prepare("delete from pool_instrument where instrument_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pool_instrument.Instrument_id)
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


