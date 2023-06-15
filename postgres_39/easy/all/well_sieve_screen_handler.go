package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellSieveScreen(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_sieve_screen")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_sieve_screen

	for rows.Next() {
		var well_sieve_screen dto.Well_sieve_screen
		if err := rows.Scan(&well_sieve_screen.Uwi, &well_sieve_screen.Anl_source, &well_sieve_screen.Analysis_obs_no, &well_sieve_screen.Screen_obs_no, &well_sieve_screen.Active_ind, &well_sieve_screen.Cat_equip_id, &well_sieve_screen.Effective_date, &well_sieve_screen.Equipment_id, &well_sieve_screen.Expiry_date, &well_sieve_screen.Particle_held_count, &well_sieve_screen.Ppdm_guid, &well_sieve_screen.Remark, &well_sieve_screen.Screen_mesh_size, &well_sieve_screen.Row_changed_by, &well_sieve_screen.Row_changed_date, &well_sieve_screen.Row_created_by, &well_sieve_screen.Row_created_date, &well_sieve_screen.Row_effective_date, &well_sieve_screen.Row_expiry_date, &well_sieve_screen.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_sieve_screen to result
		result = append(result, well_sieve_screen)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellSieveScreen(c *fiber.Ctx) error {
	var well_sieve_screen dto.Well_sieve_screen

	setDefaults(&well_sieve_screen)

	if err := json.Unmarshal(c.Body(), &well_sieve_screen); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_sieve_screen values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	well_sieve_screen.Row_created_date = formatDate(well_sieve_screen.Row_created_date)
	well_sieve_screen.Row_changed_date = formatDate(well_sieve_screen.Row_changed_date)
	well_sieve_screen.Effective_date = formatDateString(well_sieve_screen.Effective_date)
	well_sieve_screen.Expiry_date = formatDateString(well_sieve_screen.Expiry_date)
	well_sieve_screen.Row_effective_date = formatDateString(well_sieve_screen.Row_effective_date)
	well_sieve_screen.Row_expiry_date = formatDateString(well_sieve_screen.Row_expiry_date)






	rows, err := stmt.Exec(well_sieve_screen.Uwi, well_sieve_screen.Anl_source, well_sieve_screen.Analysis_obs_no, well_sieve_screen.Screen_obs_no, well_sieve_screen.Active_ind, well_sieve_screen.Cat_equip_id, well_sieve_screen.Effective_date, well_sieve_screen.Equipment_id, well_sieve_screen.Expiry_date, well_sieve_screen.Particle_held_count, well_sieve_screen.Ppdm_guid, well_sieve_screen.Remark, well_sieve_screen.Screen_mesh_size, well_sieve_screen.Row_changed_by, well_sieve_screen.Row_changed_date, well_sieve_screen.Row_created_by, well_sieve_screen.Row_created_date, well_sieve_screen.Row_effective_date, well_sieve_screen.Row_expiry_date, well_sieve_screen.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellSieveScreen(c *fiber.Ctx) error {
	var well_sieve_screen dto.Well_sieve_screen

	setDefaults(&well_sieve_screen)

	if err := json.Unmarshal(c.Body(), &well_sieve_screen); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_sieve_screen.Uwi = id

    if well_sieve_screen.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_sieve_screen set anl_source = :1, analysis_obs_no = :2, screen_obs_no = :3, active_ind = :4, cat_equip_id = :5, effective_date = :6, equipment_id = :7, expiry_date = :8, particle_held_count = :9, ppdm_guid = :10, remark = :11, screen_mesh_size = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where uwi = :20")
	if err != nil {
		return err
	}

	well_sieve_screen.Row_changed_date = formatDate(well_sieve_screen.Row_changed_date)
	well_sieve_screen.Effective_date = formatDateString(well_sieve_screen.Effective_date)
	well_sieve_screen.Expiry_date = formatDateString(well_sieve_screen.Expiry_date)
	well_sieve_screen.Row_effective_date = formatDateString(well_sieve_screen.Row_effective_date)
	well_sieve_screen.Row_expiry_date = formatDateString(well_sieve_screen.Row_expiry_date)






	rows, err := stmt.Exec(well_sieve_screen.Anl_source, well_sieve_screen.Analysis_obs_no, well_sieve_screen.Screen_obs_no, well_sieve_screen.Active_ind, well_sieve_screen.Cat_equip_id, well_sieve_screen.Effective_date, well_sieve_screen.Equipment_id, well_sieve_screen.Expiry_date, well_sieve_screen.Particle_held_count, well_sieve_screen.Ppdm_guid, well_sieve_screen.Remark, well_sieve_screen.Screen_mesh_size, well_sieve_screen.Row_changed_by, well_sieve_screen.Row_changed_date, well_sieve_screen.Row_created_by, well_sieve_screen.Row_effective_date, well_sieve_screen.Row_expiry_date, well_sieve_screen.Row_quality, well_sieve_screen.Uwi)
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

func PatchWellSieveScreen(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_sieve_screen set "
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

func DeleteWellSieveScreen(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_sieve_screen dto.Well_sieve_screen
	well_sieve_screen.Uwi = id

	stmt, err := db.Prepare("delete from well_sieve_screen where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_sieve_screen.Uwi)
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


