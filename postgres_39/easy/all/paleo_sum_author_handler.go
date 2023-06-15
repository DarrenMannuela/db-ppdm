package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPaleoSumAuthor(c *fiber.Ctx) error {
	rows, err := db.Query("select * from paleo_sum_author")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Paleo_sum_author

	for rows.Next() {
		var paleo_sum_author dto.Paleo_sum_author
		if err := rows.Scan(&paleo_sum_author.Paleo_summary_id, &paleo_sum_author.Author_id, &paleo_sum_author.Active_ind, &paleo_sum_author.Author_ba_id, &paleo_sum_author.Author_desc, &paleo_sum_author.Author_seq_no, &paleo_sum_author.Author_type, &paleo_sum_author.Effective_date, &paleo_sum_author.Expiry_date, &paleo_sum_author.Ppdm_guid, &paleo_sum_author.Remark, &paleo_sum_author.Source, &paleo_sum_author.Row_changed_by, &paleo_sum_author.Row_changed_date, &paleo_sum_author.Row_created_by, &paleo_sum_author.Row_created_date, &paleo_sum_author.Row_effective_date, &paleo_sum_author.Row_expiry_date, &paleo_sum_author.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append paleo_sum_author to result
		result = append(result, paleo_sum_author)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPaleoSumAuthor(c *fiber.Ctx) error {
	var paleo_sum_author dto.Paleo_sum_author

	setDefaults(&paleo_sum_author)

	if err := json.Unmarshal(c.Body(), &paleo_sum_author); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into paleo_sum_author values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	paleo_sum_author.Row_created_date = formatDate(paleo_sum_author.Row_created_date)
	paleo_sum_author.Row_changed_date = formatDate(paleo_sum_author.Row_changed_date)
	paleo_sum_author.Effective_date = formatDateString(paleo_sum_author.Effective_date)
	paleo_sum_author.Expiry_date = formatDateString(paleo_sum_author.Expiry_date)
	paleo_sum_author.Row_effective_date = formatDateString(paleo_sum_author.Row_effective_date)
	paleo_sum_author.Row_expiry_date = formatDateString(paleo_sum_author.Row_expiry_date)






	rows, err := stmt.Exec(paleo_sum_author.Paleo_summary_id, paleo_sum_author.Author_id, paleo_sum_author.Active_ind, paleo_sum_author.Author_ba_id, paleo_sum_author.Author_desc, paleo_sum_author.Author_seq_no, paleo_sum_author.Author_type, paleo_sum_author.Effective_date, paleo_sum_author.Expiry_date, paleo_sum_author.Ppdm_guid, paleo_sum_author.Remark, paleo_sum_author.Source, paleo_sum_author.Row_changed_by, paleo_sum_author.Row_changed_date, paleo_sum_author.Row_created_by, paleo_sum_author.Row_created_date, paleo_sum_author.Row_effective_date, paleo_sum_author.Row_expiry_date, paleo_sum_author.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePaleoSumAuthor(c *fiber.Ctx) error {
	var paleo_sum_author dto.Paleo_sum_author

	setDefaults(&paleo_sum_author)

	if err := json.Unmarshal(c.Body(), &paleo_sum_author); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	paleo_sum_author.Paleo_summary_id = id

    if paleo_sum_author.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update paleo_sum_author set author_id = :1, active_ind = :2, author_ba_id = :3, author_desc = :4, author_seq_no = :5, author_type = :6, effective_date = :7, expiry_date = :8, ppdm_guid = :9, remark = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where paleo_summary_id = :19")
	if err != nil {
		return err
	}

	paleo_sum_author.Row_changed_date = formatDate(paleo_sum_author.Row_changed_date)
	paleo_sum_author.Effective_date = formatDateString(paleo_sum_author.Effective_date)
	paleo_sum_author.Expiry_date = formatDateString(paleo_sum_author.Expiry_date)
	paleo_sum_author.Row_effective_date = formatDateString(paleo_sum_author.Row_effective_date)
	paleo_sum_author.Row_expiry_date = formatDateString(paleo_sum_author.Row_expiry_date)






	rows, err := stmt.Exec(paleo_sum_author.Author_id, paleo_sum_author.Active_ind, paleo_sum_author.Author_ba_id, paleo_sum_author.Author_desc, paleo_sum_author.Author_seq_no, paleo_sum_author.Author_type, paleo_sum_author.Effective_date, paleo_sum_author.Expiry_date, paleo_sum_author.Ppdm_guid, paleo_sum_author.Remark, paleo_sum_author.Source, paleo_sum_author.Row_changed_by, paleo_sum_author.Row_changed_date, paleo_sum_author.Row_created_by, paleo_sum_author.Row_effective_date, paleo_sum_author.Row_expiry_date, paleo_sum_author.Row_quality, paleo_sum_author.Paleo_summary_id)
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

func PatchPaleoSumAuthor(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update paleo_sum_author set "
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
	query += " where paleo_summary_id = :id"

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

func DeletePaleoSumAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	var paleo_sum_author dto.Paleo_sum_author
	paleo_sum_author.Paleo_summary_id = id

	stmt, err := db.Prepare("delete from paleo_sum_author where paleo_summary_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(paleo_sum_author.Paleo_summary_id)
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


