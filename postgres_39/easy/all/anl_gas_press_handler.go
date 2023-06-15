package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlGasPress(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_gas_press")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_gas_press

	for rows.Next() {
		var anl_gas_press dto.Anl_gas_press
		if err := rows.Scan(&anl_gas_press.Analysis_id, &anl_gas_press.Anl_source, &anl_gas_press.Gas_anl_press_obs_no, &anl_gas_press.Active_ind, &anl_gas_press.Calculate_method_id, &anl_gas_press.Effective_date, &anl_gas_press.Expiry_date, &anl_gas_press.Gauge_press, &anl_gas_press.Gauge_press_ouom, &anl_gas_press.Gauge_temp, &anl_gas_press.Gauge_temp_ouom, &anl_gas_press.Measurement_location, &anl_gas_press.Ppdm_guid, &anl_gas_press.Problem_ind, &anl_gas_press.Remark, &anl_gas_press.Source, &anl_gas_press.Step_seq_no, &anl_gas_press.Row_changed_by, &anl_gas_press.Row_changed_date, &anl_gas_press.Row_created_by, &anl_gas_press.Row_created_date, &anl_gas_press.Row_effective_date, &anl_gas_press.Row_expiry_date, &anl_gas_press.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_gas_press to result
		result = append(result, anl_gas_press)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlGasPress(c *fiber.Ctx) error {
	var anl_gas_press dto.Anl_gas_press

	setDefaults(&anl_gas_press)

	if err := json.Unmarshal(c.Body(), &anl_gas_press); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_gas_press values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	anl_gas_press.Row_created_date = formatDate(anl_gas_press.Row_created_date)
	anl_gas_press.Row_changed_date = formatDate(anl_gas_press.Row_changed_date)
	anl_gas_press.Effective_date = formatDateString(anl_gas_press.Effective_date)
	anl_gas_press.Expiry_date = formatDateString(anl_gas_press.Expiry_date)
	anl_gas_press.Row_effective_date = formatDateString(anl_gas_press.Row_effective_date)
	anl_gas_press.Row_expiry_date = formatDateString(anl_gas_press.Row_expiry_date)






	rows, err := stmt.Exec(anl_gas_press.Analysis_id, anl_gas_press.Anl_source, anl_gas_press.Gas_anl_press_obs_no, anl_gas_press.Active_ind, anl_gas_press.Calculate_method_id, anl_gas_press.Effective_date, anl_gas_press.Expiry_date, anl_gas_press.Gauge_press, anl_gas_press.Gauge_press_ouom, anl_gas_press.Gauge_temp, anl_gas_press.Gauge_temp_ouom, anl_gas_press.Measurement_location, anl_gas_press.Ppdm_guid, anl_gas_press.Problem_ind, anl_gas_press.Remark, anl_gas_press.Source, anl_gas_press.Step_seq_no, anl_gas_press.Row_changed_by, anl_gas_press.Row_changed_date, anl_gas_press.Row_created_by, anl_gas_press.Row_created_date, anl_gas_press.Row_effective_date, anl_gas_press.Row_expiry_date, anl_gas_press.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlGasPress(c *fiber.Ctx) error {
	var anl_gas_press dto.Anl_gas_press

	setDefaults(&anl_gas_press)

	if err := json.Unmarshal(c.Body(), &anl_gas_press); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_gas_press.Analysis_id = id

    if anl_gas_press.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_gas_press set anl_source = :1, gas_anl_press_obs_no = :2, active_ind = :3, calculate_method_id = :4, effective_date = :5, expiry_date = :6, gauge_press = :7, gauge_press_ouom = :8, gauge_temp = :9, gauge_temp_ouom = :10, measurement_location = :11, ppdm_guid = :12, problem_ind = :13, remark = :14, source = :15, step_seq_no = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where analysis_id = :24")
	if err != nil {
		return err
	}

	anl_gas_press.Row_changed_date = formatDate(anl_gas_press.Row_changed_date)
	anl_gas_press.Effective_date = formatDateString(anl_gas_press.Effective_date)
	anl_gas_press.Expiry_date = formatDateString(anl_gas_press.Expiry_date)
	anl_gas_press.Row_effective_date = formatDateString(anl_gas_press.Row_effective_date)
	anl_gas_press.Row_expiry_date = formatDateString(anl_gas_press.Row_expiry_date)






	rows, err := stmt.Exec(anl_gas_press.Anl_source, anl_gas_press.Gas_anl_press_obs_no, anl_gas_press.Active_ind, anl_gas_press.Calculate_method_id, anl_gas_press.Effective_date, anl_gas_press.Expiry_date, anl_gas_press.Gauge_press, anl_gas_press.Gauge_press_ouom, anl_gas_press.Gauge_temp, anl_gas_press.Gauge_temp_ouom, anl_gas_press.Measurement_location, anl_gas_press.Ppdm_guid, anl_gas_press.Problem_ind, anl_gas_press.Remark, anl_gas_press.Source, anl_gas_press.Step_seq_no, anl_gas_press.Row_changed_by, anl_gas_press.Row_changed_date, anl_gas_press.Row_created_by, anl_gas_press.Row_effective_date, anl_gas_press.Row_expiry_date, anl_gas_press.Row_quality, anl_gas_press.Analysis_id)
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

func PatchAnlGasPress(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_gas_press set "
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

func DeleteAnlGasPress(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_gas_press dto.Anl_gas_press
	anl_gas_press.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_gas_press where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_gas_press.Analysis_id)
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


