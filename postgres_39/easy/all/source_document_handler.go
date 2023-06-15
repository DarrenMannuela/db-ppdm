package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSourceDocument(c *fiber.Ctx) error {
	rows, err := db.Query("select * from source_document")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Source_document

	for rows.Next() {
		var source_document dto.Source_document
		if err := rows.Scan(&source_document.Source_document_id, &source_document.Abbreviation, &source_document.Active_ind, &source_document.Document_title, &source_document.Document_type, &source_document.Effective_date, &source_document.Expiry_date, &source_document.Figure_reference, &source_document.Issue, &source_document.Language, &source_document.Page_reference, &source_document.Plate_reference, &source_document.Ppdm_guid, &source_document.Publication, &source_document.Publication_date, &source_document.Publication_year, &source_document.Publisher, &source_document.Remark, &source_document.Source, &source_document.Row_changed_by, &source_document.Row_changed_date, &source_document.Row_created_by, &source_document.Row_created_date, &source_document.Row_effective_date, &source_document.Row_expiry_date, &source_document.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append source_document to result
		result = append(result, source_document)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSourceDocument(c *fiber.Ctx) error {
	var source_document dto.Source_document

	setDefaults(&source_document)

	if err := json.Unmarshal(c.Body(), &source_document); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into source_document values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	source_document.Row_created_date = formatDate(source_document.Row_created_date)
	source_document.Row_changed_date = formatDate(source_document.Row_changed_date)
	source_document.Effective_date = formatDateString(source_document.Effective_date)
	source_document.Expiry_date = formatDateString(source_document.Expiry_date)
	source_document.Publication_date = formatDateString(source_document.Publication_date)
	source_document.Row_effective_date = formatDateString(source_document.Row_effective_date)
	source_document.Row_expiry_date = formatDateString(source_document.Row_expiry_date)






	rows, err := stmt.Exec(source_document.Source_document_id, source_document.Abbreviation, source_document.Active_ind, source_document.Document_title, source_document.Document_type, source_document.Effective_date, source_document.Expiry_date, source_document.Figure_reference, source_document.Issue, source_document.Language, source_document.Page_reference, source_document.Plate_reference, source_document.Ppdm_guid, source_document.Publication, source_document.Publication_date, source_document.Publication_year, source_document.Publisher, source_document.Remark, source_document.Source, source_document.Row_changed_by, source_document.Row_changed_date, source_document.Row_created_by, source_document.Row_created_date, source_document.Row_effective_date, source_document.Row_expiry_date, source_document.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSourceDocument(c *fiber.Ctx) error {
	var source_document dto.Source_document

	setDefaults(&source_document)

	if err := json.Unmarshal(c.Body(), &source_document); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	source_document.Source_document_id = id

    if source_document.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update source_document set abbreviation = :1, active_ind = :2, document_title = :3, document_type = :4, effective_date = :5, expiry_date = :6, figure_reference = :7, issue = :8, language = :9, page_reference = :10, plate_reference = :11, ppdm_guid = :12, publication = :13, publication_date = :14, publication_year = :15, publisher = :16, remark = :17, source = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where source_document_id = :26")
	if err != nil {
		return err
	}

	source_document.Row_changed_date = formatDate(source_document.Row_changed_date)
	source_document.Effective_date = formatDateString(source_document.Effective_date)
	source_document.Expiry_date = formatDateString(source_document.Expiry_date)
	source_document.Publication_date = formatDateString(source_document.Publication_date)
	source_document.Row_effective_date = formatDateString(source_document.Row_effective_date)
	source_document.Row_expiry_date = formatDateString(source_document.Row_expiry_date)






	rows, err := stmt.Exec(source_document.Abbreviation, source_document.Active_ind, source_document.Document_title, source_document.Document_type, source_document.Effective_date, source_document.Expiry_date, source_document.Figure_reference, source_document.Issue, source_document.Language, source_document.Page_reference, source_document.Plate_reference, source_document.Ppdm_guid, source_document.Publication, source_document.Publication_date, source_document.Publication_year, source_document.Publisher, source_document.Remark, source_document.Source, source_document.Row_changed_by, source_document.Row_changed_date, source_document.Row_created_by, source_document.Row_effective_date, source_document.Row_expiry_date, source_document.Row_quality, source_document.Source_document_id)
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

func PatchSourceDocument(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update source_document set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "publication_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteSourceDocument(c *fiber.Ctx) error {
	id := c.Params("id")
	var source_document dto.Source_document
	source_document.Source_document_id = id

	stmt, err := db.Prepare("delete from source_document where source_document_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(source_document.Source_document_id)
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


