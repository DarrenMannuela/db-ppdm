package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlGasAnalysis(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_gas_analysis")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_gas_analysis

	for rows.Next() {
		var anl_gas_analysis dto.Anl_gas_analysis
		if err := rows.Scan(&anl_gas_analysis.Analysis_id, &anl_gas_analysis.Anl_source, &anl_gas_analysis.Gas_anl_obs_no, &anl_gas_analysis.Active_ind, &anl_gas_analysis.Effective_date, &anl_gas_analysis.Expiry_date, &anl_gas_analysis.Fluid_type, &anl_gas_analysis.Gas_gravity, &anl_gas_analysis.Gas_gravity_ouom, &anl_gas_analysis.Ppdm_guid, &anl_gas_analysis.Preferred_ind, &anl_gas_analysis.Problem_ind, &anl_gas_analysis.Pseudo_critical_press, &anl_gas_analysis.Pseudo_critical_press_ouom, &anl_gas_analysis.Pseudo_critical_temp, &anl_gas_analysis.Pseudo_critical_temp_ouom, &anl_gas_analysis.Remark, &anl_gas_analysis.Source, &anl_gas_analysis.Step_seq_no, &anl_gas_analysis.Row_changed_by, &anl_gas_analysis.Row_changed_date, &anl_gas_analysis.Row_created_by, &anl_gas_analysis.Row_created_date, &anl_gas_analysis.Row_effective_date, &anl_gas_analysis.Row_expiry_date, &anl_gas_analysis.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_gas_analysis to result
		result = append(result, anl_gas_analysis)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlGasAnalysis(c *fiber.Ctx) error {
	var anl_gas_analysis dto.Anl_gas_analysis

	setDefaults(&anl_gas_analysis)

	if err := json.Unmarshal(c.Body(), &anl_gas_analysis); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_gas_analysis values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	anl_gas_analysis.Row_created_date = formatDate(anl_gas_analysis.Row_created_date)
	anl_gas_analysis.Row_changed_date = formatDate(anl_gas_analysis.Row_changed_date)
	anl_gas_analysis.Effective_date = formatDateString(anl_gas_analysis.Effective_date)
	anl_gas_analysis.Expiry_date = formatDateString(anl_gas_analysis.Expiry_date)
	anl_gas_analysis.Row_effective_date = formatDateString(anl_gas_analysis.Row_effective_date)
	anl_gas_analysis.Row_expiry_date = formatDateString(anl_gas_analysis.Row_expiry_date)






	rows, err := stmt.Exec(anl_gas_analysis.Analysis_id, anl_gas_analysis.Anl_source, anl_gas_analysis.Gas_anl_obs_no, anl_gas_analysis.Active_ind, anl_gas_analysis.Effective_date, anl_gas_analysis.Expiry_date, anl_gas_analysis.Fluid_type, anl_gas_analysis.Gas_gravity, anl_gas_analysis.Gas_gravity_ouom, anl_gas_analysis.Ppdm_guid, anl_gas_analysis.Preferred_ind, anl_gas_analysis.Problem_ind, anl_gas_analysis.Pseudo_critical_press, anl_gas_analysis.Pseudo_critical_press_ouom, anl_gas_analysis.Pseudo_critical_temp, anl_gas_analysis.Pseudo_critical_temp_ouom, anl_gas_analysis.Remark, anl_gas_analysis.Source, anl_gas_analysis.Step_seq_no, anl_gas_analysis.Row_changed_by, anl_gas_analysis.Row_changed_date, anl_gas_analysis.Row_created_by, anl_gas_analysis.Row_created_date, anl_gas_analysis.Row_effective_date, anl_gas_analysis.Row_expiry_date, anl_gas_analysis.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlGasAnalysis(c *fiber.Ctx) error {
	var anl_gas_analysis dto.Anl_gas_analysis

	setDefaults(&anl_gas_analysis)

	if err := json.Unmarshal(c.Body(), &anl_gas_analysis); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_gas_analysis.Analysis_id = id

    if anl_gas_analysis.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_gas_analysis set anl_source = :1, gas_anl_obs_no = :2, active_ind = :3, effective_date = :4, expiry_date = :5, fluid_type = :6, gas_gravity = :7, gas_gravity_ouom = :8, ppdm_guid = :9, preferred_ind = :10, problem_ind = :11, pseudo_critical_press = :12, pseudo_critical_press_ouom = :13, pseudo_critical_temp = :14, pseudo_critical_temp_ouom = :15, remark = :16, source = :17, step_seq_no = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where analysis_id = :26")
	if err != nil {
		return err
	}

	anl_gas_analysis.Row_changed_date = formatDate(anl_gas_analysis.Row_changed_date)
	anl_gas_analysis.Effective_date = formatDateString(anl_gas_analysis.Effective_date)
	anl_gas_analysis.Expiry_date = formatDateString(anl_gas_analysis.Expiry_date)
	anl_gas_analysis.Row_effective_date = formatDateString(anl_gas_analysis.Row_effective_date)
	anl_gas_analysis.Row_expiry_date = formatDateString(anl_gas_analysis.Row_expiry_date)






	rows, err := stmt.Exec(anl_gas_analysis.Anl_source, anl_gas_analysis.Gas_anl_obs_no, anl_gas_analysis.Active_ind, anl_gas_analysis.Effective_date, anl_gas_analysis.Expiry_date, anl_gas_analysis.Fluid_type, anl_gas_analysis.Gas_gravity, anl_gas_analysis.Gas_gravity_ouom, anl_gas_analysis.Ppdm_guid, anl_gas_analysis.Preferred_ind, anl_gas_analysis.Problem_ind, anl_gas_analysis.Pseudo_critical_press, anl_gas_analysis.Pseudo_critical_press_ouom, anl_gas_analysis.Pseudo_critical_temp, anl_gas_analysis.Pseudo_critical_temp_ouom, anl_gas_analysis.Remark, anl_gas_analysis.Source, anl_gas_analysis.Step_seq_no, anl_gas_analysis.Row_changed_by, anl_gas_analysis.Row_changed_date, anl_gas_analysis.Row_created_by, anl_gas_analysis.Row_effective_date, anl_gas_analysis.Row_expiry_date, anl_gas_analysis.Row_quality, anl_gas_analysis.Analysis_id)
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

func PatchAnlGasAnalysis(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_gas_analysis set "
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

func DeleteAnlGasAnalysis(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_gas_analysis dto.Anl_gas_analysis
	anl_gas_analysis.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_gas_analysis where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_gas_analysis.Analysis_id)
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


