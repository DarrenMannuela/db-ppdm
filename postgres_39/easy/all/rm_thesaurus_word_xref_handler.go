package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmThesaurusWordXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_thesaurus_word_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_thesaurus_word_xref

	for rows.Next() {
		var rm_thesaurus_word_xref dto.Rm_thesaurus_word_xref
		if err := rows.Scan(&rm_thesaurus_word_xref.Thesaurus_id1, &rm_thesaurus_word_xref.Thesaurus_word_id1, &rm_thesaurus_word_xref.Thesaurus_id2, &rm_thesaurus_word_xref.Thesaurus_word_id2, &rm_thesaurus_word_xref.Xref_obs_no, &rm_thesaurus_word_xref.Active_ind, &rm_thesaurus_word_xref.Effective_date, &rm_thesaurus_word_xref.Expiry_date, &rm_thesaurus_word_xref.Ppdm_guid, &rm_thesaurus_word_xref.Relationship_description, &rm_thesaurus_word_xref.Remark, &rm_thesaurus_word_xref.Source, &rm_thesaurus_word_xref.Source_document_id, &rm_thesaurus_word_xref.Thesaurus_xref_type, &rm_thesaurus_word_xref.Row_changed_by, &rm_thesaurus_word_xref.Row_changed_date, &rm_thesaurus_word_xref.Row_created_by, &rm_thesaurus_word_xref.Row_created_date, &rm_thesaurus_word_xref.Row_effective_date, &rm_thesaurus_word_xref.Row_expiry_date, &rm_thesaurus_word_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_thesaurus_word_xref to result
		result = append(result, rm_thesaurus_word_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmThesaurusWordXref(c *fiber.Ctx) error {
	var rm_thesaurus_word_xref dto.Rm_thesaurus_word_xref

	setDefaults(&rm_thesaurus_word_xref)

	if err := json.Unmarshal(c.Body(), &rm_thesaurus_word_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_thesaurus_word_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	rm_thesaurus_word_xref.Row_created_date = formatDate(rm_thesaurus_word_xref.Row_created_date)
	rm_thesaurus_word_xref.Row_changed_date = formatDate(rm_thesaurus_word_xref.Row_changed_date)
	rm_thesaurus_word_xref.Effective_date = formatDateString(rm_thesaurus_word_xref.Effective_date)
	rm_thesaurus_word_xref.Expiry_date = formatDateString(rm_thesaurus_word_xref.Expiry_date)
	rm_thesaurus_word_xref.Row_effective_date = formatDateString(rm_thesaurus_word_xref.Row_effective_date)
	rm_thesaurus_word_xref.Row_expiry_date = formatDateString(rm_thesaurus_word_xref.Row_expiry_date)






	rows, err := stmt.Exec(rm_thesaurus_word_xref.Thesaurus_id1, rm_thesaurus_word_xref.Thesaurus_word_id1, rm_thesaurus_word_xref.Thesaurus_id2, rm_thesaurus_word_xref.Thesaurus_word_id2, rm_thesaurus_word_xref.Xref_obs_no, rm_thesaurus_word_xref.Active_ind, rm_thesaurus_word_xref.Effective_date, rm_thesaurus_word_xref.Expiry_date, rm_thesaurus_word_xref.Ppdm_guid, rm_thesaurus_word_xref.Relationship_description, rm_thesaurus_word_xref.Remark, rm_thesaurus_word_xref.Source, rm_thesaurus_word_xref.Source_document_id, rm_thesaurus_word_xref.Thesaurus_xref_type, rm_thesaurus_word_xref.Row_changed_by, rm_thesaurus_word_xref.Row_changed_date, rm_thesaurus_word_xref.Row_created_by, rm_thesaurus_word_xref.Row_created_date, rm_thesaurus_word_xref.Row_effective_date, rm_thesaurus_word_xref.Row_expiry_date, rm_thesaurus_word_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmThesaurusWordXref(c *fiber.Ctx) error {
	var rm_thesaurus_word_xref dto.Rm_thesaurus_word_xref

	setDefaults(&rm_thesaurus_word_xref)

	if err := json.Unmarshal(c.Body(), &rm_thesaurus_word_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_thesaurus_word_xref.Thesaurus_id1 = id

    if rm_thesaurus_word_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_thesaurus_word_xref set thesaurus_word_id1 = :1, thesaurus_id2 = :2, thesaurus_word_id2 = :3, xref_obs_no = :4, active_ind = :5, effective_date = :6, expiry_date = :7, ppdm_guid = :8, relationship_description = :9, remark = :10, source = :11, source_document_id = :12, thesaurus_xref_type = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where thesaurus_id1 = :21")
	if err != nil {
		return err
	}

	rm_thesaurus_word_xref.Row_changed_date = formatDate(rm_thesaurus_word_xref.Row_changed_date)
	rm_thesaurus_word_xref.Effective_date = formatDateString(rm_thesaurus_word_xref.Effective_date)
	rm_thesaurus_word_xref.Expiry_date = formatDateString(rm_thesaurus_word_xref.Expiry_date)
	rm_thesaurus_word_xref.Row_effective_date = formatDateString(rm_thesaurus_word_xref.Row_effective_date)
	rm_thesaurus_word_xref.Row_expiry_date = formatDateString(rm_thesaurus_word_xref.Row_expiry_date)






	rows, err := stmt.Exec(rm_thesaurus_word_xref.Thesaurus_word_id1, rm_thesaurus_word_xref.Thesaurus_id2, rm_thesaurus_word_xref.Thesaurus_word_id2, rm_thesaurus_word_xref.Xref_obs_no, rm_thesaurus_word_xref.Active_ind, rm_thesaurus_word_xref.Effective_date, rm_thesaurus_word_xref.Expiry_date, rm_thesaurus_word_xref.Ppdm_guid, rm_thesaurus_word_xref.Relationship_description, rm_thesaurus_word_xref.Remark, rm_thesaurus_word_xref.Source, rm_thesaurus_word_xref.Source_document_id, rm_thesaurus_word_xref.Thesaurus_xref_type, rm_thesaurus_word_xref.Row_changed_by, rm_thesaurus_word_xref.Row_changed_date, rm_thesaurus_word_xref.Row_created_by, rm_thesaurus_word_xref.Row_effective_date, rm_thesaurus_word_xref.Row_expiry_date, rm_thesaurus_word_xref.Row_quality, rm_thesaurus_word_xref.Thesaurus_id1)
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

func PatchRmThesaurusWordXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_thesaurus_word_xref set "
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
	query += " where thesaurus_id1 = :id"

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

func DeleteRmThesaurusWordXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_thesaurus_word_xref dto.Rm_thesaurus_word_xref
	rm_thesaurus_word_xref.Thesaurus_id1 = id

	stmt, err := db.Prepare("delete from rm_thesaurus_word_xref where thesaurus_id1 = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_thesaurus_word_xref.Thesaurus_id1)
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


