package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetConsultIssue(c *fiber.Ctx) error {
	rows, err := db.Query("select * from consult_issue")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Consult_issue

	for rows.Next() {
		var consult_issue dto.Consult_issue
		if err := rows.Scan(&consult_issue.Consult_id, &consult_issue.Detail_id, &consult_issue.Active_ind, &consult_issue.Detail_desc, &consult_issue.Discussion_id, &consult_issue.Effective_date, &consult_issue.Expiry_date, &consult_issue.Issue_type, &consult_issue.Ppdm_guid, &consult_issue.Remark, &consult_issue.Remark_type, &consult_issue.Resolution, &consult_issue.Source, &consult_issue.Row_changed_by, &consult_issue.Row_changed_date, &consult_issue.Row_created_by, &consult_issue.Row_created_date, &consult_issue.Row_effective_date, &consult_issue.Row_expiry_date, &consult_issue.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append consult_issue to result
		result = append(result, consult_issue)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetConsultIssue(c *fiber.Ctx) error {
	var consult_issue dto.Consult_issue

	setDefaults(&consult_issue)

	if err := json.Unmarshal(c.Body(), &consult_issue); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into consult_issue values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	consult_issue.Row_created_date = formatDate(consult_issue.Row_created_date)
	consult_issue.Row_changed_date = formatDate(consult_issue.Row_changed_date)
	consult_issue.Effective_date = formatDateString(consult_issue.Effective_date)
	consult_issue.Expiry_date = formatDateString(consult_issue.Expiry_date)
	consult_issue.Row_effective_date = formatDateString(consult_issue.Row_effective_date)
	consult_issue.Row_expiry_date = formatDateString(consult_issue.Row_expiry_date)






	rows, err := stmt.Exec(consult_issue.Consult_id, consult_issue.Detail_id, consult_issue.Active_ind, consult_issue.Detail_desc, consult_issue.Discussion_id, consult_issue.Effective_date, consult_issue.Expiry_date, consult_issue.Issue_type, consult_issue.Ppdm_guid, consult_issue.Remark, consult_issue.Remark_type, consult_issue.Resolution, consult_issue.Source, consult_issue.Row_changed_by, consult_issue.Row_changed_date, consult_issue.Row_created_by, consult_issue.Row_created_date, consult_issue.Row_effective_date, consult_issue.Row_expiry_date, consult_issue.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateConsultIssue(c *fiber.Ctx) error {
	var consult_issue dto.Consult_issue

	setDefaults(&consult_issue)

	if err := json.Unmarshal(c.Body(), &consult_issue); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	consult_issue.Consult_id = id

    if consult_issue.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update consult_issue set detail_id = :1, active_ind = :2, detail_desc = :3, discussion_id = :4, effective_date = :5, expiry_date = :6, issue_type = :7, ppdm_guid = :8, remark = :9, remark_type = :10, resolution = :11, source = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where consult_id = :20")
	if err != nil {
		return err
	}

	consult_issue.Row_changed_date = formatDate(consult_issue.Row_changed_date)
	consult_issue.Effective_date = formatDateString(consult_issue.Effective_date)
	consult_issue.Expiry_date = formatDateString(consult_issue.Expiry_date)
	consult_issue.Row_effective_date = formatDateString(consult_issue.Row_effective_date)
	consult_issue.Row_expiry_date = formatDateString(consult_issue.Row_expiry_date)






	rows, err := stmt.Exec(consult_issue.Detail_id, consult_issue.Active_ind, consult_issue.Detail_desc, consult_issue.Discussion_id, consult_issue.Effective_date, consult_issue.Expiry_date, consult_issue.Issue_type, consult_issue.Ppdm_guid, consult_issue.Remark, consult_issue.Remark_type, consult_issue.Resolution, consult_issue.Source, consult_issue.Row_changed_by, consult_issue.Row_changed_date, consult_issue.Row_created_by, consult_issue.Row_effective_date, consult_issue.Row_expiry_date, consult_issue.Row_quality, consult_issue.Consult_id)
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

func PatchConsultIssue(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update consult_issue set "
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

func DeleteConsultIssue(c *fiber.Ctx) error {
	id := c.Params("id")
	var consult_issue dto.Consult_issue
	consult_issue.Consult_id = id

	stmt, err := db.Prepare("delete from consult_issue where consult_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(consult_issue.Consult_id)
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


