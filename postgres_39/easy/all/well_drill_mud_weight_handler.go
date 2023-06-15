package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDrillMudWeight(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_drill_mud_weight")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_drill_mud_weight

	for rows.Next() {
		var well_drill_mud_weight dto.Well_drill_mud_weight
		if err := rows.Scan(&well_drill_mud_weight.Uwi, &well_drill_mud_weight.Source, &well_drill_mud_weight.Depth_obs_no, &well_drill_mud_weight.Media_obs_no, &well_drill_mud_weight.Active_ind, &well_drill_mud_weight.Effective_date, &well_drill_mud_weight.End_date, &well_drill_mud_weight.Expiry_date, &well_drill_mud_weight.Mud_depth, &well_drill_mud_weight.Mud_depth_ouom, &well_drill_mud_weight.Mud_weight, &well_drill_mud_weight.Mud_weight_ouom, &well_drill_mud_weight.Mud_weight_uom, &well_drill_mud_weight.Ppdm_guid, &well_drill_mud_weight.Remark, &well_drill_mud_weight.Reported_tvd, &well_drill_mud_weight.Reported_tvd_ouom, &well_drill_mud_weight.Start_date, &well_drill_mud_weight.Row_changed_by, &well_drill_mud_weight.Row_changed_date, &well_drill_mud_weight.Row_created_by, &well_drill_mud_weight.Row_created_date, &well_drill_mud_weight.Row_effective_date, &well_drill_mud_weight.Row_expiry_date, &well_drill_mud_weight.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_drill_mud_weight to result
		result = append(result, well_drill_mud_weight)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDrillMudWeight(c *fiber.Ctx) error {
	var well_drill_mud_weight dto.Well_drill_mud_weight

	setDefaults(&well_drill_mud_weight)

	if err := json.Unmarshal(c.Body(), &well_drill_mud_weight); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_drill_mud_weight values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25)")
	if err != nil {
		return err
	}
	well_drill_mud_weight.Row_created_date = formatDate(well_drill_mud_weight.Row_created_date)
	well_drill_mud_weight.Row_changed_date = formatDate(well_drill_mud_weight.Row_changed_date)
	well_drill_mud_weight.Effective_date = formatDateString(well_drill_mud_weight.Effective_date)
	well_drill_mud_weight.End_date = formatDateString(well_drill_mud_weight.End_date)
	well_drill_mud_weight.Expiry_date = formatDateString(well_drill_mud_weight.Expiry_date)
	well_drill_mud_weight.Start_date = formatDateString(well_drill_mud_weight.Start_date)
	well_drill_mud_weight.Row_effective_date = formatDateString(well_drill_mud_weight.Row_effective_date)
	well_drill_mud_weight.Row_expiry_date = formatDateString(well_drill_mud_weight.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_mud_weight.Uwi, well_drill_mud_weight.Source, well_drill_mud_weight.Depth_obs_no, well_drill_mud_weight.Media_obs_no, well_drill_mud_weight.Active_ind, well_drill_mud_weight.Effective_date, well_drill_mud_weight.End_date, well_drill_mud_weight.Expiry_date, well_drill_mud_weight.Mud_depth, well_drill_mud_weight.Mud_depth_ouom, well_drill_mud_weight.Mud_weight, well_drill_mud_weight.Mud_weight_ouom, well_drill_mud_weight.Mud_weight_uom, well_drill_mud_weight.Ppdm_guid, well_drill_mud_weight.Remark, well_drill_mud_weight.Reported_tvd, well_drill_mud_weight.Reported_tvd_ouom, well_drill_mud_weight.Start_date, well_drill_mud_weight.Row_changed_by, well_drill_mud_weight.Row_changed_date, well_drill_mud_weight.Row_created_by, well_drill_mud_weight.Row_created_date, well_drill_mud_weight.Row_effective_date, well_drill_mud_weight.Row_expiry_date, well_drill_mud_weight.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDrillMudWeight(c *fiber.Ctx) error {
	var well_drill_mud_weight dto.Well_drill_mud_weight

	setDefaults(&well_drill_mud_weight)

	if err := json.Unmarshal(c.Body(), &well_drill_mud_weight); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_drill_mud_weight.Uwi = id

    if well_drill_mud_weight.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_drill_mud_weight set source = :1, depth_obs_no = :2, media_obs_no = :3, active_ind = :4, effective_date = :5, end_date = :6, expiry_date = :7, mud_depth = :8, mud_depth_ouom = :9, mud_weight = :10, mud_weight_ouom = :11, mud_weight_uom = :12, ppdm_guid = :13, remark = :14, reported_tvd = :15, reported_tvd_ouom = :16, start_date = :17, row_changed_by = :18, row_changed_date = :19, row_created_by = :20, row_effective_date = :21, row_expiry_date = :22, row_quality = :23 where uwi = :25")
	if err != nil {
		return err
	}

	well_drill_mud_weight.Row_changed_date = formatDate(well_drill_mud_weight.Row_changed_date)
	well_drill_mud_weight.Effective_date = formatDateString(well_drill_mud_weight.Effective_date)
	well_drill_mud_weight.End_date = formatDateString(well_drill_mud_weight.End_date)
	well_drill_mud_weight.Expiry_date = formatDateString(well_drill_mud_weight.Expiry_date)
	well_drill_mud_weight.Start_date = formatDateString(well_drill_mud_weight.Start_date)
	well_drill_mud_weight.Row_effective_date = formatDateString(well_drill_mud_weight.Row_effective_date)
	well_drill_mud_weight.Row_expiry_date = formatDateString(well_drill_mud_weight.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_mud_weight.Source, well_drill_mud_weight.Depth_obs_no, well_drill_mud_weight.Media_obs_no, well_drill_mud_weight.Active_ind, well_drill_mud_weight.Effective_date, well_drill_mud_weight.End_date, well_drill_mud_weight.Expiry_date, well_drill_mud_weight.Mud_depth, well_drill_mud_weight.Mud_depth_ouom, well_drill_mud_weight.Mud_weight, well_drill_mud_weight.Mud_weight_ouom, well_drill_mud_weight.Mud_weight_uom, well_drill_mud_weight.Ppdm_guid, well_drill_mud_weight.Remark, well_drill_mud_weight.Reported_tvd, well_drill_mud_weight.Reported_tvd_ouom, well_drill_mud_weight.Start_date, well_drill_mud_weight.Row_changed_by, well_drill_mud_weight.Row_changed_date, well_drill_mud_weight.Row_created_by, well_drill_mud_weight.Row_effective_date, well_drill_mud_weight.Row_expiry_date, well_drill_mud_weight.Row_quality, well_drill_mud_weight.Uwi)
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

func PatchWellDrillMudWeight(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_drill_mud_weight set "
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
		} else if key == "effective_date" || key == "end_date" || key == "expiry_date" || key == "start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellDrillMudWeight(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_drill_mud_weight dto.Well_drill_mud_weight
	well_drill_mud_weight.Uwi = id

	stmt, err := db.Prepare("delete from well_drill_mud_weight where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_drill_mud_weight.Uwi)
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


