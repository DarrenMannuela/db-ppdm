package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLegalGeodeticLoc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from legal_geodetic_loc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Legal_geodetic_loc

	for rows.Next() {
		var legal_geodetic_loc dto.Legal_geodetic_loc
		if err := rows.Scan(&legal_geodetic_loc.Legal_geodetic_loc_id, &legal_geodetic_loc.Location_type, &legal_geodetic_loc.Source, &legal_geodetic_loc.Active_ind, &legal_geodetic_loc.Coord_system_id, &legal_geodetic_loc.Effective_date, &legal_geodetic_loc.Expiry_date, &legal_geodetic_loc.Geodetic_event_sequence, &legal_geodetic_loc.Geodetic_loc_exception, &legal_geodetic_loc.Latitude, &legal_geodetic_loc.Longitude, &legal_geodetic_loc.Node_id, &legal_geodetic_loc.Ppdm_guid, &legal_geodetic_loc.Remark, &legal_geodetic_loc.Uwi, &legal_geodetic_loc.Row_changed_by, &legal_geodetic_loc.Row_changed_date, &legal_geodetic_loc.Row_created_by, &legal_geodetic_loc.Row_created_date, &legal_geodetic_loc.Row_effective_date, &legal_geodetic_loc.Row_expiry_date, &legal_geodetic_loc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append legal_geodetic_loc to result
		result = append(result, legal_geodetic_loc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLegalGeodeticLoc(c *fiber.Ctx) error {
	var legal_geodetic_loc dto.Legal_geodetic_loc

	setDefaults(&legal_geodetic_loc)

	if err := json.Unmarshal(c.Body(), &legal_geodetic_loc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into legal_geodetic_loc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	legal_geodetic_loc.Row_created_date = formatDate(legal_geodetic_loc.Row_created_date)
	legal_geodetic_loc.Row_changed_date = formatDate(legal_geodetic_loc.Row_changed_date)
	legal_geodetic_loc.Effective_date = formatDateString(legal_geodetic_loc.Effective_date)
	legal_geodetic_loc.Expiry_date = formatDateString(legal_geodetic_loc.Expiry_date)
	legal_geodetic_loc.Row_effective_date = formatDateString(legal_geodetic_loc.Row_effective_date)
	legal_geodetic_loc.Row_expiry_date = formatDateString(legal_geodetic_loc.Row_expiry_date)






	rows, err := stmt.Exec(legal_geodetic_loc.Legal_geodetic_loc_id, legal_geodetic_loc.Location_type, legal_geodetic_loc.Source, legal_geodetic_loc.Active_ind, legal_geodetic_loc.Coord_system_id, legal_geodetic_loc.Effective_date, legal_geodetic_loc.Expiry_date, legal_geodetic_loc.Geodetic_event_sequence, legal_geodetic_loc.Geodetic_loc_exception, legal_geodetic_loc.Latitude, legal_geodetic_loc.Longitude, legal_geodetic_loc.Node_id, legal_geodetic_loc.Ppdm_guid, legal_geodetic_loc.Remark, legal_geodetic_loc.Uwi, legal_geodetic_loc.Row_changed_by, legal_geodetic_loc.Row_changed_date, legal_geodetic_loc.Row_created_by, legal_geodetic_loc.Row_created_date, legal_geodetic_loc.Row_effective_date, legal_geodetic_loc.Row_expiry_date, legal_geodetic_loc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLegalGeodeticLoc(c *fiber.Ctx) error {
	var legal_geodetic_loc dto.Legal_geodetic_loc

	setDefaults(&legal_geodetic_loc)

	if err := json.Unmarshal(c.Body(), &legal_geodetic_loc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	legal_geodetic_loc.Legal_geodetic_loc_id = id

    if legal_geodetic_loc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update legal_geodetic_loc set location_type = :1, source = :2, active_ind = :3, coord_system_id = :4, effective_date = :5, expiry_date = :6, geodetic_event_sequence = :7, geodetic_loc_exception = :8, latitude = :9, longitude = :10, node_id = :11, ppdm_guid = :12, remark = :13, uwi = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where legal_geodetic_loc_id = :22")
	if err != nil {
		return err
	}

	legal_geodetic_loc.Row_changed_date = formatDate(legal_geodetic_loc.Row_changed_date)
	legal_geodetic_loc.Effective_date = formatDateString(legal_geodetic_loc.Effective_date)
	legal_geodetic_loc.Expiry_date = formatDateString(legal_geodetic_loc.Expiry_date)
	legal_geodetic_loc.Row_effective_date = formatDateString(legal_geodetic_loc.Row_effective_date)
	legal_geodetic_loc.Row_expiry_date = formatDateString(legal_geodetic_loc.Row_expiry_date)






	rows, err := stmt.Exec(legal_geodetic_loc.Location_type, legal_geodetic_loc.Source, legal_geodetic_loc.Active_ind, legal_geodetic_loc.Coord_system_id, legal_geodetic_loc.Effective_date, legal_geodetic_loc.Expiry_date, legal_geodetic_loc.Geodetic_event_sequence, legal_geodetic_loc.Geodetic_loc_exception, legal_geodetic_loc.Latitude, legal_geodetic_loc.Longitude, legal_geodetic_loc.Node_id, legal_geodetic_loc.Ppdm_guid, legal_geodetic_loc.Remark, legal_geodetic_loc.Uwi, legal_geodetic_loc.Row_changed_by, legal_geodetic_loc.Row_changed_date, legal_geodetic_loc.Row_created_by, legal_geodetic_loc.Row_effective_date, legal_geodetic_loc.Row_expiry_date, legal_geodetic_loc.Row_quality, legal_geodetic_loc.Legal_geodetic_loc_id)
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

func PatchLegalGeodeticLoc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update legal_geodetic_loc set "
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
	query += " where legal_geodetic_loc_id = :id"

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

func DeleteLegalGeodeticLoc(c *fiber.Ctx) error {
	id := c.Params("id")
	var legal_geodetic_loc dto.Legal_geodetic_loc
	legal_geodetic_loc.Legal_geodetic_loc_id = id

	stmt, err := db.Prepare("delete from legal_geodetic_loc where legal_geodetic_loc_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(legal_geodetic_loc.Legal_geodetic_loc_id)
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


