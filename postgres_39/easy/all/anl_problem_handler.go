package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlProblem(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_problem")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_problem

	for rows.Next() {
		var anl_problem dto.Anl_problem
		if err := rows.Scan(&anl_problem.Analysis_id, &anl_problem.Anl_source, &anl_problem.Problem_obs_no, &anl_problem.Active_ind, &anl_problem.Activity_set_id, &anl_problem.Anl_problem_ind, &anl_problem.Ba_problem_ind, &anl_problem.Confidence_type, &anl_problem.Effective_date, &anl_problem.Equip_problem_ind, &anl_problem.Expiry_date, &anl_problem.Method_id, &anl_problem.Method_problem_ind, &anl_problem.Percent_of_sample, &anl_problem.Ppdm_guid, &anl_problem.Problem_desc, &anl_problem.Problem_result, &anl_problem.Problem_result_desc, &anl_problem.Problem_severity, &anl_problem.Problem_severity_desc, &anl_problem.Problem_type, &anl_problem.Referenced_column_name, &anl_problem.Referenced_ppdm_guid, &anl_problem.Referenced_system_id, &anl_problem.Referenced_table_name, &anl_problem.Remark, &anl_problem.Resolution_method, &anl_problem.Resolution_method_desc, &anl_problem.Resolution_step_seq_no, &anl_problem.Resolved_by_ba_id, &anl_problem.Sample_problem_ind, &anl_problem.Source, &anl_problem.Step_seq_no, &anl_problem.Valid_problem_obs_no, &anl_problem.Row_changed_by, &anl_problem.Row_changed_date, &anl_problem.Row_created_by, &anl_problem.Row_created_date, &anl_problem.Row_effective_date, &anl_problem.Row_expiry_date, &anl_problem.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_problem to result
		result = append(result, anl_problem)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlProblem(c *fiber.Ctx) error {
	var anl_problem dto.Anl_problem

	setDefaults(&anl_problem)

	if err := json.Unmarshal(c.Body(), &anl_problem); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_problem values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41)")
	if err != nil {
		return err
	}
	anl_problem.Row_created_date = formatDate(anl_problem.Row_created_date)
	anl_problem.Row_changed_date = formatDate(anl_problem.Row_changed_date)
	anl_problem.Effective_date = formatDateString(anl_problem.Effective_date)
	anl_problem.Expiry_date = formatDateString(anl_problem.Expiry_date)
	anl_problem.Row_effective_date = formatDateString(anl_problem.Row_effective_date)
	anl_problem.Row_expiry_date = formatDateString(anl_problem.Row_expiry_date)






	rows, err := stmt.Exec(anl_problem.Analysis_id, anl_problem.Anl_source, anl_problem.Problem_obs_no, anl_problem.Active_ind, anl_problem.Activity_set_id, anl_problem.Anl_problem_ind, anl_problem.Ba_problem_ind, anl_problem.Confidence_type, anl_problem.Effective_date, anl_problem.Equip_problem_ind, anl_problem.Expiry_date, anl_problem.Method_id, anl_problem.Method_problem_ind, anl_problem.Percent_of_sample, anl_problem.Ppdm_guid, anl_problem.Problem_desc, anl_problem.Problem_result, anl_problem.Problem_result_desc, anl_problem.Problem_severity, anl_problem.Problem_severity_desc, anl_problem.Problem_type, anl_problem.Referenced_column_name, anl_problem.Referenced_ppdm_guid, anl_problem.Referenced_system_id, anl_problem.Referenced_table_name, anl_problem.Remark, anl_problem.Resolution_method, anl_problem.Resolution_method_desc, anl_problem.Resolution_step_seq_no, anl_problem.Resolved_by_ba_id, anl_problem.Sample_problem_ind, anl_problem.Source, anl_problem.Step_seq_no, anl_problem.Valid_problem_obs_no, anl_problem.Row_changed_by, anl_problem.Row_changed_date, anl_problem.Row_created_by, anl_problem.Row_created_date, anl_problem.Row_effective_date, anl_problem.Row_expiry_date, anl_problem.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlProblem(c *fiber.Ctx) error {
	var anl_problem dto.Anl_problem

	setDefaults(&anl_problem)

	if err := json.Unmarshal(c.Body(), &anl_problem); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_problem.Analysis_id = id

    if anl_problem.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_problem set anl_source = :1, problem_obs_no = :2, active_ind = :3, activity_set_id = :4, anl_problem_ind = :5, ba_problem_ind = :6, confidence_type = :7, effective_date = :8, equip_problem_ind = :9, expiry_date = :10, method_id = :11, method_problem_ind = :12, percent_of_sample = :13, ppdm_guid = :14, problem_desc = :15, problem_result = :16, problem_result_desc = :17, problem_severity = :18, problem_severity_desc = :19, problem_type = :20, referenced_column_name = :21, referenced_ppdm_guid = :22, referenced_system_id = :23, referenced_table_name = :24, remark = :25, resolution_method = :26, resolution_method_desc = :27, resolution_step_seq_no = :28, resolved_by_ba_id = :29, sample_problem_ind = :30, source = :31, step_seq_no = :32, valid_problem_obs_no = :33, row_changed_by = :34, row_changed_date = :35, row_created_by = :36, row_effective_date = :37, row_expiry_date = :38, row_quality = :39 where analysis_id = :41")
	if err != nil {
		return err
	}

	anl_problem.Row_changed_date = formatDate(anl_problem.Row_changed_date)
	anl_problem.Effective_date = formatDateString(anl_problem.Effective_date)
	anl_problem.Expiry_date = formatDateString(anl_problem.Expiry_date)
	anl_problem.Row_effective_date = formatDateString(anl_problem.Row_effective_date)
	anl_problem.Row_expiry_date = formatDateString(anl_problem.Row_expiry_date)






	rows, err := stmt.Exec(anl_problem.Anl_source, anl_problem.Problem_obs_no, anl_problem.Active_ind, anl_problem.Activity_set_id, anl_problem.Anl_problem_ind, anl_problem.Ba_problem_ind, anl_problem.Confidence_type, anl_problem.Effective_date, anl_problem.Equip_problem_ind, anl_problem.Expiry_date, anl_problem.Method_id, anl_problem.Method_problem_ind, anl_problem.Percent_of_sample, anl_problem.Ppdm_guid, anl_problem.Problem_desc, anl_problem.Problem_result, anl_problem.Problem_result_desc, anl_problem.Problem_severity, anl_problem.Problem_severity_desc, anl_problem.Problem_type, anl_problem.Referenced_column_name, anl_problem.Referenced_ppdm_guid, anl_problem.Referenced_system_id, anl_problem.Referenced_table_name, anl_problem.Remark, anl_problem.Resolution_method, anl_problem.Resolution_method_desc, anl_problem.Resolution_step_seq_no, anl_problem.Resolved_by_ba_id, anl_problem.Sample_problem_ind, anl_problem.Source, anl_problem.Step_seq_no, anl_problem.Valid_problem_obs_no, anl_problem.Row_changed_by, anl_problem.Row_changed_date, anl_problem.Row_created_by, anl_problem.Row_effective_date, anl_problem.Row_expiry_date, anl_problem.Row_quality, anl_problem.Analysis_id)
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

func PatchAnlProblem(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_problem set "
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
	query += " where analysis_id = :id"

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

func DeleteAnlProblem(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_problem dto.Anl_problem
	anl_problem.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_problem where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_problem.Analysis_id)
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


