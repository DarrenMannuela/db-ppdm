package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFossilNameSetFossil(c *fiber.Ctx) error {
	rows, err := db.Query("select * from fossil_name_set_fossil")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Fossil_name_set_fossil

	for rows.Next() {
		var fossil_name_set_fossil dto.Fossil_name_set_fossil
		if err := rows.Scan(&fossil_name_set_fossil.Fossil_name_set_id, &fossil_name_set_fossil.Fossil_id, &fossil_name_set_fossil.Active_ind, &fossil_name_set_fossil.Effective_date, &fossil_name_set_fossil.Expiry_date, &fossil_name_set_fossil.Ppdm_guid, &fossil_name_set_fossil.Remark, &fossil_name_set_fossil.Source, &fossil_name_set_fossil.Row_changed_by, &fossil_name_set_fossil.Row_changed_date, &fossil_name_set_fossil.Row_created_by, &fossil_name_set_fossil.Row_created_date, &fossil_name_set_fossil.Row_effective_date, &fossil_name_set_fossil.Row_expiry_date, &fossil_name_set_fossil.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append fossil_name_set_fossil to result
		result = append(result, fossil_name_set_fossil)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFossilNameSetFossil(c *fiber.Ctx) error {
	var fossil_name_set_fossil dto.Fossil_name_set_fossil

	setDefaults(&fossil_name_set_fossil)

	if err := json.Unmarshal(c.Body(), &fossil_name_set_fossil); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into fossil_name_set_fossil values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15)")
	if err != nil {
		return err
	}
	fossil_name_set_fossil.Row_created_date = formatDate(fossil_name_set_fossil.Row_created_date)
	fossil_name_set_fossil.Row_changed_date = formatDate(fossil_name_set_fossil.Row_changed_date)
	fossil_name_set_fossil.Effective_date = formatDateString(fossil_name_set_fossil.Effective_date)
	fossil_name_set_fossil.Expiry_date = formatDateString(fossil_name_set_fossil.Expiry_date)
	fossil_name_set_fossil.Row_effective_date = formatDateString(fossil_name_set_fossil.Row_effective_date)
	fossil_name_set_fossil.Row_expiry_date = formatDateString(fossil_name_set_fossil.Row_expiry_date)






	rows, err := stmt.Exec(fossil_name_set_fossil.Fossil_name_set_id, fossil_name_set_fossil.Fossil_id, fossil_name_set_fossil.Active_ind, fossil_name_set_fossil.Effective_date, fossil_name_set_fossil.Expiry_date, fossil_name_set_fossil.Ppdm_guid, fossil_name_set_fossil.Remark, fossil_name_set_fossil.Source, fossil_name_set_fossil.Row_changed_by, fossil_name_set_fossil.Row_changed_date, fossil_name_set_fossil.Row_created_by, fossil_name_set_fossil.Row_created_date, fossil_name_set_fossil.Row_effective_date, fossil_name_set_fossil.Row_expiry_date, fossil_name_set_fossil.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFossilNameSetFossil(c *fiber.Ctx) error {
	var fossil_name_set_fossil dto.Fossil_name_set_fossil

	setDefaults(&fossil_name_set_fossil)

	if err := json.Unmarshal(c.Body(), &fossil_name_set_fossil); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	fossil_name_set_fossil.Fossil_name_set_id = id

    if fossil_name_set_fossil.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update fossil_name_set_fossil set fossil_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, remark = :6, source = :7, row_changed_by = :8, row_changed_date = :9, row_created_by = :10, row_effective_date = :11, row_expiry_date = :12, row_quality = :13 where fossil_name_set_id = :15")
	if err != nil {
		return err
	}

	fossil_name_set_fossil.Row_changed_date = formatDate(fossil_name_set_fossil.Row_changed_date)
	fossil_name_set_fossil.Effective_date = formatDateString(fossil_name_set_fossil.Effective_date)
	fossil_name_set_fossil.Expiry_date = formatDateString(fossil_name_set_fossil.Expiry_date)
	fossil_name_set_fossil.Row_effective_date = formatDateString(fossil_name_set_fossil.Row_effective_date)
	fossil_name_set_fossil.Row_expiry_date = formatDateString(fossil_name_set_fossil.Row_expiry_date)






	rows, err := stmt.Exec(fossil_name_set_fossil.Fossil_id, fossil_name_set_fossil.Active_ind, fossil_name_set_fossil.Effective_date, fossil_name_set_fossil.Expiry_date, fossil_name_set_fossil.Ppdm_guid, fossil_name_set_fossil.Remark, fossil_name_set_fossil.Source, fossil_name_set_fossil.Row_changed_by, fossil_name_set_fossil.Row_changed_date, fossil_name_set_fossil.Row_created_by, fossil_name_set_fossil.Row_effective_date, fossil_name_set_fossil.Row_expiry_date, fossil_name_set_fossil.Row_quality, fossil_name_set_fossil.Fossil_name_set_id)
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

func PatchFossilNameSetFossil(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update fossil_name_set_fossil set "
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
	query += " where fossil_name_set_id = :id"

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

func DeleteFossilNameSetFossil(c *fiber.Ctx) error {
	id := c.Params("id")
	var fossil_name_set_fossil dto.Fossil_name_set_fossil
	fossil_name_set_fossil.Fossil_name_set_id = id

	stmt, err := db.Prepare("delete from fossil_name_set_fossil where fossil_name_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(fossil_name_set_fossil.Fossil_name_set_id)
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


