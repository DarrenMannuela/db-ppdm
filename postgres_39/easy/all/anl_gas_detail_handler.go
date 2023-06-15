package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlGasDetail(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_gas_detail")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_gas_detail

	for rows.Next() {
		var anl_gas_detail dto.Anl_gas_detail
		if err := rows.Scan(&anl_gas_detail.Analysis_id, &anl_gas_detail.Anl_source, &anl_gas_detail.Gas_anl_detail_obs_no, &anl_gas_detail.Active_ind, &anl_gas_detail.Analysis_property, &anl_gas_detail.Analysis_value, &anl_gas_detail.Analysis_value_code, &anl_gas_detail.Analysis_value_ouom, &anl_gas_detail.Analysis_value_uom, &anl_gas_detail.Anl_value_remark, &anl_gas_detail.Calculate_method_id, &anl_gas_detail.Date_format_desc, &anl_gas_detail.Effective_date, &anl_gas_detail.Expiry_date, &anl_gas_detail.Max_value, &anl_gas_detail.Max_value_ouom, &anl_gas_detail.Max_value_uom, &anl_gas_detail.Measurement_type, &anl_gas_detail.Min_value, &anl_gas_detail.Min_value_ouom, &anl_gas_detail.Min_value_uom, &anl_gas_detail.Ppdm_guid, &anl_gas_detail.Remark, &anl_gas_detail.Source, &anl_gas_detail.Step_seq_no, &anl_gas_detail.Substance_id, &anl_gas_detail.Row_changed_by, &anl_gas_detail.Row_changed_date, &anl_gas_detail.Row_created_by, &anl_gas_detail.Row_created_date, &anl_gas_detail.Row_effective_date, &anl_gas_detail.Row_expiry_date, &anl_gas_detail.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_gas_detail to result
		result = append(result, anl_gas_detail)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlGasDetail(c *fiber.Ctx) error {
	var anl_gas_detail dto.Anl_gas_detail

	setDefaults(&anl_gas_detail)

	if err := json.Unmarshal(c.Body(), &anl_gas_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_gas_detail values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33)")
	if err != nil {
		return err
	}
	anl_gas_detail.Row_created_date = formatDate(anl_gas_detail.Row_created_date)
	anl_gas_detail.Row_changed_date = formatDate(anl_gas_detail.Row_changed_date)
	anl_gas_detail.Date_format_desc = formatDateString(anl_gas_detail.Date_format_desc)
	anl_gas_detail.Effective_date = formatDateString(anl_gas_detail.Effective_date)
	anl_gas_detail.Expiry_date = formatDateString(anl_gas_detail.Expiry_date)
	anl_gas_detail.Row_effective_date = formatDateString(anl_gas_detail.Row_effective_date)
	anl_gas_detail.Row_expiry_date = formatDateString(anl_gas_detail.Row_expiry_date)






	rows, err := stmt.Exec(anl_gas_detail.Analysis_id, anl_gas_detail.Anl_source, anl_gas_detail.Gas_anl_detail_obs_no, anl_gas_detail.Active_ind, anl_gas_detail.Analysis_property, anl_gas_detail.Analysis_value, anl_gas_detail.Analysis_value_code, anl_gas_detail.Analysis_value_ouom, anl_gas_detail.Analysis_value_uom, anl_gas_detail.Anl_value_remark, anl_gas_detail.Calculate_method_id, anl_gas_detail.Date_format_desc, anl_gas_detail.Effective_date, anl_gas_detail.Expiry_date, anl_gas_detail.Max_value, anl_gas_detail.Max_value_ouom, anl_gas_detail.Max_value_uom, anl_gas_detail.Measurement_type, anl_gas_detail.Min_value, anl_gas_detail.Min_value_ouom, anl_gas_detail.Min_value_uom, anl_gas_detail.Ppdm_guid, anl_gas_detail.Remark, anl_gas_detail.Source, anl_gas_detail.Step_seq_no, anl_gas_detail.Substance_id, anl_gas_detail.Row_changed_by, anl_gas_detail.Row_changed_date, anl_gas_detail.Row_created_by, anl_gas_detail.Row_created_date, anl_gas_detail.Row_effective_date, anl_gas_detail.Row_expiry_date, anl_gas_detail.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlGasDetail(c *fiber.Ctx) error {
	var anl_gas_detail dto.Anl_gas_detail

	setDefaults(&anl_gas_detail)

	if err := json.Unmarshal(c.Body(), &anl_gas_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_gas_detail.Analysis_id = id

    if anl_gas_detail.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_gas_detail set anl_source = :1, gas_anl_detail_obs_no = :2, active_ind = :3, analysis_property = :4, analysis_value = :5, analysis_value_code = :6, analysis_value_ouom = :7, analysis_value_uom = :8, anl_value_remark = :9, calculate_method_id = :10, date_format_desc = :11, effective_date = :12, expiry_date = :13, max_value = :14, max_value_ouom = :15, max_value_uom = :16, measurement_type = :17, min_value = :18, min_value_ouom = :19, min_value_uom = :20, ppdm_guid = :21, remark = :22, source = :23, step_seq_no = :24, substance_id = :25, row_changed_by = :26, row_changed_date = :27, row_created_by = :28, row_effective_date = :29, row_expiry_date = :30, row_quality = :31 where analysis_id = :33")
	if err != nil {
		return err
	}

	anl_gas_detail.Row_changed_date = formatDate(anl_gas_detail.Row_changed_date)
	anl_gas_detail.Date_format_desc = formatDateString(anl_gas_detail.Date_format_desc)
	anl_gas_detail.Effective_date = formatDateString(anl_gas_detail.Effective_date)
	anl_gas_detail.Expiry_date = formatDateString(anl_gas_detail.Expiry_date)
	anl_gas_detail.Row_effective_date = formatDateString(anl_gas_detail.Row_effective_date)
	anl_gas_detail.Row_expiry_date = formatDateString(anl_gas_detail.Row_expiry_date)






	rows, err := stmt.Exec(anl_gas_detail.Anl_source, anl_gas_detail.Gas_anl_detail_obs_no, anl_gas_detail.Active_ind, anl_gas_detail.Analysis_property, anl_gas_detail.Analysis_value, anl_gas_detail.Analysis_value_code, anl_gas_detail.Analysis_value_ouom, anl_gas_detail.Analysis_value_uom, anl_gas_detail.Anl_value_remark, anl_gas_detail.Calculate_method_id, anl_gas_detail.Date_format_desc, anl_gas_detail.Effective_date, anl_gas_detail.Expiry_date, anl_gas_detail.Max_value, anl_gas_detail.Max_value_ouom, anl_gas_detail.Max_value_uom, anl_gas_detail.Measurement_type, anl_gas_detail.Min_value, anl_gas_detail.Min_value_ouom, anl_gas_detail.Min_value_uom, anl_gas_detail.Ppdm_guid, anl_gas_detail.Remark, anl_gas_detail.Source, anl_gas_detail.Step_seq_no, anl_gas_detail.Substance_id, anl_gas_detail.Row_changed_by, anl_gas_detail.Row_changed_date, anl_gas_detail.Row_created_by, anl_gas_detail.Row_effective_date, anl_gas_detail.Row_expiry_date, anl_gas_detail.Row_quality, anl_gas_detail.Analysis_id)
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

func PatchAnlGasDetail(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_gas_detail set "
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

func DeleteAnlGasDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_gas_detail dto.Anl_gas_detail
	anl_gas_detail.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_gas_detail where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_gas_detail.Analysis_id)
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


