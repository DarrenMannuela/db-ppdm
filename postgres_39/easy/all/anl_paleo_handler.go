package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlPaleo(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_paleo")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_paleo

	for rows.Next() {
		var anl_paleo dto.Anl_paleo
		if err := rows.Scan(&anl_paleo.Analysis_id, &anl_paleo.Anl_source, &anl_paleo.Palynology_obs_no, &anl_paleo.Active_ind, &anl_paleo.Amount_type, &anl_paleo.Calculated_ind, &anl_paleo.Calculate_method_id, &anl_paleo.Effective_date, &anl_paleo.Expiry_date, &anl_paleo.Ppdm_guid, &anl_paleo.Preferred_ind, &anl_paleo.Problem_ind, &anl_paleo.Remark, &anl_paleo.Reported_ind, &anl_paleo.Sample_total_value, &anl_paleo.Sample_total_value_uom, &anl_paleo.Source, &anl_paleo.Step_seq_no, &anl_paleo.Substance_id, &anl_paleo.Total_value, &anl_paleo.Total_value_ouom, &anl_paleo.Row_changed_by, &anl_paleo.Row_changed_date, &anl_paleo.Row_created_by, &anl_paleo.Row_created_date, &anl_paleo.Row_effective_date, &anl_paleo.Row_expiry_date, &anl_paleo.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_paleo to result
		result = append(result, anl_paleo)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlPaleo(c *fiber.Ctx) error {
	var anl_paleo dto.Anl_paleo

	setDefaults(&anl_paleo)

	if err := json.Unmarshal(c.Body(), &anl_paleo); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_paleo values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28)")
	if err != nil {
		return err
	}
	anl_paleo.Row_created_date = formatDate(anl_paleo.Row_created_date)
	anl_paleo.Row_changed_date = formatDate(anl_paleo.Row_changed_date)
	anl_paleo.Effective_date = formatDateString(anl_paleo.Effective_date)
	anl_paleo.Expiry_date = formatDateString(anl_paleo.Expiry_date)
	anl_paleo.Row_effective_date = formatDateString(anl_paleo.Row_effective_date)
	anl_paleo.Row_expiry_date = formatDateString(anl_paleo.Row_expiry_date)






	rows, err := stmt.Exec(anl_paleo.Analysis_id, anl_paleo.Anl_source, anl_paleo.Palynology_obs_no, anl_paleo.Active_ind, anl_paleo.Amount_type, anl_paleo.Calculated_ind, anl_paleo.Calculate_method_id, anl_paleo.Effective_date, anl_paleo.Expiry_date, anl_paleo.Ppdm_guid, anl_paleo.Preferred_ind, anl_paleo.Problem_ind, anl_paleo.Remark, anl_paleo.Reported_ind, anl_paleo.Sample_total_value, anl_paleo.Sample_total_value_uom, anl_paleo.Source, anl_paleo.Step_seq_no, anl_paleo.Substance_id, anl_paleo.Total_value, anl_paleo.Total_value_ouom, anl_paleo.Row_changed_by, anl_paleo.Row_changed_date, anl_paleo.Row_created_by, anl_paleo.Row_created_date, anl_paleo.Row_effective_date, anl_paleo.Row_expiry_date, anl_paleo.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlPaleo(c *fiber.Ctx) error {
	var anl_paleo dto.Anl_paleo

	setDefaults(&anl_paleo)

	if err := json.Unmarshal(c.Body(), &anl_paleo); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_paleo.Analysis_id = id

    if anl_paleo.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_paleo set anl_source = :1, palynology_obs_no = :2, active_ind = :3, amount_type = :4, calculated_ind = :5, calculate_method_id = :6, effective_date = :7, expiry_date = :8, ppdm_guid = :9, preferred_ind = :10, problem_ind = :11, remark = :12, reported_ind = :13, sample_total_value = :14, sample_total_value_uom = :15, source = :16, step_seq_no = :17, substance_id = :18, total_value = :19, total_value_ouom = :20, row_changed_by = :21, row_changed_date = :22, row_created_by = :23, row_effective_date = :24, row_expiry_date = :25, row_quality = :26 where analysis_id = :28")
	if err != nil {
		return err
	}

	anl_paleo.Row_changed_date = formatDate(anl_paleo.Row_changed_date)
	anl_paleo.Effective_date = formatDateString(anl_paleo.Effective_date)
	anl_paleo.Expiry_date = formatDateString(anl_paleo.Expiry_date)
	anl_paleo.Row_effective_date = formatDateString(anl_paleo.Row_effective_date)
	anl_paleo.Row_expiry_date = formatDateString(anl_paleo.Row_expiry_date)






	rows, err := stmt.Exec(anl_paleo.Anl_source, anl_paleo.Palynology_obs_no, anl_paleo.Active_ind, anl_paleo.Amount_type, anl_paleo.Calculated_ind, anl_paleo.Calculate_method_id, anl_paleo.Effective_date, anl_paleo.Expiry_date, anl_paleo.Ppdm_guid, anl_paleo.Preferred_ind, anl_paleo.Problem_ind, anl_paleo.Remark, anl_paleo.Reported_ind, anl_paleo.Sample_total_value, anl_paleo.Sample_total_value_uom, anl_paleo.Source, anl_paleo.Step_seq_no, anl_paleo.Substance_id, anl_paleo.Total_value, anl_paleo.Total_value_ouom, anl_paleo.Row_changed_by, anl_paleo.Row_changed_date, anl_paleo.Row_created_by, anl_paleo.Row_effective_date, anl_paleo.Row_expiry_date, anl_paleo.Row_quality, anl_paleo.Analysis_id)
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

func PatchAnlPaleo(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_paleo set "
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

func DeleteAnlPaleo(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_paleo dto.Anl_paleo
	anl_paleo.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_paleo where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_paleo.Analysis_id)
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


