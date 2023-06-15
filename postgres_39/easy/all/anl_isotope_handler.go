package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlIsotope(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_isotope")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_isotope

	for rows.Next() {
		var anl_isotope dto.Anl_isotope
		if err := rows.Scan(&anl_isotope.Analysis_id, &anl_isotope.Anl_source, &anl_isotope.Isotope_obs_no, &anl_isotope.Active_ind, &anl_isotope.Calculate_method_id, &anl_isotope.Delta_notation_ind, &anl_isotope.Effective_date, &anl_isotope.Expiry_date, &anl_isotope.Isotope_value, &anl_isotope.Isotope_value_ouom, &anl_isotope.Isotope_value_uom, &anl_isotope.Measurement_type, &anl_isotope.Ppdm_guid, &anl_isotope.Preferred_ind, &anl_isotope.Problem_ind, &anl_isotope.Remark, &anl_isotope.Source, &anl_isotope.Standard_id, &anl_isotope.Step_seq_no, &anl_isotope.Substance_id, &anl_isotope.Row_changed_by, &anl_isotope.Row_changed_date, &anl_isotope.Row_created_by, &anl_isotope.Row_created_date, &anl_isotope.Row_effective_date, &anl_isotope.Row_expiry_date, &anl_isotope.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_isotope to result
		result = append(result, anl_isotope)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlIsotope(c *fiber.Ctx) error {
	var anl_isotope dto.Anl_isotope

	setDefaults(&anl_isotope)

	if err := json.Unmarshal(c.Body(), &anl_isotope); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_isotope values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	anl_isotope.Row_created_date = formatDate(anl_isotope.Row_created_date)
	anl_isotope.Row_changed_date = formatDate(anl_isotope.Row_changed_date)
	anl_isotope.Effective_date = formatDateString(anl_isotope.Effective_date)
	anl_isotope.Expiry_date = formatDateString(anl_isotope.Expiry_date)
	anl_isotope.Row_effective_date = formatDateString(anl_isotope.Row_effective_date)
	anl_isotope.Row_expiry_date = formatDateString(anl_isotope.Row_expiry_date)






	rows, err := stmt.Exec(anl_isotope.Analysis_id, anl_isotope.Anl_source, anl_isotope.Isotope_obs_no, anl_isotope.Active_ind, anl_isotope.Calculate_method_id, anl_isotope.Delta_notation_ind, anl_isotope.Effective_date, anl_isotope.Expiry_date, anl_isotope.Isotope_value, anl_isotope.Isotope_value_ouom, anl_isotope.Isotope_value_uom, anl_isotope.Measurement_type, anl_isotope.Ppdm_guid, anl_isotope.Preferred_ind, anl_isotope.Problem_ind, anl_isotope.Remark, anl_isotope.Source, anl_isotope.Standard_id, anl_isotope.Step_seq_no, anl_isotope.Substance_id, anl_isotope.Row_changed_by, anl_isotope.Row_changed_date, anl_isotope.Row_created_by, anl_isotope.Row_created_date, anl_isotope.Row_effective_date, anl_isotope.Row_expiry_date, anl_isotope.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlIsotope(c *fiber.Ctx) error {
	var anl_isotope dto.Anl_isotope

	setDefaults(&anl_isotope)

	if err := json.Unmarshal(c.Body(), &anl_isotope); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_isotope.Analysis_id = id

    if anl_isotope.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_isotope set anl_source = :1, isotope_obs_no = :2, active_ind = :3, calculate_method_id = :4, delta_notation_ind = :5, effective_date = :6, expiry_date = :7, isotope_value = :8, isotope_value_ouom = :9, isotope_value_uom = :10, measurement_type = :11, ppdm_guid = :12, preferred_ind = :13, problem_ind = :14, remark = :15, source = :16, standard_id = :17, step_seq_no = :18, substance_id = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where analysis_id = :27")
	if err != nil {
		return err
	}

	anl_isotope.Row_changed_date = formatDate(anl_isotope.Row_changed_date)
	anl_isotope.Effective_date = formatDateString(anl_isotope.Effective_date)
	anl_isotope.Expiry_date = formatDateString(anl_isotope.Expiry_date)
	anl_isotope.Row_effective_date = formatDateString(anl_isotope.Row_effective_date)
	anl_isotope.Row_expiry_date = formatDateString(anl_isotope.Row_expiry_date)






	rows, err := stmt.Exec(anl_isotope.Anl_source, anl_isotope.Isotope_obs_no, anl_isotope.Active_ind, anl_isotope.Calculate_method_id, anl_isotope.Delta_notation_ind, anl_isotope.Effective_date, anl_isotope.Expiry_date, anl_isotope.Isotope_value, anl_isotope.Isotope_value_ouom, anl_isotope.Isotope_value_uom, anl_isotope.Measurement_type, anl_isotope.Ppdm_guid, anl_isotope.Preferred_ind, anl_isotope.Problem_ind, anl_isotope.Remark, anl_isotope.Source, anl_isotope.Standard_id, anl_isotope.Step_seq_no, anl_isotope.Substance_id, anl_isotope.Row_changed_by, anl_isotope.Row_changed_date, anl_isotope.Row_created_by, anl_isotope.Row_effective_date, anl_isotope.Row_expiry_date, anl_isotope.Row_quality, anl_isotope.Analysis_id)
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

func PatchAnlIsotope(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_isotope set "
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

func DeleteAnlIsotope(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_isotope dto.Anl_isotope
	anl_isotope.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_isotope where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_isotope.Analysis_id)
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


