package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlDetail(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_detail")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_detail

	for rows.Next() {
		var anl_detail dto.Anl_detail
		if err := rows.Scan(&anl_detail.Analysis_id, &anl_detail.Anl_source, &anl_detail.Detail_obs_no, &anl_detail.Active_ind, &anl_detail.Anl_detail_type, &anl_detail.Average_ratio_value, &anl_detail.Average_value, &anl_detail.Average_value_ouom, &anl_detail.Average_value_uom, &anl_detail.Calculated_ind, &anl_detail.Calculate_method_id, &anl_detail.Date_format_desc, &anl_detail.Effective_date, &anl_detail.Expiry_date, &anl_detail.Max_date, &anl_detail.Max_ratio_value, &anl_detail.Max_value, &anl_detail.Max_value_ouom, &anl_detail.Max_value_uom, &anl_detail.Measurement_type, &anl_detail.Min_date, &anl_detail.Min_ratio_value, &anl_detail.Min_value, &anl_detail.Min_value_ouom, &anl_detail.Min_value_uom, &anl_detail.Ppdm_guid, &anl_detail.Problem_ind, &anl_detail.Reference_value, &anl_detail.Reference_value_ouom, &anl_detail.Reference_value_type, &anl_detail.Reference_value_uom, &anl_detail.Remark, &anl_detail.Reported_ind, &anl_detail.Source, &anl_detail.Step_seq_no, &anl_detail.Substance_id, &anl_detail.Row_changed_by, &anl_detail.Row_changed_date, &anl_detail.Row_created_by, &anl_detail.Row_created_date, &anl_detail.Row_effective_date, &anl_detail.Row_expiry_date, &anl_detail.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_detail to result
		result = append(result, anl_detail)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlDetail(c *fiber.Ctx) error {
	var anl_detail dto.Anl_detail

	setDefaults(&anl_detail)

	if err := json.Unmarshal(c.Body(), &anl_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_detail values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43)")
	if err != nil {
		return err
	}
	anl_detail.Row_created_date = formatDate(anl_detail.Row_created_date)
	anl_detail.Row_changed_date = formatDate(anl_detail.Row_changed_date)
	anl_detail.Date_format_desc = formatDateString(anl_detail.Date_format_desc)
	anl_detail.Effective_date = formatDateString(anl_detail.Effective_date)
	anl_detail.Expiry_date = formatDateString(anl_detail.Expiry_date)
	anl_detail.Max_date = formatDateString(anl_detail.Max_date)
	anl_detail.Min_date = formatDateString(anl_detail.Min_date)
	anl_detail.Row_effective_date = formatDateString(anl_detail.Row_effective_date)
	anl_detail.Row_expiry_date = formatDateString(anl_detail.Row_expiry_date)






	rows, err := stmt.Exec(anl_detail.Analysis_id, anl_detail.Anl_source, anl_detail.Detail_obs_no, anl_detail.Active_ind, anl_detail.Anl_detail_type, anl_detail.Average_ratio_value, anl_detail.Average_value, anl_detail.Average_value_ouom, anl_detail.Average_value_uom, anl_detail.Calculated_ind, anl_detail.Calculate_method_id, anl_detail.Date_format_desc, anl_detail.Effective_date, anl_detail.Expiry_date, anl_detail.Max_date, anl_detail.Max_ratio_value, anl_detail.Max_value, anl_detail.Max_value_ouom, anl_detail.Max_value_uom, anl_detail.Measurement_type, anl_detail.Min_date, anl_detail.Min_ratio_value, anl_detail.Min_value, anl_detail.Min_value_ouom, anl_detail.Min_value_uom, anl_detail.Ppdm_guid, anl_detail.Problem_ind, anl_detail.Reference_value, anl_detail.Reference_value_ouom, anl_detail.Reference_value_type, anl_detail.Reference_value_uom, anl_detail.Remark, anl_detail.Reported_ind, anl_detail.Source, anl_detail.Step_seq_no, anl_detail.Substance_id, anl_detail.Row_changed_by, anl_detail.Row_changed_date, anl_detail.Row_created_by, anl_detail.Row_created_date, anl_detail.Row_effective_date, anl_detail.Row_expiry_date, anl_detail.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlDetail(c *fiber.Ctx) error {
	var anl_detail dto.Anl_detail

	setDefaults(&anl_detail)

	if err := json.Unmarshal(c.Body(), &anl_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_detail.Analysis_id = id

    if anl_detail.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_detail set anl_source = :1, detail_obs_no = :2, active_ind = :3, anl_detail_type = :4, average_ratio_value = :5, average_value = :6, average_value_ouom = :7, average_value_uom = :8, calculated_ind = :9, calculate_method_id = :10, date_format_desc = :11, effective_date = :12, expiry_date = :13, max_date = :14, max_ratio_value = :15, max_value = :16, max_value_ouom = :17, max_value_uom = :18, measurement_type = :19, min_date = :20, min_ratio_value = :21, min_value = :22, min_value_ouom = :23, min_value_uom = :24, ppdm_guid = :25, problem_ind = :26, reference_value = :27, reference_value_ouom = :28, reference_value_type = :29, reference_value_uom = :30, remark = :31, reported_ind = :32, source = :33, step_seq_no = :34, substance_id = :35, row_changed_by = :36, row_changed_date = :37, row_created_by = :38, row_effective_date = :39, row_expiry_date = :40, row_quality = :41 where analysis_id = :43")
	if err != nil {
		return err
	}

	anl_detail.Row_changed_date = formatDate(anl_detail.Row_changed_date)
	anl_detail.Date_format_desc = formatDateString(anl_detail.Date_format_desc)
	anl_detail.Effective_date = formatDateString(anl_detail.Effective_date)
	anl_detail.Expiry_date = formatDateString(anl_detail.Expiry_date)
	anl_detail.Max_date = formatDateString(anl_detail.Max_date)
	anl_detail.Min_date = formatDateString(anl_detail.Min_date)
	anl_detail.Row_effective_date = formatDateString(anl_detail.Row_effective_date)
	anl_detail.Row_expiry_date = formatDateString(anl_detail.Row_expiry_date)






	rows, err := stmt.Exec(anl_detail.Anl_source, anl_detail.Detail_obs_no, anl_detail.Active_ind, anl_detail.Anl_detail_type, anl_detail.Average_ratio_value, anl_detail.Average_value, anl_detail.Average_value_ouom, anl_detail.Average_value_uom, anl_detail.Calculated_ind, anl_detail.Calculate_method_id, anl_detail.Date_format_desc, anl_detail.Effective_date, anl_detail.Expiry_date, anl_detail.Max_date, anl_detail.Max_ratio_value, anl_detail.Max_value, anl_detail.Max_value_ouom, anl_detail.Max_value_uom, anl_detail.Measurement_type, anl_detail.Min_date, anl_detail.Min_ratio_value, anl_detail.Min_value, anl_detail.Min_value_ouom, anl_detail.Min_value_uom, anl_detail.Ppdm_guid, anl_detail.Problem_ind, anl_detail.Reference_value, anl_detail.Reference_value_ouom, anl_detail.Reference_value_type, anl_detail.Reference_value_uom, anl_detail.Remark, anl_detail.Reported_ind, anl_detail.Source, anl_detail.Step_seq_no, anl_detail.Substance_id, anl_detail.Row_changed_by, anl_detail.Row_changed_date, anl_detail.Row_created_by, anl_detail.Row_effective_date, anl_detail.Row_expiry_date, anl_detail.Row_quality, anl_detail.Analysis_id)
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

func PatchAnlDetail(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_detail set "
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
		} else if key == "date_format_desc" || key == "effective_date" || key == "expiry_date" || key == "max_date" || key == "min_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteAnlDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_detail dto.Anl_detail
	anl_detail.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_detail where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_detail.Analysis_id)
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


