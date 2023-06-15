package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlMaceralMaturity(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_maceral_maturity")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_maceral_maturity

	for rows.Next() {
		var anl_maceral_maturity dto.Anl_maceral_maturity
		if err := rows.Scan(&anl_maceral_maturity.Analysis_id, &anl_maceral_maturity.Anl_source, &anl_maceral_maturity.Measurement_obs_no, &anl_maceral_maturity.Active_ind, &anl_maceral_maturity.Coal_rank_id, &anl_maceral_maturity.Effective_date, &anl_maceral_maturity.Expiry_date, &anl_maceral_maturity.Fluor_color, &anl_maceral_maturity.Fluor_intensity_desc, &anl_maceral_maturity.Fluor_intensity_value, &anl_maceral_maturity.Fluor_intensity_value_ouom, &anl_maceral_maturity.Mean_max_reflectance, &anl_maceral_maturity.Mean_random_reflectance, &anl_maceral_maturity.Measurements_count, &anl_maceral_maturity.Population_name, &anl_maceral_maturity.Population_num, &anl_maceral_maturity.Population_obs_no, &anl_maceral_maturity.Ppdm_guid, &anl_maceral_maturity.Preferred_ind, &anl_maceral_maturity.Problem_ind, &anl_maceral_maturity.Reflect_max_meas_value, &anl_maceral_maturity.Reflect_max_meas_value_ouom, &anl_maceral_maturity.Reflect_max_meas_value_uom, &anl_maceral_maturity.Reflect_meas_value, &anl_maceral_maturity.Reflect_meas_value_ouom, &anl_maceral_maturity.Reflect_meas_value_uom, &anl_maceral_maturity.Reflect_min_meas_value, &anl_maceral_maturity.Reflect_min_meas_value_ouom, &anl_maceral_maturity.Reflect_min_meas_value_uom, &anl_maceral_maturity.Reflect_std_dev_value, &anl_maceral_maturity.Refl_random_meas_value, &anl_maceral_maturity.Refl_random_meas_value_ouom, &anl_maceral_maturity.Refl_random_meas_value_uom, &anl_maceral_maturity.Remark, &anl_maceral_maturity.Reworked_percent, &anl_maceral_maturity.Sample_tot_maceral_val, &anl_maceral_maturity.Sample_tot_maceral_val_ouom, &anl_maceral_maturity.Source, &anl_maceral_maturity.Step_seq_no, &anl_maceral_maturity.Substance_id, &anl_maceral_maturity.Row_changed_by, &anl_maceral_maturity.Row_changed_date, &anl_maceral_maturity.Row_created_by, &anl_maceral_maturity.Row_created_date, &anl_maceral_maturity.Row_effective_date, &anl_maceral_maturity.Row_expiry_date, &anl_maceral_maturity.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_maceral_maturity to result
		result = append(result, anl_maceral_maturity)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlMaceralMaturity(c *fiber.Ctx) error {
	var anl_maceral_maturity dto.Anl_maceral_maturity

	setDefaults(&anl_maceral_maturity)

	if err := json.Unmarshal(c.Body(), &anl_maceral_maturity); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_maceral_maturity values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47)")
	if err != nil {
		return err
	}
	anl_maceral_maturity.Row_created_date = formatDate(anl_maceral_maturity.Row_created_date)
	anl_maceral_maturity.Row_changed_date = formatDate(anl_maceral_maturity.Row_changed_date)
	anl_maceral_maturity.Effective_date = formatDateString(anl_maceral_maturity.Effective_date)
	anl_maceral_maturity.Expiry_date = formatDateString(anl_maceral_maturity.Expiry_date)
	anl_maceral_maturity.Row_effective_date = formatDateString(anl_maceral_maturity.Row_effective_date)
	anl_maceral_maturity.Row_expiry_date = formatDateString(anl_maceral_maturity.Row_expiry_date)






	rows, err := stmt.Exec(anl_maceral_maturity.Analysis_id, anl_maceral_maturity.Anl_source, anl_maceral_maturity.Measurement_obs_no, anl_maceral_maturity.Active_ind, anl_maceral_maturity.Coal_rank_id, anl_maceral_maturity.Effective_date, anl_maceral_maturity.Expiry_date, anl_maceral_maturity.Fluor_color, anl_maceral_maturity.Fluor_intensity_desc, anl_maceral_maturity.Fluor_intensity_value, anl_maceral_maturity.Fluor_intensity_value_ouom, anl_maceral_maturity.Mean_max_reflectance, anl_maceral_maturity.Mean_random_reflectance, anl_maceral_maturity.Measurements_count, anl_maceral_maturity.Population_name, anl_maceral_maturity.Population_num, anl_maceral_maturity.Population_obs_no, anl_maceral_maturity.Ppdm_guid, anl_maceral_maturity.Preferred_ind, anl_maceral_maturity.Problem_ind, anl_maceral_maturity.Reflect_max_meas_value, anl_maceral_maturity.Reflect_max_meas_value_ouom, anl_maceral_maturity.Reflect_max_meas_value_uom, anl_maceral_maturity.Reflect_meas_value, anl_maceral_maturity.Reflect_meas_value_ouom, anl_maceral_maturity.Reflect_meas_value_uom, anl_maceral_maturity.Reflect_min_meas_value, anl_maceral_maturity.Reflect_min_meas_value_ouom, anl_maceral_maturity.Reflect_min_meas_value_uom, anl_maceral_maturity.Reflect_std_dev_value, anl_maceral_maturity.Refl_random_meas_value, anl_maceral_maturity.Refl_random_meas_value_ouom, anl_maceral_maturity.Refl_random_meas_value_uom, anl_maceral_maturity.Remark, anl_maceral_maturity.Reworked_percent, anl_maceral_maturity.Sample_tot_maceral_val, anl_maceral_maturity.Sample_tot_maceral_val_ouom, anl_maceral_maturity.Source, anl_maceral_maturity.Step_seq_no, anl_maceral_maturity.Substance_id, anl_maceral_maturity.Row_changed_by, anl_maceral_maturity.Row_changed_date, anl_maceral_maturity.Row_created_by, anl_maceral_maturity.Row_created_date, anl_maceral_maturity.Row_effective_date, anl_maceral_maturity.Row_expiry_date, anl_maceral_maturity.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlMaceralMaturity(c *fiber.Ctx) error {
	var anl_maceral_maturity dto.Anl_maceral_maturity

	setDefaults(&anl_maceral_maturity)

	if err := json.Unmarshal(c.Body(), &anl_maceral_maturity); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_maceral_maturity.Analysis_id = id

    if anl_maceral_maturity.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_maceral_maturity set anl_source = :1, measurement_obs_no = :2, active_ind = :3, coal_rank_id = :4, effective_date = :5, expiry_date = :6, fluor_color = :7, fluor_intensity_desc = :8, fluor_intensity_value = :9, fluor_intensity_value_ouom = :10, mean_max_reflectance = :11, mean_random_reflectance = :12, measurements_count = :13, population_name = :14, population_num = :15, population_obs_no = :16, ppdm_guid = :17, preferred_ind = :18, problem_ind = :19, reflect_max_meas_value = :20, reflect_max_meas_value_ouom = :21, reflect_max_meas_value_uom = :22, reflect_meas_value = :23, reflect_meas_value_ouom = :24, reflect_meas_value_uom = :25, reflect_min_meas_value = :26, reflect_min_meas_value_ouom = :27, reflect_min_meas_value_uom = :28, reflect_std_dev_value = :29, refl_random_meas_value = :30, refl_random_meas_value_ouom = :31, refl_random_meas_value_uom = :32, remark = :33, reworked_percent = :34, sample_tot_maceral_val = :35, sample_tot_maceral_val_ouom = :36, source = :37, step_seq_no = :38, substance_id = :39, row_changed_by = :40, row_changed_date = :41, row_created_by = :42, row_effective_date = :43, row_expiry_date = :44, row_quality = :45 where analysis_id = :47")
	if err != nil {
		return err
	}

	anl_maceral_maturity.Row_changed_date = formatDate(anl_maceral_maturity.Row_changed_date)
	anl_maceral_maturity.Effective_date = formatDateString(anl_maceral_maturity.Effective_date)
	anl_maceral_maturity.Expiry_date = formatDateString(anl_maceral_maturity.Expiry_date)
	anl_maceral_maturity.Row_effective_date = formatDateString(anl_maceral_maturity.Row_effective_date)
	anl_maceral_maturity.Row_expiry_date = formatDateString(anl_maceral_maturity.Row_expiry_date)






	rows, err := stmt.Exec(anl_maceral_maturity.Anl_source, anl_maceral_maturity.Measurement_obs_no, anl_maceral_maturity.Active_ind, anl_maceral_maturity.Coal_rank_id, anl_maceral_maturity.Effective_date, anl_maceral_maturity.Expiry_date, anl_maceral_maturity.Fluor_color, anl_maceral_maturity.Fluor_intensity_desc, anl_maceral_maturity.Fluor_intensity_value, anl_maceral_maturity.Fluor_intensity_value_ouom, anl_maceral_maturity.Mean_max_reflectance, anl_maceral_maturity.Mean_random_reflectance, anl_maceral_maturity.Measurements_count, anl_maceral_maturity.Population_name, anl_maceral_maturity.Population_num, anl_maceral_maturity.Population_obs_no, anl_maceral_maturity.Ppdm_guid, anl_maceral_maturity.Preferred_ind, anl_maceral_maturity.Problem_ind, anl_maceral_maturity.Reflect_max_meas_value, anl_maceral_maturity.Reflect_max_meas_value_ouom, anl_maceral_maturity.Reflect_max_meas_value_uom, anl_maceral_maturity.Reflect_meas_value, anl_maceral_maturity.Reflect_meas_value_ouom, anl_maceral_maturity.Reflect_meas_value_uom, anl_maceral_maturity.Reflect_min_meas_value, anl_maceral_maturity.Reflect_min_meas_value_ouom, anl_maceral_maturity.Reflect_min_meas_value_uom, anl_maceral_maturity.Reflect_std_dev_value, anl_maceral_maturity.Refl_random_meas_value, anl_maceral_maturity.Refl_random_meas_value_ouom, anl_maceral_maturity.Refl_random_meas_value_uom, anl_maceral_maturity.Remark, anl_maceral_maturity.Reworked_percent, anl_maceral_maturity.Sample_tot_maceral_val, anl_maceral_maturity.Sample_tot_maceral_val_ouom, anl_maceral_maturity.Source, anl_maceral_maturity.Step_seq_no, anl_maceral_maturity.Substance_id, anl_maceral_maturity.Row_changed_by, anl_maceral_maturity.Row_changed_date, anl_maceral_maturity.Row_created_by, anl_maceral_maturity.Row_effective_date, anl_maceral_maturity.Row_expiry_date, anl_maceral_maturity.Row_quality, anl_maceral_maturity.Analysis_id)
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

func PatchAnlMaceralMaturity(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_maceral_maturity set "
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

func DeleteAnlMaceralMaturity(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_maceral_maturity dto.Anl_maceral_maturity
	anl_maceral_maturity.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_maceral_maturity where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_maceral_maturity.Analysis_id)
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


