package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSourceDocAuthor(c *fiber.Ctx) error {
	rows, err := db.Query("select * from source_doc_author")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Source_doc_author

	for rows.Next() {
		var source_doc_author dto.Source_doc_author
		if err := rows.Scan(&source_doc_author.Source_document_id, &source_doc_author.Author_id, &source_doc_author.Active_ind, &source_doc_author.Author_ba_id, &source_doc_author.Author_first_name, &source_doc_author.Author_initial, &source_doc_author.Author_last_name, &source_doc_author.Author_seq_no, &source_doc_author.Author_type, &source_doc_author.Effective_date, &source_doc_author.Expiry_date, &source_doc_author.Ppdm_guid, &source_doc_author.Remark, &source_doc_author.Source, &source_doc_author.Row_changed_by, &source_doc_author.Row_changed_date, &source_doc_author.Row_created_by, &source_doc_author.Row_created_date, &source_doc_author.Row_effective_date, &source_doc_author.Row_expiry_date, &source_doc_author.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append source_doc_author to result
		result = append(result, source_doc_author)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSourceDocAuthor(c *fiber.Ctx) error {
	var source_doc_author dto.Source_doc_author

	setDefaults(&source_doc_author)

	if err := json.Unmarshal(c.Body(), &source_doc_author); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into source_doc_author values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	source_doc_author.Row_created_date = formatDate(source_doc_author.Row_created_date)
	source_doc_author.Row_changed_date = formatDate(source_doc_author.Row_changed_date)
	source_doc_author.Effective_date = formatDateString(source_doc_author.Effective_date)
	source_doc_author.Expiry_date = formatDateString(source_doc_author.Expiry_date)
	source_doc_author.Row_effective_date = formatDateString(source_doc_author.Row_effective_date)
	source_doc_author.Row_expiry_date = formatDateString(source_doc_author.Row_expiry_date)






	rows, err := stmt.Exec(source_doc_author.Source_document_id, source_doc_author.Author_id, source_doc_author.Active_ind, source_doc_author.Author_ba_id, source_doc_author.Author_first_name, source_doc_author.Author_initial, source_doc_author.Author_last_name, source_doc_author.Author_seq_no, source_doc_author.Author_type, source_doc_author.Effective_date, source_doc_author.Expiry_date, source_doc_author.Ppdm_guid, source_doc_author.Remark, source_doc_author.Source, source_doc_author.Row_changed_by, source_doc_author.Row_changed_date, source_doc_author.Row_created_by, source_doc_author.Row_created_date, source_doc_author.Row_effective_date, source_doc_author.Row_expiry_date, source_doc_author.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSourceDocAuthor(c *fiber.Ctx) error {
	var source_doc_author dto.Source_doc_author

	setDefaults(&source_doc_author)

	if err := json.Unmarshal(c.Body(), &source_doc_author); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	source_doc_author.Source_document_id = id

    if source_doc_author.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update source_doc_author set author_id = :1, active_ind = :2, author_ba_id = :3, author_first_name = :4, author_initial = :5, author_last_name = :6, author_seq_no = :7, author_type = :8, effective_date = :9, expiry_date = :10, ppdm_guid = :11, remark = :12, source = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where source_document_id = :21")
	if err != nil {
		return err
	}

	source_doc_author.Row_changed_date = formatDate(source_doc_author.Row_changed_date)
	source_doc_author.Effective_date = formatDateString(source_doc_author.Effective_date)
	source_doc_author.Expiry_date = formatDateString(source_doc_author.Expiry_date)
	source_doc_author.Row_effective_date = formatDateString(source_doc_author.Row_effective_date)
	source_doc_author.Row_expiry_date = formatDateString(source_doc_author.Row_expiry_date)






	rows, err := stmt.Exec(source_doc_author.Author_id, source_doc_author.Active_ind, source_doc_author.Author_ba_id, source_doc_author.Author_first_name, source_doc_author.Author_initial, source_doc_author.Author_last_name, source_doc_author.Author_seq_no, source_doc_author.Author_type, source_doc_author.Effective_date, source_doc_author.Expiry_date, source_doc_author.Ppdm_guid, source_doc_author.Remark, source_doc_author.Source, source_doc_author.Row_changed_by, source_doc_author.Row_changed_date, source_doc_author.Row_created_by, source_doc_author.Row_effective_date, source_doc_author.Row_expiry_date, source_doc_author.Row_quality, source_doc_author.Source_document_id)
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

func PatchSourceDocAuthor(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update source_doc_author set "
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

func DeleteSourceDocAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	var source_doc_author dto.Source_doc_author
	source_doc_author.Source_document_id = id

	stmt, err := db.Prepare("delete from source_doc_author where source_document_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(source_doc_author.Source_document_id)
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


