package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFossilXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from fossil_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Fossil_xref

	for rows.Next() {
		var fossil_xref dto.Fossil_xref
		if err := rows.Scan(&fossil_xref.Fossil_id1, &fossil_xref.Fossil_id2, &fossil_xref.Fossil_xref_id, &fossil_xref.Active_ind, &fossil_xref.Effective_date, &fossil_xref.Expiry_date, &fossil_xref.Fossil_xref_type, &fossil_xref.Ppdm_guid, &fossil_xref.Remark, &fossil_xref.Source, &fossil_xref.Source_document_id, &fossil_xref.Row_changed_by, &fossil_xref.Row_changed_date, &fossil_xref.Row_created_by, &fossil_xref.Row_created_date, &fossil_xref.Row_effective_date, &fossil_xref.Row_expiry_date, &fossil_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append fossil_xref to result
		result = append(result, fossil_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFossilXref(c *fiber.Ctx) error {
	var fossil_xref dto.Fossil_xref

	setDefaults(&fossil_xref)

	if err := json.Unmarshal(c.Body(), &fossil_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into fossil_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	fossil_xref.Row_created_date = formatDate(fossil_xref.Row_created_date)
	fossil_xref.Row_changed_date = formatDate(fossil_xref.Row_changed_date)
	fossil_xref.Effective_date = formatDateString(fossil_xref.Effective_date)
	fossil_xref.Expiry_date = formatDateString(fossil_xref.Expiry_date)
	fossil_xref.Row_effective_date = formatDateString(fossil_xref.Row_effective_date)
	fossil_xref.Row_expiry_date = formatDateString(fossil_xref.Row_expiry_date)






	rows, err := stmt.Exec(fossil_xref.Fossil_id1, fossil_xref.Fossil_id2, fossil_xref.Fossil_xref_id, fossil_xref.Active_ind, fossil_xref.Effective_date, fossil_xref.Expiry_date, fossil_xref.Fossil_xref_type, fossil_xref.Ppdm_guid, fossil_xref.Remark, fossil_xref.Source, fossil_xref.Source_document_id, fossil_xref.Row_changed_by, fossil_xref.Row_changed_date, fossil_xref.Row_created_by, fossil_xref.Row_created_date, fossil_xref.Row_effective_date, fossil_xref.Row_expiry_date, fossil_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFossilXref(c *fiber.Ctx) error {
	var fossil_xref dto.Fossil_xref

	setDefaults(&fossil_xref)

	if err := json.Unmarshal(c.Body(), &fossil_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	fossil_xref.Fossil_id1 = id

    if fossil_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update fossil_xref set fossil_id2 = :1, fossil_xref_id = :2, active_ind = :3, effective_date = :4, expiry_date = :5, fossil_xref_type = :6, ppdm_guid = :7, remark = :8, source = :9, source_document_id = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where fossil_id1 = :18")
	if err != nil {
		return err
	}

	fossil_xref.Row_changed_date = formatDate(fossil_xref.Row_changed_date)
	fossil_xref.Effective_date = formatDateString(fossil_xref.Effective_date)
	fossil_xref.Expiry_date = formatDateString(fossil_xref.Expiry_date)
	fossil_xref.Row_effective_date = formatDateString(fossil_xref.Row_effective_date)
	fossil_xref.Row_expiry_date = formatDateString(fossil_xref.Row_expiry_date)






	rows, err := stmt.Exec(fossil_xref.Fossil_id2, fossil_xref.Fossil_xref_id, fossil_xref.Active_ind, fossil_xref.Effective_date, fossil_xref.Expiry_date, fossil_xref.Fossil_xref_type, fossil_xref.Ppdm_guid, fossil_xref.Remark, fossil_xref.Source, fossil_xref.Source_document_id, fossil_xref.Row_changed_by, fossil_xref.Row_changed_date, fossil_xref.Row_created_by, fossil_xref.Row_effective_date, fossil_xref.Row_expiry_date, fossil_xref.Row_quality, fossil_xref.Fossil_id1)
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

func PatchFossilXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update fossil_xref set "
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
	query += " where fossil_id1 = :id"

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

func DeleteFossilXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var fossil_xref dto.Fossil_xref
	fossil_xref.Fossil_id1 = id

	stmt, err := db.Prepare("delete from fossil_xref where fossil_id1 = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(fossil_xref.Fossil_id1)
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


