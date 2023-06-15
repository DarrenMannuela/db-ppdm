package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmKeyword(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_keyword")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_keyword

	for rows.Next() {
		var rm_keyword dto.Rm_keyword
		if err := rows.Scan(&rm_keyword.Info_item_subtype, &rm_keyword.Information_item_id, &rm_keyword.Keyword_id, &rm_keyword.Active_ind, &rm_keyword.Effective_date, &rm_keyword.Expiry_date, &rm_keyword.Ppdm_guid, &rm_keyword.Remark, &rm_keyword.Reported_keyword, &rm_keyword.Source, &rm_keyword.Thesaurus_id, &rm_keyword.Thesaurus_word_id, &rm_keyword.Row_changed_by, &rm_keyword.Row_changed_date, &rm_keyword.Row_created_by, &rm_keyword.Row_created_date, &rm_keyword.Row_effective_date, &rm_keyword.Row_expiry_date, &rm_keyword.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_keyword to result
		result = append(result, rm_keyword)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmKeyword(c *fiber.Ctx) error {
	var rm_keyword dto.Rm_keyword

	setDefaults(&rm_keyword)

	if err := json.Unmarshal(c.Body(), &rm_keyword); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_keyword values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	rm_keyword.Row_created_date = formatDate(rm_keyword.Row_created_date)
	rm_keyword.Row_changed_date = formatDate(rm_keyword.Row_changed_date)
	rm_keyword.Effective_date = formatDateString(rm_keyword.Effective_date)
	rm_keyword.Expiry_date = formatDateString(rm_keyword.Expiry_date)
	rm_keyword.Row_effective_date = formatDateString(rm_keyword.Row_effective_date)
	rm_keyword.Row_expiry_date = formatDateString(rm_keyword.Row_expiry_date)






	rows, err := stmt.Exec(rm_keyword.Info_item_subtype, rm_keyword.Information_item_id, rm_keyword.Keyword_id, rm_keyword.Active_ind, rm_keyword.Effective_date, rm_keyword.Expiry_date, rm_keyword.Ppdm_guid, rm_keyword.Remark, rm_keyword.Reported_keyword, rm_keyword.Source, rm_keyword.Thesaurus_id, rm_keyword.Thesaurus_word_id, rm_keyword.Row_changed_by, rm_keyword.Row_changed_date, rm_keyword.Row_created_by, rm_keyword.Row_created_date, rm_keyword.Row_effective_date, rm_keyword.Row_expiry_date, rm_keyword.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmKeyword(c *fiber.Ctx) error {
	var rm_keyword dto.Rm_keyword

	setDefaults(&rm_keyword)

	if err := json.Unmarshal(c.Body(), &rm_keyword); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_keyword.Info_item_subtype = id

    if rm_keyword.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_keyword set information_item_id = :1, keyword_id = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, reported_keyword = :8, source = :9, thesaurus_id = :10, thesaurus_word_id = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where info_item_subtype = :19")
	if err != nil {
		return err
	}

	rm_keyword.Row_changed_date = formatDate(rm_keyword.Row_changed_date)
	rm_keyword.Effective_date = formatDateString(rm_keyword.Effective_date)
	rm_keyword.Expiry_date = formatDateString(rm_keyword.Expiry_date)
	rm_keyword.Row_effective_date = formatDateString(rm_keyword.Row_effective_date)
	rm_keyword.Row_expiry_date = formatDateString(rm_keyword.Row_expiry_date)






	rows, err := stmt.Exec(rm_keyword.Information_item_id, rm_keyword.Keyword_id, rm_keyword.Active_ind, rm_keyword.Effective_date, rm_keyword.Expiry_date, rm_keyword.Ppdm_guid, rm_keyword.Remark, rm_keyword.Reported_keyword, rm_keyword.Source, rm_keyword.Thesaurus_id, rm_keyword.Thesaurus_word_id, rm_keyword.Row_changed_by, rm_keyword.Row_changed_date, rm_keyword.Row_created_by, rm_keyword.Row_effective_date, rm_keyword.Row_expiry_date, rm_keyword.Row_quality, rm_keyword.Info_item_subtype)
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

func PatchRmKeyword(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_keyword set "
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
	query += " where info_item_subtype = :id"

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

func DeleteRmKeyword(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_keyword dto.Rm_keyword
	rm_keyword.Info_item_subtype = id

	stmt, err := db.Prepare("delete from rm_keyword where info_item_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_keyword.Info_item_subtype)
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


