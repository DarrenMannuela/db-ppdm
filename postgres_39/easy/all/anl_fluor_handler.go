package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlFluor(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_fluor")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_fluor

	for rows.Next() {
		var anl_fluor dto.Anl_fluor
		if err := rows.Scan(&anl_fluor.Analysis_id, &anl_fluor.Anl_source, &anl_fluor.Fluor_obs_no, &anl_fluor.Active_ind, &anl_fluor.Color_remark, &anl_fluor.Effective_date, &anl_fluor.Expiry_date, &anl_fluor.Fluor_percent, &anl_fluor.Fluor_remark, &anl_fluor.From_color, &anl_fluor.From_color_frequency, &anl_fluor.From_color_frequency_uom, &anl_fluor.From_color_intensity, &anl_fluor.From_color_intensity_uom, &anl_fluor.From_mobility_type, &anl_fluor.Hydrocarbon_show_ind, &anl_fluor.Ppdm_guid, &anl_fluor.Preferred_ind, &anl_fluor.Problem_ind, &anl_fluor.Remark, &anl_fluor.Source, &anl_fluor.Step_seq_no, &anl_fluor.To_color, &anl_fluor.To_color_frequency, &anl_fluor.To_color_frequency_uom, &anl_fluor.To_color_intensity, &anl_fluor.To_color_intensity_uom, &anl_fluor.To_mobility_type, &anl_fluor.Row_changed_by, &anl_fluor.Row_changed_date, &anl_fluor.Row_created_by, &anl_fluor.Row_created_date, &anl_fluor.Row_effective_date, &anl_fluor.Row_expiry_date, &anl_fluor.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_fluor to result
		result = append(result, anl_fluor)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlFluor(c *fiber.Ctx) error {
	var anl_fluor dto.Anl_fluor

	setDefaults(&anl_fluor)

	if err := json.Unmarshal(c.Body(), &anl_fluor); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_fluor values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35)")
	if err != nil {
		return err
	}
	anl_fluor.Row_created_date = formatDate(anl_fluor.Row_created_date)
	anl_fluor.Row_changed_date = formatDate(anl_fluor.Row_changed_date)
	anl_fluor.Effective_date = formatDateString(anl_fluor.Effective_date)
	anl_fluor.Expiry_date = formatDateString(anl_fluor.Expiry_date)
	anl_fluor.Row_effective_date = formatDateString(anl_fluor.Row_effective_date)
	anl_fluor.Row_expiry_date = formatDateString(anl_fluor.Row_expiry_date)






	rows, err := stmt.Exec(anl_fluor.Analysis_id, anl_fluor.Anl_source, anl_fluor.Fluor_obs_no, anl_fluor.Active_ind, anl_fluor.Color_remark, anl_fluor.Effective_date, anl_fluor.Expiry_date, anl_fluor.Fluor_percent, anl_fluor.Fluor_remark, anl_fluor.From_color, anl_fluor.From_color_frequency, anl_fluor.From_color_frequency_uom, anl_fluor.From_color_intensity, anl_fluor.From_color_intensity_uom, anl_fluor.From_mobility_type, anl_fluor.Hydrocarbon_show_ind, anl_fluor.Ppdm_guid, anl_fluor.Preferred_ind, anl_fluor.Problem_ind, anl_fluor.Remark, anl_fluor.Source, anl_fluor.Step_seq_no, anl_fluor.To_color, anl_fluor.To_color_frequency, anl_fluor.To_color_frequency_uom, anl_fluor.To_color_intensity, anl_fluor.To_color_intensity_uom, anl_fluor.To_mobility_type, anl_fluor.Row_changed_by, anl_fluor.Row_changed_date, anl_fluor.Row_created_by, anl_fluor.Row_created_date, anl_fluor.Row_effective_date, anl_fluor.Row_expiry_date, anl_fluor.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlFluor(c *fiber.Ctx) error {
	var anl_fluor dto.Anl_fluor

	setDefaults(&anl_fluor)

	if err := json.Unmarshal(c.Body(), &anl_fluor); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_fluor.Analysis_id = id

    if anl_fluor.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_fluor set anl_source = :1, fluor_obs_no = :2, active_ind = :3, color_remark = :4, effective_date = :5, expiry_date = :6, fluor_percent = :7, fluor_remark = :8, from_color = :9, from_color_frequency = :10, from_color_frequency_uom = :11, from_color_intensity = :12, from_color_intensity_uom = :13, from_mobility_type = :14, hydrocarbon_show_ind = :15, ppdm_guid = :16, preferred_ind = :17, problem_ind = :18, remark = :19, source = :20, step_seq_no = :21, to_color = :22, to_color_frequency = :23, to_color_frequency_uom = :24, to_color_intensity = :25, to_color_intensity_uom = :26, to_mobility_type = :27, row_changed_by = :28, row_changed_date = :29, row_created_by = :30, row_effective_date = :31, row_expiry_date = :32, row_quality = :33 where analysis_id = :35")
	if err != nil {
		return err
	}

	anl_fluor.Row_changed_date = formatDate(anl_fluor.Row_changed_date)
	anl_fluor.Effective_date = formatDateString(anl_fluor.Effective_date)
	anl_fluor.Expiry_date = formatDateString(anl_fluor.Expiry_date)
	anl_fluor.Row_effective_date = formatDateString(anl_fluor.Row_effective_date)
	anl_fluor.Row_expiry_date = formatDateString(anl_fluor.Row_expiry_date)






	rows, err := stmt.Exec(anl_fluor.Anl_source, anl_fluor.Fluor_obs_no, anl_fluor.Active_ind, anl_fluor.Color_remark, anl_fluor.Effective_date, anl_fluor.Expiry_date, anl_fluor.Fluor_percent, anl_fluor.Fluor_remark, anl_fluor.From_color, anl_fluor.From_color_frequency, anl_fluor.From_color_frequency_uom, anl_fluor.From_color_intensity, anl_fluor.From_color_intensity_uom, anl_fluor.From_mobility_type, anl_fluor.Hydrocarbon_show_ind, anl_fluor.Ppdm_guid, anl_fluor.Preferred_ind, anl_fluor.Problem_ind, anl_fluor.Remark, anl_fluor.Source, anl_fluor.Step_seq_no, anl_fluor.To_color, anl_fluor.To_color_frequency, anl_fluor.To_color_frequency_uom, anl_fluor.To_color_intensity, anl_fluor.To_color_intensity_uom, anl_fluor.To_mobility_type, anl_fluor.Row_changed_by, anl_fluor.Row_changed_date, anl_fluor.Row_created_by, anl_fluor.Row_effective_date, anl_fluor.Row_expiry_date, anl_fluor.Row_quality, anl_fluor.Analysis_id)
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

func PatchAnlFluor(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_fluor set "
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

func DeleteAnlFluor(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_fluor dto.Anl_fluor
	anl_fluor.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_fluor where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_fluor.Analysis_id)
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


