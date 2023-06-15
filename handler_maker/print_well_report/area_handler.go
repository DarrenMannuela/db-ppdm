package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_print_well_report/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetArea(c *fiber.Ctx) error {
	rows, err := db.Query("select * from area")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Area

	for rows.Next() {
		var area dto.Area
		if err := rows.Scan(&area.Area_id, &area.Area_type, &area.Active_ind, &area.Area_max_latitude, &area.Area_max_longitude, &area.Area_min_latitude, &area.Area_min_longitude, &area.Coord_acquisition_id, &area.Coord_system_id, &area.Effective_date, &area.Expiry_date, &area.Local_coord_system_id, &area.Ppdm_guid, &area.Preferred_name, &area.Remark, &area.Source, &area.Source_document_id, &area.Row_changed_by, &area.Row_changed_date, &area.Row_created_by, &area.Row_created_date, &area.Row_effective_date, &area.Row_expiry_date, &area.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append area to result
		result = append(result, area)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetArea(c *fiber.Ctx) error {
	var area dto.Area

	setDefaults(&area)

	if err := json.Unmarshal(c.Body(), &area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into area values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	area.Row_created_date = formatDate(area.Row_created_date)
	area.Row_changed_date = formatDate(area.Row_changed_date)
	area.Effective_date = formatDateString(area.Effective_date)
	area.Expiry_date = formatDateString(area.Expiry_date)
	area.Row_effective_date = formatDateString(area.Row_effective_date)
	area.Row_expiry_date = formatDateString(area.Row_expiry_date)






	rows, err := stmt.Exec(area.Area_id, area.Area_type, area.Active_ind, area.Area_max_latitude, area.Area_max_longitude, area.Area_min_latitude, area.Area_min_longitude, area.Coord_acquisition_id, area.Coord_system_id, area.Effective_date, area.Expiry_date, area.Local_coord_system_id, area.Ppdm_guid, area.Preferred_name, area.Remark, area.Source, area.Source_document_id, area.Row_changed_by, area.Row_changed_date, area.Row_created_by, area.Row_created_date, area.Row_effective_date, area.Row_expiry_date, area.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateArea(c *fiber.Ctx) error {
	var area dto.Area

	setDefaults(&area)

	if err := json.Unmarshal(c.Body(), &area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	area.Area_id = id

    if area.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update area set area_type = :1, active_ind = :2, area_max_latitude = :3, area_max_longitude = :4, area_min_latitude = :5, area_min_longitude = :6, coord_acquisition_id = :7, coord_system_id = :8, effective_date = :9, expiry_date = :10, local_coord_system_id = :11, ppdm_guid = :12, preferred_name = :13, remark = :14, source = :15, source_document_id = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where area_id = :24")
	if err != nil {
		return err
	}

	area.Row_changed_date = formatDate(area.Row_changed_date)
	area.Effective_date = formatDateString(area.Effective_date)
	area.Expiry_date = formatDateString(area.Expiry_date)
	area.Row_effective_date = formatDateString(area.Row_effective_date)
	area.Row_expiry_date = formatDateString(area.Row_expiry_date)






	rows, err := stmt.Exec(area.Area_type, area.Active_ind, area.Area_max_latitude, area.Area_max_longitude, area.Area_min_latitude, area.Area_min_longitude, area.Coord_acquisition_id, area.Coord_system_id, area.Effective_date, area.Expiry_date, area.Local_coord_system_id, area.Ppdm_guid, area.Preferred_name, area.Remark, area.Source, area.Source_document_id, area.Row_changed_by, area.Row_changed_date, area.Row_created_by, area.Row_effective_date, area.Row_expiry_date, area.Row_quality, area.Area_id)
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

func PatchArea(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update area set "
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
				formattedDate := formatDate(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where area_id = :id"

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

func DeleteArea(c *fiber.Ctx) error {
	id := c.Params("id")
	var area dto.Area
	area.Area_id = id

	stmt, err := db.Prepare("delete from area where area_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(area.Area_id)
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


