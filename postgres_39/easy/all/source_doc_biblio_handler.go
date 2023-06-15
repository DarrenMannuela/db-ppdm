package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSourceDocBiblio(c *fiber.Ctx) error {
	rows, err := db.Query("select * from source_doc_biblio")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Source_doc_biblio

	for rows.Next() {
		var source_doc_biblio dto.Source_doc_biblio
		if err := rows.Scan(&source_doc_biblio.Source_document_id, &source_doc_biblio.Biblio_obs_no, &source_doc_biblio.Active_ind, &source_doc_biblio.Document_name, &source_doc_biblio.Effective_date, &source_doc_biblio.Expiry_date, &source_doc_biblio.Ppdm_guid, &source_doc_biblio.Referenced_document, &source_doc_biblio.Remark, &source_doc_biblio.Source, &source_doc_biblio.Row_changed_by, &source_doc_biblio.Row_changed_date, &source_doc_biblio.Row_created_by, &source_doc_biblio.Row_created_date, &source_doc_biblio.Row_effective_date, &source_doc_biblio.Row_expiry_date, &source_doc_biblio.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append source_doc_biblio to result
		result = append(result, source_doc_biblio)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSourceDocBiblio(c *fiber.Ctx) error {
	var source_doc_biblio dto.Source_doc_biblio

	setDefaults(&source_doc_biblio)

	if err := json.Unmarshal(c.Body(), &source_doc_biblio); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into source_doc_biblio values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	source_doc_biblio.Row_created_date = formatDate(source_doc_biblio.Row_created_date)
	source_doc_biblio.Row_changed_date = formatDate(source_doc_biblio.Row_changed_date)
	source_doc_biblio.Effective_date = formatDateString(source_doc_biblio.Effective_date)
	source_doc_biblio.Expiry_date = formatDateString(source_doc_biblio.Expiry_date)
	source_doc_biblio.Row_effective_date = formatDateString(source_doc_biblio.Row_effective_date)
	source_doc_biblio.Row_expiry_date = formatDateString(source_doc_biblio.Row_expiry_date)






	rows, err := stmt.Exec(source_doc_biblio.Source_document_id, source_doc_biblio.Biblio_obs_no, source_doc_biblio.Active_ind, source_doc_biblio.Document_name, source_doc_biblio.Effective_date, source_doc_biblio.Expiry_date, source_doc_biblio.Ppdm_guid, source_doc_biblio.Referenced_document, source_doc_biblio.Remark, source_doc_biblio.Source, source_doc_biblio.Row_changed_by, source_doc_biblio.Row_changed_date, source_doc_biblio.Row_created_by, source_doc_biblio.Row_created_date, source_doc_biblio.Row_effective_date, source_doc_biblio.Row_expiry_date, source_doc_biblio.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSourceDocBiblio(c *fiber.Ctx) error {
	var source_doc_biblio dto.Source_doc_biblio

	setDefaults(&source_doc_biblio)

	if err := json.Unmarshal(c.Body(), &source_doc_biblio); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	source_doc_biblio.Source_document_id = id

    if source_doc_biblio.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update source_doc_biblio set biblio_obs_no = :1, active_ind = :2, document_name = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, referenced_document = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where source_document_id = :17")
	if err != nil {
		return err
	}

	source_doc_biblio.Row_changed_date = formatDate(source_doc_biblio.Row_changed_date)
	source_doc_biblio.Effective_date = formatDateString(source_doc_biblio.Effective_date)
	source_doc_biblio.Expiry_date = formatDateString(source_doc_biblio.Expiry_date)
	source_doc_biblio.Row_effective_date = formatDateString(source_doc_biblio.Row_effective_date)
	source_doc_biblio.Row_expiry_date = formatDateString(source_doc_biblio.Row_expiry_date)






	rows, err := stmt.Exec(source_doc_biblio.Biblio_obs_no, source_doc_biblio.Active_ind, source_doc_biblio.Document_name, source_doc_biblio.Effective_date, source_doc_biblio.Expiry_date, source_doc_biblio.Ppdm_guid, source_doc_biblio.Referenced_document, source_doc_biblio.Remark, source_doc_biblio.Source, source_doc_biblio.Row_changed_by, source_doc_biblio.Row_changed_date, source_doc_biblio.Row_created_by, source_doc_biblio.Row_effective_date, source_doc_biblio.Row_expiry_date, source_doc_biblio.Row_quality, source_doc_biblio.Source_document_id)
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

func PatchSourceDocBiblio(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update source_doc_biblio set "
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
	query += " where source_document_id = :id"

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

func DeleteSourceDocBiblio(c *fiber.Ctx) error {
	id := c.Params("id")
	var source_doc_biblio dto.Source_doc_biblio
	source_doc_biblio.Source_document_id = id

	stmt, err := db.Prepare("delete from source_doc_biblio where source_document_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(source_doc_biblio.Source_document_id)
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


