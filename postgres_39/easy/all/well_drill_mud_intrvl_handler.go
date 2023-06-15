package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDrillMudIntrvl(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_drill_mud_intrvl")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_drill_mud_intrvl

	for rows.Next() {
		var well_drill_mud_intrvl dto.Well_drill_mud_intrvl
		if err := rows.Scan(&well_drill_mud_intrvl.Uwi, &well_drill_mud_intrvl.Source, &well_drill_mud_intrvl.Media_type, &well_drill_mud_intrvl.Depth_obs_no, &well_drill_mud_intrvl.Active_ind, &well_drill_mud_intrvl.Casing_depth, &well_drill_mud_intrvl.Casing_depth_ouom, &well_drill_mud_intrvl.Effective_date, &well_drill_mud_intrvl.End_date, &well_drill_mud_intrvl.Expiry_date, &well_drill_mud_intrvl.Max_mud_weight, &well_drill_mud_intrvl.Max_mud_weight_ouom, &well_drill_mud_intrvl.Max_mud_weight_uom, &well_drill_mud_intrvl.Max_weight_depth, &well_drill_mud_intrvl.Max_weight_depth_ouom, &well_drill_mud_intrvl.Mud_end_depth, &well_drill_mud_intrvl.Mud_end_depth_ouom, &well_drill_mud_intrvl.Mud_start_depth, &well_drill_mud_intrvl.Mud_start_depth_ouom, &well_drill_mud_intrvl.Ppdm_guid, &well_drill_mud_intrvl.Remark, &well_drill_mud_intrvl.Start_date, &well_drill_mud_intrvl.Row_changed_by, &well_drill_mud_intrvl.Row_changed_date, &well_drill_mud_intrvl.Row_created_by, &well_drill_mud_intrvl.Row_created_date, &well_drill_mud_intrvl.Row_effective_date, &well_drill_mud_intrvl.Row_expiry_date, &well_drill_mud_intrvl.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_drill_mud_intrvl to result
		result = append(result, well_drill_mud_intrvl)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDrillMudIntrvl(c *fiber.Ctx) error {
	var well_drill_mud_intrvl dto.Well_drill_mud_intrvl

	setDefaults(&well_drill_mud_intrvl)

	if err := json.Unmarshal(c.Body(), &well_drill_mud_intrvl); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_drill_mud_intrvl values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29)")
	if err != nil {
		return err
	}
	well_drill_mud_intrvl.Row_created_date = formatDate(well_drill_mud_intrvl.Row_created_date)
	well_drill_mud_intrvl.Row_changed_date = formatDate(well_drill_mud_intrvl.Row_changed_date)
	well_drill_mud_intrvl.Effective_date = formatDateString(well_drill_mud_intrvl.Effective_date)
	well_drill_mud_intrvl.End_date = formatDateString(well_drill_mud_intrvl.End_date)
	well_drill_mud_intrvl.Expiry_date = formatDateString(well_drill_mud_intrvl.Expiry_date)
	well_drill_mud_intrvl.Start_date = formatDateString(well_drill_mud_intrvl.Start_date)
	well_drill_mud_intrvl.Row_effective_date = formatDateString(well_drill_mud_intrvl.Row_effective_date)
	well_drill_mud_intrvl.Row_expiry_date = formatDateString(well_drill_mud_intrvl.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_mud_intrvl.Uwi, well_drill_mud_intrvl.Source, well_drill_mud_intrvl.Media_type, well_drill_mud_intrvl.Depth_obs_no, well_drill_mud_intrvl.Active_ind, well_drill_mud_intrvl.Casing_depth, well_drill_mud_intrvl.Casing_depth_ouom, well_drill_mud_intrvl.Effective_date, well_drill_mud_intrvl.End_date, well_drill_mud_intrvl.Expiry_date, well_drill_mud_intrvl.Max_mud_weight, well_drill_mud_intrvl.Max_mud_weight_ouom, well_drill_mud_intrvl.Max_mud_weight_uom, well_drill_mud_intrvl.Max_weight_depth, well_drill_mud_intrvl.Max_weight_depth_ouom, well_drill_mud_intrvl.Mud_end_depth, well_drill_mud_intrvl.Mud_end_depth_ouom, well_drill_mud_intrvl.Mud_start_depth, well_drill_mud_intrvl.Mud_start_depth_ouom, well_drill_mud_intrvl.Ppdm_guid, well_drill_mud_intrvl.Remark, well_drill_mud_intrvl.Start_date, well_drill_mud_intrvl.Row_changed_by, well_drill_mud_intrvl.Row_changed_date, well_drill_mud_intrvl.Row_created_by, well_drill_mud_intrvl.Row_created_date, well_drill_mud_intrvl.Row_effective_date, well_drill_mud_intrvl.Row_expiry_date, well_drill_mud_intrvl.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDrillMudIntrvl(c *fiber.Ctx) error {
	var well_drill_mud_intrvl dto.Well_drill_mud_intrvl

	setDefaults(&well_drill_mud_intrvl)

	if err := json.Unmarshal(c.Body(), &well_drill_mud_intrvl); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_drill_mud_intrvl.Uwi = id

    if well_drill_mud_intrvl.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_drill_mud_intrvl set source = :1, media_type = :2, depth_obs_no = :3, active_ind = :4, casing_depth = :5, casing_depth_ouom = :6, effective_date = :7, end_date = :8, expiry_date = :9, max_mud_weight = :10, max_mud_weight_ouom = :11, max_mud_weight_uom = :12, max_weight_depth = :13, max_weight_depth_ouom = :14, mud_end_depth = :15, mud_end_depth_ouom = :16, mud_start_depth = :17, mud_start_depth_ouom = :18, ppdm_guid = :19, remark = :20, start_date = :21, row_changed_by = :22, row_changed_date = :23, row_created_by = :24, row_effective_date = :25, row_expiry_date = :26, row_quality = :27 where uwi = :29")
	if err != nil {
		return err
	}

	well_drill_mud_intrvl.Row_changed_date = formatDate(well_drill_mud_intrvl.Row_changed_date)
	well_drill_mud_intrvl.Effective_date = formatDateString(well_drill_mud_intrvl.Effective_date)
	well_drill_mud_intrvl.End_date = formatDateString(well_drill_mud_intrvl.End_date)
	well_drill_mud_intrvl.Expiry_date = formatDateString(well_drill_mud_intrvl.Expiry_date)
	well_drill_mud_intrvl.Start_date = formatDateString(well_drill_mud_intrvl.Start_date)
	well_drill_mud_intrvl.Row_effective_date = formatDateString(well_drill_mud_intrvl.Row_effective_date)
	well_drill_mud_intrvl.Row_expiry_date = formatDateString(well_drill_mud_intrvl.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_mud_intrvl.Source, well_drill_mud_intrvl.Media_type, well_drill_mud_intrvl.Depth_obs_no, well_drill_mud_intrvl.Active_ind, well_drill_mud_intrvl.Casing_depth, well_drill_mud_intrvl.Casing_depth_ouom, well_drill_mud_intrvl.Effective_date, well_drill_mud_intrvl.End_date, well_drill_mud_intrvl.Expiry_date, well_drill_mud_intrvl.Max_mud_weight, well_drill_mud_intrvl.Max_mud_weight_ouom, well_drill_mud_intrvl.Max_mud_weight_uom, well_drill_mud_intrvl.Max_weight_depth, well_drill_mud_intrvl.Max_weight_depth_ouom, well_drill_mud_intrvl.Mud_end_depth, well_drill_mud_intrvl.Mud_end_depth_ouom, well_drill_mud_intrvl.Mud_start_depth, well_drill_mud_intrvl.Mud_start_depth_ouom, well_drill_mud_intrvl.Ppdm_guid, well_drill_mud_intrvl.Remark, well_drill_mud_intrvl.Start_date, well_drill_mud_intrvl.Row_changed_by, well_drill_mud_intrvl.Row_changed_date, well_drill_mud_intrvl.Row_created_by, well_drill_mud_intrvl.Row_effective_date, well_drill_mud_intrvl.Row_expiry_date, well_drill_mud_intrvl.Row_quality, well_drill_mud_intrvl.Uwi)
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

func PatchWellDrillMudIntrvl(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_drill_mud_intrvl set "
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

func DeleteWellDrillMudIntrvl(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_drill_mud_intrvl dto.Well_drill_mud_intrvl
	well_drill_mud_intrvl.Uwi = id

	stmt, err := db.Prepare("delete from well_drill_mud_intrvl where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_drill_mud_intrvl.Uwi)
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


