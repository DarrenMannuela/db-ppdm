package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPaleoClimate(c *fiber.Ctx) error {
	rows, err := db.Query("select * from paleo_climate")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Paleo_climate

	for rows.Next() {
		var paleo_climate dto.Paleo_climate
		if err := rows.Scan(&paleo_climate.Climate_id, &paleo_climate.Active_ind, &paleo_climate.Climate_name, &paleo_climate.Climate_type, &paleo_climate.Effective_date, &paleo_climate.Expiry_date, &paleo_climate.Ppdm_guid, &paleo_climate.Remark, &paleo_climate.Source, &paleo_climate.Row_changed_by, &paleo_climate.Row_changed_date, &paleo_climate.Row_created_by, &paleo_climate.Row_created_date, &paleo_climate.Row_effective_date, &paleo_climate.Row_expiry_date, &paleo_climate.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append paleo_climate to result
		result = append(result, paleo_climate)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPaleoClimate(c *fiber.Ctx) error {
	var paleo_climate dto.Paleo_climate

	setDefaults(&paleo_climate)

	if err := json.Unmarshal(c.Body(), &paleo_climate); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into paleo_climate values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	paleo_climate.Row_created_date = formatDate(paleo_climate.Row_created_date)
	paleo_climate.Row_changed_date = formatDate(paleo_climate.Row_changed_date)
	paleo_climate.Effective_date = formatDateString(paleo_climate.Effective_date)
	paleo_climate.Expiry_date = formatDateString(paleo_climate.Expiry_date)
	paleo_climate.Row_effective_date = formatDateString(paleo_climate.Row_effective_date)
	paleo_climate.Row_expiry_date = formatDateString(paleo_climate.Row_expiry_date)






	rows, err := stmt.Exec(paleo_climate.Climate_id, paleo_climate.Active_ind, paleo_climate.Climate_name, paleo_climate.Climate_type, paleo_climate.Effective_date, paleo_climate.Expiry_date, paleo_climate.Ppdm_guid, paleo_climate.Remark, paleo_climate.Source, paleo_climate.Row_changed_by, paleo_climate.Row_changed_date, paleo_climate.Row_created_by, paleo_climate.Row_created_date, paleo_climate.Row_effective_date, paleo_climate.Row_expiry_date, paleo_climate.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePaleoClimate(c *fiber.Ctx) error {
	var paleo_climate dto.Paleo_climate

	setDefaults(&paleo_climate)

	if err := json.Unmarshal(c.Body(), &paleo_climate); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	paleo_climate.Climate_id = id

    if paleo_climate.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update paleo_climate set active_ind = :1, climate_name = :2, climate_type = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where climate_id = :16")
	if err != nil {
		return err
	}

	paleo_climate.Row_changed_date = formatDate(paleo_climate.Row_changed_date)
	paleo_climate.Effective_date = formatDateString(paleo_climate.Effective_date)
	paleo_climate.Expiry_date = formatDateString(paleo_climate.Expiry_date)
	paleo_climate.Row_effective_date = formatDateString(paleo_climate.Row_effective_date)
	paleo_climate.Row_expiry_date = formatDateString(paleo_climate.Row_expiry_date)






	rows, err := stmt.Exec(paleo_climate.Active_ind, paleo_climate.Climate_name, paleo_climate.Climate_type, paleo_climate.Effective_date, paleo_climate.Expiry_date, paleo_climate.Ppdm_guid, paleo_climate.Remark, paleo_climate.Source, paleo_climate.Row_changed_by, paleo_climate.Row_changed_date, paleo_climate.Row_created_by, paleo_climate.Row_effective_date, paleo_climate.Row_expiry_date, paleo_climate.Row_quality, paleo_climate.Climate_id)
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

func PatchPaleoClimate(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update paleo_climate set "
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
	query += " where climate_id = :id"

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

func DeletePaleoClimate(c *fiber.Ctx) error {
	id := c.Params("id")
	var paleo_climate dto.Paleo_climate
	paleo_climate.Climate_id = id

	stmt, err := db.Prepare("delete from paleo_climate where climate_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(paleo_climate.Climate_id)
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


