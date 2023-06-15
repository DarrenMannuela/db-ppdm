package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlOilDistill(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_oil_distill")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_oil_distill

	for rows.Next() {
		var anl_oil_distill dto.Anl_oil_distill
		if err := rows.Scan(&anl_oil_distill.Analysis_id, &anl_oil_distill.Anl_source, &anl_oil_distill.Dstl_summary_obs_no, &anl_oil_distill.Active_ind, &anl_oil_distill.Atmosp_dstl_press, &anl_oil_distill.Atmosp_dstl_press_ouom, &anl_oil_distill.Atmosp_dstl_temp, &anl_oil_distill.Atmosp_dstl_temp_ouom, &anl_oil_distill.Calculated_ind, &anl_oil_distill.Calculate_method_id, &anl_oil_distill.Effective_date, &anl_oil_distill.Expiry_date, &anl_oil_distill.Final_boil_pt_temp, &anl_oil_distill.Final_boil_pt_temp_ouom, &anl_oil_distill.Measurement_temp, &anl_oil_distill.Measurement_temp_ouom, &anl_oil_distill.Ppdm_guid, &anl_oil_distill.Preferred_ind, &anl_oil_distill.Problem_ind, &anl_oil_distill.Remark, &anl_oil_distill.Reported_ind, &anl_oil_distill.Source, &anl_oil_distill.Start_boil_pt_temp, &anl_oil_distill.Start_boil_pt_temp_ouom, &anl_oil_distill.Step_seq_no, &anl_oil_distill.Substance_id, &anl_oil_distill.Volume_fraction_type, &anl_oil_distill.Vol_fraction_temp, &anl_oil_distill.Vol_fraction_temp_ouom, &anl_oil_distill.Vol_fraction_value, &anl_oil_distill.Vol_fraction_value_ouom, &anl_oil_distill.Weight_cut, &anl_oil_distill.Weight_cut_ouom, &anl_oil_distill.Row_changed_by, &anl_oil_distill.Row_changed_date, &anl_oil_distill.Row_created_by, &anl_oil_distill.Row_created_date, &anl_oil_distill.Row_effective_date, &anl_oil_distill.Row_expiry_date, &anl_oil_distill.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_oil_distill to result
		result = append(result, anl_oil_distill)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlOilDistill(c *fiber.Ctx) error {
	var anl_oil_distill dto.Anl_oil_distill

	setDefaults(&anl_oil_distill)

	if err := json.Unmarshal(c.Body(), &anl_oil_distill); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_oil_distill values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40)")
	if err != nil {
		return err
	}
	anl_oil_distill.Row_created_date = formatDate(anl_oil_distill.Row_created_date)
	anl_oil_distill.Row_changed_date = formatDate(anl_oil_distill.Row_changed_date)
	anl_oil_distill.Effective_date = formatDateString(anl_oil_distill.Effective_date)
	anl_oil_distill.Expiry_date = formatDateString(anl_oil_distill.Expiry_date)
	anl_oil_distill.Row_effective_date = formatDateString(anl_oil_distill.Row_effective_date)
	anl_oil_distill.Row_expiry_date = formatDateString(anl_oil_distill.Row_expiry_date)






	rows, err := stmt.Exec(anl_oil_distill.Analysis_id, anl_oil_distill.Anl_source, anl_oil_distill.Dstl_summary_obs_no, anl_oil_distill.Active_ind, anl_oil_distill.Atmosp_dstl_press, anl_oil_distill.Atmosp_dstl_press_ouom, anl_oil_distill.Atmosp_dstl_temp, anl_oil_distill.Atmosp_dstl_temp_ouom, anl_oil_distill.Calculated_ind, anl_oil_distill.Calculate_method_id, anl_oil_distill.Effective_date, anl_oil_distill.Expiry_date, anl_oil_distill.Final_boil_pt_temp, anl_oil_distill.Final_boil_pt_temp_ouom, anl_oil_distill.Measurement_temp, anl_oil_distill.Measurement_temp_ouom, anl_oil_distill.Ppdm_guid, anl_oil_distill.Preferred_ind, anl_oil_distill.Problem_ind, anl_oil_distill.Remark, anl_oil_distill.Reported_ind, anl_oil_distill.Source, anl_oil_distill.Start_boil_pt_temp, anl_oil_distill.Start_boil_pt_temp_ouom, anl_oil_distill.Step_seq_no, anl_oil_distill.Substance_id, anl_oil_distill.Volume_fraction_type, anl_oil_distill.Vol_fraction_temp, anl_oil_distill.Vol_fraction_temp_ouom, anl_oil_distill.Vol_fraction_value, anl_oil_distill.Vol_fraction_value_ouom, anl_oil_distill.Weight_cut, anl_oil_distill.Weight_cut_ouom, anl_oil_distill.Row_changed_by, anl_oil_distill.Row_changed_date, anl_oil_distill.Row_created_by, anl_oil_distill.Row_created_date, anl_oil_distill.Row_effective_date, anl_oil_distill.Row_expiry_date, anl_oil_distill.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlOilDistill(c *fiber.Ctx) error {
	var anl_oil_distill dto.Anl_oil_distill

	setDefaults(&anl_oil_distill)

	if err := json.Unmarshal(c.Body(), &anl_oil_distill); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_oil_distill.Analysis_id = id

    if anl_oil_distill.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_oil_distill set anl_source = :1, dstl_summary_obs_no = :2, active_ind = :3, atmosp_dstl_press = :4, atmosp_dstl_press_ouom = :5, atmosp_dstl_temp = :6, atmosp_dstl_temp_ouom = :7, calculated_ind = :8, calculate_method_id = :9, effective_date = :10, expiry_date = :11, final_boil_pt_temp = :12, final_boil_pt_temp_ouom = :13, measurement_temp = :14, measurement_temp_ouom = :15, ppdm_guid = :16, preferred_ind = :17, problem_ind = :18, remark = :19, reported_ind = :20, source = :21, start_boil_pt_temp = :22, start_boil_pt_temp_ouom = :23, step_seq_no = :24, substance_id = :25, volume_fraction_type = :26, vol_fraction_temp = :27, vol_fraction_temp_ouom = :28, vol_fraction_value = :29, vol_fraction_value_ouom = :30, weight_cut = :31, weight_cut_ouom = :32, row_changed_by = :33, row_changed_date = :34, row_created_by = :35, row_effective_date = :36, row_expiry_date = :37, row_quality = :38 where analysis_id = :40")
	if err != nil {
		return err
	}

	anl_oil_distill.Row_changed_date = formatDate(anl_oil_distill.Row_changed_date)
	anl_oil_distill.Effective_date = formatDateString(anl_oil_distill.Effective_date)
	anl_oil_distill.Expiry_date = formatDateString(anl_oil_distill.Expiry_date)
	anl_oil_distill.Row_effective_date = formatDateString(anl_oil_distill.Row_effective_date)
	anl_oil_distill.Row_expiry_date = formatDateString(anl_oil_distill.Row_expiry_date)






	rows, err := stmt.Exec(anl_oil_distill.Anl_source, anl_oil_distill.Dstl_summary_obs_no, anl_oil_distill.Active_ind, anl_oil_distill.Atmosp_dstl_press, anl_oil_distill.Atmosp_dstl_press_ouom, anl_oil_distill.Atmosp_dstl_temp, anl_oil_distill.Atmosp_dstl_temp_ouom, anl_oil_distill.Calculated_ind, anl_oil_distill.Calculate_method_id, anl_oil_distill.Effective_date, anl_oil_distill.Expiry_date, anl_oil_distill.Final_boil_pt_temp, anl_oil_distill.Final_boil_pt_temp_ouom, anl_oil_distill.Measurement_temp, anl_oil_distill.Measurement_temp_ouom, anl_oil_distill.Ppdm_guid, anl_oil_distill.Preferred_ind, anl_oil_distill.Problem_ind, anl_oil_distill.Remark, anl_oil_distill.Reported_ind, anl_oil_distill.Source, anl_oil_distill.Start_boil_pt_temp, anl_oil_distill.Start_boil_pt_temp_ouom, anl_oil_distill.Step_seq_no, anl_oil_distill.Substance_id, anl_oil_distill.Volume_fraction_type, anl_oil_distill.Vol_fraction_temp, anl_oil_distill.Vol_fraction_temp_ouom, anl_oil_distill.Vol_fraction_value, anl_oil_distill.Vol_fraction_value_ouom, anl_oil_distill.Weight_cut, anl_oil_distill.Weight_cut_ouom, anl_oil_distill.Row_changed_by, anl_oil_distill.Row_changed_date, anl_oil_distill.Row_created_by, anl_oil_distill.Row_effective_date, anl_oil_distill.Row_expiry_date, anl_oil_distill.Row_quality, anl_oil_distill.Analysis_id)
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

func PatchAnlOilDistill(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_oil_distill set "
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

func DeleteAnlOilDistill(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_oil_distill dto.Anl_oil_distill
	anl_oil_distill.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_oil_distill where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_oil_distill.Analysis_id)
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


