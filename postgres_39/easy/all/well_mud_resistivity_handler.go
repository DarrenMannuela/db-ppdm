package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellMudResistivity(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_mud_resistivity")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_mud_resistivity

	for rows.Next() {
		var well_mud_resistivity dto.Well_mud_resistivity
		if err := rows.Scan(&well_mud_resistivity.Uwi, &well_mud_resistivity.Source, &well_mud_resistivity.Sample_id, &well_mud_resistivity.Resistivity_obs_no, &well_mud_resistivity.Active_ind, &well_mud_resistivity.Effective_date, &well_mud_resistivity.Expiry_date, &well_mud_resistivity.Mud_resistivity, &well_mud_resistivity.Mud_resistivity_ouom, &well_mud_resistivity.Ppdm_guid, &well_mud_resistivity.Remark, &well_mud_resistivity.Resistivity_temperature, &well_mud_resistivity.Resistivity_temperature_ouom, &well_mud_resistivity.Sample_type, &well_mud_resistivity.Row_changed_by, &well_mud_resistivity.Row_changed_date, &well_mud_resistivity.Row_created_by, &well_mud_resistivity.Row_created_date, &well_mud_resistivity.Row_effective_date, &well_mud_resistivity.Row_expiry_date, &well_mud_resistivity.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_mud_resistivity to result
		result = append(result, well_mud_resistivity)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellMudResistivity(c *fiber.Ctx) error {
	var well_mud_resistivity dto.Well_mud_resistivity

	setDefaults(&well_mud_resistivity)

	if err := json.Unmarshal(c.Body(), &well_mud_resistivity); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_mud_resistivity values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	well_mud_resistivity.Row_created_date = formatDate(well_mud_resistivity.Row_created_date)
	well_mud_resistivity.Row_changed_date = formatDate(well_mud_resistivity.Row_changed_date)
	well_mud_resistivity.Effective_date = formatDateString(well_mud_resistivity.Effective_date)
	well_mud_resistivity.Expiry_date = formatDateString(well_mud_resistivity.Expiry_date)
	well_mud_resistivity.Row_effective_date = formatDateString(well_mud_resistivity.Row_effective_date)
	well_mud_resistivity.Row_expiry_date = formatDateString(well_mud_resistivity.Row_expiry_date)






	rows, err := stmt.Exec(well_mud_resistivity.Uwi, well_mud_resistivity.Source, well_mud_resistivity.Sample_id, well_mud_resistivity.Resistivity_obs_no, well_mud_resistivity.Active_ind, well_mud_resistivity.Effective_date, well_mud_resistivity.Expiry_date, well_mud_resistivity.Mud_resistivity, well_mud_resistivity.Mud_resistivity_ouom, well_mud_resistivity.Ppdm_guid, well_mud_resistivity.Remark, well_mud_resistivity.Resistivity_temperature, well_mud_resistivity.Resistivity_temperature_ouom, well_mud_resistivity.Sample_type, well_mud_resistivity.Row_changed_by, well_mud_resistivity.Row_changed_date, well_mud_resistivity.Row_created_by, well_mud_resistivity.Row_created_date, well_mud_resistivity.Row_effective_date, well_mud_resistivity.Row_expiry_date, well_mud_resistivity.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellMudResistivity(c *fiber.Ctx) error {
	var well_mud_resistivity dto.Well_mud_resistivity

	setDefaults(&well_mud_resistivity)

	if err := json.Unmarshal(c.Body(), &well_mud_resistivity); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_mud_resistivity.Uwi = id

    if well_mud_resistivity.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_mud_resistivity set source = :1, sample_id = :2, resistivity_obs_no = :3, active_ind = :4, effective_date = :5, expiry_date = :6, mud_resistivity = :7, mud_resistivity_ouom = :8, ppdm_guid = :9, remark = :10, resistivity_temperature = :11, resistivity_temperature_ouom = :12, sample_type = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where uwi = :21")
	if err != nil {
		return err
	}

	well_mud_resistivity.Row_changed_date = formatDate(well_mud_resistivity.Row_changed_date)
	well_mud_resistivity.Effective_date = formatDateString(well_mud_resistivity.Effective_date)
	well_mud_resistivity.Expiry_date = formatDateString(well_mud_resistivity.Expiry_date)
	well_mud_resistivity.Row_effective_date = formatDateString(well_mud_resistivity.Row_effective_date)
	well_mud_resistivity.Row_expiry_date = formatDateString(well_mud_resistivity.Row_expiry_date)






	rows, err := stmt.Exec(well_mud_resistivity.Source, well_mud_resistivity.Sample_id, well_mud_resistivity.Resistivity_obs_no, well_mud_resistivity.Active_ind, well_mud_resistivity.Effective_date, well_mud_resistivity.Expiry_date, well_mud_resistivity.Mud_resistivity, well_mud_resistivity.Mud_resistivity_ouom, well_mud_resistivity.Ppdm_guid, well_mud_resistivity.Remark, well_mud_resistivity.Resistivity_temperature, well_mud_resistivity.Resistivity_temperature_ouom, well_mud_resistivity.Sample_type, well_mud_resistivity.Row_changed_by, well_mud_resistivity.Row_changed_date, well_mud_resistivity.Row_created_by, well_mud_resistivity.Row_effective_date, well_mud_resistivity.Row_expiry_date, well_mud_resistivity.Row_quality, well_mud_resistivity.Uwi)
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

func PatchWellMudResistivity(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_mud_resistivity set "
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
	query += " where uwi = :id"

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

func DeleteWellMudResistivity(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_mud_resistivity dto.Well_mud_resistivity
	well_mud_resistivity.Uwi = id

	stmt, err := db.Prepare("delete from well_mud_resistivity where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_mud_resistivity.Uwi)
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


