package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisProcPlanStep(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_proc_plan_step")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_proc_plan_step

	for rows.Next() {
		var seis_proc_plan_step dto.Seis_proc_plan_step
		if err := rows.Scan(&seis_proc_plan_step.Seis_proc_plan_id, &seis_proc_plan_step.Proc_plan_step_id, &seis_proc_plan_step.Active_ind, &seis_proc_plan_step.Description, &seis_proc_plan_step.Effective_date, &seis_proc_plan_step.Expiry_date, &seis_proc_plan_step.Ppdm_guid, &seis_proc_plan_step.Remark, &seis_proc_plan_step.Source, &seis_proc_plan_step.Step_name, &seis_proc_plan_step.Step_seq_no, &seis_proc_plan_step.Step_type, &seis_proc_plan_step.Row_changed_by, &seis_proc_plan_step.Row_changed_date, &seis_proc_plan_step.Row_created_by, &seis_proc_plan_step.Row_created_date, &seis_proc_plan_step.Row_effective_date, &seis_proc_plan_step.Row_expiry_date, &seis_proc_plan_step.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_proc_plan_step to result
		result = append(result, seis_proc_plan_step)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisProcPlanStep(c *fiber.Ctx) error {
	var seis_proc_plan_step dto.Seis_proc_plan_step

	setDefaults(&seis_proc_plan_step)

	if err := json.Unmarshal(c.Body(), &seis_proc_plan_step); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_proc_plan_step values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	seis_proc_plan_step.Row_created_date = formatDate(seis_proc_plan_step.Row_created_date)
	seis_proc_plan_step.Row_changed_date = formatDate(seis_proc_plan_step.Row_changed_date)
	seis_proc_plan_step.Effective_date = formatDateString(seis_proc_plan_step.Effective_date)
	seis_proc_plan_step.Expiry_date = formatDateString(seis_proc_plan_step.Expiry_date)
	seis_proc_plan_step.Row_effective_date = formatDateString(seis_proc_plan_step.Row_effective_date)
	seis_proc_plan_step.Row_expiry_date = formatDateString(seis_proc_plan_step.Row_expiry_date)






	rows, err := stmt.Exec(seis_proc_plan_step.Seis_proc_plan_id, seis_proc_plan_step.Proc_plan_step_id, seis_proc_plan_step.Active_ind, seis_proc_plan_step.Description, seis_proc_plan_step.Effective_date, seis_proc_plan_step.Expiry_date, seis_proc_plan_step.Ppdm_guid, seis_proc_plan_step.Remark, seis_proc_plan_step.Source, seis_proc_plan_step.Step_name, seis_proc_plan_step.Step_seq_no, seis_proc_plan_step.Step_type, seis_proc_plan_step.Row_changed_by, seis_proc_plan_step.Row_changed_date, seis_proc_plan_step.Row_created_by, seis_proc_plan_step.Row_created_date, seis_proc_plan_step.Row_effective_date, seis_proc_plan_step.Row_expiry_date, seis_proc_plan_step.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisProcPlanStep(c *fiber.Ctx) error {
	var seis_proc_plan_step dto.Seis_proc_plan_step

	setDefaults(&seis_proc_plan_step)

	if err := json.Unmarshal(c.Body(), &seis_proc_plan_step); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_proc_plan_step.Seis_proc_plan_id = id

    if seis_proc_plan_step.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_proc_plan_step set proc_plan_step_id = :1, active_ind = :2, description = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, step_name = :9, step_seq_no = :10, step_type = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where seis_proc_plan_id = :19")
	if err != nil {
		return err
	}

	seis_proc_plan_step.Row_changed_date = formatDate(seis_proc_plan_step.Row_changed_date)
	seis_proc_plan_step.Effective_date = formatDateString(seis_proc_plan_step.Effective_date)
	seis_proc_plan_step.Expiry_date = formatDateString(seis_proc_plan_step.Expiry_date)
	seis_proc_plan_step.Row_effective_date = formatDateString(seis_proc_plan_step.Row_effective_date)
	seis_proc_plan_step.Row_expiry_date = formatDateString(seis_proc_plan_step.Row_expiry_date)






	rows, err := stmt.Exec(seis_proc_plan_step.Proc_plan_step_id, seis_proc_plan_step.Active_ind, seis_proc_plan_step.Description, seis_proc_plan_step.Effective_date, seis_proc_plan_step.Expiry_date, seis_proc_plan_step.Ppdm_guid, seis_proc_plan_step.Remark, seis_proc_plan_step.Source, seis_proc_plan_step.Step_name, seis_proc_plan_step.Step_seq_no, seis_proc_plan_step.Step_type, seis_proc_plan_step.Row_changed_by, seis_proc_plan_step.Row_changed_date, seis_proc_plan_step.Row_created_by, seis_proc_plan_step.Row_effective_date, seis_proc_plan_step.Row_expiry_date, seis_proc_plan_step.Row_quality, seis_proc_plan_step.Seis_proc_plan_id)
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

func PatchSeisProcPlanStep(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_proc_plan_step set "
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
	query += " where seis_proc_plan_id = :id"

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

func DeleteSeisProcPlanStep(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_proc_plan_step dto.Seis_proc_plan_step
	seis_proc_plan_step.Seis_proc_plan_id = id

	stmt, err := db.Prepare("delete from seis_proc_plan_step where seis_proc_plan_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_proc_plan_step.Seis_proc_plan_id)
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


