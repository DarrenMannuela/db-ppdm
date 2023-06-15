package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlGasChro(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_gas_chro")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_gas_chro

	for rows.Next() {
		var anl_gas_chro dto.Anl_gas_chro
		if err := rows.Scan(&anl_gas_chro.Analysis_id, &anl_gas_chro.Anl_source, &anl_gas_chro.Anl_obs_no, &anl_gas_chro.Active_ind, &anl_gas_chro.Anl_value, &anl_gas_chro.Anl_value_code, &anl_gas_chro.Anl_value_ouom, &anl_gas_chro.Anl_value_remark, &anl_gas_chro.Anl_value_type, &anl_gas_chro.Anl_value_uom, &anl_gas_chro.Calculated_ind, &anl_gas_chro.Calculate_method_id, &anl_gas_chro.Daughter_ion_id, &anl_gas_chro.Effective_date, &anl_gas_chro.Error_bar_value, &anl_gas_chro.Error_bar_value_ouom, &anl_gas_chro.Error_bar_value_uom, &anl_gas_chro.Expiry_date, &anl_gas_chro.Gas_anl_value_type, &anl_gas_chro.Measurement_type, &anl_gas_chro.Parent_ion_id, &anl_gas_chro.Ppdm_guid, &anl_gas_chro.Preferred_ind, &anl_gas_chro.Problem_ind, &anl_gas_chro.Qualifier_ion_id, &anl_gas_chro.Quantif_additive_amt, &anl_gas_chro.Quantif_additive_amt_ouom, &anl_gas_chro.Quantif_additive_amt_uom, &anl_gas_chro.Quantif_additive_desc, &anl_gas_chro.Quantif_additive_type, &anl_gas_chro.Remark, &anl_gas_chro.Reported_ind, &anl_gas_chro.Source, &anl_gas_chro.Step_seq_no, &anl_gas_chro.Substance_id, &anl_gas_chro.Row_changed_by, &anl_gas_chro.Row_changed_date, &anl_gas_chro.Row_created_by, &anl_gas_chro.Row_created_date, &anl_gas_chro.Row_effective_date, &anl_gas_chro.Row_expiry_date, &anl_gas_chro.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_gas_chro to result
		result = append(result, anl_gas_chro)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlGasChro(c *fiber.Ctx) error {
	var anl_gas_chro dto.Anl_gas_chro

	setDefaults(&anl_gas_chro)

	if err := json.Unmarshal(c.Body(), &anl_gas_chro); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_gas_chro values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42)")
	if err != nil {
		return err
	}
	anl_gas_chro.Row_created_date = formatDate(anl_gas_chro.Row_created_date)
	anl_gas_chro.Row_changed_date = formatDate(anl_gas_chro.Row_changed_date)
	anl_gas_chro.Effective_date = formatDateString(anl_gas_chro.Effective_date)
	anl_gas_chro.Expiry_date = formatDateString(anl_gas_chro.Expiry_date)
	anl_gas_chro.Row_effective_date = formatDateString(anl_gas_chro.Row_effective_date)
	anl_gas_chro.Row_expiry_date = formatDateString(anl_gas_chro.Row_expiry_date)






	rows, err := stmt.Exec(anl_gas_chro.Analysis_id, anl_gas_chro.Anl_source, anl_gas_chro.Anl_obs_no, anl_gas_chro.Active_ind, anl_gas_chro.Anl_value, anl_gas_chro.Anl_value_code, anl_gas_chro.Anl_value_ouom, anl_gas_chro.Anl_value_remark, anl_gas_chro.Anl_value_type, anl_gas_chro.Anl_value_uom, anl_gas_chro.Calculated_ind, anl_gas_chro.Calculate_method_id, anl_gas_chro.Daughter_ion_id, anl_gas_chro.Effective_date, anl_gas_chro.Error_bar_value, anl_gas_chro.Error_bar_value_ouom, anl_gas_chro.Error_bar_value_uom, anl_gas_chro.Expiry_date, anl_gas_chro.Gas_anl_value_type, anl_gas_chro.Measurement_type, anl_gas_chro.Parent_ion_id, anl_gas_chro.Ppdm_guid, anl_gas_chro.Preferred_ind, anl_gas_chro.Problem_ind, anl_gas_chro.Qualifier_ion_id, anl_gas_chro.Quantif_additive_amt, anl_gas_chro.Quantif_additive_amt_ouom, anl_gas_chro.Quantif_additive_amt_uom, anl_gas_chro.Quantif_additive_desc, anl_gas_chro.Quantif_additive_type, anl_gas_chro.Remark, anl_gas_chro.Reported_ind, anl_gas_chro.Source, anl_gas_chro.Step_seq_no, anl_gas_chro.Substance_id, anl_gas_chro.Row_changed_by, anl_gas_chro.Row_changed_date, anl_gas_chro.Row_created_by, anl_gas_chro.Row_created_date, anl_gas_chro.Row_effective_date, anl_gas_chro.Row_expiry_date, anl_gas_chro.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlGasChro(c *fiber.Ctx) error {
	var anl_gas_chro dto.Anl_gas_chro

	setDefaults(&anl_gas_chro)

	if err := json.Unmarshal(c.Body(), &anl_gas_chro); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_gas_chro.Analysis_id = id

    if anl_gas_chro.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_gas_chro set anl_source = :1, anl_obs_no = :2, active_ind = :3, anl_value = :4, anl_value_code = :5, anl_value_ouom = :6, anl_value_remark = :7, anl_value_type = :8, anl_value_uom = :9, calculated_ind = :10, calculate_method_id = :11, daughter_ion_id = :12, effective_date = :13, error_bar_value = :14, error_bar_value_ouom = :15, error_bar_value_uom = :16, expiry_date = :17, gas_anl_value_type = :18, measurement_type = :19, parent_ion_id = :20, ppdm_guid = :21, preferred_ind = :22, problem_ind = :23, qualifier_ion_id = :24, quantif_additive_amt = :25, quantif_additive_amt_ouom = :26, quantif_additive_amt_uom = :27, quantif_additive_desc = :28, quantif_additive_type = :29, remark = :30, reported_ind = :31, source = :32, step_seq_no = :33, substance_id = :34, row_changed_by = :35, row_changed_date = :36, row_created_by = :37, row_effective_date = :38, row_expiry_date = :39, row_quality = :40 where analysis_id = :42")
	if err != nil {
		return err
	}

	anl_gas_chro.Row_changed_date = formatDate(anl_gas_chro.Row_changed_date)
	anl_gas_chro.Effective_date = formatDateString(anl_gas_chro.Effective_date)
	anl_gas_chro.Expiry_date = formatDateString(anl_gas_chro.Expiry_date)
	anl_gas_chro.Row_effective_date = formatDateString(anl_gas_chro.Row_effective_date)
	anl_gas_chro.Row_expiry_date = formatDateString(anl_gas_chro.Row_expiry_date)






	rows, err := stmt.Exec(anl_gas_chro.Anl_source, anl_gas_chro.Anl_obs_no, anl_gas_chro.Active_ind, anl_gas_chro.Anl_value, anl_gas_chro.Anl_value_code, anl_gas_chro.Anl_value_ouom, anl_gas_chro.Anl_value_remark, anl_gas_chro.Anl_value_type, anl_gas_chro.Anl_value_uom, anl_gas_chro.Calculated_ind, anl_gas_chro.Calculate_method_id, anl_gas_chro.Daughter_ion_id, anl_gas_chro.Effective_date, anl_gas_chro.Error_bar_value, anl_gas_chro.Error_bar_value_ouom, anl_gas_chro.Error_bar_value_uom, anl_gas_chro.Expiry_date, anl_gas_chro.Gas_anl_value_type, anl_gas_chro.Measurement_type, anl_gas_chro.Parent_ion_id, anl_gas_chro.Ppdm_guid, anl_gas_chro.Preferred_ind, anl_gas_chro.Problem_ind, anl_gas_chro.Qualifier_ion_id, anl_gas_chro.Quantif_additive_amt, anl_gas_chro.Quantif_additive_amt_ouom, anl_gas_chro.Quantif_additive_amt_uom, anl_gas_chro.Quantif_additive_desc, anl_gas_chro.Quantif_additive_type, anl_gas_chro.Remark, anl_gas_chro.Reported_ind, anl_gas_chro.Source, anl_gas_chro.Step_seq_no, anl_gas_chro.Substance_id, anl_gas_chro.Row_changed_by, anl_gas_chro.Row_changed_date, anl_gas_chro.Row_created_by, anl_gas_chro.Row_effective_date, anl_gas_chro.Row_expiry_date, anl_gas_chro.Row_quality, anl_gas_chro.Analysis_id)
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

func PatchAnlGasChro(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_gas_chro set "
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

func DeleteAnlGasChro(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_gas_chro dto.Anl_gas_chro
	anl_gas_chro.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_gas_chro where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_gas_chro.Analysis_id)
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


