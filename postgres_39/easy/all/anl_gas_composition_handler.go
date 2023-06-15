package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlGasComposition(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_gas_composition")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_gas_composition

	for rows.Next() {
		var anl_gas_composition dto.Anl_gas_composition
		if err := rows.Scan(&anl_gas_composition.Analysis_id, &anl_gas_composition.Anl_source, &anl_gas_composition.Gas_anl_composition_obs_no, &anl_gas_composition.Active_ind, &anl_gas_composition.Calculated_ind, &anl_gas_composition.Calculate_method_id, &anl_gas_composition.Effective_date, &anl_gas_composition.Expiry_date, &anl_gas_composition.Not_present_ind, &anl_gas_composition.Ppdm_guid, &anl_gas_composition.Problem_ind, &anl_gas_composition.Remark, &anl_gas_composition.Reported_ind, &anl_gas_composition.Source, &anl_gas_composition.Step_seq_no, &anl_gas_composition.Substance_amount, &anl_gas_composition.Substance_amount_ouom, &anl_gas_composition.Substance_amount_uom, &anl_gas_composition.Substance_id, &anl_gas_composition.Substance_percent, &anl_gas_composition.Trace_ind, &anl_gas_composition.Row_changed_by, &anl_gas_composition.Row_changed_date, &anl_gas_composition.Row_created_by, &anl_gas_composition.Row_created_date, &anl_gas_composition.Row_effective_date, &anl_gas_composition.Row_expiry_date, &anl_gas_composition.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_gas_composition to result
		result = append(result, anl_gas_composition)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlGasComposition(c *fiber.Ctx) error {
	var anl_gas_composition dto.Anl_gas_composition

	setDefaults(&anl_gas_composition)

	if err := json.Unmarshal(c.Body(), &anl_gas_composition); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_gas_composition values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28)")
	if err != nil {
		return err
	}
	anl_gas_composition.Row_created_date = formatDate(anl_gas_composition.Row_created_date)
	anl_gas_composition.Row_changed_date = formatDate(anl_gas_composition.Row_changed_date)
	anl_gas_composition.Effective_date = formatDateString(anl_gas_composition.Effective_date)
	anl_gas_composition.Expiry_date = formatDateString(anl_gas_composition.Expiry_date)
	anl_gas_composition.Row_effective_date = formatDateString(anl_gas_composition.Row_effective_date)
	anl_gas_composition.Row_expiry_date = formatDateString(anl_gas_composition.Row_expiry_date)






	rows, err := stmt.Exec(anl_gas_composition.Analysis_id, anl_gas_composition.Anl_source, anl_gas_composition.Gas_anl_composition_obs_no, anl_gas_composition.Active_ind, anl_gas_composition.Calculated_ind, anl_gas_composition.Calculate_method_id, anl_gas_composition.Effective_date, anl_gas_composition.Expiry_date, anl_gas_composition.Not_present_ind, anl_gas_composition.Ppdm_guid, anl_gas_composition.Problem_ind, anl_gas_composition.Remark, anl_gas_composition.Reported_ind, anl_gas_composition.Source, anl_gas_composition.Step_seq_no, anl_gas_composition.Substance_amount, anl_gas_composition.Substance_amount_ouom, anl_gas_composition.Substance_amount_uom, anl_gas_composition.Substance_id, anl_gas_composition.Substance_percent, anl_gas_composition.Trace_ind, anl_gas_composition.Row_changed_by, anl_gas_composition.Row_changed_date, anl_gas_composition.Row_created_by, anl_gas_composition.Row_created_date, anl_gas_composition.Row_effective_date, anl_gas_composition.Row_expiry_date, anl_gas_composition.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlGasComposition(c *fiber.Ctx) error {
	var anl_gas_composition dto.Anl_gas_composition

	setDefaults(&anl_gas_composition)

	if err := json.Unmarshal(c.Body(), &anl_gas_composition); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_gas_composition.Analysis_id = id

    if anl_gas_composition.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_gas_composition set anl_source = :1, gas_anl_composition_obs_no = :2, active_ind = :3, calculated_ind = :4, calculate_method_id = :5, effective_date = :6, expiry_date = :7, not_present_ind = :8, ppdm_guid = :9, problem_ind = :10, remark = :11, reported_ind = :12, source = :13, step_seq_no = :14, substance_amount = :15, substance_amount_ouom = :16, substance_amount_uom = :17, substance_id = :18, substance_percent = :19, trace_ind = :20, row_changed_by = :21, row_changed_date = :22, row_created_by = :23, row_effective_date = :24, row_expiry_date = :25, row_quality = :26 where analysis_id = :28")
	if err != nil {
		return err
	}

	anl_gas_composition.Row_changed_date = formatDate(anl_gas_composition.Row_changed_date)
	anl_gas_composition.Effective_date = formatDateString(anl_gas_composition.Effective_date)
	anl_gas_composition.Expiry_date = formatDateString(anl_gas_composition.Expiry_date)
	anl_gas_composition.Row_effective_date = formatDateString(anl_gas_composition.Row_effective_date)
	anl_gas_composition.Row_expiry_date = formatDateString(anl_gas_composition.Row_expiry_date)






	rows, err := stmt.Exec(anl_gas_composition.Anl_source, anl_gas_composition.Gas_anl_composition_obs_no, anl_gas_composition.Active_ind, anl_gas_composition.Calculated_ind, anl_gas_composition.Calculate_method_id, anl_gas_composition.Effective_date, anl_gas_composition.Expiry_date, anl_gas_composition.Not_present_ind, anl_gas_composition.Ppdm_guid, anl_gas_composition.Problem_ind, anl_gas_composition.Remark, anl_gas_composition.Reported_ind, anl_gas_composition.Source, anl_gas_composition.Step_seq_no, anl_gas_composition.Substance_amount, anl_gas_composition.Substance_amount_ouom, anl_gas_composition.Substance_amount_uom, anl_gas_composition.Substance_id, anl_gas_composition.Substance_percent, anl_gas_composition.Trace_ind, anl_gas_composition.Row_changed_by, anl_gas_composition.Row_changed_date, anl_gas_composition.Row_created_by, anl_gas_composition.Row_effective_date, anl_gas_composition.Row_expiry_date, anl_gas_composition.Row_quality, anl_gas_composition.Analysis_id)
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

func PatchAnlGasComposition(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_gas_composition set "
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

func DeleteAnlGasComposition(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_gas_composition dto.Anl_gas_composition
	anl_gas_composition.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_gas_composition where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_gas_composition.Analysis_id)
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


