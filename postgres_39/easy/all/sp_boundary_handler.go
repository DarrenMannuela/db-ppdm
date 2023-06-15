package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpBoundary(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_boundary")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_boundary

	for rows.Next() {
		var sp_boundary dto.Sp_boundary
		if err := rows.Scan(&sp_boundary.Polygon_id, &sp_boundary.Point_seq_no, &sp_boundary.Active_ind, &sp_boundary.Depth, &sp_boundary.Depth_ouom, &sp_boundary.Effective_date, &sp_boundary.Elevation, &sp_boundary.Elevation_ouom, &sp_boundary.Expiry_date, &sp_boundary.Latitude, &sp_boundary.Location_quality, &sp_boundary.Longitude, &sp_boundary.Point_label, &sp_boundary.Point_no, &sp_boundary.Point_position, &sp_boundary.Ppdm_guid, &sp_boundary.Remark, &sp_boundary.Source, &sp_boundary.Row_changed_by, &sp_boundary.Row_changed_date, &sp_boundary.Row_created_by, &sp_boundary.Row_created_date, &sp_boundary.Row_effective_date, &sp_boundary.Row_expiry_date, &sp_boundary.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_boundary to result
		result = append(result, sp_boundary)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpBoundary(c *fiber.Ctx) error {
	var sp_boundary dto.Sp_boundary

	setDefaults(&sp_boundary)

	if err := json.Unmarshal(c.Body(), &sp_boundary); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_boundary values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25)")
	if err != nil {
		return err
	}
	sp_boundary.Row_created_date = formatDate(sp_boundary.Row_created_date)
	sp_boundary.Row_changed_date = formatDate(sp_boundary.Row_changed_date)
	sp_boundary.Effective_date = formatDateString(sp_boundary.Effective_date)
	sp_boundary.Expiry_date = formatDateString(sp_boundary.Expiry_date)
	sp_boundary.Row_effective_date = formatDateString(sp_boundary.Row_effective_date)
	sp_boundary.Row_expiry_date = formatDateString(sp_boundary.Row_expiry_date)






	rows, err := stmt.Exec(sp_boundary.Polygon_id, sp_boundary.Point_seq_no, sp_boundary.Active_ind, sp_boundary.Depth, sp_boundary.Depth_ouom, sp_boundary.Effective_date, sp_boundary.Elevation, sp_boundary.Elevation_ouom, sp_boundary.Expiry_date, sp_boundary.Latitude, sp_boundary.Location_quality, sp_boundary.Longitude, sp_boundary.Point_label, sp_boundary.Point_no, sp_boundary.Point_position, sp_boundary.Ppdm_guid, sp_boundary.Remark, sp_boundary.Source, sp_boundary.Row_changed_by, sp_boundary.Row_changed_date, sp_boundary.Row_created_by, sp_boundary.Row_created_date, sp_boundary.Row_effective_date, sp_boundary.Row_expiry_date, sp_boundary.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpBoundary(c *fiber.Ctx) error {
	var sp_boundary dto.Sp_boundary

	setDefaults(&sp_boundary)

	if err := json.Unmarshal(c.Body(), &sp_boundary); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_boundary.Polygon_id = id

    if sp_boundary.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_boundary set point_seq_no = :1, active_ind = :2, depth = :3, depth_ouom = :4, effective_date = :5, elevation = :6, elevation_ouom = :7, expiry_date = :8, latitude = :9, location_quality = :10, longitude = :11, point_label = :12, point_no = :13, point_position = :14, ppdm_guid = :15, remark = :16, source = :17, row_changed_by = :18, row_changed_date = :19, row_created_by = :20, row_effective_date = :21, row_expiry_date = :22, row_quality = :23 where polygon_id = :25")
	if err != nil {
		return err
	}

	sp_boundary.Row_changed_date = formatDate(sp_boundary.Row_changed_date)
	sp_boundary.Effective_date = formatDateString(sp_boundary.Effective_date)
	sp_boundary.Expiry_date = formatDateString(sp_boundary.Expiry_date)
	sp_boundary.Row_effective_date = formatDateString(sp_boundary.Row_effective_date)
	sp_boundary.Row_expiry_date = formatDateString(sp_boundary.Row_expiry_date)






	rows, err := stmt.Exec(sp_boundary.Point_seq_no, sp_boundary.Active_ind, sp_boundary.Depth, sp_boundary.Depth_ouom, sp_boundary.Effective_date, sp_boundary.Elevation, sp_boundary.Elevation_ouom, sp_boundary.Expiry_date, sp_boundary.Latitude, sp_boundary.Location_quality, sp_boundary.Longitude, sp_boundary.Point_label, sp_boundary.Point_no, sp_boundary.Point_position, sp_boundary.Ppdm_guid, sp_boundary.Remark, sp_boundary.Source, sp_boundary.Row_changed_by, sp_boundary.Row_changed_date, sp_boundary.Row_created_by, sp_boundary.Row_effective_date, sp_boundary.Row_expiry_date, sp_boundary.Row_quality, sp_boundary.Polygon_id)
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

func PatchSpBoundary(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_boundary set "
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
	query += " where polygon_id = :id"

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

func DeleteSpBoundary(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_boundary dto.Sp_boundary
	sp_boundary.Polygon_id = id

	stmt, err := db.Prepare("delete from sp_boundary where polygon_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_boundary.Polygon_id)
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


