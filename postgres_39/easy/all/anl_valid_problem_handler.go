package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlValidProblem(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_valid_problem")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_valid_problem

	for rows.Next() {
		var anl_valid_problem dto.Anl_valid_problem
		if err := rows.Scan(&anl_valid_problem.Method_set_id, &anl_valid_problem.Method_id, &anl_valid_problem.Problem_obs_no, &anl_valid_problem.Accuracy_type, &anl_valid_problem.Active_ind, &anl_valid_problem.Confidence_type, &anl_valid_problem.Effective_date, &anl_valid_problem.Expiry_date, &anl_valid_problem.Ppdm_guid, &anl_valid_problem.Problem_desc, &anl_valid_problem.Problem_result, &anl_valid_problem.Problem_severity, &anl_valid_problem.Problem_type, &anl_valid_problem.Remark, &anl_valid_problem.Resolution_method, &anl_valid_problem.Source, &anl_valid_problem.Substance_id, &anl_valid_problem.Row_changed_by, &anl_valid_problem.Row_changed_date, &anl_valid_problem.Row_created_by, &anl_valid_problem.Row_created_date, &anl_valid_problem.Row_effective_date, &anl_valid_problem.Row_expiry_date, &anl_valid_problem.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_valid_problem to result
		result = append(result, anl_valid_problem)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlValidProblem(c *fiber.Ctx) error {
	var anl_valid_problem dto.Anl_valid_problem

	setDefaults(&anl_valid_problem)

	if err := json.Unmarshal(c.Body(), &anl_valid_problem); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_valid_problem values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	anl_valid_problem.Row_created_date = formatDate(anl_valid_problem.Row_created_date)
	anl_valid_problem.Row_changed_date = formatDate(anl_valid_problem.Row_changed_date)
	anl_valid_problem.Effective_date = formatDateString(anl_valid_problem.Effective_date)
	anl_valid_problem.Expiry_date = formatDateString(anl_valid_problem.Expiry_date)
	anl_valid_problem.Row_effective_date = formatDateString(anl_valid_problem.Row_effective_date)
	anl_valid_problem.Row_expiry_date = formatDateString(anl_valid_problem.Row_expiry_date)






	rows, err := stmt.Exec(anl_valid_problem.Method_set_id, anl_valid_problem.Method_id, anl_valid_problem.Problem_obs_no, anl_valid_problem.Accuracy_type, anl_valid_problem.Active_ind, anl_valid_problem.Confidence_type, anl_valid_problem.Effective_date, anl_valid_problem.Expiry_date, anl_valid_problem.Ppdm_guid, anl_valid_problem.Problem_desc, anl_valid_problem.Problem_result, anl_valid_problem.Problem_severity, anl_valid_problem.Problem_type, anl_valid_problem.Remark, anl_valid_problem.Resolution_method, anl_valid_problem.Source, anl_valid_problem.Substance_id, anl_valid_problem.Row_changed_by, anl_valid_problem.Row_changed_date, anl_valid_problem.Row_created_by, anl_valid_problem.Row_created_date, anl_valid_problem.Row_effective_date, anl_valid_problem.Row_expiry_date, anl_valid_problem.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlValidProblem(c *fiber.Ctx) error {
	var anl_valid_problem dto.Anl_valid_problem

	setDefaults(&anl_valid_problem)

	if err := json.Unmarshal(c.Body(), &anl_valid_problem); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_valid_problem.Method_set_id = id

    if anl_valid_problem.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_valid_problem set method_id = :1, problem_obs_no = :2, accuracy_type = :3, active_ind = :4, confidence_type = :5, effective_date = :6, expiry_date = :7, ppdm_guid = :8, problem_desc = :9, problem_result = :10, problem_severity = :11, problem_type = :12, remark = :13, resolution_method = :14, source = :15, substance_id = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where method_set_id = :24")
	if err != nil {
		return err
	}

	anl_valid_problem.Row_changed_date = formatDate(anl_valid_problem.Row_changed_date)
	anl_valid_problem.Effective_date = formatDateString(anl_valid_problem.Effective_date)
	anl_valid_problem.Expiry_date = formatDateString(anl_valid_problem.Expiry_date)
	anl_valid_problem.Row_effective_date = formatDateString(anl_valid_problem.Row_effective_date)
	anl_valid_problem.Row_expiry_date = formatDateString(anl_valid_problem.Row_expiry_date)






	rows, err := stmt.Exec(anl_valid_problem.Method_id, anl_valid_problem.Problem_obs_no, anl_valid_problem.Accuracy_type, anl_valid_problem.Active_ind, anl_valid_problem.Confidence_type, anl_valid_problem.Effective_date, anl_valid_problem.Expiry_date, anl_valid_problem.Ppdm_guid, anl_valid_problem.Problem_desc, anl_valid_problem.Problem_result, anl_valid_problem.Problem_severity, anl_valid_problem.Problem_type, anl_valid_problem.Remark, anl_valid_problem.Resolution_method, anl_valid_problem.Source, anl_valid_problem.Substance_id, anl_valid_problem.Row_changed_by, anl_valid_problem.Row_changed_date, anl_valid_problem.Row_created_by, anl_valid_problem.Row_effective_date, anl_valid_problem.Row_expiry_date, anl_valid_problem.Row_quality, anl_valid_problem.Method_set_id)
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

func PatchAnlValidProblem(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_valid_problem set "
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
	query += " where method_set_id = :id"

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

func DeleteAnlValidProblem(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_valid_problem dto.Anl_valid_problem
	anl_valid_problem.Method_set_id = id

	stmt, err := db.Prepare("delete from anl_valid_problem where method_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_valid_problem.Method_set_id)
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


