package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDrillShaker(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_drill_shaker")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_drill_shaker

	for rows.Next() {
		var well_drill_shaker dto.Well_drill_shaker
		if err := rows.Scan(&well_drill_shaker.Uwi, &well_drill_shaker.Period_obs_no, &well_drill_shaker.Shaker_id, &well_drill_shaker.Screen_obs_no, &well_drill_shaker.Active_ind, &well_drill_shaker.Catalogue_equip_id, &well_drill_shaker.Effective_date, &well_drill_shaker.Equipment_id, &well_drill_shaker.Expiry_date, &well_drill_shaker.New_screen_ind, &well_drill_shaker.Ppdm_guid, &well_drill_shaker.Remark, &well_drill_shaker.Screen_change_ind, &well_drill_shaker.Screen_location, &well_drill_shaker.Screen_mesh_desc, &well_drill_shaker.Source, &well_drill_shaker.Row_changed_by, &well_drill_shaker.Row_changed_date, &well_drill_shaker.Row_created_by, &well_drill_shaker.Row_created_date, &well_drill_shaker.Row_effective_date, &well_drill_shaker.Row_expiry_date, &well_drill_shaker.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_drill_shaker to result
		result = append(result, well_drill_shaker)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDrillShaker(c *fiber.Ctx) error {
	var well_drill_shaker dto.Well_drill_shaker

	setDefaults(&well_drill_shaker)

	if err := json.Unmarshal(c.Body(), &well_drill_shaker); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_drill_shaker values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	well_drill_shaker.Row_created_date = formatDate(well_drill_shaker.Row_created_date)
	well_drill_shaker.Row_changed_date = formatDate(well_drill_shaker.Row_changed_date)
	well_drill_shaker.Effective_date = formatDateString(well_drill_shaker.Effective_date)
	well_drill_shaker.Expiry_date = formatDateString(well_drill_shaker.Expiry_date)
	well_drill_shaker.Row_effective_date = formatDateString(well_drill_shaker.Row_effective_date)
	well_drill_shaker.Row_expiry_date = formatDateString(well_drill_shaker.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_shaker.Uwi, well_drill_shaker.Period_obs_no, well_drill_shaker.Shaker_id, well_drill_shaker.Screen_obs_no, well_drill_shaker.Active_ind, well_drill_shaker.Catalogue_equip_id, well_drill_shaker.Effective_date, well_drill_shaker.Equipment_id, well_drill_shaker.Expiry_date, well_drill_shaker.New_screen_ind, well_drill_shaker.Ppdm_guid, well_drill_shaker.Remark, well_drill_shaker.Screen_change_ind, well_drill_shaker.Screen_location, well_drill_shaker.Screen_mesh_desc, well_drill_shaker.Source, well_drill_shaker.Row_changed_by, well_drill_shaker.Row_changed_date, well_drill_shaker.Row_created_by, well_drill_shaker.Row_created_date, well_drill_shaker.Row_effective_date, well_drill_shaker.Row_expiry_date, well_drill_shaker.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDrillShaker(c *fiber.Ctx) error {
	var well_drill_shaker dto.Well_drill_shaker

	setDefaults(&well_drill_shaker)

	if err := json.Unmarshal(c.Body(), &well_drill_shaker); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_drill_shaker.Uwi = id

    if well_drill_shaker.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_drill_shaker set period_obs_no = :1, shaker_id = :2, screen_obs_no = :3, active_ind = :4, catalogue_equip_id = :5, effective_date = :6, equipment_id = :7, expiry_date = :8, new_screen_ind = :9, ppdm_guid = :10, remark = :11, screen_change_ind = :12, screen_location = :13, screen_mesh_desc = :14, source = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where uwi = :23")
	if err != nil {
		return err
	}

	well_drill_shaker.Row_changed_date = formatDate(well_drill_shaker.Row_changed_date)
	well_drill_shaker.Effective_date = formatDateString(well_drill_shaker.Effective_date)
	well_drill_shaker.Expiry_date = formatDateString(well_drill_shaker.Expiry_date)
	well_drill_shaker.Row_effective_date = formatDateString(well_drill_shaker.Row_effective_date)
	well_drill_shaker.Row_expiry_date = formatDateString(well_drill_shaker.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_shaker.Period_obs_no, well_drill_shaker.Shaker_id, well_drill_shaker.Screen_obs_no, well_drill_shaker.Active_ind, well_drill_shaker.Catalogue_equip_id, well_drill_shaker.Effective_date, well_drill_shaker.Equipment_id, well_drill_shaker.Expiry_date, well_drill_shaker.New_screen_ind, well_drill_shaker.Ppdm_guid, well_drill_shaker.Remark, well_drill_shaker.Screen_change_ind, well_drill_shaker.Screen_location, well_drill_shaker.Screen_mesh_desc, well_drill_shaker.Source, well_drill_shaker.Row_changed_by, well_drill_shaker.Row_changed_date, well_drill_shaker.Row_created_by, well_drill_shaker.Row_effective_date, well_drill_shaker.Row_expiry_date, well_drill_shaker.Row_quality, well_drill_shaker.Uwi)
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

func PatchWellDrillShaker(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_drill_shaker set "
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

func DeleteWellDrillShaker(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_drill_shaker dto.Well_drill_shaker
	well_drill_shaker.Uwi = id

	stmt, err := db.Prepare("delete from well_drill_shaker where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_drill_shaker.Uwi)
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


