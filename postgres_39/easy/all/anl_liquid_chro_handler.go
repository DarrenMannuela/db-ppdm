package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlLiquidChro(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_liquid_chro")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_liquid_chro

	for rows.Next() {
		var anl_liquid_chro dto.Anl_liquid_chro
		if err := rows.Scan(&anl_liquid_chro.Analysis_id, &anl_liquid_chro.Anl_source, &anl_liquid_chro.Chro_obs_no, &anl_liquid_chro.Active_ind, &anl_liquid_chro.Aged_oil_used, &anl_liquid_chro.Aged_oil_used_uom, &anl_liquid_chro.Aged_weight_of_oil, &anl_liquid_chro.Aged_weight_of_oil_uom, &anl_liquid_chro.Aromatic_hc_value, &anl_liquid_chro.Aromatic_hc_value_uom, &anl_liquid_chro.Asphaltene_free_value, &anl_liquid_chro.Asphaltene_free_value_uom, &anl_liquid_chro.Asphaltene_value, &anl_liquid_chro.Asphaltene_value_uom, &anl_liquid_chro.Effective_date, &anl_liquid_chro.Eom_used, &anl_liquid_chro.Eom_used_uom, &anl_liquid_chro.Eom_value, &anl_liquid_chro.Eom_value_uom, &anl_liquid_chro.Expiry_date, &anl_liquid_chro.Loss_on_column, &anl_liquid_chro.Loss_on_column_uom, &anl_liquid_chro.Measurement_type, &anl_liquid_chro.Naphthene_value, &anl_liquid_chro.Naphthene_value_uom, &anl_liquid_chro.Nso_value, &anl_liquid_chro.Nso_value_uom, &anl_liquid_chro.N_alkane_fraction, &anl_liquid_chro.N_alkane_fraction_uom, &anl_liquid_chro.Ppdm_guid, &anl_liquid_chro.Preferred_ind, &anl_liquid_chro.Problem_ind, &anl_liquid_chro.Refraction_factor, &anl_liquid_chro.Remark, &anl_liquid_chro.Reported_hc_fraction, &anl_liquid_chro.Reported_hc_fraction_uom, &anl_liquid_chro.Reported_hc_in_toc, &anl_liquid_chro.Reported_hc_in_toc_uom, &anl_liquid_chro.Saturated_hc_value, &anl_liquid_chro.Saturated_hc_value_uom, &anl_liquid_chro.Source, &anl_liquid_chro.Step_seq_no, &anl_liquid_chro.Total_soluble_extract, &anl_liquid_chro.Total_soluble_extract_uom, &anl_liquid_chro.Row_changed_by, &anl_liquid_chro.Row_changed_date, &anl_liquid_chro.Row_created_by, &anl_liquid_chro.Row_created_date, &anl_liquid_chro.Row_effective_date, &anl_liquid_chro.Row_expiry_date, &anl_liquid_chro.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_liquid_chro to result
		result = append(result, anl_liquid_chro)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlLiquidChro(c *fiber.Ctx) error {
	var anl_liquid_chro dto.Anl_liquid_chro

	setDefaults(&anl_liquid_chro)

	if err := json.Unmarshal(c.Body(), &anl_liquid_chro); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_liquid_chro values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51)")
	if err != nil {
		return err
	}
	anl_liquid_chro.Row_created_date = formatDate(anl_liquid_chro.Row_created_date)
	anl_liquid_chro.Row_changed_date = formatDate(anl_liquid_chro.Row_changed_date)
	anl_liquid_chro.Effective_date = formatDateString(anl_liquid_chro.Effective_date)
	anl_liquid_chro.Expiry_date = formatDateString(anl_liquid_chro.Expiry_date)
	anl_liquid_chro.Row_effective_date = formatDateString(anl_liquid_chro.Row_effective_date)
	anl_liquid_chro.Row_expiry_date = formatDateString(anl_liquid_chro.Row_expiry_date)






	rows, err := stmt.Exec(anl_liquid_chro.Analysis_id, anl_liquid_chro.Anl_source, anl_liquid_chro.Chro_obs_no, anl_liquid_chro.Active_ind, anl_liquid_chro.Aged_oil_used, anl_liquid_chro.Aged_oil_used_uom, anl_liquid_chro.Aged_weight_of_oil, anl_liquid_chro.Aged_weight_of_oil_uom, anl_liquid_chro.Aromatic_hc_value, anl_liquid_chro.Aromatic_hc_value_uom, anl_liquid_chro.Asphaltene_free_value, anl_liquid_chro.Asphaltene_free_value_uom, anl_liquid_chro.Asphaltene_value, anl_liquid_chro.Asphaltene_value_uom, anl_liquid_chro.Effective_date, anl_liquid_chro.Eom_used, anl_liquid_chro.Eom_used_uom, anl_liquid_chro.Eom_value, anl_liquid_chro.Eom_value_uom, anl_liquid_chro.Expiry_date, anl_liquid_chro.Loss_on_column, anl_liquid_chro.Loss_on_column_uom, anl_liquid_chro.Measurement_type, anl_liquid_chro.Naphthene_value, anl_liquid_chro.Naphthene_value_uom, anl_liquid_chro.Nso_value, anl_liquid_chro.Nso_value_uom, anl_liquid_chro.N_alkane_fraction, anl_liquid_chro.N_alkane_fraction_uom, anl_liquid_chro.Ppdm_guid, anl_liquid_chro.Preferred_ind, anl_liquid_chro.Problem_ind, anl_liquid_chro.Refraction_factor, anl_liquid_chro.Remark, anl_liquid_chro.Reported_hc_fraction, anl_liquid_chro.Reported_hc_fraction_uom, anl_liquid_chro.Reported_hc_in_toc, anl_liquid_chro.Reported_hc_in_toc_uom, anl_liquid_chro.Saturated_hc_value, anl_liquid_chro.Saturated_hc_value_uom, anl_liquid_chro.Source, anl_liquid_chro.Step_seq_no, anl_liquid_chro.Total_soluble_extract, anl_liquid_chro.Total_soluble_extract_uom, anl_liquid_chro.Row_changed_by, anl_liquid_chro.Row_changed_date, anl_liquid_chro.Row_created_by, anl_liquid_chro.Row_created_date, anl_liquid_chro.Row_effective_date, anl_liquid_chro.Row_expiry_date, anl_liquid_chro.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlLiquidChro(c *fiber.Ctx) error {
	var anl_liquid_chro dto.Anl_liquid_chro

	setDefaults(&anl_liquid_chro)

	if err := json.Unmarshal(c.Body(), &anl_liquid_chro); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_liquid_chro.Analysis_id = id

    if anl_liquid_chro.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_liquid_chro set anl_source = :1, chro_obs_no = :2, active_ind = :3, aged_oil_used = :4, aged_oil_used_uom = :5, aged_weight_of_oil = :6, aged_weight_of_oil_uom = :7, aromatic_hc_value = :8, aromatic_hc_value_uom = :9, asphaltene_free_value = :10, asphaltene_free_value_uom = :11, asphaltene_value = :12, asphaltene_value_uom = :13, effective_date = :14, eom_used = :15, eom_used_uom = :16, eom_value = :17, eom_value_uom = :18, expiry_date = :19, loss_on_column = :20, loss_on_column_uom = :21, measurement_type = :22, naphthene_value = :23, naphthene_value_uom = :24, nso_value = :25, nso_value_uom = :26, n_alkane_fraction = :27, n_alkane_fraction_uom = :28, ppdm_guid = :29, preferred_ind = :30, problem_ind = :31, refraction_factor = :32, remark = :33, reported_hc_fraction = :34, reported_hc_fraction_uom = :35, reported_hc_in_toc = :36, reported_hc_in_toc_uom = :37, saturated_hc_value = :38, saturated_hc_value_uom = :39, source = :40, step_seq_no = :41, total_soluble_extract = :42, total_soluble_extract_uom = :43, row_changed_by = :44, row_changed_date = :45, row_created_by = :46, row_effective_date = :47, row_expiry_date = :48, row_quality = :49 where analysis_id = :51")
	if err != nil {
		return err
	}

	anl_liquid_chro.Row_changed_date = formatDate(anl_liquid_chro.Row_changed_date)
	anl_liquid_chro.Effective_date = formatDateString(anl_liquid_chro.Effective_date)
	anl_liquid_chro.Expiry_date = formatDateString(anl_liquid_chro.Expiry_date)
	anl_liquid_chro.Row_effective_date = formatDateString(anl_liquid_chro.Row_effective_date)
	anl_liquid_chro.Row_expiry_date = formatDateString(anl_liquid_chro.Row_expiry_date)






	rows, err := stmt.Exec(anl_liquid_chro.Anl_source, anl_liquid_chro.Chro_obs_no, anl_liquid_chro.Active_ind, anl_liquid_chro.Aged_oil_used, anl_liquid_chro.Aged_oil_used_uom, anl_liquid_chro.Aged_weight_of_oil, anl_liquid_chro.Aged_weight_of_oil_uom, anl_liquid_chro.Aromatic_hc_value, anl_liquid_chro.Aromatic_hc_value_uom, anl_liquid_chro.Asphaltene_free_value, anl_liquid_chro.Asphaltene_free_value_uom, anl_liquid_chro.Asphaltene_value, anl_liquid_chro.Asphaltene_value_uom, anl_liquid_chro.Effective_date, anl_liquid_chro.Eom_used, anl_liquid_chro.Eom_used_uom, anl_liquid_chro.Eom_value, anl_liquid_chro.Eom_value_uom, anl_liquid_chro.Expiry_date, anl_liquid_chro.Loss_on_column, anl_liquid_chro.Loss_on_column_uom, anl_liquid_chro.Measurement_type, anl_liquid_chro.Naphthene_value, anl_liquid_chro.Naphthene_value_uom, anl_liquid_chro.Nso_value, anl_liquid_chro.Nso_value_uom, anl_liquid_chro.N_alkane_fraction, anl_liquid_chro.N_alkane_fraction_uom, anl_liquid_chro.Ppdm_guid, anl_liquid_chro.Preferred_ind, anl_liquid_chro.Problem_ind, anl_liquid_chro.Refraction_factor, anl_liquid_chro.Remark, anl_liquid_chro.Reported_hc_fraction, anl_liquid_chro.Reported_hc_fraction_uom, anl_liquid_chro.Reported_hc_in_toc, anl_liquid_chro.Reported_hc_in_toc_uom, anl_liquid_chro.Saturated_hc_value, anl_liquid_chro.Saturated_hc_value_uom, anl_liquid_chro.Source, anl_liquid_chro.Step_seq_no, anl_liquid_chro.Total_soluble_extract, anl_liquid_chro.Total_soluble_extract_uom, anl_liquid_chro.Row_changed_by, anl_liquid_chro.Row_changed_date, anl_liquid_chro.Row_created_by, anl_liquid_chro.Row_effective_date, anl_liquid_chro.Row_expiry_date, anl_liquid_chro.Row_quality, anl_liquid_chro.Analysis_id)
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

func PatchAnlLiquidChro(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_liquid_chro set "
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

func DeleteAnlLiquidChro(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_liquid_chro dto.Anl_liquid_chro
	anl_liquid_chro.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_liquid_chro where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_liquid_chro.Analysis_id)
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


