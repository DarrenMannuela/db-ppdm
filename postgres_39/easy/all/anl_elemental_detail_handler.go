package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlElementalDetail(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_elemental_detail")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_elemental_detail

	for rows.Next() {
		var anl_elemental_detail dto.Anl_elemental_detail
		if err := rows.Scan(&anl_elemental_detail.Analysis_id, &anl_elemental_detail.Anl_source, &anl_elemental_detail.Detail_anl_obs_no, &anl_elemental_detail.Active_ind, &anl_elemental_detail.Analysis_value, &anl_elemental_detail.Analysis_value_ouom, &anl_elemental_detail.Analysis_value_type, &anl_elemental_detail.Analysis_value_uom, &anl_elemental_detail.Average_value, &anl_elemental_detail.Average_value_ouom, &anl_elemental_detail.Average_value_uom, &anl_elemental_detail.Calculated_ind, &anl_elemental_detail.Calculate_method_id, &anl_elemental_detail.Effective_date, &anl_elemental_detail.Expiry_date, &anl_elemental_detail.Max_date, &anl_elemental_detail.Max_value, &anl_elemental_detail.Max_value_ouom, &anl_elemental_detail.Max_value_uom, &anl_elemental_detail.Measurement_type, &anl_elemental_detail.Min_date, &anl_elemental_detail.Min_value, &anl_elemental_detail.Min_value_ouom, &anl_elemental_detail.Min_value_uom, &anl_elemental_detail.Ppdm_guid, &anl_elemental_detail.Preferred_ind, &anl_elemental_detail.Remark, &anl_elemental_detail.Reported_ind, &anl_elemental_detail.Source, &anl_elemental_detail.Step_seq_no, &anl_elemental_detail.Substance_id, &anl_elemental_detail.Valid_value_ind, &anl_elemental_detail.Valid_value_remark, &anl_elemental_detail.Value_code, &anl_elemental_detail.Value_desc, &anl_elemental_detail.Value_quality, &anl_elemental_detail.Row_changed_by, &anl_elemental_detail.Row_changed_date, &anl_elemental_detail.Row_created_by, &anl_elemental_detail.Row_created_date, &anl_elemental_detail.Row_effective_date, &anl_elemental_detail.Row_expiry_date, &anl_elemental_detail.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_elemental_detail to result
		result = append(result, anl_elemental_detail)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlElementalDetail(c *fiber.Ctx) error {
	var anl_elemental_detail dto.Anl_elemental_detail

	setDefaults(&anl_elemental_detail)

	if err := json.Unmarshal(c.Body(), &anl_elemental_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_elemental_detail values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43)")
	if err != nil {
		return err
	}
	anl_elemental_detail.Row_created_date = formatDate(anl_elemental_detail.Row_created_date)
	anl_elemental_detail.Row_changed_date = formatDate(anl_elemental_detail.Row_changed_date)
	anl_elemental_detail.Effective_date = formatDateString(anl_elemental_detail.Effective_date)
	anl_elemental_detail.Expiry_date = formatDateString(anl_elemental_detail.Expiry_date)
	anl_elemental_detail.Max_date = formatDateString(anl_elemental_detail.Max_date)
	anl_elemental_detail.Min_date = formatDateString(anl_elemental_detail.Min_date)
	anl_elemental_detail.Row_effective_date = formatDateString(anl_elemental_detail.Row_effective_date)
	anl_elemental_detail.Row_expiry_date = formatDateString(anl_elemental_detail.Row_expiry_date)






	rows, err := stmt.Exec(anl_elemental_detail.Analysis_id, anl_elemental_detail.Anl_source, anl_elemental_detail.Detail_anl_obs_no, anl_elemental_detail.Active_ind, anl_elemental_detail.Analysis_value, anl_elemental_detail.Analysis_value_ouom, anl_elemental_detail.Analysis_value_type, anl_elemental_detail.Analysis_value_uom, anl_elemental_detail.Average_value, anl_elemental_detail.Average_value_ouom, anl_elemental_detail.Average_value_uom, anl_elemental_detail.Calculated_ind, anl_elemental_detail.Calculate_method_id, anl_elemental_detail.Effective_date, anl_elemental_detail.Expiry_date, anl_elemental_detail.Max_date, anl_elemental_detail.Max_value, anl_elemental_detail.Max_value_ouom, anl_elemental_detail.Max_value_uom, anl_elemental_detail.Measurement_type, anl_elemental_detail.Min_date, anl_elemental_detail.Min_value, anl_elemental_detail.Min_value_ouom, anl_elemental_detail.Min_value_uom, anl_elemental_detail.Ppdm_guid, anl_elemental_detail.Preferred_ind, anl_elemental_detail.Remark, anl_elemental_detail.Reported_ind, anl_elemental_detail.Source, anl_elemental_detail.Step_seq_no, anl_elemental_detail.Substance_id, anl_elemental_detail.Valid_value_ind, anl_elemental_detail.Valid_value_remark, anl_elemental_detail.Value_code, anl_elemental_detail.Value_desc, anl_elemental_detail.Value_quality, anl_elemental_detail.Row_changed_by, anl_elemental_detail.Row_changed_date, anl_elemental_detail.Row_created_by, anl_elemental_detail.Row_created_date, anl_elemental_detail.Row_effective_date, anl_elemental_detail.Row_expiry_date, anl_elemental_detail.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlElementalDetail(c *fiber.Ctx) error {
	var anl_elemental_detail dto.Anl_elemental_detail

	setDefaults(&anl_elemental_detail)

	if err := json.Unmarshal(c.Body(), &anl_elemental_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_elemental_detail.Analysis_id = id

    if anl_elemental_detail.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_elemental_detail set anl_source = :1, detail_anl_obs_no = :2, active_ind = :3, analysis_value = :4, analysis_value_ouom = :5, analysis_value_type = :6, analysis_value_uom = :7, average_value = :8, average_value_ouom = :9, average_value_uom = :10, calculated_ind = :11, calculate_method_id = :12, effective_date = :13, expiry_date = :14, max_date = :15, max_value = :16, max_value_ouom = :17, max_value_uom = :18, measurement_type = :19, min_date = :20, min_value = :21, min_value_ouom = :22, min_value_uom = :23, ppdm_guid = :24, preferred_ind = :25, remark = :26, reported_ind = :27, source = :28, step_seq_no = :29, substance_id = :30, valid_value_ind = :31, valid_value_remark = :32, value_code = :33, value_desc = :34, value_quality = :35, row_changed_by = :36, row_changed_date = :37, row_created_by = :38, row_effective_date = :39, row_expiry_date = :40, row_quality = :41 where analysis_id = :43")
	if err != nil {
		return err
	}

	anl_elemental_detail.Row_changed_date = formatDate(anl_elemental_detail.Row_changed_date)
	anl_elemental_detail.Effective_date = formatDateString(anl_elemental_detail.Effective_date)
	anl_elemental_detail.Expiry_date = formatDateString(anl_elemental_detail.Expiry_date)
	anl_elemental_detail.Max_date = formatDateString(anl_elemental_detail.Max_date)
	anl_elemental_detail.Min_date = formatDateString(anl_elemental_detail.Min_date)
	anl_elemental_detail.Row_effective_date = formatDateString(anl_elemental_detail.Row_effective_date)
	anl_elemental_detail.Row_expiry_date = formatDateString(anl_elemental_detail.Row_expiry_date)






	rows, err := stmt.Exec(anl_elemental_detail.Anl_source, anl_elemental_detail.Detail_anl_obs_no, anl_elemental_detail.Active_ind, anl_elemental_detail.Analysis_value, anl_elemental_detail.Analysis_value_ouom, anl_elemental_detail.Analysis_value_type, anl_elemental_detail.Analysis_value_uom, anl_elemental_detail.Average_value, anl_elemental_detail.Average_value_ouom, anl_elemental_detail.Average_value_uom, anl_elemental_detail.Calculated_ind, anl_elemental_detail.Calculate_method_id, anl_elemental_detail.Effective_date, anl_elemental_detail.Expiry_date, anl_elemental_detail.Max_date, anl_elemental_detail.Max_value, anl_elemental_detail.Max_value_ouom, anl_elemental_detail.Max_value_uom, anl_elemental_detail.Measurement_type, anl_elemental_detail.Min_date, anl_elemental_detail.Min_value, anl_elemental_detail.Min_value_ouom, anl_elemental_detail.Min_value_uom, anl_elemental_detail.Ppdm_guid, anl_elemental_detail.Preferred_ind, anl_elemental_detail.Remark, anl_elemental_detail.Reported_ind, anl_elemental_detail.Source, anl_elemental_detail.Step_seq_no, anl_elemental_detail.Substance_id, anl_elemental_detail.Valid_value_ind, anl_elemental_detail.Valid_value_remark, anl_elemental_detail.Value_code, anl_elemental_detail.Value_desc, anl_elemental_detail.Value_quality, anl_elemental_detail.Row_changed_by, anl_elemental_detail.Row_changed_date, anl_elemental_detail.Row_created_by, anl_elemental_detail.Row_effective_date, anl_elemental_detail.Row_expiry_date, anl_elemental_detail.Row_quality, anl_elemental_detail.Analysis_id)
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

func PatchAnlElementalDetail(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_elemental_detail set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "max_date" || key == "min_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteAnlElementalDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_elemental_detail dto.Anl_elemental_detail
	anl_elemental_detail.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_elemental_detail where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_elemental_detail.Analysis_id)
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


