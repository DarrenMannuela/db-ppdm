package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellHorizDrillPoe(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_horiz_drill_poe")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_horiz_drill_poe

	for rows.Next() {
		var well_horiz_drill_poe dto.Well_horiz_drill_poe
		if err := rows.Scan(&well_horiz_drill_poe.Uwi, &well_horiz_drill_poe.Source, &well_horiz_drill_poe.Point_of_entry_obs_no, &well_horiz_drill_poe.Active_ind, &well_horiz_drill_poe.Effective_date, &well_horiz_drill_poe.Expiry_date, &well_horiz_drill_poe.Lateral_hole_id, &well_horiz_drill_poe.Node_id, &well_horiz_drill_poe.Point_of_entry_md, &well_horiz_drill_poe.Point_of_entry_md_ouom, &well_horiz_drill_poe.Point_of_entry_strat_unit, &well_horiz_drill_poe.Point_of_entry_tvd, &well_horiz_drill_poe.Point_of_entry_tvd_ouom, &well_horiz_drill_poe.Ppdm_guid, &well_horiz_drill_poe.Remark, &well_horiz_drill_poe.Strat_name_set_id, &well_horiz_drill_poe.Row_changed_by, &well_horiz_drill_poe.Row_changed_date, &well_horiz_drill_poe.Row_created_by, &well_horiz_drill_poe.Row_created_date, &well_horiz_drill_poe.Row_effective_date, &well_horiz_drill_poe.Row_expiry_date, &well_horiz_drill_poe.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_horiz_drill_poe to result
		result = append(result, well_horiz_drill_poe)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellHorizDrillPoe(c *fiber.Ctx) error {
	var well_horiz_drill_poe dto.Well_horiz_drill_poe

	setDefaults(&well_horiz_drill_poe)

	if err := json.Unmarshal(c.Body(), &well_horiz_drill_poe); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_horiz_drill_poe values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	well_horiz_drill_poe.Row_created_date = formatDate(well_horiz_drill_poe.Row_created_date)
	well_horiz_drill_poe.Row_changed_date = formatDate(well_horiz_drill_poe.Row_changed_date)
	well_horiz_drill_poe.Effective_date = formatDateString(well_horiz_drill_poe.Effective_date)
	well_horiz_drill_poe.Expiry_date = formatDateString(well_horiz_drill_poe.Expiry_date)
	well_horiz_drill_poe.Row_effective_date = formatDateString(well_horiz_drill_poe.Row_effective_date)
	well_horiz_drill_poe.Row_expiry_date = formatDateString(well_horiz_drill_poe.Row_expiry_date)






	rows, err := stmt.Exec(well_horiz_drill_poe.Uwi, well_horiz_drill_poe.Source, well_horiz_drill_poe.Point_of_entry_obs_no, well_horiz_drill_poe.Active_ind, well_horiz_drill_poe.Effective_date, well_horiz_drill_poe.Expiry_date, well_horiz_drill_poe.Lateral_hole_id, well_horiz_drill_poe.Node_id, well_horiz_drill_poe.Point_of_entry_md, well_horiz_drill_poe.Point_of_entry_md_ouom, well_horiz_drill_poe.Point_of_entry_strat_unit, well_horiz_drill_poe.Point_of_entry_tvd, well_horiz_drill_poe.Point_of_entry_tvd_ouom, well_horiz_drill_poe.Ppdm_guid, well_horiz_drill_poe.Remark, well_horiz_drill_poe.Strat_name_set_id, well_horiz_drill_poe.Row_changed_by, well_horiz_drill_poe.Row_changed_date, well_horiz_drill_poe.Row_created_by, well_horiz_drill_poe.Row_created_date, well_horiz_drill_poe.Row_effective_date, well_horiz_drill_poe.Row_expiry_date, well_horiz_drill_poe.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellHorizDrillPoe(c *fiber.Ctx) error {
	var well_horiz_drill_poe dto.Well_horiz_drill_poe

	setDefaults(&well_horiz_drill_poe)

	if err := json.Unmarshal(c.Body(), &well_horiz_drill_poe); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_horiz_drill_poe.Uwi = id

    if well_horiz_drill_poe.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_horiz_drill_poe set source = :1, point_of_entry_obs_no = :2, active_ind = :3, effective_date = :4, expiry_date = :5, lateral_hole_id = :6, node_id = :7, point_of_entry_md = :8, point_of_entry_md_ouom = :9, point_of_entry_strat_unit = :10, point_of_entry_tvd = :11, point_of_entry_tvd_ouom = :12, ppdm_guid = :13, remark = :14, strat_name_set_id = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where uwi = :23")
	if err != nil {
		return err
	}

	well_horiz_drill_poe.Row_changed_date = formatDate(well_horiz_drill_poe.Row_changed_date)
	well_horiz_drill_poe.Effective_date = formatDateString(well_horiz_drill_poe.Effective_date)
	well_horiz_drill_poe.Expiry_date = formatDateString(well_horiz_drill_poe.Expiry_date)
	well_horiz_drill_poe.Row_effective_date = formatDateString(well_horiz_drill_poe.Row_effective_date)
	well_horiz_drill_poe.Row_expiry_date = formatDateString(well_horiz_drill_poe.Row_expiry_date)






	rows, err := stmt.Exec(well_horiz_drill_poe.Source, well_horiz_drill_poe.Point_of_entry_obs_no, well_horiz_drill_poe.Active_ind, well_horiz_drill_poe.Effective_date, well_horiz_drill_poe.Expiry_date, well_horiz_drill_poe.Lateral_hole_id, well_horiz_drill_poe.Node_id, well_horiz_drill_poe.Point_of_entry_md, well_horiz_drill_poe.Point_of_entry_md_ouom, well_horiz_drill_poe.Point_of_entry_strat_unit, well_horiz_drill_poe.Point_of_entry_tvd, well_horiz_drill_poe.Point_of_entry_tvd_ouom, well_horiz_drill_poe.Ppdm_guid, well_horiz_drill_poe.Remark, well_horiz_drill_poe.Strat_name_set_id, well_horiz_drill_poe.Row_changed_by, well_horiz_drill_poe.Row_changed_date, well_horiz_drill_poe.Row_created_by, well_horiz_drill_poe.Row_effective_date, well_horiz_drill_poe.Row_expiry_date, well_horiz_drill_poe.Row_quality, well_horiz_drill_poe.Uwi)
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

func PatchWellHorizDrillPoe(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_horiz_drill_poe set "
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

func DeleteWellHorizDrillPoe(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_horiz_drill_poe dto.Well_horiz_drill_poe
	well_horiz_drill_poe.Uwi = id

	stmt, err := db.Prepare("delete from well_horiz_drill_poe where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_horiz_drill_poe.Uwi)
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


