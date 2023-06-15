package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFossilDesc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from fossil_desc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Fossil_desc

	for rows.Next() {
		var fossil_desc dto.Fossil_desc
		if err := rows.Scan(&fossil_desc.Fossil_id, &fossil_desc.Fossil_desc_id, &fossil_desc.Active_ind, &fossil_desc.Date_format_desc, &fossil_desc.Description, &fossil_desc.Description_code, &fossil_desc.Description_type, &fossil_desc.Effective_date, &fossil_desc.Expiry_date, &fossil_desc.Max_value, &fossil_desc.Max_value_uom, &fossil_desc.Min_value, &fossil_desc.Min_value_uom, &fossil_desc.Ppdm_guid, &fossil_desc.Remark, &fossil_desc.Source, &fossil_desc.Source_document_id, &fossil_desc.Value_ouom, &fossil_desc.Row_changed_by, &fossil_desc.Row_changed_date, &fossil_desc.Row_created_by, &fossil_desc.Row_created_date, &fossil_desc.Row_effective_date, &fossil_desc.Row_expiry_date, &fossil_desc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append fossil_desc to result
		result = append(result, fossil_desc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFossilDesc(c *fiber.Ctx) error {
	var fossil_desc dto.Fossil_desc

	setDefaults(&fossil_desc)

	if err := json.Unmarshal(c.Body(), &fossil_desc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into fossil_desc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25)")
	if err != nil {
		return err
	}
	fossil_desc.Row_created_date = formatDate(fossil_desc.Row_created_date)
	fossil_desc.Row_changed_date = formatDate(fossil_desc.Row_changed_date)
	fossil_desc.Date_format_desc = formatDateString(fossil_desc.Date_format_desc)
	fossil_desc.Effective_date = formatDateString(fossil_desc.Effective_date)
	fossil_desc.Expiry_date = formatDateString(fossil_desc.Expiry_date)
	fossil_desc.Row_effective_date = formatDateString(fossil_desc.Row_effective_date)
	fossil_desc.Row_expiry_date = formatDateString(fossil_desc.Row_expiry_date)






	rows, err := stmt.Exec(fossil_desc.Fossil_id, fossil_desc.Fossil_desc_id, fossil_desc.Active_ind, fossil_desc.Date_format_desc, fossil_desc.Description, fossil_desc.Description_code, fossil_desc.Description_type, fossil_desc.Effective_date, fossil_desc.Expiry_date, fossil_desc.Max_value, fossil_desc.Max_value_uom, fossil_desc.Min_value, fossil_desc.Min_value_uom, fossil_desc.Ppdm_guid, fossil_desc.Remark, fossil_desc.Source, fossil_desc.Source_document_id, fossil_desc.Value_ouom, fossil_desc.Row_changed_by, fossil_desc.Row_changed_date, fossil_desc.Row_created_by, fossil_desc.Row_created_date, fossil_desc.Row_effective_date, fossil_desc.Row_expiry_date, fossil_desc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFossilDesc(c *fiber.Ctx) error {
	var fossil_desc dto.Fossil_desc

	setDefaults(&fossil_desc)

	if err := json.Unmarshal(c.Body(), &fossil_desc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	fossil_desc.Fossil_id = id

    if fossil_desc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update fossil_desc set fossil_desc_id = :1, active_ind = :2, date_format_desc = :3, description = :4, description_code = :5, description_type = :6, effective_date = :7, expiry_date = :8, max_value = :9, max_value_uom = :10, min_value = :11, min_value_uom = :12, ppdm_guid = :13, remark = :14, source = :15, source_document_id = :16, value_ouom = :17, row_changed_by = :18, row_changed_date = :19, row_created_by = :20, row_effective_date = :21, row_expiry_date = :22, row_quality = :23 where fossil_id = :25")
	if err != nil {
		return err
	}

	fossil_desc.Row_changed_date = formatDate(fossil_desc.Row_changed_date)
	fossil_desc.Date_format_desc = formatDateString(fossil_desc.Date_format_desc)
	fossil_desc.Effective_date = formatDateString(fossil_desc.Effective_date)
	fossil_desc.Expiry_date = formatDateString(fossil_desc.Expiry_date)
	fossil_desc.Row_effective_date = formatDateString(fossil_desc.Row_effective_date)
	fossil_desc.Row_expiry_date = formatDateString(fossil_desc.Row_expiry_date)






	rows, err := stmt.Exec(fossil_desc.Fossil_desc_id, fossil_desc.Active_ind, fossil_desc.Date_format_desc, fossil_desc.Description, fossil_desc.Description_code, fossil_desc.Description_type, fossil_desc.Effective_date, fossil_desc.Expiry_date, fossil_desc.Max_value, fossil_desc.Max_value_uom, fossil_desc.Min_value, fossil_desc.Min_value_uom, fossil_desc.Ppdm_guid, fossil_desc.Remark, fossil_desc.Source, fossil_desc.Source_document_id, fossil_desc.Value_ouom, fossil_desc.Row_changed_by, fossil_desc.Row_changed_date, fossil_desc.Row_created_by, fossil_desc.Row_effective_date, fossil_desc.Row_expiry_date, fossil_desc.Row_quality, fossil_desc.Fossil_id)
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

func PatchFossilDesc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update fossil_desc set "
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
		} else if key == "date_format_desc" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where fossil_id = :id"

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

func DeleteFossilDesc(c *fiber.Ctx) error {
	id := c.Params("id")
	var fossil_desc dto.Fossil_desc
	fossil_desc.Fossil_id = id

	stmt, err := db.Prepare("delete from fossil_desc where fossil_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(fossil_desc.Fossil_id)
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


