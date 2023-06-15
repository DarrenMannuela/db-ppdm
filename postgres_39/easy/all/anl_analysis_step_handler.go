package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlAnalysisStep(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_analysis_step")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_analysis_step

	for rows.Next() {
		var anl_analysis_step dto.Anl_analysis_step
		if err := rows.Scan(&anl_analysis_step.Analysis_id, &anl_analysis_step.Anl_source, &anl_analysis_step.Step_seq_no, &anl_analysis_step.Active_ind, &anl_analysis_step.Analysis_phase, &anl_analysis_step.Anl_date, &anl_analysis_step.Complete_date, &anl_analysis_step.Effective_date, &anl_analysis_step.End_date, &anl_analysis_step.Expiry_date, &anl_analysis_step.Final_volume, &anl_analysis_step.Final_volume_ouom, &anl_analysis_step.Final_volume_percent, &anl_analysis_step.Final_weight, &anl_analysis_step.Final_weight_ouom, &anl_analysis_step.Method_id, &anl_analysis_step.Method_set_id, &anl_analysis_step.Ppdm_guid, &anl_analysis_step.Problem_ind, &anl_analysis_step.Received_date, &anl_analysis_step.Recovered_percent, &anl_analysis_step.Remark, &anl_analysis_step.Reported_date, &anl_analysis_step.Results_received_date, &anl_analysis_step.Results_received_ind, &anl_analysis_step.Sample_fraction_type, &anl_analysis_step.Sample_quality, &anl_analysis_step.Sample_quality_desc, &anl_analysis_step.Source, &anl_analysis_step.Start_date, &anl_analysis_step.Step_completed_ind, &anl_analysis_step.Step_desc, &anl_analysis_step.Step_quality_desc, &anl_analysis_step.Step_quality_type, &anl_analysis_step.Step_requested_ind, &anl_analysis_step.Row_changed_by, &anl_analysis_step.Row_changed_date, &anl_analysis_step.Row_created_by, &anl_analysis_step.Row_created_date, &anl_analysis_step.Row_effective_date, &anl_analysis_step.Row_expiry_date, &anl_analysis_step.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_analysis_step to result
		result = append(result, anl_analysis_step)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlAnalysisStep(c *fiber.Ctx) error {
	var anl_analysis_step dto.Anl_analysis_step

	setDefaults(&anl_analysis_step)

	if err := json.Unmarshal(c.Body(), &anl_analysis_step); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_analysis_step values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42)")
	if err != nil {
		return err
	}
	anl_analysis_step.Row_created_date = formatDate(anl_analysis_step.Row_created_date)
	anl_analysis_step.Row_changed_date = formatDate(anl_analysis_step.Row_changed_date)
	anl_analysis_step.Anl_date = formatDateString(anl_analysis_step.Anl_date)
	anl_analysis_step.Complete_date = formatDateString(anl_analysis_step.Complete_date)
	anl_analysis_step.Effective_date = formatDateString(anl_analysis_step.Effective_date)
	anl_analysis_step.End_date = formatDateString(anl_analysis_step.End_date)
	anl_analysis_step.Expiry_date = formatDateString(anl_analysis_step.Expiry_date)
	anl_analysis_step.Received_date = formatDateString(anl_analysis_step.Received_date)
	anl_analysis_step.Reported_date = formatDateString(anl_analysis_step.Reported_date)
	anl_analysis_step.Results_received_date = formatDateString(anl_analysis_step.Results_received_date)
	anl_analysis_step.Start_date = formatDateString(anl_analysis_step.Start_date)
	anl_analysis_step.Row_effective_date = formatDateString(anl_analysis_step.Row_effective_date)
	anl_analysis_step.Row_expiry_date = formatDateString(anl_analysis_step.Row_expiry_date)






	rows, err := stmt.Exec(anl_analysis_step.Analysis_id, anl_analysis_step.Anl_source, anl_analysis_step.Step_seq_no, anl_analysis_step.Active_ind, anl_analysis_step.Analysis_phase, anl_analysis_step.Anl_date, anl_analysis_step.Complete_date, anl_analysis_step.Effective_date, anl_analysis_step.End_date, anl_analysis_step.Expiry_date, anl_analysis_step.Final_volume, anl_analysis_step.Final_volume_ouom, anl_analysis_step.Final_volume_percent, anl_analysis_step.Final_weight, anl_analysis_step.Final_weight_ouom, anl_analysis_step.Method_id, anl_analysis_step.Method_set_id, anl_analysis_step.Ppdm_guid, anl_analysis_step.Problem_ind, anl_analysis_step.Received_date, anl_analysis_step.Recovered_percent, anl_analysis_step.Remark, anl_analysis_step.Reported_date, anl_analysis_step.Results_received_date, anl_analysis_step.Results_received_ind, anl_analysis_step.Sample_fraction_type, anl_analysis_step.Sample_quality, anl_analysis_step.Sample_quality_desc, anl_analysis_step.Source, anl_analysis_step.Start_date, anl_analysis_step.Step_completed_ind, anl_analysis_step.Step_desc, anl_analysis_step.Step_quality_desc, anl_analysis_step.Step_quality_type, anl_analysis_step.Step_requested_ind, anl_analysis_step.Row_changed_by, anl_analysis_step.Row_changed_date, anl_analysis_step.Row_created_by, anl_analysis_step.Row_created_date, anl_analysis_step.Row_effective_date, anl_analysis_step.Row_expiry_date, anl_analysis_step.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlAnalysisStep(c *fiber.Ctx) error {
	var anl_analysis_step dto.Anl_analysis_step

	setDefaults(&anl_analysis_step)

	if err := json.Unmarshal(c.Body(), &anl_analysis_step); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_analysis_step.Analysis_id = id

    if anl_analysis_step.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_analysis_step set anl_source = :1, step_seq_no = :2, active_ind = :3, analysis_phase = :4, anl_date = :5, complete_date = :6, effective_date = :7, end_date = :8, expiry_date = :9, final_volume = :10, final_volume_ouom = :11, final_volume_percent = :12, final_weight = :13, final_weight_ouom = :14, method_id = :15, method_set_id = :16, ppdm_guid = :17, problem_ind = :18, received_date = :19, recovered_percent = :20, remark = :21, reported_date = :22, results_received_date = :23, results_received_ind = :24, sample_fraction_type = :25, sample_quality = :26, sample_quality_desc = :27, source = :28, start_date = :29, step_completed_ind = :30, step_desc = :31, step_quality_desc = :32, step_quality_type = :33, step_requested_ind = :34, row_changed_by = :35, row_changed_date = :36, row_created_by = :37, row_effective_date = :38, row_expiry_date = :39, row_quality = :40 where analysis_id = :42")
	if err != nil {
		return err
	}

	anl_analysis_step.Row_changed_date = formatDate(anl_analysis_step.Row_changed_date)
	anl_analysis_step.Anl_date = formatDateString(anl_analysis_step.Anl_date)
	anl_analysis_step.Complete_date = formatDateString(anl_analysis_step.Complete_date)
	anl_analysis_step.Effective_date = formatDateString(anl_analysis_step.Effective_date)
	anl_analysis_step.End_date = formatDateString(anl_analysis_step.End_date)
	anl_analysis_step.Expiry_date = formatDateString(anl_analysis_step.Expiry_date)
	anl_analysis_step.Received_date = formatDateString(anl_analysis_step.Received_date)
	anl_analysis_step.Reported_date = formatDateString(anl_analysis_step.Reported_date)
	anl_analysis_step.Results_received_date = formatDateString(anl_analysis_step.Results_received_date)
	anl_analysis_step.Start_date = formatDateString(anl_analysis_step.Start_date)
	anl_analysis_step.Row_effective_date = formatDateString(anl_analysis_step.Row_effective_date)
	anl_analysis_step.Row_expiry_date = formatDateString(anl_analysis_step.Row_expiry_date)






	rows, err := stmt.Exec(anl_analysis_step.Anl_source, anl_analysis_step.Step_seq_no, anl_analysis_step.Active_ind, anl_analysis_step.Analysis_phase, anl_analysis_step.Anl_date, anl_analysis_step.Complete_date, anl_analysis_step.Effective_date, anl_analysis_step.End_date, anl_analysis_step.Expiry_date, anl_analysis_step.Final_volume, anl_analysis_step.Final_volume_ouom, anl_analysis_step.Final_volume_percent, anl_analysis_step.Final_weight, anl_analysis_step.Final_weight_ouom, anl_analysis_step.Method_id, anl_analysis_step.Method_set_id, anl_analysis_step.Ppdm_guid, anl_analysis_step.Problem_ind, anl_analysis_step.Received_date, anl_analysis_step.Recovered_percent, anl_analysis_step.Remark, anl_analysis_step.Reported_date, anl_analysis_step.Results_received_date, anl_analysis_step.Results_received_ind, anl_analysis_step.Sample_fraction_type, anl_analysis_step.Sample_quality, anl_analysis_step.Sample_quality_desc, anl_analysis_step.Source, anl_analysis_step.Start_date, anl_analysis_step.Step_completed_ind, anl_analysis_step.Step_desc, anl_analysis_step.Step_quality_desc, anl_analysis_step.Step_quality_type, anl_analysis_step.Step_requested_ind, anl_analysis_step.Row_changed_by, anl_analysis_step.Row_changed_date, anl_analysis_step.Row_created_by, anl_analysis_step.Row_effective_date, anl_analysis_step.Row_expiry_date, anl_analysis_step.Row_quality, anl_analysis_step.Analysis_id)
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

func PatchAnlAnalysisStep(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_analysis_step set "
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
		} else if key == "anl_date" || key == "complete_date" || key == "effective_date" || key == "end_date" || key == "expiry_date" || key == "received_date" || key == "reported_date" || key == "results_received_date" || key == "start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteAnlAnalysisStep(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_analysis_step dto.Anl_analysis_step
	anl_analysis_step.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_analysis_step where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_analysis_step.Analysis_id)
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


