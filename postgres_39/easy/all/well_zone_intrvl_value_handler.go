package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellZoneIntrvlValue(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_zone_intrvl_value")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_zone_intrvl_value

	for rows.Next() {
		var well_zone_intrvl_value dto.Well_zone_intrvl_value
		if err := rows.Scan(&well_zone_intrvl_value.Uwi, &well_zone_intrvl_value.Interval_source, &well_zone_intrvl_value.Interval_id, &well_zone_intrvl_value.Zone_id, &well_zone_intrvl_value.Zone_source, &well_zone_intrvl_value.Zone_value_obs_no, &well_zone_intrvl_value.Active_ind, &well_zone_intrvl_value.Date_format_desc, &well_zone_intrvl_value.Effective_date, &well_zone_intrvl_value.Expiry_date, &well_zone_intrvl_value.Parameter, &well_zone_intrvl_value.Ppdm_guid, &well_zone_intrvl_value.Remark, &well_zone_intrvl_value.Source, &well_zone_intrvl_value.Value_type, &well_zone_intrvl_value.Zone_value, &well_zone_intrvl_value.Zone_value_ouom, &well_zone_intrvl_value.Zone_value_uom, &well_zone_intrvl_value.Row_changed_by, &well_zone_intrvl_value.Row_changed_date, &well_zone_intrvl_value.Row_created_by, &well_zone_intrvl_value.Row_created_date, &well_zone_intrvl_value.Row_effective_date, &well_zone_intrvl_value.Row_expiry_date, &well_zone_intrvl_value.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_zone_intrvl_value to result
		result = append(result, well_zone_intrvl_value)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellZoneIntrvlValue(c *fiber.Ctx) error {
	var well_zone_intrvl_value dto.Well_zone_intrvl_value

	setDefaults(&well_zone_intrvl_value)

	if err := json.Unmarshal(c.Body(), &well_zone_intrvl_value); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_zone_intrvl_value values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25)")
	if err != nil {
		return err
	}
	well_zone_intrvl_value.Row_created_date = formatDate(well_zone_intrvl_value.Row_created_date)
	well_zone_intrvl_value.Row_changed_date = formatDate(well_zone_intrvl_value.Row_changed_date)
	well_zone_intrvl_value.Date_format_desc = formatDateString(well_zone_intrvl_value.Date_format_desc)
	well_zone_intrvl_value.Effective_date = formatDateString(well_zone_intrvl_value.Effective_date)
	well_zone_intrvl_value.Expiry_date = formatDateString(well_zone_intrvl_value.Expiry_date)
	well_zone_intrvl_value.Row_effective_date = formatDateString(well_zone_intrvl_value.Row_effective_date)
	well_zone_intrvl_value.Row_expiry_date = formatDateString(well_zone_intrvl_value.Row_expiry_date)






	rows, err := stmt.Exec(well_zone_intrvl_value.Uwi, well_zone_intrvl_value.Interval_source, well_zone_intrvl_value.Interval_id, well_zone_intrvl_value.Zone_id, well_zone_intrvl_value.Zone_source, well_zone_intrvl_value.Zone_value_obs_no, well_zone_intrvl_value.Active_ind, well_zone_intrvl_value.Date_format_desc, well_zone_intrvl_value.Effective_date, well_zone_intrvl_value.Expiry_date, well_zone_intrvl_value.Parameter, well_zone_intrvl_value.Ppdm_guid, well_zone_intrvl_value.Remark, well_zone_intrvl_value.Source, well_zone_intrvl_value.Value_type, well_zone_intrvl_value.Zone_value, well_zone_intrvl_value.Zone_value_ouom, well_zone_intrvl_value.Zone_value_uom, well_zone_intrvl_value.Row_changed_by, well_zone_intrvl_value.Row_changed_date, well_zone_intrvl_value.Row_created_by, well_zone_intrvl_value.Row_created_date, well_zone_intrvl_value.Row_effective_date, well_zone_intrvl_value.Row_expiry_date, well_zone_intrvl_value.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellZoneIntrvlValue(c *fiber.Ctx) error {
	var well_zone_intrvl_value dto.Well_zone_intrvl_value

	setDefaults(&well_zone_intrvl_value)

	if err := json.Unmarshal(c.Body(), &well_zone_intrvl_value); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_zone_intrvl_value.Uwi = id

    if well_zone_intrvl_value.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_zone_intrvl_value set interval_source = :1, interval_id = :2, zone_id = :3, zone_source = :4, zone_value_obs_no = :5, active_ind = :6, date_format_desc = :7, effective_date = :8, expiry_date = :9, parameter = :10, ppdm_guid = :11, remark = :12, source = :13, value_type = :14, zone_value = :15, zone_value_ouom = :16, zone_value_uom = :17, row_changed_by = :18, row_changed_date = :19, row_created_by = :20, row_effective_date = :21, row_expiry_date = :22, row_quality = :23 where uwi = :25")
	if err != nil {
		return err
	}

	well_zone_intrvl_value.Row_changed_date = formatDate(well_zone_intrvl_value.Row_changed_date)
	well_zone_intrvl_value.Date_format_desc = formatDateString(well_zone_intrvl_value.Date_format_desc)
	well_zone_intrvl_value.Effective_date = formatDateString(well_zone_intrvl_value.Effective_date)
	well_zone_intrvl_value.Expiry_date = formatDateString(well_zone_intrvl_value.Expiry_date)
	well_zone_intrvl_value.Row_effective_date = formatDateString(well_zone_intrvl_value.Row_effective_date)
	well_zone_intrvl_value.Row_expiry_date = formatDateString(well_zone_intrvl_value.Row_expiry_date)






	rows, err := stmt.Exec(well_zone_intrvl_value.Interval_source, well_zone_intrvl_value.Interval_id, well_zone_intrvl_value.Zone_id, well_zone_intrvl_value.Zone_source, well_zone_intrvl_value.Zone_value_obs_no, well_zone_intrvl_value.Active_ind, well_zone_intrvl_value.Date_format_desc, well_zone_intrvl_value.Effective_date, well_zone_intrvl_value.Expiry_date, well_zone_intrvl_value.Parameter, well_zone_intrvl_value.Ppdm_guid, well_zone_intrvl_value.Remark, well_zone_intrvl_value.Source, well_zone_intrvl_value.Value_type, well_zone_intrvl_value.Zone_value, well_zone_intrvl_value.Zone_value_ouom, well_zone_intrvl_value.Zone_value_uom, well_zone_intrvl_value.Row_changed_by, well_zone_intrvl_value.Row_changed_date, well_zone_intrvl_value.Row_created_by, well_zone_intrvl_value.Row_effective_date, well_zone_intrvl_value.Row_expiry_date, well_zone_intrvl_value.Row_quality, well_zone_intrvl_value.Uwi)
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

func PatchWellZoneIntrvlValue(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_zone_intrvl_value set "
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
		} else if key == "date_format_desc" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellZoneIntrvlValue(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_zone_intrvl_value dto.Well_zone_intrvl_value
	well_zone_intrvl_value.Uwi = id

	stmt, err := db.Prepare("delete from well_zone_intrvl_value where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_zone_intrvl_value.Uwi)
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


