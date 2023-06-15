package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmThesaurusGlossary(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_thesaurus_glossary")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_thesaurus_glossary

	for rows.Next() {
		var rm_thesaurus_glossary dto.Rm_thesaurus_glossary
		if err := rows.Scan(&rm_thesaurus_glossary.Thesaurus_id, &rm_thesaurus_glossary.Thesaurus_word_id, &rm_thesaurus_glossary.Glossary_id, &rm_thesaurus_glossary.Active_ind, &rm_thesaurus_glossary.Effective_date, &rm_thesaurus_glossary.Expiry_date, &rm_thesaurus_glossary.Glossary_definition, &rm_thesaurus_glossary.Owner_ba_id, &rm_thesaurus_glossary.Ppdm_guid, &rm_thesaurus_glossary.Preferred_ind, &rm_thesaurus_glossary.Remark, &rm_thesaurus_glossary.Source, &rm_thesaurus_glossary.Source_document_id, &rm_thesaurus_glossary.Row_changed_by, &rm_thesaurus_glossary.Row_changed_date, &rm_thesaurus_glossary.Row_created_by, &rm_thesaurus_glossary.Row_created_date, &rm_thesaurus_glossary.Row_effective_date, &rm_thesaurus_glossary.Row_expiry_date, &rm_thesaurus_glossary.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_thesaurus_glossary to result
		result = append(result, rm_thesaurus_glossary)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmThesaurusGlossary(c *fiber.Ctx) error {
	var rm_thesaurus_glossary dto.Rm_thesaurus_glossary

	setDefaults(&rm_thesaurus_glossary)

	if err := json.Unmarshal(c.Body(), &rm_thesaurus_glossary); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_thesaurus_glossary values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	rm_thesaurus_glossary.Row_created_date = formatDate(rm_thesaurus_glossary.Row_created_date)
	rm_thesaurus_glossary.Row_changed_date = formatDate(rm_thesaurus_glossary.Row_changed_date)
	rm_thesaurus_glossary.Effective_date = formatDateString(rm_thesaurus_glossary.Effective_date)
	rm_thesaurus_glossary.Expiry_date = formatDateString(rm_thesaurus_glossary.Expiry_date)
	rm_thesaurus_glossary.Row_effective_date = formatDateString(rm_thesaurus_glossary.Row_effective_date)
	rm_thesaurus_glossary.Row_expiry_date = formatDateString(rm_thesaurus_glossary.Row_expiry_date)






	rows, err := stmt.Exec(rm_thesaurus_glossary.Thesaurus_id, rm_thesaurus_glossary.Thesaurus_word_id, rm_thesaurus_glossary.Glossary_id, rm_thesaurus_glossary.Active_ind, rm_thesaurus_glossary.Effective_date, rm_thesaurus_glossary.Expiry_date, rm_thesaurus_glossary.Glossary_definition, rm_thesaurus_glossary.Owner_ba_id, rm_thesaurus_glossary.Ppdm_guid, rm_thesaurus_glossary.Preferred_ind, rm_thesaurus_glossary.Remark, rm_thesaurus_glossary.Source, rm_thesaurus_glossary.Source_document_id, rm_thesaurus_glossary.Row_changed_by, rm_thesaurus_glossary.Row_changed_date, rm_thesaurus_glossary.Row_created_by, rm_thesaurus_glossary.Row_created_date, rm_thesaurus_glossary.Row_effective_date, rm_thesaurus_glossary.Row_expiry_date, rm_thesaurus_glossary.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmThesaurusGlossary(c *fiber.Ctx) error {
	var rm_thesaurus_glossary dto.Rm_thesaurus_glossary

	setDefaults(&rm_thesaurus_glossary)

	if err := json.Unmarshal(c.Body(), &rm_thesaurus_glossary); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_thesaurus_glossary.Thesaurus_id = id

    if rm_thesaurus_glossary.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_thesaurus_glossary set thesaurus_word_id = :1, glossary_id = :2, active_ind = :3, effective_date = :4, expiry_date = :5, glossary_definition = :6, owner_ba_id = :7, ppdm_guid = :8, preferred_ind = :9, remark = :10, source = :11, source_document_id = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where thesaurus_id = :20")
	if err != nil {
		return err
	}

	rm_thesaurus_glossary.Row_changed_date = formatDate(rm_thesaurus_glossary.Row_changed_date)
	rm_thesaurus_glossary.Effective_date = formatDateString(rm_thesaurus_glossary.Effective_date)
	rm_thesaurus_glossary.Expiry_date = formatDateString(rm_thesaurus_glossary.Expiry_date)
	rm_thesaurus_glossary.Row_effective_date = formatDateString(rm_thesaurus_glossary.Row_effective_date)
	rm_thesaurus_glossary.Row_expiry_date = formatDateString(rm_thesaurus_glossary.Row_expiry_date)






	rows, err := stmt.Exec(rm_thesaurus_glossary.Thesaurus_word_id, rm_thesaurus_glossary.Glossary_id, rm_thesaurus_glossary.Active_ind, rm_thesaurus_glossary.Effective_date, rm_thesaurus_glossary.Expiry_date, rm_thesaurus_glossary.Glossary_definition, rm_thesaurus_glossary.Owner_ba_id, rm_thesaurus_glossary.Ppdm_guid, rm_thesaurus_glossary.Preferred_ind, rm_thesaurus_glossary.Remark, rm_thesaurus_glossary.Source, rm_thesaurus_glossary.Source_document_id, rm_thesaurus_glossary.Row_changed_by, rm_thesaurus_glossary.Row_changed_date, rm_thesaurus_glossary.Row_created_by, rm_thesaurus_glossary.Row_effective_date, rm_thesaurus_glossary.Row_expiry_date, rm_thesaurus_glossary.Row_quality, rm_thesaurus_glossary.Thesaurus_id)
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

func PatchRmThesaurusGlossary(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_thesaurus_glossary set "
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

func DeleteRmThesaurusGlossary(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_thesaurus_glossary dto.Rm_thesaurus_glossary
	rm_thesaurus_glossary.Thesaurus_id = id

	stmt, err := db.Prepare("delete from rm_thesaurus_glossary where thesaurus_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_thesaurus_glossary.Thesaurus_id)
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


