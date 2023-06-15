package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPaleoAbundQualifier(c *fiber.Ctx) error {
	rows, err := db.Query("select * from paleo_abund_qualifier")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Paleo_abund_qualifier

	for rows.Next() {
		var paleo_abund_qualifier dto.Paleo_abund_qualifier
		if err := rows.Scan(&paleo_abund_qualifier.Qualifier_id, &paleo_abund_qualifier.Active_ind, &paleo_abund_qualifier.Effective_date, &paleo_abund_qualifier.Expiry_date, &paleo_abund_qualifier.Ppdm_guid, &paleo_abund_qualifier.Qualifier_code, &paleo_abund_qualifier.Qualifier_name, &paleo_abund_qualifier.Remark, &paleo_abund_qualifier.Source, &paleo_abund_qualifier.Row_changed_by, &paleo_abund_qualifier.Row_changed_date, &paleo_abund_qualifier.Row_created_by, &paleo_abund_qualifier.Row_created_date, &paleo_abund_qualifier.Row_effective_date, &paleo_abund_qualifier.Row_expiry_date, &paleo_abund_qualifier.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append paleo_abund_qualifier to result
		result = append(result, paleo_abund_qualifier)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPaleoAbundQualifier(c *fiber.Ctx) error {
	var paleo_abund_qualifier dto.Paleo_abund_qualifier

	setDefaults(&paleo_abund_qualifier)

	if err := json.Unmarshal(c.Body(), &paleo_abund_qualifier); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into paleo_abund_qualifier values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	paleo_abund_qualifier.Row_created_date = formatDate(paleo_abund_qualifier.Row_created_date)
	paleo_abund_qualifier.Row_changed_date = formatDate(paleo_abund_qualifier.Row_changed_date)
	paleo_abund_qualifier.Effective_date = formatDateString(paleo_abund_qualifier.Effective_date)
	paleo_abund_qualifier.Expiry_date = formatDateString(paleo_abund_qualifier.Expiry_date)
	paleo_abund_qualifier.Row_effective_date = formatDateString(paleo_abund_qualifier.Row_effective_date)
	paleo_abund_qualifier.Row_expiry_date = formatDateString(paleo_abund_qualifier.Row_expiry_date)






	rows, err := stmt.Exec(paleo_abund_qualifier.Qualifier_id, paleo_abund_qualifier.Active_ind, paleo_abund_qualifier.Effective_date, paleo_abund_qualifier.Expiry_date, paleo_abund_qualifier.Ppdm_guid, paleo_abund_qualifier.Qualifier_code, paleo_abund_qualifier.Qualifier_name, paleo_abund_qualifier.Remark, paleo_abund_qualifier.Source, paleo_abund_qualifier.Row_changed_by, paleo_abund_qualifier.Row_changed_date, paleo_abund_qualifier.Row_created_by, paleo_abund_qualifier.Row_created_date, paleo_abund_qualifier.Row_effective_date, paleo_abund_qualifier.Row_expiry_date, paleo_abund_qualifier.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePaleoAbundQualifier(c *fiber.Ctx) error {
	var paleo_abund_qualifier dto.Paleo_abund_qualifier

	setDefaults(&paleo_abund_qualifier)

	if err := json.Unmarshal(c.Body(), &paleo_abund_qualifier); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	paleo_abund_qualifier.Qualifier_id = id

    if paleo_abund_qualifier.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update paleo_abund_qualifier set active_ind = :1, effective_date = :2, expiry_date = :3, ppdm_guid = :4, qualifier_code = :5, qualifier_name = :6, remark = :7, source = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where qualifier_id = :16")
	if err != nil {
		return err
	}

	paleo_abund_qualifier.Row_changed_date = formatDate(paleo_abund_qualifier.Row_changed_date)
	paleo_abund_qualifier.Effective_date = formatDateString(paleo_abund_qualifier.Effective_date)
	paleo_abund_qualifier.Expiry_date = formatDateString(paleo_abund_qualifier.Expiry_date)
	paleo_abund_qualifier.Row_effective_date = formatDateString(paleo_abund_qualifier.Row_effective_date)
	paleo_abund_qualifier.Row_expiry_date = formatDateString(paleo_abund_qualifier.Row_expiry_date)






	rows, err := stmt.Exec(paleo_abund_qualifier.Active_ind, paleo_abund_qualifier.Effective_date, paleo_abund_qualifier.Expiry_date, paleo_abund_qualifier.Ppdm_guid, paleo_abund_qualifier.Qualifier_code, paleo_abund_qualifier.Qualifier_name, paleo_abund_qualifier.Remark, paleo_abund_qualifier.Source, paleo_abund_qualifier.Row_changed_by, paleo_abund_qualifier.Row_changed_date, paleo_abund_qualifier.Row_created_by, paleo_abund_qualifier.Row_effective_date, paleo_abund_qualifier.Row_expiry_date, paleo_abund_qualifier.Row_quality, paleo_abund_qualifier.Qualifier_id)
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

func PatchPaleoAbundQualifier(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update paleo_abund_qualifier set "
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
	query += " where qualifier_id = :id"

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

func DeletePaleoAbundQualifier(c *fiber.Ctx) error {
	id := c.Params("id")
	var paleo_abund_qualifier dto.Paleo_abund_qualifier
	paleo_abund_qualifier.Qualifier_id = id

	stmt, err := db.Prepare("delete from paleo_abund_qualifier where qualifier_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(paleo_abund_qualifier.Qualifier_id)
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


