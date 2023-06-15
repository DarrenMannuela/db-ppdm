package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellPlugback(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_plugback")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_plugback

	for rows.Next() {
		var well_plugback dto.Well_plugback
		if err := rows.Scan(&well_plugback.Uwi, &well_plugback.Source, &well_plugback.Plugback_obs_no, &well_plugback.Active_ind, &well_plugback.Base_depth, &well_plugback.Base_depth_ouom, &well_plugback.Base_strat_unit_id, &well_plugback.Cement_amount, &well_plugback.Cement_amount_ouom, &well_plugback.Cement_amount_uom, &well_plugback.Completion_obs_no, &well_plugback.Completion_source, &well_plugback.Effective_date, &well_plugback.Expiry_date, &well_plugback.Plugback_ba_id, &well_plugback.Plugback_date, &well_plugback.Plug_type, &well_plugback.Ppdm_guid, &well_plugback.Remark, &well_plugback.Strat_name_set_id, &well_plugback.Surface_abandon_date, &well_plugback.Top_depth, &well_plugback.Top_depth_ouom, &well_plugback.Top_found_depth, &well_plugback.Top_found_depth_ouom, &well_plugback.Top_strat_unit_id, &well_plugback.Row_changed_by, &well_plugback.Row_changed_date, &well_plugback.Row_created_by, &well_plugback.Row_created_date, &well_plugback.Row_effective_date, &well_plugback.Row_expiry_date, &well_plugback.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_plugback to result
		result = append(result, well_plugback)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellPlugback(c *fiber.Ctx) error {
	var well_plugback dto.Well_plugback

	setDefaults(&well_plugback)

	if err := json.Unmarshal(c.Body(), &well_plugback); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_plugback values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33)")
	if err != nil {
		return err
	}
	well_plugback.Row_created_date = formatDate(well_plugback.Row_created_date)
	well_plugback.Row_changed_date = formatDate(well_plugback.Row_changed_date)
	well_plugback.Effective_date = formatDateString(well_plugback.Effective_date)
	well_plugback.Expiry_date = formatDateString(well_plugback.Expiry_date)
	well_plugback.Plugback_date = formatDateString(well_plugback.Plugback_date)
	well_plugback.Surface_abandon_date = formatDateString(well_plugback.Surface_abandon_date)
	well_plugback.Row_effective_date = formatDateString(well_plugback.Row_effective_date)
	well_plugback.Row_expiry_date = formatDateString(well_plugback.Row_expiry_date)






	rows, err := stmt.Exec(well_plugback.Uwi, well_plugback.Source, well_plugback.Plugback_obs_no, well_plugback.Active_ind, well_plugback.Base_depth, well_plugback.Base_depth_ouom, well_plugback.Base_strat_unit_id, well_plugback.Cement_amount, well_plugback.Cement_amount_ouom, well_plugback.Cement_amount_uom, well_plugback.Completion_obs_no, well_plugback.Completion_source, well_plugback.Effective_date, well_plugback.Expiry_date, well_plugback.Plugback_ba_id, well_plugback.Plugback_date, well_plugback.Plug_type, well_plugback.Ppdm_guid, well_plugback.Remark, well_plugback.Strat_name_set_id, well_plugback.Surface_abandon_date, well_plugback.Top_depth, well_plugback.Top_depth_ouom, well_plugback.Top_found_depth, well_plugback.Top_found_depth_ouom, well_plugback.Top_strat_unit_id, well_plugback.Row_changed_by, well_plugback.Row_changed_date, well_plugback.Row_created_by, well_plugback.Row_created_date, well_plugback.Row_effective_date, well_plugback.Row_expiry_date, well_plugback.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellPlugback(c *fiber.Ctx) error {
	var well_plugback dto.Well_plugback

	setDefaults(&well_plugback)

	if err := json.Unmarshal(c.Body(), &well_plugback); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_plugback.Uwi = id

    if well_plugback.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_plugback set source = :1, plugback_obs_no = :2, active_ind = :3, base_depth = :4, base_depth_ouom = :5, base_strat_unit_id = :6, cement_amount = :7, cement_amount_ouom = :8, cement_amount_uom = :9, completion_obs_no = :10, completion_source = :11, effective_date = :12, expiry_date = :13, plugback_ba_id = :14, plugback_date = :15, plug_type = :16, ppdm_guid = :17, remark = :18, strat_name_set_id = :19, surface_abandon_date = :20, top_depth = :21, top_depth_ouom = :22, top_found_depth = :23, top_found_depth_ouom = :24, top_strat_unit_id = :25, row_changed_by = :26, row_changed_date = :27, row_created_by = :28, row_effective_date = :29, row_expiry_date = :30, row_quality = :31 where uwi = :33")
	if err != nil {
		return err
	}

	well_plugback.Row_changed_date = formatDate(well_plugback.Row_changed_date)
	well_plugback.Effective_date = formatDateString(well_plugback.Effective_date)
	well_plugback.Expiry_date = formatDateString(well_plugback.Expiry_date)
	well_plugback.Plugback_date = formatDateString(well_plugback.Plugback_date)
	well_plugback.Surface_abandon_date = formatDateString(well_plugback.Surface_abandon_date)
	well_plugback.Row_effective_date = formatDateString(well_plugback.Row_effective_date)
	well_plugback.Row_expiry_date = formatDateString(well_plugback.Row_expiry_date)






	rows, err := stmt.Exec(well_plugback.Source, well_plugback.Plugback_obs_no, well_plugback.Active_ind, well_plugback.Base_depth, well_plugback.Base_depth_ouom, well_plugback.Base_strat_unit_id, well_plugback.Cement_amount, well_plugback.Cement_amount_ouom, well_plugback.Cement_amount_uom, well_plugback.Completion_obs_no, well_plugback.Completion_source, well_plugback.Effective_date, well_plugback.Expiry_date, well_plugback.Plugback_ba_id, well_plugback.Plugback_date, well_plugback.Plug_type, well_plugback.Ppdm_guid, well_plugback.Remark, well_plugback.Strat_name_set_id, well_plugback.Surface_abandon_date, well_plugback.Top_depth, well_plugback.Top_depth_ouom, well_plugback.Top_found_depth, well_plugback.Top_found_depth_ouom, well_plugback.Top_strat_unit_id, well_plugback.Row_changed_by, well_plugback.Row_changed_date, well_plugback.Row_created_by, well_plugback.Row_effective_date, well_plugback.Row_expiry_date, well_plugback.Row_quality, well_plugback.Uwi)
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

func PatchWellPlugback(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_plugback set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "plugback_date" || key == "surface_abandon_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellPlugback(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_plugback dto.Well_plugback
	well_plugback.Uwi = id

	stmt, err := db.Prepare("delete from well_plugback where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_plugback.Uwi)
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


