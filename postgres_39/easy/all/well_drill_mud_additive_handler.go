package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDrillMudAdditive(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_drill_mud_additive")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_drill_mud_additive

	for rows.Next() {
		var well_drill_mud_additive dto.Well_drill_mud_additive
		if err := rows.Scan(&well_drill_mud_additive.Uwi, &well_drill_mud_additive.Drill_period_obs_no, &well_drill_mud_additive.Additive_id, &well_drill_mud_additive.Additive_seq_no, &well_drill_mud_additive.Active_ind, &well_drill_mud_additive.Additive_method, &well_drill_mud_additive.Additive_period, &well_drill_mud_additive.Additive_period_ouom, &well_drill_mud_additive.Additive_period_uom, &well_drill_mud_additive.Effective_date, &well_drill_mud_additive.End_time, &well_drill_mud_additive.Expiry_date, &well_drill_mud_additive.Ppdm_guid, &well_drill_mud_additive.Quantity_count, &well_drill_mud_additive.Quantity_value, &well_drill_mud_additive.Quantity_value_ouom, &well_drill_mud_additive.Quantity_value_uom, &well_drill_mud_additive.Remark, &well_drill_mud_additive.Source, &well_drill_mud_additive.Start_time, &well_drill_mud_additive.Timezone, &well_drill_mud_additive.Row_changed_by, &well_drill_mud_additive.Row_changed_date, &well_drill_mud_additive.Row_created_by, &well_drill_mud_additive.Row_created_date, &well_drill_mud_additive.Row_effective_date, &well_drill_mud_additive.Row_expiry_date, &well_drill_mud_additive.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_drill_mud_additive to result
		result = append(result, well_drill_mud_additive)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDrillMudAdditive(c *fiber.Ctx) error {
	var well_drill_mud_additive dto.Well_drill_mud_additive

	setDefaults(&well_drill_mud_additive)

	if err := json.Unmarshal(c.Body(), &well_drill_mud_additive); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_drill_mud_additive values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28)")
	if err != nil {
		return err
	}
	well_drill_mud_additive.Row_created_date = formatDate(well_drill_mud_additive.Row_created_date)
	well_drill_mud_additive.Row_changed_date = formatDate(well_drill_mud_additive.Row_changed_date)
	well_drill_mud_additive.Effective_date = formatDateString(well_drill_mud_additive.Effective_date)
	well_drill_mud_additive.End_time = formatDateString(well_drill_mud_additive.End_time)
	well_drill_mud_additive.Expiry_date = formatDateString(well_drill_mud_additive.Expiry_date)
	well_drill_mud_additive.Start_time = formatDateString(well_drill_mud_additive.Start_time)
	well_drill_mud_additive.Row_effective_date = formatDateString(well_drill_mud_additive.Row_effective_date)
	well_drill_mud_additive.Row_expiry_date = formatDateString(well_drill_mud_additive.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_mud_additive.Uwi, well_drill_mud_additive.Drill_period_obs_no, well_drill_mud_additive.Additive_id, well_drill_mud_additive.Additive_seq_no, well_drill_mud_additive.Active_ind, well_drill_mud_additive.Additive_method, well_drill_mud_additive.Additive_period, well_drill_mud_additive.Additive_period_ouom, well_drill_mud_additive.Additive_period_uom, well_drill_mud_additive.Effective_date, well_drill_mud_additive.End_time, well_drill_mud_additive.Expiry_date, well_drill_mud_additive.Ppdm_guid, well_drill_mud_additive.Quantity_count, well_drill_mud_additive.Quantity_value, well_drill_mud_additive.Quantity_value_ouom, well_drill_mud_additive.Quantity_value_uom, well_drill_mud_additive.Remark, well_drill_mud_additive.Source, well_drill_mud_additive.Start_time, well_drill_mud_additive.Timezone, well_drill_mud_additive.Row_changed_by, well_drill_mud_additive.Row_changed_date, well_drill_mud_additive.Row_created_by, well_drill_mud_additive.Row_created_date, well_drill_mud_additive.Row_effective_date, well_drill_mud_additive.Row_expiry_date, well_drill_mud_additive.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDrillMudAdditive(c *fiber.Ctx) error {
	var well_drill_mud_additive dto.Well_drill_mud_additive

	setDefaults(&well_drill_mud_additive)

	if err := json.Unmarshal(c.Body(), &well_drill_mud_additive); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_drill_mud_additive.Uwi = id

    if well_drill_mud_additive.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_drill_mud_additive set drill_period_obs_no = :1, additive_id = :2, additive_seq_no = :3, active_ind = :4, additive_method = :5, additive_period = :6, additive_period_ouom = :7, additive_period_uom = :8, effective_date = :9, end_time = :10, expiry_date = :11, ppdm_guid = :12, quantity_count = :13, quantity_value = :14, quantity_value_ouom = :15, quantity_value_uom = :16, remark = :17, source = :18, start_time = :19, timezone = :20, row_changed_by = :21, row_changed_date = :22, row_created_by = :23, row_effective_date = :24, row_expiry_date = :25, row_quality = :26 where uwi = :28")
	if err != nil {
		return err
	}

	well_drill_mud_additive.Row_changed_date = formatDate(well_drill_mud_additive.Row_changed_date)
	well_drill_mud_additive.Effective_date = formatDateString(well_drill_mud_additive.Effective_date)
	well_drill_mud_additive.End_time = formatDateString(well_drill_mud_additive.End_time)
	well_drill_mud_additive.Expiry_date = formatDateString(well_drill_mud_additive.Expiry_date)
	well_drill_mud_additive.Start_time = formatDateString(well_drill_mud_additive.Start_time)
	well_drill_mud_additive.Row_effective_date = formatDateString(well_drill_mud_additive.Row_effective_date)
	well_drill_mud_additive.Row_expiry_date = formatDateString(well_drill_mud_additive.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_mud_additive.Drill_period_obs_no, well_drill_mud_additive.Additive_id, well_drill_mud_additive.Additive_seq_no, well_drill_mud_additive.Active_ind, well_drill_mud_additive.Additive_method, well_drill_mud_additive.Additive_period, well_drill_mud_additive.Additive_period_ouom, well_drill_mud_additive.Additive_period_uom, well_drill_mud_additive.Effective_date, well_drill_mud_additive.End_time, well_drill_mud_additive.Expiry_date, well_drill_mud_additive.Ppdm_guid, well_drill_mud_additive.Quantity_count, well_drill_mud_additive.Quantity_value, well_drill_mud_additive.Quantity_value_ouom, well_drill_mud_additive.Quantity_value_uom, well_drill_mud_additive.Remark, well_drill_mud_additive.Source, well_drill_mud_additive.Start_time, well_drill_mud_additive.Timezone, well_drill_mud_additive.Row_changed_by, well_drill_mud_additive.Row_changed_date, well_drill_mud_additive.Row_created_by, well_drill_mud_additive.Row_effective_date, well_drill_mud_additive.Row_expiry_date, well_drill_mud_additive.Row_quality, well_drill_mud_additive.Uwi)
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

func PatchWellDrillMudAdditive(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_drill_mud_additive set "
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
		} else if key == "effective_date" || key == "end_time" || key == "expiry_date" || key == "start_time" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellDrillMudAdditive(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_drill_mud_additive dto.Well_drill_mud_additive
	well_drill_mud_additive.Uwi = id

	stmt, err := db.Prepare("delete from well_drill_mud_additive where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_drill_mud_additive.Uwi)
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


