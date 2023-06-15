package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlOilViscosity(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_oil_viscosity")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_oil_viscosity

	for rows.Next() {
		var anl_oil_viscosity dto.Anl_oil_viscosity
		if err := rows.Scan(&anl_oil_viscosity.Analysis_id, &anl_oil_viscosity.Anl_source, &anl_oil_viscosity.Viscosity_obs_no, &anl_oil_viscosity.Active_ind, &anl_oil_viscosity.Effective_date, &anl_oil_viscosity.Expiry_date, &anl_oil_viscosity.Oil_temperature, &anl_oil_viscosity.Oil_temperature_ouom, &anl_oil_viscosity.Oil_type, &anl_oil_viscosity.Oil_viscosity, &anl_oil_viscosity.Oil_viscosity_ouom, &anl_oil_viscosity.Oil_viscosity_uom, &anl_oil_viscosity.Pour_point_temperature, &anl_oil_viscosity.Pour_point_temperature_ouom, &anl_oil_viscosity.Ppdm_guid, &anl_oil_viscosity.Preferred_ind, &anl_oil_viscosity.Remark, &anl_oil_viscosity.Source, &anl_oil_viscosity.Step_seq_no, &anl_oil_viscosity.Viscosity_press, &anl_oil_viscosity.Viscosity_press_ouom, &anl_oil_viscosity.Viscosity_temp, &anl_oil_viscosity.Viscosity_temp_ouom, &anl_oil_viscosity.Viscosity_value, &anl_oil_viscosity.Viscosity_value_ouom, &anl_oil_viscosity.Viscosity_value_uom, &anl_oil_viscosity.Row_changed_by, &anl_oil_viscosity.Row_changed_date, &anl_oil_viscosity.Row_created_by, &anl_oil_viscosity.Row_created_date, &anl_oil_viscosity.Row_effective_date, &anl_oil_viscosity.Row_expiry_date, &anl_oil_viscosity.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_oil_viscosity to result
		result = append(result, anl_oil_viscosity)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlOilViscosity(c *fiber.Ctx) error {
	var anl_oil_viscosity dto.Anl_oil_viscosity

	setDefaults(&anl_oil_viscosity)

	if err := json.Unmarshal(c.Body(), &anl_oil_viscosity); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_oil_viscosity values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33)")
	if err != nil {
		return err
	}
	anl_oil_viscosity.Row_created_date = formatDate(anl_oil_viscosity.Row_created_date)
	anl_oil_viscosity.Row_changed_date = formatDate(anl_oil_viscosity.Row_changed_date)
	anl_oil_viscosity.Effective_date = formatDateString(anl_oil_viscosity.Effective_date)
	anl_oil_viscosity.Expiry_date = formatDateString(anl_oil_viscosity.Expiry_date)
	anl_oil_viscosity.Row_effective_date = formatDateString(anl_oil_viscosity.Row_effective_date)
	anl_oil_viscosity.Row_expiry_date = formatDateString(anl_oil_viscosity.Row_expiry_date)






	rows, err := stmt.Exec(anl_oil_viscosity.Analysis_id, anl_oil_viscosity.Anl_source, anl_oil_viscosity.Viscosity_obs_no, anl_oil_viscosity.Active_ind, anl_oil_viscosity.Effective_date, anl_oil_viscosity.Expiry_date, anl_oil_viscosity.Oil_temperature, anl_oil_viscosity.Oil_temperature_ouom, anl_oil_viscosity.Oil_type, anl_oil_viscosity.Oil_viscosity, anl_oil_viscosity.Oil_viscosity_ouom, anl_oil_viscosity.Oil_viscosity_uom, anl_oil_viscosity.Pour_point_temperature, anl_oil_viscosity.Pour_point_temperature_ouom, anl_oil_viscosity.Ppdm_guid, anl_oil_viscosity.Preferred_ind, anl_oil_viscosity.Remark, anl_oil_viscosity.Source, anl_oil_viscosity.Step_seq_no, anl_oil_viscosity.Viscosity_press, anl_oil_viscosity.Viscosity_press_ouom, anl_oil_viscosity.Viscosity_temp, anl_oil_viscosity.Viscosity_temp_ouom, anl_oil_viscosity.Viscosity_value, anl_oil_viscosity.Viscosity_value_ouom, anl_oil_viscosity.Viscosity_value_uom, anl_oil_viscosity.Row_changed_by, anl_oil_viscosity.Row_changed_date, anl_oil_viscosity.Row_created_by, anl_oil_viscosity.Row_created_date, anl_oil_viscosity.Row_effective_date, anl_oil_viscosity.Row_expiry_date, anl_oil_viscosity.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlOilViscosity(c *fiber.Ctx) error {
	var anl_oil_viscosity dto.Anl_oil_viscosity

	setDefaults(&anl_oil_viscosity)

	if err := json.Unmarshal(c.Body(), &anl_oil_viscosity); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_oil_viscosity.Analysis_id = id

    if anl_oil_viscosity.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_oil_viscosity set anl_source = :1, viscosity_obs_no = :2, active_ind = :3, effective_date = :4, expiry_date = :5, oil_temperature = :6, oil_temperature_ouom = :7, oil_type = :8, oil_viscosity = :9, oil_viscosity_ouom = :10, oil_viscosity_uom = :11, pour_point_temperature = :12, pour_point_temperature_ouom = :13, ppdm_guid = :14, preferred_ind = :15, remark = :16, source = :17, step_seq_no = :18, viscosity_press = :19, viscosity_press_ouom = :20, viscosity_temp = :21, viscosity_temp_ouom = :22, viscosity_value = :23, viscosity_value_ouom = :24, viscosity_value_uom = :25, row_changed_by = :26, row_changed_date = :27, row_created_by = :28, row_effective_date = :29, row_expiry_date = :30, row_quality = :31 where analysis_id = :33")
	if err != nil {
		return err
	}

	anl_oil_viscosity.Row_changed_date = formatDate(anl_oil_viscosity.Row_changed_date)
	anl_oil_viscosity.Effective_date = formatDateString(anl_oil_viscosity.Effective_date)
	anl_oil_viscosity.Expiry_date = formatDateString(anl_oil_viscosity.Expiry_date)
	anl_oil_viscosity.Row_effective_date = formatDateString(anl_oil_viscosity.Row_effective_date)
	anl_oil_viscosity.Row_expiry_date = formatDateString(anl_oil_viscosity.Row_expiry_date)






	rows, err := stmt.Exec(anl_oil_viscosity.Anl_source, anl_oil_viscosity.Viscosity_obs_no, anl_oil_viscosity.Active_ind, anl_oil_viscosity.Effective_date, anl_oil_viscosity.Expiry_date, anl_oil_viscosity.Oil_temperature, anl_oil_viscosity.Oil_temperature_ouom, anl_oil_viscosity.Oil_type, anl_oil_viscosity.Oil_viscosity, anl_oil_viscosity.Oil_viscosity_ouom, anl_oil_viscosity.Oil_viscosity_uom, anl_oil_viscosity.Pour_point_temperature, anl_oil_viscosity.Pour_point_temperature_ouom, anl_oil_viscosity.Ppdm_guid, anl_oil_viscosity.Preferred_ind, anl_oil_viscosity.Remark, anl_oil_viscosity.Source, anl_oil_viscosity.Step_seq_no, anl_oil_viscosity.Viscosity_press, anl_oil_viscosity.Viscosity_press_ouom, anl_oil_viscosity.Viscosity_temp, anl_oil_viscosity.Viscosity_temp_ouom, anl_oil_viscosity.Viscosity_value, anl_oil_viscosity.Viscosity_value_ouom, anl_oil_viscosity.Viscosity_value_uom, anl_oil_viscosity.Row_changed_by, anl_oil_viscosity.Row_changed_date, anl_oil_viscosity.Row_created_by, anl_oil_viscosity.Row_effective_date, anl_oil_viscosity.Row_expiry_date, anl_oil_viscosity.Row_quality, anl_oil_viscosity.Analysis_id)
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

func PatchAnlOilViscosity(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_oil_viscosity set "
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

func DeleteAnlOilViscosity(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_oil_viscosity dto.Anl_oil_viscosity
	anl_oil_viscosity.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_oil_viscosity where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_oil_viscosity.Analysis_id)
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


