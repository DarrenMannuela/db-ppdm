package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlLiquidChroDetail(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_liquid_chro_detail")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_liquid_chro_detail

	for rows.Next() {
		var anl_liquid_chro_detail dto.Anl_liquid_chro_detail
		if err := rows.Scan(&anl_liquid_chro_detail.Analysis_id, &anl_liquid_chro_detail.Anl_source, &anl_liquid_chro_detail.Detail_obs_no, &anl_liquid_chro_detail.Active_ind, &anl_liquid_chro_detail.Anl_value, &anl_liquid_chro_detail.Anl_value_ouom, &anl_liquid_chro_detail.Anl_value_remark, &anl_liquid_chro_detail.Anl_value_uom, &anl_liquid_chro_detail.Calculated_ind, &anl_liquid_chro_detail.Calculate_method_id, &anl_liquid_chro_detail.Chro_property_type, &anl_liquid_chro_detail.Effective_date, &anl_liquid_chro_detail.Error_bar_value, &anl_liquid_chro_detail.Error_bar_value_ouom, &anl_liquid_chro_detail.Error_bar_value_uom, &anl_liquid_chro_detail.Expiry_date, &anl_liquid_chro_detail.Measured_ind, &anl_liquid_chro_detail.Measurement_type, &anl_liquid_chro_detail.Ppdm_guid, &anl_liquid_chro_detail.Preferred_ind, &anl_liquid_chro_detail.Problem_ind, &anl_liquid_chro_detail.Quantif_additive_amt, &anl_liquid_chro_detail.Quantif_additive_amt_ouom, &anl_liquid_chro_detail.Quantif_additive_amt_uom, &anl_liquid_chro_detail.Quantif_additive_desc, &anl_liquid_chro_detail.Quantif_additive_type, &anl_liquid_chro_detail.Remark, &anl_liquid_chro_detail.Reported_ind, &anl_liquid_chro_detail.Source, &anl_liquid_chro_detail.Step_seq_no, &anl_liquid_chro_detail.Substance_id, &anl_liquid_chro_detail.Row_changed_by, &anl_liquid_chro_detail.Row_changed_date, &anl_liquid_chro_detail.Row_created_by, &anl_liquid_chro_detail.Row_created_date, &anl_liquid_chro_detail.Row_effective_date, &anl_liquid_chro_detail.Row_expiry_date, &anl_liquid_chro_detail.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_liquid_chro_detail to result
		result = append(result, anl_liquid_chro_detail)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlLiquidChroDetail(c *fiber.Ctx) error {
	var anl_liquid_chro_detail dto.Anl_liquid_chro_detail

	setDefaults(&anl_liquid_chro_detail)

	if err := json.Unmarshal(c.Body(), &anl_liquid_chro_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_liquid_chro_detail values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38)")
	if err != nil {
		return err
	}
	anl_liquid_chro_detail.Row_created_date = formatDate(anl_liquid_chro_detail.Row_created_date)
	anl_liquid_chro_detail.Row_changed_date = formatDate(anl_liquid_chro_detail.Row_changed_date)
	anl_liquid_chro_detail.Effective_date = formatDateString(anl_liquid_chro_detail.Effective_date)
	anl_liquid_chro_detail.Expiry_date = formatDateString(anl_liquid_chro_detail.Expiry_date)
	anl_liquid_chro_detail.Row_effective_date = formatDateString(anl_liquid_chro_detail.Row_effective_date)
	anl_liquid_chro_detail.Row_expiry_date = formatDateString(anl_liquid_chro_detail.Row_expiry_date)






	rows, err := stmt.Exec(anl_liquid_chro_detail.Analysis_id, anl_liquid_chro_detail.Anl_source, anl_liquid_chro_detail.Detail_obs_no, anl_liquid_chro_detail.Active_ind, anl_liquid_chro_detail.Anl_value, anl_liquid_chro_detail.Anl_value_ouom, anl_liquid_chro_detail.Anl_value_remark, anl_liquid_chro_detail.Anl_value_uom, anl_liquid_chro_detail.Calculated_ind, anl_liquid_chro_detail.Calculate_method_id, anl_liquid_chro_detail.Chro_property_type, anl_liquid_chro_detail.Effective_date, anl_liquid_chro_detail.Error_bar_value, anl_liquid_chro_detail.Error_bar_value_ouom, anl_liquid_chro_detail.Error_bar_value_uom, anl_liquid_chro_detail.Expiry_date, anl_liquid_chro_detail.Measured_ind, anl_liquid_chro_detail.Measurement_type, anl_liquid_chro_detail.Ppdm_guid, anl_liquid_chro_detail.Preferred_ind, anl_liquid_chro_detail.Problem_ind, anl_liquid_chro_detail.Quantif_additive_amt, anl_liquid_chro_detail.Quantif_additive_amt_ouom, anl_liquid_chro_detail.Quantif_additive_amt_uom, anl_liquid_chro_detail.Quantif_additive_desc, anl_liquid_chro_detail.Quantif_additive_type, anl_liquid_chro_detail.Remark, anl_liquid_chro_detail.Reported_ind, anl_liquid_chro_detail.Source, anl_liquid_chro_detail.Step_seq_no, anl_liquid_chro_detail.Substance_id, anl_liquid_chro_detail.Row_changed_by, anl_liquid_chro_detail.Row_changed_date, anl_liquid_chro_detail.Row_created_by, anl_liquid_chro_detail.Row_created_date, anl_liquid_chro_detail.Row_effective_date, anl_liquid_chro_detail.Row_expiry_date, anl_liquid_chro_detail.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlLiquidChroDetail(c *fiber.Ctx) error {
	var anl_liquid_chro_detail dto.Anl_liquid_chro_detail

	setDefaults(&anl_liquid_chro_detail)

	if err := json.Unmarshal(c.Body(), &anl_liquid_chro_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_liquid_chro_detail.Analysis_id = id

    if anl_liquid_chro_detail.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_liquid_chro_detail set anl_source = :1, detail_obs_no = :2, active_ind = :3, anl_value = :4, anl_value_ouom = :5, anl_value_remark = :6, anl_value_uom = :7, calculated_ind = :8, calculate_method_id = :9, chro_property_type = :10, effective_date = :11, error_bar_value = :12, error_bar_value_ouom = :13, error_bar_value_uom = :14, expiry_date = :15, measured_ind = :16, measurement_type = :17, ppdm_guid = :18, preferred_ind = :19, problem_ind = :20, quantif_additive_amt = :21, quantif_additive_amt_ouom = :22, quantif_additive_amt_uom = :23, quantif_additive_desc = :24, quantif_additive_type = :25, remark = :26, reported_ind = :27, source = :28, step_seq_no = :29, substance_id = :30, row_changed_by = :31, row_changed_date = :32, row_created_by = :33, row_effective_date = :34, row_expiry_date = :35, row_quality = :36 where analysis_id = :38")
	if err != nil {
		return err
	}

	anl_liquid_chro_detail.Row_changed_date = formatDate(anl_liquid_chro_detail.Row_changed_date)
	anl_liquid_chro_detail.Effective_date = formatDateString(anl_liquid_chro_detail.Effective_date)
	anl_liquid_chro_detail.Expiry_date = formatDateString(anl_liquid_chro_detail.Expiry_date)
	anl_liquid_chro_detail.Row_effective_date = formatDateString(anl_liquid_chro_detail.Row_effective_date)
	anl_liquid_chro_detail.Row_expiry_date = formatDateString(anl_liquid_chro_detail.Row_expiry_date)






	rows, err := stmt.Exec(anl_liquid_chro_detail.Anl_source, anl_liquid_chro_detail.Detail_obs_no, anl_liquid_chro_detail.Active_ind, anl_liquid_chro_detail.Anl_value, anl_liquid_chro_detail.Anl_value_ouom, anl_liquid_chro_detail.Anl_value_remark, anl_liquid_chro_detail.Anl_value_uom, anl_liquid_chro_detail.Calculated_ind, anl_liquid_chro_detail.Calculate_method_id, anl_liquid_chro_detail.Chro_property_type, anl_liquid_chro_detail.Effective_date, anl_liquid_chro_detail.Error_bar_value, anl_liquid_chro_detail.Error_bar_value_ouom, anl_liquid_chro_detail.Error_bar_value_uom, anl_liquid_chro_detail.Expiry_date, anl_liquid_chro_detail.Measured_ind, anl_liquid_chro_detail.Measurement_type, anl_liquid_chro_detail.Ppdm_guid, anl_liquid_chro_detail.Preferred_ind, anl_liquid_chro_detail.Problem_ind, anl_liquid_chro_detail.Quantif_additive_amt, anl_liquid_chro_detail.Quantif_additive_amt_ouom, anl_liquid_chro_detail.Quantif_additive_amt_uom, anl_liquid_chro_detail.Quantif_additive_desc, anl_liquid_chro_detail.Quantif_additive_type, anl_liquid_chro_detail.Remark, anl_liquid_chro_detail.Reported_ind, anl_liquid_chro_detail.Source, anl_liquid_chro_detail.Step_seq_no, anl_liquid_chro_detail.Substance_id, anl_liquid_chro_detail.Row_changed_by, anl_liquid_chro_detail.Row_changed_date, anl_liquid_chro_detail.Row_created_by, anl_liquid_chro_detail.Row_effective_date, anl_liquid_chro_detail.Row_expiry_date, anl_liquid_chro_detail.Row_quality, anl_liquid_chro_detail.Analysis_id)
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

func PatchAnlLiquidChroDetail(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_liquid_chro_detail set "
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

func DeleteAnlLiquidChroDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_liquid_chro_detail dto.Anl_liquid_chro_detail
	anl_liquid_chro_detail.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_liquid_chro_detail where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_liquid_chro_detail.Analysis_id)
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


