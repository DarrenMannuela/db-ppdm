package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfPad(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_pad")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_pad

	for rows.Next() {
		var sf_pad dto.Sf_pad
		if err := rows.Scan(&sf_pad.Sf_subtype, &sf_pad.Pad_id, &sf_pad.Active_ind, &sf_pad.Effective_date, &sf_pad.Expiry_date, &sf_pad.Pad_type, &sf_pad.Ppdm_guid, &sf_pad.Remark, &sf_pad.Source, &sf_pad.Row_changed_by, &sf_pad.Row_changed_date, &sf_pad.Row_created_by, &sf_pad.Row_created_date, &sf_pad.Row_effective_date, &sf_pad.Row_expiry_date, &sf_pad.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_pad to result
		result = append(result, sf_pad)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfPad(c *fiber.Ctx) error {
	var sf_pad dto.Sf_pad

	setDefaults(&sf_pad)

	if err := json.Unmarshal(c.Body(), &sf_pad); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_pad values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	sf_pad.Row_created_date = formatDate(sf_pad.Row_created_date)
	sf_pad.Row_changed_date = formatDate(sf_pad.Row_changed_date)
	sf_pad.Effective_date = formatDateString(sf_pad.Effective_date)
	sf_pad.Expiry_date = formatDateString(sf_pad.Expiry_date)
	sf_pad.Row_effective_date = formatDateString(sf_pad.Row_effective_date)
	sf_pad.Row_expiry_date = formatDateString(sf_pad.Row_expiry_date)






	rows, err := stmt.Exec(sf_pad.Sf_subtype, sf_pad.Pad_id, sf_pad.Active_ind, sf_pad.Effective_date, sf_pad.Expiry_date, sf_pad.Pad_type, sf_pad.Ppdm_guid, sf_pad.Remark, sf_pad.Source, sf_pad.Row_changed_by, sf_pad.Row_changed_date, sf_pad.Row_created_by, sf_pad.Row_created_date, sf_pad.Row_effective_date, sf_pad.Row_expiry_date, sf_pad.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfPad(c *fiber.Ctx) error {
	var sf_pad dto.Sf_pad

	setDefaults(&sf_pad)

	if err := json.Unmarshal(c.Body(), &sf_pad); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_pad.Sf_subtype = id

    if sf_pad.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_pad set pad_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, pad_type = :5, ppdm_guid = :6, remark = :7, source = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where sf_subtype = :16")
	if err != nil {
		return err
	}

	sf_pad.Row_changed_date = formatDate(sf_pad.Row_changed_date)
	sf_pad.Effective_date = formatDateString(sf_pad.Effective_date)
	sf_pad.Expiry_date = formatDateString(sf_pad.Expiry_date)
	sf_pad.Row_effective_date = formatDateString(sf_pad.Row_effective_date)
	sf_pad.Row_expiry_date = formatDateString(sf_pad.Row_expiry_date)






	rows, err := stmt.Exec(sf_pad.Pad_id, sf_pad.Active_ind, sf_pad.Effective_date, sf_pad.Expiry_date, sf_pad.Pad_type, sf_pad.Ppdm_guid, sf_pad.Remark, sf_pad.Source, sf_pad.Row_changed_by, sf_pad.Row_changed_date, sf_pad.Row_created_by, sf_pad.Row_effective_date, sf_pad.Row_expiry_date, sf_pad.Row_quality, sf_pad.Sf_subtype)
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

func PatchSfPad(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_pad set "
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
	query += " where sf_subtype = :id"

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

func DeleteSfPad(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_pad dto.Sf_pad
	sf_pad.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_pad where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_pad.Sf_subtype)
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


