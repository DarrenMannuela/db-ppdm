package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPaleoObsQualifier(c *fiber.Ctx) error {
	rows, err := db.Query("select * from paleo_obs_qualifier")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Paleo_obs_qualifier

	for rows.Next() {
		var paleo_obs_qualifier dto.Paleo_obs_qualifier
		if err := rows.Scan(&paleo_obs_qualifier.Qualifier_id, &paleo_obs_qualifier.Active_ind, &paleo_obs_qualifier.Effective_date, &paleo_obs_qualifier.Expiry_date, &paleo_obs_qualifier.Ppdm_guid, &paleo_obs_qualifier.Qualifier_code, &paleo_obs_qualifier.Qualifier_long_name, &paleo_obs_qualifier.Qualifier_owner, &paleo_obs_qualifier.Qualifier_short_name, &paleo_obs_qualifier.Qualifier_type, &paleo_obs_qualifier.Remark, &paleo_obs_qualifier.Source, &paleo_obs_qualifier.Row_changed_by, &paleo_obs_qualifier.Row_changed_date, &paleo_obs_qualifier.Row_created_by, &paleo_obs_qualifier.Row_created_date, &paleo_obs_qualifier.Row_effective_date, &paleo_obs_qualifier.Row_expiry_date, &paleo_obs_qualifier.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append paleo_obs_qualifier to result
		result = append(result, paleo_obs_qualifier)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPaleoObsQualifier(c *fiber.Ctx) error {
	var paleo_obs_qualifier dto.Paleo_obs_qualifier

	setDefaults(&paleo_obs_qualifier)

	if err := json.Unmarshal(c.Body(), &paleo_obs_qualifier); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into paleo_obs_qualifier values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	paleo_obs_qualifier.Row_created_date = formatDate(paleo_obs_qualifier.Row_created_date)
	paleo_obs_qualifier.Row_changed_date = formatDate(paleo_obs_qualifier.Row_changed_date)
	paleo_obs_qualifier.Effective_date = formatDateString(paleo_obs_qualifier.Effective_date)
	paleo_obs_qualifier.Expiry_date = formatDateString(paleo_obs_qualifier.Expiry_date)
	paleo_obs_qualifier.Row_effective_date = formatDateString(paleo_obs_qualifier.Row_effective_date)
	paleo_obs_qualifier.Row_expiry_date = formatDateString(paleo_obs_qualifier.Row_expiry_date)






	rows, err := stmt.Exec(paleo_obs_qualifier.Qualifier_id, paleo_obs_qualifier.Active_ind, paleo_obs_qualifier.Effective_date, paleo_obs_qualifier.Expiry_date, paleo_obs_qualifier.Ppdm_guid, paleo_obs_qualifier.Qualifier_code, paleo_obs_qualifier.Qualifier_long_name, paleo_obs_qualifier.Qualifier_owner, paleo_obs_qualifier.Qualifier_short_name, paleo_obs_qualifier.Qualifier_type, paleo_obs_qualifier.Remark, paleo_obs_qualifier.Source, paleo_obs_qualifier.Row_changed_by, paleo_obs_qualifier.Row_changed_date, paleo_obs_qualifier.Row_created_by, paleo_obs_qualifier.Row_created_date, paleo_obs_qualifier.Row_effective_date, paleo_obs_qualifier.Row_expiry_date, paleo_obs_qualifier.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePaleoObsQualifier(c *fiber.Ctx) error {
	var paleo_obs_qualifier dto.Paleo_obs_qualifier

	setDefaults(&paleo_obs_qualifier)

	if err := json.Unmarshal(c.Body(), &paleo_obs_qualifier); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	paleo_obs_qualifier.Qualifier_id = id

    if paleo_obs_qualifier.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update paleo_obs_qualifier set active_ind = :1, effective_date = :2, expiry_date = :3, ppdm_guid = :4, qualifier_code = :5, qualifier_long_name = :6, qualifier_owner = :7, qualifier_short_name = :8, qualifier_type = :9, remark = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where qualifier_id = :19")
	if err != nil {
		return err
	}

	paleo_obs_qualifier.Row_changed_date = formatDate(paleo_obs_qualifier.Row_changed_date)
	paleo_obs_qualifier.Effective_date = formatDateString(paleo_obs_qualifier.Effective_date)
	paleo_obs_qualifier.Expiry_date = formatDateString(paleo_obs_qualifier.Expiry_date)
	paleo_obs_qualifier.Row_effective_date = formatDateString(paleo_obs_qualifier.Row_effective_date)
	paleo_obs_qualifier.Row_expiry_date = formatDateString(paleo_obs_qualifier.Row_expiry_date)






	rows, err := stmt.Exec(paleo_obs_qualifier.Active_ind, paleo_obs_qualifier.Effective_date, paleo_obs_qualifier.Expiry_date, paleo_obs_qualifier.Ppdm_guid, paleo_obs_qualifier.Qualifier_code, paleo_obs_qualifier.Qualifier_long_name, paleo_obs_qualifier.Qualifier_owner, paleo_obs_qualifier.Qualifier_short_name, paleo_obs_qualifier.Qualifier_type, paleo_obs_qualifier.Remark, paleo_obs_qualifier.Source, paleo_obs_qualifier.Row_changed_by, paleo_obs_qualifier.Row_changed_date, paleo_obs_qualifier.Row_created_by, paleo_obs_qualifier.Row_effective_date, paleo_obs_qualifier.Row_expiry_date, paleo_obs_qualifier.Row_quality, paleo_obs_qualifier.Qualifier_id)
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

func PatchPaleoObsQualifier(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update paleo_obs_qualifier set "
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

func DeletePaleoObsQualifier(c *fiber.Ctx) error {
	id := c.Params("id")
	var paleo_obs_qualifier dto.Paleo_obs_qualifier
	paleo_obs_qualifier.Qualifier_id = id

	stmt, err := db.Prepare("delete from paleo_obs_qualifier where qualifier_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(paleo_obs_qualifier.Qualifier_id)
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


