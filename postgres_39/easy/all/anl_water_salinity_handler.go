package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlWaterSalinity(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_water_salinity")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_water_salinity

	for rows.Next() {
		var anl_water_salinity dto.Anl_water_salinity
		if err := rows.Scan(&anl_water_salinity.Analysis_id, &anl_water_salinity.Anl_source, &anl_water_salinity.Water_salinity_obs_no, &anl_water_salinity.Active_ind, &anl_water_salinity.Calculated_ind, &anl_water_salinity.Calculate_method_id, &anl_water_salinity.Effective_date, &anl_water_salinity.Expiry_date, &anl_water_salinity.Ignition_ind, &anl_water_salinity.Measured_ind, &anl_water_salinity.Measured_temp, &anl_water_salinity.Measured_temp_ouom, &anl_water_salinity.Ppdm_guid, &anl_water_salinity.Preferred_ind, &anl_water_salinity.Problem_ind, &anl_water_salinity.Remark, &anl_water_salinity.Reported_ind, &anl_water_salinity.Salinity, &anl_water_salinity.Salinity_ouom, &anl_water_salinity.Salinity_type, &anl_water_salinity.Salinity_uom, &anl_water_salinity.Source, &anl_water_salinity.Step_seq_no, &anl_water_salinity.Row_changed_by, &anl_water_salinity.Row_changed_date, &anl_water_salinity.Row_created_by, &anl_water_salinity.Row_created_date, &anl_water_salinity.Row_effective_date, &anl_water_salinity.Row_expiry_date, &anl_water_salinity.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_water_salinity to result
		result = append(result, anl_water_salinity)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlWaterSalinity(c *fiber.Ctx) error {
	var anl_water_salinity dto.Anl_water_salinity

	setDefaults(&anl_water_salinity)

	if err := json.Unmarshal(c.Body(), &anl_water_salinity); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_water_salinity values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30)")
	if err != nil {
		return err
	}
	anl_water_salinity.Row_created_date = formatDate(anl_water_salinity.Row_created_date)
	anl_water_salinity.Row_changed_date = formatDate(anl_water_salinity.Row_changed_date)
	anl_water_salinity.Effective_date = formatDateString(anl_water_salinity.Effective_date)
	anl_water_salinity.Expiry_date = formatDateString(anl_water_salinity.Expiry_date)
	anl_water_salinity.Row_effective_date = formatDateString(anl_water_salinity.Row_effective_date)
	anl_water_salinity.Row_expiry_date = formatDateString(anl_water_salinity.Row_expiry_date)






	rows, err := stmt.Exec(anl_water_salinity.Analysis_id, anl_water_salinity.Anl_source, anl_water_salinity.Water_salinity_obs_no, anl_water_salinity.Active_ind, anl_water_salinity.Calculated_ind, anl_water_salinity.Calculate_method_id, anl_water_salinity.Effective_date, anl_water_salinity.Expiry_date, anl_water_salinity.Ignition_ind, anl_water_salinity.Measured_ind, anl_water_salinity.Measured_temp, anl_water_salinity.Measured_temp_ouom, anl_water_salinity.Ppdm_guid, anl_water_salinity.Preferred_ind, anl_water_salinity.Problem_ind, anl_water_salinity.Remark, anl_water_salinity.Reported_ind, anl_water_salinity.Salinity, anl_water_salinity.Salinity_ouom, anl_water_salinity.Salinity_type, anl_water_salinity.Salinity_uom, anl_water_salinity.Source, anl_water_salinity.Step_seq_no, anl_water_salinity.Row_changed_by, anl_water_salinity.Row_changed_date, anl_water_salinity.Row_created_by, anl_water_salinity.Row_created_date, anl_water_salinity.Row_effective_date, anl_water_salinity.Row_expiry_date, anl_water_salinity.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlWaterSalinity(c *fiber.Ctx) error {
	var anl_water_salinity dto.Anl_water_salinity

	setDefaults(&anl_water_salinity)

	if err := json.Unmarshal(c.Body(), &anl_water_salinity); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_water_salinity.Analysis_id = id

    if anl_water_salinity.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_water_salinity set anl_source = :1, water_salinity_obs_no = :2, active_ind = :3, calculated_ind = :4, calculate_method_id = :5, effective_date = :6, expiry_date = :7, ignition_ind = :8, measured_ind = :9, measured_temp = :10, measured_temp_ouom = :11, ppdm_guid = :12, preferred_ind = :13, problem_ind = :14, remark = :15, reported_ind = :16, salinity = :17, salinity_ouom = :18, salinity_type = :19, salinity_uom = :20, source = :21, step_seq_no = :22, row_changed_by = :23, row_changed_date = :24, row_created_by = :25, row_effective_date = :26, row_expiry_date = :27, row_quality = :28 where analysis_id = :30")
	if err != nil {
		return err
	}

	anl_water_salinity.Row_changed_date = formatDate(anl_water_salinity.Row_changed_date)
	anl_water_salinity.Effective_date = formatDateString(anl_water_salinity.Effective_date)
	anl_water_salinity.Expiry_date = formatDateString(anl_water_salinity.Expiry_date)
	anl_water_salinity.Row_effective_date = formatDateString(anl_water_salinity.Row_effective_date)
	anl_water_salinity.Row_expiry_date = formatDateString(anl_water_salinity.Row_expiry_date)






	rows, err := stmt.Exec(anl_water_salinity.Anl_source, anl_water_salinity.Water_salinity_obs_no, anl_water_salinity.Active_ind, anl_water_salinity.Calculated_ind, anl_water_salinity.Calculate_method_id, anl_water_salinity.Effective_date, anl_water_salinity.Expiry_date, anl_water_salinity.Ignition_ind, anl_water_salinity.Measured_ind, anl_water_salinity.Measured_temp, anl_water_salinity.Measured_temp_ouom, anl_water_salinity.Ppdm_guid, anl_water_salinity.Preferred_ind, anl_water_salinity.Problem_ind, anl_water_salinity.Remark, anl_water_salinity.Reported_ind, anl_water_salinity.Salinity, anl_water_salinity.Salinity_ouom, anl_water_salinity.Salinity_type, anl_water_salinity.Salinity_uom, anl_water_salinity.Source, anl_water_salinity.Step_seq_no, anl_water_salinity.Row_changed_by, anl_water_salinity.Row_changed_date, anl_water_salinity.Row_created_by, anl_water_salinity.Row_effective_date, anl_water_salinity.Row_expiry_date, anl_water_salinity.Row_quality, anl_water_salinity.Analysis_id)
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

func PatchAnlWaterSalinity(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_water_salinity set "
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

func DeleteAnlWaterSalinity(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_water_salinity dto.Anl_water_salinity
	anl_water_salinity.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_water_salinity where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_water_salinity.Analysis_id)
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


