package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetConsultDiscIssue(c *fiber.Ctx) error {
	rows, err := db.Query("select * from consult_disc_issue")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Consult_disc_issue

	for rows.Next() {
		var consult_disc_issue dto.Consult_disc_issue
		if err := rows.Scan(&consult_disc_issue.Consult_id, &consult_disc_issue.Discussion_id, &consult_disc_issue.Detail_id, &consult_disc_issue.Active_ind, &consult_disc_issue.Effective_date, &consult_disc_issue.Expiry_date, &consult_disc_issue.Ppdm_guid, &consult_disc_issue.Remark, &consult_disc_issue.Source, &consult_disc_issue.Row_changed_by, &consult_disc_issue.Row_changed_date, &consult_disc_issue.Row_created_by, &consult_disc_issue.Row_created_date, &consult_disc_issue.Row_effective_date, &consult_disc_issue.Row_expiry_date, &consult_disc_issue.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append consult_disc_issue to result
		result = append(result, consult_disc_issue)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetConsultDiscIssue(c *fiber.Ctx) error {
	var consult_disc_issue dto.Consult_disc_issue

	setDefaults(&consult_disc_issue)

	if err := json.Unmarshal(c.Body(), &consult_disc_issue); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into consult_disc_issue values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	consult_disc_issue.Row_created_date = formatDate(consult_disc_issue.Row_created_date)
	consult_disc_issue.Row_changed_date = formatDate(consult_disc_issue.Row_changed_date)
	consult_disc_issue.Effective_date = formatDateString(consult_disc_issue.Effective_date)
	consult_disc_issue.Expiry_date = formatDateString(consult_disc_issue.Expiry_date)
	consult_disc_issue.Row_effective_date = formatDateString(consult_disc_issue.Row_effective_date)
	consult_disc_issue.Row_expiry_date = formatDateString(consult_disc_issue.Row_expiry_date)






	rows, err := stmt.Exec(consult_disc_issue.Consult_id, consult_disc_issue.Discussion_id, consult_disc_issue.Detail_id, consult_disc_issue.Active_ind, consult_disc_issue.Effective_date, consult_disc_issue.Expiry_date, consult_disc_issue.Ppdm_guid, consult_disc_issue.Remark, consult_disc_issue.Source, consult_disc_issue.Row_changed_by, consult_disc_issue.Row_changed_date, consult_disc_issue.Row_created_by, consult_disc_issue.Row_created_date, consult_disc_issue.Row_effective_date, consult_disc_issue.Row_expiry_date, consult_disc_issue.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateConsultDiscIssue(c *fiber.Ctx) error {
	var consult_disc_issue dto.Consult_disc_issue

	setDefaults(&consult_disc_issue)

	if err := json.Unmarshal(c.Body(), &consult_disc_issue); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	consult_disc_issue.Consult_id = id

    if consult_disc_issue.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update consult_disc_issue set discussion_id = :1, detail_id = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where consult_id = :16")
	if err != nil {
		return err
	}

	consult_disc_issue.Row_changed_date = formatDate(consult_disc_issue.Row_changed_date)
	consult_disc_issue.Effective_date = formatDateString(consult_disc_issue.Effective_date)
	consult_disc_issue.Expiry_date = formatDateString(consult_disc_issue.Expiry_date)
	consult_disc_issue.Row_effective_date = formatDateString(consult_disc_issue.Row_effective_date)
	consult_disc_issue.Row_expiry_date = formatDateString(consult_disc_issue.Row_expiry_date)






	rows, err := stmt.Exec(consult_disc_issue.Discussion_id, consult_disc_issue.Detail_id, consult_disc_issue.Active_ind, consult_disc_issue.Effective_date, consult_disc_issue.Expiry_date, consult_disc_issue.Ppdm_guid, consult_disc_issue.Remark, consult_disc_issue.Source, consult_disc_issue.Row_changed_by, consult_disc_issue.Row_changed_date, consult_disc_issue.Row_created_by, consult_disc_issue.Row_effective_date, consult_disc_issue.Row_expiry_date, consult_disc_issue.Row_quality, consult_disc_issue.Consult_id)
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

func PatchConsultDiscIssue(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update consult_disc_issue set "
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
	query += " where consult_id = :id"

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

func DeleteConsultDiscIssue(c *fiber.Ctx) error {
	id := c.Params("id")
	var consult_disc_issue dto.Consult_disc_issue
	consult_disc_issue.Consult_id = id

	stmt, err := db.Prepare("delete from consult_disc_issue where consult_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(consult_disc_issue.Consult_id)
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


