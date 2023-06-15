package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmThesaurusWord(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_thesaurus_word")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_thesaurus_word

	for rows.Next() {
		var rm_thesaurus_word dto.Rm_thesaurus_word
		if err := rows.Scan(&rm_thesaurus_word.Thesaurus_id, &rm_thesaurus_word.Thesaurus_word_id, &rm_thesaurus_word.Active_ind, &rm_thesaurus_word.Effective_date, &rm_thesaurus_word.Expiry_date, &rm_thesaurus_word.Ppdm_guid, &rm_thesaurus_word.Remark, &rm_thesaurus_word.Source, &rm_thesaurus_word.Source_document_id, &rm_thesaurus_word.Thesaurus_word, &rm_thesaurus_word.Row_changed_by, &rm_thesaurus_word.Row_changed_date, &rm_thesaurus_word.Row_created_by, &rm_thesaurus_word.Row_created_date, &rm_thesaurus_word.Row_effective_date, &rm_thesaurus_word.Row_expiry_date, &rm_thesaurus_word.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_thesaurus_word to result
		result = append(result, rm_thesaurus_word)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmThesaurusWord(c *fiber.Ctx) error {
	var rm_thesaurus_word dto.Rm_thesaurus_word

	setDefaults(&rm_thesaurus_word)

	if err := json.Unmarshal(c.Body(), &rm_thesaurus_word); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_thesaurus_word values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	rm_thesaurus_word.Row_created_date = formatDate(rm_thesaurus_word.Row_created_date)
	rm_thesaurus_word.Row_changed_date = formatDate(rm_thesaurus_word.Row_changed_date)
	rm_thesaurus_word.Effective_date = formatDateString(rm_thesaurus_word.Effective_date)
	rm_thesaurus_word.Expiry_date = formatDateString(rm_thesaurus_word.Expiry_date)
	rm_thesaurus_word.Row_effective_date = formatDateString(rm_thesaurus_word.Row_effective_date)
	rm_thesaurus_word.Row_expiry_date = formatDateString(rm_thesaurus_word.Row_expiry_date)






	rows, err := stmt.Exec(rm_thesaurus_word.Thesaurus_id, rm_thesaurus_word.Thesaurus_word_id, rm_thesaurus_word.Active_ind, rm_thesaurus_word.Effective_date, rm_thesaurus_word.Expiry_date, rm_thesaurus_word.Ppdm_guid, rm_thesaurus_word.Remark, rm_thesaurus_word.Source, rm_thesaurus_word.Source_document_id, rm_thesaurus_word.Thesaurus_word, rm_thesaurus_word.Row_changed_by, rm_thesaurus_word.Row_changed_date, rm_thesaurus_word.Row_created_by, rm_thesaurus_word.Row_created_date, rm_thesaurus_word.Row_effective_date, rm_thesaurus_word.Row_expiry_date, rm_thesaurus_word.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmThesaurusWord(c *fiber.Ctx) error {
	var rm_thesaurus_word dto.Rm_thesaurus_word

	setDefaults(&rm_thesaurus_word)

	if err := json.Unmarshal(c.Body(), &rm_thesaurus_word); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_thesaurus_word.Thesaurus_id = id

    if rm_thesaurus_word.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_thesaurus_word set thesaurus_word_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, remark = :6, source = :7, source_document_id = :8, thesaurus_word = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where thesaurus_id = :17")
	if err != nil {
		return err
	}

	rm_thesaurus_word.Row_changed_date = formatDate(rm_thesaurus_word.Row_changed_date)
	rm_thesaurus_word.Effective_date = formatDateString(rm_thesaurus_word.Effective_date)
	rm_thesaurus_word.Expiry_date = formatDateString(rm_thesaurus_word.Expiry_date)
	rm_thesaurus_word.Row_effective_date = formatDateString(rm_thesaurus_word.Row_effective_date)
	rm_thesaurus_word.Row_expiry_date = formatDateString(rm_thesaurus_word.Row_expiry_date)






	rows, err := stmt.Exec(rm_thesaurus_word.Thesaurus_word_id, rm_thesaurus_word.Active_ind, rm_thesaurus_word.Effective_date, rm_thesaurus_word.Expiry_date, rm_thesaurus_word.Ppdm_guid, rm_thesaurus_word.Remark, rm_thesaurus_word.Source, rm_thesaurus_word.Source_document_id, rm_thesaurus_word.Thesaurus_word, rm_thesaurus_word.Row_changed_by, rm_thesaurus_word.Row_changed_date, rm_thesaurus_word.Row_created_by, rm_thesaurus_word.Row_effective_date, rm_thesaurus_word.Row_expiry_date, rm_thesaurus_word.Row_quality, rm_thesaurus_word.Thesaurus_id)
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

func PatchRmThesaurusWord(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_thesaurus_word set "
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
	query += " where thesaurus_id = :id"

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

func DeleteRmThesaurusWord(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_thesaurus_word dto.Rm_thesaurus_word
	rm_thesaurus_word.Thesaurus_id = id

	stmt, err := db.Prepare("delete from rm_thesaurus_word where thesaurus_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_thesaurus_word.Thesaurus_id)
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


