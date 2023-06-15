package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLegalFpsLoc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from legal_fps_loc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Legal_fps_loc

	for rows.Next() {
		var legal_fps_loc dto.Legal_fps_loc
		if err := rows.Scan(&legal_fps_loc.Legal_fps_loc_id, &legal_fps_loc.Location_type, &legal_fps_loc.Source, &legal_fps_loc.Active_ind, &legal_fps_loc.Coord_system_id, &legal_fps_loc.Effective_date, &legal_fps_loc.Expiry_date, &legal_fps_loc.Fps_event_sequence, &legal_fps_loc.Fps_loc_exception, &legal_fps_loc.Grid, &legal_fps_loc.Ppdm_guid, &legal_fps_loc.Reference_latitude, &legal_fps_loc.Reference_longitude, &legal_fps_loc.Remark, &legal_fps_loc.Section, &legal_fps_loc.Unit, &legal_fps_loc.Uwi, &legal_fps_loc.Well_node_id, &legal_fps_loc.Row_changed_by, &legal_fps_loc.Row_changed_date, &legal_fps_loc.Row_created_by, &legal_fps_loc.Row_created_date, &legal_fps_loc.Row_effective_date, &legal_fps_loc.Row_expiry_date, &legal_fps_loc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append legal_fps_loc to result
		result = append(result, legal_fps_loc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLegalFpsLoc(c *fiber.Ctx) error {
	var legal_fps_loc dto.Legal_fps_loc

	setDefaults(&legal_fps_loc)

	if err := json.Unmarshal(c.Body(), &legal_fps_loc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into legal_fps_loc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25)")
	if err != nil {
		return err
	}
	legal_fps_loc.Row_created_date = formatDate(legal_fps_loc.Row_created_date)
	legal_fps_loc.Row_changed_date = formatDate(legal_fps_loc.Row_changed_date)
	legal_fps_loc.Effective_date = formatDateString(legal_fps_loc.Effective_date)
	legal_fps_loc.Expiry_date = formatDateString(legal_fps_loc.Expiry_date)
	legal_fps_loc.Row_effective_date = formatDateString(legal_fps_loc.Row_effective_date)
	legal_fps_loc.Row_expiry_date = formatDateString(legal_fps_loc.Row_expiry_date)






	rows, err := stmt.Exec(legal_fps_loc.Legal_fps_loc_id, legal_fps_loc.Location_type, legal_fps_loc.Source, legal_fps_loc.Active_ind, legal_fps_loc.Coord_system_id, legal_fps_loc.Effective_date, legal_fps_loc.Expiry_date, legal_fps_loc.Fps_event_sequence, legal_fps_loc.Fps_loc_exception, legal_fps_loc.Grid, legal_fps_loc.Ppdm_guid, legal_fps_loc.Reference_latitude, legal_fps_loc.Reference_longitude, legal_fps_loc.Remark, legal_fps_loc.Section, legal_fps_loc.Unit, legal_fps_loc.Uwi, legal_fps_loc.Well_node_id, legal_fps_loc.Row_changed_by, legal_fps_loc.Row_changed_date, legal_fps_loc.Row_created_by, legal_fps_loc.Row_created_date, legal_fps_loc.Row_effective_date, legal_fps_loc.Row_expiry_date, legal_fps_loc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLegalFpsLoc(c *fiber.Ctx) error {
	var legal_fps_loc dto.Legal_fps_loc

	setDefaults(&legal_fps_loc)

	if err := json.Unmarshal(c.Body(), &legal_fps_loc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	legal_fps_loc.Legal_fps_loc_id = id

    if legal_fps_loc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update legal_fps_loc set location_type = :1, source = :2, active_ind = :3, coord_system_id = :4, effective_date = :5, expiry_date = :6, fps_event_sequence = :7, fps_loc_exception = :8, grid = :9, ppdm_guid = :10, reference_latitude = :11, reference_longitude = :12, remark = :13, section = :14, unit = :15, uwi = :16, well_node_id = :17, row_changed_by = :18, row_changed_date = :19, row_created_by = :20, row_effective_date = :21, row_expiry_date = :22, row_quality = :23 where legal_fps_loc_id = :25")
	if err != nil {
		return err
	}

	legal_fps_loc.Row_changed_date = formatDate(legal_fps_loc.Row_changed_date)
	legal_fps_loc.Effective_date = formatDateString(legal_fps_loc.Effective_date)
	legal_fps_loc.Expiry_date = formatDateString(legal_fps_loc.Expiry_date)
	legal_fps_loc.Row_effective_date = formatDateString(legal_fps_loc.Row_effective_date)
	legal_fps_loc.Row_expiry_date = formatDateString(legal_fps_loc.Row_expiry_date)






	rows, err := stmt.Exec(legal_fps_loc.Location_type, legal_fps_loc.Source, legal_fps_loc.Active_ind, legal_fps_loc.Coord_system_id, legal_fps_loc.Effective_date, legal_fps_loc.Expiry_date, legal_fps_loc.Fps_event_sequence, legal_fps_loc.Fps_loc_exception, legal_fps_loc.Grid, legal_fps_loc.Ppdm_guid, legal_fps_loc.Reference_latitude, legal_fps_loc.Reference_longitude, legal_fps_loc.Remark, legal_fps_loc.Section, legal_fps_loc.Unit, legal_fps_loc.Uwi, legal_fps_loc.Well_node_id, legal_fps_loc.Row_changed_by, legal_fps_loc.Row_changed_date, legal_fps_loc.Row_created_by, legal_fps_loc.Row_effective_date, legal_fps_loc.Row_expiry_date, legal_fps_loc.Row_quality, legal_fps_loc.Legal_fps_loc_id)
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

func PatchLegalFpsLoc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update legal_fps_loc set "
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
	query += " where legal_fps_loc_id = :id"

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

func DeleteLegalFpsLoc(c *fiber.Ctx) error {
	id := c.Params("id")
	var legal_fps_loc dto.Legal_fps_loc
	legal_fps_loc.Legal_fps_loc_id = id

	stmt, err := db.Prepare("delete from legal_fps_loc where legal_fps_loc_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(legal_fps_loc.Legal_fps_loc_id)
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


