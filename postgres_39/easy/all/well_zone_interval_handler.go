package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellZoneInterval(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_zone_interval")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_zone_interval

	for rows.Next() {
		var well_zone_interval dto.Well_zone_interval
		if err := rows.Scan(&well_zone_interval.Uwi, &well_zone_interval.Source, &well_zone_interval.Interval_id, &well_zone_interval.Zone_id, &well_zone_interval.Zone_source, &well_zone_interval.Active_ind, &well_zone_interval.Base_md, &well_zone_interval.Base_md_ouom, &well_zone_interval.Base_tvd, &well_zone_interval.Base_tvd_ouom, &well_zone_interval.Effective_date, &well_zone_interval.Expiry_date, &well_zone_interval.Ppdm_guid, &well_zone_interval.Remark, &well_zone_interval.Top_delta_x, &well_zone_interval.Top_delta_x_ouom, &well_zone_interval.Top_delta_y, &well_zone_interval.Top_delta_y_ouom, &well_zone_interval.Top_md, &well_zone_interval.Top_md_ouom, &well_zone_interval.Top_tvd, &well_zone_interval.Top_tvd_ouom, &well_zone_interval.Row_changed_by, &well_zone_interval.Row_changed_date, &well_zone_interval.Row_created_by, &well_zone_interval.Row_created_date, &well_zone_interval.Row_effective_date, &well_zone_interval.Row_expiry_date, &well_zone_interval.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_zone_interval to result
		result = append(result, well_zone_interval)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellZoneInterval(c *fiber.Ctx) error {
	var well_zone_interval dto.Well_zone_interval

	setDefaults(&well_zone_interval)

	if err := json.Unmarshal(c.Body(), &well_zone_interval); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_zone_interval values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29)")
	if err != nil {
		return err
	}
	well_zone_interval.Row_created_date = formatDate(well_zone_interval.Row_created_date)
	well_zone_interval.Row_changed_date = formatDate(well_zone_interval.Row_changed_date)
	well_zone_interval.Effective_date = formatDateString(well_zone_interval.Effective_date)
	well_zone_interval.Expiry_date = formatDateString(well_zone_interval.Expiry_date)
	well_zone_interval.Row_effective_date = formatDateString(well_zone_interval.Row_effective_date)
	well_zone_interval.Row_expiry_date = formatDateString(well_zone_interval.Row_expiry_date)






	rows, err := stmt.Exec(well_zone_interval.Uwi, well_zone_interval.Source, well_zone_interval.Interval_id, well_zone_interval.Zone_id, well_zone_interval.Zone_source, well_zone_interval.Active_ind, well_zone_interval.Base_md, well_zone_interval.Base_md_ouom, well_zone_interval.Base_tvd, well_zone_interval.Base_tvd_ouom, well_zone_interval.Effective_date, well_zone_interval.Expiry_date, well_zone_interval.Ppdm_guid, well_zone_interval.Remark, well_zone_interval.Top_delta_x, well_zone_interval.Top_delta_x_ouom, well_zone_interval.Top_delta_y, well_zone_interval.Top_delta_y_ouom, well_zone_interval.Top_md, well_zone_interval.Top_md_ouom, well_zone_interval.Top_tvd, well_zone_interval.Top_tvd_ouom, well_zone_interval.Row_changed_by, well_zone_interval.Row_changed_date, well_zone_interval.Row_created_by, well_zone_interval.Row_created_date, well_zone_interval.Row_effective_date, well_zone_interval.Row_expiry_date, well_zone_interval.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellZoneInterval(c *fiber.Ctx) error {
	var well_zone_interval dto.Well_zone_interval

	setDefaults(&well_zone_interval)

	if err := json.Unmarshal(c.Body(), &well_zone_interval); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_zone_interval.Uwi = id

    if well_zone_interval.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_zone_interval set source = :1, interval_id = :2, zone_id = :3, zone_source = :4, active_ind = :5, base_md = :6, base_md_ouom = :7, base_tvd = :8, base_tvd_ouom = :9, effective_date = :10, expiry_date = :11, ppdm_guid = :12, remark = :13, top_delta_x = :14, top_delta_x_ouom = :15, top_delta_y = :16, top_delta_y_ouom = :17, top_md = :18, top_md_ouom = :19, top_tvd = :20, top_tvd_ouom = :21, row_changed_by = :22, row_changed_date = :23, row_created_by = :24, row_effective_date = :25, row_expiry_date = :26, row_quality = :27 where uwi = :29")
	if err != nil {
		return err
	}

	well_zone_interval.Row_changed_date = formatDate(well_zone_interval.Row_changed_date)
	well_zone_interval.Effective_date = formatDateString(well_zone_interval.Effective_date)
	well_zone_interval.Expiry_date = formatDateString(well_zone_interval.Expiry_date)
	well_zone_interval.Row_effective_date = formatDateString(well_zone_interval.Row_effective_date)
	well_zone_interval.Row_expiry_date = formatDateString(well_zone_interval.Row_expiry_date)






	rows, err := stmt.Exec(well_zone_interval.Source, well_zone_interval.Interval_id, well_zone_interval.Zone_id, well_zone_interval.Zone_source, well_zone_interval.Active_ind, well_zone_interval.Base_md, well_zone_interval.Base_md_ouom, well_zone_interval.Base_tvd, well_zone_interval.Base_tvd_ouom, well_zone_interval.Effective_date, well_zone_interval.Expiry_date, well_zone_interval.Ppdm_guid, well_zone_interval.Remark, well_zone_interval.Top_delta_x, well_zone_interval.Top_delta_x_ouom, well_zone_interval.Top_delta_y, well_zone_interval.Top_delta_y_ouom, well_zone_interval.Top_md, well_zone_interval.Top_md_ouom, well_zone_interval.Top_tvd, well_zone_interval.Top_tvd_ouom, well_zone_interval.Row_changed_by, well_zone_interval.Row_changed_date, well_zone_interval.Row_created_by, well_zone_interval.Row_effective_date, well_zone_interval.Row_expiry_date, well_zone_interval.Row_quality, well_zone_interval.Uwi)
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

func PatchWellZoneInterval(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_zone_interval set "
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

func DeleteWellZoneInterval(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_zone_interval dto.Well_zone_interval
	well_zone_interval.Uwi = id

	stmt, err := db.Prepare("delete from well_zone_interval where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_zone_interval.Uwi)
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


