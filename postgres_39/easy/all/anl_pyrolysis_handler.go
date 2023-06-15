package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlPyrolysis(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_pyrolysis")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_pyrolysis

	for rows.Next() {
		var anl_pyrolysis dto.Anl_pyrolysis
		if err := rows.Scan(&anl_pyrolysis.Analysis_id, &anl_pyrolysis.Anl_source, &anl_pyrolysis.Pyrolysis_anl_obs_no, &anl_pyrolysis.Active_ind, &anl_pyrolysis.Calculated_ind, &anl_pyrolysis.Calculate_method_id, &anl_pyrolysis.Effective_date, &anl_pyrolysis.Expiry_date, &anl_pyrolysis.Max_temperature, &anl_pyrolysis.Max_temperature_ouom, &anl_pyrolysis.Max_temperature_uom, &anl_pyrolysis.Peak_temperature, &anl_pyrolysis.Peak_temperature_ouom, &anl_pyrolysis.Peak_temperature_uom, &anl_pyrolysis.Ppdm_guid, &anl_pyrolysis.Preferred_ind, &anl_pyrolysis.Problem_ind, &anl_pyrolysis.Remark, &anl_pyrolysis.Reported_ind, &anl_pyrolysis.Rep_bitumen_index, &anl_pyrolysis.Rep_bitumen_index_ouom, &anl_pyrolysis.Rep_bitumen_index_uom, &anl_pyrolysis.Rep_hydrogen_index, &anl_pyrolysis.Rep_hydrogen_index_ouom, &anl_pyrolysis.Rep_hydrogen_index_uom, &anl_pyrolysis.Rep_oxygen_index, &anl_pyrolysis.Rep_oxygen_index_ouom, &anl_pyrolysis.Rep_oxygen_index_uom, &anl_pyrolysis.Rep_production_index, &anl_pyrolysis.Rep_pyrolisable_carbon, &anl_pyrolysis.Rep_pyrolisable_carbon_ouom, &anl_pyrolysis.Rep_pyrolisable_carbon_uom, &anl_pyrolysis.Source, &anl_pyrolysis.Step_seq_no, &anl_pyrolysis.S_0, &anl_pyrolysis.S_0_ouom, &anl_pyrolysis.S_1, &anl_pyrolysis.S_1_ouom, &anl_pyrolysis.S_2, &anl_pyrolysis.S_2_ouom, &anl_pyrolysis.S_3, &anl_pyrolysis.S_3_ouom, &anl_pyrolysis.S_4, &anl_pyrolysis.S_4_ouom, &anl_pyrolysis.S_5, &anl_pyrolysis.S_5_ouom, &anl_pyrolysis.Total_organic_carbon, &anl_pyrolysis.Total_organic_carbon_ouom, &anl_pyrolysis.Total_organic_carbon_uom, &anl_pyrolysis.Valid_anl_ind, &anl_pyrolysis.Row_changed_by, &anl_pyrolysis.Row_changed_date, &anl_pyrolysis.Row_created_by, &anl_pyrolysis.Row_created_date, &anl_pyrolysis.Row_effective_date, &anl_pyrolysis.Row_expiry_date, &anl_pyrolysis.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_pyrolysis to result
		result = append(result, anl_pyrolysis)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlPyrolysis(c *fiber.Ctx) error {
	var anl_pyrolysis dto.Anl_pyrolysis

	setDefaults(&anl_pyrolysis)

	if err := json.Unmarshal(c.Body(), &anl_pyrolysis); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_pyrolysis values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57)")
	if err != nil {
		return err
	}
	anl_pyrolysis.Row_created_date = formatDate(anl_pyrolysis.Row_created_date)
	anl_pyrolysis.Row_changed_date = formatDate(anl_pyrolysis.Row_changed_date)
	anl_pyrolysis.Effective_date = formatDateString(anl_pyrolysis.Effective_date)
	anl_pyrolysis.Expiry_date = formatDateString(anl_pyrolysis.Expiry_date)
	anl_pyrolysis.Row_effective_date = formatDateString(anl_pyrolysis.Row_effective_date)
	anl_pyrolysis.Row_expiry_date = formatDateString(anl_pyrolysis.Row_expiry_date)






	rows, err := stmt.Exec(anl_pyrolysis.Analysis_id, anl_pyrolysis.Anl_source, anl_pyrolysis.Pyrolysis_anl_obs_no, anl_pyrolysis.Active_ind, anl_pyrolysis.Calculated_ind, anl_pyrolysis.Calculate_method_id, anl_pyrolysis.Effective_date, anl_pyrolysis.Expiry_date, anl_pyrolysis.Max_temperature, anl_pyrolysis.Max_temperature_ouom, anl_pyrolysis.Max_temperature_uom, anl_pyrolysis.Peak_temperature, anl_pyrolysis.Peak_temperature_ouom, anl_pyrolysis.Peak_temperature_uom, anl_pyrolysis.Ppdm_guid, anl_pyrolysis.Preferred_ind, anl_pyrolysis.Problem_ind, anl_pyrolysis.Remark, anl_pyrolysis.Reported_ind, anl_pyrolysis.Rep_bitumen_index, anl_pyrolysis.Rep_bitumen_index_ouom, anl_pyrolysis.Rep_bitumen_index_uom, anl_pyrolysis.Rep_hydrogen_index, anl_pyrolysis.Rep_hydrogen_index_ouom, anl_pyrolysis.Rep_hydrogen_index_uom, anl_pyrolysis.Rep_oxygen_index, anl_pyrolysis.Rep_oxygen_index_ouom, anl_pyrolysis.Rep_oxygen_index_uom, anl_pyrolysis.Rep_production_index, anl_pyrolysis.Rep_pyrolisable_carbon, anl_pyrolysis.Rep_pyrolisable_carbon_ouom, anl_pyrolysis.Rep_pyrolisable_carbon_uom, anl_pyrolysis.Source, anl_pyrolysis.Step_seq_no, anl_pyrolysis.S_0, anl_pyrolysis.S_0_ouom, anl_pyrolysis.S_1, anl_pyrolysis.S_1_ouom, anl_pyrolysis.S_2, anl_pyrolysis.S_2_ouom, anl_pyrolysis.S_3, anl_pyrolysis.S_3_ouom, anl_pyrolysis.S_4, anl_pyrolysis.S_4_ouom, anl_pyrolysis.S_5, anl_pyrolysis.S_5_ouom, anl_pyrolysis.Total_organic_carbon, anl_pyrolysis.Total_organic_carbon_ouom, anl_pyrolysis.Total_organic_carbon_uom, anl_pyrolysis.Valid_anl_ind, anl_pyrolysis.Row_changed_by, anl_pyrolysis.Row_changed_date, anl_pyrolysis.Row_created_by, anl_pyrolysis.Row_created_date, anl_pyrolysis.Row_effective_date, anl_pyrolysis.Row_expiry_date, anl_pyrolysis.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlPyrolysis(c *fiber.Ctx) error {
	var anl_pyrolysis dto.Anl_pyrolysis

	setDefaults(&anl_pyrolysis)

	if err := json.Unmarshal(c.Body(), &anl_pyrolysis); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_pyrolysis.Analysis_id = id

    if anl_pyrolysis.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_pyrolysis set anl_source = :1, pyrolysis_anl_obs_no = :2, active_ind = :3, calculated_ind = :4, calculate_method_id = :5, effective_date = :6, expiry_date = :7, max_temperature = :8, max_temperature_ouom = :9, max_temperature_uom = :10, peak_temperature = :11, peak_temperature_ouom = :12, peak_temperature_uom = :13, ppdm_guid = :14, preferred_ind = :15, problem_ind = :16, remark = :17, reported_ind = :18, rep_bitumen_index = :19, rep_bitumen_index_ouom = :20, rep_bitumen_index_uom = :21, rep_hydrogen_index = :22, rep_hydrogen_index_ouom = :23, rep_hydrogen_index_uom = :24, rep_oxygen_index = :25, rep_oxygen_index_ouom = :26, rep_oxygen_index_uom = :27, rep_production_index = :28, rep_pyrolisable_carbon = :29, rep_pyrolisable_carbon_ouom = :30, rep_pyrolisable_carbon_uom = :31, source = :32, step_seq_no = :33, s_0 = :34, s_0_ouom = :35, s_1 = :36, s_1_ouom = :37, s_2 = :38, s_2_ouom = :39, s_3 = :40, s_3_ouom = :41, s_4 = :42, s_4_ouom = :43, s_5 = :44, s_5_ouom = :45, total_organic_carbon = :46, total_organic_carbon_ouom = :47, total_organic_carbon_uom = :48, valid_anl_ind = :49, row_changed_by = :50, row_changed_date = :51, row_created_by = :52, row_effective_date = :53, row_expiry_date = :54, row_quality = :55 where analysis_id = :57")
	if err != nil {
		return err
	}

	anl_pyrolysis.Row_changed_date = formatDate(anl_pyrolysis.Row_changed_date)
	anl_pyrolysis.Effective_date = formatDateString(anl_pyrolysis.Effective_date)
	anl_pyrolysis.Expiry_date = formatDateString(anl_pyrolysis.Expiry_date)
	anl_pyrolysis.Row_effective_date = formatDateString(anl_pyrolysis.Row_effective_date)
	anl_pyrolysis.Row_expiry_date = formatDateString(anl_pyrolysis.Row_expiry_date)






	rows, err := stmt.Exec(anl_pyrolysis.Anl_source, anl_pyrolysis.Pyrolysis_anl_obs_no, anl_pyrolysis.Active_ind, anl_pyrolysis.Calculated_ind, anl_pyrolysis.Calculate_method_id, anl_pyrolysis.Effective_date, anl_pyrolysis.Expiry_date, anl_pyrolysis.Max_temperature, anl_pyrolysis.Max_temperature_ouom, anl_pyrolysis.Max_temperature_uom, anl_pyrolysis.Peak_temperature, anl_pyrolysis.Peak_temperature_ouom, anl_pyrolysis.Peak_temperature_uom, anl_pyrolysis.Ppdm_guid, anl_pyrolysis.Preferred_ind, anl_pyrolysis.Problem_ind, anl_pyrolysis.Remark, anl_pyrolysis.Reported_ind, anl_pyrolysis.Rep_bitumen_index, anl_pyrolysis.Rep_bitumen_index_ouom, anl_pyrolysis.Rep_bitumen_index_uom, anl_pyrolysis.Rep_hydrogen_index, anl_pyrolysis.Rep_hydrogen_index_ouom, anl_pyrolysis.Rep_hydrogen_index_uom, anl_pyrolysis.Rep_oxygen_index, anl_pyrolysis.Rep_oxygen_index_ouom, anl_pyrolysis.Rep_oxygen_index_uom, anl_pyrolysis.Rep_production_index, anl_pyrolysis.Rep_pyrolisable_carbon, anl_pyrolysis.Rep_pyrolisable_carbon_ouom, anl_pyrolysis.Rep_pyrolisable_carbon_uom, anl_pyrolysis.Source, anl_pyrolysis.Step_seq_no, anl_pyrolysis.S_0, anl_pyrolysis.S_0_ouom, anl_pyrolysis.S_1, anl_pyrolysis.S_1_ouom, anl_pyrolysis.S_2, anl_pyrolysis.S_2_ouom, anl_pyrolysis.S_3, anl_pyrolysis.S_3_ouom, anl_pyrolysis.S_4, anl_pyrolysis.S_4_ouom, anl_pyrolysis.S_5, anl_pyrolysis.S_5_ouom, anl_pyrolysis.Total_organic_carbon, anl_pyrolysis.Total_organic_carbon_ouom, anl_pyrolysis.Total_organic_carbon_uom, anl_pyrolysis.Valid_anl_ind, anl_pyrolysis.Row_changed_by, anl_pyrolysis.Row_changed_date, anl_pyrolysis.Row_created_by, anl_pyrolysis.Row_effective_date, anl_pyrolysis.Row_expiry_date, anl_pyrolysis.Row_quality, anl_pyrolysis.Analysis_id)
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

func PatchAnlPyrolysis(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_pyrolysis set "
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

func DeleteAnlPyrolysis(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_pyrolysis dto.Anl_pyrolysis
	anl_pyrolysis.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_pyrolysis where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_pyrolysis.Analysis_id)
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


