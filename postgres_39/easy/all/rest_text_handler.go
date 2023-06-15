package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRestText(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rest_text")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rest_text

	for rows.Next() {
		var rest_text dto.Rest_text
		if err := rows.Scan(&rest_text.Restriction_id, &rest_text.Restriction_version, &rest_text.Text_seq_no, &rest_text.Active_ind, &rest_text.Body_of_text, &rest_text.Effective_date, &rest_text.Expiry_date, &rest_text.Ppdm_guid, &rest_text.Remark, &rest_text.Source, &rest_text.Row_changed_by, &rest_text.Row_changed_date, &rest_text.Row_created_by, &rest_text.Row_created_date, &rest_text.Row_effective_date, &rest_text.Row_expiry_date, &rest_text.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rest_text to result
		result = append(result, rest_text)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRestText(c *fiber.Ctx) error {
	var rest_text dto.Rest_text

	setDefaults(&rest_text)

	if err := json.Unmarshal(c.Body(), &rest_text); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rest_text values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	rest_text.Row_created_date = formatDate(rest_text.Row_created_date)
	rest_text.Row_changed_date = formatDate(rest_text.Row_changed_date)
	rest_text.Effective_date = formatDateString(rest_text.Effective_date)
	rest_text.Expiry_date = formatDateString(rest_text.Expiry_date)
	rest_text.Row_effective_date = formatDateString(rest_text.Row_effective_date)
	rest_text.Row_expiry_date = formatDateString(rest_text.Row_expiry_date)






	rows, err := stmt.Exec(rest_text.Restriction_id, rest_text.Restriction_version, rest_text.Text_seq_no, rest_text.Active_ind, rest_text.Body_of_text, rest_text.Effective_date, rest_text.Expiry_date, rest_text.Ppdm_guid, rest_text.Remark, rest_text.Source, rest_text.Row_changed_by, rest_text.Row_changed_date, rest_text.Row_created_by, rest_text.Row_created_date, rest_text.Row_effective_date, rest_text.Row_expiry_date, rest_text.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRestText(c *fiber.Ctx) error {
	var rest_text dto.Rest_text

	setDefaults(&rest_text)

	if err := json.Unmarshal(c.Body(), &rest_text); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rest_text.Restriction_id = id

    if rest_text.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rest_text set restriction_version = :1, text_seq_no = :2, active_ind = :3, body_of_text = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where restriction_id = :17")
	if err != nil {
		return err
	}

	rest_text.Row_changed_date = formatDate(rest_text.Row_changed_date)
	rest_text.Effective_date = formatDateString(rest_text.Effective_date)
	rest_text.Expiry_date = formatDateString(rest_text.Expiry_date)
	rest_text.Row_effective_date = formatDateString(rest_text.Row_effective_date)
	rest_text.Row_expiry_date = formatDateString(rest_text.Row_expiry_date)






	rows, err := stmt.Exec(rest_text.Restriction_version, rest_text.Text_seq_no, rest_text.Active_ind, rest_text.Body_of_text, rest_text.Effective_date, rest_text.Expiry_date, rest_text.Ppdm_guid, rest_text.Remark, rest_text.Source, rest_text.Row_changed_by, rest_text.Row_changed_date, rest_text.Row_created_by, rest_text.Row_effective_date, rest_text.Row_expiry_date, rest_text.Row_quality, rest_text.Restriction_id)
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

func PatchRestText(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rest_text set "
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
	query += " where restriction_id = :id"

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

func DeleteRestText(c *fiber.Ctx) error {
	id := c.Params("id")
	var rest_text dto.Rest_text
	rest_text.Restriction_id = id

	stmt, err := db.Prepare("delete from rest_text where restriction_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rest_text.Restriction_id)
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


