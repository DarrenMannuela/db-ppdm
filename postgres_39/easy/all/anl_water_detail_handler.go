package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlWaterDetail(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_water_detail")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_water_detail

	for rows.Next() {
		var anl_water_detail dto.Anl_water_detail
		if err := rows.Scan(&anl_water_detail.Analysis_id, &anl_water_detail.Anl_source, &anl_water_detail.Detail_obs_no, &anl_water_detail.Active_ind, &anl_water_detail.Analysis_press, &anl_water_detail.Analysis_press_ouom, &anl_water_detail.Analysis_press_uom, &anl_water_detail.Analysis_temp, &anl_water_detail.Analysis_temp_ouom, &anl_water_detail.Analysis_temp_uom, &anl_water_detail.Analysis_value, &anl_water_detail.Analysis_value_ouom, &anl_water_detail.Analysis_value_type, &anl_water_detail.Analysis_value_uom, &anl_water_detail.Anl_value_remark, &anl_water_detail.Calculated_value_ind, &anl_water_detail.Calculate_method_id, &anl_water_detail.Date_format_desc, &anl_water_detail.Effective_date, &anl_water_detail.Expiry_date, &anl_water_detail.Max_value, &anl_water_detail.Max_value_ouom, &anl_water_detail.Max_value_uom, &anl_water_detail.Measured_value_ind, &anl_water_detail.Measurement_type, &anl_water_detail.Min_value, &anl_water_detail.Min_value_ouom, &anl_water_detail.Min_value_uom, &anl_water_detail.Ppdm_guid, &anl_water_detail.Preferred_ind, &anl_water_detail.Problem_ind, &anl_water_detail.Remark, &anl_water_detail.Reported_value_ind, &anl_water_detail.Source, &anl_water_detail.Step_seq_no, &anl_water_detail.Substance_id, &anl_water_detail.Water_property, &anl_water_detail.Water_property_code, &anl_water_detail.Row_changed_by, &anl_water_detail.Row_changed_date, &anl_water_detail.Row_created_by, &anl_water_detail.Row_created_date, &anl_water_detail.Row_effective_date, &anl_water_detail.Row_expiry_date, &anl_water_detail.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_water_detail to result
		result = append(result, anl_water_detail)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlWaterDetail(c *fiber.Ctx) error {
	var anl_water_detail dto.Anl_water_detail

	setDefaults(&anl_water_detail)

	if err := json.Unmarshal(c.Body(), &anl_water_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_water_detail values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45)")
	if err != nil {
		return err
	}
	anl_water_detail.Row_created_date = formatDate(anl_water_detail.Row_created_date)
	anl_water_detail.Row_changed_date = formatDate(anl_water_detail.Row_changed_date)
	anl_water_detail.Date_format_desc = formatDateString(anl_water_detail.Date_format_desc)
	anl_water_detail.Effective_date = formatDateString(anl_water_detail.Effective_date)
	anl_water_detail.Expiry_date = formatDateString(anl_water_detail.Expiry_date)
	anl_water_detail.Row_effective_date = formatDateString(anl_water_detail.Row_effective_date)
	anl_water_detail.Row_expiry_date = formatDateString(anl_water_detail.Row_expiry_date)






	rows, err := stmt.Exec(anl_water_detail.Analysis_id, anl_water_detail.Anl_source, anl_water_detail.Detail_obs_no, anl_water_detail.Active_ind, anl_water_detail.Analysis_press, anl_water_detail.Analysis_press_ouom, anl_water_detail.Analysis_press_uom, anl_water_detail.Analysis_temp, anl_water_detail.Analysis_temp_ouom, anl_water_detail.Analysis_temp_uom, anl_water_detail.Analysis_value, anl_water_detail.Analysis_value_ouom, anl_water_detail.Analysis_value_type, anl_water_detail.Analysis_value_uom, anl_water_detail.Anl_value_remark, anl_water_detail.Calculated_value_ind, anl_water_detail.Calculate_method_id, anl_water_detail.Date_format_desc, anl_water_detail.Effective_date, anl_water_detail.Expiry_date, anl_water_detail.Max_value, anl_water_detail.Max_value_ouom, anl_water_detail.Max_value_uom, anl_water_detail.Measured_value_ind, anl_water_detail.Measurement_type, anl_water_detail.Min_value, anl_water_detail.Min_value_ouom, anl_water_detail.Min_value_uom, anl_water_detail.Ppdm_guid, anl_water_detail.Preferred_ind, anl_water_detail.Problem_ind, anl_water_detail.Remark, anl_water_detail.Reported_value_ind, anl_water_detail.Source, anl_water_detail.Step_seq_no, anl_water_detail.Substance_id, anl_water_detail.Water_property, anl_water_detail.Water_property_code, anl_water_detail.Row_changed_by, anl_water_detail.Row_changed_date, anl_water_detail.Row_created_by, anl_water_detail.Row_created_date, anl_water_detail.Row_effective_date, anl_water_detail.Row_expiry_date, anl_water_detail.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlWaterDetail(c *fiber.Ctx) error {
	var anl_water_detail dto.Anl_water_detail

	setDefaults(&anl_water_detail)

	if err := json.Unmarshal(c.Body(), &anl_water_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_water_detail.Analysis_id = id

    if anl_water_detail.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_water_detail set anl_source = :1, detail_obs_no = :2, active_ind = :3, analysis_press = :4, analysis_press_ouom = :5, analysis_press_uom = :6, analysis_temp = :7, analysis_temp_ouom = :8, analysis_temp_uom = :9, analysis_value = :10, analysis_value_ouom = :11, analysis_value_type = :12, analysis_value_uom = :13, anl_value_remark = :14, calculated_value_ind = :15, calculate_method_id = :16, date_format_desc = :17, effective_date = :18, expiry_date = :19, max_value = :20, max_value_ouom = :21, max_value_uom = :22, measured_value_ind = :23, measurement_type = :24, min_value = :25, min_value_ouom = :26, min_value_uom = :27, ppdm_guid = :28, preferred_ind = :29, problem_ind = :30, remark = :31, reported_value_ind = :32, source = :33, step_seq_no = :34, substance_id = :35, water_property = :36, water_property_code = :37, row_changed_by = :38, row_changed_date = :39, row_created_by = :40, row_effective_date = :41, row_expiry_date = :42, row_quality = :43 where analysis_id = :45")
	if err != nil {
		return err
	}

	anl_water_detail.Row_changed_date = formatDate(anl_water_detail.Row_changed_date)
	anl_water_detail.Date_format_desc = formatDateString(anl_water_detail.Date_format_desc)
	anl_water_detail.Effective_date = formatDateString(anl_water_detail.Effective_date)
	anl_water_detail.Expiry_date = formatDateString(anl_water_detail.Expiry_date)
	anl_water_detail.Row_effective_date = formatDateString(anl_water_detail.Row_effective_date)
	anl_water_detail.Row_expiry_date = formatDateString(anl_water_detail.Row_expiry_date)






	rows, err := stmt.Exec(anl_water_detail.Anl_source, anl_water_detail.Detail_obs_no, anl_water_detail.Active_ind, anl_water_detail.Analysis_press, anl_water_detail.Analysis_press_ouom, anl_water_detail.Analysis_press_uom, anl_water_detail.Analysis_temp, anl_water_detail.Analysis_temp_ouom, anl_water_detail.Analysis_temp_uom, anl_water_detail.Analysis_value, anl_water_detail.Analysis_value_ouom, anl_water_detail.Analysis_value_type, anl_water_detail.Analysis_value_uom, anl_water_detail.Anl_value_remark, anl_water_detail.Calculated_value_ind, anl_water_detail.Calculate_method_id, anl_water_detail.Date_format_desc, anl_water_detail.Effective_date, anl_water_detail.Expiry_date, anl_water_detail.Max_value, anl_water_detail.Max_value_ouom, anl_water_detail.Max_value_uom, anl_water_detail.Measured_value_ind, anl_water_detail.Measurement_type, anl_water_detail.Min_value, anl_water_detail.Min_value_ouom, anl_water_detail.Min_value_uom, anl_water_detail.Ppdm_guid, anl_water_detail.Preferred_ind, anl_water_detail.Problem_ind, anl_water_detail.Remark, anl_water_detail.Reported_value_ind, anl_water_detail.Source, anl_water_detail.Step_seq_no, anl_water_detail.Substance_id, anl_water_detail.Water_property, anl_water_detail.Water_property_code, anl_water_detail.Row_changed_by, anl_water_detail.Row_changed_date, anl_water_detail.Row_created_by, anl_water_detail.Row_effective_date, anl_water_detail.Row_expiry_date, anl_water_detail.Row_quality, anl_water_detail.Analysis_id)
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

func PatchAnlWaterDetail(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_water_detail set "
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
		} else if key == "date_format_desc" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteAnlWaterDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_water_detail dto.Anl_water_detail
	anl_water_detail.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_water_detail where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_water_detail.Analysis_id)
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


