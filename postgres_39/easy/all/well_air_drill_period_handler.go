package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellAirDrillPeriod(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_air_drill_period")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_air_drill_period

	for rows.Next() {
		var well_air_drill_period dto.Well_air_drill_period
		if err := rows.Scan(&well_air_drill_period.Uwi, &well_air_drill_period.Interval_source, &well_air_drill_period.Air_drill_obs_no, &well_air_drill_period.Depth_obs_no, &well_air_drill_period.Period_obs_no, &well_air_drill_period.Active_ind, &well_air_drill_period.Base_depth, &well_air_drill_period.Base_depth_ouom, &well_air_drill_period.Effective_date, &well_air_drill_period.Expiry_date, &well_air_drill_period.Ppdm_guid, &well_air_drill_period.Remark, &well_air_drill_period.Source, &well_air_drill_period.Top_depth, &well_air_drill_period.Top_depth_ouom, &well_air_drill_period.Row_changed_by, &well_air_drill_period.Row_changed_date, &well_air_drill_period.Row_created_by, &well_air_drill_period.Row_created_date, &well_air_drill_period.Row_effective_date, &well_air_drill_period.Row_expiry_date, &well_air_drill_period.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_air_drill_period to result
		result = append(result, well_air_drill_period)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellAirDrillPeriod(c *fiber.Ctx) error {
	var well_air_drill_period dto.Well_air_drill_period

	setDefaults(&well_air_drill_period)

	if err := json.Unmarshal(c.Body(), &well_air_drill_period); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_air_drill_period values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	well_air_drill_period.Row_created_date = formatDate(well_air_drill_period.Row_created_date)
	well_air_drill_period.Row_changed_date = formatDate(well_air_drill_period.Row_changed_date)
	well_air_drill_period.Effective_date = formatDateString(well_air_drill_period.Effective_date)
	well_air_drill_period.Expiry_date = formatDateString(well_air_drill_period.Expiry_date)
	well_air_drill_period.Row_effective_date = formatDateString(well_air_drill_period.Row_effective_date)
	well_air_drill_period.Row_expiry_date = formatDateString(well_air_drill_period.Row_expiry_date)






	rows, err := stmt.Exec(well_air_drill_period.Uwi, well_air_drill_period.Interval_source, well_air_drill_period.Air_drill_obs_no, well_air_drill_period.Depth_obs_no, well_air_drill_period.Period_obs_no, well_air_drill_period.Active_ind, well_air_drill_period.Base_depth, well_air_drill_period.Base_depth_ouom, well_air_drill_period.Effective_date, well_air_drill_period.Expiry_date, well_air_drill_period.Ppdm_guid, well_air_drill_period.Remark, well_air_drill_period.Source, well_air_drill_period.Top_depth, well_air_drill_period.Top_depth_ouom, well_air_drill_period.Row_changed_by, well_air_drill_period.Row_changed_date, well_air_drill_period.Row_created_by, well_air_drill_period.Row_created_date, well_air_drill_period.Row_effective_date, well_air_drill_period.Row_expiry_date, well_air_drill_period.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellAirDrillPeriod(c *fiber.Ctx) error {
	var well_air_drill_period dto.Well_air_drill_period

	setDefaults(&well_air_drill_period)

	if err := json.Unmarshal(c.Body(), &well_air_drill_period); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_air_drill_period.Uwi = id

    if well_air_drill_period.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_air_drill_period set interval_source = :1, air_drill_obs_no = :2, depth_obs_no = :3, period_obs_no = :4, active_ind = :5, base_depth = :6, base_depth_ouom = :7, effective_date = :8, expiry_date = :9, ppdm_guid = :10, remark = :11, source = :12, top_depth = :13, top_depth_ouom = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where uwi = :22")
	if err != nil {
		return err
	}

	well_air_drill_period.Row_changed_date = formatDate(well_air_drill_period.Row_changed_date)
	well_air_drill_period.Effective_date = formatDateString(well_air_drill_period.Effective_date)
	well_air_drill_period.Expiry_date = formatDateString(well_air_drill_period.Expiry_date)
	well_air_drill_period.Row_effective_date = formatDateString(well_air_drill_period.Row_effective_date)
	well_air_drill_period.Row_expiry_date = formatDateString(well_air_drill_period.Row_expiry_date)






	rows, err := stmt.Exec(well_air_drill_period.Interval_source, well_air_drill_period.Air_drill_obs_no, well_air_drill_period.Depth_obs_no, well_air_drill_period.Period_obs_no, well_air_drill_period.Active_ind, well_air_drill_period.Base_depth, well_air_drill_period.Base_depth_ouom, well_air_drill_period.Effective_date, well_air_drill_period.Expiry_date, well_air_drill_period.Ppdm_guid, well_air_drill_period.Remark, well_air_drill_period.Source, well_air_drill_period.Top_depth, well_air_drill_period.Top_depth_ouom, well_air_drill_period.Row_changed_by, well_air_drill_period.Row_changed_date, well_air_drill_period.Row_created_by, well_air_drill_period.Row_effective_date, well_air_drill_period.Row_expiry_date, well_air_drill_period.Row_quality, well_air_drill_period.Uwi)
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

func PatchWellAirDrillPeriod(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_air_drill_period set "
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

func DeleteWellAirDrillPeriod(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_air_drill_period dto.Well_air_drill_period
	well_air_drill_period.Uwi = id

	stmt, err := db.Prepare("delete from well_air_drill_period where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_air_drill_period.Uwi)
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


