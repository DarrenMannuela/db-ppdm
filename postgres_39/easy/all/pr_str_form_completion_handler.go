package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPrStrFormCompletion(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pr_str_form_completion")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pr_str_form_completion

	for rows.Next() {
		var pr_str_form_completion dto.Pr_str_form_completion
		if err := rows.Scan(&pr_str_form_completion.Uwi, &pr_str_form_completion.Prod_string_source, &pr_str_form_completion.String_id, &pr_str_form_completion.Pr_str_form_obs_no, &pr_str_form_completion.Completion_source, &pr_str_form_completion.Completion_obs_no, &pr_str_form_completion.Form_completion_obs_no, &pr_str_form_completion.Active_ind, &pr_str_form_completion.Completion_status, &pr_str_form_completion.Effective_date, &pr_str_form_completion.Expiry_date, &pr_str_form_completion.Ppdm_guid, &pr_str_form_completion.Remark, &pr_str_form_completion.Source, &pr_str_form_completion.Status_type, &pr_str_form_completion.Row_changed_by, &pr_str_form_completion.Row_changed_date, &pr_str_form_completion.Row_created_by, &pr_str_form_completion.Row_created_date, &pr_str_form_completion.Row_effective_date, &pr_str_form_completion.Row_expiry_date, &pr_str_form_completion.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pr_str_form_completion to result
		result = append(result, pr_str_form_completion)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPrStrFormCompletion(c *fiber.Ctx) error {
	var pr_str_form_completion dto.Pr_str_form_completion

	setDefaults(&pr_str_form_completion)

	if err := json.Unmarshal(c.Body(), &pr_str_form_completion); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pr_str_form_completion values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	pr_str_form_completion.Row_created_date = formatDate(pr_str_form_completion.Row_created_date)
	pr_str_form_completion.Row_changed_date = formatDate(pr_str_form_completion.Row_changed_date)
	pr_str_form_completion.Effective_date = formatDateString(pr_str_form_completion.Effective_date)
	pr_str_form_completion.Expiry_date = formatDateString(pr_str_form_completion.Expiry_date)
	pr_str_form_completion.Row_effective_date = formatDateString(pr_str_form_completion.Row_effective_date)
	pr_str_form_completion.Row_expiry_date = formatDateString(pr_str_form_completion.Row_expiry_date)






	rows, err := stmt.Exec(pr_str_form_completion.Uwi, pr_str_form_completion.Prod_string_source, pr_str_form_completion.String_id, pr_str_form_completion.Pr_str_form_obs_no, pr_str_form_completion.Completion_source, pr_str_form_completion.Completion_obs_no, pr_str_form_completion.Form_completion_obs_no, pr_str_form_completion.Active_ind, pr_str_form_completion.Completion_status, pr_str_form_completion.Effective_date, pr_str_form_completion.Expiry_date, pr_str_form_completion.Ppdm_guid, pr_str_form_completion.Remark, pr_str_form_completion.Source, pr_str_form_completion.Status_type, pr_str_form_completion.Row_changed_by, pr_str_form_completion.Row_changed_date, pr_str_form_completion.Row_created_by, pr_str_form_completion.Row_created_date, pr_str_form_completion.Row_effective_date, pr_str_form_completion.Row_expiry_date, pr_str_form_completion.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePrStrFormCompletion(c *fiber.Ctx) error {
	var pr_str_form_completion dto.Pr_str_form_completion

	setDefaults(&pr_str_form_completion)

	if err := json.Unmarshal(c.Body(), &pr_str_form_completion); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pr_str_form_completion.Uwi = id

    if pr_str_form_completion.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pr_str_form_completion set prod_string_source = :1, string_id = :2, pr_str_form_obs_no = :3, completion_source = :4, completion_obs_no = :5, form_completion_obs_no = :6, active_ind = :7, completion_status = :8, effective_date = :9, expiry_date = :10, ppdm_guid = :11, remark = :12, source = :13, status_type = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where uwi = :22")
	if err != nil {
		return err
	}

	pr_str_form_completion.Row_changed_date = formatDate(pr_str_form_completion.Row_changed_date)
	pr_str_form_completion.Effective_date = formatDateString(pr_str_form_completion.Effective_date)
	pr_str_form_completion.Expiry_date = formatDateString(pr_str_form_completion.Expiry_date)
	pr_str_form_completion.Row_effective_date = formatDateString(pr_str_form_completion.Row_effective_date)
	pr_str_form_completion.Row_expiry_date = formatDateString(pr_str_form_completion.Row_expiry_date)






	rows, err := stmt.Exec(pr_str_form_completion.Prod_string_source, pr_str_form_completion.String_id, pr_str_form_completion.Pr_str_form_obs_no, pr_str_form_completion.Completion_source, pr_str_form_completion.Completion_obs_no, pr_str_form_completion.Form_completion_obs_no, pr_str_form_completion.Active_ind, pr_str_form_completion.Completion_status, pr_str_form_completion.Effective_date, pr_str_form_completion.Expiry_date, pr_str_form_completion.Ppdm_guid, pr_str_form_completion.Remark, pr_str_form_completion.Source, pr_str_form_completion.Status_type, pr_str_form_completion.Row_changed_by, pr_str_form_completion.Row_changed_date, pr_str_form_completion.Row_created_by, pr_str_form_completion.Row_effective_date, pr_str_form_completion.Row_expiry_date, pr_str_form_completion.Row_quality, pr_str_form_completion.Uwi)
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

func PatchPrStrFormCompletion(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pr_str_form_completion set "
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
	query += " where uwi = :id"

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

func DeletePrStrFormCompletion(c *fiber.Ctx) error {
	id := c.Params("id")
	var pr_str_form_completion dto.Pr_str_form_completion
	pr_str_form_completion.Uwi = id

	stmt, err := db.Prepare("delete from pr_str_form_completion where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pr_str_form_completion.Uwi)
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


