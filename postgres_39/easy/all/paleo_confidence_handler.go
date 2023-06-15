package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPaleoConfidence(c *fiber.Ctx) error {
	rows, err := db.Query("select * from paleo_confidence")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Paleo_confidence

	for rows.Next() {
		var paleo_confidence dto.Paleo_confidence
		if err := rows.Scan(&paleo_confidence.Confidence_id, &paleo_confidence.Active_ind, &paleo_confidence.Confidence_code, &paleo_confidence.Confidence_name, &paleo_confidence.Confidence_type, &paleo_confidence.Effective_date, &paleo_confidence.Expiry_date, &paleo_confidence.Ppdm_guid, &paleo_confidence.Remark, &paleo_confidence.Source, &paleo_confidence.Row_changed_by, &paleo_confidence.Row_changed_date, &paleo_confidence.Row_created_by, &paleo_confidence.Row_created_date, &paleo_confidence.Row_effective_date, &paleo_confidence.Row_expiry_date, &paleo_confidence.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append paleo_confidence to result
		result = append(result, paleo_confidence)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPaleoConfidence(c *fiber.Ctx) error {
	var paleo_confidence dto.Paleo_confidence

	setDefaults(&paleo_confidence)

	if err := json.Unmarshal(c.Body(), &paleo_confidence); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into paleo_confidence values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	paleo_confidence.Row_created_date = formatDate(paleo_confidence.Row_created_date)
	paleo_confidence.Row_changed_date = formatDate(paleo_confidence.Row_changed_date)
	paleo_confidence.Effective_date = formatDateString(paleo_confidence.Effective_date)
	paleo_confidence.Expiry_date = formatDateString(paleo_confidence.Expiry_date)
	paleo_confidence.Row_effective_date = formatDateString(paleo_confidence.Row_effective_date)
	paleo_confidence.Row_expiry_date = formatDateString(paleo_confidence.Row_expiry_date)






	rows, err := stmt.Exec(paleo_confidence.Confidence_id, paleo_confidence.Active_ind, paleo_confidence.Confidence_code, paleo_confidence.Confidence_name, paleo_confidence.Confidence_type, paleo_confidence.Effective_date, paleo_confidence.Expiry_date, paleo_confidence.Ppdm_guid, paleo_confidence.Remark, paleo_confidence.Source, paleo_confidence.Row_changed_by, paleo_confidence.Row_changed_date, paleo_confidence.Row_created_by, paleo_confidence.Row_created_date, paleo_confidence.Row_effective_date, paleo_confidence.Row_expiry_date, paleo_confidence.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePaleoConfidence(c *fiber.Ctx) error {
	var paleo_confidence dto.Paleo_confidence

	setDefaults(&paleo_confidence)

	if err := json.Unmarshal(c.Body(), &paleo_confidence); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	paleo_confidence.Confidence_id = id

    if paleo_confidence.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update paleo_confidence set active_ind = :1, confidence_code = :2, confidence_name = :3, confidence_type = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where confidence_id = :17")
	if err != nil {
		return err
	}

	paleo_confidence.Row_changed_date = formatDate(paleo_confidence.Row_changed_date)
	paleo_confidence.Effective_date = formatDateString(paleo_confidence.Effective_date)
	paleo_confidence.Expiry_date = formatDateString(paleo_confidence.Expiry_date)
	paleo_confidence.Row_effective_date = formatDateString(paleo_confidence.Row_effective_date)
	paleo_confidence.Row_expiry_date = formatDateString(paleo_confidence.Row_expiry_date)






	rows, err := stmt.Exec(paleo_confidence.Active_ind, paleo_confidence.Confidence_code, paleo_confidence.Confidence_name, paleo_confidence.Confidence_type, paleo_confidence.Effective_date, paleo_confidence.Expiry_date, paleo_confidence.Ppdm_guid, paleo_confidence.Remark, paleo_confidence.Source, paleo_confidence.Row_changed_by, paleo_confidence.Row_changed_date, paleo_confidence.Row_created_by, paleo_confidence.Row_effective_date, paleo_confidence.Row_expiry_date, paleo_confidence.Row_quality, paleo_confidence.Confidence_id)
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

func PatchPaleoConfidence(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update paleo_confidence set "
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
	query += " where confidence_id = :id"

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

func DeletePaleoConfidence(c *fiber.Ctx) error {
	id := c.Params("id")
	var paleo_confidence dto.Paleo_confidence
	paleo_confidence.Confidence_id = id

	stmt, err := db.Prepare("delete from paleo_confidence where confidence_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(paleo_confidence.Confidence_id)
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


