package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFossilDocument(c *fiber.Ctx) error {
	rows, err := db.Query("select * from fossil_document")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Fossil_document

	for rows.Next() {
		var fossil_document dto.Fossil_document
		if err := rows.Scan(&fossil_document.Fossil_id, &fossil_document.Source_document_id, &fossil_document.Active_ind, &fossil_document.Effective_date, &fossil_document.Expiry_date, &fossil_document.Ppdm_guid, &fossil_document.Remark, &fossil_document.Source, &fossil_document.Row_changed_by, &fossil_document.Row_changed_date, &fossil_document.Row_created_by, &fossil_document.Row_created_date, &fossil_document.Row_effective_date, &fossil_document.Row_expiry_date, &fossil_document.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append fossil_document to result
		result = append(result, fossil_document)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFossilDocument(c *fiber.Ctx) error {
	var fossil_document dto.Fossil_document

	setDefaults(&fossil_document)

	if err := json.Unmarshal(c.Body(), &fossil_document); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into fossil_document values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15)")
	if err != nil {
		return err
	}
	fossil_document.Row_created_date = formatDate(fossil_document.Row_created_date)
	fossil_document.Row_changed_date = formatDate(fossil_document.Row_changed_date)
	fossil_document.Effective_date = formatDateString(fossil_document.Effective_date)
	fossil_document.Expiry_date = formatDateString(fossil_document.Expiry_date)
	fossil_document.Row_effective_date = formatDateString(fossil_document.Row_effective_date)
	fossil_document.Row_expiry_date = formatDateString(fossil_document.Row_expiry_date)






	rows, err := stmt.Exec(fossil_document.Fossil_id, fossil_document.Source_document_id, fossil_document.Active_ind, fossil_document.Effective_date, fossil_document.Expiry_date, fossil_document.Ppdm_guid, fossil_document.Remark, fossil_document.Source, fossil_document.Row_changed_by, fossil_document.Row_changed_date, fossil_document.Row_created_by, fossil_document.Row_created_date, fossil_document.Row_effective_date, fossil_document.Row_expiry_date, fossil_document.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFossilDocument(c *fiber.Ctx) error {
	var fossil_document dto.Fossil_document

	setDefaults(&fossil_document)

	if err := json.Unmarshal(c.Body(), &fossil_document); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	fossil_document.Fossil_id = id

    if fossil_document.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update fossil_document set source_document_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, remark = :6, source = :7, row_changed_by = :8, row_changed_date = :9, row_created_by = :10, row_effective_date = :11, row_expiry_date = :12, row_quality = :13 where fossil_id = :15")
	if err != nil {
		return err
	}

	fossil_document.Row_changed_date = formatDate(fossil_document.Row_changed_date)
	fossil_document.Effective_date = formatDateString(fossil_document.Effective_date)
	fossil_document.Expiry_date = formatDateString(fossil_document.Expiry_date)
	fossil_document.Row_effective_date = formatDateString(fossil_document.Row_effective_date)
	fossil_document.Row_expiry_date = formatDateString(fossil_document.Row_expiry_date)






	rows, err := stmt.Exec(fossil_document.Source_document_id, fossil_document.Active_ind, fossil_document.Effective_date, fossil_document.Expiry_date, fossil_document.Ppdm_guid, fossil_document.Remark, fossil_document.Source, fossil_document.Row_changed_by, fossil_document.Row_changed_date, fossil_document.Row_created_by, fossil_document.Row_effective_date, fossil_document.Row_expiry_date, fossil_document.Row_quality, fossil_document.Fossil_id)
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

func PatchFossilDocument(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update fossil_document set "
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

func DeleteFossilDocument(c *fiber.Ctx) error {
	id := c.Params("id")
	var fossil_document dto.Fossil_document
	fossil_document.Fossil_id = id

	stmt, err := db.Prepare("delete from fossil_document where fossil_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(fossil_document.Fossil_id)
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


