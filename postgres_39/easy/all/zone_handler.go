package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetZone(c *fiber.Ctx) error {
	rows, err := db.Query("select * from zone")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Zone

	for rows.Next() {
		var zone dto.Zone
		if err := rows.Scan(&zone.Zone_id, &zone.Source, &zone.Active_ind, &zone.Effective_date, &zone.Expiry_date, &zone.Ppdm_guid, &zone.Remark, &zone.Zone_name, &zone.Row_changed_by, &zone.Row_changed_date, &zone.Row_created_by, &zone.Row_created_date, &zone.Row_effective_date, &zone.Row_expiry_date, &zone.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append zone to result
		result = append(result, zone)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetZone(c *fiber.Ctx) error {
	var zone dto.Zone

	setDefaults(&zone)

	if err := json.Unmarshal(c.Body(), &zone); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into zone values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15)")
	if err != nil {
		return err
	}
	zone.Row_created_date = formatDate(zone.Row_created_date)
	zone.Row_changed_date = formatDate(zone.Row_changed_date)
	zone.Effective_date = formatDateString(zone.Effective_date)
	zone.Expiry_date = formatDateString(zone.Expiry_date)
	zone.Row_effective_date = formatDateString(zone.Row_effective_date)
	zone.Row_expiry_date = formatDateString(zone.Row_expiry_date)






	rows, err := stmt.Exec(zone.Zone_id, zone.Source, zone.Active_ind, zone.Effective_date, zone.Expiry_date, zone.Ppdm_guid, zone.Remark, zone.Zone_name, zone.Row_changed_by, zone.Row_changed_date, zone.Row_created_by, zone.Row_created_date, zone.Row_effective_date, zone.Row_expiry_date, zone.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateZone(c *fiber.Ctx) error {
	var zone dto.Zone

	setDefaults(&zone)

	if err := json.Unmarshal(c.Body(), &zone); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	zone.Zone_id = id

    if zone.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update zone set source = :1, active_ind = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, remark = :6, zone_name = :7, row_changed_by = :8, row_changed_date = :9, row_created_by = :10, row_effective_date = :11, row_expiry_date = :12, row_quality = :13 where zone_id = :15")
	if err != nil {
		return err
	}

	zone.Row_changed_date = formatDate(zone.Row_changed_date)
	zone.Effective_date = formatDateString(zone.Effective_date)
	zone.Expiry_date = formatDateString(zone.Expiry_date)
	zone.Row_effective_date = formatDateString(zone.Row_effective_date)
	zone.Row_expiry_date = formatDateString(zone.Row_expiry_date)






	rows, err := stmt.Exec(zone.Source, zone.Active_ind, zone.Effective_date, zone.Expiry_date, zone.Ppdm_guid, zone.Remark, zone.Zone_name, zone.Row_changed_by, zone.Row_changed_date, zone.Row_created_by, zone.Row_effective_date, zone.Row_expiry_date, zone.Row_quality, zone.Zone_id)
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

func PatchZone(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update zone set "
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
	query += " where zone_id = :id"

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

func DeleteZone(c *fiber.Ctx) error {
	id := c.Params("id")
	var zone dto.Zone
	zone.Zone_id = id

	stmt, err := db.Prepare("delete from zone where zone_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(zone.Zone_id)
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


