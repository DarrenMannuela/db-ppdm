package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlQgfTsf(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_qgf_tsf")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_qgf_tsf

	for rows.Next() {
		var anl_qgf_tsf dto.Anl_qgf_tsf
		if err := rows.Scan(&anl_qgf_tsf.Analysis_id, &anl_qgf_tsf.Anl_source, &anl_qgf_tsf.Measurement_obs_no, &anl_qgf_tsf.Active_ind, &anl_qgf_tsf.Effective_date, &anl_qgf_tsf.Emission_intensity, &anl_qgf_tsf.Emission_intensity_uom, &anl_qgf_tsf.Expiry_date, &anl_qgf_tsf.Ppdm_guid, &anl_qgf_tsf.Preferred_ind, &anl_qgf_tsf.Problem_ind, &anl_qgf_tsf.Remark, &anl_qgf_tsf.Source, &anl_qgf_tsf.Step_seq_no, &anl_qgf_tsf.Wavelength_emission, &anl_qgf_tsf.Wavelength_emission_uom, &anl_qgf_tsf.Wavelenth_excitation, &anl_qgf_tsf.Wavelenth_excitation_uom, &anl_qgf_tsf.Row_changed_by, &anl_qgf_tsf.Row_changed_date, &anl_qgf_tsf.Row_created_by, &anl_qgf_tsf.Row_created_date, &anl_qgf_tsf.Row_effective_date, &anl_qgf_tsf.Row_expiry_date, &anl_qgf_tsf.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_qgf_tsf to result
		result = append(result, anl_qgf_tsf)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlQgfTsf(c *fiber.Ctx) error {
	var anl_qgf_tsf dto.Anl_qgf_tsf

	setDefaults(&anl_qgf_tsf)

	if err := json.Unmarshal(c.Body(), &anl_qgf_tsf); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_qgf_tsf values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25)")
	if err != nil {
		return err
	}
	anl_qgf_tsf.Row_created_date = formatDate(anl_qgf_tsf.Row_created_date)
	anl_qgf_tsf.Row_changed_date = formatDate(anl_qgf_tsf.Row_changed_date)
	anl_qgf_tsf.Effective_date = formatDateString(anl_qgf_tsf.Effective_date)
	anl_qgf_tsf.Expiry_date = formatDateString(anl_qgf_tsf.Expiry_date)
	anl_qgf_tsf.Row_effective_date = formatDateString(anl_qgf_tsf.Row_effective_date)
	anl_qgf_tsf.Row_expiry_date = formatDateString(anl_qgf_tsf.Row_expiry_date)






	rows, err := stmt.Exec(anl_qgf_tsf.Analysis_id, anl_qgf_tsf.Anl_source, anl_qgf_tsf.Measurement_obs_no, anl_qgf_tsf.Active_ind, anl_qgf_tsf.Effective_date, anl_qgf_tsf.Emission_intensity, anl_qgf_tsf.Emission_intensity_uom, anl_qgf_tsf.Expiry_date, anl_qgf_tsf.Ppdm_guid, anl_qgf_tsf.Preferred_ind, anl_qgf_tsf.Problem_ind, anl_qgf_tsf.Remark, anl_qgf_tsf.Source, anl_qgf_tsf.Step_seq_no, anl_qgf_tsf.Wavelength_emission, anl_qgf_tsf.Wavelength_emission_uom, anl_qgf_tsf.Wavelenth_excitation, anl_qgf_tsf.Wavelenth_excitation_uom, anl_qgf_tsf.Row_changed_by, anl_qgf_tsf.Row_changed_date, anl_qgf_tsf.Row_created_by, anl_qgf_tsf.Row_created_date, anl_qgf_tsf.Row_effective_date, anl_qgf_tsf.Row_expiry_date, anl_qgf_tsf.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlQgfTsf(c *fiber.Ctx) error {
	var anl_qgf_tsf dto.Anl_qgf_tsf

	setDefaults(&anl_qgf_tsf)

	if err := json.Unmarshal(c.Body(), &anl_qgf_tsf); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_qgf_tsf.Analysis_id = id

    if anl_qgf_tsf.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_qgf_tsf set anl_source = :1, measurement_obs_no = :2, active_ind = :3, effective_date = :4, emission_intensity = :5, emission_intensity_uom = :6, expiry_date = :7, ppdm_guid = :8, preferred_ind = :9, problem_ind = :10, remark = :11, source = :12, step_seq_no = :13, wavelength_emission = :14, wavelength_emission_uom = :15, wavelenth_excitation = :16, wavelenth_excitation_uom = :17, row_changed_by = :18, row_changed_date = :19, row_created_by = :20, row_effective_date = :21, row_expiry_date = :22, row_quality = :23 where analysis_id = :25")
	if err != nil {
		return err
	}

	anl_qgf_tsf.Row_changed_date = formatDate(anl_qgf_tsf.Row_changed_date)
	anl_qgf_tsf.Effective_date = formatDateString(anl_qgf_tsf.Effective_date)
	anl_qgf_tsf.Expiry_date = formatDateString(anl_qgf_tsf.Expiry_date)
	anl_qgf_tsf.Row_effective_date = formatDateString(anl_qgf_tsf.Row_effective_date)
	anl_qgf_tsf.Row_expiry_date = formatDateString(anl_qgf_tsf.Row_expiry_date)






	rows, err := stmt.Exec(anl_qgf_tsf.Anl_source, anl_qgf_tsf.Measurement_obs_no, anl_qgf_tsf.Active_ind, anl_qgf_tsf.Effective_date, anl_qgf_tsf.Emission_intensity, anl_qgf_tsf.Emission_intensity_uom, anl_qgf_tsf.Expiry_date, anl_qgf_tsf.Ppdm_guid, anl_qgf_tsf.Preferred_ind, anl_qgf_tsf.Problem_ind, anl_qgf_tsf.Remark, anl_qgf_tsf.Source, anl_qgf_tsf.Step_seq_no, anl_qgf_tsf.Wavelength_emission, anl_qgf_tsf.Wavelength_emission_uom, anl_qgf_tsf.Wavelenth_excitation, anl_qgf_tsf.Wavelenth_excitation_uom, anl_qgf_tsf.Row_changed_by, anl_qgf_tsf.Row_changed_date, anl_qgf_tsf.Row_created_by, anl_qgf_tsf.Row_effective_date, anl_qgf_tsf.Row_expiry_date, anl_qgf_tsf.Row_quality, anl_qgf_tsf.Analysis_id)
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

func PatchAnlQgfTsf(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_qgf_tsf set "
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

func DeleteAnlQgfTsf(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_qgf_tsf dto.Anl_qgf_tsf
	anl_qgf_tsf.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_qgf_tsf where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_qgf_tsf.Analysis_id)
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


